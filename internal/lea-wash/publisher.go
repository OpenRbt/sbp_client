package leawash

import (
	"errors"
	"sbp/internal/logic"

	logicEntities "sbp/internal/logic/entities"

	rabbitMq "sbp/internal/rabbit-mq"

	shareBusinessEntities "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
)

// check 'LeaWashPublisher' struct = logic interface 'CentralWashPublisher'
var _ = logic.LeaWashPublisher(&leaWashPublisher{})

// LeaWashPublisher ...
type leaWashPublisher struct {
	publisher *rabbitMq.Publisher
}

// NewLeaWashPublisher ...
func NewLeaWashPublisher(rabbitMqClient *rabbitMq.RabbitMqClient) (*leaWashPublisher, error) {
	if rabbitMqClient == nil {
		return nil, errors.New("NewLeaWashPublisher: rabbitMqClient = nil")
	}

	service := logicEntities.ServiceSbpClient
	p, err := rabbitMqClient.CreatePublisher(string(service))
	if err != nil {
		return nil, err
	}

	return &leaWashPublisher{
		publisher: p,
	}, nil
}

// SendToLeaError ...
func (leaWashPublisher *leaWashPublisher) SendToLeaError(
	serviceKey string,
	serverID string,
	postID string,
	orderID string,
	errorDesc string,
	errorCode int64,
) error {

	return leaWashPublisher.SendToLea(serviceKey, string(logicEntities.MessageTypePaymentError), &logicEntities.PayError{
		ServerID:  serverID,
		PostID:    postID,
		OrderID:   orderID,
		ErrorCode: errorCode,
		ErrorDesc: errorDesc,
	})
}

// SendToLea ...
func (leaWashPublisher *leaWashPublisher) SendToLea(serviceKey string, messageType string, messageStruct interface{}) error {
	if messageStruct == nil {
		return errors.New("messageStruct = nil")
	}

	ms := messageStruct
	service := logicEntities.ServiceLeaCentralWash
	mt := shareBusinessEntities.MessageType(messageType)
	rk := shareBusinessEntities.RoutingKey(serviceKey)

	return leaWashPublisher.publisher.Send(ms, service, rk, mt)
}
