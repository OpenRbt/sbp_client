package conversions

import (
	"sbp/internal/app"
	"sbp/internal/dal/dbmodels"
	"sbp/openapi/models"

	uuid "github.com/satori/go.uuid"
)

func UpdateWashServerToDb(entity app.UpdateWashServer) dbmodels.UpdateWashServer {
	return dbmodels.UpdateWashServer{
		ID:               uuid.NullUUID{UUID: entity.ID, Valid: true},
		Name:             entity.Title,
		Description:      entity.Description,
		TerminalKey:      entity.TerminalKey,
		TerminalPassword: entity.TerminalPassword,
	}
}

func UpdateWashServerFromRest(model models.WashServerUpdate) (app.UpdateWashServer, error) {
	id, err := uuid.FromString(model.ID)

	if err != nil {
		return app.UpdateWashServer{}, err
	}

	return app.UpdateWashServer{
		ID:          id,
		Title:       &model.Name,
		Description: &model.Description,
	}, nil
}
