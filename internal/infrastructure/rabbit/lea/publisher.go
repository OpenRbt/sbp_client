package leawash

import (
	"errors"
	"sbp/internal/app"
	"sbp/internal/entities"
	rabbitEntities "sbp/internal/entities/rabbit"
	"sbp/pkg/rabbitmq"

	"sbp/internal/conversions"

	"go.uber.org/zap"
)

var _ = app.LeaWashPublisher(&leaWashPublisher{})

type leaWashPublisher struct {
	publisher *rabbitmq.Publisher
}

func NewLeaWashPublisher(logger *zap.SugaredLogger, rabbitMqClient *rabbitmq.RabbitMqClient) (*leaWashPublisher, error) {
	if rabbitMqClient == nil {
		return nil, errors.New("create lea_wash_publisher failed: rabbitMqClient = nil")
	}

	sbpClientExchangeName := string(rabbitEntities.SbpClientExchange)
	leaClientExchangeName := string(rabbitEntities.LeaCentralWashExchange)

	err := rabbitMqClient.CreateExchangeName(sbpClientExchangeName)
	if err != nil {
		return nil, err
	}

	err = rabbitMqClient.CreateExchangeName(leaClientExchangeName)
	if err != nil {
		return nil, err
	}

	p, err := rabbitmq.NewPublisher(logger, rabbitMqClient, sbpClientExchangeName)
	if err != nil {
		return nil, err
	}

	return &leaWashPublisher{
		publisher: p,
	}, nil
}

func (leaPub *leaWashPublisher) SendToLeaPaymentResponse(message entities.PaymentResponse) error {
	leaMessage := conversions.PaymentResponseToLea(message)
	return leaPub.sendToLea(leaMessage.WashID, string(rabbitEntities.PaymentResponseMessageType), leaMessage)
}

func (leaPub *leaWashPublisher) SendToLeaPaymentFailedResponse(washID string, postID string, orderID string, err string) error {
	paymentResponse := entities.PaymentResponse{
		WashID:  washID,
		PostID:  postID,
		OrderID: orderID,
		UrlPay:  "",
		Failed:  true,
		Error:   err,
	}
	return leaPub.SendToLeaPaymentResponse(paymentResponse)
}

func (leaPub *leaWashPublisher) SendToLeaPaymentNotification(message entities.PaymentNotificationForLea) error {
	leaMessage := conversions.PaymentNotifcationToLea(message)
	return leaPub.sendToLea(leaMessage.WashID, string(rabbitEntities.PaymentNotificationMessageType), leaMessage)
}

func (leaPub *leaWashPublisher) sendToLea(washID string, messageType string, messageStruct interface{}) error {
	if messageStruct == nil {
		return errors.New("send to lea failed: message = nil")
	}

	ms := messageStruct
	exchangeName := rabbitEntities.LeaCentralWashExchange
	mt := rabbitEntities.MessageType(messageType)
	rk := rabbitEntities.RoutingKey(washID)

	return leaPub.publisher.Send(ms, exchangeName, rk, mt)
}
