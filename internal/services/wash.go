package services

import (
	"context"
	"math/rand"
	"sbp/internal/app"
	"sbp/internal/entities"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type washService struct {
	logger         *zap.SugaredLogger
	passwordLength int

	repository washRepository
	sharePub   app.SharePublisher
}

type washRepository interface {
	app.WashRepository

	GetGroupByID(ctx context.Context, id uuid.UUID) (entities.Group, error)
}

func NewWashService(
	ctx context.Context,
	logger *zap.SugaredLogger,

	passwordLength int,

	repository washRepository,
	sharePub app.SharePublisher,
) (*washService, error) {

	logic := washService{
		logger:         logger,
		passwordLength: passwordLength,

		repository: repository,
		sharePub:   sharePub,
	}

	return &logic, nil
}

func (svc *washService) CreateWash(ctx context.Context, auth *entities.Auth, newWash entities.WashCreation) (entities.Wash, error) {
	group, err := svc.repository.GetGroupByID(ctx, newWash.GroupID)
	if err != nil {
		return entities.Wash{}, err
	}

	if !auth.IsSystemManager() && !auth.IsAdminManageOrganization(group.OrganizationID) {
		return entities.Wash{}, entities.ErrForbidden
	}

	newWash.Password = svc.generatePassword()

	w, err := svc.repository.CreateWash(ctx, newWash)
	if err != nil {
		return entities.Wash{}, err
	}

	err = svc.sharePub.CreateRabbitUser(w.ID.String(), w.Password)
	if err != nil {
		return entities.Wash{}, err
	}

	return w, nil
}

func (svc *washService) GetWashByID(ctx context.Context, auth *entities.Auth, id uuid.UUID) (entities.Wash, error) {
	wash, err := svc.repository.GetWashByID(ctx, id)
	if err != nil {
		return entities.Wash{}, err
	}

	if auth.IsSystemManager() || auth.IsAdminManageOrganization(wash.OrganizationID.UUID) {
		return wash, nil
	}

	return entities.Wash{}, entities.ErrForbidden
}

func (svc *washService) UpdateWash(ctx context.Context, auth *entities.Auth, id uuid.UUID, updateWash entities.WashUpdate) error {
	wash, err := svc.repository.GetWashByID(ctx, id)
	if err != nil {
		return err
	}

	if auth.IsSystemManager() || auth.IsAdminManageOrganization(wash.OrganizationID.UUID) {
		return svc.repository.UpdateWash(ctx, id, updateWash)
	}

	return entities.ErrForbidden
}

func (svc *washService) DeleteWash(ctx context.Context, auth *entities.Auth, id uuid.UUID) error {
	wash, err := svc.repository.GetWashByID(ctx, id)
	if err != nil {
		return err
	}

	if auth.IsSystemManager() || auth.IsAdminManageOrganization(wash.OrganizationID.UUID) {
		return svc.repository.DeleteWash(ctx, id)
	}

	return entities.ErrForbidden
}

func (svc *washService) GetWashes(ctx context.Context, auth *entities.Auth, filter entities.WashFilter) ([]entities.Wash, error) {
	if auth.IsSystemManager() {
		return svc.repository.GetWashes(ctx, filter)
	}

	if auth.IsAdmin() && auth.User.OrganizationID != nil {
		filter.OrganizationID = auth.User.OrganizationID
		return svc.repository.GetWashes(ctx, filter)
	}

	return nil, entities.ErrForbidden
}

func (svc *washService) AssignWashToGroup(ctx context.Context, auth *entities.Auth, washID, groupID uuid.UUID) error {
	wash, err := svc.repository.GetWashByID(ctx, washID)
	if err != nil {
		return err
	}

	group, err := svc.repository.GetGroupByID(ctx, groupID)
	if err != nil {
		return err
	}

	isUserServerManager := auth.IsAdminManageOrganization(wash.OrganizationID.UUID)
	if err != nil {
		return err
	}

	isUserGroupManager := auth.IsAdminManageOrganization(group.OrganizationID)
	if err != nil {
		return err
	}

	if auth.IsSystemManager() || (isUserServerManager && isUserGroupManager) {
		return svc.repository.AssignWashToGroup(ctx, washID, groupID)
	}

	return entities.ErrForbidden
}

func (logic *washService) generatePassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	password := make([]byte, logic.passwordLength)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}
