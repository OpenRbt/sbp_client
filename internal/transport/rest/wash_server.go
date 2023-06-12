package rest

import (
	"errors"
	"fmt"
	"log"
	"sbp/internal/app"
	"sbp/internal/conversions"
	"sbp/openapi/models"
	"sbp/openapi/restapi/operations"
	"sbp/openapi/restapi/operations/wash_servers"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initWashServerHandlers(api *operations.WashSbpAPI) {
	api.WashServersGetWashServerHandler = wash_servers.GetWashServerHandlerFunc(svc.getWashServer)
	api.WashServersAddHandler = wash_servers.AddHandlerFunc(svc.addWashServer)
	api.WashServersUpdateHandler = wash_servers.UpdateHandlerFunc(svc.updateWashServer)
	api.WashServersDeleteHandler = wash_servers.DeleteHandlerFunc(svc.deleteWashServer)

	api.WashServersNotificationHandler = wash_servers.NotificationHandlerFunc(svc.notif)

	api.WashServersPayHandler = wash_servers.PayHandlerFunc(svc.pay)
	api.WashServersCancelHandler = wash_servers.CancelHandlerFunc(svc.cancel)
}

func (svc *service) cancel(params wash_servers.CancelParams, auth *app.Auth) wash_servers.CancelResponder {
	err := svc.washServers.Cancel(params.HTTPRequest.Context(), params.Body.OrderID)

	switch {
	case err == nil:
		return wash_servers.NewCancelOK()
	default:
		return wash_servers.NewCancelBadRequest()
	}
}

func (svc *service) pay(params wash_servers.PayParams, auth *app.Auth) wash_servers.PayResponder {
	url, orderID, err := svc.washServers.Pay(params.HTTPRequest.Context(), int(params.Body.Amount), params.Body.ServerID, params.Body.PostID)

	switch {
	case err == nil:
		return wash_servers.NewPayOK().WithPayload(&models.PayResponse{URL: url, OrderID: orderID})
	default:
		return wash_servers.NewPayBadRequest()
	}
}

func (svc *service) notif(params wash_servers.NotificationParams, auth *app.Auth) wash_servers.NotificationResponder {
	registerNotif := conversions.RegisterNotificationFromRest(*params.Body)

	err := svc.washServers.Notification(params.HTTPRequest.Context(), registerNotif)

	if err != nil {
		fmt.Println("Error ", err)
		wash_servers.NewNotificationOK().WithPayload("Not OK")
	}

	return wash_servers.NewNotificationOK().WithPayload("OK")
}

func (svc *service) getWashServer(params wash_servers.GetWashServerParams, auth *app.Auth) wash_servers.GetWashServerResponder {
	id, err := uuid.FromString(params.ID)

	if err != nil {
		return wash_servers.NewGetWashServerBadRequest()
	}

	res, err := svc.washServers.GetWashServer(params.HTTPRequest.Context(), auth, id)

	switch {
	case err == nil:
		return wash_servers.NewGetWashServerOK().WithPayload(conversions.WashServerToRest(res))
	case errors.Is(err, app.ErrNotFound):
		return wash_servers.NewGetWashServerNotFound()
	default:
		return wash_servers.NewGetWashServerInternalServerError()
	}
}

func (svc *service) addWashServer(params wash_servers.AddParams, auth *app.Auth) wash_servers.AddResponder {
	registerWashServerFromRest := conversions.RegisterWashServerFromRest(*params.Body)

	newServer, err := svc.washServers.RegisterWashServer(params.HTTPRequest.Context(), auth, registerWashServerFromRest)

	if err != nil {
		log.Println(err)
	}

	switch {
	case err == nil:
		return wash_servers.NewAddOK().WithPayload(conversions.WashServerToRest(newServer))
	case errors.Is(err, app.ErrNotFound):
		return wash_servers.NewAddBadRequest()
	default:
		return wash_servers.NewAddInternalServerError()
	}
}

func (svc *service) updateWashServer(params wash_servers.UpdateParams, auth *app.Auth) wash_servers.UpdateResponder {
	updateWashServerFromRest, err := conversions.UpdateWashServerFromRest(*params.Body)

	if err != nil {
		return wash_servers.NewUpdateBadRequest()
	}

	err = svc.washServers.UpdateWashServer(params.HTTPRequest.Context(), auth, updateWashServerFromRest)

	switch {
	case err == nil:
		return wash_servers.NewUpdateNoContent()
	case errors.Is(err, app.ErrNotFound):
		return wash_servers.NewUpdateNotFound()
	default:
		return wash_servers.NewUpdateInternalServerError()
	}
}

func (svc *service) deleteWashServer(params wash_servers.DeleteParams, auth *app.Auth) wash_servers.DeleteResponder {
	deleteWashServerFromRest, err := conversions.DeleteWashServerFromRest(*params.Body)

	if err != nil {
		return wash_servers.NewDeleteBadRequest()
	}

	err = svc.washServers.DeleteWashServer(params.HTTPRequest.Context(), auth, deleteWashServerFromRest)

	switch {
	case err == nil:
		return wash_servers.NewDeleteNoContent()
	case errors.Is(err, app.ErrNotFound):
		return wash_servers.NewDeleteNotFound()
	default:
		return wash_servers.NewDeleteInternalServerError()
	}
}

func (svc *service) getWashServerList(params wash_servers.ListParams, auth *app.Auth) wash_servers.ListResponder {
	res, err := svc.washServers.GetWashServerList(params.HTTPRequest.Context(), auth, conversions.PaginationFromRest(*params.Body))

	payload := conversions.WashServerListToRest(res)

	switch {
	case err == nil:
		return wash_servers.NewListOK().WithPayload(payload)
	case errors.Is(err, app.ErrNotFound):
		return wash_servers.NewListNotFound()
	default:
		return wash_servers.NewListInternalServerError()
	}
}
