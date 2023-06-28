package handlers

import (
	"context"

	logicEntities "sbp/internal/logic/entities"

	uuid "github.com/satori/go.uuid"
)

// Logic ...
type Logic interface {
	// admin
	GetOrCreateAdminIfNotExists(ctx context.Context, auth *logicEntities.Auth) (*logicEntities.SbpAdmin, error)
	GetSbpAdmin(ctx context.Context, auth *logicEntities.Auth) (*logicEntities.SbpAdmin, error)

	// auth
	Auth(token string) (*logicEntities.Auth, error)
	SignUp(ctx context.Context) (*logicEntities.Token, error)

	// wash server
	CreateWashServer(ctx context.Context, admin logicEntities.SbpAdmin, s logicEntities.RegisterWashServer) (logicEntities.WashServer, error)
	UpdateWashServer(ctx context.Context, admin logicEntities.SbpAdmin, updateWashServer logicEntities.UpdateWashServer) error
	DeleteWashServer(ctx context.Context, admin logicEntities.SbpAdmin, id uuid.UUID) error
	GetWashServer(ctx context.Context, id uuid.UUID) (logicEntities.WashServer, error)
	GetWashServerList(ctx context.Context, pagination logicEntities.Pagination) ([]logicEntities.WashServer, error)

	// payment
	Pay(ctx context.Context, req logicEntities.PayRequest) (*logicEntities.PayResponse, error)
	Cancel(ctx context.Context, req logicEntities.PayСancellationRequest) (resendNeaded bool, err error)
	Notification(ctx context.Context, notification logicEntities.PaymentRegisterNotification) error
}
