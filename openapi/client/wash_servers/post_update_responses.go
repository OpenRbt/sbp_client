// Code generated by go-swagger; DO NOT EDIT.

package wash_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"sbp/openapi/models"
)

// PostUpdateReader is a Reader for the PostUpdate structure.
type PostUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUpdateNoContent creates a PostUpdateNoContent with default headers values
func NewPostUpdateNoContent() *PostUpdateNoContent {
	return &PostUpdateNoContent{}
}

/*
PostUpdateNoContent describes a response with status code 204, with default header values.

Success update
*/
type PostUpdateNoContent struct {
}

// IsSuccess returns true when this post update no content response has a 2xx status code
func (o *PostUpdateNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post update no content response has a 3xx status code
func (o *PostUpdateNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post update no content response has a 4xx status code
func (o *PostUpdateNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post update no content response has a 5xx status code
func (o *PostUpdateNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post update no content response a status code equal to that given
func (o *PostUpdateNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the post update no content response
func (o *PostUpdateNoContent) Code() int {
	return 204
}

func (o *PostUpdateNoContent) Error() string {
	return fmt.Sprintf("[PATCH /wash-post/][%d] postUpdateNoContent ", 204)
}

func (o *PostUpdateNoContent) String() string {
	return fmt.Sprintf("[PATCH /wash-post/][%d] postUpdateNoContent ", 204)
}

func (o *PostUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUpdateBadRequest creates a PostUpdateBadRequest with default headers values
func NewPostUpdateBadRequest() *PostUpdateBadRequest {
	return &PostUpdateBadRequest{}
}

/*
PostUpdateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostUpdateBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post update bad request response has a 2xx status code
func (o *PostUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post update bad request response has a 3xx status code
func (o *PostUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post update bad request response has a 4xx status code
func (o *PostUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post update bad request response has a 5xx status code
func (o *PostUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post update bad request response a status code equal to that given
func (o *PostUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post update bad request response
func (o *PostUpdateBadRequest) Code() int {
	return 400
}

func (o *PostUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /wash-post/][%d] postUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PostUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /wash-post/][%d] postUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PostUpdateBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUpdateNotFound creates a PostUpdateNotFound with default headers values
func NewPostUpdateNotFound() *PostUpdateNotFound {
	return &PostUpdateNotFound{}
}

/*
PostUpdateNotFound describes a response with status code 404, with default header values.

WashServer not exists
*/
type PostUpdateNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this post update not found response has a 2xx status code
func (o *PostUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post update not found response has a 3xx status code
func (o *PostUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post update not found response has a 4xx status code
func (o *PostUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post update not found response has a 5xx status code
func (o *PostUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post update not found response a status code equal to that given
func (o *PostUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post update not found response
func (o *PostUpdateNotFound) Code() int {
	return 404
}

func (o *PostUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /wash-post/][%d] postUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PostUpdateNotFound) String() string {
	return fmt.Sprintf("[PATCH /wash-post/][%d] postUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PostUpdateNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUpdateInternalServerError creates a PostUpdateInternalServerError with default headers values
func NewPostUpdateInternalServerError() *PostUpdateInternalServerError {
	return &PostUpdateInternalServerError{}
}

/*
PostUpdateInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type PostUpdateInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post update internal server error response has a 2xx status code
func (o *PostUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post update internal server error response has a 3xx status code
func (o *PostUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post update internal server error response has a 4xx status code
func (o *PostUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post update internal server error response has a 5xx status code
func (o *PostUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post update internal server error response a status code equal to that given
func (o *PostUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post update internal server error response
func (o *PostUpdateInternalServerError) Code() int {
	return 500
}

func (o *PostUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /wash-post/][%d] postUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /wash-post/][%d] postUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUpdateInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
