// Code generated by go-swagger; DO NOT EDIT.

package wash

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"sbp/openapi/models"
)

// CreateOKCode is the HTTP code returned for type CreateOK
const CreateOKCode int = 200

/*
CreateOK Success creation

swagger:response createOK
*/
type CreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.Wash `json:"body,omitempty"`
}

// NewCreateOK creates CreateOK with default headers values
func NewCreateOK() *CreateOK {

	return &CreateOK{}
}

// WithPayload adds the payload to the create o k response
func (o *CreateOK) WithPayload(payload *models.Wash) *CreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create o k response
func (o *CreateOK) SetPayload(payload *models.Wash) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateOK) CreateResponder() {}

// CreateBadRequestCode is the HTTP code returned for type CreateBadRequest
const CreateBadRequestCode int = 400

/*
CreateBadRequest Bad request

swagger:response createBadRequest
*/
type CreateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBadRequest creates CreateBadRequest with default headers values
func NewCreateBadRequest() *CreateBadRequest {

	return &CreateBadRequest{}
}

// WithPayload adds the payload to the create bad request response
func (o *CreateBadRequest) WithPayload(payload *models.Error) *CreateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create bad request response
func (o *CreateBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateBadRequest) CreateResponder() {}

// CreateInternalServerErrorCode is the HTTP code returned for type CreateInternalServerError
const CreateInternalServerErrorCode int = 500

/*
CreateInternalServerError Internal error

swagger:response createInternalServerError
*/
type CreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateInternalServerError creates CreateInternalServerError with default headers values
func NewCreateInternalServerError() *CreateInternalServerError {

	return &CreateInternalServerError{}
}

// WithPayload adds the payload to the create internal server error response
func (o *CreateInternalServerError) WithPayload(payload *models.Error) *CreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create internal server error response
func (o *CreateInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateInternalServerError) CreateResponder() {}

type CreateNotImplementedResponder struct {
	middleware.Responder
}

func (*CreateNotImplementedResponder) CreateResponder() {}

func CreateNotImplemented() CreateResponder {
	return &CreateNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Create has not yet been implemented",
		),
	}
}

type CreateResponder interface {
	middleware.Responder
	CreateResponder()
}