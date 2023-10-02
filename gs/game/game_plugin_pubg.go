package game

import (
	"fmt"
	"hk4e/common/constant"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/pkg/random"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"
	"math"
)

const (
	PUBG_PHASE_WAIT = iota
	PUBG_PHASE_START
	PUBG_PHASE_II
	PUBG_PHASE_END
)

const (
	PUBG_PHASE_INV_TIME         = 180.0
	PUBG_FIRST_AREA_REDUCE_TIME = 600.0
)

const PLUGIN_NAME_PUBG = "pubg" // 插件名

// PluginPubg pubg游戏插件
type PluginPubg struct {
	*Plugin
	world                 *World                // 世界对象
	blueAreaCenterPos     *model.Vector         // 蓝区中心点
	blueAreaRadius        float64               // 蓝区半径
	safeAreaCenterPos     *model.Vector         // 安全区中心点
	safeAreaRadius        float64               // 安全区半径
	phase                 int                   // 阶段
	areaReduceRadiusSpeed float64               // 缩圈半径速度
	areaReduceXSpeed      float64               // 缩圈X速度
	areaReduceZSpeed      float64               // 缩圈Z速度
	areaPointList         []*proto.MapMarkPoint // 客户端区域地图坐标列表
}

func NewPluginPubg() *PluginPubg {
	p := &PluginPubg{
		Plugin:                NewPlugin(PLUGIN_NAME_PUBG),
		world:                 nil,
		blueAreaCenterPos:     &model.Vector{X: 0.0, Y: 0.0, Z: 0.0},
		blueAreaRadius:        0.0,
		safeAreaCenterPos:     &model.Vector{X: 0.0, Y: 0.0, Z: 0.0},
		safeAreaRadius:        0.0,
		phase:                 PUBG_PHASE_START,
		areaReduceRadiusSpeed: 0.0,
		areaReduceXSpeed:      0.0,
		areaReduceZSpeed:      0.0,
		areaPointList:         make([]*proto.MapMarkPoint, 0),
	}
	// 注册事件
	p.ListenEvent(PluginEventIdPlayerKillAvatar, PluginEventPriorityNormal, p.EventKillAvatar)
	p.ListenEvent(PluginEventIdMarkMap, PluginEventPriorityNormal, p.EventMarkMap)
	p.ListenEvent(PluginEventIdAvatarDieAnimationEnd, PluginEventPriorityNormal, p.EventAvatarDieAnimationEnd)
	// 添加全局定时器
	p.AddGlobalTick(PluginGlobalTickSecond, p.GlobalTickPubg)
	p.AddGlobalTick(PluginGlobalTickHour, p.GlobalTickHourStart)
	return p
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
	player.PubgRank += uint32(100 - alivePlayerNum)
	GAME.SendMsg(cmd.AvatarDieAnimationEndRsp, player.PlayerId, player.ClientSeq, &proto.AvatarDieAnimationEndRsp{SkillId: event.Req.SkillId, DieGuid: event.Req.DieGuid})
	event.Cancel()
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

/************************************************** 插件功能 **************************************************/

// StartPubg 开始pubg游戏
func (p *PluginPubg) StartPubg() {
	if p.IsStartPubg() {
		return
	}
	world := WORLD_MANAGER.GetAiWorld()
	p.world = world
	p.phase = PUBG_PHASE_START
	p.RefreshArea()
	world.chatMsgList = make([]*proto.ChatInfo, 0)
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
	return p.world != nil && p.phase != PUBG_PHASE_WAIT && p.phase != PUBG_PHASE_END
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
