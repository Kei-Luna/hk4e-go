package game

import (
	"encoding/base64"
	"fmt"
	"math"

	"hk4e/common/constant"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/pkg/random"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"
)

const (
	PUBG_PHASE_WAIT  = -1
	PUBG_PHASE_START = 0
	PUBG_PHASE_II    = 2
	PUBG_PHASE_END   = 16
)

const (
	PUBG_PHASE_INV_TIME         = 180.0
	PUBG_FIRST_AREA_REDUCE_TIME = 600.0
)

const (
	PUBG_ATK = 100.0
	PUBG_HP  = 1000.0
)

// PluginPubg pubg游戏插件
type PluginPubg struct {
	*Plugin
	world                    *World                // 世界对象
	blueAreaCenterPos        *model.Vector         // 蓝区中心点
	blueAreaRadius           float64               // 蓝区半径
	safeAreaCenterPos        *model.Vector         // 安全区中心点
	safeAreaRadius           float64               // 安全区半径
	phase                    int                   // 阶段
	areaReduceRadiusSpeed    float64               // 缩圈半径速度
	areaReduceXSpeed         float64               // 缩圈X速度
	areaReduceZSpeed         float64               // 缩圈Z速度
	areaPointList            []*proto.MapMarkPoint // 客户端区域地图坐标列表
	entityIdWorldGadgetIdMap map[uint32]int32      // 实体id世界物件id映射集合
}

func NewPluginPubg() *PluginPubg {
	p := &PluginPubg{
		Plugin:                   NewPlugin(),
		world:                    nil,
		blueAreaCenterPos:        &model.Vector{X: 0.0, Y: 0.0, Z: 0.0},
		blueAreaRadius:           0.0,
		safeAreaCenterPos:        &model.Vector{X: 0.0, Y: 0.0, Z: 0.0},
		safeAreaRadius:           0.0,
		phase:                    PUBG_PHASE_WAIT,
		areaReduceRadiusSpeed:    0.0,
		areaReduceXSpeed:         0.0,
		areaReduceZSpeed:         0.0,
		areaPointList:            make([]*proto.MapMarkPoint, 0),
		entityIdWorldGadgetIdMap: make(map[uint32]int32),
	}
	return p
}

// OnEnable 插件启用生命周期
func (p *PluginPubg) OnEnable() {
	// 监听事件
	p.ListenEvent(PluginEventIdPlayerKillAvatar, PluginEventPriorityNormal, p.EventKillAvatar)
	p.ListenEvent(PluginEventIdMarkMap, PluginEventPriorityNormal, p.EventMarkMap)
	p.ListenEvent(PluginEventIdAvatarDieAnimationEnd, PluginEventPriorityNormal, p.EventAvatarDieAnimationEnd)
	p.ListenEvent(PluginEventIdGadgetInteract, PluginEventPriorityNormal, p.EventGadgetInteract)
	p.ListenEvent(PluginEventIdPostEnterScene, PluginEventPriorityNormal, p.EventPostEnterScene)
	// 添加全局定时器
	p.AddGlobalTick(PluginGlobalTickSecond, p.GlobalTickPubg)
	p.AddGlobalTick(PluginGlobalTickHourChange, p.GlobalTickHourStart)
	// 注册命令
	p.RegCommandController(p.NewPubgCommandController())
}

/************************************************** 事件监听 **************************************************/

// EventKillAvatar 角色被杀死事件
func (p *PluginPubg) EventKillAvatar(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventKillAvatar)
	player := event.Player
	// 确保游戏开启
	if !p.IsStartPubg() || p.world.id != player.WorldId {
		return
	}
	p.CreateUserTimer(player.PlayerId, 10, p.UserTimerPubgDieExit)
}

// EventMarkMap 地图标点事件
func (p *PluginPubg) EventMarkMap(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventMarkMap)
	player := event.Player
	// 确保游戏开启
	if !p.IsStartPubg() || p.world.id != player.WorldId {
		return
	}
	GAME.SendMsg(cmd.MarkMapRsp, player.PlayerId, player.ClientSeq, &proto.MarkMapRsp{MarkList: p.GetAreaPointList()})
	event.Cancel()
}

// EventAvatarDieAnimationEnd 角色死亡动画结束事件
func (p *PluginPubg) EventAvatarDieAnimationEnd(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventAvatarDieAnimationEnd)
	player := event.Player
	// 确保游戏开启
	if !p.IsStartPubg() || p.world.id != player.WorldId {
		return
	}
	alivePlayerNum := len(p.GetAlivePlayerList())
	info := fmt.Sprintf("『%v』死亡了，剩余%v位存活玩家。", player.NickName, alivePlayerNum)
	GAME.PlayerChatReq(p.world.GetOwner(), &proto.PlayerChatReq{ChatInfo: &proto.ChatInfo{Content: &proto.ChatInfo_Text{Text: info}}})
	GAME.SendMsg(cmd.AvatarDieAnimationEndRsp, player.PlayerId, player.ClientSeq, &proto.AvatarDieAnimationEndRsp{SkillId: event.Req.SkillId, DieGuid: event.Req.DieGuid})
	event.Cancel()
}

func (p *PluginPubg) EventGadgetInteract(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventGadgetInteract)
	player := event.Player
	// 确保游戏开启
	if !p.IsStartPubg() || p.world.id != player.WorldId {
		return
	}
	req := event.Req
	worldGadgetId, exist := p.entityIdWorldGadgetIdMap[req.GadgetEntityId]
	if exist {
		dbAvatar := player.GetDbAvatar()
		avatarId := p.world.GetPlayerActiveAvatarId(player)
		avatar := dbAvatar.AvatarMap[avatarId]
		pubgWorldGadgetDataConfig := gdconf.GetPubgWorldGadgetDataById(worldGadgetId)
		switch pubgWorldGadgetDataConfig.Type {
		case gdconf.PubgWorldGadgetTypeIncAtk:
			avatar.FightPropMap[constant.FIGHT_PROP_BASE_ATTACK] += float32(pubgWorldGadgetDataConfig.Param[0])
			avatar.FightPropMap[constant.FIGHT_PROP_CUR_ATTACK] += float32(pubgWorldGadgetDataConfig.Param[0])
			// 提示玩家
			info := fmt.Sprintf("你的角色攻击力增加：%v，增加后的攻击力：%v。", pubgWorldGadgetDataConfig.Param[0], avatar.FightPropMap[constant.FIGHT_PROP_BASE_ATTACK])
			GAME.PlayerChatReq(p.world.GetOwner(), &proto.PlayerChatReq{ChatInfo: &proto.ChatInfo{Content: &proto.ChatInfo_Text{Text: info}}})
		case gdconf.PubgWorldGadgetTypeIncHp:
			avatar.FightPropMap[constant.FIGHT_PROP_CUR_HP] += float32(pubgWorldGadgetDataConfig.Param[0])
			if avatar.FightPropMap[constant.FIGHT_PROP_CUR_HP] > avatar.FightPropMap[constant.FIGHT_PROP_MAX_HP] {
				avatar.FightPropMap[constant.FIGHT_PROP_CUR_HP] = avatar.FightPropMap[constant.FIGHT_PROP_MAX_HP]
			}
			// 提示玩家
			info := fmt.Sprintf("你的角色生命值增加：%v，目前为：%v。", pubgWorldGadgetDataConfig.Param[0], avatar.FightPropMap[constant.FIGHT_PROP_CUR_HP])
			GAME.PlayerChatReq(p.world.GetOwner(), &proto.PlayerChatReq{ChatInfo: &proto.ChatInfo{Content: &proto.ChatInfo_Text{Text: info}}})
		}
		GAME.SendMsg(cmd.AvatarFightPropUpdateNotify, player.PlayerId, player.ClientSeq, &proto.AvatarFightPropUpdateNotify{
			AvatarGuid:   avatar.Guid,
			FightPropMap: avatar.FightPropMap,
		})
	}
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		logger.Error("get world is nil, worldId: %v, uid: %v", player.WorldId, player.PlayerId)
		return
	}
	scene := world.GetSceneById(player.SceneId)
	GAME.KillEntity(player, scene, req.GadgetEntityId, proto.PlayerDieType_PLAYER_DIE_NONE)
	rsp := &proto.GadgetInteractRsp{
		GadgetEntityId: req.GadgetEntityId,
		GadgetId:       req.GadgetId,
		OpType:         req.OpType,
		InteractType:   proto.InteractType_INTERACT_GATHER,
	}
	GAME.SendMsg(cmd.GadgetInteractRsp, player.PlayerId, player.ClientSeq, rsp)
	event.Cancel()
}

func (p *PluginPubg) EventPostEnterScene(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventPostEnterScene)
	player := event.Player
	// 确保游戏开启
	if !p.IsStartPubg() || p.world.id != player.WorldId {
		return
	}
	// 开启GM按钮 隐藏多人世界玩家位置地图标记
	// local btnGm = CS.UnityEngine.GameObject.Find("/Canvas/Pages/InLevelMainPage/GrpMainPage/GrpMainBtn/GrpMainToggle/GrpTopPanel/BtnGm")
	// btnGm:SetActive(true)
	// local miniMapMarkLayer3 = CS.UnityEngine.GameObject.Find("/Canvas/Pages/InLevelMainPage/GrpMainPage/MapInfo/GrpMiniMap/GrpMap/MarkContainer/Layer3")
	// miniMapMarkLayer3:SetActive(false)
	// local mapMarkLayer3 = CS.UnityEngine.GameObject.Find("/Canvas/Pages/InLevelMapPage/GrpMap/MarkContainer/Layer3")
	// mapMarkLayer3:SetActive(false)
	luac, err := base64.StdEncoding.DecodeString("G0x1YVMBGZMNChoKBAQICHhWAAAAAAAAAAAAAAAod0ABDkBhaV93b3JsZC5sdWEAAAAAAAAAAAABBhwAAAAkAEAAKUBAACmAQAApwEAAVgABACyAAAFdQEEA2ACAAGxAgAFkAEAAaUDAAGmAwABpwMAAloABAGyAAAGdQMEAGAEAAKxAgAGkAEAAqUBAAamAQAGpwEAB1sABAKyAAAHdQEEBWAEAAOxAgAEZAIAACAAAAAQDQ1MEDFVuaXR5RW5naW5lBAtHYW1lT2JqZWN0BAVGaW5kFFUvQ2FudmFzL1BhZ2VzL0luTGV2ZWxNYWluUGFnZS9HcnBNYWluUGFnZS9HcnBNYWluQnRuL0dycE1haW5Ub2dnbGUvR3JwVG9wUGFuZWwvQnRuR20EClNldEFjdGl2ZRRZL0NhbnZhcy9QYWdlcy9JbkxldmVsTWFpblBhZ2UvR3JwTWFpblBhZ2UvTWFwSW5mby9HcnBNaW5pTWFwL0dycE1hcC9NYXJrQ29udGFpbmVyL0xheWVyMxQ5L0NhbnZhcy9QYWdlcy9JbkxldmVsTWFwUGFnZS9HcnBNYXAvTWFya0NvbnRhaW5lci9MYXllcjMBAAAAAQAAAAAAHAAAAAEAAAABAAAAAQAAAAEAAAABAAAAAQAAAAIAAAACAAAAAgAAAAMAAAADAAAAAwAAAAMAAAADAAAAAwAAAAQAAAAEAAAABAAAAAUAAAAFAAAABQAAAAUAAAAFAAAABQAAAAYAAAAGAAAABgAAAAYAAAADAAAABmJ0bkdtBgAAABwAAAASbWluaU1hcE1hcmtMYXllcjMPAAAAHAAAAA5tYXBNYXJrTGF5ZXIzGAAAABwAAAABAAAABV9FTlY=")
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

/************************************************** 全局定时器 **************************************************/

// GlobalTickPubg pubg游戏定时器
func (p *PluginPubg) GlobalTickPubg() {
	world := p.world
	if world == nil {
		return
	}
	// 确保游戏开启
	if !p.IsStartPubg() {
		return
	}
	p.UpdateArea()
	scene := world.GetSceneById(world.GetOwner().SceneId)
	for _, scenePlayer := range scene.GetAllPlayer() {
		if !p.IsInBlueArea(scenePlayer.Pos) {
			GAME.handleEvtBeingHit(scenePlayer, scene, &proto.EvtBeingHitInfo{
				AttackResult: &proto.AttackResult{
					AttackerId: 0,
					DefenseId:  world.GetPlayerWorldAvatarEntityId(scenePlayer, world.GetPlayerActiveAvatarId(scenePlayer)),
					Damage:     10,
				},
			})
		}
	}
	alivePlayerList := p.GetAlivePlayerList()
	if len(alivePlayerList) <= 1 {
		if len(alivePlayerList) == 1 {
			info := fmt.Sprintf("『%v』大吉大利，今晚吃鸡。", alivePlayerList[0].NickName)
			GAME.PlayerChatReq(world.GetOwner(), &proto.PlayerChatReq{ChatInfo: &proto.ChatInfo{Content: &proto.ChatInfo_Text{Text: info}}})
		}
		p.StopPubg()
	}
}

// GlobalTickHourStart 每小时开启pubg游戏
func (p *PluginPubg) GlobalTickHourStart() {
	p.StartPubg()
}

/************************************************** 用户定时器 **************************************************/

// UserTimerPubgEnd pubg游戏结束后执行定时器
func (p *PluginPubg) UserTimerPubgEnd(player *model.Player, data []any) {
	logger.Debug("PubgEnd")
	world := p.world
	if world == nil {
		return
	}
	for _, worldPlayer := range world.GetAllPlayer() {
		if worldPlayer.PlayerId == world.GetOwner().PlayerId {
			continue
		}
		GAME.ReLoginPlayer(worldPlayer.PlayerId, true)
	}
}

// UserTimerPubgUpdateArea 更新游戏区域
func (p *PluginPubg) UserTimerPubgUpdateArea(player *model.Player, data []any) {
	logger.Debug("PubgUpdateArea")
	p.phase++
	p.RefreshArea()
}

// UserTimerPubgDieExit pubg死亡离开
func (p *PluginPubg) UserTimerPubgDieExit(player *model.Player, data []any) {
	logger.Debug("PubgDieExit")
	GAME.ReLoginPlayer(player.PlayerId, true)
}

/************************************************** 命令控制器 **************************************************/

// pubg游戏命令

func (p *PluginPubg) NewPubgCommandController() *CommandController {
	return &CommandController{
		Name:        "PUBG游戏",
		AliasList:   []string{"pubg"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>测试pubg游戏</color>",
		UsageList: []string{
			"{alias} <start/stop> 开始或关闭pubg游戏",
		},
		Perm: CommandPermNormal,
		Func: p.PubgCommand,
	}
}

func (p *PluginPubg) PubgCommand(c *CommandContent) bool {
	var mode string // 模式

	return c.Dynamic("string", func(param any) bool {
		// 模式
		mode = param.(string)
		return true
	}).Execute(func() bool {
		switch mode {
		case "start":
			// 开始游戏
			p.StartPubg()
			c.SendSuccMessage(c.Executor, "已开始PUBG游戏。")
		case "stop":
			// 结束游戏
			p.StopPubg()
			c.SendSuccMessage(c.Executor, "已结束PUBG游戏。")
		default:
			return false
		}
		return true
	})
}

/************************************************** 插件功能 **************************************************/

// StartPubg 开始pubg游戏
func (p *PluginPubg) StartPubg() {
	if p.IsStartPubg() {
		return
	}
	world := WORLD_MANAGER.GetAiWorld()
	p.world = world
	for _, pubgWorldGadgetDataConfig := range gdconf.GetPubgWorldGadgetDataMap() {
		rn := random.GetRandomInt32(1, 100)
		if rn > pubgWorldGadgetDataConfig.Probability {
			continue
		}
		entityId := GAME.CreateGadget(
			p.world.GetOwner(),
			&model.Vector{X: float64(pubgWorldGadgetDataConfig.X), Y: float64(pubgWorldGadgetDataConfig.Y), Z: float64(pubgWorldGadgetDataConfig.Z)},
			uint32(pubgWorldGadgetDataConfig.GadgetId),
			nil,
		)
		p.entityIdWorldGadgetIdMap[entityId] = pubgWorldGadgetDataConfig.WorldGadgetId
	}
	p.phase = PUBG_PHASE_START
	p.RefreshArea()
	world.chatMsgList = make([]*proto.ChatInfo, 0)
	for _, player := range world.GetAllPlayer() {
		dbAvatar := player.GetDbAvatar()
		avatarId := p.world.GetPlayerActiveAvatarId(player)
		avatar := dbAvatar.AvatarMap[avatarId]
		for k := range avatar.FightPropMap {
			avatar.FightPropMap[k] = 0.0
		}
		avatar.FightPropMap[constant.FIGHT_PROP_BASE_HP] = PUBG_HP
		avatar.FightPropMap[constant.FIGHT_PROP_MAX_HP] = PUBG_HP
		avatar.FightPropMap[constant.FIGHT_PROP_CUR_HP] = PUBG_HP
		avatar.FightPropMap[constant.FIGHT_PROP_BASE_ATTACK] = PUBG_ATK
		avatar.FightPropMap[constant.FIGHT_PROP_CUR_ATTACK] = PUBG_ATK
		GAME.SendMsg(cmd.AvatarFightPropUpdateNotify, player.PlayerId, player.ClientSeq, &proto.AvatarFightPropUpdateNotify{
			AvatarGuid:   avatar.Guid,
			FightPropMap: avatar.FightPropMap,
		})
	}
}

// StopPubg 结束pubg游戏
func (p *PluginPubg) StopPubg() {
	if !p.IsStartPubg() {
		return
	}
	p.phase = PUBG_PHASE_WAIT
	p.CreateUserTimer(p.world.GetOwner().PlayerId, 60, p.UserTimerPubgEnd)
}

// IsStartPubg pubg游戏是否开启
func (p *PluginPubg) IsStartPubg() bool {
	return p.world != nil && p.phase != PUBG_PHASE_WAIT
}

// GetAreaPointList 获取游戏区域标点列表
func (p *PluginPubg) GetAreaPointList() []*proto.MapMarkPoint {
	return p.areaPointList
}

// UpdateArea 更新游戏区域
func (p *PluginPubg) UpdateArea() {
	if p.areaReduceRadiusSpeed > 0.0 && p.blueAreaRadius > p.safeAreaRadius {
		p.blueAreaRadius -= p.areaReduceRadiusSpeed
		p.blueAreaCenterPos.X += p.areaReduceXSpeed
		p.blueAreaCenterPos.Z += p.areaReduceZSpeed
		p.SyncMapMarkArea()
	}
}

// IsInBlueArea 是否在蓝圈内
func (p *PluginPubg) IsInBlueArea(pos *model.Vector) bool {
	distance2D := math.Sqrt(
		(p.blueAreaCenterPos.X-pos.X)*(p.blueAreaCenterPos.X-pos.X) +
			(p.blueAreaCenterPos.Z-pos.Z)*(p.blueAreaCenterPos.Z-pos.Z),
	)
	return distance2D < p.blueAreaRadius
}

// RefreshArea 刷新游戏区域
func (p *PluginPubg) RefreshArea() {
	info := ""
	if p.phase == PUBG_PHASE_START {
		info = fmt.Sprintf("安全区已生成，当前%v位存活玩家。", len(p.GetAlivePlayerList()))
		p.blueAreaCenterPos = &model.Vector{X: 500.0, Y: 0.0, Z: -500.0}
		p.blueAreaRadius = 2000.0
		p.safeAreaCenterPos = &model.Vector{X: 0.0, Y: 0.0, Z: 0.0}
		p.safeAreaRadius = 0.0
		p.CreateUserTimer(p.world.GetOwner().PlayerId, PUBG_PHASE_INV_TIME, p.UserTimerPubgUpdateArea)
	} else if p.phase == PUBG_PHASE_END {
		info = "安全区已消失。"
		p.blueAreaRadius = 0.0
		p.safeAreaRadius = 0.0
	} else {
		switch p.phase % 3 {
		case 1:
			info = fmt.Sprintf("新的安全区已出现，进度%.1f%%。", float64(p.phase)/PUBG_PHASE_END*100.0)
			p.safeAreaCenterPos = &model.Vector{
				X: p.blueAreaCenterPos.X + random.GetRandomFloat64(-(p.blueAreaRadius*0.7/2.0), p.blueAreaRadius*0.7/2.0),
				Y: 0.0,
				Z: p.blueAreaCenterPos.Z + random.GetRandomFloat64(-(p.blueAreaRadius*0.7/2.0), p.blueAreaRadius*0.7/2.0),
			}
			p.safeAreaRadius = p.blueAreaRadius / 2.0
			p.areaReduceRadiusSpeed = 0.0
			p.CreateUserTimer(p.world.GetOwner().PlayerId, PUBG_PHASE_INV_TIME, p.UserTimerPubgUpdateArea)
		case 2:
			info = fmt.Sprintf("安全区正在缩小，进度%.1f%%。", float64(p.phase)/PUBG_PHASE_END*100.0)
			invTime := 0.0
			if p.phase == PUBG_PHASE_II {
				invTime = PUBG_FIRST_AREA_REDUCE_TIME
			} else {
				invTime = PUBG_PHASE_INV_TIME
			}
			p.areaReduceRadiusSpeed = (p.blueAreaRadius - p.safeAreaRadius) / invTime
			p.areaReduceXSpeed = (p.safeAreaCenterPos.X - p.blueAreaCenterPos.X) / invTime
			p.areaReduceZSpeed = (p.safeAreaCenterPos.Z - p.blueAreaCenterPos.Z) / invTime
			p.CreateUserTimer(p.world.GetOwner().PlayerId, uint32(invTime), p.UserTimerPubgUpdateArea)
		case 0:
			info = fmt.Sprintf("安全区缩小完毕，进度%.1f%%。", float64(p.phase)/PUBG_PHASE_END*100.0)
			p.CreateUserTimer(p.world.GetOwner().PlayerId, PUBG_PHASE_INV_TIME, p.UserTimerPubgUpdateArea)
		}
	}
	p.SyncMapMarkArea()
	GAME.PlayerChatReq(p.world.GetOwner(), &proto.PlayerChatReq{ChatInfo: &proto.ChatInfo{Content: &proto.ChatInfo_Text{Text: info}}})
}

// SyncMapMarkArea 同步地图标点区域
func (p *PluginPubg) SyncMapMarkArea() {
	p.areaPointList = make([]*proto.MapMarkPoint, 0)
	if p.blueAreaRadius > 0.0 {
		for angleStep := 0; angleStep < 360; angleStep += 5 {
			x := p.blueAreaRadius*math.Cos(float64(angleStep)/360.0*2*math.Pi) + p.blueAreaCenterPos.X
			z := p.blueAreaRadius*math.Sin(float64(angleStep)/360.0*2*math.Pi) + p.blueAreaCenterPos.Z
			p.areaPointList = append(p.areaPointList, &proto.MapMarkPoint{
				SceneId:   3,
				Name:      "",
				Pos:       &proto.Vector{X: float32(x), Y: 0, Z: float32(z)},
				PointType: proto.MapMarkPointType_SPECIAL,
			})
		}
	}
	if p.safeAreaRadius > 0.0 {
		for angleStep := 0; angleStep < 360; angleStep += 5 {
			x := p.safeAreaRadius*math.Cos(float64(angleStep)/360.0*2*math.Pi) + p.safeAreaCenterPos.X
			z := p.safeAreaRadius*math.Sin(float64(angleStep)/360.0*2*math.Pi) + p.safeAreaCenterPos.Z
			p.areaPointList = append(p.areaPointList, &proto.MapMarkPoint{
				SceneId:   3,
				Name:      "",
				Pos:       &proto.Vector{X: float32(x), Y: 0, Z: float32(z)},
				PointType: proto.MapMarkPointType_COLLECTION,
			})
		}
	}
	for _, player := range p.world.GetAllPlayer() {
		GAME.SendMsg(cmd.AllMarkPointNotify, player.PlayerId, player.ClientSeq, &proto.AllMarkPointNotify{MarkList: p.areaPointList})
	}
}

// GetAlivePlayerList 获取存活玩家列表
func (p *PluginPubg) GetAlivePlayerList() []*model.Player {
	scene := p.world.GetSceneById(p.world.GetOwner().SceneId)
	alivePlayerList := make([]*model.Player, 0)
	for _, scenePlayer := range scene.GetAllPlayer() {
		if scenePlayer.PlayerId == p.world.GetOwner().PlayerId {
			continue
		}
		avatarEntityId := p.world.GetPlayerWorldAvatarEntityId(scenePlayer, p.world.GetPlayerActiveAvatarId(scenePlayer))
		entity := scene.GetEntity(avatarEntityId)
		if entity.GetFightProp()[constant.FIGHT_PROP_CUR_HP] <= 0.0 {
			continue
		}
		alivePlayerList = append(alivePlayerList, scenePlayer)
	}
	return alivePlayerList
}
