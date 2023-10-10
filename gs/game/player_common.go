package game

import (
	"math"
	"strings"
	"time"

	"hk4e/gdconf"

	"hk4e/common/constant"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

func (g *Game) PlayerSetPauseReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.PlayerSetPauseReq)
	isPaused := req.IsPaused
	player.Pause = isPaused
	g.SendMsg(cmd.PlayerSetPauseRsp, player.PlayerId, player.ClientSeq, new(proto.PlayerSetPauseRsp))
}

func (g *Game) TowerAllDataReq(player *model.Player, payloadMsg pb.Message) {
	towerAllDataRsp := &proto.TowerAllDataRsp{
		TowerScheduleId:        29,
		TowerFloorRecordList:   []*proto.TowerFloorRecord{{FloorId: 1001}},
		CurLevelRecord:         &proto.TowerCurLevelRecord{IsEmpty: true},
		NextScheduleChangeTime: 4294967295,
		FloorOpenTimeMap: map[uint32]uint32{
			1024: 1630486800,
			1025: 1630486800,
			1026: 1630486800,
			1027: 1630486800,
		},
		ScheduleStartTime: 1630486800,
	}
	g.SendMsg(cmd.TowerAllDataRsp, player.PlayerId, player.ClientSeq, towerAllDataRsp)
}

func (g *Game) ClientRttNotify(userId uint32, clientRtt uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	// logger.Debug("client rtt notify, uid: %v, rtt: %v", userId, clientRtt)
	player.ClientRTT = clientRtt
}

func (g *Game) ClientTimeNotify(userId uint32, clientTime uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	player.ClientTime = clientTime
	now := uint32(time.Now().Unix())
	// 客户端与服务器时间相差太过严重
	if math.Abs(float64(now-player.ClientTime)) > 600.0 {
		logger.Debug("abs of client time and server time above 600s, uid: %v", userId)
	}
	player.LastKeepaliveTime = now
}

func (g *Game) ServerAnnounceNotify(announceId uint32, announceMsg string) {
	for _, onlinePlayer := range USER_MANAGER.GetAllOnlineUserList() {
		now := uint32(time.Now().Unix())
		serverAnnounceNotify := &proto.ServerAnnounceNotify{
			AnnounceDataList: []*proto.AnnounceData{{
				ConfigId:              announceId,
				BeginTime:             now + 1,
				EndTime:               now + 2,
				CenterSystemText:      announceMsg,
				CenterSystemFrequency: 1,
			}},
		}
		g.SendMsg(cmd.ServerAnnounceNotify, onlinePlayer.PlayerId, 0, serverAnnounceNotify)
	}
}

func (g *Game) ServerAnnounceRevokeNotify(announceId uint32) {
	for _, onlinePlayer := range USER_MANAGER.GetAllOnlineUserList() {
		serverAnnounceRevokeNotify := &proto.ServerAnnounceRevokeNotify{
			ConfigIdList: []uint32{announceId},
		}
		g.SendMsg(cmd.ServerAnnounceRevokeNotify, onlinePlayer.PlayerId, 0, serverAnnounceRevokeNotify)
	}
}

func (g *Game) ToTheMoonEnterSceneReq(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("player ttm enter scene, uid: %v", player.PlayerId)
	req := payloadMsg.(*proto.ToTheMoonEnterSceneReq)
	_ = req
	g.SendMsg(cmd.ToTheMoonEnterSceneRsp, player.PlayerId, player.ClientSeq, new(proto.ToTheMoonEnterSceneRsp))
}

func (g *Game) PathfindingEnterSceneReq(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("player pf enter scene, uid: %v", player.PlayerId)
	req := payloadMsg.(*proto.PathfindingEnterSceneReq)
	_ = req
	g.SendMsg(cmd.PathfindingEnterSceneRsp, player.PlayerId, player.ClientSeq, new(proto.PathfindingEnterSceneRsp))
}

func (g *Game) QueryPathReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.QueryPathReq)
	queryPathRsp := &proto.QueryPathRsp{
		QueryId:     req.QueryId,
		QueryStatus: proto.QueryPathRsp_STATUS_SUCC,
		Corners:     []*proto.Vector{req.DestinationPos[0]},
	}
	g.SendMsg(cmd.QueryPathRsp, player.PlayerId, player.ClientSeq, queryPathRsp)
}

func (g *Game) ObstacleModifyNotify(player *model.Player, payloadMsg pb.Message) {
	ntf := payloadMsg.(*proto.ObstacleModifyNotify)
	_ = ntf
	// logger.Debug("ObstacleModifyNotify: %v, uid: %v", ntf, player.PlayerId)
}

// WorldPlayerRTTNotify 世界里所有玩家的网络延迟广播
func (g *Game) WorldPlayerRTTNotify(world *World) {
	worldPlayerRTTNotify := &proto.WorldPlayerRTTNotify{
		PlayerRttList: make([]*proto.PlayerRTTInfo, 0),
	}
	for _, worldPlayer := range world.GetAllPlayer() {
		playerRTTInfo := &proto.PlayerRTTInfo{Uid: worldPlayer.PlayerId, Rtt: worldPlayer.ClientRTT}
		worldPlayerRTTNotify.PlayerRttList = append(worldPlayerRTTNotify.PlayerRttList, playerRTTInfo)
	}
	g.SendToWorldA(world, cmd.WorldPlayerRTTNotify, 0, worldPlayerRTTNotify)
}

// WorldPlayerLocationNotify 多人世界其他玩家的坐标位置广播
func (g *Game) WorldPlayerLocationNotify(world *World) {
	worldPlayerLocationNotify := &proto.WorldPlayerLocationNotify{
		PlayerWorldLocList: make([]*proto.PlayerWorldLocationInfo, 0),
	}
	for _, worldPlayer := range world.GetAllPlayer() {
		playerWorldLocationInfo := &proto.PlayerWorldLocationInfo{
			SceneId: worldPlayer.SceneId,
			PlayerLoc: &proto.PlayerLocationInfo{
				Uid: worldPlayer.PlayerId,
				Pos: &proto.Vector{
					X: float32(worldPlayer.Pos.X),
					Y: float32(worldPlayer.Pos.Y),
					Z: float32(worldPlayer.Pos.Z),
				},
				Rot: &proto.Vector{
					X: float32(worldPlayer.Rot.X),
					Y: float32(worldPlayer.Rot.Y),
					Z: float32(worldPlayer.Rot.Z),
				},
			},
		}

		if WORLD_MANAGER.IsAiWorld(world) {
			playerWorldLocationInfo.PlayerLoc.Pos = new(proto.Vector)
			playerWorldLocationInfo.PlayerLoc.Rot = new(proto.Vector)
		}

		worldPlayerLocationNotify.PlayerWorldLocList = append(worldPlayerLocationNotify.PlayerWorldLocList, playerWorldLocationInfo)
	}
	g.SendToWorldA(world, cmd.WorldPlayerLocationNotify, 0, worldPlayerLocationNotify)
}

func (g *Game) ScenePlayerLocationNotify(world *World) {
	for _, scene := range world.GetAllScene() {
		scenePlayerLocationNotify := &proto.ScenePlayerLocationNotify{
			SceneId:        scene.id,
			PlayerLocList:  make([]*proto.PlayerLocationInfo, 0),
			VehicleLocList: make([]*proto.VehicleLocationInfo, 0),
		}
		for _, scenePlayer := range scene.GetAllPlayer() {
			// 玩家位置
			playerLocationInfo := &proto.PlayerLocationInfo{
				Uid: scenePlayer.PlayerId,
				Pos: &proto.Vector{
					X: float32(scenePlayer.Pos.X),
					Y: float32(scenePlayer.Pos.Y),
					Z: float32(scenePlayer.Pos.Z),
				},
				Rot: &proto.Vector{
					X: float32(scenePlayer.Rot.X),
					Y: float32(scenePlayer.Rot.Y),
					Z: float32(scenePlayer.Rot.Z),
				},
			}

			if WORLD_MANAGER.IsAiWorld(world) {
				playerLocationInfo.Pos = new(proto.Vector)
				playerLocationInfo.Rot = new(proto.Vector)
			}

			scenePlayerLocationNotify.PlayerLocList = append(scenePlayerLocationNotify.PlayerLocList, playerLocationInfo)
			// 载具位置
			for _, entityId := range scenePlayer.VehicleInfo.CreateEntityIdMap {
				entity := scene.GetEntity(entityId)
				// 确保实体类型是否为载具
				if entity != nil && entity.GetEntityType() == constant.ENTITY_TYPE_GADGET && entity.gadgetEntity.gadgetVehicleEntity != nil {
					vehicleLocationInfo := &proto.VehicleLocationInfo{
						Rot: &proto.Vector{
							X: float32(entity.GetRot().X),
							Y: float32(entity.GetRot().Y),
							Z: float32(entity.GetRot().Z),
						},
						EntityId: entity.id,
						CurHp:    entity.fightProp[constant.FIGHT_PROP_CUR_HP],
						OwnerUid: entity.gadgetEntity.gadgetVehicleEntity.ownerUid,
						Pos: &proto.Vector{
							X: float32(entity.GetPos().X),
							Y: float32(entity.GetPos().Y),
							Z: float32(entity.GetPos().Z),
						},
						UidList:  make([]uint32, 0, len(entity.gadgetEntity.gadgetVehicleEntity.memberMap)),
						GadgetId: entity.gadgetEntity.gadgetVehicleEntity.vehicleId,
						MaxHp:    entity.fightProp[constant.FIGHT_PROP_MAX_HP],
					}
					for _, p := range entity.gadgetEntity.gadgetVehicleEntity.memberMap {
						vehicleLocationInfo.UidList = append(vehicleLocationInfo.UidList, p.PlayerId)
					}
					scenePlayerLocationNotify.VehicleLocList = append(scenePlayerLocationNotify.VehicleLocList, vehicleLocationInfo)
				}
			}
		}
		g.SendToSceneA(scene, cmd.ScenePlayerLocationNotify, 0, scenePlayerLocationNotify)
	}
}

func (g *Game) SceneTimeNotify(world *World) {
	for _, scene := range world.GetAllScene() {
		for _, player := range scene.GetAllPlayer() {
			sceneTimeNotify := &proto.SceneTimeNotify{
				SceneId:   player.SceneId,
				SceneTime: uint64(scene.GetSceneTime()),
			}
			g.SendMsg(cmd.SceneTimeNotify, player.PlayerId, 0, sceneTimeNotify)
		}
	}
}

func (g *Game) PlayerTimeNotify(world *World) {
	for _, player := range world.GetAllPlayer() {
		playerTimeNotify := &proto.PlayerTimeNotify{
			IsPaused:   player.Pause,
			PlayerTime: uint64(player.TotalOnlineTime),
			ServerTime: uint64(time.Now().UnixMilli()),
		}
		g.SendMsg(cmd.PlayerTimeNotify, player.PlayerId, 0, playerTimeNotify)
	}
}

func (g *Game) PlayerGameTimeNotify(world *World) {
	for _, player := range world.GetAllPlayer() {
		scene := world.GetSceneById(player.SceneId)
		if scene == nil {
			logger.Error("scene is nil, sceneId: %v, uid: %v", player.SceneId, player.PlayerId)
			return
		}
		for _, scenePlayer := range scene.GetAllPlayer() {
			playerGameTimeNotify := &proto.PlayerGameTimeNotify{
				GameTime: scene.GetGameTime(),
				Uid:      scenePlayer.PlayerId,
			}
			g.SendMsg(cmd.PlayerGameTimeNotify, scenePlayer.PlayerId, 0, playerGameTimeNotify)
			// 设置玩家天气
			climateType := GAME.GetWeatherAreaClimate(player.WeatherInfo.WeatherAreaId)
			// 跳过相同的天气
			if climateType == player.WeatherInfo.ClimateType {
				return
			}
			GAME.SetPlayerWeather(player, player.WeatherInfo.WeatherAreaId, climateType)
		}
	}
}

func (g *Game) GmTalkReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GmTalkReq)
	logger.Info("GmTalkReq: %v", req.Msg)

	commandMessageInput := COMMAND_MANAGER.GetCommandMessageInput()
	if strings.Contains(req.Msg, "@@") {
		commandText := req.Msg
		commandText = strings.ReplaceAll(commandText, "@@", "")
		commandText = strings.ReplaceAll(commandText, " ", "")
		beginIndex := strings.Index(commandText, "(")
		endIndex := strings.Index(commandText, ")")
		if beginIndex == 0 || beginIndex == -1 || endIndex == -1 || beginIndex >= endIndex {
			g.SendMsg(cmd.GmTalkRsp, player.PlayerId, player.ClientSeq, &proto.GmTalkRsp{Retmsg: "命令解析失败", Msg: req.Msg})
			return
		}
		funcName := commandText[:beginIndex]
		paramList := strings.Split(commandText[beginIndex+1:endIndex], ",")
		commandMessageInput <- &CommandMessage{
			GMType:    SystemFuncGM,
			FuncName:  funcName,
			ParamList: paramList,
		}
	} else {
		commandMessageInput <- &CommandMessage{
			GMType:   DevClientGM,
			Executor: player,
			Text:     req.Msg,
		}
	}
	g.SendMsg(cmd.GmTalkRsp, player.PlayerId, player.ClientSeq, &proto.GmTalkRsp{Retmsg: "执行成功", Msg: req.Msg})
}

func (g *Game) PacketOpenStateUpdateNotify() *proto.OpenStateUpdateNotify {
	openStateUpdateNotify := &proto.OpenStateUpdateNotify{
		OpenStateMap: make(map[uint32]uint32),
	}
	// 先暂时开放全部功能模块
	for _, data := range gdconf.GetOpenStateDataMap() {
		openStateUpdateNotify.OpenStateMap[uint32(data.OpenStateId)] = 1
	}
	return openStateUpdateNotify
}
