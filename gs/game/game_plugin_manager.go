package game

import (
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/protocol/proto"
	"sort"
)

// 游戏服务器插件管理器

// InitPlugin 初始化插件
func (p *PluginManager) InitPlugin() {
	iPluginList := []IPlugin{
		NewPluginPubg(),
	}
	p.RegAllPlugin(iPluginList...)
}

// PluginEventId 事件编号
type PluginEventId uint16

const (
	PluginEventIdNone = PluginEventId(iota)
	PluginEventIdPlayerKillAvatar
	PluginEventIdMarkMap
	PluginEventIdAvatarDieAnimationEnd
)

// PluginEventKillAvatar 角色被杀死
type PluginEventKillAvatar struct {
	*PluginEvent
	Player   *model.Player       // 玩家
	AvatarId uint32              // 角色id
	DieType  proto.PlayerDieType // 死亡类型
}

// PluginEventMarkMap 地图标点
type PluginEventMarkMap struct {
	*PluginEvent
	Player *model.Player     // 玩家
	Req    *proto.MarkMapReq // 请求
}

// PluginEventAvatarDieAnimationEnd 角色死亡动画结束
type PluginEventAvatarDieAnimationEnd struct {
	*PluginEvent
	Player *model.Player                   // 玩家
	Req    *proto.AvatarDieAnimationEndReq // 请求
}

type PluginEventFunc func(event IPluginEvent)

// IPluginEvent 插件事件接口
type IPluginEvent interface {
	Cancel()
	IsCancel() bool
}

// PluginEvent 插件事件
type PluginEvent struct {
	isCancel bool // 事件是否已取消
}

func NewPluginEvent() *PluginEvent {
	return &PluginEvent{}
}

// Cancel 取消事件
// 仍会继续传递事件 但不会在触发事件的地方继续执行
func (p *PluginEvent) Cancel() {
	p.isCancel = true
}

// IsCancel 事件是否已取消
func (p *PluginEvent) IsCancel() bool {
	return p.isCancel
}

// IPlugin 插件接口
type IPlugin interface {
	GetPlugin() *Plugin
	OnEnable()
	OnDisable()
}

// PluginEventPriority 插件事件优先级
type PluginEventPriority uint8

const (
	PluginEventPriorityLowest = PluginEventPriority(iota)
	PluginEventPriorityLow
	PluginEventPriorityNormal
	PluginEventPriorityHigh
	PluginEventPriorityHighest
)

// PluginEventInfo 插件事件信息
type PluginEventInfo struct {
	EventId   PluginEventId       // 事件id
	Priority  PluginEventPriority // 优先级
	EventFunc PluginEventFunc     // 事件执行函数
}

// PluginGlobalTick 全局tick
type PluginGlobalTick uint8

const (
	PluginGlobalTick50MilliSecond = PluginGlobalTick(iota)
	PluginGlobalTick100MilliSecond
	PluginGlobalTick200MilliSecond
	PluginGlobalTickSecond
	PluginGlobalTick5Second
	PluginGlobalTick10Second
	PluginGlobalTickMinute
	PluginGlobalTickHour
)

// PluginUserTimerFunc 用户timer处理函数
type PluginUserTimerFunc func(player *model.Player, data []any)

// Plugin 插件结构
type Plugin struct {
	PluginName string // 插件名 遵守小驼峰命名法

	isEnable      bool                                 // 是否启用
	eventMap      map[PluginEventId][]*PluginEventInfo // 事件集合
	globalTickMap map[PluginGlobalTick][]func()        // 全局tick集合
	userTimerMap  map[uint64]PluginUserTimerFunc       // 用户timer集合
}

func NewPlugin(pluginName string) *Plugin {
	return &Plugin{
		PluginName:    pluginName,
		isEnable:      true,
		eventMap:      make(map[PluginEventId][]*PluginEventInfo),
		globalTickMap: make(map[PluginGlobalTick][]func()),
		userTimerMap:  make(map[uint64]PluginUserTimerFunc),
	}
}

// GetPlugin 获取插件
func (p *Plugin) GetPlugin() *Plugin {
	return p
}

// OnEnable 插件启用时的生命周期
func (p *Plugin) OnEnable() {
	// 具体逻辑由插件来重写
}

// OnDisable 插件禁用时的生命周期
func (p *Plugin) OnDisable() {
	// 具体逻辑由插件来重写
}

// ListenEvent 监听事件
func (p *Plugin) ListenEvent(eventId PluginEventId, priority PluginEventPriority, eventFuncList ...PluginEventFunc) {
	for _, eventFunc := range eventFuncList {
		_, exist := p.eventMap[eventId]
		if !exist {
			p.eventMap[eventId] = make([]*PluginEventInfo, 0)
		}
		pluginEventInfo := &PluginEventInfo{
			EventId:   eventId,
			Priority:  priority,
			EventFunc: eventFunc,
		}
		p.eventMap[eventId] = append(p.eventMap[eventId], pluginEventInfo)
	}
}

// AddGlobalTick 添加全局tick
func (p *Plugin) AddGlobalTick(tick PluginGlobalTick, tickFuncList ...func()) {
	for _, tickFunc := range tickFuncList {
		_, exist := p.globalTickMap[tick]
		if !exist {
			p.globalTickMap[tick] = make([]func(), 0)
		}
		p.globalTickMap[tick] = append(p.globalTickMap[tick], tickFunc)
	}
}

// CreateUserTimer 创建用户timer
func (p *Plugin) CreateUserTimer(userId uint32, delay uint32, timerFunc PluginUserTimerFunc, data ...any) {
	PLUGIN_MANAGER.userTimerCounter++
	userTimerId := PLUGIN_MANAGER.userTimerCounter
	p.userTimerMap[userTimerId] = timerFunc
	// 用户timer编号插入到数据前
	data = append([]any{userTimerId}, data...)
	TICK_MANAGER.CreateUserTimer(userId, UserTimerActionPlugin, delay, data...)
}

type PluginManager struct {
	pluginMap        map[string]IPlugin // 插件集合
	userTimerCounter uint64             // 用户timer计数器
}

func NewPluginManager() *PluginManager {
	r := new(PluginManager)
	r.pluginMap = make(map[string]IPlugin)
	return r
}

// RegAllPlugin 注册全部插件
func (p *PluginManager) RegAllPlugin(iPluginList ...IPlugin) {
	for _, plugin := range iPluginList {
		p.RegPlugin(plugin)
	}
}

// RegPlugin 注册插件
func (p *PluginManager) RegPlugin(iPlugin IPlugin) {
	plugin := iPlugin.GetPlugin()
	// 校验插件名是否已被注册
	_, exist := p.pluginMap[plugin.PluginName]
	if exist {
		logger.Error("plugin has been register, name: %v", plugin.PluginName)
		return
	}
	logger.Info("plugin enable, name: %v", plugin.PluginName)
	// 调用插件启用的生命周期
	iPlugin.OnEnable()
	p.pluginMap[plugin.PluginName] = iPlugin
}

// DelAllPlugin 卸载全部插件
func (p *PluginManager) DelAllPlugin() {
	for _, plugin := range p.pluginMap {
		p.DelPlugin(plugin)
	}
}

// DelPlugin 卸载插件
func (p *PluginManager) DelPlugin(iPlugin IPlugin) {
	plugin := iPlugin.GetPlugin()
	// 校验插件是否注册
	_, exist := p.pluginMap[plugin.PluginName]
	if !exist {
		logger.Error("plugin not exist, name: %v", plugin.PluginName)
		return
	}
	logger.Info("plugin disable, name: %v", plugin.PluginName)
	// 调用插件禁用的生命周期
	iPlugin.OnDisable()
	delete(p.pluginMap, plugin.PluginName)
}

// GetIPlugin 获取抽象插件
func (p *PluginManager) GetIPlugin(pluginName string) IPlugin {
	iPlugin, exist := p.pluginMap[pluginName]
	if !exist {
		logger.Error("plugin not exist, name: %v", pluginName)
		return nil
	}
	return iPlugin
}

// TriggerEvent 触发事件
func (p *PluginManager) TriggerEvent(eventId PluginEventId, event IPluginEvent) bool {
	// 获取每个插件监听的事件并根据优先级排序
	eventInfoList := make([]*PluginEventInfo, 0)
	for _, iPlugin := range p.pluginMap {
		plugin := iPlugin.GetPlugin()
		// 插件未启用则跳过
		if !plugin.isEnable {
			continue
		}
		// 获取插件事件列表
		infoList, exist := plugin.eventMap[eventId]
		if !exist {
			continue
		}
		for _, info := range infoList {
			eventInfoList = append(eventInfoList, info)
		}
	}
	// 根据优先级排序
	sort.Slice(eventInfoList, func(i, j int) bool {
		return eventInfoList[i].Priority > eventInfoList[j].Priority
	})
	// 执行每个处理函数
	for _, info := range eventInfoList {
		info.EventFunc(event)
	}
	// 判断事件是否被取消
	return event.IsCancel()
}

// HandleGlobalTick 处理全局tick
func (p *PluginManager) HandleGlobalTick(tick PluginGlobalTick) {
	for _, iPlugin := range p.pluginMap {
		plugin := iPlugin.GetPlugin()
		// 插件未启用则跳过
		if !plugin.isEnable {
			continue
		}
		// 获取插件tick处理函数列表
		tickFuncList, exist := plugin.globalTickMap[tick]
		if !exist {
			continue
		}
		for _, tickFunc := range tickFuncList {
			tickFunc()
		}
	}
}

// HandleUserTimer 处理用户timer
func (p *PluginManager) HandleUserTimer(player *model.Player, data []any) {
	// 如果创建的用户timer没有id数据则报错
	if len(data) < 1 {
		logger.Error("data len less 1, len: %v", len(data))
		return
	}
	userTimerId := data[0].(uint64)
	data = data[1:]
	// 通知插件
	for _, iPlugin := range p.pluginMap {
		plugin := iPlugin.GetPlugin()
		// 插件未启用则跳过
		if !plugin.isEnable {
			continue
		}
		// 获取插件用户timer处理函数列表
		timerFunc, exist := plugin.userTimerMap[userTimerId]
		if !exist {
			continue
		}
		timerFunc(player, data)
		// 只需要执行一次
		break
	}
}
