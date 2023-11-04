package gdconf

import (
	"hk4e/pkg/logger"
)

// MonsterData 怪物配置表
type MonsterData struct {
	MonsterId int32  `csv:"ID"`
	Name      string `csv:"名称$text_name_Name,omitempty"`
}

func (g *GameDataConfig) loadMonsterData() {
	g.MonsterDataMap = make(map[int32]*MonsterData)
	monsterDataList := make([]*MonsterData, 0)
	readTable[MonsterData](g.txtPrefix+"MonsterData.txt", &monsterDataList)
	for _, monsterData := range monsterDataList {
		g.MonsterDataMap[monsterData.MonsterId] = monsterData
	}
	logger.Info("MonsterData count: %v", len(g.MonsterDataMap))
}

func GetMonsterDataById(monsterId int32) *MonsterData {
	return CONF.MonsterDataMap[monsterId]
}

func GetMonsterDataMap() map[int32]*MonsterData {
	return CONF.MonsterDataMap
}
