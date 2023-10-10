package gdconf

import (
	"hk4e/pkg/logger"
)

// pubg世界物件

type PubgWorldGadgetData struct {
	Id       int32   `csv:"Id"`
	GadgetId int32   `csv:"GadgetId"`
	X        float32 `csv:"X"`
	Y        float32 `csv:"Y"`
	Z        float32 `csv:"Z"`
}

func (g *GameDataConfig) loadPubgWorldGadgetData() {
	g.PubgWorldGadgetDataMap = make(map[int32]*PubgWorldGadgetData)
	pubgWorldGadgetDataList := make([]*PubgWorldGadgetData, 0)
	readExtCsv[PubgWorldGadgetData](g.extPrefix+"PubgWorldGadgetData.csv", &pubgWorldGadgetDataList)
	for _, pubgWorldGadgetData := range pubgWorldGadgetDataList {
		g.PubgWorldGadgetDataMap[pubgWorldGadgetData.Id] = pubgWorldGadgetData
	}
	logger.Info("PubgWorldGadgetData count: %v", len(g.PubgWorldGadgetDataMap))
}

func GetPubgWorldGadgetDataMap() map[int32]*PubgWorldGadgetData {
	return CONF.PubgWorldGadgetDataMap
}
