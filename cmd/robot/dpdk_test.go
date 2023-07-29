//go:build linux
// +build linux

package main

import (
	"testing"

	"github.com/flswld/dpdk-go/dpdk"
	"github.com/flswld/dpdk-go/engine"
)

func TestDpdk(t *testing.T) {
	ec := &engine.Config{
		DebugLog:      true,                // 调试日志
		MacAddr:       "00:0C:29:3E:3E:DF", // mac地址
		IpAddr:        "192.168.199.199",   // ip地址
		NetworkMask:   "255.255.255.0",     // 子网掩码
		GatewayIpAddr: "192.168.199.1",     // 网关ip地址
	}
	// 初始化协议栈
	e, err := engine.InitEngine(ec)
	if err != nil {
		panic(err)
	}
	dc := &dpdk.Config{
		GolangCpuCoreList: []int{4, 5},       // golang侧使用的核心编号列表 非单核模式下至少需要两个核心
		DpdkCpuCoreList:   []int{0, 1, 2, 3}, // dpdk侧使用的核心编号列表 非单核模式下至少需要四个核心
		DpdkMemChanNum:    1,                 // dpdk内存通道数
		DebugLog:          true,              // 收发包调试日志
		IdleSleep:         false,             // 空闲睡眠 降低cpu占用
		SingleCore:        false,             // 单核模式 物理单核机器需要开启
		KniBypassTargetIp: false,             // kni旁路目标ip dpdk-go只接收来自目标ip的包 其他的包全部送到kni网卡
		TargetIpAddr:      "",                // 目标ip地址
	}
	// 启动协议栈
	e.RunEngine(dc)

	// 停止协议栈
	e.StopEngine()
}
