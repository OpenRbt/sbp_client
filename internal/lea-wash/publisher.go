package leawash

import (
	"errors"
	leEntities "sbp/internal/lea-wash/entities"
	"sbp/internal/logic"

	rabbitMq "sbp/internal/rabbit-mq"

	leConverter "sbp/internal/lea-wash/converter"

	logicEntities "sbp/internal/logic/entities"

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
		return nil, errors.New("create lea_wash_publisher failed: rabbitMqClient = nil")
	}

	// create exchange
	sbpClientExchangeName := string(leEntities.ExchangeNameSbpClient)
	leaClientExchangeName := string(leEntities.ExchangeNameLeaCentralWash)
	// sbp
	err := rabbitMqClient.CreateExchangeName(sbpClientExchangeName)
	if err != nil {
		return nil, err
	}
	// lea
	err = rabbitMqClient.CreateExchangeName(leaClientExchangeName)
	if err != nil {
		return nil, err
	}

	// create publisher
	p, err := rabbitMqClient.CreatePublisher(sbpClientExchangeName)
	if err != nil {
		return nil, err
	}

	return &leaWashPublisher{
		publisher: p,
	}, nil
}

// SendToLeaPaymentResponse ...
func (leaWashPublisher *leaWashPublisher) SendToLeaPaymentResponse(message logicEntities.PaymentResponse) error {
	leaMessage := leConverter.PaymentResponseToLea(message)
	return leaWashPublisher.sendToLea(leaMessage.WashID, string(leEntities.MessageTypePaymentResponse), leaMessage)
}

// SendToLeaPaymentResponse ...
func (leaWashPublisher *leaWashPublisher) SendToLeaPaymentFailedResponse(washID string, postID string, orderID string) error {
	paymentResponse := logicEntities.PaymentResponse{
		WashID:  washID,
		PostID:  postID,
		OrderID: orderID,
		UrlPay:  "",
		Failed:  true,
	}
	return leaWashPublisher.SendToLeaPaymentResponse(paymentResponse)
}

// SendToLeaPaymentNotification ...
func (leaWashPublisher *leaWashPublisher) SendToLeaPaymentNotification(message logicEntities.PaymentNotifcation) error {
	leaMessage := leConverter.PaymentNotifcationToLea(message)
	return leaWashPublisher.sendToLea(leaMessage.WashID, string(leEntities.MessageTypePaymentNotification), leaMessage)
}

// sendToLea ...
func (leaWashPublisher *leaWashPublisher) sendToLea(washID string, messageType string, messageStruct interface{}) error {
	if messageStruct == nil {
		return errors.New("send to lea failed: message = nil")
	}

	ms := messageStruct
	exchangeName := leEntities.ExchangeNameLeaCentralWash
	mt := shareBusinessEntities.MessageType(messageType)
	rk := shareBusinessEntities.RoutingKey(washID)

	return leaWashPublisher.publisher.Send(ms, exchangeName, rk, mt)
}
