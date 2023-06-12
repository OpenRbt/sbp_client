package conversions

import (
	"sbp/internal/app"
	"sbp/internal/dal/dbmodels"
)

func SbpAdminFromDB(dbSbpAdmin dbmodels.SbpAdmin) app.SbpAdmin {
	return app.SbpAdmin{
		ID:       dbSbpAdmin.ID.UUID,
		Identity: dbSbpAdmin.Identity,
	}
}
