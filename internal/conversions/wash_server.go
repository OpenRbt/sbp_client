package conversions

import (
	"sbp/internal/app"
	"sbp/internal/dal/dbmodels"
	"sbp/openapi/models"

	"github.com/OpenRbt/share_business/wash_rabbit/entity/admin"
)

func WashServerFromDB(dbWashServer dbmodels.WashServer) app.WashServer {
	return app.WashServer{
		ID:               dbWashServer.ID.UUID,
		Title:            dbWashServer.Title,
		Description:      dbWashServer.Description,
		Owner:            dbWashServer.Owner.UUID,
		ServiceKey:       dbWashServer.ServiceKey,
		TerminalKey:      dbWashServer.TerminalKey,
		TerminalPassword: dbWashServer.TerminalPassword,
	}
}

func WashServerToRest(washServer app.WashServer) *models.WashServer {
	return &models.WashServer{
		ID:               washServer.ID.String(),
		Name:             washServer.Title,
		Description:      washServer.Description,
		ServiceKey:       washServer.ServiceKey,
		TerminalPassword: washServer.TerminalPassword,
	}
}

func RegisterWashServerFromRest(rest models.WashServerAdd) app.RegisterWashServer {
	return app.RegisterWashServer{
		Title:            *rest.Name,
		Description:      rest.Description,
		TerminalKey:      rest.TerminalKey,
		TerminalPassword: rest.TerminalPassword,
	}
}

func WashServerToRabbit(e app.WashServer) admin.ServerRegistered {
	return admin.ServerRegistered{
		ID:          e.ID.String(),
		Title:       e.Title,
		Description: e.Description,
	}
}

func WashServerUpdateToRabbit(e app.UpdateWashServer, deleted bool) admin.ServerUpdate {
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
