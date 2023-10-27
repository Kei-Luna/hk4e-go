package game

import (
	"fmt"
	"math"
	"strings"
	"time"

	"hk4e/common/constant"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/pkg/reflection"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

var cmdProtoMap *cmd.CmdProtoMap = nil
var attackResultTemplate *proto.AttackResult = nil

func DoForward[IET model.InvokeEntryType](player *model.Player, invokeHandler *model.InvokeHandler[IET],
	cmdId uint16, newNtf pb.Message, forwardField string,
	srcNtf pb.Message, copyFieldList []string) {
	if cmdProtoMap == nil {
		cmdProtoMap = cmd.NewCmdProtoMap()
	}
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	if srcNtf != nil && copyFieldList != nil {
		for _, fieldName := range copyFieldList {
			reflection.CopyStructField(newNtf, srcNtf, fieldName)
		}
	}
	if invokeHandler.AllLen() == 0 && invokeHandler.AllExceptCurLen() == 0 && invokeHandler.HostLen() == 0 {
		return
	}
	// TODO aoi漏移除玩家的bug解决了就删掉
	if WORLD_MANAGER.IsAiWorld(world) {
		aiWorldAoi := world.GetAiWorldAoi()
		gid := aiWorldAoi.GetGidByPos(float32(player.Pos.X), float32(player.Pos.Y), float32(player.Pos.Z))
		if gid == math.MaxUint32 {
			return
		}
		gridList := aiWorldAoi.GetSurrGridListByGid(gid)
		for _, grid := range gridList {
			objectList := grid.GetObjectList()
			for uid, wa := range objectList {
				playerMap := world.GetAllPlayer()
				_, exist := playerMap[uint32(uid)]
				if !exist {
					logger.Error("remove not in world player cause by aoi bug, niw uid: %v, niw wa: %+v, uid: %v", uid, wa, player.PlayerId)
					delete(objectList, uid)
				}
			}
		}
	}
	// TODO aoi漏移除玩家的bug解决了就删掉
	if WORLD_MANAGER.IsAiWorld(world) && cmdId != cmd.CombatInvocationsNotify {
		if invokeHandler.AllLen() > 0 {
			reflection.SetStructFieldValue(newNtf, forwardField, invokeHandler.EntryListForwardAll)
			GAME.SendToSceneACV(scene, cmdId, player.ClientSeq, newNtf, 0, player.ClientVersion)
		}
		if invokeHandler.AllExceptCurLen() > 0 {
			reflection.SetStructFieldValue(newNtf, forwardField, invokeHandler.EntryListForwardAllExceptCur)
			GAME.SendToSceneACV(scene, cmdId, player.ClientSeq, newNtf, player.PlayerId, player.ClientVersion)
		}
		if invokeHandler.HostLen() > 0 {
			reflection.SetStructFieldValue(newNtf, forwardField, invokeHandler.EntryListForwardHost)
			GAME.SendToWorldH(world, cmdId, player.ClientSeq, newNtf)
		}
		return
	}
	if invokeHandler.AllLen() > 0 {
		reflection.SetStructFieldValue(newNtf, forwardField, invokeHandler.EntryListForwardAll)
		GAME.SendToSceneA(scene, cmdId, player.ClientSeq, newNtf, 0)
	}
	if invokeHandler.AllExceptCurLen() > 0 {
		reflection.SetStructFieldValue(newNtf, forwardField, invokeHandler.EntryListForwardAllExceptCur)
		GAME.SendToSceneA(scene, cmdId, player.ClientSeq, newNtf, player.PlayerId)
	}
	if invokeHandler.HostLen() > 0 {
		reflection.SetStructFieldValue(newNtf, forwardField, invokeHandler.EntryListForwardHost)
		GAME.SendToWorldH(world, cmdId, player.ClientSeq, newNtf)
	}
}

func (g *Game) UnionCmdNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.UnionCmdNotify)
	_ = ntf
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	DoForward[proto.CombatInvokeEntry](player, player.CombatInvokeHandler,
		cmd.CombatInvocationsNotify, new(proto.CombatInvocationsNotify), "InvokeList",
		nil, nil)
	DoForward[proto.AbilityInvokeEntry](player, player.AbilityInvokeHandler,
		cmd.AbilityInvocationsNotify, new(proto.AbilityInvocationsNotify), "Invokes",
		nil, nil)
	player.CombatInvokeHandler.Clear()
	player.AbilityInvokeHandler.Clear()
}

func (g *Game) CombatInvocationsNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.CombatInvocationsNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	for _, entry := range ntf.InvokeList {
		switch entry.ArgumentType {
		case proto.CombatTypeArgument_COMBAT_EVT_BEING_HIT:
			evtBeingHitInfo := new(proto.EvtBeingHitInfo)
			err := pb.Unmarshal(entry.CombatData, evtBeingHitInfo)
			if err != nil {
				logger.Error("parse EvtBeingHitInfo error: %v", err)
				break
			}
			// logger.Debug("EvtBeingHitInfo: %+v, ForwardType: %v", evtBeingHitInfo, entry.ForwardType)
			g.handleEvtBeingHit(player, scene, evtBeingHitInfo)
		case proto.CombatTypeArgument_ENTITY_MOVE:
			entityMoveInfo := new(proto.EntityMoveInfo)
			err := pb.Unmarshal(entry.CombatData, entityMoveInfo)
			if err != nil {
				logger.Error("parse EntityMoveInfo error: %v", err)
				break
			}
			// logger.Debug("EntityMoveInfo: %+v, ForwardType: %v", entityMoveInfo, entry.ForwardType)
			motionInfo := entityMoveInfo.MotionInfo
			if motionInfo == nil || motionInfo.Pos == nil || motionInfo.Rot == nil {
				break
			}
			g.handleEntityMove(player, world, scene, entityMoveInfo.EntityId, &model.Vector{
				X: float64(motionInfo.Pos.X),
				Y: float64(motionInfo.Pos.Y),
				Z: float64(motionInfo.Pos.Z),
			}, &model.Vector{
				X: float64(motionInfo.Rot.X),
				Y: float64(motionInfo.Rot.Y),
				Z: float64(motionInfo.Rot.Z),
			}, false, entityMoveInfo)
			// 众里寻他千百度 蓦然回首 那人却在灯火阑珊处
			if motionInfo.State == proto.MotionState_MOTION_NOTIFY || motionInfo.State == proto.MotionState_MOTION_FIGHT {
				// 只要转发了这两个包的其中之一 客户端的动画就会被打断
				continue
			}
		case proto.CombatTypeArgument_COMBAT_ANIMATOR_PARAMETER_CHANGED:
			evtAnimatorParameterInfo := new(proto.EvtAnimatorParameterInfo)
			err := pb.Unmarshal(entry.CombatData, evtAnimatorParameterInfo)
			if err != nil {
				logger.Error("parse EvtAnimatorParameterInfo error: %v", err)
				break
			}
			// logger.Debug("EvtAnimatorParameterInfo: %+v, ForwardType: %v", evtAnimatorParameterInfo, entry.ForwardType)
		case proto.CombatTypeArgument_COMBAT_ANIMATOR_STATE_CHANGED:
			evtAnimatorStateChangedInfo := new(proto.EvtAnimatorStateChangedInfo)
			err := pb.Unmarshal(entry.CombatData, evtAnimatorStateChangedInfo)
			if err != nil {
				logger.Error("parse EvtAnimatorStateChangedInfo error: %v", err)
				break
			}
			// logger.Debug("EvtAnimatorStateChangedInfo: %+v, ForwardType: %v", evtAnimatorStateChangedInfo, entry.ForwardType)
		}
		player.CombatInvokeHandler.AddEntry(entry.ForwardType, entry)
	}
}

func (g *Game) handleEvtBeingHit(player *model.Player, scene *Scene, hitInfo *proto.EvtBeingHitInfo) {
	world := scene.GetWorld()
	attackResult := hitInfo.AttackResult
	if attackResult == nil {
		logger.Error("attackResult is nil")
		return
	}
	if attackResultTemplate == nil {
		attackResultTemplate = attackResult
	}
	defEntity := scene.GetEntity(attackResult.DefenseId)
	if defEntity == nil {
		logger.Error("not found def entity, DefenseId: %v", attackResult.DefenseId)
		return
	}

	if WORLD_MANAGER.IsAiWorld(world) {
		if defEntity.GetEntityType() == constant.ENTITY_TYPE_AVATAR &&
			defEntity.GetAvatarEntity().GetUid() == world.GetOwner().PlayerId {
			return
		}
	}

	fightProp := defEntity.GetFightProp()
	currHp := fightProp[constant.FIGHT_PROP_CUR_HP]
	if currHp == 0.0 {
		return
	}
	currHp -= attackResult.Damage
	deltaHp := -attackResult.Damage
	if currHp < 0.0 {
		deltaHp -= currHp
		currHp = 0.0
	}
	fightProp[constant.FIGHT_PROP_CUR_HP] = currHp
	g.EntityFightPropUpdateNotifyBroadcast(scene, defEntity)
	switch defEntity.GetEntityType() {
	case constant.ENTITY_TYPE_AVATAR:
		g.SendMsg(cmd.EntityFightPropChangeReasonNotify, player.PlayerId, player.ClientSeq, &proto.EntityFightPropChangeReasonNotify{
			PropDelta:      deltaHp,
			ChangeHpReason: proto.ChangHpReason_CHANGE_HP_SUB_GM,
			EntityId:       defEntity.GetId(),
			PropType:       constant.FIGHT_PROP_CUR_HP,
		})
		if currHp == 0.0 {
			defAvatarEntity := defEntity.GetAvatarEntity()
			g.KillPlayerAvatar(player, defAvatarEntity.GetAvatarId(), proto.PlayerDieType_PLAYER_DIE_GM)

			if WORLD_MANAGER.IsAiWorld(world) {
				defPlayer := USER_MANAGER.GetOnlineUser(defAvatarEntity.GetUid())
				if defPlayer == nil {
					return
				}
				atkEntity := scene.GetEntity(attackResult.AttackerId)
				if atkEntity != nil && atkEntity.GetEntityType() == constant.ENTITY_TYPE_AVATAR {
					atkAvatarEntity := atkEntity.GetAvatarEntity()
					atkPlayer := USER_MANAGER.GetOnlineUser(atkAvatarEntity.GetUid())
					if atkPlayer == nil {
						return
					}
					info := fmt.Sprintf("『%v』击败了『%v』。", atkPlayer.NickName, defPlayer.NickName)
					g.PlayerChatReq(world.GetOwner(), &proto.PlayerChatReq{ChatInfo: &proto.ChatInfo{Content: &proto.ChatInfo_Text{Text: info}}})
				}
			}
		}
	case constant.ENTITY_TYPE_MONSTER:
		if currHp == 0.0 {
			g.KillEntity(player, scene, defEntity.GetId(), proto.PlayerDieType_PLAYER_DIE_GM)
		}
	case constant.ENTITY_TYPE_GADGET:
		gadgetEntity := defEntity.GetGadgetEntity()
		gadgetDataConfig := gdconf.GetGadgetDataById(int32(gadgetEntity.GetGadgetId()))
		if gadgetDataConfig == nil {
			logger.Error("get gadget data config is nil, gadgetId: %v", gadgetEntity.GetGadgetId())
			break
		}
		logger.Debug("[EvtBeingHit] GadgetData: %+v, EntityId: %v, uid: %v", gadgetDataConfig, defEntity.GetId(), player.PlayerId)
		g.handleGadgetEntityBeHitLow(player, defEntity, attackResult.ElementType)
	}
}

func (g *Game) handleEntityMove(player *model.Player, world *World, scene *Scene, entityId uint32, pos, rot *model.Vector, force bool, moveInfo *proto.EntityMoveInfo) {
	entity := scene.GetEntity(entityId)
	if entity == nil {
		return
	}
	if entity.GetEntityType() == constant.ENTITY_TYPE_AVATAR {
		// 玩家实体在移动
		avatarEntity := entity.GetAvatarEntity()
		if avatarEntity.GetUid() != player.PlayerId {
			return
		}
		if !WORLD_MANAGER.IsAiWorld(world) {
			g.SceneBlockAoiPlayerMove(player, world, scene, player.Pos, pos, entity.GetId())
		} else {
			g.AiWorldAoiPlayerMove(player, world, scene, player.Pos, pos)
		}
		// 场景天气区域变更检测
		g.SceneWeatherAreaCheck(player, player.Pos, pos)
		// 更新玩家的位置信息
		player.Pos.X, player.Pos.Y, player.Pos.Z = pos.X, pos.Y, pos.Z
		player.Rot.X, player.Rot.Y, player.Rot.Z = rot.X, rot.Y, rot.Z
	}
	// 更新场景实体的位置信息
	entity.SetPos(pos)
	entity.SetRot(rot)
	if !force {
		motionInfo := moveInfo.MotionInfo
		switch entity.GetEntityType() {
		case constant.ENTITY_TYPE_AVATAR:
			// 玩家安全位置更新
			switch motionInfo.State {
			case proto.MotionState_MOTION_DANGER_RUN,
				proto.MotionState_MOTION_RUN,
				proto.MotionState_MOTION_DANGER_STANDBY_MOVE,
				proto.MotionState_MOTION_DANGER_STANDBY,
				proto.MotionState_MOTION_LADDER_TO_STANDBY,
				proto.MotionState_MOTION_STANDBY_MOVE,
				proto.MotionState_MOTION_STANDBY,
				proto.MotionState_MOTION_DANGER_WALK,
				proto.MotionState_MOTION_WALK,
				proto.MotionState_MOTION_DASH:
				// 仅在陆地时更新玩家安全位置
				player.SafePos.X, player.SafePos.Y, player.SafePos.Z = player.Pos.X, player.Pos.Y, player.Pos.Z
			}
			// 处理耐力消耗
			g.ImmediateStamina(player, motionInfo.State)
		case constant.ENTITY_TYPE_GADGET:
			// 载具耐力消耗
			gadgetEntity := entity.GetGadgetEntity()
			if gadgetEntity.GetGadgetVehicleEntity() != nil {
				// 处理耐力消耗
				g.ImmediateStamina(player, motionInfo.State)
				// 处理载具销毁请求
				g.VehicleDestroyMotion(player, entity, motionInfo.State)
			}
		}
		entity.SetMoveState(uint16(motionInfo.State))
		entity.SetLastMoveSceneTimeMs(moveInfo.SceneTime)
		entity.SetLastMoveReliableSeq(moveInfo.ReliableSeq)
	}
}

func (g *Game) SceneBlockAoiPlayerMove(player *model.Player, world *World, scene *Scene, oldPos *model.Vector, newPos *model.Vector, avatarEntityId uint32) {
	if !world.IsValidSceneBlockPos(scene.GetId(), float32(newPos.X), 0.0, float32(newPos.Z)) {
		return
	}
	// 服务器处理玩家移动场景区块aoi事件频率限制
	now := uint64(time.Now().UnixMilli())
	if now-player.LastSceneBlockAoiMoveTime < 200 {
		return
	}
	player.LastSceneBlockAoiMoveTime = now
	sceneBlockAoiMap := WORLD_MANAGER.GetSceneBlockAoiMap()
	aoiManager, exist := sceneBlockAoiMap[player.SceneId]
	if !exist {
		logger.Error("get scene block aoi is nil, sceneId: %v, uid: %v", player.SceneId, player.PlayerId)
		return
	}
	oldGid := aoiManager.GetGidByPos(float32(oldPos.X), 0.0, float32(oldPos.Z))
	newGid := aoiManager.GetGidByPos(float32(newPos.X), 0.0, float32(newPos.Z))
	if oldGid != newGid {
		// 跨越了block格子
		logger.Debug("player cross scene block grid, oldGid: %v, newGid: %v, uid: %v", oldGid, newGid, player.PlayerId)
	}
	// 加载和卸载的group
	oldNeighborGroupMap := g.GetNeighborGroup(player.SceneId, oldPos)
	newNeighborGroupMap := g.GetNeighborGroup(player.SceneId, newPos)
	for groupId, groupConfig := range oldNeighborGroupMap {
		_, exist := newNeighborGroupMap[groupId]
		if exist {
			continue
		}
		// 旧有新没有的group即为卸载的
		if !world.GetMultiplayer() {
			// 单人世界直接卸载group
			g.RemoveSceneGroup(player, scene, groupConfig)
		} else {
			// 多人世界group附近没有任何玩家则卸载
			remove := true
			for _, otherPlayer := range scene.GetAllPlayer() {
				dx := int32(otherPlayer.Pos.X) - int32(groupConfig.Pos.X)
				if dx < 0 {
					dx *= -1
				}
				dy := int32(otherPlayer.Pos.Z) - int32(groupConfig.Pos.Z)
				if dy < 0 {
					dy *= -1
				}
				if dx <= GROUP_LOAD_DISTANCE || dy <= GROUP_LOAD_DISTANCE {
					remove = false
					break
				}
			}
			if remove {
				g.RemoveSceneGroup(player, scene, groupConfig)
			}
		}
	}
	for groupId, groupConfig := range newNeighborGroupMap {
		_, exist := oldNeighborGroupMap[groupId]
		if exist {
			continue
		}
		// 新有旧没有的group即为加载的
		g.AddSceneGroup(player, scene, groupConfig)
	}
	// 消失和出现的场景实体
	oldVisionEntityMap := g.GetVisionEntity(scene, oldPos)
	newVisionEntityMap := g.GetVisionEntity(scene, newPos)
	delEntityIdList := make([]uint32, 0)
	for entityId := range oldVisionEntityMap {
		_, exist := newVisionEntityMap[entityId]
		if exist {
			continue
		}
		// 旧有新没有的实体即为消失的
		delEntityIdList = append(delEntityIdList, entityId)
	}
	addEntityIdList := make([]uint32, 0)
	for entityId := range newVisionEntityMap {
		_, exist := oldVisionEntityMap[entityId]
		if exist {
			continue
		}
		// 新有旧没有的实体即为出现的
		addEntityIdList = append(addEntityIdList, entityId)
	}
	// 同步客户端消失和出现的场景实体
	if len(delEntityIdList) > 0 {
		g.RemoveSceneEntityNotifyToPlayer(player, proto.VisionType_VISION_MISS, delEntityIdList)
	}
	if len(addEntityIdList) > 0 {
		g.AddSceneEntityNotify(player, proto.VisionType_VISION_MEET, addEntityIdList, false, false)
	}
	// 场景区域触发器检测
	g.SceneRegionTriggerCheck(player, oldPos, newPos, avatarEntityId)
}

func (g *Game) AiWorldAoiPlayerMove(player *model.Player, world *World, scene *Scene, oldPos *model.Vector, newPos *model.Vector) {
	if !world.IsValidAiWorldPos(scene.GetId(), float32(newPos.X), float32(newPos.Y), float32(newPos.Z)) {
		return
	}
	aiWorldAoi := world.GetAiWorldAoi()
	oldGid := aiWorldAoi.GetGidByPos(float32(oldPos.X), float32(oldPos.Y), float32(oldPos.Z))
	newGid := aiWorldAoi.GetGidByPos(float32(newPos.X), float32(newPos.Y), float32(newPos.Z))
	if oldGid != newGid {
		// 玩家跨越了格子
		logger.Debug("player cross ai world aoi grid, oldGid: %v, oldPos: %+v, newGid: %v, newPos: %+v, uid: %v",
			oldGid, oldPos, newGid, newPos, player.PlayerId)
		// 找出本次移动所带来的消失和出现的格子
		oldGridList := aiWorldAoi.GetSurrGridListByGid(oldGid)
		newGridList := aiWorldAoi.GetSurrGridListByGid(newGid)
		delGridIdList := make([]uint32, 0)
		for _, oldGrid := range oldGridList {
			exist := false
			for _, newGrid := range newGridList {
				if oldGrid.GetGid() == newGrid.GetGid() {
					exist = true
					break
				}
			}
			if exist {
				continue
			}
			delGridIdList = append(delGridIdList, oldGrid.GetGid())
		}
		addGridIdList := make([]uint32, 0)
		for _, newGrid := range newGridList {
			exist := false
			for _, oldGrid := range oldGridList {
				if newGrid.GetGid() == oldGrid.GetGid() {
					exist = true
					break
				}
			}
			if exist {
				continue
			}
			addGridIdList = append(addGridIdList, newGrid.GetGid())
		}
		activeAvatarId := world.GetPlayerActiveAvatarId(player)
		activeWorldAvatar := world.GetPlayerWorldAvatar(player, activeAvatarId)
		// 老格子移除玩家
		logger.Debug("ai world aoi remove player, oldPos: %+v, uid: %v", oldPos, player.PlayerId)
		ok := aiWorldAoi.RemoveObjectFromGridByPos(int64(player.PlayerId), float32(oldPos.X), float32(oldPos.Y), float32(oldPos.Z))
		if !ok {
			logger.Error("ai world aoi remove player fail, uid: %v, pos: %+v", player.PlayerId, player.Pos)
		}
		// 处理消失的格子
		for _, delGridId := range delGridIdList {
			// 通知自己 老格子里的其它玩家消失
			oldOtherWorldAvatarMap := aiWorldAoi.GetObjectListByGid(delGridId)
			delEntityIdList := make([]uint32, 0)
			for _, otherWorldAvatarAny := range oldOtherWorldAvatarMap {
				otherWorldAvatar := otherWorldAvatarAny.(*WorldAvatar)
				delEntityIdList = append(delEntityIdList, otherWorldAvatar.GetAvatarEntityId())
			}
			if len(delEntityIdList) > 0 {
				g.RemoveSceneEntityNotifyToPlayer(player, proto.VisionType_VISION_MISS, delEntityIdList)
			}
			// 通知老格子里的其它玩家 自己消失
			for otherPlayerId := range oldOtherWorldAvatarMap {
				otherPlayer := USER_MANAGER.GetOnlineUser(uint32(otherPlayerId))
				if otherPlayer == nil {
					logger.Error("get player is nil, target uid: %v, uid: %v", otherPlayerId, player.PlayerId)
					continue
				}
				g.RemoveSceneEntityNotifyToPlayer(otherPlayer, proto.VisionType_VISION_MISS, []uint32{activeWorldAvatar.GetAvatarEntityId()})
			}
		}
		// 处理出现的格子
		for _, addGridId := range addGridIdList {
			// 通知自己 新格子里的其他玩家出现
			newOtherWorldAvatarMap := aiWorldAoi.GetObjectListByGid(addGridId)
			addEntityIdList := make([]uint32, 0)
			for _, otherWorldAvatarAny := range newOtherWorldAvatarMap {
				otherWorldAvatar := otherWorldAvatarAny.(*WorldAvatar)
				addEntityIdList = append(addEntityIdList, otherWorldAvatar.GetAvatarEntityId())
			}
			if len(addEntityIdList) > 0 {
				g.AddSceneEntityNotify(player, proto.VisionType_VISION_MEET, addEntityIdList, false, false)
			}
			// 通知新格子里的其他玩家 自己出现
			for otherPlayerId := range newOtherWorldAvatarMap {
				otherPlayer := USER_MANAGER.GetOnlineUser(uint32(otherPlayerId))
				if otherPlayer == nil {
					logger.Error("get player is nil, target uid: %v, uid: %v", otherPlayerId, player.PlayerId)
					continue
				}
				sceneEntityInfoAvatar := g.PacketSceneEntityInfoAvatar(scene, player, world.GetPlayerActiveAvatarId(player))
				g.AddSceneEntityNotifyToPlayer(otherPlayer, proto.VisionType_VISION_MEET, []*proto.SceneEntityInfo{sceneEntityInfoAvatar})
			}
		}
		// 新格子添加玩家
		logger.Debug("ai world aoi add player, newPos: %+v, uid: %v", newPos, player.PlayerId)
		ok = aiWorldAoi.AddObjectToGridByPos(int64(player.PlayerId), activeWorldAvatar, float32(newPos.X), float32(newPos.Y), float32(newPos.Z))
		if !ok {
			logger.Error("ai world aoi add player fail, uid: %v, pos: %+v", player.PlayerId, player.Pos)
		}
	}
	// 消失和出现的场景实体
	oldVisionEntityMap := g.GetVisionEntity(scene, oldPos)
	newVisionEntityMap := g.GetVisionEntity(scene, newPos)
	delEntityIdList := make([]uint32, 0)
	for entityId, entity := range oldVisionEntityMap {
		if entity.GetEntityType() == constant.ENTITY_TYPE_AVATAR {
			continue
		}
		_, exist := newVisionEntityMap[entityId]
		if exist {
			continue
		}
		// 旧有新没有的实体即为消失的
		delEntityIdList = append(delEntityIdList, entityId)
	}
	addEntityIdList := make([]uint32, 0)
	for entityId, entity := range newVisionEntityMap {
		if entity.GetEntityType() == constant.ENTITY_TYPE_AVATAR {
			continue
		}
		_, exist := oldVisionEntityMap[entityId]
		if exist {
			continue
		}
		// 新有旧没有的实体即为出现的
		addEntityIdList = append(addEntityIdList, entityId)
	}
	// 同步客户端消失和出现的场景实体
	if len(delEntityIdList) > 0 {
		g.RemoveSceneEntityNotifyToPlayer(player, proto.VisionType_VISION_MISS, delEntityIdList)
	}
	if len(addEntityIdList) > 0 {
		g.AddSceneEntityNotify(player, proto.VisionType_VISION_MEET, addEntityIdList, false, false)
	}
}

func (g *Game) AbilityInvocationsNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.AbilityInvocationsNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	for _, entry := range ntf.Invokes {
		player.AbilityInvokeHandler.AddEntry(entry.ForwardType, entry)
		switch entry.ArgumentType {
		case proto.AbilityInvokeArgument_ABILITY_META_MODIFIER_CHANGE:
			modifierChange := new(proto.AbilityMetaModifierChange)
			err := pb.Unmarshal(entry.AbilityData, modifierChange)
			if err != nil {
				logger.Error("parse AbilityMetaModifierChange error: %v", err)
				continue
			}
			// logger.Debug("EntityId: %v, ModifierChange: %+v", entry.EntityId, modifierChange)
			// 处理耐力消耗
			g.HandleAbilityStamina(player, entry)
			g.handleGadgetEntityAbilityLow(player, entry.EntityId, entry.ArgumentType, modifierChange)
		case proto.AbilityInvokeArgument_ABILITY_MIXIN_COST_STAMINA:
			costStamina := new(proto.AbilityMixinCostStamina)
			err := pb.Unmarshal(entry.AbilityData, costStamina)
			if err != nil {
				logger.Error("parse AbilityMixinCostStamina error: %v", err)
				continue
			}
			// logger.Debug("EntityId: %v, MixinCostStamina: %+v", entry.EntityId, costStamina)
			// 处理耐力消耗
			g.HandleAbilityStamina(player, entry)
		case proto.AbilityInvokeArgument_ABILITY_META_MODIFIER_DURABILITY_CHANGE:
			modifierDurabilityChange := new(proto.AbilityMetaModifierDurabilityChange)
			err := pb.Unmarshal(entry.AbilityData, modifierDurabilityChange)
			if err != nil {
				logger.Error("parse AbilityMetaModifierDurabilityChange error: %v", err)
				continue
			}
			// logger.Debug("EntityId: %v, DurabilityChange: %+v", entry.EntityId, modifierDurabilityChange)
		}
	}
}

func (g *Game) ClientAbilityInitFinishNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.ClientAbilityInitFinishNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	invokeHandler := model.NewInvokeHandler[proto.AbilityInvokeEntry]()
	for _, entry := range ntf.Invokes {
		// logger.Debug("ClientAbilityInitFinishNotify: %+v", entry)
		invokeHandler.AddEntry(entry.ForwardType, entry)
	}
	DoForward[proto.AbilityInvokeEntry](player, invokeHandler,
		cmd.ClientAbilityInitFinishNotify, new(proto.ClientAbilityInitFinishNotify), "Invokes",
		ntf, []string{"EntityId"})
}

func (g *Game) ClientAbilityChangeNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.ClientAbilityChangeNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	invokeHandler := model.NewInvokeHandler[proto.AbilityInvokeEntry]()
	for _, entry := range ntf.Invokes {
		// logger.Debug("ClientAbilityChangeNotify: %+v", entry)
		invokeHandler.AddEntry(entry.ForwardType, entry)
	}
	DoForward[proto.AbilityInvokeEntry](player, invokeHandler,
		cmd.ClientAbilityChangeNotify, new(proto.ClientAbilityChangeNotify), "Invokes",
		ntf, []string{"IsInitHash", "EntityId"})
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	for _, abilityInvokeEntry := range ntf.Invokes {
		switch abilityInvokeEntry.ArgumentType {
		case proto.AbilityInvokeArgument_ABILITY_META_ADD_NEW_ABILITY:
			abilityMetaAddAbility := new(proto.AbilityMetaAddAbility)
			err := pb.Unmarshal(abilityInvokeEntry.AbilityData, abilityMetaAddAbility)
			if err != nil {
				logger.Error("parse AbilityMetaAddAbility error: %v", err)
				continue
			}
			worldAvatar := world.GetWorldAvatarByEntityId(abilityInvokeEntry.EntityId)
			if worldAvatar == nil {
				continue
			}
			if abilityMetaAddAbility.Ability == nil {
				continue
			}
			worldAvatar.AddAbility(abilityMetaAddAbility.Ability)
		case proto.AbilityInvokeArgument_ABILITY_META_MODIFIER_CHANGE:
			abilityMetaModifierChange := new(proto.AbilityMetaModifierChange)
			err := pb.Unmarshal(abilityInvokeEntry.AbilityData, abilityMetaModifierChange)
			if err != nil {
				logger.Error("parse AbilityMetaModifierChange error: %v", err)
				continue
			}
			abilityAppliedModifier := &proto.AbilityAppliedModifier{
				ModifierLocalId:           abilityMetaModifierChange.ModifierLocalId,
				ParentAbilityEntityId:     0,
				ParentAbilityName:         abilityMetaModifierChange.ParentAbilityName,
				ParentAbilityOverride:     abilityMetaModifierChange.ParentAbilityOverride,
				InstancedAbilityId:        abilityInvokeEntry.Head.InstancedAbilityId,
				InstancedModifierId:       abilityInvokeEntry.Head.InstancedModifierId,
				ExistDuration:             0,
				AttachedInstancedModifier: abilityMetaModifierChange.AttachedInstancedModifier,
				ApplyEntityId:             abilityMetaModifierChange.ApplyEntityId,
				IsAttachedParentAbility:   abilityMetaModifierChange.IsAttachedParentAbility,
				ModifierDurability:        nil,
				SbuffUid:                  0,
				IsServerbuffModifier:      abilityInvokeEntry.Head.IsServerbuffModifier,
			}
			worldAvatar := world.GetWorldAvatarByEntityId(abilityInvokeEntry.EntityId)
			if worldAvatar == nil {
				continue
			}
			worldAvatar.AddModifier(abilityAppliedModifier)
		}
	}
}

func (g *Game) MassiveEntityElementOpBatchNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.MassiveEntityElementOpBatchNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	ntf.OpIdx = scene.GetMeeoIndex()
	scene.SetMeeoIndex(scene.GetMeeoIndex() + 1)
	g.SendToSceneA(scene, cmd.MassiveEntityElementOpBatchNotify, player.ClientSeq, ntf, 0)
}

func (g *Game) EvtDoSkillSuccNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EvtDoSkillSuccNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	// logger.Debug("EvtDoSkillSuccNotify: %+v", ntf)

	// 触发事件
	if PLUGIN_MANAGER.TriggerEvent(PluginEventIdEvtDoSkillSucc, &PluginEventEvtDoSkillSucc{
		PluginEvent: NewPluginEvent(),
		Player:      player,
		Ntf:         ntf,
	}) {
		return
	}

	// 处理技能开始的耐力消耗
	g.SkillStartStamina(player, ntf.CasterId, ntf.SkillId)
	g.TriggerQuest(player, constant.QUEST_FINISH_COND_TYPE_SKILL, "", int32(ntf.SkillId))
}

func (g *Game) EvtAvatarEnterFocusNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EvtAvatarEnterFocusNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	// logger.Debug("EvtAvatarEnterFocusNotify: %+v", ntf)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	g.SendToSceneA(scene, cmd.EvtAvatarEnterFocusNotify, player.ClientSeq, ntf, 0)
}

func (g *Game) EvtAvatarUpdateFocusNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EvtAvatarUpdateFocusNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	// logger.Debug("EvtAvatarUpdateFocusNotify: %+v", ntf)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	g.SendToSceneA(scene, cmd.EvtAvatarUpdateFocusNotify, player.ClientSeq, ntf, 0)
}

func (g *Game) EvtAvatarExitFocusNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EvtAvatarExitFocusNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	// logger.Debug("EvtAvatarExitFocusNotify: %+v", ntf)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	g.SendToSceneA(scene, cmd.EvtAvatarExitFocusNotify, player.ClientSeq, ntf, 0)
}

func (g *Game) EvtEntityRenderersChangedNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EvtEntityRenderersChangedNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	// logger.Debug("EvtEntityRenderersChangedNotify: %+v", ntf)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	g.SendToSceneA(scene, cmd.EvtEntityRenderersChangedNotify, player.ClientSeq, ntf, 0)
}

func (g *Game) EvtBulletDeactiveNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EvtBulletDeactiveNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	// logger.Debug("EvtBulletDeactiveNotify: %+v", ntf)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	g.SendToSceneA(scene, cmd.EvtBulletDeactiveNotify, player.ClientSeq, ntf, 0)
}

func (g *Game) EvtBulletHitNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EvtBulletHitNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	// logger.Debug("EvtBulletHitNotify: %+v", ntf)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	g.SendToSceneA(scene, cmd.EvtBulletHitNotify, player.ClientSeq, ntf, 0)

	if WORLD_MANAGER.IsAiWorld(world) {
		bulletPhysicsEngine := world.GetBulletPhysicsEngine()
		if bulletPhysicsEngine.IsRigidBody(ntf.EntityId) {
			bulletPhysicsEngine.DestroyRigidBody(ntf.EntityId)
			_ = ntf.HitPoint
		}
	}
}

func (g *Game) EvtBulletMoveNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EvtBulletMoveNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	// logger.Debug("EvtBulletMoveNotify: %+v", ntf)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	g.SendToSceneA(scene, cmd.EvtBulletMoveNotify, player.ClientSeq, ntf, 0)
}

func (g *Game) EvtCreateGadgetNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EvtCreateGadgetNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	// logger.Debug("EvtCreateGadgetNotify: %+v", ntf)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	if ntf.InitPos == nil {
		return
	}
	ok := scene.CreateEntityGadgetClient(&model.Vector{
		X: float64(ntf.InitPos.X),
		Y: float64(ntf.InitPos.Y),
		Z: float64(ntf.InitPos.Z),
	}, &model.Vector{
		X: float64(ntf.InitEulerAngles.X),
		Y: float64(ntf.InitEulerAngles.Y),
		Z: float64(ntf.InitEulerAngles.Z),
	}, ntf.EntityId, ntf.ConfigId, ntf.CampId, ntf.CampType, ntf.OwnerEntityId, ntf.TargetEntityId, ntf.PropOwnerEntityId)
	if !ok {
		return
	}
	g.AddSceneEntityNotify(player, proto.VisionType_VISION_BORN, []uint32{ntf.EntityId}, true, true)

	if WORLD_MANAGER.IsAiWorld(world) {
		gadgetDataConfig := gdconf.GetGadgetDataById(int32(ntf.ConfigId))
		if gadgetDataConfig == nil {
			logger.Error("gadget data config is nil, gadgetId: %v", ntf.ConfigId)
			return
		}
		// 蓄力箭
		if gadgetDataConfig.PrefabPath != "ART/Others/Bullet/Bullet_ArrowAiming" &&
			gadgetDataConfig.PrefabPath != "ART/Others/Bullet/Bullet_Venti_ArrowAiming" {
			return
		}
		pitchAngleRaw := ntf.InitEulerAngles.X
		pitchAngle := float32(0.0)
		if pitchAngleRaw < 90.0 {
			pitchAngle = -pitchAngleRaw
		} else if pitchAngleRaw > 270.0 {
			pitchAngle = 360.0 - pitchAngleRaw
		} else {
			logger.Error("invalid raw pitch angle: %v, uid: %v", pitchAngleRaw, player.PlayerId)
			return
		}
		yawAngle := ntf.InitEulerAngles.Y
		bulletPhysicsEngine := world.GetBulletPhysicsEngine()
		activeAvatarId := world.GetPlayerActiveAvatarId(player)
		bulletPhysicsEngine.CreateRigidBody(
			ntf.EntityId,
			world.GetPlayerWorldAvatarEntityId(player, activeAvatarId),
			player.SceneId,
			ntf.InitPos.X, ntf.InitPos.Y, ntf.InitPos.Z,
			pitchAngle, yawAngle,
		)
	}
}

func (g *Game) EvtDestroyGadgetNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EvtDestroyGadgetNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	// logger.Debug("EvtDestroyGadgetNotify: %+v", ntf)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	scene.DestroyEntity(ntf.EntityId)
	g.RemoveSceneEntityNotifyBroadcast(scene, proto.VisionType_VISION_MISS, []uint32{ntf.EntityId})
}

func (g *Game) EvtAiSyncSkillCdNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EvtAiSyncSkillCdNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	// logger.Debug("EvtAiSyncSkillCdNotify: %+v", ntf)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	g.SendToSceneA(scene, cmd.EvtAiSyncSkillCdNotify, player.ClientSeq, ntf, 0)
}

func (g *Game) EvtAiSyncCombatThreatInfoNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EvtAiSyncCombatThreatInfoNotify)
	if player.SceneLoadState != model.SceneEnterDone {
		return
	}
	// logger.Debug("EvtAiSyncCombatThreatInfoNotify: %+v", ntf)
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	g.SendToSceneA(scene, cmd.EvtAiSyncCombatThreatInfoNotify, player.ClientSeq, ntf, 0)
}

func (g *Game) EntityConfigHashNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EntityConfigHashNotify)
	_ = ntf
}

func (g *Game) MonsterAIConfigHashNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.MonsterAIConfigHashNotify)
	_ = ntf
}

func (g *Game) SetEntityClientDataNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.SetEntityClientDataNotify)
	g.SendMsg(cmd.SetEntityClientDataNotify, player.PlayerId, player.ClientSeq, ntf)
}

func (g *Game) EntityAiSyncNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.EntityAiSyncNotify)
	entityAiSyncNotify := &proto.EntityAiSyncNotify{
		InfoList: make([]*proto.AiSyncInfo, 0),
	}
	for _, monsterId := range ntf.LocalAvatarAlertedMonsterList {
		entityAiSyncNotify.InfoList = append(entityAiSyncNotify.InfoList, &proto.AiSyncInfo{
			EntityId:        monsterId,
			HasPathToTarget: true,
			IsSelfKilling:   false,
		})
	}
	g.SendMsg(cmd.EntityAiSyncNotify, player.PlayerId, player.ClientSeq, entityAiSyncNotify)
}

// TODO 一些很low的解决方案 我本来是不想写的 有多low？要多low有多low！

func (g *Game) handleGadgetEntityBeHitLow(player *model.Player, entity *Entity, hitElementType uint32) {
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	if entity.GetEntityType() != constant.ENTITY_TYPE_GADGET {
		return
	}
	gadgetEntity := entity.GetGadgetEntity()
	gadgetId := gadgetEntity.GetGadgetId()
	gadgetDataConfig := gdconf.GetGadgetDataById(int32(gadgetId))
	if gadgetDataConfig == nil {
		logger.Error("get gadget data config is nil, gadgetId: %v", gadgetEntity.GetGadgetId())
		return
	}
	if strings.Contains(gadgetDataConfig.Name, "火把") ||
		strings.Contains(gadgetDataConfig.Name, "火盆") ||
		strings.Contains(gadgetDataConfig.Name, "篝火") {
		// 火把点燃
		if hitElementType != constant.ELEMENT_TYPE_FIRE {
			return
		}
		g.ChangeGadgetState(player, entity.GetId(), constant.GADGET_STATE_GEAR_START)
	} else if strings.Contains(gadgetDataConfig.ServerLuaScript, "Controller") {
		// 元素方碑点亮
		gadgetElementType := uint32(0)
		if strings.Contains(gadgetDataConfig.ServerLuaScript, "Fire") {
			gadgetElementType = constant.ELEMENT_TYPE_FIRE
		} else if strings.Contains(gadgetDataConfig.ServerLuaScript, "Water") {
			gadgetElementType = constant.ELEMENT_TYPE_WATER
		} else if strings.Contains(gadgetDataConfig.ServerLuaScript, "Grass") {
			gadgetElementType = constant.ELEMENT_TYPE_GRASS
		} else if strings.Contains(gadgetDataConfig.ServerLuaScript, "Elec") {
			gadgetElementType = constant.ELEMENT_TYPE_ELEC
		} else if strings.Contains(gadgetDataConfig.ServerLuaScript, "Ice") {
			gadgetElementType = constant.ELEMENT_TYPE_ICE
		} else if strings.Contains(gadgetDataConfig.ServerLuaScript, "Wind") {
			gadgetElementType = constant.ELEMENT_TYPE_WIND
		} else if strings.Contains(gadgetDataConfig.ServerLuaScript, "Rock") {
			gadgetElementType = constant.ELEMENT_TYPE_ROCK
		}
		if hitElementType != gadgetElementType {
			return
		}
		g.ChangeGadgetState(player, entity.GetId(), constant.GADGET_STATE_GEAR_START)
	} else if strings.Contains(gadgetDataConfig.ServerLuaScript, "SubfieldDrop_WoodenObject_Broken") {
		// 木箱破碎
		g.KillEntity(player, scene, entity.GetId(), proto.PlayerDieType_PLAYER_DIE_GM)
	}
}

func (g *Game) handleGadgetEntityAbilityLow(player *model.Player, entityId uint32, argument proto.AbilityInvokeArgument, entry pb.Message) {
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	entity := scene.GetEntity(entityId)
	if entity == nil {
		return
	}
	switch argument {
	case proto.AbilityInvokeArgument_ABILITY_META_MODIFIER_CHANGE:
		// 物件破碎
		modifierChange := entry.(*proto.AbilityMetaModifierChange)
		if modifierChange.Action != proto.ModifierAction_REMOVED {
			return
		}
		if entity.GetEntityType() != constant.ENTITY_TYPE_GADGET {
			return
		}
		gadgetEntity := entity.GetGadgetEntity()
		gadgetId := gadgetEntity.GetGadgetId()
		if gadgetId == 0 {
			return
		}
		gadgetDataConfig := gdconf.GetGadgetDataById(int32(gadgetId))
		if gadgetDataConfig == nil {
			logger.Error("get gadget data config is nil, gadgetId: %v", gadgetEntity.GetGadgetId())
			return
		}
		if strings.Contains(gadgetDataConfig.Name, "碎石堆") ||
			strings.Contains(gadgetDataConfig.ServerLuaScript, "SubfieldDrop_WoodenObject_Broken") {
			logger.Debug("物件破碎, entityId: %v, modifierChange: %v, uid: %v", entityId, modifierChange, player.PlayerId)
			g.KillEntity(player, scene, entity.GetId(), proto.PlayerDieType_PLAYER_DIE_GM)
		} else if strings.Contains(gadgetDataConfig.ServerLuaScript, "SubfieldDrop_Ore") {
			g.KillEntity(player, scene, entity.GetId(), proto.PlayerDieType_PLAYER_DIE_GM)
			g.CreateDropGadget(player, entity.GetPos(), 70900001, 233, 1)
		}
	}
}
