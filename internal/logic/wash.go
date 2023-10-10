package logic

import (
	"context"
	"math/rand"
	logicEntities "sbp/internal/logic/entities"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// WashLogic ...
type WashLogic struct {
	logger            *zap.SugaredLogger
	repository        WashRepository
	passwordLength    int
	brokerUserCreator BrokerUserCreator
}

// WashRepository ...
type WashRepository interface {
	CreateWash(ctx context.Context, new logicEntities.RegisterWash) (logicEntities.Wash, error)
	UpdateWash(ctx context.Context, updateWash logicEntities.UpdateWash) error
	DeleteWash(ctx context.Context, id uuid.UUID) error
	GetWash(ctx context.Context, id uuid.UUID) (logicEntities.Wash, error)
	GetWashList(ctx context.Context, pagination logicEntities.Pagination) ([]logicEntities.Wash, error)
}

// BrokerUserCreator
type BrokerUserCreator interface {
	CreateUser(login string, password string) error
}

// newWashLogic ...
func newWashLogic(
	ctx context.Context,
	logger *zap.SugaredLogger,
	repository WashRepository,
	brokerUserCreator BrokerUserCreator,
	passwordLength int,
) (*WashLogic, error) {

	logic := WashLogic{
		logger:            logger,
		repository:        repository,
		passwordLength:    passwordLength,
		brokerUserCreator: brokerUserCreator,
	}

	return &logic, nil
}

// CreateWash ...
func (logic *WashLogic) CreateWash(ctx context.Context, newWash logicEntities.RegisterWash) (logicEntities.Wash, error) {
	// generate password
	newWash.Password = logic.generatePassword()

	// create
	w, err := logic.repository.CreateWash(ctx, newWash)
	if err != nil {
		return logicEntities.Wash{}, err
	}

	// create broker user
	err = logic.brokerUserCreator.CreateUser(w.ID.String(), w.Password)
	if err != nil {
		return logicEntities.Wash{}, err
	}

	return w, nil
}

// GetWash ...
func (logic *WashLogic) GetWash(ctx context.Context, id uuid.UUID) (logicEntities.Wash, error) {
	wash, err := logic.repository.GetWash(ctx, id)
	if err != nil {
		return logicEntities.Wash{}, err
	}

	return wash, nil
}

// UpdateWash ...
func (logic *WashLogic) UpdateWash(ctx context.Context, updateWash logicEntities.UpdateWash) error {
	wash, err := logic.repository.GetWash(ctx, updateWash.ID)
	if err != nil {
		return err
	}
	if wash.OwnerID != updateWash.OwnerID {
		return logicEntities.ErrUserNotOwner
	}

	err = logic.repository.UpdateWash(ctx, updateWash)
	if err != nil {
		return err
	}
	return nil
}

// DeleteWash ...
func (logic *WashLogic) DeleteWash(ctx context.Context, ownerID uuid.UUID, id uuid.UUID) error {
	wash, err := logic.repository.GetWash(ctx, id)
	if err != nil {
		return err
	}
	if wash.OwnerID != ownerID {
		return logicEntities.ErrUserNotOwner
	}

	err = logic.repository.DeleteWash(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// GetWashList ...
func (logic *WashLogic) GetWashList(ctx context.Context, pagination logicEntities.Pagination) ([]logicEntities.Wash, error) {
	washs, err := logic.repository.GetWashList(ctx, pagination)
	if err != nil {
		return nil, err
	}

	return washs, nil
}

// generatePassword ...
func (logic *WashLogic) generatePassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	// Генерируем пароль
	password := make([]byte, logic.passwordLength)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}
