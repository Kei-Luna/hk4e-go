package game

import (
	"sync/atomic"
	"time"

	"hk4e/common/constant"
	"hk4e/common/mq"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

func (g *Game) PlayerLoginReq(userId uint32, clientSeq uint32, gateAppId string, payloadMsg pb.Message) {
	logger.Info("player login req, uid: %v, gateAppId: %v", userId, gateAppId)
	req := payloadMsg.(*proto.PlayerLoginReq)
	logger.Debug("login data: %v", req)
	USER_MANAGER.OnlineUser(userId, clientSeq, gateAppId, req.TargetUid)
}

func (g *Game) SetPlayerBornDataReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetPlayerBornDataReq)
	logger.Debug("avatar id: %v, nickname: %v", req.AvatarId, req.NickName)

	if player.IsBorn {
		logger.Error("player is already born, uid: %v", player.PlayerId)
		return
	}
	player.IsBorn = true

	mainCharAvatarId := req.AvatarId
	if mainCharAvatarId != 10000005 && mainCharAvatarId != 10000007 {
		logger.Error("invalid main char avatar id: %v", mainCharAvatarId)
		return
	}
	player.NickName = req.NickName
	player.HeadImage = mainCharAvatarId

	dbAvatar := player.GetDbAvatar()
	dbAvatar.MainCharAvatarId = mainCharAvatarId
	// 添加选定的主角
	dbAvatar.AddAvatar(player, dbAvatar.MainCharAvatarId)
	// 添加主角初始武器
	avatarDataConfig := gdconf.GetAvatarDataById(int32(dbAvatar.MainCharAvatarId))
	if avatarDataConfig == nil {
		logger.Error("get avatar data config is nil, avatarId: %v", dbAvatar.MainCharAvatarId)
		return
	}
	weaponId := uint64(g.snowflake.GenId())
	dbWeapon := player.GetDbWeapon()
	dbWeapon.AddWeapon(player, uint32(avatarDataConfig.InitialWeapon), weaponId)
	weapon := dbWeapon.WeaponMap[weaponId]
	dbAvatar.WearWeapon(dbAvatar.MainCharAvatarId, weapon)
	dbTeam := player.GetDbTeam()
	dbTeam.GetActiveTeam().SetAvatarIdList([]uint32{dbAvatar.MainCharAvatarId})

	g.AcceptQuest(player, false)

	g.LoginNotify(player.PlayerId, player.ClientSeq, player)

	// 创建世界
	world := WORLD_MANAGER.CreateWorld(player)
	world.AddPlayer(player, player.SceneId)
	player.WorldId = world.GetId()
	// 进入场景
	player.SceneJump = true
	player.SceneLoadState = model.SceneNone
	player.SceneEnterReason = uint32(proto.EnterReason_ENTER_REASON_LOGIN)
	g.SendMsg(cmd.PlayerEnterSceneNotify, player.PlayerId, player.ClientSeq, g.PacketPlayerEnterSceneNotifyLogin(player, proto.EnterType_ENTER_SELF))

	g.SendMsg(cmd.SetPlayerBornDataRsp, player.PlayerId, player.ClientSeq, new(proto.SetPlayerBornDataRsp))
}

func (g *Game) OnLogin(userId uint32, clientSeq uint32, gateAppId string, player *model.Player, joinHostUserId uint32) {
	if player == nil {
		player = g.CreatePlayer(userId)
		USER_MANAGER.ChangeUserDbState(player, model.DbInsert)
	}
	USER_MANAGER.AddUser(player)

	SELF = player

	player.OnlineTime = uint32(time.Now().UnixMilli())
	player.Online = true
	player.GateAppId = gateAppId

	// 初始化
	player.InitOnlineData()

	// 确保玩家位置安全
	player.Pos.X = player.SafePos.X
	player.Pos.Y = player.SafePos.Y
	player.Pos.Z = player.SafePos.Z
	if player.SceneId > 100 {
		player.SceneId = 3
		player.Pos = &model.Vector{X: 2747, Y: 194, Z: -1719}
		player.Rot = &model.Vector{X: 0, Y: 307, Z: 0}
	}

	TICK_MANAGER.CreateUserGlobalTick(userId)
	TICK_MANAGER.CreateUserTimer(userId, UserTimerActionTest, 100, player.NickName)

	if player.IsBorn {
		g.LoginNotify(userId, clientSeq, player)
		if joinHostUserId != 0 {
			hostPlayer := USER_MANAGER.GetOnlineUser(joinHostUserId)
			if hostPlayer != nil {
				g.JoinOtherWorld(player, hostPlayer)
			} else {
				logger.Error("player is nil, uid: %v", joinHostUserId)
			}
		} else {
			// 创建世界
			world := WORLD_MANAGER.CreateWorld(player)
			world.AddPlayer(player, player.SceneId)
			player.WorldId = world.GetId()
			// 进入场景
			player.SceneJump = true
			player.SceneLoadState = model.SceneNone
			player.SceneEnterReason = uint32(proto.EnterReason_ENTER_REASON_LOGIN)
			g.SendMsg(cmd.PlayerEnterSceneNotify, userId, clientSeq, g.PacketPlayerEnterSceneNotifyLogin(player, proto.EnterType_ENTER_SELF))
		}
	} else {
		g.SendMsg(cmd.DoSetPlayerBornDataNotify, userId, clientSeq, new(proto.DoSetPlayerBornDataNotify))
	}

	playerLoginRsp := &proto.PlayerLoginRsp{
		IsUseAbilityHash:        true,
		AbilityHashCode:         0,
		IsEnableClientHashDebug: true,
		IsScOpen:                false,
		ScInfo:                  []byte{},
		TotalTickTime:           0.0,
		GameBiz:                 "hk4e_global",
		RegisterCps:             "mihoyo",
		CountryCode:             "US",
		Birthday:                "2000-01-01",
	}
	g.SendMsg(cmd.PlayerLoginRsp, userId, clientSeq, playerLoginRsp)

	MESSAGE_QUEUE.SendToAll(&mq.NetMsg{
		MsgType: mq.MsgTypeServer,
		EventId: mq.ServerUserOnlineStateChangeNotify,
		ServerMsg: &mq.ServerMsg{
			UserId:   userId,
			IsOnline: true,
		},
	})
	atomic.AddInt32(&ONLINE_PLAYER_NUM, 1)

	SELF = nil
}

func (g *Game) CreatePlayer(userId uint32) *model.Player {
	player := new(model.Player)
	player.PlayerId = userId
	player.NickName = "旅行者"
	player.Signature = ""
	player.HeadImage = 10000007
	player.Birthday = []uint8{0, 0}
	player.NameCard = 210001
	player.NameCardList = make([]uint32, 0)
	player.FriendList = make(map[uint32]bool)
	player.FriendApplyList = make(map[uint32]bool)
	player.PropertiesMap = make(map[uint16]uint32)
	player.FlyCloakList = make([]uint32, 0)
	player.CostumeList = make([]uint32, 0)
	player.ChatMsgMap = make(map[uint32][]*model.ChatMsg)

	player.SceneId = 3

	player.NameCardList = append(player.NameCardList, 210001, 210042)

	player.PropertiesMap[constant.PLAYER_PROP_PLAYER_LEVEL] = 1
	player.PropertiesMap[constant.PLAYER_PROP_PLAYER_WORLD_LEVEL] = 0
	player.PropertiesMap[constant.PLAYER_PROP_IS_SPRING_AUTO_USE] = 1
	player.PropertiesMap[constant.PLAYER_PROP_SPRING_AUTO_USE_PERCENT] = 100
	player.PropertiesMap[constant.PLAYER_PROP_IS_FLYABLE] = 1
	player.PropertiesMap[constant.PLAYER_PROP_IS_TRANSFERABLE] = 1
	player.PropertiesMap[constant.PLAYER_PROP_MAX_STAMINA] = 24000
	player.PropertiesMap[constant.PLAYER_PROP_CUR_PERSIST_STAMINA] = 24000
	player.PropertiesMap[constant.PLAYER_PROP_PLAYER_RESIN] = 160
	player.PropertiesMap[constant.PLAYER_PROP_PLAYER_MP_SETTING_TYPE] = 2
	player.PropertiesMap[constant.PLAYER_PROP_IS_MP_MODE_AVAILABLE] = 1

	sceneLuaConfig := gdconf.GetSceneLuaConfigById(int32(player.SceneId))
	if sceneLuaConfig != nil {
		bornPos := sceneLuaConfig.SceneConfig.BornPos
		bornRot := sceneLuaConfig.SceneConfig.BornRot
		player.SafePos = &model.Vector{X: float64(bornPos.X), Y: float64(bornPos.Y), Z: float64(bornPos.Z)}
		player.Pos = &model.Vector{X: float64(bornPos.X), Y: float64(bornPos.Y), Z: float64(bornPos.Z)}
		player.Rot = &model.Vector{X: float64(bornRot.X), Y: float64(bornRot.Y), Z: float64(bornRot.Z)}
	} else {
		logger.Error("get scene lua config is nil, sceneId: %v, uid: %v", player.SceneId, player.PlayerId)
		player.SafePos = &model.Vector{X: 2747, Y: 194, Z: -1719}
		player.Pos = &model.Vector{X: 2747, Y: 194, Z: -1719}
		player.Rot = &model.Vector{X: 0, Y: 307, Z: 0}
	}
	return player
}

func (g *Game) ServerAppidBindNotify(userId uint32, anticheatAppId string) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	logger.Debug("server appid bind notify, uid: %v, anticheatAppId: %v", userId, anticheatAppId)
	player.AnticheatAppId = anticheatAppId
}

func (g *Game) OnOffline(userId uint32, changeGsInfo *ChangeGsInfo) {
	logger.Info("player offline, uid: %v", userId)
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	TICK_MANAGER.DestroyUserGlobalTick(userId)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world != nil {
		g.WorldRemovePlayer(world, player)
	}
	player.OfflineTime = uint32(time.Now().Unix())
	player.Online = false
	player.TotalOnlineTime += uint32(time.Now().UnixMilli()) - player.OnlineTime
	USER_MANAGER.OfflineUser(player, changeGsInfo)
	atomic.AddInt32(&ONLINE_PLAYER_NUM, -1)
}

func (g *Game) LoginNotify(userId uint32, clientSeq uint32, player *model.Player) {
	g.SendMsg(cmd.PlayerDataNotify, userId, clientSeq, g.PacketPlayerDataNotify(player))
	g.SendMsg(cmd.StoreWeightLimitNotify, userId, clientSeq, g.PacketStoreWeightLimitNotify())
	g.SendMsg(cmd.PlayerStoreNotify, userId, clientSeq, g.PacketPlayerStoreNotify(player))
	g.SendMsg(cmd.AvatarDataNotify, userId, clientSeq, g.PacketAvatarDataNotify(player))
	g.SendMsg(cmd.OpenStateUpdateNotify, userId, clientSeq, g.PacketOpenStateUpdateNotify())
	g.SendMsg(cmd.QuestListNotify, userId, clientSeq, g.PacketQuestListNotify(player))
	g.SendMsg(cmd.FinishedParentQuestNotify, userId, clientSeq, g.PacketFinishedParentQuestNotify(player))
	g.SendMsg(cmd.AllMarkPointNotify, player.PlayerId, player.ClientSeq, &proto.AllMarkPointNotify{MarkList: g.PacketMapMarkPointList(player)})
	// g.GCGLogin(player) // 发送GCG登录相关的通知包
}
