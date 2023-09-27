package game

import (
	"fmt"

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
	targetAvatarGuid := req.Guid
	targetAvatar, ok := player.GameObjectGuidMap[targetAvatarGuid].(*model.Avatar)
	if !ok {
		logger.Error("target avatar error, avatarGuid: %v", targetAvatarGuid)
		return
	}

	g.ChangeAvatar(player, targetAvatar.AvatarId)

	changeAvatarRsp := &proto.ChangeAvatarRsp{
		CurGuid: targetAvatarGuid,
	}
	g.SendMsg(cmd.ChangeAvatarRsp, player.PlayerId, player.ClientSeq, changeAvatarRsp)
}

func (g *Game) SetUpAvatarTeamReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetUpAvatarTeamReq)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		logger.Error("get world is nil, worldId: %v, uid: %v", player.WorldId, player.PlayerId)
		return
	}
	if world.GetMultiplayer() {
		g.SendError(cmd.SetUpAvatarTeamRsp, player, &proto.SetUpAvatarTeamRsp{})
		return
	}
	teamId := req.TeamId
	if teamId <= 0 || teamId >= 5 {
		g.SendError(cmd.SetUpAvatarTeamRsp, player, &proto.SetUpAvatarTeamRsp{})
		return
	}
	avatarGuidList := req.AvatarTeamGuidList
	dbTeam := player.GetDbTeam()
	selfTeam := teamId == uint32(dbTeam.GetActiveTeamId())
	if (selfTeam && len(avatarGuidList) == 0) || len(avatarGuidList) > 4 {
		g.SendError(cmd.SetUpAvatarTeamRsp, player, &proto.SetUpAvatarTeamRsp{})
		return
	}
	avatarIdList := make([]uint32, 0)
	dbAvatar := player.GetDbAvatar()
	for _, avatarGuid := range avatarGuidList {
		for avatarId, avatar := range dbAvatar.AvatarMap {
			if avatarGuid == avatar.Guid {
				avatarIdList = append(avatarIdList, avatarId)
			}
		}
	}
	dbTeam.GetTeamByIndex(uint8(teamId - 1)).SetAvatarIdList(avatarIdList)

	avatarTeamUpdateNotify := &proto.AvatarTeamUpdateNotify{
		AvatarTeamMap: make(map[uint32]*proto.AvatarTeam),
	}
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

	if selfTeam {
		// player.TeamConfig.UpdateTeam()
		world.SetPlayerLocalTeam(player, avatarIdList)
		world.UpdateMultiplayerTeam()
		world.InitPlayerWorldAvatar(player)

		currAvatarGuid := req.CurAvatarGuid
		currAvatar, ok := player.GameObjectGuidMap[currAvatarGuid].(*model.Avatar)
		if !ok {
			logger.Error("avatar error, avatarGuid: %v", currAvatarGuid)
			return
		}
		currAvatarId := currAvatar.AvatarId
		currAvatarIndex := world.GetPlayerAvatarIndexByAvatarId(player, currAvatarId)
		dbTeam.CurrAvatarIndex = uint8(currAvatarIndex)
		world.SetPlayerAvatarIndex(player, currAvatarIndex)

		sceneTeamUpdateNotify := g.PacketSceneTeamUpdateNotify(world, player)
		g.SendMsg(cmd.SceneTeamUpdateNotify, player.PlayerId, player.ClientSeq, sceneTeamUpdateNotify)
	}

	setUpAvatarTeamRsp := &proto.SetUpAvatarTeamRsp{
		TeamId:             req.TeamId,
		CurAvatarGuid:      req.CurAvatarGuid,
		AvatarTeamGuidList: req.AvatarTeamGuidList,
	}
	g.SendMsg(cmd.SetUpAvatarTeamRsp, player.PlayerId, player.ClientSeq, setUpAvatarTeamRsp)
}

func (g *Game) ChooseCurAvatarTeamReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ChooseCurAvatarTeamReq)
	teamId := req.TeamId
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		logger.Error("get world is nil, worldId: %v, uid: %v", player.WorldId, player.PlayerId)
		return
	}
	if world.GetMultiplayer() {
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
	// player.TeamConfig.UpdateTeam()
	world.SetPlayerAvatarIndex(player, 0)
	world.SetPlayerLocalTeam(player, team.GetAvatarIdList())
	world.UpdateMultiplayerTeam()
	world.InitPlayerWorldAvatar(player)

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
	if WORLD_MANAGER.IsAiWorld(world) || !world.GetMultiplayer() || len(avatarGuidList) == 0 || len(avatarGuidList) > 4 {
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
	world.InitPlayerWorldAvatar(player)

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
	g.SendToWorldA(world, cmd.SceneTeamUpdateNotify, player.ClientSeq, sceneTeamUpdateNotify)

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

	if WORLD_MANAGER.IsAiWorld(world) {
		pubg := world.GetPubg()
		if pubg != nil {
			alivePlayerNum := len(pubg.GetAlivePlayerList())
			info := fmt.Sprintf("『%v』死亡了，剩余%v位存活玩家。", player.NickName, alivePlayerNum)
			g.PlayerChatReq(world.GetOwner(), &proto.PlayerChatReq{ChatInfo: &proto.ChatInfo{Content: &proto.ChatInfo_Text{Text: info}}})
			player.PubgRank += uint32(100 - alivePlayerNum)
			g.SendMsg(cmd.AvatarDieAnimationEndRsp, player.PlayerId, player.ClientSeq, &proto.AvatarDieAnimationEndRsp{SkillId: req.SkillId, DieGuid: req.DieGuid})
			return
		}
	}

	entity := scene.GetEntity(uint32(req.DieGuid))
	if entity.GetLastDieType() == int32(proto.PlayerDieType_PLAYER_DIE_DRAWN) {
		maxStamina := player.PropertiesMap[constant.PLAYER_PROP_MAX_STAMINA]
		// 设置玩家耐力为一半
		g.SetPlayerStamina(player, maxStamina/2)
		// 传送玩家至安全位置
		g.TeleportPlayer(
			player,
			proto.EnterReason_ENTER_REASON_REVIVAL,
			player.SceneId,
			&model.Vector{
				X: player.SafePos.X,
				Y: player.SafePos.Y,
				Z: player.SafePos.Z,
			},
			new(model.Vector),
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
		&model.Vector{
			X: player.SafePos.X,
			Y: player.SafePos.Y,
			Z: player.SafePos.Z,
		},
		new(model.Vector),
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
	if !world.GetMultiplayer() {
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
	g.SendToSceneA(scene, cmd.SceneEntityDisappearNotify, player.ClientSeq, sceneEntityDisappearNotify)

	newAvatarId := world.GetPlayerActiveAvatarId(player)
	newAvatarEntity := g.PacketSceneEntityInfoAvatar(scene, player, newAvatarId)
	sceneEntityAppearNotify := &proto.SceneEntityAppearNotify{
		AppearType: proto.VisionType_VISION_REPLACE,
		Param:      oldAvatarEntity.GetId(),
		EntityList: []*proto.SceneEntityInfo{newAvatarEntity},
	}
	g.SendToSceneA(scene, cmd.SceneEntityAppearNotify, player.ClientSeq, sceneEntityAppearNotify)
}

/************************************************** 打包封装 **************************************************/

func (g *Game) PacketSceneTeamUpdateNotify(world *World, player *model.Player) *proto.SceneTeamUpdateNotify {
	sceneTeamUpdateNotify := &proto.SceneTeamUpdateNotify{
		IsInMp: world.GetMultiplayer(),
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
			IsPlayerCurAvatar: world.GetPlayerActiveAvatarId(worldPlayer) == worldAvatar.GetAvatarId(),
			IsOnScene:         world.GetPlayerActiveAvatarId(worldPlayer) == worldAvatar.GetAvatarId(),
			AvatarAbilityInfo: &proto.AbilitySyncStateInfo{
				IsInited:           len(worldAvatar.GetAbilityList()) != 0,
				DynamicValueMap:    nil,
				AppliedAbilities:   worldAvatar.GetAbilityList(),
				AppliedModifiers:   worldAvatar.GetModifierList(),
				MixinRecoverInfos:  nil,
				SgvDynamicValueMap: nil,
			},
			WeaponAbilityInfo:   empty,
			AbilityControlBlock: new(proto.AbilityControlBlock),
		}
		if world.GetMultiplayer() {
			sceneTeamAvatar.AvatarInfo = g.PacketAvatarInfo(worldPlayerAvatar)
			sceneTeamAvatar.SceneAvatarInfo = g.PacketSceneAvatarInfo(worldPlayerScene, worldPlayer, worldAvatar.GetAvatarId())
		}
		// 角色的ability控制块
		acb := sceneTeamAvatar.AbilityControlBlock
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
		avatarDataConfig := gdconf.GetAvatarDataById(int32(worldAvatar.GetAvatarId()))
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
		skillDepot := gdconf.GetAvatarSkillDepotDataById(int32(worldPlayerAvatar.SkillDepotId))
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
		sceneTeamUpdateNotify.SceneTeamAvatarList = append(sceneTeamUpdateNotify.SceneTeamAvatarList, sceneTeamAvatar)
	}
	return sceneTeamUpdateNotify
}
