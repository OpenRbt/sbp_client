package handlers

import (
	// "context"
	"context"
	"errors"
	logicEntities "sbp/internal/logic/entities"
	openapiEntities "sbp/openapi/models"
	wash "sbp/openapi/restapi/operations/wash"
)

// Auth ...
func (handler Handler) Auth(token string) (*logicEntities.AuthExtended, error) {
	auth, err := handler.logic.Auth(token)
	if err != nil {
		return nil, err
	}
	if auth == nil {
		return nil, errors.New("auth failed: auth = nil")
	}

	ctx := context.TODO()
	user, err := handler.logic.CreateUser(ctx, auth)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("auth failed: user = nil")
	}

	return &logicEntities.AuthExtended{
		Auth: *auth,
		User: *user,
	}, nil
}

// SignUP ...
func (handler Handler) SignUP(params wash.SignupParams, auth *logicEntities.AuthExtended) wash.SignupResponder {
	res, err := handler.logic.SignUp(context.TODO())
	switch {
	case err == nil:
		return wash.NewSignupOK().WithPayload(&openapiEntities.FirebaseToken{
			Value: res.Value,
		})
	}
	return wash.NewSignupOK()
}
