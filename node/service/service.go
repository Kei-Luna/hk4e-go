package service

import (
	"hk4e/common/mq"
	"hk4e/node/api"

	"github.com/byebyebruce/natsrpc"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/protobuf"
)

type Service struct {
	discoveryService *DiscoveryService
}

func NewService(conn *nats.Conn, messageQueue *mq.MessageQueue) (*Service, error) {
	enc, err := nats.NewEncodedConn(conn, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		return nil, err
	}
	svr, err := natsrpc.NewServer(enc)
	if err != nil {
		return nil, err
	}
	discoveryService := NewDiscoveryService(messageQueue)
	_, err = api.RegisterDiscoveryNATSRPCServer(svr, discoveryService)
	if err != nil {
		return nil, err
	}
	s := &Service{
		discoveryService: discoveryService,
	}
	return s, nil
}

func (s *Service) Close() {
}
