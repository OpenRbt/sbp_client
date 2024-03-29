// Code generated by go-swagger; DO NOT EDIT.

package washes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"sbp/internal/entities"
)

// GetWashByIDHandlerFunc turns a function with the right signature into a get wash by Id handler
type GetWashByIDHandlerFunc func(GetWashByIDParams, *entities.Auth) GetWashByIDResponder

// Handle executing the request and returning a response
func (fn GetWashByIDHandlerFunc) Handle(params GetWashByIDParams, principal *entities.Auth) GetWashByIDResponder {
	return fn(params, principal)
}

// GetWashByIDHandler interface for that can handle valid get wash by Id params
type GetWashByIDHandler interface {
	Handle(GetWashByIDParams, *entities.Auth) GetWashByIDResponder
}

// NewGetWashByID creates a new http.Handler for the get wash by Id operation
func NewGetWashByID(ctx *middleware.Context, handler GetWashByIDHandler) *GetWashByID {
	return &GetWashByID{Context: ctx, Handler: handler}
}

/*
	GetWashByID swagger:route GET /washes/{id} washes getWashById

GetWashByID get wash by Id API
*/
type GetWashByID struct {
	Context *middleware.Context
	Handler GetWashByIDHandler
}

func (o *GetWashByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetWashByIDParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *entities.Auth
	if uprinc != nil {
		principal = uprinc.(*entities.Auth) // this is really a entities.Auth, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
