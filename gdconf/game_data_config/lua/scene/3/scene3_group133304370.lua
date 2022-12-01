-- 基础信息
local base_info = {
	group_id = 133304370
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 370001, gadget_id = 70220008, pos = { x = -1716.277, y = 241.334, z = 2767.675 }, rot = { x = 76.269, y = 358.658, z = 11.832 }, level = 30, area_id = 21 },
	{ config_id = 370002, gadget_id = 70220008, pos = { x = -1716.374, y = 242.318, z = 2773.399 }, rot = { x = 37.054, y = 354.344, z = 6.646 }, level = 30, area_id = 21 },
	{ config_id = 370003, gadget_id = 70220008, pos = { x = -1718.522, y = 242.383, z = 2768.398 }, rot = { x = 319.060, y = 180.000, z = 180.000 }, level = 30, area_id = 21 },
	{ config_id = 370004, gadget_id = 70300094, pos = { x = -1712.282, y = 243.483, z = 2777.044 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1, area_id = 21 },
	{ config_id = 370005, gadget_id = 70300104, pos = { x = -1711.952, y = 243.701, z = 2777.669 }, rot = { x = 0.000, y = 227.812, z = 0.000 }, level = 30, area_id = 21 },
	{ config_id = 370006, gadget_id = 70220008, pos = { x = -1717.254, y = 241.954, z = 2771.125 }, rot = { x = 76.269, y = 358.658, z = 343.154 }, level = 30, area_id = 21 },
	{ config_id = 370007, gadget_id = 70220008, pos = { x = -1717.005, y = 240.779, z = 2765.191 }, rot = { x = 319.060, y = 180.000, z = 180.000 }, level = 30, area_id = 21 },
	{ config_id = 370008, gadget_id = 70220008, pos = { x = -1718.661, y = 240.057, z = 2763.310 }, rot = { x = 319.060, y = 180.000, z = 180.000 }, level = 30, area_id = 21 },
	{ config_id = 370009, gadget_id = 70300094, pos = { x = -1713.989, y = 243.301, z = 2775.864 }, rot = { x = 0.000, y = 133.639, z = 0.000 }, level = 1, area_id = 21 }
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
		monsters = { },
		gadgets = { 370001, 370002, 370003, 370004, 370005, 370006, 370007, 370008, 370009 },
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