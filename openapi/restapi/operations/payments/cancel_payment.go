// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"sbp/internal/entities"
)

// CancelPaymentHandlerFunc turns a function with the right signature into a cancel payment handler
type CancelPaymentHandlerFunc func(CancelPaymentParams, *entities.Auth) CancelPaymentResponder

// Handle executing the request and returning a response
func (fn CancelPaymentHandlerFunc) Handle(params CancelPaymentParams, principal *entities.Auth) CancelPaymentResponder {
	return fn(params, principal)
}

// CancelPaymentHandler interface for that can handle valid cancel payment params
type CancelPaymentHandler interface {
	Handle(CancelPaymentParams, *entities.Auth) CancelPaymentResponder
}

// NewCancelPayment creates a new http.Handler for the cancel payment operation
func NewCancelPayment(ctx *middleware.Context, handler CancelPaymentHandler) *CancelPayment {
	return &CancelPayment{Context: ctx, Handler: handler}
}

/*
	CancelPayment swagger:route POST /payments/cancel payments cancelPayment

CancelPayment cancel payment API
*/
type CancelPayment struct {
	Context *middleware.Context
	Handler CancelPaymentHandler
}

func (o *CancelPayment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCancelPaymentParams()
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
