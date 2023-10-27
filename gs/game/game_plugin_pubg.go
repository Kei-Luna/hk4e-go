package game

import (
	"encoding/base64"
	"math"
	"time"

	"hk4e/common/constant"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/pkg/random"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

const (
	PUBG_STATE_WAIT = iota
	PUBG_STATE_START
	PUBG_STATE_STOP
)

const (
	PUBG_START_MIN_PLAYER_COUNT = 2  // pubg游戏开始需要的最少玩家数
	PUBG_WAIT_TICK              = 30 // pubg游戏玩家人数足够后需等待的时间
	PUBG_STOP_TICK              = 5  // pubg游戏结束后后需等待的时间
)

const (
	PUBG_PHASE_START = 0
	PUBG_PHASE_II    = 2
	PUBG_PHASE_END   = 16
)

const (
	PUBG_PHASE_INV_TIME         = 180.0
	PUBG_FIRST_AREA_REDUCE_TIME = 300.0
)

const (
	PUBG_ATK                         = 100.0
	PUBG_HP                          = 1000.0
	PUBG_HP_LOST                     = 10.0
	PUBG_BOW_ATTACK_ATK_RATIO        = 2.0
	PUBG_NORMAL_ATTACK_DISTANCE      = 3.0
	PUBG_NORMAL_ATTACK_INTERVAL_TIME = 500
	PUBG_NORMAL_ATTACK_ATK_RATIO     = 5.0
)

// PubgGame pubg游戏
type PubgGame struct {
	gameTick                 uint64                // 游戏tick
	gameState                uint8                 // 游戏状态
	waitTick                 uint64                // 等待tick
	stopTick                 uint64                // 等待tick
	phase                    uint16                // 阶段
	ai                       *model.Player         // ai
	world                    *World                // 世界对象
	gameId                   uint32                // 匹配服游戏id
	roomId                   uint32                // 匹配服房间id
	playerUidList            []uint32              // 玩家uid列表
	blueAreaCenterPos        *model.Vector         // 蓝区中心点
	blueAreaRadius           float64               // 蓝区半径
	safeAreaCenterPos        *model.Vector         // 安全区中心点
	safeAreaRadius           float64               // 安全区半径
	areaReduceRadiusSpeed    float64               // 缩圈半径速度
	areaReduceXSpeed         float64               // 缩圈X速度
	areaReduceZSpeed         float64               // 缩圈Z速度
	areaPointList            []*proto.MapMarkPoint // 客户端区域地图坐标列表
	entityIdWorldGadgetIdMap map[uint32]int32      // 实体id世界物件id映射集合
	playerHitTimeMap         map[uint32]int64      // 玩家攻击命中时间集合

	pluginPubg *PluginPubg
}

// PluginPubg pubg游戏插件
type PluginPubg struct {
	*Plugin
	pubgGameMap map[uint32]*PubgGame // pubg游戏集合
}

func NewPluginPubg() *PluginPubg {
	p := &PluginPubg{
		Plugin:      NewPlugin(),
		pubgGameMap: make(map[uint32]*PubgGame),
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
	p.ListenEvent(PluginEventIdMatchCreateAi, PluginEventPriorityNormal, p.EventMatchCreateAi)
	p.ListenEvent(PluginEventIdJoinOtherWorld, PluginEventPriorityNormal, p.EventJoinOtherWorld)
	p.ListenEvent(PluginEventIdEvtDoSkillSucc, PluginEventPriorityNormal, p.EventEvtDoSkillSucc)
	p.ListenEvent(PluginEventIdBackMyWorld, PluginEventPriorityNormal, p.EventBackMyWorld)
	p.ListenEvent(PluginEventIdUserOffline, PluginEventPriorityNormal, p.EventUserOffline)
	// 添加全局定时器
	p.AddGlobalTick(PluginGlobalTickSecond, p.GlobalTickPubg)
	p.AddGlobalTick(PluginGlobalTick100MilliSecond, p.GlobalTickPubgHit)
	// 注册命令
	// p.RegCommandController(p.NewPubgCommandController())
}

/************************************************** 事件监听 **************************************************/

// EventMatchCreateAi 游戏房间ai创建
func (p *PluginPubg) EventMatchCreateAi(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventMatchCreateAi)
	// 判断是否为pubg游戏类型
	if event.GameType != constant.MATCH_GAME_TYPE_PUBG {
		return
	}
	// 创建pubg游戏
	p.CreatePubgGame(event.Ai, event.GameId, event.RoomId, p)
}

// EventJoinOtherWorld 玩家加入游戏世界
func (p *PluginPubg) EventJoinOtherWorld(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventJoinOtherWorld)
	player := event.Player
	// 通过ai获取pubg游戏
	pubgGame, exist := p.pubgGameMap[event.HostPlayer.PlayerId]
	if !exist {
		return
	}
	p.PlayerJoinGame(player, pubgGame)
}

// EventBackMyWorld 玩家返回自己的世界
func (p *PluginPubg) EventBackMyWorld(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventBackMyWorld)
	player := event.Player
	// 通过ai获取pubg游戏
	pubgGame, exist := p.pubgGameMap[event.HostWorld.GetOwner().PlayerId]
	if !exist {
		return
	}
	p.PlayerExitGame(player, pubgGame)
}

// EventUserOffline 玩家离线
func (p *PluginPubg) EventUserOffline(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventUserOffline)
	player := event.Player
	// 获取玩家所在的世界
	world := WORLD_MANAGER.GetWorldById(player.WorldId)
	if world == nil {
		return
	}
	// 通过ai获取pubg游戏
	pubgGame, exist := p.pubgGameMap[world.GetOwner().PlayerId]
	if !exist {
		return
	}
	p.PlayerExitGame(player, pubgGame)
}

// EventEvtDoSkillSucc 使用技能完毕
func (p *PluginPubg) EventEvtDoSkillSucc(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventEvtDoSkillSucc)
	player := event.Player
	// 获取玩家所在的pubg游戏
	pubgGame := p.GetPlayerInPubgGame(player)
	if pubgGame == nil {
		return
	}
	// 确保游戏开启
	if !pubgGame.IsStartPubg() || pubgGame.world.GetId() != player.WorldId {
		return
	}
	worldAvatar := pubgGame.world.GetWorldAvatarByEntityId(event.Ntf.CasterId)
	if worldAvatar == nil {
		return
	}
	avatarDataConfig := gdconf.GetAvatarDataById(int32(worldAvatar.GetAvatarId()))
	if avatarDataConfig == nil {
		return
	}
	logger.Debug("avatar normal attack, avatarId: %v, weaponType: %v, uid: %v", avatarDataConfig.AvatarId, avatarDataConfig.WeaponType, player.PlayerId)
	switch avatarDataConfig.WeaponType {
	case constant.WEAPON_TYPE_SWORD_ONE_HAND, constant.WEAPON_TYPE_CLAYMORE, constant.WEAPON_TYPE_POLE, constant.WEAPON_TYPE_CATALYST, constant.WEAPON_TYPE_BOW:
		scene := pubgGame.world.GetSceneById(player.SceneId)
		avatarEntity := scene.GetEntity(worldAvatar.GetAvatarEntityId())
		for _, entity := range scene.GetAllEntity() {
			if entity.GetId() == avatarEntity.GetId() || entity.GetEntityType() != constant.ENTITY_TYPE_AVATAR {
				continue
			}
			distance3D := math.Sqrt(
				(avatarEntity.GetPos().X-entity.GetPos().X)*(avatarEntity.GetPos().X-entity.GetPos().X) +
					(avatarEntity.GetPos().Y-entity.GetPos().Y)*(avatarEntity.GetPos().Y-entity.GetPos().Y) +
					(avatarEntity.GetPos().Z-entity.GetPos().Z)*(avatarEntity.GetPos().Z-entity.GetPos().Z),
			)
			if distance3D > PUBG_NORMAL_ATTACK_DISTANCE {
				continue
			}
			pubgGame.PubgHit(scene, entity.GetId(), avatarEntity.GetId(), false)
		}
	default:
	}
}

// EventKillAvatar 角色被杀死事件
func (p *PluginPubg) EventKillAvatar(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventKillAvatar)
	player := event.Player
	// 获取玩家所在的pubg游戏
	pubgGame := p.GetPlayerInPubgGame(player)
	if pubgGame == nil {
		return
	}
	// 确保游戏开启
	if !pubgGame.IsStartPubg() || pubgGame.world.GetId() != player.WorldId {
		return
	}
	p.CreateUserTimer(player.PlayerId, 10, pubgGame.UserTimerPubgDieExit, pubgGame.world.GetId())
}

// EventMarkMap 地图标点事件
func (p *PluginPubg) EventMarkMap(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventMarkMap)
	player := event.Player
	// 获取玩家所在的pubg游戏
	pubgGame := p.GetPlayerInPubgGame(player)
	if pubgGame == nil {
		return
	}
	// 确保游戏开启
	if !pubgGame.IsStartPubg() || pubgGame.world.GetId() != player.WorldId {
		return
	}
	GAME.SendMsg(cmd.MarkMapRsp, player.PlayerId, player.ClientSeq, &proto.MarkMapRsp{MarkList: pubgGame.GetAreaPointList()})
	event.Cancel()
}

// EventAvatarDieAnimationEnd 角色死亡动画结束事件
func (p *PluginPubg) EventAvatarDieAnimationEnd(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventAvatarDieAnimationEnd)
	player := event.Player
	// 获取玩家所在的pubg游戏
	pubgGame := p.GetPlayerInPubgGame(player)
	if pubgGame == nil {
		return
	}
	// 确保游戏开启
	if !pubgGame.IsStartPubg() || pubgGame.world.GetId() != player.WorldId {
		return
	}
	alivePlayerNum := len(pubgGame.GetAlivePlayerList())
	GAME.SendWorldChat(pubgGame.world, "『%v』死亡了，剩余%v位存活玩家。", player.NickName, alivePlayerNum)
	event.Cancel()
}

func (p *PluginPubg) EventGadgetInteract(iEvent IPluginEvent) {
	event := iEvent.(*PluginEventGadgetInteract)
	player := event.Player
	// 获取玩家所在的pubg游戏
	pubgGame := p.GetPlayerInPubgGame(player)
	if pubgGame == nil {
		return
	}
	// 确保游戏开启
	if !pubgGame.IsStartPubg() || pubgGame.world.GetId() != player.WorldId {
		return
	}
	req := event.Req
	worldGadgetId, exist := pubgGame.entityIdWorldGadgetIdMap[req.GadgetEntityId]
	if exist {
		dbAvatar := player.GetDbAvatar()
		avatarId := pubgGame.world.GetPlayerActiveAvatarId(player)
		avatar := dbAvatar.AvatarMap[avatarId]
		pubgWorldGadgetDataConfig := gdconf.GetPubgWorldGadgetDataById(worldGadgetId)
		switch pubgWorldGadgetDataConfig.Type {
		case gdconf.PubgWorldGadgetTypeIncAtk:
			avatar.FightPropMap[constant.FIGHT_PROP_BASE_ATTACK] += float32(pubgWorldGadgetDataConfig.Param[0])
			avatar.FightPropMap[constant.FIGHT_PROP_CUR_ATTACK] += float32(pubgWorldGadgetDataConfig.Param[0])
			// 提示玩家
			GAME.SendWorldChat(pubgGame.world, "你的角色攻击力增加：%v，增加后的攻击力：%v。", pubgWorldGadgetDataConfig.Param[0], avatar.FightPropMap[constant.FIGHT_PROP_BASE_ATTACK])
		case gdconf.PubgWorldGadgetTypeIncHp:
			avatar.FightPropMap[constant.FIGHT_PROP_CUR_HP] += float32(pubgWorldGadgetDataConfig.Param[0])
			if avatar.FightPropMap[constant.FIGHT_PROP_CUR_HP] > avatar.FightPropMap[constant.FIGHT_PROP_MAX_HP] {
				avatar.FightPropMap[constant.FIGHT_PROP_CUR_HP] = avatar.FightPropMap[constant.FIGHT_PROP_MAX_HP]
			}
			// 提示玩家
			GAME.SendWorldChat(pubgGame.world, "你的角色生命值增加：%v，目前为：%v。", pubgWorldGadgetDataConfig.Param[0], avatar.FightPropMap[constant.FIGHT_PROP_CUR_HP])
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
	// 获取玩家所在的pubg游戏
	pubgGame := p.GetPlayerInPubgGame(player)
	if pubgGame == nil {
		return
	}
	// 确保游戏开启
	if !pubgGame.IsStartPubg() || pubgGame.world.GetId() != player.WorldId {
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
func (p *PluginPubg) GlobalTickPubg(now int64) {
	for _, game := range p.pubgGameMap {
		// 游戏没有玩家则不进行tick
		if len(game.playerUidList) == 0 {
			continue
		}
		game.gameTick++
		// 处理当前阶段
		switch game.gameState {
		case PUBG_STATE_WAIT:
			// 等待阶段
			game.PhaseWait()
		case PUBG_STATE_START:
			// 开始阶段
			game.HandleStateStart()
		case PUBG_STATE_STOP:
			// 结束阶段
			game.HandleStateStop()
		}
	}
}

// GlobalTickPubgHit pubg命中
func (p *PluginPubg) GlobalTickPubgHit(now int64) {
	for _, game := range p.pubgGameMap {
		if !game.IsStartPubg() {
			return
		}
		bulletPhysicsEngine := game.world.GetBulletPhysicsEngine()
		hitList := bulletPhysicsEngine.Update(now)
		for _, rigidBody := range hitList {
			scene := game.world.GetSceneById(rigidBody.sceneId)
			game.PubgHit(scene, rigidBody.hitAvatarEntityId, rigidBody.avatarEntityId, true)
		}
	}
}

/************************************************** 用户定时器 **************************************************/

// UserTimerPubgUpdateArea 更新游戏区域
func (p *PubgGame) UserTimerPubgUpdateArea(player *model.Player, data []any) {
	logger.Debug("UserTimerPubgUpdateArea, gameId: %v, roomId: %v", p.gameId, p.roomId)
	if !p.IsStartPubg() {
		return
	}
	p.phase++
	p.RefreshArea()
}

// UserTimerPubgDieExit pubg死亡离开
func (p *PubgGame) UserTimerPubgDieExit(player *model.Player, data []any) {
	logger.Debug("UserTimerPubgDieExit, gameId: %v, roomId: %v", p.gameId, p.roomId)
	pubgWorldId := data[0].(uint64)
	if player.WorldId != pubgWorldId {
		return
	}
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
		Perm: CommandPermGM,
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
			// p.StartPubg()
			c.SendSuccMessage(c.Executor, "已开始PUBG游戏。")
		case "stop":
			// 结束游戏
			// p.StopPubg()
			c.SendSuccMessage(c.Executor, "已结束PUBG游戏。")
		default:
			return false
		}
		return true
	})
}

/************************************************** 插件功能 **************************************************/

// CreatePubgGame 创建pubg游戏
func (p *PluginPubg) CreatePubgGame(ai *model.Player, gameId uint32, roomId uint32, pluginPubg *PluginPubg) {
	world := WORLD_MANAGER.GetWorldById(ai.WorldId)
	if world == nil {
		logger.Error("world is nil, worldId: %v, uid: %v", ai.WorldId, ai.PlayerId)
		return
	}
	p.pubgGameMap[ai.PlayerId] = &PubgGame{
		gameTick:                 0,
		gameState:                PUBG_STATE_WAIT,
		waitTick:                 PUBG_WAIT_TICK,
		stopTick:                 PUBG_STOP_TICK,
		phase:                    PUBG_PHASE_START,
		ai:                       ai,
		world:                    world,
		gameId:                   gameId,
		roomId:                   roomId,
		playerUidList:            make([]uint32, 0),
		blueAreaCenterPos:        &model.Vector{X: 0.0, Y: 0.0, Z: 0.0},
		blueAreaRadius:           0.0,
		safeAreaCenterPos:        &model.Vector{X: 0.0, Y: 0.0, Z: 0.0},
		safeAreaRadius:           0.0,
		areaReduceRadiusSpeed:    0.0,
		areaReduceXSpeed:         0.0,
		areaReduceZSpeed:         0.0,
		areaPointList:            make([]*proto.MapMarkPoint, 0),
		entityIdWorldGadgetIdMap: make(map[uint32]int32),
		playerHitTimeMap:         make(map[uint32]int64),
		pluginPubg:               pluginPubg,
	}
}

// DeletePubgGame 删除pubg游戏
func (p *PluginPubg) DeletePubgGame(game *PubgGame) {
	_, exist := p.pubgGameMap[game.ai.PlayerId]
	if !exist {
		logger.Error("pubg game not exist, gameId: %v, roomId: %v, aiUid: %v", game.gameId, game.roomId, game.ai.PlayerId)
		return
	}
	delete(p.pubgGameMap, game.ai.PlayerId)
}

// PlayerJoinGame 玩家加入游戏
func (p *PluginPubg) PlayerJoinGame(player *model.Player, game *PubgGame) {
	// 校验是否已加入
	for _, uid := range game.playerUidList {
		if uid == player.PlayerId {
			logger.Error("player has been join, uid: %v", player.PlayerId)
			return
		}
	}
	game.playerUidList = append(game.playerUidList, player.PlayerId)
}

// PlayerExitGame 玩家离开游戏
func (p *PluginPubg) PlayerExitGame(player *model.Player, game *PubgGame) {
	// 清除uid
	for i, uid := range game.playerUidList {
		if uid == player.PlayerId {
			game.playerUidList = append(game.playerUidList[:i], game.playerUidList[i+1:]...)
		}
	}
}

// GetPlayerInPubgGame 获取玩家所在的pubg游戏
func (p *PluginPubg) GetPlayerInPubgGame(player *model.Player) *PubgGame {
	for _, game := range p.pubgGameMap {
		for _, uid := range game.playerUidList {
			if uid == player.PlayerId {
				return game
			}
		}
	}
	return nil
}

// PhaseWait 等待阶段处理
func (p *PubgGame) PhaseWait() {
	// 玩家数是否到达设定值 到达则开始倒计时 没有则提示玩家
	if len(p.playerUidList) >= PUBG_START_MIN_PLAYER_COUNT {
		switch p.waitTick {
		case PUBG_WAIT_TICK:
			GAME.SendWorldChat(p.world, "玩家数量足够，开始倒计时！\n感谢各位玩家参与PUBG游戏～")
		case 20, 10, 5, 4, 3, 2, 1:
			GAME.SendWorldChat(p.world, "游戏将在%v秒后开始！", p.waitTick)
		case 0:
			// 倒计时到了则开始游戏
			p.StartPubg()
			return
		}
		p.waitTick-- // 计时-1

		// 如果人少于设定值并正在倒计时的话那就是取消了
	} else if p.waitTick != PUBG_WAIT_TICK {
		GAME.SendWorldChat(p.world, "玩家数量不足！暂停游戏开始倒计时。")
		// 重置计时器
		p.waitTick = PUBG_WAIT_TICK

		// 玩家不足每5s提示一次
	} else if p.gameTick%5 == 0 {
		GAME.SendWorldChat(p.world, "玩家数量不足，还需要%v名玩家加入。", PUBG_START_MIN_PLAYER_COUNT-len(p.playerUidList))
	}
}

// HandleStateStart 开始状态处理
func (p *PubgGame) HandleStateStart() {
	// 更新区域
	p.UpdateArea()
	scene := p.world.GetSceneById(p.world.GetOwner().SceneId)
	for _, scenePlayer := range scene.GetAllPlayer() {
		if !p.IsInBlueArea(scenePlayer.Pos) {
			GAME.handleEvtBeingHit(scenePlayer, scene, &proto.EvtBeingHitInfo{
				AttackResult: &proto.AttackResult{
					AttackerId: 0,
					DefenseId:  p.world.GetPlayerWorldAvatarEntityId(scenePlayer, p.world.GetPlayerActiveAvatarId(scenePlayer)),
					Damage:     PUBG_HP_LOST,
				},
			})
		}
	}
	// 判断玩家存活状况
	alivePlayerList := p.GetAlivePlayerList()
	if len(alivePlayerList) <= 1 {
		if len(alivePlayerList) == 1 {
			GAME.SendWorldChat(p.world, "『%v』大吉大利，今晚吃鸡。", alivePlayerList[0].NickName)
		}
		p.StopPubg()
	}
}

// HandleStateStop 结束状态处理
func (p *PubgGame) HandleStateStop() {
	if p.stopTick <= 0 {
		GAME.SendWorldChat(p.world, "游戏结束。")
		// 清除游戏
		p.pluginPubg.DeletePubgGame(p)
		// 通知匹配服游戏结束
		p.pluginPubg.MatchGameStop(p.ai, p.gameId, p.roomId)
		return
	}
	p.stopTick--
}

// WaitPubg 等待pubg游戏开始
func (p *PubgGame) WaitPubg() {
	logger.Debug("WaitPubg, gameId: %v, roomId: %v", p.gameId, p.roomId)
	// 确保上个阶段为结束状态
	if p.gameState != PUBG_STATE_STOP {
		logger.Error("pubg game state not end, gameState: %v", p.gameState)
		return
	}
	// 切换阶段
	p.gameState = PUBG_STATE_WAIT
}

// StartPubg 开始pubg游戏
func (p *PubgGame) StartPubg() {
	logger.Debug("StartPubg, gameId: %v, roomId: %v", p.gameId, p.roomId)
	// 确保上个阶段为等待状态
	if p.gameState != PUBG_STATE_WAIT {
		logger.Error("pubg game state not wait, gameState: %v", p.gameState)
		return
	}
	// 切换阶段
	p.gameState = PUBG_STATE_START
	// 重置game tick
	p.gameTick = 0
	// 初始化pubg游戏
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
	// 刷新区域
	p.RefreshArea()
	for _, player := range p.world.GetAllPlayer() {
		// 跳过ai
		if player.PlayerId == p.ai.PlayerId {
			continue
		}
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
		p.playerHitTimeMap[player.PlayerId] = 0
		// 传送玩家至ai身边
		GAME.TeleportPlayer(
			player,
			proto.EnterReason_ENTER_REASON_GM,
			p.ai.SceneId,
			&model.Vector{X: p.ai.Pos.X, Y: p.ai.Pos.Y, Z: p.ai.Pos.Z},
			new(model.Vector),
			0,
			0,
		)
	}
	GAME.SendWorldChat(p.world, "游戏开始。祝您游戏愉快～")
	// 通知匹配服游戏开始
	p.pluginPubg.MatchGameStart(p.ai, p.gameId, p.roomId)
}

// StopPubg 结束pubg游戏
func (p *PubgGame) StopPubg() {
	logger.Debug("StopPubg, gameId: %v, roomId: %v", p.gameId, p.roomId)
	// 确保上个阶段为不为等待状态和结束状态
	if !p.IsStartPubg() {
		logger.Error("pubg game state not end, gameState: %v", p.gameState)
		return
	}
	// 切换阶段
	p.gameState = PUBG_STATE_STOP
}

// IsStartPubg pubg游戏是否开启
func (p *PubgGame) IsStartPubg() bool {
	return p.gameState != PUBG_STATE_WAIT && p.gameState != PUBG_STATE_STOP
}

// GetAreaPointList 获取游戏区域标点列表
func (p *PubgGame) GetAreaPointList() []*proto.MapMarkPoint {
	return p.areaPointList
}

// UpdateArea 更新游戏区域
func (p *PubgGame) UpdateArea() {
	if p.areaReduceRadiusSpeed > 0.0 && p.blueAreaRadius > p.safeAreaRadius {
		p.blueAreaRadius -= p.areaReduceRadiusSpeed
		p.blueAreaCenterPos.X += p.areaReduceXSpeed
		p.blueAreaCenterPos.Z += p.areaReduceZSpeed
		p.SyncMapMarkArea()
	}
}

// IsInBlueArea 是否在蓝圈内
func (p *PubgGame) IsInBlueArea(pos *model.Vector) bool {
	distance2D := math.Sqrt(
		(p.blueAreaCenterPos.X-pos.X)*(p.blueAreaCenterPos.X-pos.X) +
			(p.blueAreaCenterPos.Z-pos.Z)*(p.blueAreaCenterPos.Z-pos.Z),
	)
	return distance2D < p.blueAreaRadius
}

// RefreshArea 刷新游戏区域
func (p *PubgGame) RefreshArea() {
	if p.phase == PUBG_PHASE_START {
		GAME.SendWorldChat(p.world, "安全区已生成，当前%v位存活玩家。", len(p.GetAlivePlayerList()))
		p.blueAreaCenterPos = &model.Vector{X: 500.0, Y: 0.0, Z: -500.0}
		p.blueAreaRadius = 2000.0
		p.safeAreaCenterPos = &model.Vector{X: 0.0, Y: 0.0, Z: 0.0}
		p.safeAreaRadius = 0.0
		p.pluginPubg.CreateUserTimer(p.world.GetOwner().PlayerId, PUBG_PHASE_INV_TIME, p.UserTimerPubgUpdateArea)
	} else if p.phase == PUBG_PHASE_END {
		GAME.SendWorldChat(p.world, "安全区已消失。")
		p.blueAreaRadius = 0.0
		p.safeAreaRadius = 0.0
	} else {
		switch p.phase % 3 {
		case 1:
			GAME.SendWorldChat(p.world, "新的安全区已出现，进度%.1f%%。", float64(p.phase)/PUBG_PHASE_END*100.0)
			p.safeAreaCenterPos = &model.Vector{
				X: p.blueAreaCenterPos.X + random.GetRandomFloat64(-(p.blueAreaRadius*0.7/2.0), p.blueAreaRadius*0.7/2.0),
				Y: 0.0,
				Z: p.blueAreaCenterPos.Z + random.GetRandomFloat64(-(p.blueAreaRadius*0.7/2.0), p.blueAreaRadius*0.7/2.0),
			}
			p.safeAreaRadius = p.blueAreaRadius / 2.0
			p.areaReduceRadiusSpeed = 0.0
			p.pluginPubg.CreateUserTimer(p.world.GetOwner().PlayerId, PUBG_PHASE_INV_TIME, p.UserTimerPubgUpdateArea)
		case 2:
			GAME.SendWorldChat(p.world, "安全区正在缩小，进度%.1f%%。", float64(p.phase)/PUBG_PHASE_END*100.0)
			invTime := 0.0
			if p.phase == PUBG_PHASE_II {
				invTime = PUBG_FIRST_AREA_REDUCE_TIME
			} else {
				invTime = PUBG_PHASE_INV_TIME
			}
			p.areaReduceRadiusSpeed = (p.blueAreaRadius - p.safeAreaRadius) / invTime
			p.areaReduceXSpeed = (p.safeAreaCenterPos.X - p.blueAreaCenterPos.X) / invTime
			p.areaReduceZSpeed = (p.safeAreaCenterPos.Z - p.blueAreaCenterPos.Z) / invTime
			p.pluginPubg.CreateUserTimer(p.world.GetOwner().PlayerId, uint32(invTime), p.UserTimerPubgUpdateArea)
		case 0:
			GAME.SendWorldChat(p.world, "安全区缩小完毕，进度%.1f%%。", float64(p.phase)/PUBG_PHASE_END*100.0)
			p.pluginPubg.CreateUserTimer(p.world.GetOwner().PlayerId, PUBG_PHASE_INV_TIME, p.UserTimerPubgUpdateArea)
		}
	}
	p.SyncMapMarkArea()
}

// SyncMapMarkArea 同步地图标点区域
func (p *PubgGame) SyncMapMarkArea() {
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
func (p *PubgGame) GetAlivePlayerList() []*model.Player {
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

func (p *PubgGame) PubgHit(scene *Scene, defAvatarEntityId uint32, atkAvatarEntityId uint32, isBow bool) {
	defAvatarEntity := scene.GetEntity(defAvatarEntityId)
	if defAvatarEntity == nil {
		return
	}
	defPlayer := USER_MANAGER.GetOnlineUser(defAvatarEntity.GetAvatarEntity().GetUid())
	if defPlayer == nil {
		return
	}
	atkAvatarEntity := scene.GetEntity(atkAvatarEntityId)
	if atkAvatarEntity == nil {
		return
	}
	atkPlayer := USER_MANAGER.GetOnlineUser(atkAvatarEntity.GetAvatarEntity().GetUid())
	if atkPlayer == nil {
		return
	}
	now := time.Now().UnixMilli()
	lastHitTime := p.playerHitTimeMap[atkPlayer.PlayerId]
	if now-lastHitTime < PUBG_NORMAL_ATTACK_INTERVAL_TIME {
		return
	}
	p.playerHitTimeMap[atkPlayer.PlayerId] = now
	atk := atkAvatarEntity.GetFightProp()[constant.FIGHT_PROP_CUR_ATTACK]
	dmg := float32(0.0)
	if isBow {
		dmg = atk / PUBG_BOW_ATTACK_ATK_RATIO
	} else {
		dmg = atk / PUBG_NORMAL_ATTACK_ATK_RATIO
	}
	GAME.handleEvtBeingHit(defPlayer, scene, &proto.EvtBeingHitInfo{
		AttackResult: &proto.AttackResult{
			AttackerId: atkAvatarEntity.GetId(),
			DefenseId:  defAvatarEntity.GetId(),
			Damage:     dmg,
		},
	})
	if attackResultTemplate == nil {
		return
	}
	evtBeingHitInfo := &proto.EvtBeingHitInfo{
		PeerId:       0,
		AttackResult: attackResultTemplate,
		FrameNum:     0,
	}
	evtBeingHitInfo.AttackResult.AttackerId = atkAvatarEntity.GetId()
	evtBeingHitInfo.AttackResult.DefenseId = defAvatarEntity.GetId()
	evtBeingHitInfo.AttackResult.Damage = dmg
	if evtBeingHitInfo.AttackResult.HitCollision == nil {
		return
	}
	evtBeingHitInfo.AttackResult.HitCollision.HitPoint = &proto.Vector{X: float32(defPlayer.Pos.X), Y: float32(defPlayer.Pos.Y), Z: float32(defPlayer.Pos.Z)}
	combatData, err := pb.Marshal(evtBeingHitInfo)
	if err != nil {
		return
	}
	GAME.SendToSceneA(scene, cmd.CombatInvocationsNotify, 0, &proto.CombatInvocationsNotify{
		InvokeList: []*proto.CombatInvokeEntry{{
			CombatData:   combatData,
			ForwardType:  proto.ForwardType_FORWARD_TO_ALL,
			ArgumentType: proto.CombatTypeArgument_COMBAT_EVT_BEING_HIT,
		}},
	}, 0)
}
