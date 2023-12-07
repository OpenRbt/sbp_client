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

// GetWashByIDOKCode is the HTTP code returned for type GetWashByIDOK
const GetWashByIDOKCode int = 200

/*GetWashByIDOK OK

swagger:response getWashByIdOK
*/
type GetWashByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Wash `json:"body,omitempty"`
}

// NewGetWashByIDOK creates GetWashByIDOK with default headers values
func NewGetWashByIDOK() *GetWashByIDOK {

	return &GetWashByIDOK{}
}

// WithPayload adds the payload to the get wash by Id o k response
func (o *GetWashByIDOK) WithPayload(payload *models.Wash) *GetWashByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get wash by Id o k response
func (o *GetWashByIDOK) SetPayload(payload *models.Wash) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWashByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetWashByIDOK) GetWashByIDResponder() {}

/*GetWashByIDDefault Generic error response

swagger:response getWashByIdDefault
*/
type GetWashByIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetWashByIDDefault creates GetWashByIDDefault with default headers values
func NewGetWashByIDDefault(code int) *GetWashByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetWashByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get wash by Id default response
func (o *GetWashByIDDefault) WithStatusCode(code int) *GetWashByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get wash by Id default response
func (o *GetWashByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get wash by Id default response
func (o *GetWashByIDDefault) WithPayload(payload *models.Error) *GetWashByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get wash by Id default response
func (o *GetWashByIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWashByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetWashByIDDefault) GetWashByIDResponder() {}

type GetWashByIDNotImplementedResponder struct {
	middleware.Responder
}

func (*GetWashByIDNotImplementedResponder) GetWashByIDResponder() {}

func GetWashByIDNotImplemented() GetWashByIDResponder {
	return &GetWashByIDNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetWashByID has not yet been implemented",
		),
	}
}

type GetWashByIDResponder interface {
	middleware.Responder
	GetWashByIDResponder()
}