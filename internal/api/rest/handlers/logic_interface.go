package handlers

import (
	"context"

	logicEntities "sbp/internal/logic/entities"

	uuid "github.com/satori/go.uuid"
)

// Logic ...
type Logic interface {
	// user
	CreateUser(ctx context.Context, auth *logicEntities.Auth) (*logicEntities.User, error)
	GetUser(ctx context.Context, auth *logicEntities.Auth) (*logicEntities.User, error)

	// auth
	Auth(token string) (*logicEntities.Auth, error)
	SignUp(ctx context.Context) (*logicEntities.Token, error)

	// wash
	CreateWash(ctx context.Context, registerWash logicEntities.RegisterWash) (logicEntities.Wash, error)
	UpdateWash(ctx context.Context, updateWash logicEntities.UpdateWash) error
	DeleteWash(ctx context.Context, id uuid.UUID, ownerId uuid.UUID) error
	GetWash(ctx context.Context, id uuid.UUID) (logicEntities.Wash, error)
	GetWashList(ctx context.Context, pagination logicEntities.Pagination) ([]logicEntities.Wash, error)

	// payment
	Pay(ctx context.Context, req logicEntities.PayRequest) (*logicEntities.PayResponse, error)
	Cancel(ctx context.Context, req logicEntities.PayСancellationRequest) (resendNeaded bool, err error)
	Notification(ctx context.Context, notification logicEntities.PaymentRegisterNotification) error
}
