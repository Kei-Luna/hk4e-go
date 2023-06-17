package net

import (
	"reflect"

	"hk4e/pkg/logger"
)

const (
	KcpConnForceClose = iota
	KcpAllConnForceClose
	KcpConnCloseNotify
	KcpConnEstNotify
	KcpConnAddrChangeNotify
)

type KcpEvent struct {
	ConvId       uint64
	EventId      int
	EventMessage any
}

func (k *KcpConnectManager) GetKcpEventInputChan() chan *KcpEvent {
	return k.kcpEventInput
}

func (k *KcpConnectManager) GetKcpEventOutputChan() chan *KcpEvent {
	return k.kcpEventOutput
}

func (k *KcpConnectManager) eventHandle() {
	logger.Debug("event handle start")
	// 事件处理
	for {
		event := <-k.kcpEventInput
		logger.Info("kcp manager recv event, ConvId: %v, EventId: %v, EventMessage Type: %v", event.ConvId, event.EventId, reflect.TypeOf(event.EventMessage))
		switch event.EventId {
		case KcpConnForceClose:
			reason, ok := event.EventMessage.(uint32)
			if !ok {
				logger.Error("event KcpConnForceClose msg type error")
				return
			}
			k.forceCloseKcpConn(event.ConvId, reason)
		case KcpAllConnForceClose:
			// 强制关闭所有连接
			k.closeAllKcpConn()
			logger.Info("all conn has been force close")
		}
	}
}
