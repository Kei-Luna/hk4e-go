package game

import (
	"hk4e/common/constant"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/endec"
	"hk4e/pkg/logger"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

/************************************************** 接口请求 **************************************************/

func (g *Game) ChangeAvatarReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ChangeAvatarReq)
	targetAvatar, ok := player.GameObjectGuidMap[req.Guid].(*model.Avatar)
	if !ok {
		logger.Error("target avatar error, avatarGuid: %v", req.Guid)
		g.SendError(cmd.ChangeAvatarRsp, player, &proto.ChangeAvatarRsp{})
		return
	}

	g.ChangeAvatar(player, targetAvatar.AvatarId)

	rsp := &proto.ChangeAvatarRsp{
		CurGuid: req.Guid,
	}
	g.SendMsg(cmd.ChangeAvatarRsp, player.PlayerId, player.ClientSeq, rsp)
}

func (g *Game) SetUpAvatarTeamReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetUpAvatarTeamReq)
	teamId := req.TeamId
	avatarGuidList := req.AvatarTeamGuidList
	currAvatarGuid := req.CurAvatarGuid

	if teamId <= 0 || teamId >= 5 {
		g.SendError(cmd.SetUpAvatarTeamRsp, player, &proto.SetUpAvatarTeamRsp{})
		return
	}
	dbTeam := player.GetDbTeam()
	if (teamId == uint32(dbTeam.GetActiveTeamId()) && len(avatarGuidList) == 0) || len(avatarGuidList) > 4 {
		g.SendError(cmd.SetUpAvatarTeamRsp, player, &proto.SetUpAvatarTeamRsp{})
		return
	}
	avatarIdList := make([]uint32, 0)
	for _, avatarGuid := range avatarGuidList {
		avatar, ok := player.GameObjectGuidMap[avatarGuid].(*model.Avatar)
		if !ok {
			g.SendError(cmd.SetUpAvatarTeamRsp, player, &proto.SetUpAvatarTeamRsp{})
			return
		}
		avatarIdList = append(avatarIdList, avatar.AvatarId)
	}
	currAvatar, ok := player.GameObjectGuidMap[currAvatarGuid].(*model.Avatar)
	if !ok {
		logger.Error("get curr avatar error, avatarGuid: %v", currAvatarGuid)
		g.SendError(cmd.SetUpAvatarTeamRsp, player, &proto.SetUpAvatarTeamRsp{})
		return
	}

	g.ChangeTeam(player, teamId, avatarIdList, currAvatar.AvatarId)

	rsp := &proto.SetUpAvatarTeamRsp{
		TeamId:             req.TeamId,
		CurAvatarGuid:      req.CurAvatarGuid,
		AvatarTeamGuidList: req.AvatarTeamGuidList,
	}
	g.SendMsg(cmd.SetUpAvatarTeamRsp, player.PlayerId, player.ClientSeq, rsp)
}

func (g *Game) ChooseCurAvatarTeamReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ChooseCurAvatarTeamReq)
	teamId := req.TeamId
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		logger.Error("get world is nil, worldId: %v, uid: %v", player.WorldId, player.PlayerId)
		return
	}
	if world.IsMultiplayerWorld() {
		g.SendError(cmd.ChooseCurAvatarTeamRsp, player, &proto.ChooseCurAvatarTeamRsp{})
		return
	}
	dbTeam := player.GetDbTeam()
	team := dbTeam.GetTeamByIndex(uint8(teamId) - 1)
	if team == nil || len(team.GetAvatarIdList()) == 0 {
		return
	}
	dbTeam.CurrTeamIndex = uint8(teamId) - 1
	dbTeam.CurrAvatarIndex = 0
	world.SetPlayerAvatarIndex(player, 0)
	world.SetPlayerLocalTeam(player, team.GetAvatarIdList())
	world.UpdateMultiplayerTeam()
	world.UpdatePlayerWorldAvatar(player)

	sceneTeamUpdateNotify := g.PacketSceneTeamUpdateNotify(world, player)
	g.SendMsg(cmd.SceneTeamUpdateNotify, player.PlayerId, player.ClientSeq, sceneTeamUpdateNotify)

	chooseCurAvatarTeamRsp := &proto.ChooseCurAvatarTeamRsp{
		CurTeamId: teamId,
	}
	g.SendMsg(cmd.ChooseCurAvatarTeamRsp, player.PlayerId, player.ClientSeq, chooseCurAvatarTeamRsp)
}

func (g *Game) ChangeMpTeamAvatarReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ChangeMpTeamAvatarReq)
	avatarGuidList := req.AvatarGuidList
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		logger.Error("get world is nil, worldId: %v, uid: %v", player.WorldId, player.PlayerId)
		return
	}
	if WORLD_MANAGER.IsAiWorld(world) || !world.IsMultiplayerWorld() || len(avatarGuidList) == 0 || len(avatarGuidList) > 4 {
		g.SendError(cmd.ChangeMpTeamAvatarRsp, player, &proto.ChangeMpTeamAvatarRsp{})
		return
	}
	avatarIdList := make([]uint32, 0)
	for _, avatarGuid := range avatarGuidList {
		avatar, ok := player.GameObjectGuidMap[avatarGuid].(*model.Avatar)
		if !ok {
			logger.Error("avatar error, avatarGuid: %v", avatarGuid)
			return
		}
		avatarId := avatar.AvatarId
		avatarIdList = append(avatarIdList, avatarId)
	}
	world.SetPlayerLocalTeam(player, avatarIdList)
	world.UpdateMultiplayerTeam()
	world.UpdatePlayerWorldAvatar(player)

	currAvatarGuid := req.CurAvatarGuid
	currAvatar, ok := player.GameObjectGuidMap[currAvatarGuid].(*model.Avatar)
	if !ok {
		logger.Error("avatar error, avatarGuid: %v", currAvatarGuid)
		return
	}
	currAvatarId := currAvatar.AvatarId
	newAvatarIndex := world.GetPlayerAvatarIndexByAvatarId(player, currAvatarId)
	world.SetPlayerAvatarIndex(player, newAvatarIndex)

	sceneTeamUpdateNotify := g.PacketSceneTeamUpdateNotify(world, player)
	g.SendToWorldA(world, cmd.SceneTeamUpdateNotify, player.ClientSeq, sceneTeamUpdateNotify, 0)

	changeMpTeamAvatarRsp := &proto.ChangeMpTeamAvatarRsp{
		CurAvatarGuid:  req.CurAvatarGuid,
		AvatarGuidList: req.AvatarGuidList,
	}
	g.SendMsg(cmd.ChangeMpTeamAvatarRsp, player.PlayerId, player.ClientSeq, changeMpTeamAvatarRsp)
}

func (g *Game) AvatarDieAnimationEndReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.AvatarDieAnimationEndReq)

	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)

	// 触发事件
	if PLUGIN_MANAGER.TriggerEvent(PluginEventIdAvatarDieAnimationEnd, &PluginEventAvatarDieAnimationEnd{
		PluginEvent: NewPluginEvent(),
		Player:      player,
		Req:         req,
	}) {
		return
	}

	entity := scene.GetEntity(uint32(req.DieGuid))
	if entity.GetLastDieType() == int32(proto.PlayerDieType_PLAYER_DIE_DRAWN) {
		maxStamina := player.PropMap[constant.PLAYER_PROP_MAX_STAMINA]
		// 设置玩家耐力为一半
		g.SetPlayerStamina(player, maxStamina/2)
		// 传送玩家至安全位置
		g.TeleportPlayer(
			player,
			proto.EnterReason_ENTER_REASON_REVIVAL,
			player.SceneId,
			player.GetPos(),
			player.GetRot(),
			0,
			0,
		)
	} else {
		targetAvatarId := uint32(0)
		for _, worldAvatar := range world.GetPlayerWorldAvatarList(player) {
			dbAvatar := player.GetDbAvatar()
			avatar, exist := dbAvatar.AvatarMap[worldAvatar.GetAvatarId()]
			if !exist {
				logger.Error("get db avatar is nil, avatarId: %v", worldAvatar.GetAvatarId())
				continue
			}
			if avatar.LifeState != constant.LIFE_STATE_ALIVE {
				continue
			}
			targetAvatarId = worldAvatar.GetAvatarId()
		}
		if targetAvatarId == 0 {
			g.SendMsg(cmd.WorldPlayerDieNotify, player.PlayerId, player.ClientSeq, &proto.WorldPlayerDieNotify{
				DieType: proto.PlayerDieType(entity.GetLastDieType()),
			})
		} else {
			g.ChangeAvatar(player, targetAvatarId)
		}
	}

	g.SendMsg(cmd.AvatarDieAnimationEndRsp, player.PlayerId, player.ClientSeq, &proto.AvatarDieAnimationEndRsp{SkillId: req.SkillId, DieGuid: req.DieGuid})
}

func (g *Game) WorldPlayerReviveReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.WorldPlayerReviveReq)
	_ = req
	world := WORLD_MANAGER.GetWorldById(player.WorldId)

	if WORLD_MANAGER.IsAiWorld(world) {
		GAME.ReLoginPlayer(player.PlayerId, true)
		return
	}

	g.TeleportPlayer(
		player,
		proto.EnterReason_ENTER_REASON_REVIVAL,
		player.SceneId,
		player.GetPos(),
		player.GetRot(),
		0,
		0,
	)
	g.SendMsg(cmd.WorldPlayerReviveRsp, player.PlayerId, player.ClientSeq, new(proto.WorldPlayerReviveRsp))
}

/************************************************** 游戏功能 **************************************************/

func (g *Game) ChangeAvatar(player *model.Player, targetAvatarId uint32) {
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		logger.Error("get world is nil, worldId: %v, uid: %v", player.WorldId, player.PlayerId)
		return
	}
	scene := world.GetSceneById(player.SceneId)
	oldAvatarId := world.GetPlayerActiveAvatarId(player)
	if targetAvatarId == oldAvatarId {
		logger.Error("can not change to the same avatar, uid: %v, oldAvatarId: %v, targetAvatarId: %v", player.PlayerId, oldAvatarId, targetAvatarId)
		return
	}
	newAvatarIndex := world.GetPlayerAvatarIndexByAvatarId(player, targetAvatarId)
	if newAvatarIndex == -1 {
		logger.Error("can not find the target avatar in team, uid: %v, targetAvatarId: %v", player.PlayerId, targetAvatarId)
		return
	}
	if !world.IsMultiplayerWorld() {
		dbTeam := player.GetDbTeam()
		dbTeam.CurrAvatarIndex = uint8(newAvatarIndex)
	}
	world.SetPlayerAvatarIndex(player, newAvatarIndex)
	oldAvatarEntityId := world.GetPlayerWorldAvatarEntityId(player, oldAvatarId)
	oldAvatarEntity := scene.GetEntity(oldAvatarEntityId)
	if oldAvatarEntity == nil {
		logger.Error("can not find old avatar entity, entity id: %v", oldAvatarEntityId)
		return
	}
	oldAvatarEntity.SetMoveState(uint16(proto.MotionState_MOTION_STANDBY))

	sceneEntityDisappearNotify := &proto.SceneEntityDisappearNotify{
		DisappearType: proto.VisionType_VISION_REPLACE,
		EntityList:    []uint32{oldAvatarEntity.GetId()},
	}
	g.SendToSceneA(scene, cmd.SceneEntityDisappearNotify, player.ClientSeq, sceneEntityDisappearNotify, 0)

	newAvatarEntity := g.PacketSceneEntityInfoAvatar(scene, player, targetAvatarId)
	sceneEntityAppearNotify := &proto.SceneEntityAppearNotify{
		AppearType: proto.VisionType_VISION_REPLACE,
		Param:      oldAvatarEntity.GetId(),
		EntityList: []*proto.SceneEntityInfo{newAvatarEntity},
	}
	g.SendToSceneA(scene, cmd.SceneEntityAppearNotify, player.ClientSeq, sceneEntityAppearNotify, 0)
}

func (g *Game) ChangeTeam(player *model.Player, teamId uint32, avatarIdList []uint32, currAvatarId uint32) {
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		logger.Error("get world is nil, worldId: %v, uid: %v", player.WorldId, player.PlayerId)
		return
	}
	if world.IsMultiplayerWorld() {
		return
	}

	dbTeam := player.GetDbTeam()
	dbTeam.GetTeamByIndex(uint8(teamId - 1)).SetAvatarIdList(avatarIdList)

	avatarTeamUpdateNotify := &proto.AvatarTeamUpdateNotify{
		AvatarTeamMap: make(map[uint32]*proto.AvatarTeam),
	}
	dbAvatar := player.GetDbAvatar()
	for teamIndex, team := range dbTeam.TeamList {
		avatarTeam := &proto.AvatarTeam{
			TeamName:       team.Name,
			AvatarGuidList: make([]uint64, 0),
		}
		for _, avatarId := range team.GetAvatarIdList() {
			avatarTeam.AvatarGuidList = append(avatarTeam.AvatarGuidList, dbAvatar.AvatarMap[avatarId].Guid)
		}
		avatarTeamUpdateNotify.AvatarTeamMap[uint32(teamIndex)+1] = avatarTeam
	}
	g.SendMsg(cmd.AvatarTeamUpdateNotify, player.PlayerId, player.ClientSeq, avatarTeamUpdateNotify)

	if teamId == uint32(dbTeam.GetActiveTeamId()) {
		world.SetPlayerLocalTeam(player, avatarIdList)
		world.UpdateMultiplayerTeam()
		world.UpdatePlayerWorldAvatar(player)

		currAvatarIndex := world.GetPlayerAvatarIndexByAvatarId(player, currAvatarId)
		dbTeam.CurrAvatarIndex = uint8(currAvatarIndex)
		world.SetPlayerAvatarIndex(player, currAvatarIndex)

		sceneTeamUpdateNotify := g.PacketSceneTeamUpdateNotify(world, player)
		g.SendMsg(cmd.SceneTeamUpdateNotify, player.PlayerId, player.ClientSeq, sceneTeamUpdateNotify)
	}
}

/************************************************** 打包封装 **************************************************/

func (g *Game) PacketSceneTeamUpdateNotify(world *World, player *model.Player) *proto.SceneTeamUpdateNotify {
	sceneTeamUpdateNotify := &proto.SceneTeamUpdateNotify{
		IsInMp: world.IsMultiplayerWorld(),
	}
	empty := new(proto.AbilitySyncStateInfo)
	for _, worldAvatar := range world.GetWorldAvatarList() {
		if WORLD_MANAGER.IsAiWorld(world) && worldAvatar.uid != player.PlayerId {
			continue
		}

		worldPlayer := USER_MANAGER.GetOnlineUser(worldAvatar.GetUid())
		if worldPlayer == nil {
			logger.Error("player is nil, uid: %v", worldAvatar.GetUid())
			continue
		}
		worldPlayerScene := world.GetSceneById(worldPlayer.SceneId)
		worldPlayerDbAvatar := worldPlayer.GetDbAvatar()
		worldPlayerAvatar := worldPlayerDbAvatar.AvatarMap[worldAvatar.GetAvatarId()]
		equipIdList := make([]uint32, 0)
		weapon := worldPlayerAvatar.EquipWeapon
		equipIdList = append(equipIdList, weapon.ItemId)
		for _, reliquary := range worldPlayerAvatar.EquipReliquaryMap {
			equipIdList = append(equipIdList, reliquary.ItemId)
		}
		sceneTeamAvatar := &proto.SceneTeamAvatar{
			PlayerUid:         worldPlayer.PlayerId,
			AvatarGuid:        worldPlayerAvatar.Guid,
			SceneId:           worldPlayer.SceneId,
			EntityId:          world.GetPlayerWorldAvatarEntityId(worldPlayer, worldAvatar.GetAvatarId()),
			SceneEntityInfo:   g.PacketSceneEntityInfoAvatar(worldPlayerScene, worldPlayer, worldAvatar.GetAvatarId()),
			WeaponGuid:        worldPlayerAvatar.EquipWeapon.Guid,
			WeaponEntityId:    world.GetPlayerWorldAvatarWeaponEntityId(worldPlayer, worldAvatar.GetAvatarId()),
			IsPlayerCurAvatar: world.IsPlayerActiveAvatarEntity(worldPlayer, worldAvatar.GetAvatarEntityId()),
			IsOnScene:         world.IsPlayerActiveAvatarEntity(worldPlayer, worldAvatar.GetAvatarEntityId()),
			AvatarAbilityInfo: &proto.AbilitySyncStateInfo{
				IsInited:           len(worldAvatar.GetAbilityList()) != 0,
				DynamicValueMap:    nil,
				AppliedAbilities:   worldAvatar.GetAbilityList(),
				AppliedModifiers:   worldAvatar.GetModifierList(),
				MixinRecoverInfos:  nil,
				SgvDynamicValueMap: nil,
			},
			WeaponAbilityInfo:   empty,
			AbilityControlBlock: g.PacketAbilityControlBlock(worldAvatar.GetAvatarId(), worldPlayerAvatar.SkillDepotId),
		}
		if world.IsMultiplayerWorld() {
			sceneTeamAvatar.AvatarInfo = g.PacketAvatarInfo(worldPlayerAvatar)
			sceneTeamAvatar.SceneAvatarInfo = g.PacketSceneAvatarInfo(worldPlayerScene, worldPlayer, worldAvatar.GetAvatarId())
		}
		sceneTeamUpdateNotify.SceneTeamAvatarList = append(sceneTeamUpdateNotify.SceneTeamAvatarList, sceneTeamAvatar)
	}
	return sceneTeamUpdateNotify
}

// PacketAbilityControlBlock 角色的ability控制块
func (g *Game) PacketAbilityControlBlock(avatarId uint32, skillDepotId uint32) *proto.AbilityControlBlock {
	acb := &proto.AbilityControlBlock{
		AbilityEmbryoList: make([]*proto.AbilityEmbryo, 0),
	}
	abilityId := 0
	// 默认ability
	for _, abilityHashCode := range constant.DEFAULT_ABILITY_HASH_CODE {
		abilityId++
		ae := &proto.AbilityEmbryo{
			AbilityId:               uint32(abilityId),
			AbilityNameHash:         uint32(abilityHashCode),
			AbilityOverrideNameHash: uint32(endec.Hk4eAbilityHashCode("Default")),
		}
		acb.AbilityEmbryoList = append(acb.AbilityEmbryoList, ae)
	}
	// 角色ability
	avatarDataConfig := gdconf.GetAvatarDataById(int32(avatarId))
	if avatarDataConfig != nil {
		for _, abilityHashCode := range avatarDataConfig.AbilityHashCodeList {
			abilityId++
			ae := &proto.AbilityEmbryo{
				AbilityId:               uint32(abilityId),
				AbilityNameHash:         uint32(abilityHashCode),
				AbilityOverrideNameHash: uint32(endec.Hk4eAbilityHashCode("Default")),
			}
			acb.AbilityEmbryoList = append(acb.AbilityEmbryoList, ae)
		}
	}
	// 技能库ability
	skillDepot := gdconf.GetAvatarSkillDepotDataById(int32(skillDepotId))
	if skillDepot != nil && len(skillDepot.AbilityHashCodeList) != 0 {
		for _, abilityHashCode := range skillDepot.AbilityHashCodeList {
			abilityId++
			ae := &proto.AbilityEmbryo{
				AbilityId:               uint32(abilityId),
				AbilityNameHash:         uint32(abilityHashCode),
				AbilityOverrideNameHash: uint32(endec.Hk4eAbilityHashCode("Default")),
			}
			acb.AbilityEmbryoList = append(acb.AbilityEmbryoList, ae)
		}
	}
	// TODO 队伍ability
	// TODO 装备ability
	return acb
}
