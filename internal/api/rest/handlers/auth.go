package handlers

import (
	// "context"
	"context"
	logicEntities "sbp/internal/logic/entities"
	openapiEntities "sbp/openapi/models"
	wash "sbp/openapi/restapi/operations/wash"
)

// Auth ...
func (handler Handler) Auth(token string) (*logicEntities.AuthExtended, error) {
	auth, err := handler.logic.Auth(token)
	if err != nil {
		return &logicEntities.AuthExtended{}, nil
	}
	if auth == nil {
		return &logicEntities.AuthExtended{}, nil
	}

	ctx := context.TODO()
	user, err := handler.logic.CreateUser(ctx, auth)
	if err != nil {
		return &logicEntities.AuthExtended{}, nil
	}
	if user == nil {
		return &logicEntities.AuthExtended{}, nil
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
