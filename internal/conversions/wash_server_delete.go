package conversions

import (
	"sbp/internal/dal/dbmodels"
	"sbp/openapi/models"

	uuid "github.com/satori/go.uuid"
)

func DeleteWashServerToDB(id uuid.UUID) dbmodels.DeleteWashServer {
	return dbmodels.DeleteWashServer{
		ID: uuid.NullUUID{UUID: id, Valid: true},
	}
}

func DeleteWashServerFromRest(model models.WashServerDelete) (uuid.UUID, error) {
	id, err := uuid.FromString(*model.ID)

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
