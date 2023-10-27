package handle

import (
	"context"
	"hk4e/common/mq"
	"hk4e/common/rpc"
	"hk4e/gate/kcp"
	"hk4e/node/api"
	"hk4e/pkg/logger"
	"hk4e/protocol/cmd"
	"time"

	pb "google.golang.org/protobuf/proto"
)

type Handle struct {
	messageQueue         *mq.MessageQueue
	discoveryClient      *rpc.DiscoveryClient
	playerAcCtxMap       map[uint32]*AnticheatContext
	worldStatic          *WorldStatic
	match                *Match
	minLoadGsServerAppId string
}

func NewHandle(messageQueue *mq.MessageQueue, discoveryClient *rpc.DiscoveryClient) (r *Handle) {
	r = new(Handle)
	r.messageQueue = messageQueue
	r.discoveryClient = discoveryClient
	r.playerAcCtxMap = make(map[uint32]*AnticheatContext)
	r.worldStatic = NewWorldStatic()
	r.worldStatic.InitTerrain()
	r.syncMinLoadServerAppid()
	go r.autoSyncMinLoadServerAppid()
	r.match = NewMatch(r)
	go r.run()
	return r
}

func (h *Handle) run() {
	logger.Info("start handle")
	for {
		netMsg := <-h.messageQueue.GetNetMsg()
		switch netMsg.MsgType {
		case mq.MsgTypeGame:
			if netMsg.OriginServerType != api.GATE {
				continue
			}
			if netMsg.EventId != mq.NormalMsg {
				continue
			}
			gameMsg := netMsg.GameMsg
			switch gameMsg.CmdId {
			case cmd.CombatInvocationsNotify:
				h.CombatInvocationsNotify(gameMsg.UserId, netMsg.OriginServerAppId, gameMsg.PayloadMessage)
			case cmd.ToTheMoonEnterSceneReq:
				h.ToTheMoonEnterSceneReq(gameMsg.UserId, netMsg.OriginServerAppId, gameMsg.PayloadMessage)
			case cmd.QueryPathReq:
				h.QueryPath(gameMsg.UserId, netMsg.OriginServerAppId, gameMsg.PayloadMessage)
			case cmd.ObstacleModifyNotify:
				h.ObstacleModifyNotify(gameMsg.UserId, netMsg.OriginServerAppId, gameMsg.PayloadMessage)
			}
		case mq.MsgTypeServer:
			serverMsg := netMsg.ServerMsg
			switch netMsg.EventId {
			case mq.ServerUserOnlineStateChangeNotify:
				logger.Info("player online state change, state: %v, uid: %v", serverMsg.IsOnline, serverMsg.UserId)
				if serverMsg.IsOnline {
					h.AddPlayerAcCtx(serverMsg.UserId)
				} else {
					h.DelPlayerAcCtx(serverMsg.UserId)
				}
			case mq.ServerGetMatchGameListReq:
				h.ServerGetMatchGameListReq(serverMsg.UserId, netMsg.OriginServerAppId)
			case mq.ServerGetMatchRoomAiUidReq:
				h.ServerGetMatchRoomAiUidReq(serverMsg.UserId, serverMsg.MatchGameId, netMsg.OriginServerAppId)
			case mq.ServerMatchCreateAiRsp:
				h.ServerMatchCreateAiRsp(serverMsg.MatchGameId, serverMsg.MatchRoomId, serverMsg.MatchAiUid)
			case mq.ServerMatchPlayerJoinGameNotify:
				h.ServerMatchPlayerJoinGameNotify(serverMsg.UserId, serverMsg.MatchGameId, serverMsg.MatchRoomId)
			case mq.ServerMatchPlayerExitGameNotify:
				h.ServerMatchPlayerExitGameNotify(serverMsg.UserId, serverMsg.MatchGameId, serverMsg.MatchRoomId)
			case mq.ServerMatchGameStartNotify:
				h.ServerMatchGameStartNotify(serverMsg.MatchGameId, serverMsg.MatchRoomId)
			case mq.ServerMatchGameStopNotify:
				h.ServerMatchGameStopNotify(serverMsg.MatchGameId, serverMsg.MatchRoomId)
			}
		}
	}
}

func (h *Handle) syncMinLoadServerAppid() {
	gsServerAppId, err := h.discoveryClient.GetServerAppId(context.TODO(), &api.GetServerAppIdReq{
		ServerType: api.GS,
	})
	if err != nil {
		logger.Error("get gs server appid error: %v", err)
		h.minLoadGsServerAppId = ""
	} else {
		h.minLoadGsServerAppId = gsServerAppId.AppId
	}
}

func (h *Handle) autoSyncMinLoadServerAppid() {
	ticker := time.NewTicker(time.Second * 15)
	for {
		<-ticker.C
		h.syncMinLoadServerAppid()

		// 同步完负载最小的gs后重新尝试请求创建ai
		h.match.HandleWaitCreateAiRoom()
	}
}

func (h *Handle) KickPlayer(userId uint32, gateAppId string) {
	if !KickCheatPlayer {
		return
	}
	h.messageQueue.SendToGate(gateAppId, &mq.NetMsg{
		MsgType: mq.MsgTypeConnCtrl,
		EventId: mq.KickPlayerNotify,
		ConnCtrlMsg: &mq.ConnCtrlMsg{
			KickUserId: userId,
			KickReason: kcp.EnetServerKillClient,
		},
	})
}

// SendMsg 发送消息给客户端
func (h *Handle) SendMsg(cmdId uint16, userId uint32, gateAppId string, payloadMsg pb.Message) {
	if payloadMsg == nil {
		return
	}
	gameMsg := new(mq.GameMsg)
	gameMsg.UserId = userId
	gameMsg.CmdId = cmdId
	gameMsg.ClientSeq = 0
	// 在这里直接序列化成二进制数据 防止发送的消息内包含各种游戏数据指针 而造成并发读写的问题
	payloadMessageData, err := pb.Marshal(payloadMsg)
	if err != nil {
		logger.Error("parse payload msg to bin error: %v", err)
		return
	}
	gameMsg.PayloadMessageData = payloadMessageData
	h.messageQueue.SendToGate(gateAppId, &mq.NetMsg{
		MsgType: mq.MsgTypeGame,
		EventId: mq.NormalMsg,
		GameMsg: gameMsg,
	})
}
