package net

import (
	"context"
	"encoding/binary"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"hk4e/common/config"
	"hk4e/common/mq"
	"hk4e/common/region"
	"hk4e/common/rpc"
	"hk4e/gate/client_proto"
	"hk4e/gate/kcp"
	"hk4e/node/api"
	"hk4e/pkg/logger"
	"hk4e/pkg/random"
	"hk4e/protocol/cmd"
)

const (
	ConnEstFreqLimit      = 100        // 每秒连接建立频率限制
	RecvPacketFreqLimit   = 1000       // 客户端上行每秒发包频率限制
	SendPacketFreqLimit   = 1000       // 服务器下行每秒发包频率限制
	PacketMaxLen          = 343 * 1024 // 最大应用层包长度
	ConnRecvTimeout       = 30         // 收包超时时间 秒
	ConnSendTimeout       = 10         // 发包超时时间 秒
	MaxClientConnNumLimit = 1000       // 最大客户端连接数限制
)

var CLIENT_CONN_NUM int32 = 0 // 当前客户端连接数

type KcpConnectManager struct {
	discovery *rpc.DiscoveryClient // node服务器客户端
	// 会话
	sessionMap            map[uint32]*Session
	sessionUserIdMap      map[uint32]*Session
	sessionMapLock        sync.RWMutex
	createSessionChan     chan *Session
	destroySessionChan    chan *Session
	globalGsOnlineMap     map[uint32]string
	globalGsOnlineMapLock sync.RWMutex
	// 连接事件
	kcpEventInput            chan *KcpEvent
	kcpEventOutput           chan *KcpEvent
	reLoginRemoteKickRegChan chan *RemoteKick
	// 协议
	serverCmdProtoMap *cmd.CmdProtoMap
	clientCmdProtoMap *client_proto.ClientCmdProtoMap
	// 输入输出管道
	messageQueue *mq.MessageQueue
	// 密钥
	dispatchKey  []byte
	signRsaKey   []byte
	encRsaKeyMap map[string][]byte
}

func NewKcpConnectManager(messageQueue *mq.MessageQueue, discovery *rpc.DiscoveryClient) (r *KcpConnectManager) {
	r = new(KcpConnectManager)
	r.discovery = discovery
	r.sessionMap = make(map[uint32]*Session)
	r.sessionUserIdMap = make(map[uint32]*Session)
	r.createSessionChan = make(chan *Session, 1000)
	r.destroySessionChan = make(chan *Session, 1000)
	r.globalGsOnlineMap = make(map[uint32]string)
	r.kcpEventInput = make(chan *KcpEvent, 1000)
	r.kcpEventOutput = make(chan *KcpEvent, 1000)
	r.reLoginRemoteKickRegChan = make(chan *RemoteKick, 1000)
	r.serverCmdProtoMap = cmd.NewCmdProtoMap()
	if config.GetConfig().Hk4e.ClientProtoProxyEnable {
		r.clientCmdProtoMap = client_proto.NewClientCmdProtoMap()
	}
	r.messageQueue = messageQueue
	r.run()
	return r
}

func (k *KcpConnectManager) run() {
	// 读取密钥相关文件
	k.signRsaKey, k.encRsaKeyMap, _ = region.LoadRegionRsaKey()
	// key
	rsp, err := k.discovery.GetRegionEc2B(context.TODO(), &api.NullMsg{})
	if err != nil {
		logger.Error("get region ec2b error: %v", err)
		return
	}
	ec2b, err := random.LoadEc2bKey(rsp.Data)
	if err != nil {
		logger.Error("parse region ec2b error: %v", err)
		return
	}
	regionEc2b := random.NewEc2b()
	regionEc2b.SetSeed(ec2b.Seed())
	k.dispatchKey = regionEc2b.XorKey()
	// kcp
	port := strconv.Itoa(int(config.GetConfig().Hk4e.KcpPort))
	listener, err := kcp.ListenWithOptions("0.0.0.0:" + port)
	if err != nil {
		logger.Error("listen kcp err: %v", err)
		return
	}
	go k.enetHandle(listener)
	go k.eventHandle()
	if !config.GetConfig().Hk4e.ForwardModeEnable {
		go k.forwardServerMsgToClientHandle()
	}
	go k.acceptHandle(listener)
	go k.gateNetInfo()
	k.syncGlobalGsOnlineMap()
	go k.autoSyncGlobalGsOnlineMap()
}

func (k *KcpConnectManager) Close() {
	k.closeAllKcpConn()
}

func (k *KcpConnectManager) gateNetInfo() {
	ticker := time.NewTicker(time.Second * 60)
	kcpErrorCount := uint64(0)
	for {
		<-ticker.C
		snmp := kcp.DefaultSnmp.Copy()
		kcpErrorCount += snmp.KCPInErrors
		logger.Info("kcp send: %v B/s, kcp recv: %v B/s", snmp.BytesSent/60, snmp.BytesReceived/60)
		logger.Info("udp send: %v B/s, udp recv: %v B/s", snmp.OutBytes/60, snmp.InBytes/60)
		logger.Info("udp send: %v pps, udp recv: %v pps", snmp.OutPkts/60, snmp.InPkts/60)
		clientConnNum := atomic.LoadInt32(&CLIENT_CONN_NUM)
		logger.Info("conn num: %v, new conn num: %v, kcp error num: %v", clientConnNum, snmp.CurrEstab, kcpErrorCount)
		kcp.DefaultSnmp.Reset()
	}
}

// 接收并创建新连接处理函数
func (k *KcpConnectManager) acceptHandle(listener *kcp.Listener) {
	logger.Info("accept handle start")
	connEstFreqLimitCounter := 0
	connEstFreqLimitTimer := time.Now().UnixNano()
	for {
		conn, err := listener.AcceptKCP()
		if err != nil {
			logger.Error("accept kcp err: %v", err)
			return
		}
		// 连接建立频率限制
		connEstFreqLimitCounter++
		if connEstFreqLimitCounter > ConnEstFreqLimit {
			now := time.Now().UnixNano()
			if now-connEstFreqLimitTimer > int64(time.Second) {
				connEstFreqLimitCounter = 0
				connEstFreqLimitTimer = now
			} else {
				logger.Error("conn est freq limit, now: %v conn/s", connEstFreqLimitCounter)
				_ = conn.Close()
				continue
			}
		}
		if config.GetConfig().Hk4e.ForwardModeEnable {
			clientConnNum := atomic.LoadInt32(&CLIENT_CONN_NUM)
			if clientConnNum != 0 {
				logger.Error("forward mode only support one client conn now")
				_ = conn.Close()
				continue
			}
		}
		if k.GetSession(conn.GetSessionId()) != nil {
			logger.Error("session already exist, sessionId: %v", conn.GetSessionId())
			_ = conn.Close()
			continue
		}
		conn.SetACKNoDelay(true)
		conn.SetWriteDelay(false)
		conn.SetWindowSize(256, 256)
		logger.Info("client connect, conv: %v, sessionId: %v", conn.GetConv(), conn.GetSessionId())
		kcpRawSendChan := make(chan *ProtoMsg, 1000)
		session := &Session{
			sessionId:              conn.GetSessionId(),
			conn:                   conn,
			connState:              ConnEst,
			userId:                 0,
			kcpRawSendChan:         kcpRawSendChan,
			seed:                   0,
			xorKey:                 k.dispatchKey,
			changeXorKeyFin:        false,
			gsServerAppId:          "",
			anticheatServerAppId:   "",
			pathfindingServerAppId: "",
			robotServerAppId:       "",
			useMagicSeed:           false,
			keyId:                  0,
			clientRandKey:          "",
		}
		if config.GetConfig().Hk4e.ForwardModeEnable {
			robotServerAppId, err := k.discovery.GetServerAppId(context.TODO(), &api.GetServerAppIdReq{
				ServerType: api.ROBOT,
			})
			if err != nil {
				logger.Error("get robot server appid error: %v", err)
				_ = conn.Close()
				continue
			}
			session.robotServerAppId = robotServerAppId.AppId
			k.messageQueue.SendToRobot(session.robotServerAppId, &mq.NetMsg{
				MsgType: mq.MsgTypeServer,
				EventId: mq.ServerForwardModeClientConnNotify,
			})
		}
		go k.recvHandle(session)
		go k.sendHandle(session)
		if config.GetConfig().Hk4e.ForwardModeEnable {
			go k.forwardRobotMsgToClientHandle(session)
		}
		// 连接建立成功通知
		k.kcpEventOutput <- &KcpEvent{
			Conv:         conn.GetConv(),
			SessionId:    conn.GetSessionId(),
			EventId:      KcpConnEstNotify,
			EventMessage: conn.RemoteAddr().String(),
		}
		atomic.AddInt32(&CLIENT_CONN_NUM, 1)
	}
}

// 连接事件处理函数
func (k *KcpConnectManager) enetHandle(listener *kcp.Listener) {
	logger.Info("enet handle start")
	sessionIdCounter := uint32(0)
	for {
		enetNotify := <-listener.EnetNotify
		logger.Info("[Enet Notify], addr: %v, conv: %v, sessionId: %v, connType: %v, enetType: %v",
			enetNotify.Addr, enetNotify.Conv, enetNotify.SessionId, enetNotify.ConnType, enetNotify.EnetType)
		switch enetNotify.ConnType {
		case kcp.ConnEnetSyn:
			if enetNotify.EnetType != kcp.EnetClientConnectKey {
				logger.Error("enet type not match, sessionId: %v", enetNotify.SessionId)
				continue
			}
			sessionIdCounter++
			listener.SendEnetNotifyToPeer(&kcp.Enet{
				Addr:      enetNotify.Addr,
				Conv:      binary.BigEndian.Uint32(random.GetRandomByte(4)),
				SessionId: sessionIdCounter,
				ConnType:  kcp.ConnEnetEst,
				EnetType:  enetNotify.EnetType,
			})
		case kcp.ConnEnetEst:
		case kcp.ConnEnetFin:
			session := k.GetSession(enetNotify.SessionId)
			if session == nil {
				logger.Error("session not exist, sessionId: %v", enetNotify.SessionId)
				continue
			}
			if session.conn.GetConv() != enetNotify.Conv {
				logger.Error("conv not match, sessionId: %v", enetNotify.SessionId)
				continue
			}
			session.conn.SendEnetNotifyToPeer(&kcp.Enet{
				ConnType: kcp.ConnEnetFin,
				EnetType: enetNotify.EnetType,
			})
			_ = session.conn.Close()
		case kcp.ConnEnetAddrChange:
			// 连接地址改变通知
			k.kcpEventOutput <- &KcpEvent{
				Conv:         enetNotify.Conv,
				SessionId:    enetNotify.SessionId,
				EventId:      KcpConnAddrChangeNotify,
				EventMessage: enetNotify.Addr,
			}
		default:
		}
	}
}

// Session 连接会话结构 只允许定义并发安全或者简单的基础数据结构
type Session struct {
	sessionId              uint32
	conn                   *kcp.UDPSession
	connState              uint8
	userId                 uint32
	kcpRawSendChan         chan *ProtoMsg
	seed                   uint64
	xorKey                 []byte
	changeXorKeyFin        bool
	gsServerAppId          string
	anticheatServerAppId   string
	pathfindingServerAppId string
	robotServerAppId       string
	useMagicSeed           bool
	keyId                  uint32
	clientRandKey          string
}

// 接收
func (k *KcpConnectManager) recvHandle(session *Session) {
	logger.Info("recv handle start")
	conn := session.conn
	sessionId := conn.GetSessionId()
	recvBuf := make([]byte, PacketMaxLen)
	pktFreqLimitCounter := 0
	pktFreqLimitTimer := time.Now().UnixNano()
	for {
		_ = conn.SetReadDeadline(time.Now().Add(time.Second * ConnRecvTimeout))
		recvLen, err := conn.Read(recvBuf)
		if err != nil {
			logger.Error("exit recv loop, conn read err: %v, sessionId: %v", err, sessionId)
			k.closeKcpConn(session, kcp.EnetServerKick)
			break
		}
		// 收包频率限制
		pktFreqLimitCounter++
		if pktFreqLimitCounter > RecvPacketFreqLimit {
			now := time.Now().UnixNano()
			if now-pktFreqLimitTimer > int64(time.Second) {
				pktFreqLimitCounter = 0
				pktFreqLimitTimer = now
			} else {
				logger.Error("exit recv loop, client packet send freq too high, sessionId: %v, pps: %v",
					sessionId, pktFreqLimitCounter)
				k.closeKcpConn(session, kcp.EnetPacketFreqTooHigh)
				break
			}
		}
		recvData := recvBuf[:recvLen]
		kcpMsgList := make([]*KcpMsg, 0)
		DecodeBinToPayload(recvData, sessionId, &kcpMsgList, session.xorKey)
		for _, v := range kcpMsgList {
			protoMsgList := ProtoDecode(v, k.serverCmdProtoMap, k.clientCmdProtoMap)
			for _, vv := range protoMsgList {
				if config.GetConfig().Hk4e.ForwardModeEnable {
					k.forwardClientMsgToRobotHandle(vv, session)
				} else {
					k.forwardClientMsgToServerHandle(vv, session)
				}
			}
		}
	}
}

// 发送
func (k *KcpConnectManager) sendHandle(session *Session) {
	logger.Info("send handle start")
	conn := session.conn
	sessionId := conn.GetSessionId()
	pktFreqLimitCounter := 0
	pktFreqLimitTimer := time.Now().UnixNano()
	for {
		protoMsg, ok := <-session.kcpRawSendChan
		if !ok {
			logger.Error("exit send loop, send chan close, sessionId: %v", sessionId)
			k.closeKcpConn(session, kcp.EnetServerKick)
			break
		}
		kcpMsg := ProtoEncode(protoMsg, k.serverCmdProtoMap, k.clientCmdProtoMap)
		if kcpMsg == nil {
			logger.Error("decode kcp msg is nil, sessionId: %v", sessionId)
			continue
		}
		bin := EncodePayloadToBin(kcpMsg, session.xorKey)
		_ = conn.SetWriteDeadline(time.Now().Add(time.Second * ConnSendTimeout))
		_, err := conn.Write(bin)
		if err != nil {
			logger.Error("exit send loop, conn write err: %v, sessionId: %v", err, sessionId)
			k.closeKcpConn(session, kcp.EnetServerKick)
			break
		}
		// 发包频率限制
		pktFreqLimitCounter++
		if pktFreqLimitCounter > SendPacketFreqLimit {
			now := time.Now().UnixNano()
			if now-pktFreqLimitTimer > int64(time.Second) {
				pktFreqLimitCounter = 0
				pktFreqLimitTimer = now
			} else {
				logger.Error("exit send loop, server packet send freq too high, sessionId: %v, pps: %v",
					sessionId, pktFreqLimitCounter)
				k.closeKcpConn(session, kcp.EnetPacketFreqTooHigh)
				break
			}
		}
		if session.changeXorKeyFin == false && protoMsg.CmdId == cmd.GetPlayerTokenRsp {
			// XOR密钥切换
			logger.Info("change session xor key, sessionId: %v", sessionId)
			session.changeXorKeyFin = true
			keyBlock := random.NewKeyBlock(session.seed, session.useMagicSeed)
			xorKey := keyBlock.XorKey()
			key := make([]byte, 4096)
			copy(key, xorKey[:])
			session.xorKey = key
		}
	}
}

// 强制关闭指定连接
func (k *KcpConnectManager) forceCloseKcpConn(sessionId uint32, reason uint32) {
	session := k.GetSession(sessionId)
	if session == nil {
		logger.Error("session not exist, sessionId: %v", sessionId)
		return
	}
	k.closeKcpConn(session, reason)
	logger.Info("conn has been force close, sessionId: %v", sessionId)
}

// 关闭指定连接
func (k *KcpConnectManager) closeKcpConn(session *Session, enetType uint32) {
	if session.connState == ConnClose {
		return
	}
	session.connState = ConnClose
	conn := session.conn
	// 清理数据
	k.DeleteSession(session.sessionId, session.userId)
	// 关闭连接
	err := conn.Close()
	if err == nil {
		conn.SendEnetNotifyToPeer(&kcp.Enet{
			ConnType: kcp.ConnEnetFin,
			EnetType: enetType,
		})
	}
	// 连接关闭通知
	k.kcpEventOutput <- &KcpEvent{
		Conv:      conn.GetConv(),
		SessionId: conn.GetSessionId(),
		EventId:   KcpConnCloseNotify,
	}
	if !config.GetConfig().Hk4e.ForwardModeEnable {
		// 通知GS玩家下线
		connCtrlMsg := new(mq.ConnCtrlMsg)
		connCtrlMsg.UserId = session.userId
		k.messageQueue.SendToGs(session.gsServerAppId, &mq.NetMsg{
			MsgType:     mq.MsgTypeConnCtrl,
			EventId:     mq.UserOfflineNotify,
			ConnCtrlMsg: connCtrlMsg,
		})
		logger.Info("send to gs user offline, SessionId: %v, UserId: %v", conn.GetSessionId(), connCtrlMsg.UserId)
		k.destroySessionChan <- session
	} else {
		k.messageQueue.SendToRobot(session.robotServerAppId, &mq.NetMsg{
			MsgType: mq.MsgTypeServer,
			EventId: mq.ServerForwardModeClientCloseNotify,
		})
	}
	atomic.AddInt32(&CLIENT_CONN_NUM, -1)
}

// 关闭所有连接
func (k *KcpConnectManager) closeAllKcpConn() {
	sessionList := make([]*Session, 0)
	k.sessionMapLock.RLock()
	for _, session := range k.sessionMap {
		sessionList = append(sessionList, session)
	}
	k.sessionMapLock.RUnlock()
	for _, session := range sessionList {
		k.closeKcpConn(session, kcp.EnetServerShutdown)
	}
}

func (k *KcpConnectManager) GetSession(sessionId uint32) *Session {
	k.sessionMapLock.RLock()
	session, _ := k.sessionMap[sessionId]
	k.sessionMapLock.RUnlock()
	return session
}

func (k *KcpConnectManager) GetSessionByUserId(userId uint32) *Session {
	k.sessionMapLock.RLock()
	session, _ := k.sessionUserIdMap[userId]
	k.sessionMapLock.RUnlock()
	return session
}

func (k *KcpConnectManager) SetSession(session *Session, sessionId uint32, userId uint32) {
	k.sessionMapLock.Lock()
	k.sessionMap[sessionId] = session
	k.sessionUserIdMap[userId] = session
	k.sessionMapLock.Unlock()
}

func (k *KcpConnectManager) DeleteSession(sessionId uint32, userId uint32) {
	k.sessionMapLock.Lock()
	delete(k.sessionMap, sessionId)
	delete(k.sessionUserIdMap, userId)
	k.sessionMapLock.Unlock()
}

func (k *KcpConnectManager) autoSyncGlobalGsOnlineMap() {
	ticker := time.NewTicker(time.Second * 60)
	for {
		<-ticker.C
		k.syncGlobalGsOnlineMap()
	}
}

func (k *KcpConnectManager) syncGlobalGsOnlineMap() {
	rsp, err := k.discovery.GetGlobalGsOnlineMap(context.TODO(), nil)
	if err != nil {
		logger.Error("get global gs online map error: %v", err)
		return
	}
	copyMap := make(map[uint32]string)
	for k, v := range rsp.OnlineMap {
		copyMap[k] = v
	}
	copyMapLen := len(copyMap)
	k.globalGsOnlineMapLock.Lock()
	k.globalGsOnlineMap = copyMap
	k.globalGsOnlineMapLock.Unlock()
	logger.Info("sync global gs online map finish, len: %v", copyMapLen)
}
