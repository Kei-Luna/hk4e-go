package game

import (
	"os"
	"runtime"
	"strconv"
	"time"

	"hk4e/common/mq"
	"hk4e/common/rpc"
	"hk4e/gate/kcp"
	"hk4e/gs/dao"
	"hk4e/gs/model"
	"hk4e/pkg/alg"
	"hk4e/pkg/logger"
	"hk4e/pkg/reflection"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

const (
	PlayerBaseUid    = 100000000
	MaxPlayerBaseUid = 200000000
	AiBaseUid        = 10000
	AiName           = "小可爱"
	AiSign           = "快捷指令"
)

var GAME *Game = nil
var LOCAL_EVENT_MANAGER *LocalEventManager = nil
var ROUTE_MANAGER *RouteManager = nil
var USER_MANAGER *UserManager = nil
var WORLD_MANAGER *WorldManager = nil
var TICK_MANAGER *TickManager = nil
var COMMAND_MANAGER *CommandManager = nil
var GCG_MANAGER *GCGManager = nil
var PLUGIN_MANAGER *PluginManager = nil
var AI_MANAGER *AiManager = nil
var MESSAGE_QUEUE *mq.MessageQueue

var ONLINE_PLAYER_NUM int32 = 0 // 当前在线玩家数

var SELF *model.Player

type Game struct {
	discovery      *rpc.DiscoveryClient // node节点服务器的natsrpc客户端
	db             *dao.Dao
	snowflake      *alg.SnowflakeWorker
	gsId           uint32
	gsAppid        string
	gsAppVersion   string
	ai             *model.Player // 本服的Ai玩家对象
	isStop         bool
	dispatchCancel bool
}

func NewGameCore(db *dao.Dao, messageQueue *mq.MessageQueue, gsId uint32, gsAppid string, gsAppVersion string, discovery *rpc.DiscoveryClient) (r *Game) {
	r = new(Game)
	r.discovery = discovery
	r.db = db
	MESSAGE_QUEUE = messageQueue
	r.snowflake = alg.NewSnowflakeWorker(int64(gsId))
	r.gsId = gsId
	r.gsAppid = gsAppid
	r.gsAppVersion = gsAppVersion
	GAME = r
	LOCAL_EVENT_MANAGER = NewLocalEventManager()
	ROUTE_MANAGER = NewRouteManager()
	USER_MANAGER = NewUserManager(db)
	WORLD_MANAGER = NewWorldManager(r.snowflake)
	TICK_MANAGER = NewTickManager()
	COMMAND_MANAGER = NewCommandManager()
	GCG_MANAGER = NewGCGManager()
	PLUGIN_MANAGER = NewPluginManager()
	AI_MANAGER = NewAiManager()
	RegLuaScriptLibFunc()
	// 创建本服的Ai
	name := AiName
	sign := AiSign + " GS:" + strconv.Itoa(int(gsId))
	r.ai = AI_MANAGER.CreateAi(name, sign, 10000007, 210001, "")
	COMMAND_MANAGER.SetSystem(r.ai)
	// 初始化插件 最后再调用以免插件需要访问其他模块导致出错
	PLUGIN_MANAGER.InitPlugin()
	r.run()
	r.isStop = false
	r.dispatchCancel = false
	return r
}

func (g *Game) GetGsId() uint32 {
	return g.gsId
}

func (g *Game) GetGsAppid() string {
	return g.gsAppid
}

// GetAi 获取本服的Ai玩家对象
func (g *Game) GetAi() *model.Player {
	return g.ai
}

func (g *Game) run() {
	go g.gameMainLoopD()
}

func (g *Game) gameMainLoopD() {
	times := 1
	panicCounter := 0
	lastPanicTime := time.Now().UnixNano()
	for {
		logger.Warn("game main loop start, times: %v", times)
		g.gameMainLoop()
		logger.Warn("game main loop stop, times: %v", times)
		times++
		panicCounter++
		if panicCounter > 10 {
			now := time.Now().UnixNano()
			if now-lastPanicTime > int64(time.Second) {
				panicCounter = 0
				lastPanicTime = now
			} else {
				logger.Error("!!! GAME MAIN LOOP STOP !!!")
				time.Sleep(time.Second * 10)
				os.Exit(-1)
			}
		}
	}
}

func (g *Game) gameMainLoop() {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GAME MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			if SELF != nil {
				logger.Error("the motherfucker player uid: %v", SELF.PlayerId)
				// info, _ := json.Marshal(SELF)
				// logger.Error("the motherfucker player info: %v", string(info))
				GAME.KickPlayer(SELF.PlayerId, kcp.EnetServerKick)
			}
		}
	}()
	intervalTime := time.Second.Nanoseconds() * 60
	lastTime := time.Now().UnixNano()
	routeCost := int64(0)
	tickCost := int64(0)
	localEventCost := int64(0)
	commandCost := int64(0)
	routeCount := int64(0)
	maxRouteCost := int64(0)
	maxRouteCmdId := uint16(0)
	runtime.LockOSThread()
	for {
		// 消耗CPU时间性能统计
		now := time.Now().UnixNano()
		if now-lastTime > intervalTime {
			routeCost /= 1e6
			tickCost /= 1e6
			localEventCost /= 1e6
			commandCost /= 1e6
			maxRouteCost /= 1e6
			logger.Info("[GAME MAIN LOOP] cpu time cost detail, routeCost: %v ms, tickCost: %v ms, localEventCost: %v ms, commandCost: %v ms",
				routeCost, tickCost, localEventCost, commandCost)
			totalCost := routeCost + tickCost + localEventCost + commandCost
			logger.Info("[GAME MAIN LOOP] cpu time cost percent, routeCost: %v%%, tickCost: %v%%, localEventCost: %v%%, commandCost: %v%%",
				float32(routeCost)/float32(totalCost)*100.0,
				float32(tickCost)/float32(totalCost)*100.0,
				float32(localEventCost)/float32(totalCost)*100.0,
				float32(commandCost)/float32(totalCost)*100.0)
			logger.Info("[GAME MAIN LOOP] total cpu time cost detail, totalCost: %v ms",
				totalCost)
			logger.Info("[GAME MAIN LOOP] total cpu time cost percent, totalCost: %v%%",
				float32(totalCost)/float32(intervalTime/1e6)*100.0)
			avgRouteCost := float32(0)
			if routeCount != 0 {
				avgRouteCost = float32(routeCost) / float32(routeCount)
			}
			logger.Info("[GAME MAIN LOOP] avg route cost: %v ms", avgRouteCost)
			logger.Info("[GAME MAIN LOOP] max route cost: %v ms, cmdId: %v", maxRouteCost, maxRouteCmdId)
			lastTime = now
			routeCost = 0
			tickCost = 0
			localEventCost = 0
			commandCost = 0
			routeCount = 0
			maxRouteCost = 0
			maxRouteCmdId = 0
		}
		select {
		case netMsg := <-MESSAGE_QUEUE.GetNetMsg():
			// 接收客户端消息
			start := time.Now().UnixNano()
			ROUTE_MANAGER.RouteHandle(netMsg)
			end := time.Now().UnixNano()
			if netMsg.MsgType == mq.MsgTypeGame && (end-start) > maxRouteCost {
				maxRouteCost = end - start
				maxRouteCmdId = netMsg.GameMsg.CmdId
			}
			routeCost += end - start
			routeCount++
		case <-TICK_MANAGER.GetGlobalTick().C:
			// 游戏服务器定时帧
			start := time.Now().UnixNano()
			TICK_MANAGER.OnGameServerTick()
			end := time.Now().UnixNano()
			tickCost += end - start
		case localEvent := <-LOCAL_EVENT_MANAGER.GetLocalEventChan():
			// 处理本地事件
			start := time.Now().UnixNano()
			LOCAL_EVENT_MANAGER.LocalEventHandle(localEvent)
			end := time.Now().UnixNano()
			localEventCost += end - start
		case command := <-COMMAND_MANAGER.GetCommandMessageInput():
			// 处理GM命令
			start := time.Now().UnixNano()
			COMMAND_MANAGER.HandleCommand(command)
			end := time.Now().UnixNano()
			commandCost += end - start
			logger.Info("run gm cmd cost: %v ns", end-start)
		}
	}
}

var EXIT_SAVE_FIN_CHAN chan bool

func (g *Game) ServerStopNotify() {
	go func() {
		info := "服务器即将关闭。"
		GAME.ServerAnnounceNotify(1, info)
		logger.Warn("stop game server last 1 minute")
		time.Sleep(time.Minute)
		delay := GAME.GetGsId()
		logger.Warn("stop game server last %v second", delay)
		time.Sleep(time.Second * time.Duration(delay))
		GAME.Close()
	}()
}

func (g *Game) Close() {
	if g.isStop {
		return
	}
	g.isStop = true
	logger.Warn("stop game server begin")
	// 保存玩家数据
	EXIT_SAVE_FIN_CHAN = make(chan bool)
	LOCAL_EVENT_MANAGER.GetLocalEventChan() <- &LocalEvent{
		EventId: ExitRunUserCopyAndSave,
	}
	<-EXIT_SAVE_FIN_CHAN
	logger.Warn("stop game server save player finish")
	// 告诉网关下线玩家并全服广播玩家离线
	userList := USER_MANAGER.GetAllOnlineUserList()
	for _, player := range userList {
		g.KickPlayer(player.PlayerId, kcp.EnetServerShutdown)
		MESSAGE_QUEUE.SendToAll(&mq.NetMsg{
			MsgType: mq.MsgTypeServer,
			EventId: mq.ServerUserOnlineStateChangeNotify,
			ServerMsg: &mq.ServerMsg{
				UserId:   player.PlayerId,
				IsOnline: false,
			},
		})
		time.Sleep(time.Millisecond * 100)
	}
	// 卸载插件
	PLUGIN_MANAGER.DelAllPlugin()
	logger.Warn("stop game server finish")
}

func (g *Game) ServerDispatchCancelNotify(appVersion string) {
	if appVersion != g.gsAppVersion {
		return
	}
	logger.Warn("game server dispatch cancel")
	g.dispatchCancel = true
}

// SendMsgToGate 发送消息给客户端 指定网关
func (g *Game) SendMsgToGate(cmdId uint16, userId uint32, clientSeq uint32, gateAppId string, payloadMsg pb.Message) {
	if userId < PlayerBaseUid {
		return
	}
	if payloadMsg == nil {
		logger.Error("payload msg is nil, stack: %v", logger.Stack())
		return
	}
	// 在这里直接序列化成二进制数据 防止发送的消息内包含各种游戏数据指针 而造成并发读写的问题
	payloadMessageData, err := pb.Marshal(payloadMsg)
	if err != nil {
		logger.Error("parse payload msg to bin error: %v, stack: %v", err, logger.Stack())
		return
	}
	gameMsg := &mq.GameMsg{
		UserId:             userId,
		CmdId:              cmdId,
		ClientSeq:          clientSeq,
		PayloadMessageData: payloadMessageData,
	}
	MESSAGE_QUEUE.SendToGate(gateAppId, &mq.NetMsg{
		MsgType: mq.MsgTypeGame,
		EventId: mq.NormalMsg,
		GameMsg: gameMsg,
	})
}

// SendMsg 发送消息给客户端
func (g *Game) SendMsg(cmdId uint16, userId uint32, clientSeq uint32, payloadMsg pb.Message) {
	if userId < PlayerBaseUid {
		return
	}
	if payloadMsg == nil {
		logger.Error("payload msg is nil, stack: %v", logger.Stack())
		return
	}
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player not exist, uid: %v, stack: %v", userId, logger.Stack())
		return
	}
	if !player.Online {
		return
	}
	if player.NetFreeze {
		return
	}
	gameMsg := new(mq.GameMsg)
	gameMsg.UserId = userId
	gameMsg.CmdId = cmdId
	gameMsg.ClientSeq = clientSeq
	// 在这里直接序列化成二进制数据 防止发送的消息内包含各种游戏数据指针 而造成并发读写的问题
	payloadMessageData, err := pb.Marshal(payloadMsg)
	if err != nil {
		logger.Error("parse payload msg to bin error: %v, stack: %v", err, logger.Stack())
		return
	}
	gameMsg.PayloadMessageData = payloadMessageData
	MESSAGE_QUEUE.SendToGate(player.GateAppId, &mq.NetMsg{
		MsgType: mq.MsgTypeGame,
		EventId: mq.NormalMsg,
		GameMsg: gameMsg,
	})
}

// SendError 通用返回错误码
func (g *Game) SendError(cmdId uint16, player *model.Player, rsp pb.Message, retCode ...proto.Retcode) {
	if rsp == nil {
		return
	}
	ret := int32(proto.Retcode_RET_FAIL)
	if len(retCode) == 0 {
		ret = int32(proto.Retcode_RET_SVR_ERROR)
	} else if len(retCode) == 1 {
		ret = int32(retCode[0])
	} else {
		return
	}
	ok := reflection.SetStructFieldValue(rsp, "Retcode", ret)
	if !ok {
		return
	}
	logger.Debug("send common error: %v", rsp)
	g.SendMsg(cmdId, player.PlayerId, player.ClientSeq, rsp)
}

// SendSucc 通用返回成功
func (g *Game) SendSucc(cmdId uint16, player *model.Player, rsp pb.Message) {
	if rsp == nil {
		return
	}
	ok := reflection.SetStructFieldValue(rsp, "Retcode", int32(proto.Retcode_RET_SUCC))
	if !ok {
		return
	}
	g.SendMsg(cmdId, player.PlayerId, player.ClientSeq, rsp)
}

// SendToWorldA 给世界内所有玩家发消息
func (g *Game) SendToWorldA(world *World, cmdId uint16, seq uint32, msg pb.Message, aecUid uint32) {
	for _, v := range world.GetAllPlayer() {
		if aecUid == v.PlayerId {
			continue
		}
		g.SendMsg(cmdId, v.PlayerId, seq, msg)
	}
}

// SendToWorldH 给世界房主发消息
func (g *Game) SendToWorldH(world *World, cmdId uint16, seq uint32, msg pb.Message) {
	g.SendMsg(cmdId, world.GetOwner().PlayerId, seq, msg)
}

// SendToSceneA 给场景内所有玩家发消息
func (g *Game) SendToSceneA(scene *Scene, cmdId uint16, seq uint32, msg pb.Message, aecUid uint32) {
	world := scene.GetWorld()
	if WORLD_MANAGER.IsAiWorld(world) && SELF != nil {
		aiWorldAoi := world.GetAiWorldAoi()
		otherWorldAvatarMap := aiWorldAoi.GetObjectListByPos(float32(SELF.Pos.X), float32(SELF.Pos.Y), float32(SELF.Pos.Z))
		for uid := range otherWorldAvatarMap {
			if aecUid == uint32(uid) {
				continue
			}
			g.SendMsg(cmdId, uint32(uid), seq, msg)
		}
	} else {
		for _, v := range scene.GetAllPlayer() {
			if aecUid == v.PlayerId {
				continue
			}
			g.SendMsg(cmdId, v.PlayerId, seq, msg)
		}
	}
}

// SendToSceneACV 给场景内所有指定客户端版本的玩家发消息
func (g *Game) SendToSceneACV(scene *Scene, cmdId uint16, seq uint32, msg pb.Message, aecUid uint32, clientVersion int) {
	world := scene.GetWorld()
	if WORLD_MANAGER.IsAiWorld(world) && SELF != nil {
		aiWorldAoi := world.GetAiWorldAoi()
		otherWorldAvatarMap := aiWorldAoi.GetObjectListByPos(float32(SELF.Pos.X), float32(SELF.Pos.Y), float32(SELF.Pos.Z))
		for uid := range otherWorldAvatarMap {
			player := USER_MANAGER.GetOnlineUser(uint32(uid))
			if player == nil {
				logger.Error("player not exist, uid: %v, stack: %v", uid, logger.Stack())
				continue
			}
			if aecUid == player.PlayerId {
				continue
			}
			if player.ClientVersion != clientVersion {
				continue
			}
			g.SendMsg(cmdId, uint32(uid), seq, msg)
		}
	} else {
		for _, v := range scene.GetAllPlayer() {
			if aecUid == v.PlayerId {
				continue
			}
			if v.ClientVersion != clientVersion {
				continue
			}
			g.SendMsg(cmdId, v.PlayerId, seq, msg)
		}
	}
}

func (g *Game) ReLoginPlayer(userId uint32, isQuitMp bool) {
	reason := proto.ClientReconnectReason_CLIENT_RECONNNECT_NONE
	if isQuitMp {
		reason = proto.ClientReconnectReason_CLIENT_RECONNNECT_QUIT_MP
	}
	g.SendMsg(cmd.ClientReconnectNotify, userId, 0, &proto.ClientReconnectNotify{
		Reason: reason,
	})
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		return
	}
	player.NetFreeze = true
}

func (g *Game) LogoutPlayer(userId uint32) {
	g.SendMsg(cmd.PlayerLogoutNotify, userId, 0, &proto.PlayerLogoutNotify{})
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		return
	}
	// 冻结掉服务器对该玩家的下行 避免大量发包对整个系统造成压力
	player.NetFreeze = true
}

func (g *Game) KickPlayer(userId uint32, reason uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		return
	}
	MESSAGE_QUEUE.SendToGate(player.GateAppId, &mq.NetMsg{
		MsgType: mq.MsgTypeConnCtrl,
		EventId: mq.KickPlayerNotify,
		ConnCtrlMsg: &mq.ConnCtrlMsg{
			KickUserId: userId,
			KickReason: reason,
		},
	})
	player.NetFreeze = true
}
