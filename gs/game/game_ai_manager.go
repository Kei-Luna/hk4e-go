package game

import (
	"hk4e/common/mq"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/protocol/proto"
)

// AiManager ai管理器
type AiManager struct {
	aiMap            map[uint32]*model.Player     // ai集合
	aiUidCounter     uint32                       // ai的uid计数器
	gameRoomAiUidMap map[uint32]map[uint32]uint32 // 游戏房间ai的uid集合
}

func NewAiManager() *AiManager {
	r := new(AiManager)

	r.aiMap = make(map[uint32]*model.Player)
	r.gameRoomAiUidMap = make(map[uint32]map[uint32]uint32)

	return r
}

// CreateAi 创建ai机器人
func (a *AiManager) CreateAi(name string, sign string, headImage uint32, nameCard uint32, multiServerAppId string) *model.Player {
	a.aiUidCounter++
	uid := GAME.gsId*AiBaseUid + a.aiUidCounter
	GAME.OnLogin(uid, 0, "", nil, new(proto.PlayerLoginReq), true)
	ai := USER_MANAGER.GetOnlineUser(uid)
	ai.DbState = model.DbNormal

	// 记录ai
	a.aiMap[ai.PlayerId] = ai

	GAME.SetPlayerBornDataReq(ai, &proto.SetPlayerBornDataReq{AvatarId: 10000007, NickName: name})
	ai.Signature = sign
	ai.HeadImage = headImage
	ai.NameCard = nameCard
	ai.MultiServerAppId = multiServerAppId
	world := WORLD_MANAGER.GetWorldById(ai.WorldId)
	GAME.EnterSceneReadyReq(ai, &proto.EnterSceneReadyReq{
		EnterSceneToken: world.GetEnterSceneToken(),
	})
	GAME.SceneInitFinishReq(ai, &proto.SceneInitFinishReq{
		EnterSceneToken: world.GetEnterSceneToken(),
	})
	GAME.EnterSceneDoneReq(ai, &proto.EnterSceneDoneReq{
		EnterSceneToken: world.GetEnterSceneToken(),
	})
	GAME.PostEnterSceneReq(ai, &proto.PostEnterSceneReq{
		EnterSceneToken: world.GetEnterSceneToken(),
	})
	GAME.EntityForceSyncReq(ai, &proto.EntityForceSyncReq{
		MotionInfo: &proto.MotionInfo{
			Pos: &proto.Vector{X: 500.0, Y: 900.0, Z: -500.0},
			Rot: new(proto.Vector),
		},
		EntityId: world.GetPlayerWorldAvatarEntityId(ai, 10000007),
	})
	return ai
}

// DeleteAi 清除ai机器人
func (a *AiManager) DeleteAi(uid uint32) {
	// 确保uid为ai
	_, exist := a.aiMap[uid]
	if !exist {
		logger.Error("ai not exist, uid: %v", uid)
		return
	}
	// ai离线
	GAME.OnOffline(uid, &ChangeGsInfo{
		IsChangeGs: false,
	})
	delete(a.aiMap, uid)
}

// IsAi uid是否为ai
func (a *AiManager) IsAi(uid uint32) bool {
	_, exist := a.aiMap[uid]
	return exist
}

// GetAiUidGameId 获取ai的uid存在的游戏id
func (a *AiManager) GetAiUidGameId(uid uint32) (ok bool, matchGameId uint32, matchRoomId uint32) {
	for gameId, roomUidMap := range a.gameRoomAiUidMap {
		for roomId, u := range roomUidMap {
			if u == uid {
				return true, gameId, roomId
			}
		}
	}
	return
}

// ServerMatchCreateAiReq 匹配服创建ai请求
func (g *Game) ServerMatchCreateAiReq(gameType uint8, gameId uint32, roomId uint32, aiInfo *mq.AiInfo, multiAppId string) {
	// 校验ai是否已经注册
	roomUidMap, exist := AI_MANAGER.gameRoomAiUidMap[gameId]
	if !exist {
		roomUidMap = make(map[uint32]uint32)
		AI_MANAGER.gameRoomAiUidMap[gameId] = roomUidMap
	}
	uid, exist := roomUidMap[roomId]
	if exist {
		logger.Error("room has been create ai, gameId: %v, roomId: %v, uid: %v", gameId, roomId, uid)
		return
	}

	ai := AI_MANAGER.CreateAi(aiInfo.Name, aiInfo.Sign, aiInfo.HeadImage, aiInfo.NameCard, multiAppId)
	WORLD_MANAGER.InitAiWorld(ai)

	// 记录游戏房间ai的uid
	roomUidMap[roomId] = ai.PlayerId

	// 触发事件
	if PLUGIN_MANAGER.TriggerEvent(PluginEventIdMatchCreateAi, &PluginEventMatchCreateAi{
		PluginEvent: NewPluginEvent(),
		GameType:    gameType,
		GameId:      gameId,
		RoomId:      roomId,
		Ai:          ai,
	}) {
		return
	}

	MESSAGE_QUEUE.SendToMulti(multiAppId, &mq.NetMsg{
		MsgType: mq.MsgTypeServer,
		EventId: mq.ServerMatchCreateAiRsp,
		ServerMsg: &mq.ServerMsg{
			MatchGameId: gameId,
			MatchRoomId: roomId,
			MatchAiUid:  ai.PlayerId,
		},
	})
}
