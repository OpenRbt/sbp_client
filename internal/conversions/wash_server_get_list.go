package conversions

import (
	"sbp/internal/app"
	"sbp/internal/dal/dbmodels"
	"sbp/openapi/models"
)

func PaginationToDB(entity app.Pagination) dbmodels.Pagination {
	return dbmodels.Pagination{
		Limit:  entity.Limit,
		Offset: entity.Offset,
	}
}

func PaginationFromRest(model models.Pagination) app.Pagination {
	return app.Pagination{
		Limit:  model.Limit,
		Offset: model.Offset,
	}
}

func WashServerListFromDB(washServerList []dbmodels.WashServer) []app.WashServer {
	res := make([]app.WashServer, len(washServerList))

	for i, value := range washServerList {
		res[i] = WashServerFromDB(value)
	}

	return res
}

func WashServerListToRest(washServerEntity []app.WashServer) []*models.WashServer {
	res := make([]*models.WashServer, len(washServerEntity))

	for i, value := range washServerEntity {
		rest := WashServerToRest(value)
		res[i] = rest
	}

	return res
}
