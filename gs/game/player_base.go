package game

import (
	"time"

	"hk4e/protocol/cmd"

	"hk4e/common/constant"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/protocol/proto"
)

/************************************************** 接口请求 **************************************************/

/************************************************** 游戏功能 **************************************************/

// HandlePlayerExpAdd 玩家冒险阅历增加处理
func (g *Game) HandlePlayerExpAdd(userId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	// 玩家升级
	for {
		playerLevel := player.PropertiesMap[constant.PLAYER_PROP_PLAYER_LEVEL]
		// 读取玩家等级配置表
		playerLevelConfig := gdconf.GetPlayerLevelDataByLevel(int32(playerLevel))
		if playerLevelConfig == nil {
			// 获取不到代表已经到达最大等级
			break
		}
		// 确保拥有下一级的配置表
		if gdconf.GetPlayerLevelDataByLevel(int32(playerLevel+1)) == nil {
			// 获取不到代表已经到达最大等级
			break
		}
		// 玩家冒险阅历不足则跳出循环
		if player.PropertiesMap[constant.PLAYER_PROP_PLAYER_EXP] < uint32(playerLevelConfig.Exp) {
			break
		}
		// 玩家增加冒险等阶
		player.PropertiesMap[constant.PLAYER_PROP_PLAYER_LEVEL]++
		player.PropertiesMap[constant.PLAYER_PROP_PLAYER_EXP] -= uint32(playerLevelConfig.Exp)
	}
	// 更新玩家属性
	playerPropNotify := &proto.PlayerPropNotify{
		PropMap: make(map[uint32]*proto.PropValue),
	}
	playerPropNotify.PropMap[uint32(constant.PLAYER_PROP_PLAYER_LEVEL)] = &proto.PropValue{
		Type: uint32(constant.PLAYER_PROP_PLAYER_LEVEL),
		Val:  int64(player.PropertiesMap[constant.PLAYER_PROP_PLAYER_LEVEL]),
		Value: &proto.PropValue_Ival{
			Ival: int64(player.PropertiesMap[constant.PLAYER_PROP_PLAYER_LEVEL]),
		},
	}
	playerPropNotify.PropMap[uint32(constant.PLAYER_PROP_PLAYER_EXP)] = &proto.PropValue{
		Type: uint32(constant.PLAYER_PROP_PLAYER_EXP),
		Val:  int64(player.PropertiesMap[constant.PLAYER_PROP_PLAYER_EXP]),
		Value: &proto.PropValue_Ival{
			Ival: int64(player.PropertiesMap[constant.PLAYER_PROP_PLAYER_EXP]),
		},
	}
	g.SendMsg(cmd.PlayerPropNotify, userId, player.ClientSeq, playerPropNotify)
}

/************************************************** 打包封装 **************************************************/

func (g *Game) PacketPlayerDataNotify(player *model.Player) *proto.PlayerDataNotify {
	playerDataNotify := &proto.PlayerDataNotify{
		NickName:          player.NickName,
		ServerTime:        uint64(time.Now().UnixMilli()),
		IsFirstLoginToday: true,
		RegionId:          1,
		PropMap:           make(map[uint32]*proto.PropValue),
	}
	for k, v := range player.PropertiesMap {
		propValue := &proto.PropValue{
			Type:  uint32(k),
			Value: &proto.PropValue_Ival{Ival: int64(v)},
			Val:   int64(v),
		}
		playerDataNotify.PropMap[uint32(k)] = propValue
	}
	return playerDataNotify
}
