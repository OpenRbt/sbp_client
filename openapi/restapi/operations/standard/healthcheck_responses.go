// Code generated by go-swagger; DO NOT EDIT.

package standard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// HealthcheckOKCode is the HTTP code returned for type HealthcheckOK
const HealthcheckOKCode int = 200

/*
HealthcheckOK OK

swagger:response healthcheckOK
*/
type HealthcheckOK struct {

	/*
	  In: Body
	*/
	Payload *HealthcheckOKBody `json:"body,omitempty"`
}

// NewHealthcheckOK creates HealthcheckOK with default headers values
func NewHealthcheckOK() *HealthcheckOK {

	return &HealthcheckOK{}
}

// WithPayload adds the payload to the healthcheck o k response
func (o *HealthcheckOK) WithPayload(payload *HealthcheckOKBody) *HealthcheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the healthcheck o k response
func (o *HealthcheckOK) SetPayload(payload *HealthcheckOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HealthcheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *HealthcheckOK) HealthcheckResponder() {}

type HealthcheckNotImplementedResponder struct {
	middleware.Responder
}

func (*HealthcheckNotImplementedResponder) HealthcheckResponder() {}

func HealthcheckNotImplemented() HealthcheckResponder {
	return &HealthcheckNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Healthcheck has not yet been implemented",
		),
	}
}

type HealthcheckResponder interface {
	middleware.Responder
	HealthcheckResponder()
}
