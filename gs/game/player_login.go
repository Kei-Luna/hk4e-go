package game

import (
	"hk4e/common/constant"
	"hk4e/common/region"
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
	USER_MANAGER.UserLoginLoad(userId, clientSeq, gateAppId, req)
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

func (g *Game) OnLogin(userId uint32, clientSeq uint32, gateAppId string, player *model.Player, req *proto.PlayerLoginReq, ok bool) {
	if !ok {
		g.SendMsgToGate(cmd.PlayerLoginRsp, userId, clientSeq, gateAppId, &proto.PlayerLoginRsp{Retcode: int32(proto.Retcode_RET_LOGIN_INIT_FAIL)})
		return
	}

	if player == nil {
		logger.Info("reg new player, uid: %v", userId)
		player = g.CreatePlayer(userId)
		USER_MANAGER.ChangeUserDbState(player, model.DbInsert)
	}
	USER_MANAGER.OnlineUser(player)

	TICK_MANAGER.CreateUserGlobalTick(userId)
	TICK_MANAGER.CreateUserTimer(userId, UserTimerActionTest, 100, player.NickName)

	player.GateAppId = gateAppId

	SELF = player

	// 初始化
	player.InitOnlineData()

	if player.SceneId > 100 {
		player.SceneId = 3
		player.Pos = &model.Vector{X: 2747, Y: 194, Z: -1719}
		player.Rot = &model.Vector{X: 0, Y: 307, Z: 0}
	}

	dbQuest := player.GetDbQuest()
	for _, quest := range dbQuest.GetQuestMap() {
		if quest.State == constant.QUEST_STATE_UNFINISHED {
			quest.State = constant.QUEST_STATE_UNSTARTED
			g.StartQuest(player, quest.QuestId, false)
		}
	}

	if player.IsBorn {
		g.LoginNotify(userId, clientSeq, player)
		if req.TargetUid != 0 {
			hostPlayer := USER_MANAGER.GetOnlineUser(req.TargetUid)
			if hostPlayer != nil {
				g.JoinOtherWorld(player, hostPlayer)
			} else {
				logger.Error("player is nil, uid: %v", req.TargetUid)
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

	clientVersion, _ := region.GetClientVersionByName(req.ChecksumClientVersion)
	player.ClientVersion = clientVersion

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

	SELF = nil
}

func (g *Game) CreatePlayer(userId uint32) *model.Player {
	player := new(model.Player)
	player.PlayerId = userId
	player.NickName = "旅行者"
	player.Signature = ""
	player.HeadImage = 10000007
	player.PropMap = make(map[uint32]uint32)
	player.OpenStateMap = make(map[uint32]uint32)
	player.ChatMsgMap = make(map[uint32][]*model.ChatMsg)

	player.SceneId = 3

	player.PropMap[constant.PLAYER_PROP_PLAYER_WORLD_LEVEL] = 0
	player.PropMap[constant.PLAYER_PROP_CUR_PERSIST_STAMINA] = 10000
	player.PropMap[constant.PLAYER_PROP_CUR_TEMPORARY_STAMINA] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_LEVEL] = 1
	player.PropMap[constant.PLAYER_PROP_PLAYER_EXP] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_HCOIN] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_SCOIN] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_MP_SETTING_TYPE] = 2
	player.PropMap[constant.PLAYER_PROP_PLAYER_RESIN] = 160
	player.PropMap[constant.PLAYER_PROP_PLAYER_WAIT_SUB_HCOIN] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_WAIT_SUB_SCOIN] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_MCOIN] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_WAIT_SUB_MCOIN] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_LEGENDARY_KEY] = 0
	player.PropMap[constant.PLAYER_PROP_CUR_CLIMATE_METER] = 0
	player.PropMap[constant.PLAYER_PROP_CUR_CLIMATE_TYPE] = 0
	player.PropMap[constant.PLAYER_PROP_CUR_CLIMATE_AREA_ID] = 0
	player.PropMap[constant.PLAYER_PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_WORLD_LEVEL_LIMIT] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_WORLD_LEVEL_ADJUST_CD] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_HOME_COIN] = 0
	player.PropMap[constant.PLAYER_PROP_PLAYER_WAIT_SUB_HOME_COIN] = 0
	player.PropMap[constant.PLAYER_PROP_IS_AUTO_UNLOCK_SPECIFIC_EQUIP] = 0
	player.PropMap[constant.PLAYER_PROP_IS_SPRING_AUTO_USE] = 1
	player.PropMap[constant.PLAYER_PROP_SPRING_AUTO_USE_PERCENT] = 50
	player.PropMap[constant.PLAYER_PROP_IS_FLYABLE] = 0
	player.PropMap[constant.PLAYER_PROP_IS_WEATHER_LOCKED] = 1
	player.PropMap[constant.PLAYER_PROP_IS_GAME_TIME_LOCKED] = 1
	player.PropMap[constant.PLAYER_PROP_IS_TRANSFERABLE] = 1
	player.PropMap[constant.PLAYER_PROP_MAX_STAMINA] = 10000

	player.OpenStateMap[constant.OPEN_STATE_DERIVATIVE_MALL] = 1
	player.OpenStateMap[constant.OPEN_STATE_PHOTOGRAPH] = 1
	player.OpenStateMap[constant.OPEN_STATE_GUIDE_RELIC_PROM] = 1
	player.OpenStateMap[constant.OPEN_STATE_GUIDE_TALENT] = 1
	player.OpenStateMap[constant.OPEN_STATE_RELIQUARY_PROMOTE] = 1
	player.OpenStateMap[constant.OPEN_STATE_SHOP_TYPE_RECOMMANDED] = 1
	player.OpenStateMap[constant.OPEN_STATE_RELIQUARY_UPGRADE] = 1
	player.OpenStateMap[constant.OPEN_STATE_WEAPON_AWAKEN] = 1
	player.OpenStateMap[constant.OPEN_STATE_WEAPON_PROMOTE] = 1
	player.OpenStateMap[constant.OPEN_STATE_WEAPON_UPGRADE] = 1
	player.OpenStateMap[constant.OPEN_STATE_AVATAR_TALENT] = 1
	player.OpenStateMap[constant.OPEN_STATE_AVATAR_PROMOTE] = 1
	player.OpenStateMap[constant.OPEN_STATE_RESIN] = 1
	player.OpenStateMap[constant.OPEN_STATE_SHOP_TYPE_GENESISCRYSTAL] = 1
	player.OpenStateMap[constant.OPEN_STATE_SHOP_TYPE_GIFTPACKAGE] = 1
	player.OpenStateMap[constant.OPEN_STATE_BATTLE_PASS] = 1
	player.OpenStateMap[constant.OPEN_STATE_SHOP_TYPE_BLACKSMITH] = 1
	player.OpenStateMap[constant.OPEN_STATE_SHOP_TYPE_PAIMON] = 1
	player.OpenStateMap[constant.OPEN_STATE_SHOP_TYPE_VIRTUAL_SHOP] = 1

	sceneLuaConfig := gdconf.GetSceneLuaConfigById(int32(player.SceneId))
	if sceneLuaConfig != nil {
		bornPos := sceneLuaConfig.SceneConfig.BornPos
		bornRot := sceneLuaConfig.SceneConfig.BornRot
		player.Pos = &model.Vector{X: float64(bornPos.X), Y: float64(bornPos.Y), Z: float64(bornPos.Z)}
		player.Rot = &model.Vector{X: float64(bornRot.X), Y: float64(bornRot.Y), Z: float64(bornRot.Z)}
	} else {
		logger.Error("get scene lua config is nil, sceneId: %v, uid: %v", player.SceneId, player.PlayerId)
		player.Pos = &model.Vector{X: 2747, Y: 194, Z: -1719}
		player.Rot = &model.Vector{X: 0, Y: 307, Z: 0}
	}
	return player
}

func (g *Game) ServerAppidBindNotify(userId uint32, multiServerAppId string) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	logger.Debug("server appid bind notify, uid: %v, multiServerAppId: %v", userId, multiServerAppId)
	player.MultiServerAppId = multiServerAppId
}

func (g *Game) OnOffline(userId uint32, changeGsInfo *ChangeGsInfo) {
	logger.Info("player offline, uid: %v", userId)
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}

	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world != nil {
		g.WorldRemovePlayer(world, player)
	}

	TICK_MANAGER.DestroyUserGlobalTick(userId)

	USER_MANAGER.UserOfflineSave(player, changeGsInfo)
}

func (g *Game) LoginNotify(userId uint32, clientSeq uint32, player *model.Player) {
	g.SendMsg(cmd.PlayerDataNotify, userId, clientSeq, g.PacketPlayerDataNotify(player))
	g.SendMsg(cmd.StoreWeightLimitNotify, userId, clientSeq, g.PacketStoreWeightLimitNotify())
	g.SendMsg(cmd.PlayerStoreNotify, userId, clientSeq, g.PacketPlayerStoreNotify(player))
	g.SendMsg(cmd.AvatarDataNotify, userId, clientSeq, g.PacketAvatarDataNotify(player))
	g.SendMsg(cmd.OpenStateUpdateNotify, userId, clientSeq, g.PacketOpenStateUpdateNotify(player))
	g.SendMsg(cmd.QuestListNotify, userId, clientSeq, g.PacketQuestListNotify(player))
	g.SendMsg(cmd.FinishedParentQuestNotify, userId, clientSeq, g.PacketFinishedParentQuestNotify(player))
	g.SendMsg(cmd.AllMarkPointNotify, player.PlayerId, player.ClientSeq, &proto.AllMarkPointNotify{MarkList: g.PacketMapMarkPointList(player)})
	g.GCGLogin(player) // 发送GCG登录相关的通知包
}
