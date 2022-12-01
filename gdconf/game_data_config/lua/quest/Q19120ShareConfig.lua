-- 任务配置数据开始-----------------------------

main_id = 19120

sub_ids = 
{
	1912001,
	1912002,
	1912003,
}
-- 任务配置数据结束---------------------------------
-- >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

-- 父任务执行项数据开始-----------------------------
finish_action = 
{
	CLIENT = { },
	SERVER = { },
}

fail_action = 
{
	CLIENT = { },
	SERVER = { },
}

cancel_action = 
{
	CLIENT = { },
	SERVER = { },
}
-- 父任务执行项数据结束-----------------------------
-- >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
-- Actor模块数据开始--------------------------------
-- 空
-- Actor模块数据结束--------------------------------
-- >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
-- 文本模块数据开始---------------------------------
-- 空
-- 文本模块数据结束---------------------------------
-- >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
-- 路点模块数据开始---------------------------------
-- 空
-- 路点模块数据结束---------------------------------
-- >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>


-- 断线重连生成内容 开始----------------------------
-- 和questdata配的存档点对应
rewind_data = 
{
	["1912001"] = { },
	["1912002"] = { },
	["1912003"] = { },
}
-- 断线重连生成内容 结束----------------------------
-- >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

-- 校验数据 开始----------------------------------
-- 和任务lua中生成NPC/Monster/Gadget/Item等对应
quest_data = 
{
	["1912001"] = 
	{
		npcs = 
		{
			{
				id = 503,
				alias = "Coop_Tohma",
				script = "Actor/Npc/TempNPC",
				pos = "Q1912001tuoma",
				scene_id = 3,
				room_id = 0,
				data_index = 1,
			},
		},
	},
	["1912002"] = {
		npcs = 
		{
			{
				id = 503,
				alias = "Coop_Tohma",
				script = "Actor/Npc/TempNPC",
				pos = "Q1911903tuoma",
				scene_id = 3,
				room_id = 0,
				data_index = 1,
			},
			{
				id = 503,
				alias = "Coop_Tohma",
				script = "Actor/Npc/TempNPC",
				pos = "Q1911903tuoma",
				scene_id = 3,
				room_id = 0,
				data_index = 1,
			},
			{
				id = 503,
				alias = "Coop_Tohma",
				script = "Actor/Npc/TempNPC",
				pos = "Q1911903tuoma",
				scene_id = 3,
				room_id = 0,
				data_index = 1,
			},
			{
				id = 503,
				alias = "Coop_Tohma",
				script = "Actor/Npc/TempNPC",
				pos = "Q1911903tuoma",
				scene_id = 3,
				room_id = 0,
				data_index = 1,
			},
			{
				id = 503,
				alias = "Coop_Tohma",
				script = "Actor/Npc/TempNPC",
				pos = "Q1911903tuoma",
				scene_id = 3,
				room_id = 0,
				data_index = 1,
			},
			{
				id = 503,
				alias = "Coop_Tohma",
				script = "Actor/Npc/TempNPC",
				pos = "Q1911903tuoma",
				scene_id = 3,
				room_id = 0,
				data_index = 1,
			},
		},
	 },
	["1912003"] = 
	{
		npcs = 
		{
			{
				id = 503,
				alias = "Coop_Tohma",
				script = "Actor/Npc/TempNPC",
				pos = "Q1912003tuoma",
				scene_id = 3,
				room_id = 0,
				data_index = 1,
			},
		},
	},
}
-- 校验数据 结束----------------------------------
-- >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>