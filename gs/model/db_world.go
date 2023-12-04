package model

import (
	"hk4e/gdconf"
)

type DbWorld struct {
	SceneMap    map[uint32]*DbScene
	MapMarkList []*MapMark
}

type DbScene struct {
	SceneId        uint32
	UnlockPointMap map[uint32]bool
	UnHidePointMap map[uint32]bool
	UnlockAreaMap  map[uint32]bool
}

type MapMark struct {
	SceneId   uint32
	Pos       *Vector
	PointType uint32
	Name      string
}

func (p *Player) GetDbWorld() *DbWorld {
	if p.DbWorld == nil {
		p.DbWorld = new(DbWorld)
	}
	if p.DbWorld.SceneMap == nil {
		p.DbWorld.SceneMap = make(map[uint32]*DbScene)
	}
	if p.DbWorld.MapMarkList == nil {
		p.DbWorld.MapMarkList = make([]*MapMark, 0)
	}
	return p.DbWorld
}

func (w *DbWorld) GetSceneById(sceneId uint32) *DbScene {
	scene, exist := w.SceneMap[sceneId]
	// 不存在自动创建场景
	if !exist {
		// 拒绝创建配置表中不存在的非法场景
		sceneDataConfig := gdconf.GetSceneDataById(int32(sceneId))
		if sceneDataConfig == nil {
			return nil
		}
		scene = new(DbScene)
		w.SceneMap[sceneId] = scene
	}
	if scene.SceneId == 0 {
		scene.SceneId = sceneId
	}
	if scene.UnlockPointMap == nil {
		scene.UnlockPointMap = make(map[uint32]bool)
	}
	if scene.UnHidePointMap == nil {
		scene.UnHidePointMap = make(map[uint32]bool)
	}
	if scene.UnlockAreaMap == nil {
		scene.UnlockAreaMap = make(map[uint32]bool)
	}
	return scene
}

func (s *DbScene) GetUnHidePointList() []uint32 {
	unHidePointList := make([]uint32, 0, len(s.UnHidePointMap))
	for pointId := range s.UnHidePointMap {
		unHidePointList = append(unHidePointList, pointId)
	}
	return unHidePointList
}

func (s *DbScene) GetUnlockPointList() []uint32 {
	unlockPointList := make([]uint32, 0, len(s.UnlockPointMap))
	for pointId := range s.UnlockPointMap {
		unlockPointList = append(unlockPointList, pointId)
	}
	return unlockPointList
}

func (s *DbScene) UnlockPoint(pointId uint32) {
	pointDataConfig := gdconf.GetScenePointBySceneIdAndPointId(int32(s.SceneId), int32(pointId))
	if pointDataConfig == nil {
		return
	}
	s.UnlockPointMap[pointId] = true
	// 隐藏锚点取消隐藏
	if pointDataConfig.IsModelHidden {
		s.UnHidePointMap[pointId] = true
	}
}

func (s *DbScene) CheckPointUnlock(pointId uint32) bool {
	_, exist := s.UnlockPointMap[pointId]
	return exist
}

func (s *DbScene) GetUnlockAreaList() []uint32 {
	unlockAreaList := make([]uint32, 0, len(s.UnlockAreaMap))
	for areaId := range s.UnlockAreaMap {
		unlockAreaList = append(unlockAreaList, areaId)
	}
	return unlockAreaList
}

func (s *DbScene) UnlockArea(areaId uint32) {
	exist := false
	for _, worldAreaData := range gdconf.GetWorldAreaDataMap() {
		if uint32(worldAreaData.SceneId) == s.SceneId && uint32(worldAreaData.AreaId1) == areaId {
			exist = true
			break
		}
	}
	if !exist {
		return
	}
	s.UnlockAreaMap[areaId] = true
}

func (s *DbScene) CheckAreaUnlock(areaId uint32) bool {
	_, exist := s.UnlockAreaMap[areaId]
	return exist
}
