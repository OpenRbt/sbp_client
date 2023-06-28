package repconverter

import (
	logicEntities "sbp/internal/logic/entities"
	repEntities "sbp/internal/repository/entities"
)

func ConvertSbpAdminFromDB(dbSbpAdmin repEntities.SbpAdmin) logicEntities.SbpAdmin {
	return logicEntities.SbpAdmin{
		ID:       dbSbpAdmin.ID.UUID,
		Identity: dbSbpAdmin.Identity,
	}
}
