-- 基础信息
local base_info = {
	group_id = 133314087
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 87001, monster_id = 28040101, pos = { x = -203.907, y = -49.000, z = 4798.453 }, rot = { x = 0.000, y = 49.043, z = 0.000 }, level = 32, drop_tag = "水族", area_id = 32 },
	{ config_id = 87002, monster_id = 28040102, pos = { x = -192.289, y = -49.000, z = 4777.157 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, drop_tag = "水族", area_id = 32 },
	{ config_id = 87003, monster_id = 28040103, pos = { x = -192.899, y = -49.000, z = 4787.090 }, rot = { x = 0.000, y = 15.953, z = 0.000 }, level = 32, drop_tag = "水族", area_id = 32 },
	{ config_id = 87004, monster_id = 28040102, pos = { x = -163.514, y = -49.000, z = 4861.854 }, rot = { x = 0.000, y = 317.954, z = 0.000 }, level = 32, drop_tag = "水族", area_id = 32 },
	{ config_id = 87005, monster_id = 28040102, pos = { x = -258.347, y = -38.900, z = 4721.713 }, rot = { x = 0.000, y = 221.366, z = 0.000 }, level = 32, drop_tag = "水族", area_id = 32 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 变量
variables = {
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 0,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { 87001, 87002, 87003, 87004, 87005 },
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================