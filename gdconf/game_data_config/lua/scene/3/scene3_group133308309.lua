-- 基础信息
local base_info = {
	group_id = 133308309
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 309001, monster_id = 28010404, pos = { x = -1988.691, y = 341.724, z = 3825.881 }, rot = { x = 0.000, y = 182.698, z = 0.000 }, level = 32, drop_tag = "采集动物", pose_id = 1, area_id = 26 },
	{ config_id = 309002, monster_id = 28010404, pos = { x = -1993.357, y = 346.772, z = 3817.661 }, rot = { x = 0.000, y = 124.099, z = 0.000 }, level = 32, drop_tag = "采集动物", pose_id = 1, area_id = 26 },
	{ config_id = 309003, monster_id = 28010404, pos = { x = -1998.747, y = 345.815, z = 3820.869 }, rot = { x = 0.000, y = 110.960, z = 0.000 }, level = 32, drop_tag = "采集动物", pose_id = 1, area_id = 26 }
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
		monsters = { 309001, 309002, 309003 },
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