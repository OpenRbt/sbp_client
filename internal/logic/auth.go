package logic

import (
	"context"
	logicEntities "sbp/internal/logic/entities"

	"go.uber.org/zap"
)

// AuthLogic ...
type AuthLogic struct {
	logger     *zap.SugaredLogger
	authClient AuthClient
}

// AuthClient ...
type AuthClient interface {
	Auth(token string) (*logicEntities.Auth, error)
	SignUP() (*logicEntities.Token, error)
}

// newAuthLogic ...
func newAuthLogic(ctx context.Context, logger *zap.SugaredLogger, authClient AuthClient) (*AuthLogic, error) {
	return &AuthLogic{
		logger:     logger,
		authClient: authClient,
	}, nil
}

// Auth ...
func (logic *AuthLogic) Auth(token string) (*logicEntities.Auth, error) {
	return logic.authClient.Auth(token)
}

// SignUp ...
func (logic *AuthLogic) SignUp(ctx context.Context) (*logicEntities.Token, error) {
	return logic.authClient.SignUP()
}
