package game

import (
	"fmt"
	"math"

	"hk4e/common/constant"
	"hk4e/gs/model"
	"hk4e/pkg/random"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"
)

const (
	PUBG_PHASE_START = 0
	PUBG_PHASE_II    = 2
	PUBG_PHASE_END   = 16
)

const (
	PUBG_PHASE_INV_TIME         = 180.0
	PUBG_FIRST_AREA_REDUCE_TIME = 600.0
)

type Pubg struct {
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

func (w *World) NewPubg() {
	w.pubg = &Pubg{
		world:                 w,
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
}

func (p *Pubg) GetAreaPointList() []*proto.MapMarkPoint {
	return p.areaPointList
}

func (p *Pubg) UpdateArea() {
	if p.areaReduceRadiusSpeed > 0.0 && p.blueAreaRadius > p.safeAreaRadius {
		p.blueAreaRadius -= p.areaReduceRadiusSpeed
		p.blueAreaCenterPos.X += p.areaReduceXSpeed
		p.blueAreaCenterPos.Z += p.areaReduceZSpeed
		p.SyncMapMarkArea()
	}
}

func (p *Pubg) IsInBlueArea(pos *model.Vector) bool {
	distance2D := math.Sqrt(
		(p.blueAreaCenterPos.X-pos.X)*(p.blueAreaCenterPos.X-pos.X) +
			(p.blueAreaCenterPos.Z-pos.Z)*(p.blueAreaCenterPos.Z-pos.Z),
	)
	return distance2D < p.blueAreaRadius
}

func (p *Pubg) RefreshArea() {
	info := ""
	if p.phase == PUBG_PHASE_START {
		info = fmt.Sprintf("安全区已出现，当前%v位存活玩家。", len(p.GetAlivePlayerList()))
		p.blueAreaCenterPos = &model.Vector{X: 500.0, Y: 0.0, Z: -500.0}
		p.blueAreaRadius = 2000.0
		p.safeAreaCenterPos = &model.Vector{X: 0.0, Y: 0.0, Z: 0.0}
		p.safeAreaRadius = 0.0
		TICK_MANAGER.CreateUserTimer(p.world.GetOwner().PlayerId, UserTimerActionPubgUpdateArea, PUBG_PHASE_INV_TIME)
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
			TICK_MANAGER.CreateUserTimer(p.world.GetOwner().PlayerId, UserTimerActionPubgUpdateArea, PUBG_PHASE_INV_TIME)
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
			TICK_MANAGER.CreateUserTimer(p.world.GetOwner().PlayerId, UserTimerActionPubgUpdateArea, uint32(invTime))
		case 0:
			info = fmt.Sprintf("安全区缩小完毕，进度%.1f%%。", float64(p.phase)/PUBG_PHASE_END*100.0)
			TICK_MANAGER.CreateUserTimer(p.world.GetOwner().PlayerId, UserTimerActionPubgUpdateArea, PUBG_PHASE_INV_TIME)
		}
	}
	p.SyncMapMarkArea()
	GAME.PlayerChatReq(p.world.GetOwner(), &proto.PlayerChatReq{ChatInfo: &proto.ChatInfo{Content: &proto.ChatInfo_Text{Text: info}}})
}

func (p *Pubg) SyncMapMarkArea() {
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

func (p *Pubg) GetAlivePlayerList() []*model.Player {
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
