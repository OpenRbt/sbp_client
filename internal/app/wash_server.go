package app

import (
	"context"

	rabbit_vo "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type WashServerService interface {
	GetWashServer(ctx context.Context, auth *Auth, id uuid.UUID) (WashServer, error)
	RegisterWashServer(ctx context.Context, auth *Auth, newServer RegisterWashServer) (WashServer, error)
	UpdateWashServer(ctx context.Context, auth *Auth, updateWashServer UpdateWashServer) error
	DeleteWashServer(ctx context.Context, auth *Auth, id uuid.UUID) error
	GetWashServerList(ctx context.Context, auth *Auth, getWashServerList Pagination) ([]WashServer, error)

	Pay(ctx context.Context, amount int, serverID string, postID string) (string, string, error)
	Cancel(ctx context.Context, paymentID string) error

	Notification(ctx context.Context, notification RegisterNotification) error
}

type Repository interface {
	GetOrCreateAdminIfNotExists(ctx context.Context, identity string) (SbpAdmin, error)
	GetSbpAdmin(ctx context.Context, identity string) (SbpAdmin, error)
	GetWashServer(ctx context.Context, id uuid.UUID) (WashServer, error)

	RegisterWashServer(ctx context.Context, admin uuid.UUID, newServer RegisterWashServer) (WashServer, error)
	UpdateWashServer(ctx context.Context, updateWashServer UpdateWashServer) error
	DeleteWashServer(ctx context.Context, id uuid.UUID) error
	GetWashServerList(ctx context.Context, pagination Pagination) ([]WashServer, error)

	NewTransaction(ctx context.Context, serverID string, postID string, amount int) (Transaction, error)
	UpdateTransaction(ctx context.Context, id uuid.UUID, paymentID *string, status *string) error
	GetTransaction(ctx context.Context, orderID uuid.UUID) (Transaction, error)
}

type WashServerSvc struct {
	l    *zap.SugaredLogger
	repo Repository

	r RabbitSvc

	t TinkoffSvc
}

type RabbitSvc interface {
	CreateRabbitUser(userID, userKey string) (err error)
	SendMessage(msg interface{}, service rabbit_vo.Service, routingKey rabbit_vo.RoutingKey, messageType rabbit_vo.MessageType) error
}

type TinkoffSvc interface {
	Init(terminalKey string, amount int64, orderID string) (Init, error)
	GetQr(terminalKey string, paymentId string, token string) (GetQr, error)
	Cancel(terminalKey string, paymentId string, token string) (Cancel, error)
}

func NewWashServerService(logger *zap.SugaredLogger, repo Repository, rabbit RabbitSvc, tinkoff TinkoffSvc) WashServerService {
	return &WashServerSvc{
		l:    logger,
		repo: repo,
		r:    rabbit,
		t:    tinkoff,
	}
}

func (svc *WashServerSvc) RegisterWashServer(ctx context.Context, auth *Auth, newServer RegisterWashServer) (WashServer, error) {

	admin, err := svc.repo.GetOrCreateAdminIfNotExists(ctx, auth.UID)

	if err != nil {
		return WashServer{}, err
	}

	registered, err := svc.repo.RegisterWashServer(ctx, admin.ID, newServer)
	if err != nil {
		return WashServer{}, err
	}

	// err = svc.r.CreateRabbitUser(registered.ID.String(), registered.ServiceKey)
	// if err != nil {
	// 	return WashServer{}, err
	// }

	return registered, nil
}

func (svc *WashServerSvc) GetWashServer(ctx context.Context, auth *Auth, id uuid.UUID) (WashServer, error) {
	// owner, err := svc.repo.GetOrCreateAdminIfNotExists(ctx, auth.UID)

	// if err != nil {
	// 	return WashServer{}, err
	// }

	return svc.repo.GetWashServer(ctx, id)
}

func (svc *WashServerSvc) UpdateWashServer(ctx context.Context, auth *Auth, updateWashServer UpdateWashServer) error {
	owner, err := svc.repo.GetSbpAdmin(ctx, auth.UID)

	if err != nil {
		return err
	}

	washServer, err := svc.repo.GetWashServer(ctx, updateWashServer.ID)

	if err != nil {
		return err
	}

	if washServer.Owner != owner.ID {
		return ErrUserNotOwner
	}

	err = svc.repo.UpdateWashServer(ctx, updateWashServer)
	if err != nil {
		return err
	}

	return nil
}

func (svc *WashServerSvc) DeleteWashServer(ctx context.Context, auth *Auth, id uuid.UUID) error {
	owner, err := svc.repo.GetSbpAdmin(ctx, auth.UID)
	if err != nil {
		return err
	}

	washServer, err := svc.repo.GetWashServer(ctx, id)
	if err != nil {
		return err
	}

	if washServer.Owner != owner.ID {
		return ErrUserNotOwner
	}
	err = svc.repo.DeleteWashServer(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (svc *WashServerSvc) GetWashServerList(ctx context.Context, auth *Auth, pagination Pagination) ([]WashServer, error) {
	// owner, err := svc.repo.GetOrCreateAdminIfNotExists(ctx, auth.UID)

	// if err != nil {
	// 	return []WashServer{}, err
	// }

	return svc.repo.GetWashServerList(ctx, pagination)
}
