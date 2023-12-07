package leawash

import (
	"errors"
	"sbp/internal/app"
	"sbp/internal/entities"
	rabbitEntities "sbp/internal/entities/rabbit"
	"sbp/pkg/rabbitmq"

	leConverter "sbp/internal/conversions"

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

func (leaWashPublisher *leaWashPublisher) SendToLeaPaymentResponse(message entities.PaymentResponse) error {
	leaMessage := leConverter.PaymentResponseToLea(message)
	return leaWashPublisher.sendToLea(leaMessage.WashID, string(rabbitEntities.PaymentResponseMessage), leaMessage)
}

func (leaWashPublisher *leaWashPublisher) SendToLeaPaymentFailedResponse(washID string, postID string, orderID string, err string) error {
	paymentResponse := entities.PaymentResponse{
		WashID:  washID,
		PostID:  postID,
		OrderID: orderID,
		UrlPay:  "",
		Failed:  true,
		Error:   err,
	}
	return leaWashPublisher.SendToLeaPaymentResponse(paymentResponse)
}

func (leaWashPublisher *leaWashPublisher) SendToLeaPaymentNotification(message entities.PaymentNotificationForLea) error {
	leaMessage := leConverter.PaymentNotifcationToLea(message)
	return leaWashPublisher.sendToLea(leaMessage.WashID, string(rabbitEntities.PaymentNotificationMessage), leaMessage)
}

func (leaWashPublisher *leaWashPublisher) sendToLea(washID string, messageType string, messageStruct interface{}) error {
	if messageStruct == nil {
		return errors.New("send to lea failed: message = nil")
	}

	ms := messageStruct
	exchangeName := rabbitEntities.LeaCentralWashExchange
	mt := rabbitEntities.Message(messageType)
	rk := rabbitEntities.RoutingKey(washID)

	return leaWashPublisher.publisher.Send(ms, exchangeName, rk, mt)
}
