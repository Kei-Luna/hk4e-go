package game

import (
	"encoding/base64"

	"hk4e/common/constant"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/pkg/random"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

// GM函数模块
// GM函数只支持基本类型的简单参数传入

type GMCmd struct {
}

// 玩家通用GM指令

// GMTeleportPlayer 传送玩家
func (g *GMCmd) GMTeleportPlayer(userId, sceneId uint32, posX, posY, posZ float64) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		logger.Error("world is nil, worldId: %v, uid: %v", player.WorldId, player.PlayerId)
		return
	}

	// ai世界禁止用命令传送
	if WORLD_MANAGER.IsAiWorld(world) && player.CmdPerm != uint8(CommandPermGM) {
		return
	}

	GAME.TeleportPlayer(
		player,
		proto.EnterReason_ENTER_REASON_GM,
		sceneId,
		&model.Vector{X: posX, Y: posY, Z: posZ},
		new(model.Vector),
		0,
		0,
	)
}

// GMAddItem 给予玩家物品
func (g *GMCmd) GMAddItem(userId, itemId, itemCount uint32) {
	GAME.AddPlayerItem(userId, []*ChangeItem{
		{
			ItemId:      itemId,
			ChangeCount: itemCount,
		},
	}, true, 0)
}

// GMAddWeapon 给予玩家武器
func (g *GMCmd) GMAddWeapon(userId, itemId, itemCount uint32, level, promote, refinement uint8) {
	// 武器数量
	for i := uint32(0); i < itemCount; i++ {
		// 给予武器
		weaponId := GAME.AddPlayerWeapon(userId, itemId)
		// 获取玩家
		player := USER_MANAGER.GetOnlineUser(userId)
		if player == nil {
			logger.Error("player is nil, uid: %v", userId)
			return
		}
		// 获取武器
		weapon := player.GetDbWeapon().GetWeapon(weaponId)
		if weapon == nil {
			logger.Error("weapon is nil, weaponId: %v", weaponId)
			return
		}
		// 设置武器的突破等级
		weapon.Promote = promote
		// 设置武器等级
		weapon.Level = level
		weapon.Exp = 0
		// 设置武器精炼
		weapon.Refinement = refinement
		// 更新武器的物品数据
		GAME.SendMsg(cmd.StoreItemChangeNotify, player.PlayerId, player.ClientSeq, GAME.PacketStoreItemChangeNotifyByWeapon(weapon))
	}
}

// GMAddReliquary 给予玩家圣遗物
func (g *GMCmd) GMAddReliquary(userId, itemId, itemCount uint32) {
	// 圣遗物数量
	for i := uint32(0); i < itemCount; i++ {
		// 给予圣遗物
		GAME.AddPlayerReliquary(userId, itemId)
	}
}

// GMAddAvatar 给予玩家角色
func (g *GMCmd) GMAddAvatar(userId, avatarId uint32, level, promote uint8) {
	// 添加角色
	GAME.AddPlayerAvatar(userId, avatarId)
	// 获取玩家
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	// 获取角色
	avatar, ok := player.GetDbAvatar().AvatarMap[avatarId]
	if !ok {
		logger.Error("avatar not exist, avatarId: %v", avatarId)
		return
	}
	// 设置角色的突破等级
	avatar.Promote = promote
	// 设置角色的等级
	avatar.Level = level
	avatar.Exp = 0
	// 角色更新面板
	GAME.UpdatePlayerAvatarFightProp(player.PlayerId, avatar.AvatarId)
	// 角色属性表更新通知
	GAME.SendMsg(cmd.AvatarPropNotify, player.PlayerId, player.ClientSeq, GAME.PacketAvatarPropNotify(avatar))
}

// GMAddCostume 给予玩家时装
func (g *GMCmd) GMAddCostume(userId, costumeId uint32) {
	// 添加时装
	GAME.AddPlayerCostume(userId, costumeId)
}

// GMAddFlycloak 给予玩家风之翼
func (g *GMCmd) GMAddFlycloak(userId, flycloakId uint32) {
	// 添加风之翼
	GAME.AddPlayerFlycloak(userId, flycloakId)
}

// GMAddAllItem 给予玩家所有物品
func (g *GMCmd) GMAddAllItem(userId, itemCount uint32) {
	GAME.LogoutPlayer(userId)
	itemList := make([]*ChangeItem, 0)
	for itemId := range GAME.GetAllItemDataConfig() {
		itemList = append(itemList, &ChangeItem{
			ItemId:      uint32(itemId),
			ChangeCount: itemCount,
		})
	}
	GAME.AddPlayerItem(userId, itemList, false, 0)
}

// GMAddAllWeapon 给予玩家所有武器
func (g *GMCmd) GMAddAllWeapon(userId, itemCount uint32, level, promote, refinement uint8) {
	for itemId := range GAME.GetAllWeaponDataConfig() {
		g.GMAddWeapon(userId, uint32(itemId), itemCount, level, promote, refinement)
	}
}

// GMAddAllReliquary 给予玩家所有圣遗物
func (g *GMCmd) GMAddAllReliquary(userId, itemCount uint32) {
	GAME.LogoutPlayer(userId)
	for itemId := range GAME.GetAllReliquaryDataConfig() {
		g.GMAddReliquary(userId, uint32(itemId), itemCount)
	}
}

// GMAddAllAvatar 给予玩家所有角色
func (g *GMCmd) GMAddAllAvatar(userId uint32, level, promote uint8) {
	for avatarId := range GAME.GetAllAvatarDataConfig() {
		g.GMAddAvatar(userId, uint32(avatarId), level, promote)
	}
}

// GMAddAllCostume 给予玩家所有时装
func (g *GMCmd) GMAddAllCostume(userId uint32) {
	for costumeId := range gdconf.GetAvatarCostumeDataMap() {
		g.GMAddCostume(userId, uint32(costumeId))
	}
}

// GMAddAllFlycloak 给予玩家所有风之翼
func (g *GMCmd) GMAddAllFlycloak(userId uint32) {
	for flycloakId := range gdconf.GetAvatarFlycloakDataMap() {
		g.GMAddFlycloak(userId, uint32(flycloakId))
	}
}

// GMAddAll 给予玩家所有内容
func (g *GMCmd) GMAddAll(userId uint32) {
	GAME.LogoutPlayer(userId)
	// 给予玩家所有物品
	g.GMAddAllItem(userId, 9999)
	// 给予玩家所有武器
	g.GMAddAllWeapon(userId, 1, 90, 5, 5)
	// 给予玩家所有圣遗物
	g.GMAddAllReliquary(userId, 5)
	// 给予玩家所有角色
	g.GMAddAllAvatar(userId, 90, 5)
	// 给予玩家所有时装
	g.GMAddAllCostume(userId)
	// 给予玩家所有风之翼
	g.GMAddAllFlycloak(userId)
}

// GMKillSelf 杀死自己
func (g *GMCmd) GMKillSelf(userId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	// 激活的角色id
	activeAvatarId := player.GetDbTeam().GetActiveAvatarId()
	// 杀死角色
	GAME.KillPlayerAvatar(player, activeAvatarId, proto.PlayerDieType_PLAYER_DIE_GM)
}

// GMKillMonster 杀死某个怪物
func (g *GMCmd) GMKillMonster(userId uint32, entityId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		logger.Error("world is nil, worldId: %v, uid: %v", player.WorldId, player.PlayerId)
		return
	}
	scene := world.GetSceneById(player.SceneId)
	if scene == nil {
		logger.Error("scene is nil, sceneId: %v, uid: %v", player.SceneId, player.PlayerId)
		return
	}
	// 获取实体
	entity := scene.GetEntity(entityId)
	if entity == nil {
		logger.Error("entity is nil, entityId: %v, uid: %v", entityId, player.PlayerId)
		return
	}
	// 确保为怪物
	if entity.GetEntityType() != constant.ENTITY_TYPE_MONSTER {
		return
	}
	// 杀死怪物
	GAME.KillEntity(player, scene, entity.GetId(), proto.PlayerDieType_PLAYER_DIE_GM)
}

// GMKillAllMonster 杀死所有怪物
func (g *GMCmd) GMKillAllMonster(userId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		logger.Error("world is nil, worldId: %v, uid: %v", player.WorldId, player.PlayerId)
		return
	}
	scene := world.GetSceneById(player.SceneId)
	if scene == nil {
		logger.Error("scene is nil, sceneId: %v, uid: %v", player.SceneId, player.PlayerId)
		return
	}
	// 杀死所有怪物实体
	for _, entity := range scene.GetAllEntity() {
		// 确保为怪物
		if entity.GetEntityType() != constant.ENTITY_TYPE_MONSTER {
			continue
		}
		// 杀死怪物
		GAME.KillEntity(player, scene, entity.GetId(), proto.PlayerDieType_PLAYER_DIE_GM)
	}
}

// GMAddQuest 添加任务
func (g *GMCmd) GMAddQuest(userId uint32, questId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	dbQuest := player.GetDbQuest()
	dbQuest.AddQuest(questId)
	dbQuest.StartQuest(questId)
	ntf := &proto.QuestListUpdateNotify{
		QuestList: make([]*proto.Quest, 0),
	}
	ntf.QuestList = append(ntf.QuestList, GAME.PacketQuest(player, questId))
	GAME.SendMsg(cmd.QuestListUpdateNotify, player.PlayerId, player.ClientSeq, ntf)
}

// GMFinishQuest 完成任务
func (g *GMCmd) GMFinishQuest(userId uint32, questId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	dbQuest := player.GetDbQuest()
	dbQuest.ForceFinishQuest(questId)
	ntf := &proto.QuestListUpdateNotify{
		QuestList: make([]*proto.Quest, 0),
	}
	ntf.QuestList = append(ntf.QuestList, GAME.PacketQuest(player, questId))
	GAME.SendMsg(cmd.QuestListUpdateNotify, player.PlayerId, player.ClientSeq, ntf)
	GAME.AcceptQuest(player, true)
}

// GMForceFinishAllQuest 强制完成当前所有任务
func (g *GMCmd) GMForceFinishAllQuest(userId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	dbQuest := player.GetDbQuest()
	ntf := &proto.QuestListUpdateNotify{
		QuestList: make([]*proto.Quest, 0),
	}
	for _, quest := range dbQuest.GetQuestMap() {
		dbQuest.ForceFinishQuest(quest.QuestId)
		pbQuest := GAME.PacketQuest(player, quest.QuestId)
		if pbQuest == nil {
			continue
		}
		ntf.QuestList = append(ntf.QuestList, pbQuest)
	}
	GAME.SendMsg(cmd.QuestListUpdateNotify, player.PlayerId, player.ClientSeq, ntf)
}

// GMUnlockPoint 解锁场景某个锚点
func (g *GMCmd) GMUnlockPoint(userId uint32, sceneId uint32, pointId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	dbWorld := player.GetDbWorld()
	dbScene := dbWorld.GetSceneById(sceneId)
	if dbScene == nil {
		logger.Error("db scene is nil, uid: %v", sceneId)
		return
	}
	scenePointMapConfig := gdconf.GetScenePointMapBySceneId(int32(sceneId))
	if scenePointMapConfig == nil {
		logger.Error("scene point config is nil, sceneId: %v", sceneId)
		return
	}
	// 解锁锚点
	dbScene.UnlockPoint(pointId)
	GAME.SendMsg(cmd.ScenePointUnlockNotify, player.PlayerId, player.ClientSeq, &proto.ScenePointUnlockNotify{
		SceneId:   sceneId,
		PointList: []uint32{pointId},
	})
}

// GMUnlockAllPoint 解锁场景全部锚点
func (g *GMCmd) GMUnlockAllPoint(userId uint32, sceneId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	dbWorld := player.GetDbWorld()
	dbScene := dbWorld.GetSceneById(sceneId)
	if dbScene == nil {
		logger.Error("db scene is nil, sceneId: %v, uid: %v", sceneId, userId)
		return
	}
	scenePointMapConfig := gdconf.GetScenePointMapBySceneId(int32(sceneId))
	if scenePointMapConfig == nil {
		logger.Error("scene point config is nil, sceneId: %v", sceneId)
		return
	}
	for _, pointData := range scenePointMapConfig {
		dbScene.UnlockPoint(uint32(pointData.Id))
	}
	GAME.SendMsg(cmd.ScenePointUnlockNotify, player.PlayerId, player.ClientSeq, &proto.ScenePointUnlockNotify{
		SceneId:   sceneId,
		PointList: dbScene.GetUnlockPointList(),
	})
}

// GMSetWeather 设置天气
func (g *GMCmd) GMSetWeather(userId uint32, weatherAreaId uint32, climateType uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	GAME.SetPlayerWeather(player, weatherAreaId, climateType)
}

// GMCreateMonster 在玩家附近创建怪物
func (g *GMCmd) GMCreateMonster(userId uint32, monsterId uint32, posX, posY, posZ float64, count uint32, level uint8) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		logger.Error("world is nil, worldId: %v, uid: %v", player.WorldId, player.PlayerId)
		return
	}
	scene := world.GetSceneById(player.SceneId)
	if scene == nil {
		logger.Error("scene is nil, sceneId: %v, uid: %v", player.SceneId, player.PlayerId)
		return
	}
	if count > 100 {
		logger.Error("monster count too large, uid: %v", userId)
		return
	}
	for i := 0; i < int(count); i++ {
		GAME.CreateMonster(player, &model.Vector{
			X: posX,
			Y: posY,
			Z: posZ,
		}, monsterId, level)
	}
}

// GMCreateGadget 在玩家附近创建物件
func (g *GMCmd) GMCreateGadget(userId uint32, gadgetId uint32, count uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	if count > 100 {
		logger.Error("gadget count too large, uid: %v", userId)
		return
	}
	for i := 0; i < int(count); i++ {
		GAME.CreateGadget(player, nil, gadgetId, nil)
	}
}

// GMClearPlayer 清除账号数据
func (g *GMCmd) GMClearPlayer(userId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	player.OfflineClear = true
	GAME.LogoutPlayer(userId)
}

// GMClearItem 清除道具
func (g *GMCmd) GMClearItem(userId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	player.DbItem = nil
	GAME.LogoutPlayer(userId)
}

// GMClearQuest 清除任务
func (g *GMCmd) GMClearQuest(userId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	player.DbQuest = nil
	GAME.LogoutPlayer(userId)
}

// GMClearWorld 清除大世界数据
func (g *GMCmd) GMClearWorld(userId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	player.DbWorld = nil
	GAME.LogoutPlayer(userId)
}

// 系统级GM指令

// TODO 不知道为什么0个参数的函数会反射调用失败

func (g *GMCmd) ChangePlayerCmdPerm(userId uint32, cmdPerm uint8) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	player.CmdPerm = cmdPerm
}

func (g *GMCmd) ReloadGameDataConfig(reloadSceneLua bool) {
	LOCAL_EVENT_MANAGER.GetLocalEventChan() <- &LocalEvent{
		EventId: ReloadGameDataConfig,
		Msg:     reloadSceneLua,
	}
}

func (g *GMCmd) XLuaDebug(userId uint32, luacBase64 string) {
	logger.Debug("xlua debug, uid: %v, luac: %v", userId, luacBase64)
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	// 只有在线玩家主动开启之后才能发送
	if !player.XLuaDebug {
		logger.Error("player xlua debug not enable, uid: %v", userId)
		return
	}
	luac, err := base64.StdEncoding.DecodeString(luacBase64)
	if err != nil {
		logger.Error("decode luac error: %v", err)
		return
	}
	GAME.SendMsg(cmd.PlayerLuaShellNotify, player.PlayerId, 0, &proto.PlayerLuaShellNotify{
		ShellType: proto.LuaShellType_LUASHELL_NORMAL,
		Id:        1,
		LuaShell:  luac,
		UseType:   1,
	})
}

func (g *GMCmd) AvPlayAudio(fileDataBase64 string) {
	fileData, err := base64.StdEncoding.DecodeString(fileDataBase64)
	if err != nil {
		logger.Error("file data base64 format error: %v", err)
		return
	}
	go PlayAudio(fileData)
}

func (g *GMCmd) AvUpdateFrame(fileDataBase64 string, rgb bool, posX, posY, posZ float64) {
	fileData, err := base64.StdEncoding.DecodeString(fileDataBase64)
	if err != nil {
		logger.Error("file data base64 format error: %v", err)
		return
	}
	basePos := &model.Vector{X: posX, Y: posY, Z: posZ}
	if basePos.X == 0.0 && basePos.Y == 0.0 && basePos.Z == 0.0 {
		basePos = &model.Vector{X: 2700, Y: 200, Z: -1800}
	}
	UpdateFrame(fileData, basePos, rgb)
}

func (g *GMCmd) CreateRobotInAiWorld(aiUid uint32, name string, avatarId uint32, posX, posY, posZ float64) {
	if name == "" {
		name = random.GetRandomStr(8)
	}
	if avatarId == 0 {
		for _, avatarData := range gdconf.GetAvatarDataMap() {
			avatarId = uint32(avatarData.AvatarId)
			break
		}
	}
	ai := USER_MANAGER.GetOnlineUser(aiUid)
	if ai == nil {
		logger.Error("ai is nil, uid: %v", aiUid)
		return
	}
	world := WORLD_MANAGER.GetWorldById(ai.WorldId)
	robot := AI_MANAGER.CreateAi(name, name, 10000007, 210001, "")
	GAME.AddPlayerAvatar(robot.PlayerId, avatarId)
	dbAvatar := robot.GetDbAvatar()
	GAME.SetUpAvatarTeamReq(robot, &proto.SetUpAvatarTeamReq{
		TeamId:             1,
		AvatarTeamGuidList: []uint64{dbAvatar.AvatarMap[avatarId].Guid},
		CurAvatarGuid:      dbAvatar.AvatarMap[avatarId].Guid,
	})
	GAME.SetPlayerHeadImageReq(robot, &proto.SetPlayerHeadImageReq{
		AvatarId: avatarId,
	})
	GAME.JoinPlayerSceneReq(robot, &proto.JoinPlayerSceneReq{
		TargetUid: world.GetOwner().PlayerId,
	})
	GAME.EnterSceneReadyReq(robot, &proto.EnterSceneReadyReq{
		EnterSceneToken: world.GetEnterSceneToken(),
	})
	GAME.SceneInitFinishReq(robot, &proto.SceneInitFinishReq{
		EnterSceneToken: world.GetEnterSceneToken(),
	})
	GAME.EnterSceneDoneReq(robot, &proto.EnterSceneDoneReq{
		EnterSceneToken: world.GetEnterSceneToken(),
	})
	GAME.PostEnterSceneReq(robot, &proto.PostEnterSceneReq{
		EnterSceneToken: world.GetEnterSceneToken(),
	})
	activeAvatarId := world.GetPlayerActiveAvatarId(robot)
	entityMoveInfo := &proto.EntityMoveInfo{
		EntityId: world.GetPlayerWorldAvatarEntityId(robot, activeAvatarId),
		MotionInfo: &proto.MotionInfo{
			Pos:   &proto.Vector{X: float32(posX), Y: float32(posY), Z: float32(posZ)},
			Rot:   &proto.Vector{X: float32(0.0), Y: float32(0.0), Z: float32(0.0)},
			State: proto.MotionState_MOTION_STANDBY,
		},
		SceneTime:   0,
		ReliableSeq: 0,
	}
	combatData, err := pb.Marshal(entityMoveInfo)
	if err != nil {
		return
	}
	GAME.CombatInvocationsNotify(robot, &proto.CombatInvocationsNotify{
		InvokeList: []*proto.CombatInvokeEntry{{
			CombatData:   combatData,
			ForwardType:  proto.ForwardType_FORWARD_TO_ALL_EXCEPT_CUR,
			ArgumentType: proto.CombatTypeArgument_ENTITY_MOVE,
		}},
	})
	GAME.UnionCmdNotify(robot, &proto.UnionCmdNotify{})
}

func (g *GMCmd) ServerAnnounce(announceId uint32, announceMsg string, isRevoke bool) {
	if !isRevoke {
		GAME.ServerAnnounceNotify(announceId, announceMsg)
	} else {
		GAME.ServerAnnounceRevokeNotify(announceId)
	}
}

func (g *GMCmd) SendMsgToPlayer(cmdName string, userId uint32, msgJson string) {
	if cmdProtoMap == nil {
		cmdProtoMap = cmd.NewCmdProtoMap()
	}
	cmdId := cmdProtoMap.GetCmdIdByCmdName(cmdName)
	if cmdId == 0 {
		logger.Error("cmd name not found")
		return
	}
	if cmdId == cmd.WindSeedClientNotify || cmdId == cmd.PlayerLuaShellNotify {
		logger.Error("what are you doing ???")
		return
	}
	msg := cmdProtoMap.GetProtoObjByCmdId(cmdId)
	err := protojson.Unmarshal([]byte(msgJson), msg)
	if err != nil {
		logger.Error("parse msg error: %v", err)
		return
	}
	GAME.SendMsg(cmdId, userId, 0, msg)
}

func (g *GMCmd) SetPhysicsEngineParam(aiUid uint32, pathTracing bool, acc float32, drag float32, pao float32, is float32, ayo float32) {
	ai := USER_MANAGER.GetOnlineUser(aiUid)
	if ai == nil {
		logger.Error("ai is nil, uid: %v", aiUid)
		return
	}
	world := WORLD_MANAGER.GetWorldById(ai.WorldId)
	engine := world.GetBulletPhysicsEngine()
	engine.SetPhysicsEngineParam(pathTracing, acc, drag, pao, is, ayo)
}

func (g *GMCmd) ShowAvatarCollider(aiUid uint32) {
	ai := USER_MANAGER.GetOnlineUser(aiUid)
	if ai == nil {
		logger.Error("ai is nil, uid: %v", aiUid)
		return
	}
	world := WORLD_MANAGER.GetWorldById(ai.WorldId)
	engine := world.GetBulletPhysicsEngine()
	engine.ShowAvatarCollider()
}

func (g *GMCmd) AiWorldAoiDebug(aiUid uint32) {
	ai := USER_MANAGER.GetOnlineUser(aiUid)
	if ai == nil {
		logger.Error("ai is nil, uid: %v", aiUid)
		return
	}
	world := WORLD_MANAGER.GetWorldById(ai.WorldId)
	scene := world.GetSceneById(world.GetOwner().SceneId)
	aiWorldAoi := world.GetAiWorldAoi()
	gridMap := aiWorldAoi.Debug()
	logger.Debug("total grid num: %v", len(gridMap))
	for _, grid := range gridMap {
		objectMap := grid.GetObjectList()
		if len(objectMap) == 0 {
			continue
		}
		logger.Debug("================================================== GRID gid:%v ==================================================", grid.GetGid())
		for objectId, object := range objectMap {
			wa := object.(*WorldAvatar)
			var pos *model.Vector = nil
			entity := scene.GetEntity(wa.avatarEntityId)
			if entity != nil {
				pos = entity.GetPos()
			}
			logger.Debug("uid: %v, wa.uid: %v, wa.avatarId: %v, wa.entityId: %v, pos: %+v", objectId, wa.uid, wa.avatarId, wa.avatarEntityId, pos)
		}
	}
}

func (g *GMCmd) GetPlayerData(userId uint32) *model.Player {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return nil
	}
	return player
}

func (g *GMCmd) GetPlayerPos(userId uint32) *model.Vector {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return nil
	}
	return player.Pos
}

func (g *GMCmd) NotSavePlayer(userId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	player.OfflineNotSave = true
}
