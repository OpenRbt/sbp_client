package handlers

import (
	"context"
	logicEntities "sbp/internal/logic/entities"
	openapiEntities "sbp/openapi/models"
	washServers "sbp/openapi/restapi/operations/wash_servers"
)

// Auth ...
func (handler Handler) Auth(token string) (*logicEntities.Auth, error) {
	auth, err := handler.logic.Auth(token)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

// SignUP ...
func (handler Handler) SignUP(params washServers.SignupParams, auth *logicEntities.Auth) washServers.SignupResponder {
	res, err := handler.logic.SignUp(context.TODO())
	switch {
	case err == nil:
		return washServers.NewSignupOK().WithPayload(&openapiEntities.FirebaseToken{
			Value: res.Value,
		})
	}
	return washServers.NewSignupOK()
}
