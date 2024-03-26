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

// UpdateWashNoContentCode is the HTTP code returned for type UpdateWashNoContent
const UpdateWashNoContentCode int = 204

/*
UpdateWashNoContent Success update

swagger:response updateWashNoContent
*/
type UpdateWashNoContent struct {
}

// NewUpdateWashNoContent creates UpdateWashNoContent with default headers values
func NewUpdateWashNoContent() *UpdateWashNoContent {

	return &UpdateWashNoContent{}
}

// WriteResponse to the client
func (o *UpdateWashNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *UpdateWashNoContent) UpdateWashResponder() {}

/*
UpdateWashDefault Generic error response

swagger:response updateWashDefault
*/
type UpdateWashDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateWashDefault creates UpdateWashDefault with default headers values
func NewUpdateWashDefault(code int) *UpdateWashDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateWashDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update wash default response
func (o *UpdateWashDefault) WithStatusCode(code int) *UpdateWashDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update wash default response
func (o *UpdateWashDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update wash default response
func (o *UpdateWashDefault) WithPayload(payload *models.Error) *UpdateWashDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update wash default response
func (o *UpdateWashDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateWashDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateWashDefault) UpdateWashResponder() {}

type UpdateWashNotImplementedResponder struct {
	middleware.Responder
}

func (*UpdateWashNotImplementedResponder) UpdateWashResponder() {}

func UpdateWashNotImplemented() UpdateWashResponder {
	return &UpdateWashNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.UpdateWash has not yet been implemented",
		),
	}
}

type UpdateWashResponder interface {
	middleware.Responder
	UpdateWashResponder()
}
