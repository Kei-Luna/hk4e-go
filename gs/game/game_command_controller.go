package game

import (
	"fmt"
	"hk4e/gdconf"
	"hk4e/pkg/logger"
	"strconv"
	"strings"
)

// 玩家游戏内GM命令格式解析模块

// CommandController 命令控制器
type CommandController struct {
	Name        string      // 名称
	AliasList   []string    // 别名列表
	Description string      // 命令描述
	UsageList   []string    // 用法描述
	Perm        CommandPerm // 权限
	Func        CommandFunc // 命令执行函数
}

// InitControllerList 初始化控制器列表
func (c *CommandManager) InitControllerList() {
	c.commandControllerList = []*CommandController{
		// 权限等级 0: 普通玩家
		HelpCommandController,
		GotoCommandController,
		JumpCommandController,
		EquipCommandController,
		ItemCommandController,
		AvatarCommandController,
		GiveCommandController,
		KillCommandController,
		QuestCommandController,
		PointCommandController,
		XLuaDebugCommandController,
		GcgCommandController,
		// 权限等级 1: GM 1级
		AssignCommandController,
	}
}

// 指定命令

var AssignCommandController = &CommandController{
	Name:        "指定",
	AliasList:   []string{"assign"},
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>设置命令指定玩家</color>",
	UsageList: []string{
		"{alias} <目标UID> 指定某个玩家",
	},
	Perm: CommandPermGM,
	Func: AssignCommand,
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
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>查看简要帮助信息</color>",
	UsageList: []string{
		"{alias} 查看简要帮助信息",
		"{alias} <序号/命令别名> 查看详细帮助信息",
	},
	Perm: CommandPermNormal,
	Func: HelpCommand,
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
			helpText += fmt.Sprintf("<color=%v>%v. %v</color> <color=#FFE5CC>-</color> %v\n", permColor, strconv.Itoa(i+1), controller.Name, strings.ReplaceAll(controller.Description, "{alias}", controller.AliasList[0]))
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
		usage := "命令用法：\n"
		for i, s := range controller.UsageList {
			s = strings.ReplaceAll(s, "{alias}", alias)
			usage += fmt.Sprintf("%v. %v", i+1, s)
			// 换行
			if i != len(controller.UsageList)-1 {
				usage += "\n"
			}
		}
		text := fmt.Sprintf("<color=#FFFFCC>%v</color><color=#CCCCFF> 命令详细帮助：</color>\n\n%v\n\n<color=#CCE5FF>所有别名：</color><color=#E0E0E0>%v</color>", controller.Name, usage, controller.AliasList)
		c.SendMessage(c.Executor, text)
		return true
	})
}

// 传送坐标命令

var GotoCommandController = &CommandController{
	Name:        "传送坐标",
	AliasList:   []string{"goto"},
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>传送到指定坐标</color>",
	UsageList: []string{
		"{alias} <坐标X> <坐标Y> <坐标Z> 传送至指定坐标",
	},
	Perm: CommandPermNormal,
	Func: GotoCommand,
}

func GotoCommand(c *CommandContent) {
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
	// 传送玩家到场景以及坐标
	var sceneId = c.AssignPlayer.SceneId
	var posX, posY, posZ float64

	// 解析命令
	c.Dynamic("string", func(param any) bool {
		// 坐标x
		value := param.(string)
		pos, ok := parseRelativePosFunc(value, c.AssignPlayer.Pos.X)
		posX = pos
		return ok
	}).Dynamic("string", func(param any) bool {
		// 坐标y
		value := param.(string)
		pos, ok := parseRelativePosFunc(value, c.AssignPlayer.Pos.Y)
		posY = pos
		return ok
	}).Dynamic("string", func(param any) bool {
		// 坐标z
		value := param.(string)
		pos, ok := parseRelativePosFunc(value, c.AssignPlayer.Pos.Z)
		posZ = pos
		return ok
	}).Execute(func() bool {
		// 传送玩家至指定的位置
		COMMAND_MANAGER.gmCmd.GMTeleportPlayer(c.AssignPlayer.PlayerId, sceneId, posX, posY, posZ)
		// 发送消息给执行者
		c.SendSuccMessage(c.Executor, "已传送至指定位置，指定UID：%v，场景ID：%v，X：%.2f，Y：%.2f，Z：%.2f。", c.AssignPlayer.PlayerId, c.AssignPlayer.SceneId, posX, posY, posZ)
		return true
	})
}

// 传送场景命令

var JumpCommandController = &CommandController{
	Name:        "传送场景",
	AliasList:   []string{"jump"},
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>传送到至指定场景</color>",
	UsageList: []string{
		"{alias} <场景ID> 传送至指定场景",
	},
	Perm: CommandPermNormal,
	Func: JumpCommand,
}

func JumpCommand(c *CommandContent) {
	var sceneId uint32 // 场景id

	c.Dynamic("uint32", func(param any) bool {
		// 场景id
		sceneId = param.(uint32)
		return true
	}).Execute(func() bool {
		var posX float64
		var posY float64
		var posZ float64
		// 读取场景初始位置
		sceneLuaConfig := gdconf.GetSceneLuaConfigById(int32(sceneId))
		if sceneLuaConfig != nil {
			bornPos := sceneLuaConfig.SceneConfig.BornPos
			posX = float64(bornPos.X)
			posY = float64(bornPos.Y)
			posZ = float64(bornPos.Z)
		} else {
			logger.Error("get scene lua config is nil, sceneId: %v, uid: %v", sceneId, c.AssignPlayer.PlayerId)
		}
		// 传送玩家至指定的位置
		COMMAND_MANAGER.gmCmd.GMTeleportPlayer(c.AssignPlayer.PlayerId, sceneId, posX, posY, posZ)
		// 发送消息给执行者
		c.SendSuccMessage(c.Executor, "已传送至指定场景，指定UID：%v，场景ID：%v，X：%.2f，Y：%.2f，Z：%.2f。", c.AssignPlayer.PlayerId, c.AssignPlayer.SceneId, posX, posY, posZ)
		return true
	})
}

// 管理武器命令

var EquipCommandController = &CommandController{
	Name:        "管理武器",
	AliasList:   []string{"equip", "weapon"},
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>管理你的武器</color>",
	UsageList: []string{
		"{alias} add <武器ID/all> [武器等级] [突破等级] 添加武器",
	},
	Perm: CommandPermNormal,
	Func: EquipCommand,
}

func EquipCommand(c *CommandContent) {
	var mode string     // 模式
	var param1 string   // 参数1
	var level uint8 = 1 // 武器等级
	var promote uint8   // 突破等级

	c.Dynamic("string", func(param any) bool {
		// 模式
		mode = param.(string)
		return true
	}).Dynamic("string", func(param any) bool {
		// 参数1
		param1 = param.(string)
		return true
	}).Option("uint8", func(param any) bool {
		// 武器等级
		level = param.(uint8)
		return true
	}).Option("uint8", func(param any) bool {
		// 突破等级
		promote = param.(uint8)
		return true
	}).Execute(func() bool {
		switch mode {
		case "add":
			// 添加武器
			// 判断是否要添加全部武器
			if param1 == "all" {
				COMMAND_MANAGER.gmCmd.GMAddAllWeapon(c.AssignPlayer.PlayerId, 1, level, promote, 1)
				c.SendSuccMessage(c.Executor, "已给予所有武器，指定UID：%v，武器等级：%v，突破等级：%v。", c.AssignPlayer.PlayerId, level, promote)
				return true
			}
			// 物品id
			itemId, err := strconv.ParseUint(param1, 10, 32)
			if err != nil {
				return false
			}
			COMMAND_MANAGER.gmCmd.GMAddWeapon(c.AssignPlayer.PlayerId, uint32(itemId), 1, level, promote, 1)
			c.SendSuccMessage(c.Executor, "已给予武器，指定UID：%v，武器ID：%v，武器等级：%v，突破等级：%v。", c.AssignPlayer.PlayerId, itemId, level, promote)
		default:
			return false
		}
		return true
	})
}

// 管理物品命令

var ItemCommandController = &CommandController{
	Name:        "管理物品",
	AliasList:   []string{"item"},
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>管理你的物品</color>",
	UsageList: []string{
		"{alias} add <物品ID/all> [数量] 添加物品",
	},
	Perm: CommandPermNormal,
	Func: ItemCommand,
}

func ItemCommand(c *CommandContent) {
	var mode string      // 模式
	var param1 string    // 参数1
	var count uint32 = 1 // 数量

	c.Dynamic("string", func(param any) bool {
		// 模式
		mode = param.(string)
		return true
	}).Dynamic("string", func(param any) bool {
		// 参数1
		param1 = param.(string)
		return true
	}).Option("uint32", func(param any) bool {
		// 数量
		count = param.(uint32)
		return true
	}).Execute(func() bool {
		switch mode {
		case "add":
			// 添加物品
			// 判断是否要添加全部物品
			if param1 == "all" {
				COMMAND_MANAGER.gmCmd.GMAddAllItem(c.AssignPlayer.PlayerId, count)
				c.SendSuccMessage(c.Executor, "已给予所有物品，指定UID：%v，数量：%v。", c.AssignPlayer.PlayerId, count)
				return true
			}
			// 物品id
			itemId, err := strconv.ParseUint(param1, 10, 32)
			if err != nil {
				return false
			}
			COMMAND_MANAGER.gmCmd.GMAddItem(c.AssignPlayer.PlayerId, uint32(itemId), count)
			c.SendSuccMessage(c.Executor, "已给予物品，指定UID：%v，物品ID：%v，数量：%v。", c.AssignPlayer.PlayerId, itemId, count)
		default:
			return false
		}
		return true
	})
}

// 管理角色命令

var AvatarCommandController = &CommandController{
	Name:        "管理角色",
	AliasList:   []string{"avatar"},
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>管理你的角色</color>",
	UsageList: []string{
		"{alias} add <角色ID/all>",
	},
	Perm: CommandPermNormal,
	Func: AvatarCommand,
}

func AvatarCommand(c *CommandContent) {
	var mode string   // 模式
	var param1 string // 参数1

	c.Dynamic("string", func(param any) bool {
		// 模式
		mode = param.(string)
		return true
	}).Dynamic("string", func(param any) bool {
		// 参数1
		param1 = param.(string)
		return true
	}).Execute(func() bool {
		switch mode {
		case "add":
			// 添加角色
			// 判断是否要添加全部角色
			if param1 == "all" {
				COMMAND_MANAGER.gmCmd.GMAddAllAvatar(c.AssignPlayer.PlayerId, 1, 0)
				c.SendSuccMessage(c.Executor, "已给予所有角色，指定UID：%v。", c.AssignPlayer.PlayerId)
				return true
			}
			// 角色id
			avatarId, err := strconv.ParseUint(param1, 10, 32)
			if err != nil {
				return false
			}
			COMMAND_MANAGER.gmCmd.GMAddAvatar(c.AssignPlayer.PlayerId, uint32(avatarId), 1, 0)
			c.SendSuccMessage(c.Executor, "已给予角色，指定UID：%v，角色ID：%v。", c.AssignPlayer.PlayerId, avatarId)
		default:
			return false
		}
		return true
	})
}

// 给予命令

var GiveCommandController = &CommandController{
	Name:        "给予物品",
	AliasList:   []string{"give"},
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>获得全部物品</color>",
	UsageList: []string{
		"模式：ID / item (所有物品) / weapon (所有武器) / reliquary (所有圣遗物) / avatar (所有角色) / costume (所有时装) / flycloak (所有风之翼) / all (全部)\n不要加上括号内的中文！！",
		"{alias} <模式> [数量] 给予指定物品",
		"数量仅物品、武器、圣遗物可用",
	},
	Perm: CommandPermNormal,
	Func: GiveCommand,
}

func GiveCommand(c *CommandContent) {
	// 给予物品
	var mode string      // 给予的模式
	var itemId uint32    // 物品id
	var count uint32 = 1 // 数量

	c.Dynamic("string", func(param any) bool {
		value := param.(string)
		// 给予的物品
		switch value {
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
		count = value
		return true
	}).Execute(func() bool {
		switch mode {
		case "id":
			// 判断是否为物品
			_, ok := GAME.GetAllItemDataConfig()[int32(itemId)]
			if ok {
				// 给予玩家物品
				COMMAND_MANAGER.gmCmd.GMAddItem(c.AssignPlayer.PlayerId, itemId, count)
				c.SendSuccMessage(c.Executor, "已给予物品，指定UID：%v，ID：%v，数量：%v。", c.AssignPlayer.PlayerId, itemId, count)
				return true
			}
			// 判断是否为武器
			_, ok = GAME.GetAllWeaponDataConfig()[int32(itemId)]
			if ok {
				// 给予玩家武器
				COMMAND_MANAGER.gmCmd.GMAddWeapon(c.AssignPlayer.PlayerId, itemId, count, 1, 0, 1)
				c.SendSuccMessage(c.Executor, "已给予武器，指定UID：%v，ID：%v，数量：%v。", c.AssignPlayer.PlayerId, itemId, count)
				return true
			}
			// 判断是否为圣遗物
			_, ok = GAME.GetAllReliquaryDataConfig()[int32(itemId)]
			if ok {
				// 给予玩家圣遗物
				COMMAND_MANAGER.gmCmd.GMAddReliquary(c.AssignPlayer.PlayerId, itemId, count)
				c.SendSuccMessage(c.Executor, "已给予圣遗物，指定UID：%v，ID：%v，数量：%v。", c.AssignPlayer.PlayerId, itemId, count)
				return true
			}
			// 判断是否为角色
			_, ok = GAME.GetAllAvatarDataConfig()[int32(itemId)]
			if ok {
				// 给予玩家角色
				COMMAND_MANAGER.gmCmd.GMAddAvatar(c.AssignPlayer.PlayerId, itemId, 1, 0)
				c.SendSuccMessage(c.Executor, "已给予角色，指定UID：%v，ID：%v。", c.AssignPlayer.PlayerId, itemId)
				return true
			}
			// 判断是否为时装
			if gdconf.GetAvatarCostumeDataById(int32(itemId)) != nil {
				// 给予玩家时装
				COMMAND_MANAGER.gmCmd.GMAddCostume(c.AssignPlayer.PlayerId, itemId)
				c.SendSuccMessage(c.Executor, "已给予时装，指定UID：%v，ID：%v。", c.AssignPlayer.PlayerId, itemId)
				return true
			}
			// 判断是否为风之翼
			if gdconf.GetAvatarFlycloakDataById(int32(itemId)) != nil {
				// 给予玩家风之翼
				COMMAND_MANAGER.gmCmd.GMAddFlycloak(c.AssignPlayer.PlayerId, itemId)
				c.SendSuccMessage(c.Executor, "已给予风之翼，指定UID：%v，ID：%v。", c.AssignPlayer.PlayerId, itemId)
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
			COMMAND_MANAGER.gmCmd.GMAddAllItem(c.AssignPlayer.PlayerId, count)
			c.SendSuccMessage(c.Executor, "已给予所有物品，指定UID：%v，数量：%v。", c.AssignPlayer.PlayerId, count)
		case "weapon":
			// 给予玩家所有武器
			COMMAND_MANAGER.gmCmd.GMAddAllWeapon(c.AssignPlayer.PlayerId, count, 1, 0, 1)
			c.SendSuccMessage(c.Executor, "已给予所有武器，指定UID：%v，数量：%v。", c.AssignPlayer.PlayerId, itemId, count)
		case "reliquary":
			// 给予玩家所有圣遗物
			COMMAND_MANAGER.gmCmd.GMAddAllReliquary(c.AssignPlayer.PlayerId, count)
			c.SendSuccMessage(c.Executor, "已给予所有圣遗物，指定UID：%v，数量：%v。", c.AssignPlayer.PlayerId, count)
		case "avatar":
			// 给予玩家所有角色
			COMMAND_MANAGER.gmCmd.GMAddAllAvatar(c.AssignPlayer.PlayerId, 1, 0)
			c.SendSuccMessage(c.Executor, "已给予所有角色，指定UID：%v。", c.AssignPlayer.PlayerId)
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

// 杀死实体命令

var KillCommandController = &CommandController{
	Name:        "杀死实体",
	AliasList:   []string{"kill"},
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>杀死讨厌的实体</color>",
	UsageList: []string{
		"{alias} self 杀死自己",
		"{alias} monster <实体ID/all> 杀死怪物",
	},
	Perm: CommandPermNormal,
	Func: KillCommand,
}

func KillCommand(c *CommandContent) {
	var mode string   // 模式
	var param1 string // 参数

	c.Dynamic("string", func(param any) bool {
		// 模式
		mode = param.(string)
		return true
	}).Option("string", func(param any) bool {
		// 参数1
		param1 = param.(string)
		return true
	}).Execute(func() bool {
		switch mode {
		case "self":
			// 杀死自己
			COMMAND_MANAGER.gmCmd.GMKillSelf(c.AssignPlayer.PlayerId)
			c.SendSuccMessage(c.Executor, "已杀死自己，指定UID：%v。", c.AssignPlayer.PlayerId)
		case "monster":
			// 杀死怪物
			switch param1 {
			case "":
				// 怪物的话必须指定目标
				c.SetElse(func() {
					c.SendFailMessage(c.Executor, "参数不足，必须指定杀死的怪物。")
				})
				return false
			case "all":
				// 目标为全部怪物
				COMMAND_MANAGER.gmCmd.GMKillAllMonster(c.AssignPlayer.PlayerId)
				c.SendSuccMessage(c.Executor, "已杀死所有怪物，指定UID：%v。", c.AssignPlayer.PlayerId)
			default:
				// 实体id
				entityId, err := strconv.ParseUint(param1, 10, 32)
				if err != nil {
					return false
				}
				COMMAND_MANAGER.gmCmd.GMKillMonster(c.AssignPlayer.PlayerId, uint32(entityId))
				c.SendSuccMessage(c.Executor, "已杀死目标怪物，指定UID：%v，实体ID：%v。", c.AssignPlayer.PlayerId, entityId)
			}
		default:
			return false
		}
		return true
	})
}

// 管理任务命令

var QuestCommandController = &CommandController{
	Name:        "管理任务",
	AliasList:   []string{"quest"},
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>管理你的任务</color>",
	UsageList: []string{
		"{alias} <add/accept> <任务ID> 接受任务",
		"{alias} finish <任务ID/all> 完成任务",
	},
	Perm: CommandPermNormal,
	Func: QuestCommand,
}

func QuestCommand(c *CommandContent) {
	var mode string   // 模式
	var param1 string // 参数1

	c.Dynamic("string", func(param any) bool {
		// 模式
		mode = param.(string)
		return true
	}).Dynamic("string", func(param any) bool {
		// 参数1
		param1 = param.(string)
		return true
	}).Execute(func() bool {
		switch mode {
		case "add", "accept":
			// 任务id
			questId, err := strconv.ParseUint(param1, 10, 32)
			if err != nil {
				return false
			}
			// 添加指定任务
			// 接受指定任务 暂时与添加相同
			COMMAND_MANAGER.gmCmd.GMAddQuest(c.AssignPlayer.PlayerId, uint32(questId))
			c.SendSuccMessage(c.Executor, "已添加任务，指定UID：%v，任务ID：%v。", c.AssignPlayer.PlayerId, questId)
		case "finish":
			// 完成指定任务
			if param1 == "all" {
				// 强制完成当前所有任务
				COMMAND_MANAGER.gmCmd.GMForceFinishAllQuest(c.AssignPlayer.PlayerId)
				c.SendSuccMessage(c.Executor, "已完成当前全部任务，指定UID：%v。", c.AssignPlayer.PlayerId)
				return true
			}
			// 任务id
			questId, err := strconv.ParseUint(param1, 10, 32)
			if err != nil {
				return false
			}
			COMMAND_MANAGER.gmCmd.GMFinishQuest(c.AssignPlayer.PlayerId, uint32(questId))
			c.SendSuccMessage(c.Executor, "已完成玩家任务，指定UID：%v，任务ID：%v。", c.AssignPlayer.PlayerId, questId)
		default:
			return false
		}
		return true
	})
}

// 解锁锚点命令

var PointCommandController = &CommandController{
	Name:        "解锁锚点",
	AliasList:   []string{"point"},
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>解锁地图上的锚点</color>",
	UsageList: []string{
		"{alias} [场景ID] <锚点ID/all> 解锁锚点",
	},
	Perm: CommandPermNormal,
	Func: PointCommand,
}

func PointCommand(c *CommandContent) {
	var sceneId = c.AssignPlayer.SceneId // 场景id
	var param1 string                    // 参数1

	switch len(c.ParamList) {
	case 1:
		c.Dynamic("string", func(param any) bool {
			// 参数1
			param1 = param.(string)
			return true
		})
	case 2:
		c.Dynamic("uint32", func(param any) bool {
			// 场景id
			sceneId = param.(uint32)
			return true
		}).Dynamic("string", func(param any) bool {
			// 参数1
			param1 = param.(string)
			return true
		})
	}
	c.Execute(func() bool {
		if param1 == "all" {
			// 解锁当前场景所有锚点
			COMMAND_MANAGER.gmCmd.GMUnlockAllPoint(c.AssignPlayer.PlayerId, sceneId)
			c.SendSuccMessage(c.Executor, "已解锁所有锚点，指定UID：%v，场景ID：%v。", c.AssignPlayer.PlayerId, c.AssignPlayer.SceneId)
			return true
		}
		// 锚点id
		pointId, err := strconv.ParseUint(param1, 10, 32)
		if err != nil {
			return false
		}
		COMMAND_MANAGER.gmCmd.GMUnlockPoint(c.AssignPlayer.PlayerId, sceneId, uint32(pointId))
		c.SendSuccMessage(c.Executor, "已解锁锚点，指定UID：%v，场景ID：%v，锚点ID：%v。", c.AssignPlayer.PlayerId, c.AssignPlayer.SceneId, pointId)
		return true
	})
}

// xLua调试命令

var XLuaDebugCommandController = &CommandController{
	Name:        "xLua调试",
	AliasList:   []string{"xluadebug"},
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>开关xLua调试</color>",
	UsageList: []string{
		"{alias} 开关xLua调试",
	},
	Perm: CommandPermNormal,
	Func: XLuaDebugCommand,
}

func XLuaDebugCommand(c *CommandContent) {
	c.Execute(func() bool {
		// 主动开关客户端XLUA调试
		if !c.AssignPlayer.XLuaDebug {
			c.AssignPlayer.XLuaDebug = true
			c.SendSuccMessage(c.Executor, "已开启客户端XLUA调试，指定UID：%v。", c.AssignPlayer.PlayerId)
		} else {
			c.AssignPlayer.XLuaDebug = false
			c.SendSuccMessage(c.Executor, "已关闭客户端XLUA调试，指定UID：%v。", c.AssignPlayer.PlayerId)
		}
		return true
	})
}

// 七圣召唤测试命令

var GcgCommandController = &CommandController{
	Name:        "七圣召唤测试",
	AliasList:   []string{"gcgtest"},
	Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>测试七圣召唤</color>",
	UsageList: []string{
		"{alias} 测试七圣召唤",
	},
	Perm: CommandPermNormal,
	Func: GcgCommand,
}

func GcgCommand(c *CommandContent) {
	c.Execute(func() bool {
		// 开始七圣召唤对局
		GAME.GCGStartChallenge(c.AssignPlayer)
		c.SendSuccMessage(c.Executor, "已开始七圣召唤对局，指定UID：%v。", c.AssignPlayer.PlayerId)
		return true
	})
}
