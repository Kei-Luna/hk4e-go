package gdconf

import (
	"hk4e/pkg/logger"
)

// OpenStateData 开放状态配置表
type OpenStateData struct {
	OpenStateId    int32 `csv:"ID"`
	DefaultOpen    int32 `csv:"默认是否开启,omitempty"`
	AllowClientReq int32 `csv:"客户端能否发起功能开启,omitempty"`
}

func (g *GameDataConfig) loadOpenStateData() {
	g.OpenStateDataMap = make(map[int32]*OpenStateData)
	openStateDataList := make([]*OpenStateData, 0)
	readTable[OpenStateData](g.txtPrefix+"OpenStateData.txt", &openStateDataList)
	for _, openStateData := range openStateDataList {
		g.OpenStateDataMap[openStateData.OpenStateId] = openStateData
	}
	logger.Info("OpenStateData count: %v", len(g.OpenStateDataMap))
}

func GetOpenStateDataById(openStateId int32) *OpenStateData {
	return CONF.OpenStateDataMap[openStateId]
}

func GetOpenStateDataMap() map[int32]*OpenStateData {
	return CONF.OpenStateDataMap
}
