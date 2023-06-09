package net

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"hk4e/common/config"
	"hk4e/common/mq"
	"hk4e/dispatch/controller"
	"hk4e/gate/kcp"
	"hk4e/node/api"
	"hk4e/pkg/endec"
	"hk4e/pkg/httpclient"
	"hk4e/pkg/logger"
	"hk4e/pkg/random"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

const (
	ConnEst = iota
	ConnWaitLogin
	ConnActive
	ConnClose
)

// 转发客户端消息到其他服务器 每个连接独立协程
func (k *KcpConnectManager) forwardClientMsgToServerHandle(protoMsg *ProtoMsg, session *Session) {
	if protoMsg.HeadMessage == nil {
		logger.Error("recv null head msg: %v", protoMsg)
		return
	}
	// 网关服务器本地处理的请求
	switch protoMsg.CmdId {
	case cmd.GetPlayerTokenReq:
		// GATE登录包
		if session.connState != ConnEst {
			return
		}
		session.connState = ConnWaitLogin
		req := protoMsg.PayloadMessage.(*proto.GetPlayerTokenReq)
		rsp := k.doGateLogin(req, session)
		if rsp == nil {
			return
		}
		// 返回数据到客户端
		msg := &ProtoMsg{
			ConvId:         protoMsg.ConvId,
			CmdId:          cmd.GetPlayerTokenRsp,
			HeadMessage:    k.getHeadMsg(protoMsg.HeadMessage.ClientSequenceId),
			PayloadMessage: rsp,
		}
		session.kcpRawSendChan <- msg
	case cmd.PlayerForceExitReq:
		// 退出游戏
		if session.connState != ConnActive {
			return
		}
		k.kcpEventInput <- &KcpEvent{
			ConvId:       protoMsg.ConvId,
			EventId:      KcpConnForceClose,
			EventMessage: uint32(kcp.EnetClientClose),
		}
	case cmd.PingReq:
		// ping
		pingReq := protoMsg.PayloadMessage.(*proto.PingReq)
		logger.Debug("user ping req, data: %v", pingReq.String())
		// 返回数据到客户端
		pingRsp := new(proto.PingRsp)
		pingRsp.ClientTime = pingReq.ClientTime
		msg := &ProtoMsg{
			ConvId:         protoMsg.ConvId,
			CmdId:          cmd.PingRsp,
			HeadMessage:    k.getHeadMsg(protoMsg.HeadMessage.ClientSequenceId),
			PayloadMessage: pingRsp,
		}
		session.kcpRawSendChan <- msg
		logger.Debug("convId: %v, RTO: %v, SRTT: %v, RTTVar: %v",
			protoMsg.ConvId, session.conn.GetRTO(), session.conn.GetSRTT(), session.conn.GetSRTTVar())
		if session.connState != ConnActive {
			return
		}
		// 通知GS玩家客户端往返时延
		rtt := session.conn.GetSRTT()
		connCtrlMsg := &mq.ConnCtrlMsg{
			UserId:     session.userId,
			ClientRtt:  uint32(rtt),
			ClientTime: 0,
		}
		k.messageQueue.SendToGs(session.gsServerAppId, &mq.NetMsg{
			MsgType:     mq.MsgTypeConnCtrl,
			EventId:     mq.ClientRttNotify,
			ConnCtrlMsg: connCtrlMsg,
		})
		// 通知GS玩家客户端的本地时钟
		connCtrlMsg = &mq.ConnCtrlMsg{
			UserId:     session.userId,
			ClientTime: pingReq.ClientTime,
		}
		k.messageQueue.SendToGs(session.gsServerAppId, &mq.NetMsg{
			MsgType:     mq.MsgTypeConnCtrl,
			EventId:     mq.ClientTimeNotify,
			ConnCtrlMsg: connCtrlMsg,
		})
	case cmd.PlayerLoginReq:
		// GS登录包
		if session.connState != ConnWaitLogin {
			return
		}
		playerLoginReq := protoMsg.PayloadMessage.(*proto.PlayerLoginReq)
		playerLoginReq.TargetUid = 0
		playerLoginReq.TargetHomeOwnerUid = 0
		gameMsg := &mq.GameMsg{
			UserId:         session.userId,
			CmdId:          protoMsg.CmdId,
			ClientSeq:      protoMsg.HeadMessage.ClientSequenceId,
			PayloadMessage: playerLoginReq,
		}
		// 转发到GS
		k.messageQueue.SendToGs(session.gsServerAppId, &mq.NetMsg{
			MsgType: mq.MsgTypeGame,
			EventId: mq.NormalMsg,
			GameMsg: gameMsg,
		})
	default:
		if session.connState != ConnActive {
			logger.Error("conn not active so drop packet, cmdId: %v, userId: %v, convId: %v", protoMsg.CmdId, session.userId, protoMsg.ConvId)
			return
		}
		gameMsg := &mq.GameMsg{
			UserId:             session.userId,
			CmdId:              protoMsg.CmdId,
			ClientSeq:          protoMsg.HeadMessage.ClientSequenceId,
			PayloadMessageData: nil,
		}
		// 在这里直接序列化成二进制数据 终结PayloadMessage的生命周期并回收进缓存池
		payloadMessageData, err := pb.Marshal(protoMsg.PayloadMessage)
		if err != nil {
			logger.Error("parse payload msg to bin error: %v, stack: %v", err, logger.Stack())
			return
		}
		k.serverCmdProtoMap.PutProtoObjCache(protoMsg.CmdId, protoMsg.PayloadMessage)
		gameMsg.PayloadMessageData = payloadMessageData
		// 转发到寻路服务器
		if session.pathfindingServerAppId != "" {
			if protoMsg.CmdId == cmd.QueryPathReq ||
				protoMsg.CmdId == cmd.ObstacleModifyNotify {
				k.messageQueue.SendToPathfinding(session.pathfindingServerAppId, &mq.NetMsg{
					MsgType: mq.MsgTypeGame,
					EventId: mq.NormalMsg,
					GameMsg: gameMsg,
				})
				return
			}
		}
		// 转发到反作弊服务器
		if session.anticheatServerAppId != "" {
			if protoMsg.CmdId == cmd.CombatInvocationsNotify ||
				protoMsg.CmdId == cmd.ToTheMoonEnterSceneReq {
				k.messageQueue.SendToAnticheat(session.anticheatServerAppId, &mq.NetMsg{
					MsgType: mq.MsgTypeGame,
					EventId: mq.NormalMsg,
					GameMsg: gameMsg,
				})
			}
		}
		// 转发到GS
		k.messageQueue.SendToGs(session.gsServerAppId, &mq.NetMsg{
			MsgType: mq.MsgTypeGame,
			EventId: mq.NormalMsg,
			GameMsg: gameMsg,
		})
	}
}

// 转发其他服务器的消息到客户端 所有连接共享一个协程
func (k *KcpConnectManager) forwardServerMsgToClientHandle() {
	logger.Debug("server msg forward handle start")
	// 函数栈内缓存 添加删除事件走chan 避免频繁加锁
	convSessionMap := make(map[uint64]*Session)
	userIdConvMap := make(map[uint32]uint64)
	// 远程全局顶号注册列表
	reLoginRemoteKickRegMap := make(map[uint32]chan bool)
	for {
		select {
		case session := <-k.createSessionChan:
			convSessionMap[session.conn.GetConv()] = session
			userIdConvMap[session.userId] = session.conn.GetConv()
		case session := <-k.destroySessionChan:
			delete(convSessionMap, session.conn.GetConv())
			delete(userIdConvMap, session.userId)
			close(session.kcpRawSendChan)
		case remoteKick := <-k.reLoginRemoteKickRegChan:
			reLoginRemoteKickRegMap[remoteKick.userId] = remoteKick.kickFinishNotifyChan
			remoteKick.regFinishNotifyChan <- true
		case netMsg := <-k.messageQueue.GetNetMsg():
			switch netMsg.MsgType {
			case mq.MsgTypeGame:
				gameMsg := netMsg.GameMsg
				switch netMsg.EventId {
				case mq.NormalMsg:
					// 分发到每个连接具体的发送协程
					convId, exist := userIdConvMap[gameMsg.UserId]
					if !exist {
						logger.Error("can not find convId by userId")
						continue
					}
					protoMsg := &ProtoMsg{
						ConvId:         convId,
						CmdId:          gameMsg.CmdId,
						HeadMessage:    k.getHeadMsg(gameMsg.ClientSeq),
						PayloadMessage: gameMsg.PayloadMessage,
					}
					session := convSessionMap[protoMsg.ConvId]
					if session == nil {
						logger.Error("session is nil, convId: %v", protoMsg.ConvId)
						continue
					}
					kcpRawSendChan := session.kcpRawSendChan
					if kcpRawSendChan == nil {
						logger.Error("kcpRawSendChan is nil, convId: %v", protoMsg.ConvId)
						continue
					}
					if len(kcpRawSendChan) == 1000 {
						logger.Error("kcpRawSendChan is full, convId: %v", protoMsg.ConvId)
						continue
					}
					if protoMsg.CmdId == cmd.PlayerLoginRsp {
						logger.Debug("session active, convId: %v", protoMsg.ConvId)
						session.connState = ConnActive
						// 通知GS玩家各个服务器的appid
						serverMsg := &mq.ServerMsg{
							UserId:               session.userId,
							AnticheatServerAppId: session.anticheatServerAppId,
						}
						k.messageQueue.SendToGs(session.gsServerAppId, &mq.NetMsg{
							MsgType:   mq.MsgTypeServer,
							EventId:   mq.ServerAppidBindNotify,
							ServerMsg: serverMsg,
						})
					}
					kcpRawSendChan <- protoMsg
				}
			case mq.MsgTypeConnCtrl:
				connCtrlMsg := netMsg.ConnCtrlMsg
				switch netMsg.EventId {
				case mq.KickPlayerNotify:
					convId, exist := userIdConvMap[connCtrlMsg.KickUserId]
					if !exist {
						logger.Error("can not find convId by userId")
						continue
					}
					k.kcpEventInput <- &KcpEvent{
						ConvId:       convId,
						EventId:      KcpConnForceClose,
						EventMessage: connCtrlMsg.KickReason,
					}
				}
			case mq.MsgTypeServer:
				serverMsg := netMsg.ServerMsg
				switch netMsg.EventId {
				case mq.ServerUserGsChangeNotify:
					convId, exist := userIdConvMap[serverMsg.UserId]
					if !exist {
						logger.Error("can not find convId by userId")
						continue
					}
					session := convSessionMap[convId]
					if session == nil {
						logger.Error("session is nil, convId: %v", convId)
						continue
					}
					session.gsServerAppId = serverMsg.GameServerAppId
					session.anticheatServerAppId = ""
					// 网关代发登录请求到新的GS
					gameMsg := &mq.GameMsg{
						UserId:    serverMsg.UserId,
						CmdId:     cmd.PlayerLoginReq,
						ClientSeq: 0,
						PayloadMessage: &proto.PlayerLoginReq{
							TargetUid:          serverMsg.JoinHostUserId,
							TargetHomeOwnerUid: 0,
						},
					}
					k.messageQueue.SendToGs(session.gsServerAppId, &mq.NetMsg{
						MsgType: mq.MsgTypeGame,
						EventId: mq.NormalMsg,
						GameMsg: gameMsg,
					})
				case mq.ServerUserOnlineStateChangeNotify:
					// 收到GS玩家离线完成通知
					logger.Debug("global player online state change, uid: %v, online: %v, gs appid: %v",
						serverMsg.UserId, serverMsg.IsOnline, netMsg.OriginServerAppId)
					if serverMsg.IsOnline {
						k.globalGsOnlineMapLock.Lock()
						k.globalGsOnlineMap[serverMsg.UserId] = netMsg.OriginServerAppId
						k.globalGsOnlineMapLock.Unlock()
					} else {
						k.globalGsOnlineMapLock.Lock()
						delete(k.globalGsOnlineMap, serverMsg.UserId)
						k.globalGsOnlineMapLock.Unlock()
						kickFinishNotifyChan, exist := reLoginRemoteKickRegMap[serverMsg.UserId]
						if !exist {
							continue
						}
						// 唤醒存在的顶号登录流程
						logger.Info("awake interrupt login, uid: %v", serverMsg.UserId)
						kickFinishNotifyChan <- true
						delete(reLoginRemoteKickRegMap, serverMsg.UserId)
					}
				}
			}
		}
	}
}

// 转发客户端消息到robot
func (k *KcpConnectManager) forwardClientMsgToRobotHandle(protoMsg *ProtoMsg, session *Session) {
	if protoMsg.HeadMessage == nil {
		logger.Error("recv null head msg: %v", protoMsg)
		return
	}
	// 网关服务器本地处理的请求
	switch protoMsg.CmdId {
	case cmd.GetPlayerTokenReq:
		// GATE登录包
		if session.connState != ConnEst {
			return
		}
		session.connState = ConnWaitLogin
		req := protoMsg.PayloadMessage.(*proto.GetPlayerTokenReq)
		session.keyId = req.KeyId
		session.clientRandKey = req.ClientRandKey
	case cmd.PlayerForceExitReq:
		// 退出游戏
		if session.connState != ConnActive {
			return
		}
	case cmd.PlayerLoginReq:
		// GS登录包
		if session.connState != ConnWaitLogin {
			return
		}
	}
	gameMsg := &mq.GameMsg{
		UserId:             session.userId,
		CmdId:              protoMsg.CmdId,
		ClientSeq:          protoMsg.HeadMessage.ClientSequenceId,
		PayloadMessageData: nil,
	}
	// 在这里直接序列化成二进制数据 终结PayloadMessage的生命周期并回收进缓存池
	payloadMessageData, err := pb.Marshal(protoMsg.PayloadMessage)
	if err != nil {
		logger.Error("parse payload msg to bin error: %v, stack: %v", err, logger.Stack())
		return
	}
	k.serverCmdProtoMap.PutProtoObjCache(protoMsg.CmdId, protoMsg.PayloadMessage)
	gameMsg.PayloadMessageData = payloadMessageData
	// 转发到Robot
	k.messageQueue.SendToRobot(session.robotServerAppId, &mq.NetMsg{
		MsgType: mq.MsgTypeGame,
		EventId: mq.NormalMsg,
		GameMsg: gameMsg,
	})
}

// 转发robot的消息到客户端
func (k *KcpConnectManager) forwardRobotMsgToClientHandle(session *Session) {
	logger.Debug("robot msg forward handle start")
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			if session.connState == ConnClose {
				logger.Debug("robot msg forward handle stop")
				return
			}
		case netMsg := <-k.messageQueue.GetNetMsg():
			if netMsg.OriginServerType != api.ROBOT {
				continue
			}
			switch netMsg.MsgType {
			case mq.MsgTypeGame:
				gameMsg := netMsg.GameMsg
				switch netMsg.EventId {
				case mq.NormalMsg:
					switch gameMsg.CmdId {
					case cmd.GetPlayerTokenRsp:
						rsp := gameMsg.PayloadMessage.(*proto.GetPlayerTokenRsp)
						ok := k.keyExchange(session, rsp.Uid, rsp)
						if !ok {
							logger.Error("key exchange error, uid: %v", rsp.Uid)
							continue
						}
						// 关联玩家uid和连接信息
						session.userId = rsp.Uid
						k.SetSession(session, session.conn.GetConv(), session.userId)
						gameMsg.PayloadMessage = rsp
					case cmd.PlayerLoginRsp:
						logger.Debug("session active, uid: %v", gameMsg.UserId)
						session.connState = ConnActive
					}
					protoMsg := &ProtoMsg{
						ConvId:         session.conn.GetConv(),
						CmdId:          gameMsg.CmdId,
						HeadMessage:    k.getHeadMsg(gameMsg.ClientSeq),
						PayloadMessage: gameMsg.PayloadMessage,
					}
					kcpRawSendChan := session.kcpRawSendChan
					if kcpRawSendChan == nil {
						logger.Error("kcpRawSendChan is nil, convId: %v", protoMsg.ConvId)
						continue
					}
					if len(kcpRawSendChan) == 1000 {
						logger.Error("kcpRawSendChan is full, convId: %v", protoMsg.ConvId)
						continue
					}
					kcpRawSendChan <- protoMsg
				}
			case mq.MsgTypeServer:
				switch netMsg.EventId {
				case mq.ServerForwardModeServerCloseNotify:
					_ = session.conn.Close()
				}
			}
		}
	}
}

func (k *KcpConnectManager) getHeadMsg(clientSeq uint32) *proto.PacketHead {
	headMsg := new(proto.PacketHead)
	if clientSeq != 0 {
		headMsg.ClientSequenceId = clientSeq
		headMsg.SentMs = uint64(time.Now().UnixMilli())
	}
	return headMsg
}

type RemoteKick struct {
	regFinishNotifyChan  chan bool
	userId               uint32
	kickFinishNotifyChan chan bool
}

func (k *KcpConnectManager) loginFailClose(session *Session) {
	k.kcpEventInput <- &KcpEvent{
		ConvId:       session.conn.GetConv(),
		EventId:      KcpConnForceClose,
		EventMessage: uint32(kcp.EnetLoginUnfinished),
	}
}

func (k *KcpConnectManager) loginFailRsp(uid uint32, retCode proto.Retcode, isForbid bool, forbidEndTime uint32) *proto.GetPlayerTokenRsp {
	rsp := new(proto.GetPlayerTokenRsp)
	rsp.Uid = uid
	rsp.Retcode = int32(retCode)
	if isForbid {
		rsp.Msg = "FORBID_CHEATING_PLUGINS"
		rsp.BlackUidEndTime = forbidEndTime
		if rsp.BlackUidEndTime == 0 {
			rsp.BlackUidEndTime = 2051193600 // 2035-01-01 00:00:00
		}
	}
	return rsp
}

func (k *KcpConnectManager) doGateLogin(req *proto.GetPlayerTokenReq, session *Session) *proto.GetPlayerTokenRsp {
	tokenVerifyRsp, err := httpclient.PostJson[controller.TokenVerifyRsp](
		config.GetConfig().Hk4e.LoginSdkUrl+"/gate/token/verify",
		&controller.TokenVerifyReq{
			AccountId:    req.AccountUid,
			AccountToken: req.AccountToken,
		})
	if err != nil {
		logger.Error("verify token error: %v, account uid: %v", err, req.AccountUid)
		k.loginFailClose(session)
		return nil
	}
	uid := tokenVerifyRsp.PlayerID
	if !tokenVerifyRsp.Valid {
		logger.Error("token error, uid: %v", uid)
		return k.loginFailRsp(uid, proto.Retcode_RET_TOKEN_ERROR, false, 0)
	}
	// comboToken验证成功
	if tokenVerifyRsp.Forbid {
		// 封号通知
		return k.loginFailRsp(uid, proto.Retcode_RET_BLACK_UID, true, tokenVerifyRsp.ForbidEndTime)
	}
	clientConnNum := atomic.LoadInt32(&CLIENT_CONN_NUM)
	if clientConnNum > MaxClientConnNumLimit {
		logger.Error("gate conn num limit, uid: %v", uid)
		return k.loginFailRsp(uid, proto.Retcode_RET_MAX_PLAYER, false, 0)
	}
	k.globalGsOnlineMapLock.RLock()
	_, exist := k.globalGsOnlineMap[uid]
	k.globalGsOnlineMapLock.RUnlock()
	if exist {
		// 注册回调通知
		regFinishNotifyChan := make(chan bool, 1)
		kickFinishNotifyChan := make(chan bool, 1)
		k.reLoginRemoteKickRegChan <- &RemoteKick{
			regFinishNotifyChan:  regFinishNotifyChan,
			userId:               uid,
			kickFinishNotifyChan: kickFinishNotifyChan,
		}
		// 注册等待
		logger.Info("run global interrupt login reg wait, uid: %v", uid)
		timer := time.NewTimer(time.Second * 1)
		select {
		case <-timer.C:
			logger.Error("global interrupt login reg wait timeout, uid: %v", uid)
			timer.Stop()
			k.loginFailClose(session)
			return nil
		case <-regFinishNotifyChan:
			timer.Stop()
		}
		oldSession := k.GetSessionByUserId(uid)
		if oldSession != nil {
			// 本地顶号
			k.kcpEventInput <- &KcpEvent{
				ConvId:  oldSession.conn.GetConv(),
				EventId: KcpConnRelogin,
			}
		} else {
			// 远程顶号
			connCtrlMsg := new(mq.ConnCtrlMsg)
			connCtrlMsg.KickUserId = uid
			connCtrlMsg.KickReason = kcp.EnetServerRelogin
			k.messageQueue.SendToAll(&mq.NetMsg{
				MsgType:     mq.MsgTypeConnCtrl,
				EventId:     mq.KickPlayerNotify,
				ConnCtrlMsg: connCtrlMsg,
			})
		}
		// 顶号等待
		logger.Info("run global interrupt login kick wait, uid: %v", uid)
		timer = time.NewTimer(time.Second * 10)
		select {
		case <-timer.C:
			logger.Error("global interrupt login kick wait timeout, uid: %v", uid)
			timer.Stop()
			k.loginFailClose(session)
			return nil
		case <-kickFinishNotifyChan:
			timer.Stop()
		}
	}
	// 关联玩家uid和连接信息
	session.userId = uid
	k.SetSession(session, session.conn.GetConv(), session.userId)
	k.createSessionChan <- session
	// 绑定各个服务器appid
	gsServerAppId, err := k.discovery.GetServerAppId(context.TODO(), &api.GetServerAppIdReq{
		ServerType: api.GS,
	})
	if err != nil {
		logger.Error("get gs server appid error: %v, uid: %v", err, uid)
		k.loginFailClose(session)
		return nil
	}
	session.gsServerAppId = gsServerAppId.AppId
	anticheatServerAppId, err := k.discovery.GetServerAppId(context.TODO(), &api.GetServerAppIdReq{
		ServerType: api.ANTICHEAT,
	})
	if err != nil {
		logger.Error("get anticheat server appid error: %v, uid: %v", err, uid)
	}
	session.anticheatServerAppId = anticheatServerAppId.AppId
	pathfindingServerAppId, err := k.discovery.GetServerAppId(context.TODO(), &api.GetServerAppIdReq{
		ServerType: api.PATHFINDING,
	})
	if err != nil {
		logger.Error("get pathfinding server appid error: %v, uid: %v", err, uid)
	}
	session.pathfindingServerAppId = pathfindingServerAppId.AppId
	logger.Debug("session gs appid: %v, uid: %v", session.gsServerAppId, uid)
	logger.Debug("session anticheat appid: %v, uid: %v", session.anticheatServerAppId, uid)
	logger.Debug("session pathfinding appid: %v, uid: %v", session.pathfindingServerAppId, uid)
	// 构造响应
	rsp := k.buildGateLoginRsp(session, uid, req.AccountUid, req.AccountToken)
	// 密钥交换
	session.keyId = req.KeyId
	session.clientRandKey = req.ClientRandKey
	ok := k.keyExchange(session, uid, rsp)
	if !ok {
		logger.Error("key exchange error, uid: %v", uid)
		return nil
	}
	return rsp
}

func (k *KcpConnectManager) buildGateLoginRsp(session *Session, uid uint32, accountUid string, token string) *proto.GetPlayerTokenRsp {
	addr := session.conn.RemoteAddr().String()
	addrSplit := strings.Split(addr, ":")
	rsp := &proto.GetPlayerTokenRsp{
		Uid:                    uid,
		AccountUid:             accountUid,
		Token:                  token,
		AccountType:            1,
		ChannelId:              1,
		SubChannelId:           1,
		PlatformType:           3,
		RegPlatform:            3,
		IsProficientPlayer:     false,
		CountryCode:            "US",
		Birthday:               "2000-01-01",
		ClientIpStr:            addrSplit[0],
		SecurityCmdBuffer:      random.GetRandomByte(32),
		ClientVersionRandomKey: fmt.Sprintf("%03x-%012x", random.GetRandomByte(3), random.GetRandomByte(12)),
	}
	return rsp
}

func (k *KcpConnectManager) keyExchange(session *Session, uid uint32, rsp *proto.GetPlayerTokenRsp) bool {
	timeRand := random.GetTimeRand()
	serverSeedUint64 := timeRand.Uint64()
	session.seed = serverSeedUint64
	if session.keyId != 0 {
		session.useMagicSeed = true
		keyId := strconv.Itoa(int(session.keyId))
		encPubPrivKey, exist := k.encRsaKeyMap[keyId]
		if !exist {
			logger.Error("can not found key id: %v, uid: %v", keyId, uid)
			k.loginFailClose(session)
			return false
		}
		pubKey, err := endec.RsaParsePubKeyByPrivKey(encPubPrivKey)
		if err != nil {
			logger.Error("parse rsa pub key error: %v, uid: %v", err, uid)
			k.loginFailClose(session)
			return false
		}
		signPrivkey, err := endec.RsaParsePrivKey(k.signRsaKey)
		if err != nil {
			logger.Error("parse rsa priv key error: %v, uid: %v", err, uid)
			k.loginFailClose(session)
			return false
		}
		clientSeedBase64 := session.clientRandKey
		clientSeedEnc, err := base64.StdEncoding.DecodeString(clientSeedBase64)
		if err != nil {
			logger.Error("parse client seed base64 error: %v, uid: %v", err, uid)
			k.loginFailClose(session)
			return false
		}
		clientSeed, err := endec.RsaDecrypt(clientSeedEnc, signPrivkey)
		if err != nil {
			logger.Error("rsa dec error: %v, uid: %v", err, uid)
			k.loginFailClose(session)
			return false
		}
		clientSeedUint64 := uint64(0)
		err = binary.Read(bytes.NewReader(clientSeed), binary.BigEndian, &clientSeedUint64)
		if err != nil {
			logger.Error("parse client seed to uint64 error: %v, uid: %v", err, uid)
			k.loginFailClose(session)
			return false
		}
		logger.Debug("clientSeed: %v, clientSeedUint64: %v", clientSeed, clientSeedUint64)
		logger.Debug("serverSeedUint64: %v", serverSeedUint64)
		seedUint64 := serverSeedUint64 ^ clientSeedUint64
		seedBuf := new(bytes.Buffer)
		err = binary.Write(seedBuf, binary.BigEndian, seedUint64)
		if err != nil {
			logger.Error("conv seed uint64 to bytes error: %v, uid: %v", err, uid)
			k.loginFailClose(session)
			return false
		}
		seed := seedBuf.Bytes()
		logger.Debug("seed: %v, seedUint64: %v", seed, seedUint64)
		seedEnc, err := endec.RsaEncrypt(seed, pubKey)
		if err != nil {
			logger.Error("rsa enc error: %v, uid: %v", err, uid)
			k.loginFailClose(session)
			return false
		}
		seedSign, err := endec.RsaSign(seed, signPrivkey)
		if err != nil {
			logger.Error("rsa sign error: %v, uid: %v", err, uid)
			k.loginFailClose(session)
			return false
		}
		rsp.KeyId = session.keyId
		rsp.ServerRandKey = base64.StdEncoding.EncodeToString(seedEnc)
		rsp.Sign = base64.StdEncoding.EncodeToString(seedSign)
	} else {
		session.useMagicSeed = false
		rsp.SecretKeySeed = serverSeedUint64
		rsp.SecretKey = fmt.Sprintf("%03x-%012x", random.GetRandomByte(3), random.GetRandomByte(12))
	}
	return true
}
