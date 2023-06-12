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

// PostAddReader is a Reader for the PostAdd structure.
type PostAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAddBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAddOK creates a PostAddOK with default headers values
func NewPostAddOK() *PostAddOK {
	return &PostAddOK{}
}

/*
PostAddOK describes a response with status code 200, with default header values.

Success creation
*/
type PostAddOK struct {
	Payload *models.WashPost
}

// IsSuccess returns true when this post add o k response has a 2xx status code
func (o *PostAddOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post add o k response has a 3xx status code
func (o *PostAddOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post add o k response has a 4xx status code
func (o *PostAddOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post add o k response has a 5xx status code
func (o *PostAddOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post add o k response a status code equal to that given
func (o *PostAddOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post add o k response
func (o *PostAddOK) Code() int {
	return 200
}

func (o *PostAddOK) Error() string {
	return fmt.Sprintf("[PUT /wash-post/][%d] postAddOK  %+v", 200, o.Payload)
}

func (o *PostAddOK) String() string {
	return fmt.Sprintf("[PUT /wash-post/][%d] postAddOK  %+v", 200, o.Payload)
}

func (o *PostAddOK) GetPayload() *models.WashPost {
	return o.Payload
}

func (o *PostAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WashPost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddBadRequest creates a PostAddBadRequest with default headers values
func NewPostAddBadRequest() *PostAddBadRequest {
	return &PostAddBadRequest{}
}

/*
PostAddBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostAddBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post add bad request response has a 2xx status code
func (o *PostAddBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post add bad request response has a 3xx status code
func (o *PostAddBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post add bad request response has a 4xx status code
func (o *PostAddBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post add bad request response has a 5xx status code
func (o *PostAddBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post add bad request response a status code equal to that given
func (o *PostAddBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post add bad request response
func (o *PostAddBadRequest) Code() int {
	return 400
}

func (o *PostAddBadRequest) Error() string {
	return fmt.Sprintf("[PUT /wash-post/][%d] postAddBadRequest  %+v", 400, o.Payload)
}

func (o *PostAddBadRequest) String() string {
	return fmt.Sprintf("[PUT /wash-post/][%d] postAddBadRequest  %+v", 400, o.Payload)
}

func (o *PostAddBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostAddBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddInternalServerError creates a PostAddInternalServerError with default headers values
func NewPostAddInternalServerError() *PostAddInternalServerError {
	return &PostAddInternalServerError{}
}

/*
PostAddInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type PostAddInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post add internal server error response has a 2xx status code
func (o *PostAddInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post add internal server error response has a 3xx status code
func (o *PostAddInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post add internal server error response has a 4xx status code
func (o *PostAddInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post add internal server error response has a 5xx status code
func (o *PostAddInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post add internal server error response a status code equal to that given
func (o *PostAddInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post add internal server error response
func (o *PostAddInternalServerError) Code() int {
	return 500
}

func (o *PostAddInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /wash-post/][%d] postAddInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAddInternalServerError) String() string {
	return fmt.Sprintf("[PUT /wash-post/][%d] postAddInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAddInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostAddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
