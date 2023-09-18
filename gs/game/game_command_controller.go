package game

import (
	"fmt"
	"hk4e/gdconf"
	"strconv"
	"strings"
)

// 玩家游戏内GM命令格式解析模块

// CommandController 命令控制器
type CommandController struct {
	Name        string      // 名称
	AliasList   []string    // 别名列表
	Description string      // 命令描述
	Usage       string      // 用法描述
	Perm        CommandPerm // 权限
	Func        CommandFunc // 命令执行函数
}

// InitControllerList 初始化控制器列表
func (c *CommandManager) InitControllerList() {
	c.commandControllerList = []*CommandController{
		// 权限等级 0: 普通玩家
		HelpCommandController,
		TeleportCommandController,
		GiveCommandController,
		QuestCommandController,
		UnlockCommandController,
		GcgCommandController,
		// 权限等级 1: GM 1级
		AssignCommandController,
	}
}

// 指定命令

var AssignCommandController = &CommandController{
	Name:        "指定",
	AliasList:   []string{"assign"},
	Description: "<color=#FFFFCC>assign</color> <color=#FFCC99>设置命令指定玩家</color>",
	Usage:       "命令格式：\n1. {alias} <目标UID> 指定某个玩家",
	Perm:        CommandPermGM,
	Func:        AssignCommand,
}

func AssignCommand(c *CommandContent) {
	var assignUid uint32

	c.Dynamic("uint32", func(param any) bool {
		value := param.(uint32)
		// 指定uid
		assignUid = value
		return true
	}).Execute(func() bool {
		// 设置命令指定uid
		c.Executor.CommandAssignUid = assignUid
		c.SendSuccMessage(c.Executor, "已指定玩家，指定UID：%v", assignUid)
		return true
	})
}

// 帮助命令

var HelpCommandController = &CommandController{
	Name:        "帮助",
	AliasList:   []string{"help"},
	Description: "<color=#FFFFCC>help</color> <color=#FFCC99>我需要帮助！</color>",
	Usage:       "命令格式：\n1. {alias} 查看简要帮助信息\n\n2. {alias} <序号/命令别名> 查看详细帮助信息",
	Perm:        CommandPermNormal,
	Func:        HelpCommand,
}

func HelpCommand(c *CommandContent) {
	var controller *CommandController // 命令控制器
	var alias string                  // 别名

	c.SetElse(func() {
		// 显示简要帮助信息
		helpText := "<color=#66B2FF>================</color><color=#CCE5FF>/ 帮 助 /</color><color=#66B2FF>================</color>\n"
		for i, controller := range COMMAND_MANAGER.commandControllerList {
			// 权限不足跳过
			if c.Executor.CmdPerm < uint8(controller.Perm) {
				continue
			}
			// GM命令和普通命令区分颜色
			var permColor string
			switch controller.Perm {
			case CommandPermNormal:
				permColor = "#CCFFCC"
			case CommandPermGM:
				permColor = "#FF9999"
			}
			helpText += fmt.Sprintf("<color=%v>%v. %v命令</color> <color=#FFE5CC>-</color> %v\n", permColor, strconv.Itoa(i+1), controller.Name, controller.Description)
		}
		helpText += "\n<color=#FFFFCC>help</color> <color=#FFCCE5><命令别名></color> <color=#FF9999>能查看详细用法哦~</color>\n"
		helpText += "<color=#FF6347><></color> <color=#87CEFA>代表必填参数</color> <color=#FF6347>[]</color> <color=#87CEFA>代表可选参数</color> <color=#FF6347>/</color> <color=#87CEFA>代表或者</color>"
		c.SendMessage(c.Executor, helpText)
	}).Dynamic("string", func(param any) bool {
		value := param.(string)
		// 通过别名获取
		controller = COMMAND_MANAGER.commandControllerMap[value]
		if controller == nil {
			return false
		}
		alias = value
		return true
	}).Execute(func() bool {
		text := fmt.Sprintf("<color=#FFFFCC>%v</color><color=#CCCCFF> 命令详细帮助：</color>\n\n%v\n\n<color=#CCE5FF>所有别名：</color><color=#E0E0E0>%v</color>", controller.Name, controller.Usage, controller.AliasList)
		text = strings.ReplaceAll(text, "{alias}", alias)
		c.SendMessage(c.Executor, text)
		return true
	})
}

// 传送命令

var TeleportCommandController = &CommandController{
	Name:        "传送",
	AliasList:   []string{"tp", "teleport", "goto"},
	Description: "<color=#FFFFCC>tp</color> <color=#FFCC99>传送到世界的任何角落～</color>",
	Usage:       "命令格式：\n1. {alias} <目标UID> 传送至目标玩家\n\n2. {alias} <坐标X> <坐标Y> <坐标Z> [场景ID] 传送至指定场景及坐标",
	Perm:        CommandPermNormal,
	Func:        TeleportCommand,
}

func TeleportCommand(c *CommandContent) {
	// 计算相对坐标
	parseRelativePosFunc := func(param string, pos float64) (float64, bool) {
		// 不以 ~ 开头代表使用绝对坐标
		if !strings.HasPrefix(param, "~") {
			value, err := strconv.ParseFloat(param, 64)
			return value, err == nil
		}
		// 用户只输入 ~ 获取为玩家当前位置
		if param == "~" {
			return pos, true
		}
		// 以 ~ 开头 此时位置加 ~ 后的数
		param = param[1:] // 去除 ~
		addPos, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return 0, false
		}
		// 计算坐标
		pos += addPos
		return pos, true
	}
	// 根据参数数量做不同的逻辑
	switch len(c.ParamList) {
	case 1:
		// 传送至某个玩家
		var targetUid uint32 // 目标玩家uid

		c.Dynamic("uint32", func(param any) bool {
			value := param.(uint32)
			// 目标玩家
			targetUid = value
			return true
		}).Execute(func() bool {
			// 判断目标用户是否在线
			target := USER_MANAGER.GetOnlineUser(targetUid)
			// 目标玩家属于非本地玩家
			if target == nil && !USER_MANAGER.GetRemoteUserOnlineState(targetUid) {
				c.SetElse(func() {
					// 全服不存在该在线玩家
					c.SendFailMessage(c.Executor, "目标玩家不在线，UID：%v。", targetUid)
				})
				return false
			}
			// 如果玩家不与目标玩家同一世界或不同服务器
			if target == nil || c.AssignPlayer.WorldId != target.WorldId {
				// 请求进入目标玩家世界
				GAME.PlayerApplyEnterWorld(c.AssignPlayer, target.PlayerId)
				// 发送消息给执行者
				c.SendSuccMessage(c.Executor, "已请求加入目标玩家世界，指定UID：%v，目标UID：%v。", c.AssignPlayer.PlayerId, target.PlayerId)
			} else {
				// 传送玩家至目标玩家的位置
				COMMAND_MANAGER.gmCmd.GMTeleportPlayer(c.AssignPlayer.PlayerId, target.SceneId, target.Pos.X, target.Pos.Y, target.Pos.Z)
				// 发送消息给执行者
				c.SendSuccMessage(c.Executor, "已传送至目标玩家，指定UID：%v，目标UID：%v。", c.AssignPlayer.PlayerId, target.PlayerId)
			}
			return true
		})
	case 5, 4, 3:
		// 传送玩家到场景以及坐标
		var sceneId = c.AssignPlayer.SceneId
		var posX, posY, posZ float64

		// 解析命令
		c.Dynamic("string", func(param any) bool {
			value := param.(string)
			// 坐标x
			pos, ok := parseRelativePosFunc(value, c.AssignPlayer.Pos.X)
			posX = pos
			return ok
		}).Dynamic("string", func(param any) bool {
			value := param.(string)
			// 坐标y
			pos, ok := parseRelativePosFunc(value, c.AssignPlayer.Pos.Y)
			posY = pos
			return ok
		}).Dynamic("string", func(param any) bool {
			value := param.(string)
			// 坐标z
			pos, ok := parseRelativePosFunc(value, c.AssignPlayer.Pos.Z)
			posZ = pos
			return ok
		}).Option("uint32", func(param any) bool {
			value := param.(uint32)
			// 场景id
			sceneId = value
			return true
		}).Execute(func() bool {
			// 传送玩家至指定的位置
			COMMAND_MANAGER.gmCmd.GMTeleportPlayer(c.AssignPlayer.PlayerId, sceneId, posX, posY, posZ)
			// 发送消息给执行者
			c.SendSuccMessage(c.Executor, "已传送至指定位置，指定UID：%v，场景：%v，X：%.2f，Y：%.2f，Z：%.2f。", c.AssignPlayer.PlayerId, c.AssignPlayer.SceneId, posX, posY, posZ)
			return true
		})
	}
}

// 给予命令

var GiveCommandController = &CommandController{
	Name:        "给予",
	AliasList:   []string{"give"},
	Description: "<color=#FFFFCC>give</color> <color=#FFCC99>获得些不得了的物品</color>",
	Usage:       "模式：ID / item (所有物品) / weapon (所有武器) / reliquary (所有圣遗物) / avatar (所有角色) / costume (所有时装) / flycloak (所有风之翼) / all (全部)\n不要加上括号内的中文！！\n\n命令格式：\n1. {alias} <模式> [参数1] [参数2] [参数3] 给予指定物品\n\n通常 参数1为数量\n给予角色时 参数1为等级\n给予武器时 参数2 等级 参数3 为精炼等级",
	Perm:        CommandPermNormal,
	Func:        GiveCommand,
}

func GiveCommand(c *CommandContent) {
	// 给予物品
	var mode string     // 给予的模式
	var itemId uint32   // 物品id
	var arg1 uint32 = 1 // 参数1 数量或等级
	var arg2 uint32     // 参数2 武器等级
	var arg3 uint32     // 参数2 精炼等级

	c.Dynamic("string", func(param any) bool {
		value := param.(string)
		// 给予的物品
		switch param {
		case "item", "weapon", "reliquary", "avatar", "costume", "flycloak", "all":
			// 所有武器 所有圣遗物 所有角色 所有时装 所有风之翼 一切
			mode = value
		default:
			id, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				return false
			}
			mode = "id"
			itemId = uint32(id)
		}
		return true
	}).Option("uint32", func(param any) bool {
		value := param.(uint32)
		// 参数1
		arg1 = value
		return true
	}).Option("uint32", func(param any) bool {
		value := param.(uint32)
		// 参数2
		arg2 = value
		return true
	}).Option("uint32", func(param any) bool {
		value := param.(uint32)
		// 参数3
		arg2 = value
		return true
	}).Execute(func() bool {
		switch mode {
		case "id":
			// 判断是否为物品
			_, ok := GAME.GetAllItemDataConfig()[int32(itemId)]
			if ok {
				// 给予玩家物品
				COMMAND_MANAGER.gmCmd.GMAddItem(c.AssignPlayer.PlayerId, itemId, arg1)
				c.SendSuccMessage(c.Executor, "已给予物品，指定UID：%v，ID：%v，数量：%v。", c.AssignPlayer.PlayerId, itemId, arg1)
				return true
			}
			// 判断是否为武器
			_, ok = GAME.GetAllWeaponDataConfig()[int32(itemId)]
			if ok {
				// 给予玩家武器
				COMMAND_MANAGER.gmCmd.GMAddWeapon(c.AssignPlayer.PlayerId, itemId, arg1, uint8(arg2), uint8(arg3))
				c.SendSuccMessage(c.Executor, "已给予武器，指定UID：%v，ID：%v，数量：%v，等级：%v，精炼等级：%v。", c.AssignPlayer.PlayerId, itemId, arg1, arg2, arg3)
				return true
			}
			// 判断是否为圣遗物
			_, ok = GAME.GetAllReliquaryDataConfig()[int32(itemId)]
			if ok {
				// 给予玩家圣遗物
				COMMAND_MANAGER.gmCmd.GMAddReliquary(c.AssignPlayer.PlayerId, itemId, arg1)
				c.SendSuccMessage(c.Executor, "已给予圣遗物，指定UID：%v，ID：%v，数量：%v。", c.AssignPlayer.PlayerId, itemId, arg1)
				return true
			}
			// 判断是否为角色
			_, ok = GAME.GetAllAvatarDataConfig()[int32(itemId)]
			if ok {
				// 给予玩家角色
				COMMAND_MANAGER.gmCmd.GMAddAvatar(c.AssignPlayer.PlayerId, itemId, uint8(arg1))
				c.SendSuccMessage(c.Executor, "已给予角色，指定UID：%v，ID：%v，等级：%v。", c.AssignPlayer.PlayerId, itemId, arg1)
				return true
			}
			// 判断是否为时装
			if gdconf.GetAvatarCostumeDataById(int32(itemId)) != nil {
				// 给予玩家时装
				COMMAND_MANAGER.gmCmd.GMAddCostume(c.AssignPlayer.PlayerId, itemId)
				c.SendSuccMessage(c.Executor, "已给予时装，指定UID：%v，ID：%v，数量：%v。", c.AssignPlayer.PlayerId, itemId, arg1)
				return true
			}
			// 判断是否为风之翼
			if gdconf.GetAvatarFlycloakDataById(int32(itemId)) != nil {
				// 给予玩家风之翼
				COMMAND_MANAGER.gmCmd.GMAddFlycloak(c.AssignPlayer.PlayerId, itemId)
				c.SendSuccMessage(c.Executor, "已给予风之翼，指定UID：%v，ID：%v，数量：%v。", c.AssignPlayer.PlayerId, itemId, arg1)
				return true
			}
			// 都执行到这里那肯定是都不匹配
			c.SetElse(func() {
				// 物品id不存在
				c.SendFailMessage(c.Executor, "不存在的物品，ID：%v。", itemId)
			})
			return false
		case "item":
			// 给予玩家所有物品
			COMMAND_MANAGER.gmCmd.GMAddAllItem(c.AssignPlayer.PlayerId, arg1)
			c.SendSuccMessage(c.Executor, "已给予所有物品，指定UID：%v，数量：%v。", c.AssignPlayer.PlayerId, arg1)
		case "weapon":
			// 给予玩家所有武器
			COMMAND_MANAGER.gmCmd.GMAddAllWeapon(c.AssignPlayer.PlayerId, arg1, uint8(arg2), uint8(arg3))
			c.SendSuccMessage(c.Executor, "已给予所有武器，指定UID：%v，数量：%v，等级：%v，精炼等级：%v。", c.AssignPlayer.PlayerId, itemId, arg1, arg2, arg3)
		case "reliquary":
			// 给予玩家所有圣遗物
			COMMAND_MANAGER.gmCmd.GMAddAllReliquary(c.AssignPlayer.PlayerId, arg1)
			c.SendSuccMessage(c.Executor, "已给予所有圣遗物，指定UID：%v，数量：%v。", c.AssignPlayer.PlayerId, arg1)
		case "avatar":
			// 给予玩家所有角色
			COMMAND_MANAGER.gmCmd.GMAddAllAvatar(c.AssignPlayer.PlayerId, uint8(arg1))
			c.SendSuccMessage(c.Executor, "已给予所有角色，指定UID：%v，等级：%v。", c.AssignPlayer.PlayerId, arg1)
		case "costume":
			// 给予玩家所有时装
			COMMAND_MANAGER.gmCmd.GMAddAllCostume(c.AssignPlayer.PlayerId)
			c.SendSuccMessage(c.Executor, "已给予所有时装，指定UID：%v。", c.AssignPlayer.PlayerId)
		case "flycloak":
			// 给予玩家所有风之翼
			COMMAND_MANAGER.gmCmd.GMAddAllFlycloak(c.AssignPlayer.PlayerId)
			c.SendSuccMessage(c.Executor, "已给予所有风之翼，指定UID：%v。", c.AssignPlayer.PlayerId)
		case "all":
			// 给予玩家所有内容
			COMMAND_MANAGER.gmCmd.GMAddAll(c.AssignPlayer.PlayerId)
			c.SendSuccMessage(c.Executor, "已给予所有内容，指定UID：%v。", c.AssignPlayer.PlayerId)
		default:
			return false
		}
		return true
	})
}

// 任务命令

var QuestCommandController = &CommandController{
	Name:        "任务",
	AliasList:   []string{"quest"},
	Description: "<color=#FFFFCC>quest</color> <color=#FFCC99>管理你的任务</color>",
	Usage:       "命令格式：\n1. {alias} <add/finish/allfinish> <任务ID> 添加/完成/完成全部 任务",
	Perm:        CommandPermNormal,
	Func:        QuestCommand,
}

func QuestCommand(c *CommandContent) {
	var mode string    // 模式
	var questId uint32 // 任务id

	c.Dynamic("string", func(param any) bool {
		value := param.(string)
		// 模式
		switch value {
		case "add", "finish", "allfinish":
			// 添加 完成 完成全部
			mode = value
			return true
		}
		return false
	}).Dynamic("uint32", func(param any) bool {
		value := param.(uint32)
		// 任务id
		questId = value
		return true
	}).Execute(func() bool {
		switch mode {
		case "add":
			// 添加指定任务
			COMMAND_MANAGER.gmCmd.GMAddQuest(c.AssignPlayer.PlayerId, questId)
			c.SendSuccMessage(c.Executor, "已添加任务，指定UID：%v，任务ID：%v。", c.AssignPlayer.PlayerId, questId)
		case "finish":
			// 完成指定任务
			COMMAND_MANAGER.gmCmd.GMFinishQuest(c.AssignPlayer.PlayerId, questId)
			c.SendSuccMessage(c.Executor, "已完成玩家任务，指定UID：%v，任务ID：%v。", c.AssignPlayer.PlayerId, questId)
		case "allfinish":
			// 强制完成当前所有任务
			COMMAND_MANAGER.gmCmd.GMForceFinishAllQuest(c.AssignPlayer.PlayerId)
			c.SendSuccMessage(c.Executor, "已完成当前全部任务，指定UID：%v。", c.AssignPlayer.PlayerId, questId)
		default:
			return false
		}
		return true
	})
}

// 解锁命令

var UnlockCommandController = &CommandController{
	Name:        "解锁",
	AliasList:   []string{"unlock"},
	Description: "<color=#FFFFCC>unlock</color> <color=#FFCC99>解锁游戏进度</color>",
	Usage:       "命令格式：\n1. {alias} <allpoint/xluadebug> 解锁全部锚点/主动开关客户端XLUA调试",
	Perm:        CommandPermNormal,
	Func:        UnlockCommand,
}

func UnlockCommand(c *CommandContent) {
	var mode string // 模式

	c.Dynamic("string", func(param any) bool {
		value := param.(string)
		// 模式
		switch value {
		case "allpoint", "xluadebug":
			// 解锁全部锚点 主动开关客户端XLUA调试
			mode = value
			return true
		}
		return false
	}).Execute(func() bool {
		switch mode {
		case "allpoint":
			// 解锁全部锚点
			COMMAND_MANAGER.gmCmd.GMUnlockAllPoint(c.AssignPlayer.PlayerId, c.AssignPlayer.SceneId)
			c.SendSuccMessage(c.Executor, "已解锁所有锚点，指定UID：%v，场景：%v。", c.AssignPlayer.PlayerId, c.AssignPlayer.SceneId)
		case "xluadebug":
			// 主动开关客户端XLUA调试
			if !c.AssignPlayer.XLuaDebug {
				c.AssignPlayer.XLuaDebug = true
				c.SendSuccMessage(c.Executor, "已开启客户端XLUA调试，指定UID：%v。", c.AssignPlayer.PlayerId)
			} else {
				c.AssignPlayer.XLuaDebug = false
				c.SendSuccMessage(c.Executor, "已关闭客户端XLUA调试，指定UID：%v。", c.AssignPlayer.PlayerId)
			}
			return true
		default:
			return false
		}
		return true
	})
}

// 七圣召唤测试命令

var GcgCommandController = &CommandController{
	Name:        "七圣召唤测试",
	AliasList:   []string{"gcgtest"},
	Description: "<color=#FFFFCC>unlock</color> <color=#FFCC99>七圣召唤测试命令</color>",
	Usage:       "命令格式：\n1. {alias} 测试七圣召唤",
	Perm:        CommandPermNormal,
	Func:        GcgCommand,
}

func GcgCommand(c *CommandContent) {
	c.Execute(func() bool {
		// 开始七圣召唤对局
		GAME.GCGStartChallenge(c.AssignPlayer)
		c.SendSuccMessage(c.Executor, "已开始七圣召唤对局，指定UID：%v。", c.AssignPlayer.PlayerId)
		return true
	})
}
