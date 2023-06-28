package repconverter

import (
	repEntities "sbp/internal/repository/entities"

	logicEntities "sbp/internal/logic/entities"

	"github.com/OpenRbt/share_business/wash_rabbit/entity/admin"
	uuid "github.com/satori/go.uuid"
)

func ConvertDeleteWashServerToDB(id uuid.UUID) repEntities.DeleteWashServer {
	return repEntities.DeleteWashServer{
		ID: uuid.NullUUID{UUID: id, Valid: true},
	}
}

func ConvertPaginationToDB(entity logicEntities.Pagination) repEntities.Pagination {
	return repEntities.Pagination{
		Limit:  entity.Limit,
		Offset: entity.Offset,
	}
}

func ConvertWashServerListFromDB(washServerList []repEntities.WashServer) []logicEntities.WashServer {
	res := make([]logicEntities.WashServer, len(washServerList))

	for i, value := range washServerList {
		res[i] = ConvertWashServerFromDB(value)
	}

	return res
}

func ConvertWashServerFromDB(dbWashServer repEntities.WashServer) logicEntities.WashServer {
	return logicEntities.WashServer{
		ID:               dbWashServer.ID.UUID,
		Title:            dbWashServer.Title,
		Description:      dbWashServer.Description,
		Owner:            dbWashServer.Owner.UUID,
		ServiceKey:       dbWashServer.ServiceKey,
		TerminalKey:      dbWashServer.TerminalKey,
		TerminalPassword: dbWashServer.TerminalPassword,
	}
}

func ConvertWashServerToRabbit(e logicEntities.WashServer) admin.ServerRegistered {
	return admin.ServerRegistered{
		ID:          e.ID.String(),
		Title:       e.Title,
		Description: e.Description,
	}
}

func ConvertWashServerUpdateToRabbit(e logicEntities.UpdateWashServer, deleted bool) admin.ServerUpdate {
	var del *bool
	if deleted {
		t := true
		del = &t
	}

	return admin.ServerUpdate{
		ID:          e.ID.String(),
		Title:       e.Title,
		Description: e.Description,
		Deleted:     del,
	}
}

func ConvertUpdateWashServerToDb(entity logicEntities.UpdateWashServer) repEntities.UpdateWashServer {
	return repEntities.UpdateWashServer{
		ID:               uuid.NullUUID{UUID: entity.ID, Valid: true},
		Name:             entity.Title,
		Description:      entity.Description,
		TerminalKey:      entity.TerminalKey,
		TerminalPassword: entity.TerminalPassword,
	}
}
