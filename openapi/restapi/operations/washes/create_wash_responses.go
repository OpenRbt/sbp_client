// Code generated by go-swagger; DO NOT EDIT.

package washes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"sbp/openapi/models"
)

// CreateWashOKCode is the HTTP code returned for type CreateWashOK
const CreateWashOKCode int = 200

/*
CreateWashOK Success creation

swagger:response createWashOK
*/
type CreateWashOK struct {

	/*
	  In: Body
	*/
	Payload *models.Wash `json:"body,omitempty"`
}

// NewCreateWashOK creates CreateWashOK with default headers values
func NewCreateWashOK() *CreateWashOK {

	return &CreateWashOK{}
}

// WithPayload adds the payload to the create wash o k response
func (o *CreateWashOK) WithPayload(payload *models.Wash) *CreateWashOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create wash o k response
func (o *CreateWashOK) SetPayload(payload *models.Wash) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateWashOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateWashOK) CreateWashResponder() {}

/*
CreateWashDefault Generic error response

swagger:response createWashDefault
*/
type CreateWashDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateWashDefault creates CreateWashDefault with default headers values
func NewCreateWashDefault(code int) *CreateWashDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateWashDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create wash default response
func (o *CreateWashDefault) WithStatusCode(code int) *CreateWashDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create wash default response
func (o *CreateWashDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create wash default response
func (o *CreateWashDefault) WithPayload(payload *models.Error) *CreateWashDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create wash default response
func (o *CreateWashDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateWashDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateWashDefault) CreateWashResponder() {}

type CreateWashNotImplementedResponder struct {
	middleware.Responder
}

func (*CreateWashNotImplementedResponder) CreateWashResponder() {}

func CreateWashNotImplemented() CreateWashResponder {
	return &CreateWashNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.CreateWash has not yet been implemented",
		),
	}
}

type CreateWashResponder interface {
	middleware.Responder
	CreateWashResponder()
}
