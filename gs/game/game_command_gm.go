package game

import (
	"encoding/base64"
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
func (g *GMCmd) GMAddWeapon(userId, itemId, itemCount uint32, level, refinement uint8) {
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
		maxLevel := 90
		maxPromote := 6
		weapon.Promote = level / (uint8(maxLevel / maxPromote))
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
func (g *GMCmd) GMAddAvatar(userId, avatarId uint32, level uint8) {
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
	maxLevel := 90
	maxPromote := 6
	avatar.Promote = level / (uint8(maxLevel / maxPromote))
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
func (g *GMCmd) GMAddAllWeapon(userId, itemCount uint32, level, refinement uint8) {
	for itemId := range GAME.GetAllWeaponDataConfig() {
		g.GMAddWeapon(userId, uint32(itemId), itemCount, level, refinement)
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
func (g *GMCmd) GMAddAllAvatar(userId uint32, level uint8) {
	for avatarId := range GAME.GetAllAvatarDataConfig() {
		g.GMAddAvatar(userId, uint32(avatarId), level)
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
	g.GMAddAllWeapon(userId, 5, 90, 5)
	// 给予玩家所有圣遗物
	g.GMAddAllReliquary(userId, 5)
	// 给予玩家所有角色
	g.GMAddAllAvatar(userId, 90)
	// 给予玩家所有时装
	g.GMAddAllCostume(userId)
	// 给予玩家所有风之翼
	g.GMAddAllFlycloak(userId)
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
	GAME.AcceptQuest(player, true)
}

// GMUnlockAllPoint 解锁场景全部传送点
func (g *GMCmd) GMUnlockAllPoint(userId uint32, sceneId uint32) {
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
	for _, pointData := range scenePointMapConfig {
		dbScene.UnlockPoint(uint32(pointData.Id))
	}
	GAME.SendMsg(cmd.ScenePointUnlockNotify, player.PlayerId, player.ClientSeq, &proto.ScenePointUnlockNotify{
		SceneId:         sceneId,
		PointList:       dbScene.GetUnlockPointList(),
		UnhidePointList: nil,
	})
}

// GMCreateMonster 在玩家附近创建怪物
func (g *GMCmd) GMCreateMonster(userId uint32, monsterId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	GAME.CreateMonster(player, nil, monsterId)
}

// GMCreateGadget 在玩家附近创建物件
func (g *GMCmd) GMCreateGadget(userId uint32, gadgetId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	GAME.CreateGadget(player, nil, gadgetId, nil)
}

// 系统级GM指令

func (g *GMCmd) ChangePlayerCmdPerm(userId uint32, cmdPerm uint8) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	player.CmdPerm = cmdPerm
}

func (g *GMCmd) ReloadGameDataConfig() {
	LOCAL_EVENT_MANAGER.GetLocalEventChan() <- &LocalEvent{
		EventId: ReloadGameDataConfig,
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
	GAME.SendMsg(cmd.WindSeedClientNotify, player.PlayerId, 0, &proto.WindSeedClientNotify{
		Notify: &proto.WindSeedClientNotify_AreaNotify_{
			AreaNotify: &proto.WindSeedClientNotify_AreaNotify{
				AreaCode: luac,
				AreaId:   1,
				AreaType: 1,
			},
		},
	})
}

func (g *GMCmd) PlayAudio(v bool) {
	go PlayAudio()
}

func (g *GMCmd) UpdateFrame(rgb bool) {
	UpdateFrame(rgb)
}

var RobotUidCounter uint32 = 0

func (g *GMCmd) CreateRobotInBigWorld(uid uint32, name string, avatarId uint32) {
	if !GAME.IsMainGs() {
		return
	}
	if uid == 0 {
		RobotUidCounter++
		uid = 1000000 + RobotUidCounter
	}
	if name == "" {
		name = random.GetRandomStr(8)
	}
	if avatarId == 0 {
		for _, avatarData := range gdconf.GetAvatarDataMap() {
			avatarId = uint32(avatarData.AvatarId)
			break
		}
	}
	aiWorld := WORLD_MANAGER.GetAiWorld()
	robot := GAME.CreateRobot(uid, name, name)
	GAME.AddPlayerAvatar(uid, avatarId)
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
		TargetUid: aiWorld.owner.PlayerId,
	})
	GAME.EnterSceneReadyReq(robot, &proto.EnterSceneReadyReq{
		EnterSceneToken: aiWorld.GetEnterSceneToken(),
	})
	GAME.SceneInitFinishReq(robot, &proto.SceneInitFinishReq{
		EnterSceneToken: aiWorld.GetEnterSceneToken(),
	})
	GAME.EnterSceneDoneReq(robot, &proto.EnterSceneDoneReq{
		EnterSceneToken: aiWorld.GetEnterSceneToken(),
	})
	GAME.PostEnterSceneReq(robot, &proto.PostEnterSceneReq{
		EnterSceneToken: aiWorld.GetEnterSceneToken(),
	})
	activeAvatarId := aiWorld.GetPlayerActiveAvatarId(robot)
	pos := new(model.Vector)
	rot := new(model.Vector)
	for _, targetPlayer := range aiWorld.GetAllPlayer() {
		if targetPlayer.PlayerId < PlayerBaseUid {
			continue
		}
		pos = &model.Vector{X: targetPlayer.Pos.X, Y: targetPlayer.Pos.Y, Z: targetPlayer.Pos.Z}
		rot = &model.Vector{X: targetPlayer.Rot.X, Y: targetPlayer.Rot.Y, Z: targetPlayer.Rot.Z}
	}
	entityMoveInfo := &proto.EntityMoveInfo{
		EntityId: aiWorld.GetPlayerWorldAvatarEntityId(robot, activeAvatarId),
		MotionInfo: &proto.MotionInfo{
			Pos:   &proto.Vector{X: float32(pos.X), Y: float32(pos.Y), Z: float32(pos.Z)},
			Rot:   &proto.Vector{X: float32(rot.X), Y: float32(rot.Y), Z: float32(rot.Z)},
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
	if cmdId == cmd.WindSeedClientNotify {
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

func (g *GMCmd) StartPubg(v bool) {
	if world := WORLD_MANAGER.GetAiWorld(); WORLD_MANAGER.IsBigWorld(world) {
		world.StartPubg()
	}
}

func (g *GMCmd) StopPubg(v bool) {
	if world := WORLD_MANAGER.GetAiWorld(); WORLD_MANAGER.IsBigWorld(world) {
		world.StopPubg()
	}
}
