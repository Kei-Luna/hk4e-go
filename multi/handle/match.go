package handle

import (
	"fmt"
	"hk4e/common/constant"
	"hk4e/common/mq"
	"hk4e/pkg/logger"
)

// InitGame 初始化游戏
func (m *Match) InitGame() {
	m.CreateGame(constant.MATCH_GAME_TYPE_PUBG, 100, "<color=#FFCCE5>PUBG游戏</color>", "<color=#87CEFA>在原的世界中体验大逃杀吧～</color>", 10000007, 210001)
}

// MatchPhase 阶段
type MatchPhase uint8

const (
	MatchPhaseWait  = MatchPhase(iota) // 等待游戏开始
	MatchPhaseStart                    // 游戏开始
	MatchPhaseStop                     // 游戏结束
)

// MatchRoom 匹配房间
type MatchRoom struct {
	RoomId        uint32     // 房间id
	AiUid         uint32     // 房间世界对应ai的uid
	PlayerUidList []uint32   // 玩家uid列表
	Phase         MatchPhase // 游戏阶段
}

// MatchGame 匹配游戏
type MatchGame struct {
	GameId               uint32                // 游戏id
	GameType             uint8                 // 游戏类型
	MaxPlayerCount       uint16                // 最大玩家数量
	Name                 string                // 游戏名
	Sign                 string                // 描述签名
	HeadImage            uint32                // 头像id
	NameCard             uint32                // 名片id
	RoomMap              map[uint32]*MatchRoom // 房间集合 TODO gs关闭清除无用room
	roomIdCounter        uint32                // 房间id计数器
	WaitCreateAiRoomList []uint32              // 等待创建ai的房间请求列表
}

// Match 匹配功能
type Match struct {
	GameMap       map[uint32]*MatchGame // 游戏集合
	gameIdCounter uint32                // 游戏id计数器

	handle *Handle
}

func NewMatch(handle *Handle) *Match {
	m := new(Match)

	m.handle = handle
	m.GameMap = make(map[uint32]*MatchGame)

	m.InitGame()
	return m
}

// CreateGame 创建游戏
func (m *Match) CreateGame(gameType uint8, maxPlayerCount uint16, name, sign string, headImage, nameCard uint32) {
	m.gameIdCounter++
	m.GameMap[m.gameIdCounter] = &MatchGame{
		GameId:               m.gameIdCounter,
		GameType:             gameType,
		MaxPlayerCount:       maxPlayerCount,
		Name:                 name,
		Sign:                 sign,
		HeadImage:            headImage,
		NameCard:             nameCard,
		RoomMap:              make(map[uint32]*MatchRoom),
		WaitCreateAiRoomList: make([]uint32, 0),
	}
	// 创建游戏后添加默认房间
	m.AddRoom(m.gameIdCounter)
}

// RoomCreateAi 房间请求gs创建ai
func (m *Match) RoomCreateAi(gameId uint32, roomId uint32) {
	game, exist := m.GameMap[gameId]
	if !exist {
		logger.Error("game not exist, gameId: %v", gameId)
		return
	}
	// 请求负载最低的gs创建ai
	m.handle.messageQueue.SendToGs(m.handle.minLoadGsServerAppId, &mq.NetMsg{
		MsgType: mq.MsgTypeServer,
		EventId: mq.ServerMatchCreateAiReq,
		ServerMsg: &mq.ServerMsg{
			MatchGameId:   game.GameId,
			MatchRoomId:   roomId,
			MatchGameType: game.GameType,
			MatchAiInfo: &mq.AiInfo{
				Name:      game.Name,
				Sign:      game.Sign,
				HeadImage: game.HeadImage,
				NameCard:  game.NameCard,
			},
		},
	})
	// 添加到等待列表
	for _, id := range game.WaitCreateAiRoomList {
		if id == roomId {
			return
		}
	}
	game.WaitCreateAiRoomList = append(game.WaitCreateAiRoomList, roomId)
}

// CreateRoom 创建房间
func (m *Match) CreateRoom(gameId uint32, oldRoomId uint32) {
	game, exist := m.GameMap[gameId]
	if !exist {
		logger.Error("game not exist, gameId: %v", gameId)
		return
	}
	// 没有新房间则直接创建
	if game.roomIdCounter != oldRoomId {
		// 如果已有新房间校验人数
		room, exist := game.RoomMap[game.roomIdCounter]
		if !exist {
			logger.Error("room not exist, gameId: %v, roomId: %v", game.GameId, game.roomIdCounter)
			return
		}
		// 人数未满不创建
		if len(room.PlayerUidList) < int(game.MaxPlayerCount) {
			return
		}
	}
	m.AddRoom(gameId)
}

// AddRoom 添加房间
func (m *Match) AddRoom(gameId uint32) {
	game, exist := m.GameMap[gameId]
	if !exist {
		logger.Error("game not exist, gameId: %v", gameId)
		return
	}
	game.roomIdCounter++
	game.RoomMap[game.roomIdCounter] = &MatchRoom{
		RoomId:        game.roomIdCounter,
		AiUid:         0, // 请求负载最低的gs创建ai
		PlayerUidList: make([]uint32, 0),
		Phase:         MatchPhaseWait,
	}
	// 请求gs创建ai
	m.RoomCreateAi(game.GameId, game.roomIdCounter)
}

// HandleWaitCreateAiRoom 同步完负载最小的gs后请求创建ai
func (m *Match) HandleWaitCreateAiRoom() {
	for _, game := range m.GameMap {
		for _, roomId := range game.WaitCreateAiRoomList {
			m.RoomCreateAi(game.GameId, roomId)
		}
	}
}

// ServerGetMatchGameListReq 获取游戏列表请求
func (h *Handle) ServerGetMatchGameListReq(userId uint32, gsAppId string) {
	matchGameList := make([]*mq.MatchGameInfo, 0)
	// 遍历游戏列表
	for _, game := range h.match.GameMap {
		curRoom, exist := game.RoomMap[game.roomIdCounter]
		if !exist {
			logger.Error("room not exist, gameId: %v, roomId: %v", game.GameId, game.roomIdCounter)
			return
		}
		// 校验房间是否满人 满人创建新房间
		if len(curRoom.PlayerUidList) >= int(game.MaxPlayerCount) {
			h.match.CreateRoom(game.GameId, curRoom.RoomId)
		}
		matchGameList = append(matchGameList, &mq.MatchGameInfo{
			GameId: game.GameId,
			AiInfo: &mq.AiInfo{
				Name:      game.Name + " " + fmt.Sprintf("<color=#FF9999>[ %v / %v ]</color>", len(curRoom.PlayerUidList), game.MaxPlayerCount), // 添加玩家人数显示
				Sign:      game.Sign,
				HeadImage: game.HeadImage,
				NameCard:  game.NameCard,
			},
			PlayerCount: uint32(len(curRoom.PlayerUidList)),
		})
	}
	h.messageQueue.SendToGs(gsAppId, &mq.NetMsg{
		MsgType: mq.MsgTypeServer,
		EventId: mq.ServerGetMatchGameListRsp,
		ServerMsg: &mq.ServerMsg{
			UserId:        userId,
			MatchGameList: matchGameList,
		},
	})
}

// ServerGetMatchRoomAiUidReq 获取房间ai的uid请求
func (h *Handle) ServerGetMatchRoomAiUidReq(userId, gameId uint32, gsAppId string) {
	game, exist := h.match.GameMap[gameId]
	if !exist {
		logger.Error("game not exist, gameId: %v", game.GameId)
		return
	}
	// 获取现行房间
	curRoom, exist := game.RoomMap[game.roomIdCounter]
	if !exist {
		logger.Error("room not exist, gameId: %v, roomId: %v", game.GameId, game.roomIdCounter)
		return
	}
	aiUid := curRoom.AiUid
	// 确保房间处于等待状态
	if curRoom.Phase != MatchPhaseWait {
		aiUid = 0
	}
	// 确保房间没有满人
	if len(curRoom.PlayerUidList) >= int(game.MaxPlayerCount) {
		aiUid = 0
	}
	h.messageQueue.SendToGs(gsAppId, &mq.NetMsg{
		MsgType: mq.MsgTypeServer,
		EventId: mq.ServerGetMatchRoomAiUidRsp,
		ServerMsg: &mq.ServerMsg{
			UserId:     userId,
			MatchAiUid: aiUid,
		},
	})
}

// ServerMatchCreateAiRsp 创建ai响应
func (h *Handle) ServerMatchCreateAiRsp(gameId, roomId, aiUid uint32) {
	game, exist := h.match.GameMap[gameId]
	if !exist {
		logger.Error("game not exist, gameId: %v", game.GameId)
		return
	}
	// 获取房间
	room, exist := game.RoomMap[roomId]
	if !exist {
		logger.Error("room not exist, gameId: %v, roomId: %v", game.GameId, roomId)
		return
	}
	// 如果不为0代表已经设置过了
	if room.AiUid != 0 {
		logger.Error("ai uid has been set, aiUid: %v", room.AiUid)
		return
	}
	// 设置房间ai的uid
	room.AiUid = aiUid
	// 清除等待
	for i, u := range game.WaitCreateAiRoomList {
		if u == roomId {
			game.WaitCreateAiRoomList = append(game.WaitCreateAiRoomList[:i], game.WaitCreateAiRoomList[i+1:]...)
			break
		}
	}
}

// ServerMatchPlayerJoinGameNotify 玩家加入游戏通知
func (h *Handle) ServerMatchPlayerJoinGameNotify(userId, gameId, roomId uint32) {
	game, exist := h.match.GameMap[gameId]
	if !exist {
		logger.Error("game not exist, gameId: %v", game.GameId)
		return
	}
	// 获取房间
	room, exist := game.RoomMap[roomId]
	if !exist {
		logger.Error("room not exist, gameId: %v, roomId: %v", game.GameId, roomId)
		return
	}
	// 校验玩家是否已加入
	for _, uid := range room.PlayerUidList {
		if uid == userId {
			logger.Error("player has been join, uid: %v", userId)
			return
		}
	}
	// 添加玩家列表
	room.PlayerUidList = append(room.PlayerUidList, userId)
}

// ServerMatchPlayerExitGameNotify 玩家离开游戏通知
func (h *Handle) ServerMatchPlayerExitGameNotify(userId, gameId, roomId uint32) {
	game, exist := h.match.GameMap[gameId]
	if !exist {
		logger.Error("game not exist, gameId: %v", game.GameId)
		return
	}
	// 获取房间
	room, exist := game.RoomMap[roomId]
	if !exist {
		logger.Error("room not exist, gameId: %v, roomId: %v", game.GameId, roomId)
		return
	}
	// 清理玩家uid
	for i, uid := range room.PlayerUidList {
		if uid == userId {
			room.PlayerUidList = append(room.PlayerUidList[:i], room.PlayerUidList[i+1:]...)
		}
	}
}

// ServerMatchGameStartNotify 游戏开始通知
func (h *Handle) ServerMatchGameStartNotify(gameId, roomId uint32) {
	game, exist := h.match.GameMap[gameId]
	if !exist {
		logger.Error("game not exist, gameId: %v", game.GameId)
		return
	}
	// 获取房间
	room, exist := game.RoomMap[roomId]
	if !exist {
		logger.Error("room not exist, gameId: %v, roomId: %v", game.GameId, roomId)
		return
	}
	// 确保房间处于等待阶段
	if room.Phase != MatchPhaseWait {
		logger.Error("room phase not wait, gameId: %v, roomId: %v, phase: %v", game.GameId, roomId, room.Phase)
		return
	}
	// 游戏创建新房间
	h.match.CreateRoom(gameId, roomId)
	// 切换阶段
	room.Phase = MatchPhaseStart
}

// ServerMatchGameStopNotify 游戏结束通知
func (h *Handle) ServerMatchGameStopNotify(gameId, roomId uint32) {
	game, exist := h.match.GameMap[gameId]
	if !exist {
		logger.Error("game not exist, gameId: %v", game.GameId)
		return
	}
	// 获取房间
	room, exist := game.RoomMap[roomId]
	if !exist {
		logger.Error("room not exist, gameId: %v, roomId: %v", game.GameId, roomId)
		return
	}
	// 确保房间处于开始阶段
	if room.Phase != MatchPhaseStart {
		logger.Error("room phase not start, gameId: %v, roomId: %v, phase: %v", game.GameId, roomId, room.Phase)
		return
	}
	// 切换阶段
	room.Phase = MatchPhaseStop
	// 清除房间
	delete(game.RoomMap, roomId)
}
