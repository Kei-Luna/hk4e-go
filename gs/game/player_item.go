package game

import (
	"hk4e/common/constant"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"
)

/************************************************** 接口请求 **************************************************/

/************************************************** 游戏功能 **************************************************/

type ChangeItem struct {
	ItemId      uint32
	ChangeCount uint32
}

func (g *Game) GetAllItemDataConfig() map[int32]*gdconf.ItemData {
	allItemDataConfig := make(map[int32]*gdconf.ItemData)
	for itemId, itemData := range gdconf.GetItemDataMap() {
		if itemData.Type == constant.ITEM_TYPE_WEAPON {
			// 排除武器
			continue
		}
		if itemData.Type == constant.ITEM_TYPE_RELIQUARY {
			// 排除圣遗物
			continue
		}
		allItemDataConfig[itemId] = itemData
	}
	return allItemDataConfig
}

func (g *Game) GetPlayerItemCount(userId uint32, itemId uint32) uint32 {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return 0
	}
	prop, ok := constant.VIRTUAL_ITEM_PROP[itemId]
	if ok {
		value := player.PropertiesMap[prop]
		return value
	} else {
		dbItem := player.GetDbItem()
		value := dbItem.GetItemCount(itemId)
		return value
	}
}

// AddPlayerItem 玩家添加物品
func (g *Game) AddPlayerItem(userId uint32, itemList []*ChangeItem, isHint bool, hintReason uint16) bool {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return false
	}
	dbItem := player.GetDbItem()
	playerPropNotify := &proto.PlayerPropNotify{
		PropMap: make(map[uint32]*proto.PropValue),
	}
	storeItemChangeNotify := &proto.StoreItemChangeNotify{
		StoreType: proto.StoreType_STORE_PACK,
		ItemList:  make([]*proto.Item, 0),
	}
	for _, changeItem := range itemList {
		prop, exist := constant.VIRTUAL_ITEM_PROP[changeItem.ItemId]
		if exist {
			// 物品为虚拟物品 角色属性物品数量增加
			player.PropertiesMap[prop] += changeItem.ChangeCount
			playerPropNotify.PropMap[uint32(prop)] = &proto.PropValue{
				Type: uint32(prop),
				Val:  int64(player.PropertiesMap[prop]),
				Value: &proto.PropValue_Ival{
					Ival: int64(player.PropertiesMap[prop]),
				},
			}
			// 特殊属性变化处理函数
			switch changeItem.ItemId {
			case constant.ITEM_ID_PLAYER_EXP:
				// 冒险阅历
				g.HandlePlayerExpAdd(userId)
			}
		} else {
			// 物品为普通物品 直接进背包
			// 校验背包物品容量 目前物品包括材料和家具
			if dbItem.GetItemMapLen() > constant.STORE_PACK_LIMIT_MATERIAL+constant.STORE_PACK_LIMIT_FURNITURE {
				return false
			}
			dbItem.AddItem(player, changeItem.ItemId, changeItem.ChangeCount)
		}
		pbItem := &proto.Item{
			ItemId: changeItem.ItemId,
			Guid:   dbItem.GetItemGuid(changeItem.ItemId),
			Detail: &proto.Item_Material{
				Material: &proto.Material{
					Count: dbItem.GetItemCount(changeItem.ItemId),
				},
			},
		}
		storeItemChangeNotify.ItemList = append(storeItemChangeNotify.ItemList, pbItem)
	}
	if len(playerPropNotify.PropMap) > 0 {
		g.SendMsg(cmd.PlayerPropNotify, userId, player.ClientSeq, playerPropNotify)
	}
	g.SendMsg(cmd.StoreItemChangeNotify, userId, player.ClientSeq, storeItemChangeNotify)
	if isHint {
		if hintReason == 0 {
			hintReason = uint16(proto.ActionReasonType_ACTION_REASON_SUBFIELD_DROP)
		}
		itemAddHintNotify := &proto.ItemAddHintNotify{
			Reason:   uint32(hintReason),
			ItemList: make([]*proto.ItemHint, 0),
		}
		for _, changeItem := range itemList {
			itemAddHintNotify.ItemList = append(itemAddHintNotify.ItemList, &proto.ItemHint{
				ItemId: changeItem.ItemId,
				Count:  changeItem.ChangeCount,
				IsNew:  false,
			})
		}
		g.SendMsg(cmd.ItemAddHintNotify, userId, player.ClientSeq, itemAddHintNotify)
	}
	return true
}

func (g *Game) CostPlayerItem(userId uint32, itemList []*ChangeItem) bool {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return false
	}
	dbItem := player.GetDbItem()
	playerPropNotify := &proto.PlayerPropNotify{
		PropMap: make(map[uint32]*proto.PropValue),
	}
	storeItemChangeNotify := &proto.StoreItemChangeNotify{
		StoreType: proto.StoreType_STORE_PACK,
		ItemList:  make([]*proto.Item, 0),
	}
	storeItemDelNotify := &proto.StoreItemDelNotify{
		StoreType: proto.StoreType_STORE_PACK,
		GuidList:  make([]uint64, 0),
	}
	for _, changeItem := range itemList {
		// 检查剩余道具数量
		count := g.GetPlayerItemCount(player.PlayerId, changeItem.ItemId)
		if count < changeItem.ChangeCount {
			return false
		}
		prop, exist := constant.VIRTUAL_ITEM_PROP[changeItem.ItemId]
		if exist {
			// 物品为虚拟物品 角色属性物品数量减少
			player.PropertiesMap[prop] -= changeItem.ChangeCount
			playerPropNotify.PropMap[uint32(prop)] = &proto.PropValue{
				Type: uint32(prop),
				Val:  int64(player.PropertiesMap[prop]),
				Value: &proto.PropValue_Ival{
					Ival: int64(player.PropertiesMap[prop]),
				},
			}
			// 特殊属性变化处理函数
			switch changeItem.ItemId {
			case constant.ITEM_ID_PLAYER_EXP:
				// 冒险阅历应该也没人会去扣吧?
				g.HandlePlayerExpAdd(userId)
			}
		} else {
			// 物品为普通物品 直接扣除
			dbItem.CostItem(player, changeItem.ItemId, changeItem.ChangeCount)
		}
		count = g.GetPlayerItemCount(player.PlayerId, changeItem.ItemId)
		if count > 0 {
			pbItem := &proto.Item{
				ItemId: changeItem.ItemId,
				Guid:   dbItem.GetItemGuid(changeItem.ItemId),
				Detail: &proto.Item_Material{
					Material: &proto.Material{
						Count: count,
					},
				},
			}
			storeItemChangeNotify.ItemList = append(storeItemChangeNotify.ItemList, pbItem)
		} else if count == 0 {
			storeItemDelNotify.GuidList = append(storeItemDelNotify.GuidList, dbItem.GetItemGuid(changeItem.ItemId))
		}
	}

	if len(playerPropNotify.PropMap) > 0 {
		g.SendMsg(cmd.PlayerPropNotify, userId, player.ClientSeq, playerPropNotify)
	}
	if len(storeItemChangeNotify.ItemList) > 0 {
		g.SendMsg(cmd.StoreItemChangeNotify, userId, player.ClientSeq, storeItemChangeNotify)
	}
	if len(storeItemDelNotify.GuidList) > 0 {
		g.SendMsg(cmd.StoreItemDelNotify, userId, player.ClientSeq, storeItemDelNotify)
	}

	return true
}

/************************************************** 打包封装 **************************************************/

func (g *Game) PacketStoreWeightLimitNotify() *proto.StoreWeightLimitNotify {
	storeWeightLimitNotify := &proto.StoreWeightLimitNotify{
		StoreType: proto.StoreType_STORE_PACK,
		// 背包容量限制
		WeightLimit:         constant.STORE_PACK_LIMIT_WEIGHT,
		WeaponCountLimit:    constant.STORE_PACK_LIMIT_WEAPON,
		ReliquaryCountLimit: constant.STORE_PACK_LIMIT_RELIQUARY,
		MaterialCountLimit:  constant.STORE_PACK_LIMIT_MATERIAL,
		FurnitureCountLimit: constant.STORE_PACK_LIMIT_FURNITURE,
	}
	return storeWeightLimitNotify
}

func (g *Game) PacketPlayerStoreNotify(player *model.Player) *proto.PlayerStoreNotify {
	dbItem := player.GetDbItem()
	dbWeapon := player.GetDbWeapon()
	dbReliquary := player.GetDbReliquary()
	playerStoreNotify := &proto.PlayerStoreNotify{
		StoreType:   proto.StoreType_STORE_PACK,
		WeightLimit: constant.STORE_PACK_LIMIT_WEIGHT,
		ItemList:    make([]*proto.Item, 0, len(dbItem.ItemMap)+len(dbWeapon.WeaponMap)+len(dbReliquary.ReliquaryMap)),
	}
	for _, weapon := range dbWeapon.WeaponMap {
		itemDataConfig := gdconf.GetItemDataById(int32(weapon.ItemId))
		if itemDataConfig == nil {
			logger.Error("get item data config is nil, itemId: %v", weapon.ItemId)
			continue
		}
		if itemDataConfig.Type != constant.ITEM_TYPE_WEAPON {
			continue
		}
		affixMap := make(map[uint32]uint32)
		for _, affixId := range weapon.AffixIdList {
			affixMap[affixId] = uint32(weapon.Refinement)
		}
		pbItem := &proto.Item{
			ItemId: weapon.ItemId,
			Guid:   weapon.Guid,
			Detail: &proto.Item_Equip{
				Equip: &proto.Equip{
					Detail: &proto.Equip_Weapon{
						Weapon: &proto.Weapon{
							Level:        uint32(weapon.Level),
							Exp:          weapon.Exp,
							PromoteLevel: uint32(weapon.Promote),
							AffixMap:     affixMap,
						},
					},
					IsLocked: weapon.Lock,
				},
			},
		}
		playerStoreNotify.ItemList = append(playerStoreNotify.ItemList, pbItem)
	}
	for _, reliquary := range dbReliquary.ReliquaryMap {
		itemDataConfig := gdconf.GetItemDataById(int32(reliquary.ItemId))
		if itemDataConfig == nil {
			logger.Error("get item data config is nil, itemId: %v", reliquary.ItemId)
			continue
		}
		if itemDataConfig.Type != constant.ITEM_TYPE_RELIQUARY {
			continue
		}
		pbItem := &proto.Item{
			ItemId: reliquary.ItemId,
			Guid:   reliquary.Guid,
			Detail: &proto.Item_Equip{
				Equip: &proto.Equip{
					Detail: &proto.Equip_Reliquary{
						Reliquary: &proto.Reliquary{
							Level:            uint32(reliquary.Level),
							Exp:              reliquary.Exp,
							PromoteLevel:     uint32(reliquary.Promote),
							MainPropId:       reliquary.MainPropId,
							AppendPropIdList: reliquary.AppendPropIdList,
						},
					},
					IsLocked: reliquary.Lock,
				},
			},
		}
		playerStoreNotify.ItemList = append(playerStoreNotify.ItemList, pbItem)
	}
	for _, item := range dbItem.ItemMap {
		itemDataConfig := gdconf.GetItemDataById(int32(item.ItemId))
		if itemDataConfig == nil {
			logger.Error("get item data config is nil, itemId: %v", item.ItemId)
			continue
		}
		pbItem := &proto.Item{
			ItemId: item.ItemId,
			Guid:   item.Guid,
			Detail: nil,
		}
		if itemDataConfig != nil && itemDataConfig.Type == constant.ITEM_TYPE_FURNITURE {
			pbItem.Detail = &proto.Item_Furniture{
				Furniture: &proto.Furniture{
					Count: item.Count,
				},
			}
		} else {
			pbItem.Detail = &proto.Item_Material{
				Material: &proto.Material{
					Count:      item.Count,
					DeleteInfo: nil,
				},
			}
		}
		playerStoreNotify.ItemList = append(playerStoreNotify.ItemList, pbItem)
	}
	return playerStoreNotify
}
