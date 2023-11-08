package game

import (
	"fmt"
	"strconv"
	"strings"

	"hk4e/gdconf"
	"hk4e/pkg/logger"
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

// InitController 初始化命令控制器
func (c *CommandManager) InitController() {
	controllerList := []*CommandController{
		c.NewAssignCommandController(),
		c.NewHelpCommandController(),
		c.NewGotoCommandController(),
		c.NewJumpCommandController(),
		c.NewEquipCommandController(),
		c.NewItemCommandController(),
		c.NewAvatarCommandController(),
		c.NewGiveCommandController(),
		c.NewKillCommandController(),
		c.NewMonsterCommandController(),
		c.NewGadgetCommandController(),
		c.NewQuestCommandController(),
		c.NewPointCommandController(),
		c.NewWeatherCommandController(),
		c.NewClearCommandController(),
		c.NewDebugCommandController(),
		c.NewWudiCommandController(),
		c.NewEnergyCommandController(),
		c.NewStaminaCommandController(),
	}
	c.RegAllController(controllerList...)
}

// 指定命令

func (c *CommandManager) NewAssignCommandController() *CommandController {
	return &CommandController{
		Name:        "指定",
		AliasList:   []string{"assign"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>设置命令指定玩家</color>",
		UsageList: []string{
			"{alias} <目标UID> 指定某个玩家",
		},
		Perm: CommandPermGM,
		Func: c.AssignCommand,
	}
}

func (c *CommandManager) AssignCommand(content *CommandContent) bool {
	var assignUid uint32

	return content.Dynamic("uint32", func(param any) bool {
		value := param.(uint32)
		// 指定uid
		assignUid = value
		return true
	}).Execute(func() bool {
		// 设置命令指定uid
		content.Executor.CommandAssignUid = assignUid
		content.SendSuccMessage(content.Executor, "已指定玩家，指定UID：%v", assignUid)
		return true
	})
}

// 帮助命令

func (c *CommandManager) NewHelpCommandController() *CommandController {
	return &CommandController{
		Name:        "帮助",
		AliasList:   []string{"help"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>查看简要帮助信息</color>",
		UsageList: []string{
			"{alias} 查看简要帮助信息",
			"{alias} <序号/命令别名> 查看详细帮助信息",
		},
		Perm: CommandPermNormal,
		Func: c.HelpCommand,
	}
}

func (c *CommandManager) HelpCommand(content *CommandContent) bool {
	var controller *CommandController // 命令控制器
	var alias string                  // 别名

	return content.Option("string", func(param any) bool {
		value := param.(string)
		// 通过别名获取
		controller = c.commandControllerMap[value]
		if controller == nil {
			return false
		}
		alias = value
		return true
	}).Execute(func() bool {
		if alias == "" {
			// 显示简要帮助信息
			helpText := "<color=#66B2FF>================</color><color=#CCE5FF>/ 帮 助 /</color><color=#66B2FF>================</color>\n"
			commandCount := 0 // 权限足够的命令
			for _, controller := range c.commandControllerList {
				// 权限不足跳过
				if content.Executor.CmdPerm < uint8(controller.Perm) {
					continue
				}
				commandCount++
				// GM命令和普通命令区分颜色
				var permColor string
				switch controller.Perm {
				case CommandPermNormal:
					permColor = "#CCFFCC"
				case CommandPermGM:
					permColor = "#FF9999"
				}
				helpText += fmt.Sprintf("<color=%v>%v. %v</color> <color=#FFE5CC>-</color> %v\n", permColor, commandCount, controller.Name, strings.ReplaceAll(controller.Description, "{alias}", controller.AliasList[0]))
			}
			helpText += "\n<color=#FFFFCC>help</color> <color=#FFCCE5><命令别名></color> <color=#FF9999>能查看详细用法哦~</color>\n"
			helpText += "<color=#FF6347><></color> <color=#87CEFA>代表必填参数</color> <color=#FF6347>[]</color> <color=#87CEFA>代表可选参数</color> <color=#FF6347>/</color> <color=#87CEFA>代表或者</color>"
			content.SendMessage(content.Executor, helpText)
			return true
		}
		// 命令详细用法
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
		content.SendMessage(content.Executor, text)
		return true
	})
}

// 传送坐标命令

func (c *CommandManager) NewGotoCommandController() *CommandController {
	return &CommandController{
		Name:        "传送坐标",
		AliasList:   []string{"goto"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>传送到指定坐标</color>",
		UsageList: []string{
			"{alias} <坐标X> <坐标Y> <坐标Z> 传送至指定坐标",
		},
		Perm: CommandPermNormal,
		Func: c.GotoCommand,
	}
}

func (c *CommandManager) GotoCommand(content *CommandContent) bool {
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
	var sceneId = content.AssignPlayer.SceneId
	var posX, posY, posZ float64

	// 解析命令
	playerPos := GAME.GetPlayerPos(content.AssignPlayer)
	return content.Dynamic("string", func(param any) bool {
		// 坐标x
		value := param.(string)
		pos, ok := parseRelativePosFunc(value, playerPos.X)
		posX = pos
		return ok
	}).Dynamic("string", func(param any) bool {
		// 坐标y
		value := param.(string)
		pos, ok := parseRelativePosFunc(value, playerPos.Y)
		posY = pos
		return ok
	}).Dynamic("string", func(param any) bool {
		// 坐标z
		value := param.(string)
		pos, ok := parseRelativePosFunc(value, playerPos.Z)
		posZ = pos
		return ok
	}).Execute(func() bool {
		// 传送玩家至指定的位置
		c.gmCmd.GMTeleportPlayer(content.AssignPlayer.PlayerId, sceneId, posX, posY, posZ)
		// 发送消息给执行者
		content.SendSuccMessage(content.Executor, "已传送至指定位置，指定UID：%v，场景ID：%v，X：%.2f，Y：%.2f，Z：%.2f。", content.AssignPlayer.PlayerId, content.AssignPlayer.SceneId, posX, posY, posZ)
		return true
	})
}

// 传送场景命令

func (c *CommandManager) NewJumpCommandController() *CommandController {
	return &CommandController{
		Name:        "传送场景",
		AliasList:   []string{"jump"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>传送到至指定场景</color>",
		UsageList: []string{
			"{alias} <场景ID> 传送至指定场景",
		},
		Perm: CommandPermNormal,
		Func: c.JumpCommand,
	}
}

func (c *CommandManager) JumpCommand(content *CommandContent) bool {
	var sceneId uint32 // 场景id

	return content.Dynamic("uint32", func(param any) bool {
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
			logger.Error("get scene lua config is nil, sceneId: %v, uid: %v", sceneId, content.AssignPlayer.PlayerId)
		}
		// 传送玩家至指定的位置
		c.gmCmd.GMTeleportPlayer(content.AssignPlayer.PlayerId, sceneId, posX, posY, posZ)
		// 发送消息给执行者
		content.SendSuccMessage(content.Executor, "已传送至指定场景，指定UID：%v，场景ID：%v，X：%.2f，Y：%.2f，Z：%.2f。", content.AssignPlayer.PlayerId, content.AssignPlayer.SceneId, posX, posY, posZ)
		return true
	})
}

// 管理武器命令

func (c *CommandManager) NewEquipCommandController() *CommandController {
	return &CommandController{
		Name:        "管理武器",
		AliasList:   []string{"equip", "weapon"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>管理你的武器</color>",
		UsageList: []string{
			"{alias} add <武器ID/all> [武器等级] [突破等级] 添加武器",
		},
		Perm: CommandPermNormal,
		Func: c.EquipCommand,
	}
}

func (c *CommandManager) EquipCommand(content *CommandContent) bool {
	var mode string     // 模式
	var param1 string   // 参数1
	var level uint8 = 1 // 武器等级
	var promote uint8   // 突破等级

	return content.Dynamic("string", func(param any) bool {
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
				c.gmCmd.GMAddAllWeapon(content.AssignPlayer.PlayerId, 1, level, promote, 1)
				content.SendSuccMessage(content.Executor, "已给予所有武器，指定UID：%v，武器等级：%v，突破等级：%v。", content.AssignPlayer.PlayerId, level, promote)
				return true
			}
			// 物品id
			itemId, err := strconv.ParseUint(param1, 10, 32)
			if err != nil {
				return false
			}
			c.gmCmd.GMAddWeapon(content.AssignPlayer.PlayerId, uint32(itemId), 1, level, promote, 1)
			content.SendSuccMessage(content.Executor, "已给予武器，指定UID：%v，武器ID：%v，武器等级：%v，突破等级：%v。", content.AssignPlayer.PlayerId, itemId, level, promote)
		default:
			return false
		}
		return true
	})
}

// 管理物品命令

func (c *CommandManager) NewItemCommandController() *CommandController {
	return &CommandController{
		Name:        "管理物品",
		AliasList:   []string{"item"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>管理你的物品</color>",
		UsageList: []string{
			"{alias} add <物品ID/all> [数量] 添加物品",
			"{alias} clear 清除全部物品",
		},
		Perm: CommandPermNormal,
		Func: c.ItemCommand,
	}
}

func (c *CommandManager) ItemCommand(content *CommandContent) bool {
	var mode string      // 模式
	var param1 string    // 参数1
	var count uint32 = 1 // 数量

	return content.Dynamic("string", func(param any) bool {
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
				c.gmCmd.GMAddAllItem(content.AssignPlayer.PlayerId, count)
				content.SendSuccMessage(content.Executor, "已给予所有物品，指定UID：%v，数量：%v。", content.AssignPlayer.PlayerId, count)
				return true
			}
			// 物品id
			itemId, err := strconv.ParseUint(param1, 10, 32)
			if err != nil {
				return false
			}
			c.gmCmd.GMAddItem(content.AssignPlayer.PlayerId, uint32(itemId), count)
			content.SendSuccMessage(content.Executor, "已给予物品，指定UID：%v，物品ID：%v，数量：%v。", content.AssignPlayer.PlayerId, itemId, count)
		case "clear":
			c.gmCmd.GMClearItem(content.AssignPlayer.PlayerId)
			content.SendSuccMessage(content.Executor, "已清除全部物品，指定UID：%v。", content.AssignPlayer.PlayerId)
		default:
			return false
		}
		return true
	})
}

// 管理角色命令

func (c *CommandManager) NewAvatarCommandController() *CommandController {
	return &CommandController{
		Name:        "管理角色",
		AliasList:   []string{"avatar"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>管理你的角色</color>",
		UsageList: []string{
			"{alias} add <角色ID/all>",
		},
		Perm: CommandPermNormal,
		Func: c.AvatarCommand,
	}
}

func (c *CommandManager) AvatarCommand(content *CommandContent) bool {
	var mode string   // 模式
	var param1 string // 参数1

	return content.Dynamic("string", func(param any) bool {
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
				c.gmCmd.GMAddAllAvatar(content.AssignPlayer.PlayerId, 1, 0)
				content.SendSuccMessage(content.Executor, "已给予所有角色，指定UID：%v。", content.AssignPlayer.PlayerId)
				return true
			}
			// 角色id
			avatarId, err := strconv.ParseUint(param1, 10, 32)
			if err != nil {
				return false
			}
			c.gmCmd.GMAddAvatar(content.AssignPlayer.PlayerId, uint32(avatarId), 1, 0)
			content.SendSuccMessage(content.Executor, "已给予角色，指定UID：%v，角色ID：%v。", content.AssignPlayer.PlayerId, avatarId)
		default:
			return false
		}
		return true
	})
}

// 给予命令

func (c *CommandManager) NewGiveCommandController() *CommandController {
	return &CommandController{
		Name:        "给予物品",
		AliasList:   []string{"give"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>获得全部物品</color>",
		UsageList: []string{
			"模式：ID / item (所有物品) / weapon (所有武器) / reliquary (所有圣遗物) / avatar (所有角色) / costume (所有时装) / flycloak (所有风之翼) / all (全部)\n不要加上括号内的中文！！",
			"{alias} <模式> [数量] 给予指定物品",
			"数量仅物品、武器、圣遗物可用",
		},
		Perm: CommandPermNormal,
		Func: c.GiveCommand,
	}
}

func (c *CommandManager) GiveCommand(content *CommandContent) bool {
	// 给予物品
	var mode string      // 给予的模式
	var itemId uint32    // 物品id
	var count uint32 = 1 // 数量

	return content.Dynamic("string", func(param any) bool {
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
				c.gmCmd.GMAddItem(content.AssignPlayer.PlayerId, itemId, count)
				content.SendSuccMessage(content.Executor, "已给予物品，指定UID：%v，ID：%v，数量：%v。", content.AssignPlayer.PlayerId, itemId, count)
				return true
			}
			// 判断是否为武器
			_, ok = GAME.GetAllWeaponDataConfig()[int32(itemId)]
			if ok {
				// 给予玩家武器
				c.gmCmd.GMAddWeapon(content.AssignPlayer.PlayerId, itemId, count, 1, 0, 1)
				content.SendSuccMessage(content.Executor, "已给予武器，指定UID：%v，ID：%v，数量：%v。", content.AssignPlayer.PlayerId, itemId, count)
				return true
			}
			// 判断是否为圣遗物
			_, ok = GAME.GetAllReliquaryDataConfig()[int32(itemId)]
			if ok {
				// 给予玩家圣遗物
				c.gmCmd.GMAddReliquary(content.AssignPlayer.PlayerId, itemId, count)
				content.SendSuccMessage(content.Executor, "已给予圣遗物，指定UID：%v，ID：%v，数量：%v。", content.AssignPlayer.PlayerId, itemId, count)
				return true
			}
			// 判断是否为角色
			_, ok = GAME.GetAllAvatarDataConfig()[int32(itemId)]
			if ok {
				// 给予玩家角色
				c.gmCmd.GMAddAvatar(content.AssignPlayer.PlayerId, itemId, 1, 0)
				content.SendSuccMessage(content.Executor, "已给予角色，指定UID：%v，ID：%v。", content.AssignPlayer.PlayerId, itemId)
				return true
			}
			// 判断是否为时装
			if gdconf.GetAvatarCostumeDataById(int32(itemId)) != nil {
				// 给予玩家时装
				c.gmCmd.GMAddCostume(content.AssignPlayer.PlayerId, itemId)
				content.SendSuccMessage(content.Executor, "已给予时装，指定UID：%v，ID：%v。", content.AssignPlayer.PlayerId, itemId)
				return true
			}
			// 判断是否为风之翼
			if gdconf.GetAvatarFlycloakDataById(int32(itemId)) != nil {
				// 给予玩家风之翼
				c.gmCmd.GMAddFlycloak(content.AssignPlayer.PlayerId, itemId)
				content.SendSuccMessage(content.Executor, "已给予风之翼，指定UID：%v，ID：%v。", content.AssignPlayer.PlayerId, itemId)
				return true
			}
			// 都执行到这里那肯定是都不匹配
			content.SetElse(func() {
				// 物品id不存在
				content.SendFailMessage(content.Executor, "不存在的物品，ID：%v。", itemId)
			})
			return false
		case "item":
			// 给予玩家所有物品
			c.gmCmd.GMAddAllItem(content.AssignPlayer.PlayerId, count)
			content.SendSuccMessage(content.Executor, "已给予所有物品，指定UID：%v，数量：%v。", content.AssignPlayer.PlayerId, count)
		case "weapon":
			// 给予玩家所有武器
			c.gmCmd.GMAddAllWeapon(content.AssignPlayer.PlayerId, count, 1, 0, 1)
			content.SendSuccMessage(content.Executor, "已给予所有武器，指定UID：%v，数量：%v。", content.AssignPlayer.PlayerId, count)
		case "reliquary":
			// 给予玩家所有圣遗物
			c.gmCmd.GMAddAllReliquary(content.AssignPlayer.PlayerId, count)
			content.SendSuccMessage(content.Executor, "已给予所有圣遗物，指定UID：%v，数量：%v。", content.AssignPlayer.PlayerId, count)
		case "avatar":
			// 给予玩家所有角色
			c.gmCmd.GMAddAllAvatar(content.AssignPlayer.PlayerId, 1, 0)
			content.SendSuccMessage(content.Executor, "已给予所有角色，指定UID：%v。", content.AssignPlayer.PlayerId)
		case "costume":
			// 给予玩家所有时装
			c.gmCmd.GMAddAllCostume(content.AssignPlayer.PlayerId)
			content.SendSuccMessage(content.Executor, "已给予所有时装，指定UID：%v。", content.AssignPlayer.PlayerId)
		case "flycloak":
			// 给予玩家所有风之翼
			c.gmCmd.GMAddAllFlycloak(content.AssignPlayer.PlayerId)
			content.SendSuccMessage(content.Executor, "已给予所有风之翼，指定UID：%v。", content.AssignPlayer.PlayerId)
		case "all":
			// 给予玩家所有内容
			c.gmCmd.GMAddAll(content.AssignPlayer.PlayerId)
			content.SendSuccMessage(content.Executor, "已给予所有内容，指定UID：%v。", content.AssignPlayer.PlayerId)
		default:
			return false
		}
		return true
	})
}

// 杀死实体命令

func (c *CommandManager) NewKillCommandController() *CommandController {
	return &CommandController{
		Name:        "杀死实体",
		AliasList:   []string{"kill"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>杀死讨厌的实体</color>",
		UsageList: []string{
			"{alias} self 杀死自己",
			"{alias} monster <实体ID/all> 杀死怪物",
		},
		Perm: CommandPermNormal,
		Func: c.KillCommand,
	}
}

func (c *CommandManager) KillCommand(content *CommandContent) bool {
	var mode string   // 模式
	var param1 string // 参数

	return content.Dynamic("string", func(param any) bool {
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
			c.gmCmd.GMKillSelf(content.AssignPlayer.PlayerId)
			content.SendSuccMessage(content.Executor, "已杀死自己，指定UID：%v。", content.AssignPlayer.PlayerId)
		case "monster":
			// 杀死怪物
			switch param1 {
			case "":
				// 怪物的话必须指定目标
				content.SetElse(func() {
					content.SendFailMessage(content.Executor, "参数不足，必须指定杀死的怪物。")
				})
				return false
			case "all":
				// 目标为全部怪物
				c.gmCmd.GMKillAllMonster(content.AssignPlayer.PlayerId)
				content.SendSuccMessage(content.Executor, "已杀死所有怪物，指定UID：%v。", content.AssignPlayer.PlayerId)
			default:
				// 实体id
				entityId, err := strconv.ParseUint(param1, 10, 32)
				if err != nil {
					return false
				}
				c.gmCmd.GMKillMonster(content.AssignPlayer.PlayerId, uint32(entityId))
				content.SendSuccMessage(content.Executor, "已杀死目标怪物，指定UID：%v，实体ID：%v。", content.AssignPlayer.PlayerId, entityId)
			}
		default:
			return false
		}
		return true
	})
}

// 生成怪物命令

func (c *CommandManager) NewMonsterCommandController() *CommandController {
	return &CommandController{
		Name:        "生成怪物",
		AliasList:   []string{"monster"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>生成原魔</color>",
		UsageList: []string{
			"{alias} <怪物ID> [数量] [等级] [姿势 (暂时无效)] [坐标X] [坐标Y] [坐标Z] 生成怪物",
		},
		Perm: CommandPermNormal,
		Func: c.MonsterCommand,
	}
}

func (c *CommandManager) MonsterCommand(content *CommandContent) bool {
	var monsterId uint32 // 怪物id
	var count uint32 = 1 // 数量
	var level uint8 = 1  // 等级
	// var pose uint32      // 姿势
	pos := GAME.GetPlayerPos(content.AssignPlayer)
	var posX = pos.X // 坐标x
	var posY = pos.Y // 坐标y
	var posZ = pos.Z // 坐标z

	return content.Dynamic("uint32", func(param any) bool {
		monsterId = param.(uint32)
		return true
	}).Option("uint32", func(param any) bool {
		count = param.(uint32)
		return true
	}).Option("uint8", func(param any) bool {
		level = param.(uint8)
		return true
	}).Option("uint32", func(param any) bool {
		// pose = param.(uint32)
		return true
	}).Option("float64", func(param any) bool {
		posX = param.(float64)
		return true
	}).Option("float64", func(param any) bool {
		posY = param.(float64)
		return true
	}).Option("float64", func(param any) bool {
		posZ = param.(float64)
		return true
	}).Execute(func() bool {
		c.gmCmd.GMCreateMonster(content.AssignPlayer.PlayerId, monsterId, posX, posY, posZ, count, level)
		return true
	})
}

// 生成物件命令

func (c *CommandManager) NewGadgetCommandController() *CommandController {
	return &CommandController{
		Name:        "生成物件",
		AliasList:   []string{"gadget"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>生成物件</color>",
		UsageList: []string{
			"{alias} <物件ID> [数量] 附近生成物件",
		},
		Perm: CommandPermNormal,
		Func: c.GadgetCommand,
	}
}

func (c *CommandManager) GadgetCommand(content *CommandContent) bool {
	var gadgetId uint32  // 物件id
	var count uint32 = 1 // 数量

	return content.Dynamic("uint32", func(param any) bool {
		gadgetId = param.(uint32)
		return true
	}).Option("uint32", func(param any) bool {
		count = param.(uint32)
		return true
	}).Execute(func() bool {
		c.gmCmd.GMCreateGadget(content.AssignPlayer.PlayerId, gadgetId, count)
		return true
	})
}

// 管理任务命令

func (c *CommandManager) NewQuestCommandController() *CommandController {
	return &CommandController{
		Name:        "管理任务",
		AliasList:   []string{"quest"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>管理你的任务</color>",
		UsageList: []string{
			"{alias} <add/accept> <任务ID> 接受任务",
			"{alias} finish <任务ID/all> 完成任务",
			"{alias} clear all 清除全部任务",
		},
		Perm: CommandPermNormal,
		Func: c.QuestCommand,
	}
}

func (c *CommandManager) QuestCommand(content *CommandContent) bool {
	var mode string   // 模式
	var param1 string // 参数1

	return content.Dynamic("string", func(param any) bool {
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
			c.gmCmd.GMAddQuest(content.AssignPlayer.PlayerId, uint32(questId))
			content.SendSuccMessage(content.Executor, "已添加任务，指定UID：%v，任务ID：%v。", content.AssignPlayer.PlayerId, questId)
		case "finish":
			// 完成指定任务
			if param1 == "all" {
				// 强制完成当前所有任务
				c.gmCmd.GMForceFinishAllQuest(content.AssignPlayer.PlayerId)
				content.SendSuccMessage(content.Executor, "已完成当前全部任务，指定UID：%v。", content.AssignPlayer.PlayerId)
				return true
			}
			// 任务id
			questId, err := strconv.ParseUint(param1, 10, 32)
			if err != nil {
				return false
			}
			c.gmCmd.GMFinishQuest(content.AssignPlayer.PlayerId, uint32(questId))
			content.SendSuccMessage(content.Executor, "已完成玩家任务，指定UID：%v，任务ID：%v。", content.AssignPlayer.PlayerId, questId)
		case "clear":
			c.gmCmd.GMClearQuest(content.AssignPlayer.PlayerId)
			content.SendSuccMessage(content.Executor, "已清除全部任务，指定UID：%v。", content.AssignPlayer.PlayerId)
		default:
			return false
		}
		return true
	})
}

// 解锁锚点命令

func (c *CommandManager) NewPointCommandController() *CommandController {
	return &CommandController{
		Name:        "解锁锚点",
		AliasList:   []string{"point"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>解锁地图上的锚点</color>",
		UsageList: []string{
			"{alias} [场景ID] <锚点ID/all> 解锁锚点",
		},
		Perm: CommandPermNormal,
		Func: c.PointCommand,
	}
}

func (c *CommandManager) PointCommand(content *CommandContent) bool {
	var sceneId = content.AssignPlayer.SceneId // 场景id
	var param1 string                          // 参数1

	return content.Option("uint32", func(param any) bool {
		// 场景id
		sceneId = param.(uint32)
		return true
	}).Dynamic("string", func(param any) bool {
		// 参数1
		param1 = param.(string)
		return true
	}).Execute(func() bool {
		if param1 == "all" {
			// 解锁当前场景所有锚点
			c.gmCmd.GMUnlockAllPoint(content.AssignPlayer.PlayerId, sceneId)
			content.SendSuccMessage(content.Executor, "已解锁所有锚点，指定UID：%v，场景ID：%v。", content.AssignPlayer.PlayerId, content.AssignPlayer.SceneId)
			return true
		}
		// 锚点id
		pointId, err := strconv.ParseUint(param1, 10, 32)
		if err != nil {
			return false
		}
		c.gmCmd.GMUnlockPoint(content.AssignPlayer.PlayerId, sceneId, uint32(pointId))
		content.SendSuccMessage(content.Executor, "已解锁锚点，指定UID：%v，场景ID：%v，锚点ID：%v。", content.AssignPlayer.PlayerId, content.AssignPlayer.SceneId, pointId)
		return true
	})
}

// 更改天气命令

func (c *CommandManager) NewWeatherCommandController() *CommandController {
	return &CommandController{
		Name:        "更改天气",
		AliasList:   []string{"weather"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>更改天气</color>",
		UsageList: []string{
			"{alias} [天气区域ID] <气象类型> 更改天气",
		},
		Perm: CommandPermNormal,
		Func: c.WeatherCommand,
	}
}

func (c *CommandManager) WeatherCommand(content *CommandContent) bool {
	var weatherAreaId = content.AssignPlayer.WeatherInfo.WeatherAreaId // 天气区域id
	var climateType uint32                                             // 气象类型

	return content.Option("uint32", func(param any) bool {
		// 天气id
		weatherAreaId = param.(uint32)
		return true
	}).Dynamic("uint32", func(param any) bool {
		// 气象类型
		climateType = param.(uint32)
		return true
	}).Execute(func() bool {
		// 设置天气
		c.gmCmd.GMSetWeather(content.AssignPlayer.PlayerId, weatherAreaId, climateType)
		content.SendSuccMessage(content.Executor, "已更改天气，指定UID：%v，天气区域ID：%v，气象类型：%v。", content.AssignPlayer.PlayerId, weatherAreaId, climateType)
		return true
	})
}

// 清除命令

func (c *CommandManager) NewClearCommandController() *CommandController {
	return &CommandController{
		Name:        "清除",
		AliasList:   []string{"clear"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>清除</color>",
		UsageList: []string{
			"{alias} all 清除玩家数据",
		},
		Perm: CommandPermNormal,
		Func: c.ClearCommand,
	}
}

func (c *CommandManager) ClearCommand(content *CommandContent) bool {
	var mode string // 模式

	return content.Dynamic("string", func(param any) bool {
		// 模式
		mode = param.(string)
		return true
	}).Execute(func() bool {
		switch mode {
		case "all":
			c.gmCmd.GMClearPlayer(content.AssignPlayer.PlayerId)
			content.SendSuccMessage(content.Executor, "已清除玩家数据，指定UID：%v。", content.AssignPlayer.PlayerId)
			return true
		default:
			return false
		}
	})
}

// 调试命令

func (c *CommandManager) NewDebugCommandController() *CommandController {
	return &CommandController{
		Name:        "调试",
		AliasList:   []string{"debug"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>调试</color>",
		UsageList: []string{
			"{alias} freemode 自由探索模式",
			"{alias} openstate 解锁全部功能",
			"{alias} clearworld 清除大世界数据",
			"{alias} notsave 本次离线回档",
			"{alias} xluaswitch 开关xLua",
			"{alias} gcgtest 七圣召唤测试",
		},
		Perm: CommandPermNormal,
		Func: c.DebugCommand,
	}
}

func (c *CommandManager) DebugCommand(content *CommandContent) bool {
	var mode string // 模式

	return content.Dynamic("string", func(param any) bool {
		// 模式
		mode = param.(string)
		return true
	}).Execute(func() bool {
		switch mode {
		case "freemode":
			c.gmCmd.GMFreeMode(content.AssignPlayer.PlayerId)
			content.SendSuccMessage(content.Executor, "已开启自由探索模式，指定UID：%v。", content.AssignPlayer.PlayerId)
			return true
		case "openstate":
			c.gmCmd.GMUnlockAllOpenState(content.AssignPlayer.PlayerId)
			content.SendSuccMessage(content.Executor, "已解锁全部功能，指定UID：%v。", content.AssignPlayer.PlayerId)
			return true
		case "clearworld":
			c.gmCmd.GMClearWorld(content.AssignPlayer.PlayerId)
			content.SendSuccMessage(content.Executor, "已清除大世界数据，指定UID：%v。", content.AssignPlayer.PlayerId)
			return true
		case "notsave":
			c.gmCmd.GMNotSave(content.AssignPlayer.PlayerId)
			content.SendSuccMessage(content.Executor, "已设置本次离线回档，指定UID：%v。", content.AssignPlayer.PlayerId)
			return true
		case "xluaswitch":
			if !content.AssignPlayer.XLuaDebug {
				content.AssignPlayer.XLuaDebug = true
				content.SendSuccMessage(content.Executor, "已开启客户端XLUA调试，指定UID：%v。", content.AssignPlayer.PlayerId)
			} else {
				content.AssignPlayer.XLuaDebug = false
				content.SendSuccMessage(content.Executor, "已关闭客户端XLUA调试，指定UID：%v。", content.AssignPlayer.PlayerId)
			}
			return true
		case "gcgtest":
			GAME.GCGStartChallenge(content.AssignPlayer)
			content.SendSuccMessage(content.Executor, "已开始七圣召唤对局，指定UID：%v。", content.AssignPlayer.PlayerId)
			return true
		default:
			return false
		}
	})
}

func (c *CommandManager) NewWudiCommandController() *CommandController {
	return &CommandController{
		Name:        "开关无敌",
		AliasList:   []string{"wudi"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>开关无敌</color>",
		UsageList: []string{
			"{alias} global avatar <on/off> 开关玩家无敌",
		},
		Perm: CommandPermNormal,
		Func: c.WudiCommand,
	}
}

func (c *CommandManager) WudiCommand(content *CommandContent) bool {
	var mode1 string  // 模式
	var mode2 string  // 模式
	var param1 string // 参数

	return content.Dynamic("string", func(param any) bool {
		// 模式
		mode1 = param.(string)
		return true
	}).Dynamic("string", func(param any) bool {
		// 模式
		mode2 = param.(string)
		return true
	}).Option("string", func(param any) bool {
		// 参数1
		param1 = param.(string)
		return true
	}).Execute(func() bool {
		switch mode1 {
		case "global":
			switch mode2 {
			case "avatar":
				switch param1 {
				case "on":
					c.gmCmd.GMSetPlayerWuDi(content.AssignPlayer.PlayerId, true)
					content.SendSuccMessage(content.Executor, "已开启玩家无敌，指定UID：%v。", content.AssignPlayer.PlayerId)
					return true
				case "off":
					c.gmCmd.GMSetPlayerWuDi(content.AssignPlayer.PlayerId, false)
					content.SendSuccMessage(content.Executor, "已关闭玩家无敌，指定UID：%v。", content.AssignPlayer.PlayerId)
					return true
				default:
					return false
				}
			default:
				return false
			}
		default:
			return false
		}
	})
}

func (c *CommandManager) NewEnergyCommandController() *CommandController {
	return &CommandController{
		Name:        "元素能量",
		AliasList:   []string{"energy"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>元素能量</color>",
		UsageList: []string{
			"{alias} infinite <on/off> 开关无限元素爆发",
		},
		Perm: CommandPermNormal,
		Func: c.EnergyCommand,
	}
}

func (c *CommandManager) EnergyCommand(content *CommandContent) bool {
	var mode1 string  // 模式
	var param1 string // 参数

	return content.Dynamic("string", func(param any) bool {
		// 模式
		mode1 = param.(string)
		return true
	}).Option("string", func(param any) bool {
		// 参数1
		param1 = param.(string)
		return true
	}).Execute(func() bool {
		switch mode1 {
		case "infinite":
			switch param1 {
			case "on":
				c.gmCmd.GMSetPlayerEnergyInf(content.AssignPlayer.PlayerId, true)
				content.SendSuccMessage(content.Executor, "已开启无限元素爆发，指定UID：%v。", content.AssignPlayer.PlayerId)
				return true
			case "off":
				c.gmCmd.GMSetPlayerEnergyInf(content.AssignPlayer.PlayerId, false)
				content.SendSuccMessage(content.Executor, "已关闭无限元素爆发，指定UID：%v。", content.AssignPlayer.PlayerId)
				return true
			default:
				return false
			}
		default:
			return false
		}
	})
}

func (c *CommandManager) NewStaminaCommandController() *CommandController {
	return &CommandController{
		Name:        "无限耐力",
		AliasList:   []string{"stamina"},
		Description: "<color=#FFFFCC>{alias}</color> <color=#FFCC99>无限耐力</color>",
		UsageList: []string{
			"{alias} infinite <on/off> 开关无限耐力",
		},
		Perm: CommandPermNormal,
		Func: c.StaminaCommand,
	}
}

func (c *CommandManager) StaminaCommand(content *CommandContent) bool {
	var mode1 string  // 模式
	var param1 string // 参数

	return content.Dynamic("string", func(param any) bool {
		// 模式
		mode1 = param.(string)
		return true
	}).Option("string", func(param any) bool {
		// 参数1
		param1 = param.(string)
		return true
	}).Execute(func() bool {
		switch mode1 {
		case "infinite":
			switch param1 {
			case "on":
				c.gmCmd.GMSetPlayerStaminaInf(content.AssignPlayer.PlayerId, true)
				content.SendSuccMessage(content.Executor, "已开启无限耐力，指定UID：%v。", content.AssignPlayer.PlayerId)
				return true
			case "off":
				c.gmCmd.GMSetPlayerStaminaInf(content.AssignPlayer.PlayerId, false)
				content.SendSuccMessage(content.Executor, "已关闭无限耐力，指定UID：%v。", content.AssignPlayer.PlayerId)
				return true
			default:
				return false
			}
		default:
			return false
		}
	})
}
