package game

import (
	"strconv"
	"strings"

	"hk4e/gdconf"

	"hk4e/gs/model"
)

// 玩家游戏内GM命令格式解析模块

// 以后不妨考虑改成米哈游版本的GM命令格式

func (c *CommandManager) GotoCommand(cmd *CommandMessage) {
	split := strings.Split(cmd.Text, " ")
	if len(split) != 4 {
		return
	}
	if split[0] != "goto" {
		return
	}
	x, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
		return
	}
	y, err := strconv.ParseFloat(split[2], 64)
	if err != nil {
		return
	}
	z, err := strconv.ParseFloat(split[3], 64)
	if err != nil {
		return
	}
	c.gmCmd.GMTeleportPlayer(cmd.Executor.PlayerId, cmd.Executor.SceneId, x, y, z)
}

// HelpCommand 帮助命令
func (c *CommandManager) HelpCommand(cmd *CommandMessage) {
	c.SendMessage(cmd.Executor,
		"========== 帮助 / Help ==========\n\n"+
			"传送：tp [--u <UID>] [--s <场景ID>] {--t <目标UID> | --x <坐标X> | --y <坐标Y> | --z <坐标Z>}\n\n"+
			"给予：give [--u <UID>] [--c <数量>] --i <ID / 物品 / 武器 / 圣遗物 / 角色 / 时装 / 风之翼 / 全部>\n\n"+
			"任务：quest [--u <UID>] [--q <任务ID>] --i <添加/完成/完成全部>\n\n"+
			"解锁所有锚点：unlock\n",
	)
}

// TeleportCommand 传送玩家命令
// tp [--u <uid>] [--s <sceneId>] {--t <targetUid> --x <posX> | --y <posY> | --z <posZ>}
func (c *CommandManager) TeleportCommand(cmd *CommandMessage) {
	// 执行者如果不是玩家则必须输入UID
	player := cmd.Executor

	// 判断是否填写必备参数
	// 目前传送的必备参数是任意包含一个就行
	if cmd.Args["t"] == "" && cmd.Args["x"] == "" && cmd.Args["y"] == "" && cmd.Args["z"] == "" {
		c.SendMessage(cmd.Executor, "参数不足，正确用法：%v [--u <UID>] [--s <场景ID>] {--t <目标UID> | --x <坐标X> | --y <坐标Y> | --z <坐标Z>}", cmd.Name)
		return
	}
	// 输入了目标UID则不能指定坐标或场景ID
	if cmd.Args["t"] != "" && (cmd.Args["x"] != "" || cmd.Args["y"] != "" || cmd.Args["z"] != "" || cmd.Args["s"] != "") {
		c.SendMessage(cmd.Executor, "你已指定目标玩家，无法指定传送位置。")
		return
	}

	// 初始值
	var target *model.Player  // 目标
	targetUid := uint32(0)    // 目标玩家uid
	sceneId := player.SceneId // 场景Id
	pos := &model.Vector{
		X: player.Pos.X,
		Y: player.Pos.Y,
		Z: player.Pos.Z,
	} // 坐标初始值为玩家当前所在位置

	// 选择每个参数
	for k, v := range cmd.Args {
		var err error

		switch k {
		case "u":
			var uid uint64
			if uid, err = strconv.ParseUint(v, 10, 32); err == nil {
				// 判断目标用户是否在线
				if user := USER_MANAGER.GetOnlineUser(uint32(uid)); user != nil {
					player = user
					// 防止覆盖用户指定过的sceneId
					if player.SceneId != sceneId {
						sceneId = player.SceneId
					}
				} else {
					c.SendMessage(cmd.Executor, "玩家不在线，UID：%v。", v)
					return
				}
			}
		case "s":
			var sid uint64
			if sid, err = strconv.ParseUint(v, 10, 32); err == nil {
				sceneId = uint32(sid)
			}
		case "t":
			var uid uint64
			if uid, err = strconv.ParseUint(v, 10, 32); err == nil {
				// 判断目标用户是否在线
				user := USER_MANAGER.GetOnlineUser(uint32(uid))
				if user == nil {
					// 目标玩家属于非本地玩家
					if !USER_MANAGER.GetRemoteUserOnlineState(uint32(uid)) {
						// 全服不存在该在线玩家
						c.SendMessage(cmd.Executor, "目标玩家不在线，UID：%v。", v)
						return
					}
				}
				target = user
				targetUid = uint32(uid)
			}
		case "x":
			// 玩家此时的位置X
			var nowX float64
			// 如果以 ~ 开头则 此时位置加 ~ 后的数
			if strings.HasPrefix(v, "~") {
				v = v[1:]           // 去除 ~
				nowX = player.Pos.X // 先记录
			}
			// 为空代表用户只输入 ~ 获取为玩家当前位置
			if v != "" {
				var x float64
				if x, err = strconv.ParseFloat(v, 64); err == nil {
					pos.X = x + nowX // 如果不以 ~ 开头则加 0
				}
			}
		case "y":
			// 玩家此时的位置Z
			var nowY float64
			// 如果以 ~ 开头则 此时位置加 ~ 后的数
			if strings.HasPrefix(v, "~") {
				v = v[1:]           // 去除 ~
				nowY = player.Pos.Y // 先记录
			}
			// 为空代表用户只输入 ~ 获取为玩家当前位置
			if v != "" {
				var y float64
				if y, err = strconv.ParseFloat(v, 64); err == nil {
					pos.Y = y + nowY
				}
			}
		case "z":
			// 玩家此时的位置Z
			var nowZ float64
			// 如果以 ~ 开头则 此时位置加 ~ 后的数
			if strings.HasPrefix(v, "~") {
				v = v[1:]           // 去除 ~
				nowZ = player.Pos.Z // 先记录
			}
			// 为空代表用户只输入 ~ 获取为玩家当前位置
			if v != "" {
				var z float64
				if z, err = strconv.ParseFloat(v, 64); err == nil {
					pos.Z = z + nowZ
				}
			}
		default:
			c.SendMessage(cmd.Executor, "参数 --%v 冗余。", k)
			return
		}

		// 解析错误的话应该是参数类型问题
		if err != nil {
			c.SendMessage(cmd.Executor, "参数 --%v 有误，类型错误。", k)
			return
		}
	}

	// 玩家是否指定目标UID
	if cmd.Args["t"] != "" {
		// 如果玩家不与目标玩家同一世界或不同服务器
		if target == nil || player.WorldId != target.WorldId {
			// 请求进入目标玩家世界
			GAME.PlayerApplyEnterWorld(player, targetUid)
			// 发送消息给执行者
			c.SendMessage(cmd.Executor, "已将玩家 UID：%v 请求加入目标玩家 UID：%v 的世界。", player.PlayerId, targetUid)
		} else {
			// 传送玩家至目标玩家的位置
			c.gmCmd.GMTeleportPlayer(player.PlayerId, target.SceneId, target.Pos.X, target.Pos.Y, target.Pos.Z)
			// 发送消息给执行者
			c.SendMessage(cmd.Executor, "已将玩家 UID：%v 传送至 目标玩家 UID：%v。", player.PlayerId, targetUid)
		}
	} else {
		// 传送玩家至指定的位置
		c.gmCmd.GMTeleportPlayer(player.PlayerId, sceneId, pos.X, pos.Y, pos.Z)
		// 发送消息给执行者
		c.SendMessage(cmd.Executor, "已将玩家 UID：%v 传送至 场景：%v, X：%.2f, Y：%.2f, Z：%.2f。", player.PlayerId, sceneId, pos.X, pos.Y, pos.Z)
	}

}

// GiveCommand 给予物品命令
// give [--u <userId>] [--c <count>] --i <id/item/weapon/reliquary/avatar/costume/flycloak/all>
func (c *CommandManager) GiveCommand(cmd *CommandMessage) {
	player := cmd.Executor

	// 判断是否填写必备参数
	if cmd.Args["i"] == "" {
		c.SendMessage(cmd.Executor, "参数不足，正确用法：%v [--u <UID>] [--c <数量>] --i <ID / 物品 / 武器 / 圣遗物 / 角色 / 时装 / 风之翼 / 全部>。", cmd.Name)
		return
	}

	// 初始值
	count := uint32(1) // 数量
	id := uint32(0)    // id
	// 给予物品的模式
	// once 单个 / all 所有物品
	// item 物品 / weapon 武器
	mode := "once"

	// 选择每个参数
	for k, v := range cmd.Args {
		var err error

		switch k {
		case "u":
			var uid uint64
			if uid, err = strconv.ParseUint(v, 10, 32); err == nil {
				// 判断目标用户是否在线
				if user := USER_MANAGER.GetOnlineUser(uint32(uid)); user != nil {
					player = user
				} else {
					c.SendMessage(cmd.Executor, "目标玩家不在线，UID：%v。", v)
					return
				}
			}
		case "c":
			var cnt uint64
			if cnt, err = strconv.ParseUint(v, 10, 32); err == nil {
				count = uint32(cnt)
			}
		case "i":
			switch v {
			case "item", "物品", "weapon", "武器", "reliquary", "圣遗物", "avatar", "角色", "costume", "时装", "flycloak", "风之翼", "all", "全部":
				// 将模式修改为参数的值
				mode = v
			default:
				var tempId uint64
				if tempId, err = strconv.ParseUint(v, 10, 32); err != nil {
					c.SendMessage(cmd.Executor, "参数 --%v 有误，允许内容：<ID / 物品 / 武器 / 圣遗物 / 角色 / 时装 / 风之翼 / 内容>。", k)
					return
				}
				id = uint32(tempId)
			}
		default:
			c.SendMessage(cmd.Executor, "参数 --%v 冗余。", k)
			return
		}

		// 解析错误的话应该是参数类型问题
		if err != nil {
			c.SendMessage(cmd.Executor, "参数 --%v 有误，类型错误。", k)
			return
		}
	}

	switch mode {
	case "once":
		// 判断是否为物品
		_, ok := GAME.GetAllItemDataConfig()[int32(id)]
		if ok {
			// 给予玩家物品
			c.gmCmd.GMAddItem(player.PlayerId, id, count)
			c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 物品ID：%v 数量：%v。", player.PlayerId, id, count)
			return
		}
		// 判断是否为武器
		_, ok = GAME.GetAllWeaponDataConfig()[int32(id)]
		if ok {
			// 给予玩家武器
			c.gmCmd.GMAddWeapon(player.PlayerId, id, count)
			c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 武器 物品ID：%v 数量：%v。", player.PlayerId, id, count)
			return

		}
		// 判断是否为圣遗物
		_, ok = GAME.GetAllReliquaryDataConfig()[int32(id)]
		if ok {
			// 给予玩家圣遗物
			c.gmCmd.GMAddReliquary(player.PlayerId, id, count)
			c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 圣遗物 物品ID：%v 数量：%v。", player.PlayerId, id, count)
			return

		}
		// 判断是否为角色
		_, ok = GAME.GetAllAvatarDataConfig()[int32(id)]
		if ok {
			// 给予玩家角色
			c.gmCmd.GMAddAvatar(player.PlayerId, id)
			c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 角色ID：%v 数量：%v。", player.PlayerId, id, count)
			return
		}
		// 判断是否为时装
		if gdconf.GetAvatarCostumeDataById(int32(id)) != nil {
			// 给予玩家时装
			c.gmCmd.GMAddCostume(player.PlayerId, id)
			c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 时装ID：%v 数量：%v。", player.PlayerId, id, count)
			return
		}
		// 判断是否为风之翼
		if gdconf.GetAvatarFlycloakDataById(int32(id)) != nil {
			// 给予玩家风之翼
			c.gmCmd.GMAddFlycloak(player.PlayerId, id)
			c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 风之翼ID：%v 数量：%v。", player.PlayerId, id, count)
			return
		}
		// 都执行到这里那肯定是都不匹配
		c.SendMessage(cmd.Executor, "ID：%v 不存在。", id)
	case "item", "物品":
		// 给予玩家所有物品
		c.gmCmd.GMAddAllItem(player.PlayerId, count)
		c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 所有物品 数量：%v。", player.PlayerId, count)
	case "weapon", "武器":
		// 给予玩家所有武器
		c.gmCmd.GMAddAllWeapon(player.PlayerId, count)
		c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 所有武器 数量：%v。", player.PlayerId, count)
	case "reliquary", "圣遗物":
		// 给予玩家所有圣遗物
		c.gmCmd.GMAddAllReliquary(player.PlayerId, count)
		c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 所有圣遗物 数量：%v。", player.PlayerId, count)
	case "avatar", "角色":
		// 给予玩家所有角色
		c.gmCmd.GMAddAllAvatar(player.PlayerId)
		c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 所有角色。", player.PlayerId)
	case "costume", "时装":
		// 给予玩家所有时装
		c.gmCmd.GMAddAllCostume(player.PlayerId)
		c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 所有时装。", player.PlayerId)
	case "flycloak", "风之翼":
		// 给予玩家所有风之翼
		c.gmCmd.GMAddAllFlycloak(player.PlayerId)
		c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 所有风之翼。", player.PlayerId)
	case "all", "全部":
		// 给予玩家所有内容
		c.gmCmd.GMAddAll(player.PlayerId, count) // TODO 武器额外获取数量
		c.SendMessage(cmd.Executor, "已给予玩家 UID：%v, 所有内容。", player.PlayerId)
	}
}

// GcgCommand Gcg测试命令
func (c *CommandManager) GcgCommand(cmd *CommandMessage) {
	player := cmd.Executor
	GAME.GCGStartChallenge(player)
	c.SendMessage(cmd.Executor, "收到命令")
}

// QuestCommand 任务控制命令
// quest [--u <userId>] [--q <questId>] --i <add/finish/finishall>
func (c *CommandManager) QuestCommand(cmd *CommandMessage) {
	player := cmd.Executor

	// 判断是否填写必备参数
	if cmd.Args["i"] == "" {
		c.SendMessage(cmd.Executor, "参数不足，正确用法：%v [--u <UID>] [--q <任务ID>] --i <添加/完成/完成全部>。", cmd.Name)
		return
	}
	// 当参数不为完成全部时需要指定任务Id
	if cmd.Args["i"] != "finishall" && cmd.Args["i"] != "完成全部" {
		c.SendMessage(cmd.Executor, "你需要指定任务ID。", cmd.Name)
		return
	}

	// 初始值
	var questId uint32 // 任务Id
	// 操控任务的模式
	// add 添加任务
	// finish 完成任务
	// finishall 完成全部任务
	var mode string

	// 选择每个参数
	for k, v := range cmd.Args {
		var err error

		switch k {
		case "u":
			var uid uint64
			if uid, err = strconv.ParseUint(v, 10, 32); err == nil {
				// 判断目标用户是否在线
				if user := USER_MANAGER.GetOnlineUser(uint32(uid)); user != nil {
					player = user
				} else {
					c.SendMessage(cmd.Executor, "目标玩家不在线，UID：%v。", v)
					return
				}
			}
		case "q":
			var id uint64
			if id, err = strconv.ParseUint(v, 10, 32); err == nil {
				questId = uint32(id)
			}
		case "i":
			switch v {
			case "add", "添加", "finish", "完成", "finishall", "完成全部":
				// 将模式修改为参数的值
				mode = v
			default:
				c.SendMessage(cmd.Executor, "参数 --%v 只允许使用 <添加/完成/完成全部>。", k)
				return
			}
		default:
			c.SendMessage(cmd.Executor, "参数 --%v 冗余。", k)
			return
		}

		// 解析错误的话应该是参数类型问题
		if err != nil {
			c.SendMessage(cmd.Executor, "参数 --%v 有误，类型错误。", k)
			return
		}
	}

	switch mode {
	case "add", "添加":
		// 添加指定任务
		c.gmCmd.GMAddQuest(player.PlayerId, questId)
		c.SendMessage(cmd.Executor, "已添加玩家 UID：%v, 的任务，任务ID：%v。", player.PlayerId, questId)
	case "finish", "完成":
		// 完成指定任务
		c.gmCmd.GMFinishQuest(player.PlayerId, questId)
		c.SendMessage(cmd.Executor, "已完成玩家 UID：%v, 的任务，任务ID：%v。", player.PlayerId, questId)
	case "finishall", "完成全部":
		// 强制完成当前所有任务
		c.gmCmd.GMForceFinishAllQuest(player.PlayerId)
		c.SendMessage(cmd.Executor, "已完成玩家 UID：%v, 当前全部任务。", player.PlayerId, questId)
	}
}

// UnlockAllPointCommand 解锁所有锚点命令
func (c *CommandManager) UnlockAllPointCommand(cmd *CommandMessage) {
	player := cmd.Executor
	c.gmCmd.GMUnlockAllPoint(player.PlayerId, player.SceneId)
	c.SendMessage(cmd.Executor, "已解锁玩家 UID：%v, 场景：%v，所有锚点。", player.PlayerId, player.SceneId)
}

// XLuaDebugCommand 主动开启客户端XLUA调试命令
func (c *CommandManager) XLuaDebugCommand(cmd *CommandMessage) {
	player := cmd.Executor
	player.XLuaDebug = true
	c.SendMessage(cmd.Executor, "XLua Debug Enable")
}
