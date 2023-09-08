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

// DeleteNoContentCode is the HTTP code returned for type DeleteNoContent
const DeleteNoContentCode int = 204

/*
DeleteNoContent OK

swagger:response deleteNoContent
*/
type DeleteNoContent struct {
}

// NewDeleteNoContent creates DeleteNoContent with default headers values
func NewDeleteNoContent() *DeleteNoContent {

	return &DeleteNoContent{}
}

// WriteResponse to the client
func (o *DeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *DeleteNoContent) DeleteResponder() {}

// DeleteBadRequestCode is the HTTP code returned for type DeleteBadRequest
const DeleteBadRequestCode int = 400

/*
DeleteBadRequest Bad request

swagger:response deleteBadRequest
*/
type DeleteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteBadRequest creates DeleteBadRequest with default headers values
func NewDeleteBadRequest() *DeleteBadRequest {

	return &DeleteBadRequest{}
}

// WithPayload adds the payload to the delete bad request response
func (o *DeleteBadRequest) WithPayload(payload *models.Error) *DeleteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete bad request response
func (o *DeleteBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteBadRequest) DeleteResponder() {}

// DeleteForbiddenCode is the HTTP code returned for type DeleteForbidden
const DeleteForbiddenCode int = 403

/*
DeleteForbidden Access denied

swagger:response deleteForbidden
*/
type DeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteForbidden creates DeleteForbidden with default headers values
func NewDeleteForbidden() *DeleteForbidden {

	return &DeleteForbidden{}
}

// WithPayload adds the payload to the delete forbidden response
func (o *DeleteForbidden) WithPayload(payload *models.Error) *DeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete forbidden response
func (o *DeleteForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteForbidden) DeleteResponder() {}

// DeleteNotFoundCode is the HTTP code returned for type DeleteNotFound
const DeleteNotFoundCode int = 404

/*
DeleteNotFound Wash not exists

swagger:response deleteNotFound
*/
type DeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteNotFound creates DeleteNotFound with default headers values
func NewDeleteNotFound() *DeleteNotFound {

	return &DeleteNotFound{}
}

// WithPayload adds the payload to the delete not found response
func (o *DeleteNotFound) WithPayload(payload *models.Error) *DeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete not found response
func (o *DeleteNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteNotFound) DeleteResponder() {}

// DeleteInternalServerErrorCode is the HTTP code returned for type DeleteInternalServerError
const DeleteInternalServerErrorCode int = 500

/*
DeleteInternalServerError Internal error

swagger:response deleteInternalServerError
*/
type DeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteInternalServerError creates DeleteInternalServerError with default headers values
func NewDeleteInternalServerError() *DeleteInternalServerError {

	return &DeleteInternalServerError{}
}

// WithPayload adds the payload to the delete internal server error response
func (o *DeleteInternalServerError) WithPayload(payload *models.Error) *DeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete internal server error response
func (o *DeleteInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteInternalServerError) DeleteResponder() {}

type DeleteNotImplementedResponder struct {
	middleware.Responder
}

func (*DeleteNotImplementedResponder) DeleteResponder() {}

func DeleteNotImplemented() DeleteResponder {
	return &DeleteNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Delete has not yet been implemented",
		),
	}
}

type DeleteResponder interface {
	middleware.Responder
	DeleteResponder()
}
