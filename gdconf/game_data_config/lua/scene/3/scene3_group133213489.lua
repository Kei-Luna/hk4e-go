-- 基础信息
local base_info = {
	group_id = 133213489
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
	{ config_id = 489002, gadget_id = 70210101, pos = { x = -3379.304, y = 203.525, z = -3259.144 }, rot = { x = 0.000, y = 268.013, z = 0.000 }, level = 26, drop_tag = "搜刮点解谜破损稻妻", area_id = 12 },
	{ config_id = 489003, gadget_id = 70210101, pos = { x = -3256.882, y = 201.731, z = -3271.694 }, rot = { x = 0.104, y = 268.016, z = 3.003 }, level = 26, drop_tag = "搜刮点解谜果蔬稻妻", area_id = 12 },
	{ config_id = 489006, gadget_id = 70210101, pos = { x = -3273.439, y = 200.865, z = -3249.993 }, rot = { x = 354.463, y = 268.307, z = 353.946 }, level = 26, drop_tag = "搜刮点解谜破损稻妻", area_id = 12 },
	{ config_id = 489007, gadget_id = 70210101, pos = { x = -3270.376, y = 200.865, z = -3253.775 }, rot = { x = 354.463, y = 268.307, z = 353.946 }, level = 26, drop_tag = "搜刮点解谜破损稻妻", area_id = 12 },
	{ config_id = 489009, gadget_id = 70210101, pos = { x = -3250.570, y = 200.260, z = -3259.092 }, rot = { x = 359.079, y = 269.445, z = 357.325 }, level = 26, drop_tag = "搜刮点解谜遗物稻妻", area_id = 12 },
	{ config_id = 489010, gadget_id = 70210101, pos = { x = -3252.209, y = 200.260, z = -3255.784 }, rot = { x = 359.079, y = 269.445, z = 357.325 }, level = 26, drop_tag = "搜刮点解谜破损稻妻", area_id = 12 },
	{ config_id = 489012, gadget_id = 70210101, pos = { x = -3263.844, y = 201.414, z = -3271.909 }, rot = { x = 357.380, y = 267.993, z = 1.881 }, level = 26, drop_tag = "搜刮点解谜破损稻妻", area_id = 12 },
	{ config_id = 489013, gadget_id = 70210101, pos = { x = -3269.716, y = 202.805, z = -3270.466 }, rot = { x = 0.000, y = 268.013, z = 0.000 }, level = 26, drop_tag = "搜刮点解谜果蔬稻妻", area_id = 12 }
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

-- 废弃数据
garbages = {
	gadgets = {
		{ config_id = 489001, gadget_id = 70210101, pos = { x = -3265.573, y = 202.805, z = -3267.946 }, rot = { x = 0.000, y = 268.013, z = 0.000 }, level = 26, drop_tag = "搜刮点解谜果蔬稻妻", area_id = 12 },
		{ config_id = 489004, gadget_id = 70210101, pos = { x = -3255.749, y = 200.520, z = -3271.316 }, rot = { x = 354.646, y = 268.276, z = 0.161 }, level = 26, drop_tag = "搜刮点解谜果蔬稻妻", area_id = 12 },
		{ config_id = 489005, gadget_id = 70210101, pos = { x = -3254.645, y = 200.689, z = -3273.498 }, rot = { x = 314.965, y = 275.689, z = 347.441 }, level = 26, drop_tag = "搜刮点解谜果蔬稻妻", area_id = 12 },
		{ config_id = 489008, gadget_id = 70210101, pos = { x = -3267.671, y = 202.805, z = -3262.483 }, rot = { x = 0.000, y = 268.013, z = 0.000 }, level = 26, drop_tag = "搜刮点解谜果蔬稻妻", area_id = 12 },
		{ config_id = 489011, gadget_id = 70210101, pos = { x = -3262.676, y = 202.171, z = -3270.022 }, rot = { x = 359.072, y = 268.016, z = 359.694 }, level = 26, drop_tag = "搜刮点解谜果蔬稻妻", area_id = 12 }
	}
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
		gadgets = { 489002, 489003, 489006, 489007, 489009, 489010, 489012, 489013 },
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