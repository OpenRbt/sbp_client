package repconverter

import (
	repEntities "sbp/internal/repository/entities"

	logicEntities "sbp/internal/logic/entities"

	uuid "github.com/satori/go.uuid"
)

// ConvertDeleteWashToDB ...
func ConvertDeleteWashToDB(id uuid.UUID) repEntities.DeleteWash {
	return repEntities.DeleteWash{
		ID: uuid.NullUUID{UUID: id, Valid: true},
	}
}

// ConvertWashListFromDB ...
func ConvertWashListFromDB(washList []repEntities.Wash) []logicEntities.Wash {
	res := make([]logicEntities.Wash, len(washList))

	for i, value := range washList {
		res[i] = ConvertWashFromDB(value)
	}

	return res
}

// ConvertWashFromDB ...
func ConvertWashFromDB(dbWash repEntities.Wash) logicEntities.Wash {
	return logicEntities.Wash{
		ID:               dbWash.ID,
		Password:         dbWash.Password,
		Title:            dbWash.Title,
		Description:      dbWash.Description,
		OwnerID:          dbWash.OwnerID,
		TerminalKey:      dbWash.TerminalKey,
		TerminalPassword: dbWash.TerminalPassword,
		CreatedAt:        dbWash.CreatedAt,
		UpdatedAt:        dbWash.UpdatedAt,
	}
}

// ConvertUpdateWashToDb ...
func ConvertUpdateWashToDb(entity logicEntities.UpdateWash) repEntities.UpdateWash {
	return repEntities.UpdateWash{
		ID:               uuid.NullUUID{UUID: entity.ID, Valid: true},
		Name:             entity.Title,
		Description:      entity.Description,
		TerminalKey:      entity.TerminalKey,
		TerminalPassword: entity.TerminalPassword,
	}
}
