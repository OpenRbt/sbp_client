package app

import (
	"context"
	"net/http"
	"sbp/internal/entities"
	"sbp/openapi/restapi/operations"
)

type (
	Ctx  = context.Context
	Auth = entities.Auth
)

type LeaWashPublisher interface {
	SendToLeaPaymentResponse(entities.PaymentResponse) error
	SendToLeaPaymentNotification(entities.PaymentNotificationForLea) error
	SendToLeaPaymentFailedResponse(washID string, postID string, orderID string, err string) error
}

type SharePublisher interface {
	SendDataRequest() error
	CreateRabbitUser(login, password string) error
}

type Repository interface {
	PaymentRepository
	UserRepository
	WashRepository
	GroupRepository
	OrganizationRepository
	Close() error
}

type Service interface {
	WashService
	PaymentService
	OrganizationService
	GroupService
	UserService
}

type Api interface {
	GetSwaggerApi() *operations.WashSbpAPI
	GetHandler() http.Handler
}
