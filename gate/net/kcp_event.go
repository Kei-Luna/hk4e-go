package net

import (
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
	Conv         uint32
	SessionId    uint32
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
		logger.Info("kcp manager recv event, Conv: %v, SessionId: %v, EventId: %v, EventMessage: %v",
			event.Conv, event.SessionId, event.EventId, event.EventMessage)
		switch event.EventId {
		case KcpConnForceClose:
			reason, ok := event.EventMessage.(uint32)
			if !ok {
				logger.Error("event KcpConnForceClose msg type error")
				return
			}
			k.forceCloseKcpConn(event.SessionId, reason)
		case KcpAllConnForceClose:
			// 强制关闭所有连接
			k.closeAllKcpConn()
			logger.Info("all conn has been force close")
		}
	}
}
