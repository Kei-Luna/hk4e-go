package model

import (
	"hk4e/common/constant"
	"hk4e/gdconf"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SceneBlock struct {
	ID            primitive.ObjectID     `bson:"_id,omitempty"`
	Uid           uint32                 `bson:"uid"`
	BlockId       uint32                 `bson:"block_id"`
	SceneGroupMap map[uint32]*SceneGroup `bson:"scene_group_map"`
	IsNew         bool                   `bson:"-"`
}

type SceneGroup struct {
	GroupId        uint32           `bson:"group_id"`
	VariableMap    map[string]int32 `bson:"variable_map"`
	KillConfigMap  map[uint32]bool  `bson:"kill_config_map"`
	GadgetStateMap map[uint32]uint8 `bson:"gadget_state_map"`
}

func (p *Player) GetSceneGroupById(groupId uint32) *SceneGroup {
	groupConfig := gdconf.GetSceneGroup(int32(groupId))
	if groupConfig == nil {
		return nil
	}
	sceneBlock := p.SceneBlockMap[uint32(groupConfig.BlockId)]
	if sceneBlock == nil {
		return nil
	}
	sceneGroup, exist := sceneBlock.SceneGroupMap[groupId]
	if !exist {
		sceneGroup = &SceneGroup{
			GroupId:        groupId,
			VariableMap:    make(map[string]int32),
			KillConfigMap:  make(map[uint32]bool),
			GadgetStateMap: make(map[uint32]uint8),
		}
		sceneBlock.SceneGroupMap[groupId] = sceneGroup
	}
	return sceneGroup
}

func (g *SceneGroup) GetVariableByName(name string) int32 {
	return g.VariableMap[name]
}

func (g *SceneGroup) SetVariable(name string, value int32) {
	g.VariableMap[name] = value
}

func (g *SceneGroup) CheckVariableExist(name string) bool {
	_, exist := g.VariableMap[name]
	return exist
}

func (g *SceneGroup) AddKill(configId uint32) {
	g.KillConfigMap[configId] = true
}

func (g *SceneGroup) CheckIsKill(configId uint32) bool {
	_, exist := g.KillConfigMap[configId]
	return exist
}

func (g *SceneGroup) RemoveAllKill() {
	g.KillConfigMap = make(map[uint32]bool)
}

func (g *SceneGroup) GetGadgetState(configId uint32) uint8 {
	state, exist := g.GadgetStateMap[configId]
	if !exist {
		return constant.GADGET_STATE_DEFAULT
	}
	return state
}

func (g *SceneGroup) ChangeGadgetState(configId uint32, state uint8) {
	g.GadgetStateMap[configId] = state
}

func (g *SceneGroup) CheckGadgetExist(configId uint32) bool {
	_, exist := g.GadgetStateMap[configId]
	return exist
}
