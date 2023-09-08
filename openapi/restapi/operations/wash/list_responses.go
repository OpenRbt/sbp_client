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

// ListOKCode is the HTTP code returned for type ListOK
const ListOKCode int = 200

/*
ListOK OK

swagger:response listOK
*/
type ListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Wash `json:"body,omitempty"`
}

// NewListOK creates ListOK with default headers values
func NewListOK() *ListOK {

	return &ListOK{}
}

// WithPayload adds the payload to the list o k response
func (o *ListOK) WithPayload(payload []*models.Wash) *ListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list o k response
func (o *ListOK) SetPayload(payload []*models.Wash) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Wash, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *ListOK) ListResponder() {}

// ListBadRequestCode is the HTTP code returned for type ListBadRequest
const ListBadRequestCode int = 400

/*
ListBadRequest Bad request

swagger:response listBadRequest
*/
type ListBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListBadRequest creates ListBadRequest with default headers values
func NewListBadRequest() *ListBadRequest {

	return &ListBadRequest{}
}

// WithPayload adds the payload to the list bad request response
func (o *ListBadRequest) WithPayload(payload *models.Error) *ListBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list bad request response
func (o *ListBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *ListBadRequest) ListResponder() {}

// ListForbiddenCode is the HTTP code returned for type ListForbidden
const ListForbiddenCode int = 403

/*
ListForbidden Access denied

swagger:response listForbidden
*/
type ListForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListForbidden creates ListForbidden with default headers values
func NewListForbidden() *ListForbidden {

	return &ListForbidden{}
}

// WithPayload adds the payload to the list forbidden response
func (o *ListForbidden) WithPayload(payload *models.Error) *ListForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list forbidden response
func (o *ListForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *ListForbidden) ListResponder() {}

// ListNotFoundCode is the HTTP code returned for type ListNotFound
const ListNotFoundCode int = 404

/*
ListNotFound Wash not exists

swagger:response listNotFound
*/
type ListNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListNotFound creates ListNotFound with default headers values
func NewListNotFound() *ListNotFound {

	return &ListNotFound{}
}

// WithPayload adds the payload to the list not found response
func (o *ListNotFound) WithPayload(payload *models.Error) *ListNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list not found response
func (o *ListNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *ListNotFound) ListResponder() {}

// ListInternalServerErrorCode is the HTTP code returned for type ListInternalServerError
const ListInternalServerErrorCode int = 500

/*
ListInternalServerError Internal error

swagger:response listInternalServerError
*/
type ListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListInternalServerError creates ListInternalServerError with default headers values
func NewListInternalServerError() *ListInternalServerError {

	return &ListInternalServerError{}
}

// WithPayload adds the payload to the list internal server error response
func (o *ListInternalServerError) WithPayload(payload *models.Error) *ListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list internal server error response
func (o *ListInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *ListInternalServerError) ListResponder() {}

type ListNotImplementedResponder struct {
	middleware.Responder
}

func (*ListNotImplementedResponder) ListResponder() {}

func ListNotImplemented() ListResponder {
	return &ListNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.List has not yet been implemented",
		),
	}
}

type ListResponder interface {
	middleware.Responder
	ListResponder()
}
