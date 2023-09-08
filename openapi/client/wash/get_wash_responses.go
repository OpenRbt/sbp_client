// Code generated by go-swagger; DO NOT EDIT.

package wash

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"sbp/openapi/models"
)

// GetWashReader is a Reader for the GetWash structure.
type GetWashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWashBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWashInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /wash/{id}] getWash", response, response.Code())
	}
}

// NewGetWashOK creates a GetWashOK with default headers values
func NewGetWashOK() *GetWashOK {
	return &GetWashOK{}
}

/*
GetWashOK describes a response with status code 200, with default header values.

OK
*/
type GetWashOK struct {
	Payload *models.Wash
}

// IsSuccess returns true when this get wash o k response has a 2xx status code
func (o *GetWashOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get wash o k response has a 3xx status code
func (o *GetWashOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get wash o k response has a 4xx status code
func (o *GetWashOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get wash o k response has a 5xx status code
func (o *GetWashOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get wash o k response a status code equal to that given
func (o *GetWashOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get wash o k response
func (o *GetWashOK) Code() int {
	return 200
}

func (o *GetWashOK) Error() string {
	return fmt.Sprintf("[GET /wash/{id}][%d] getWashOK  %+v", 200, o.Payload)
}

func (o *GetWashOK) String() string {
	return fmt.Sprintf("[GET /wash/{id}][%d] getWashOK  %+v", 200, o.Payload)
}

func (o *GetWashOK) GetPayload() *models.Wash {
	return o.Payload
}

func (o *GetWashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Wash)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWashBadRequest creates a GetWashBadRequest with default headers values
func NewGetWashBadRequest() *GetWashBadRequest {
	return &GetWashBadRequest{}
}

/*
GetWashBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetWashBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get wash bad request response has a 2xx status code
func (o *GetWashBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get wash bad request response has a 3xx status code
func (o *GetWashBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get wash bad request response has a 4xx status code
func (o *GetWashBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get wash bad request response has a 5xx status code
func (o *GetWashBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get wash bad request response a status code equal to that given
func (o *GetWashBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get wash bad request response
func (o *GetWashBadRequest) Code() int {
	return 400
}

func (o *GetWashBadRequest) Error() string {
	return fmt.Sprintf("[GET /wash/{id}][%d] getWashBadRequest  %+v", 400, o.Payload)
}

func (o *GetWashBadRequest) String() string {
	return fmt.Sprintf("[GET /wash/{id}][%d] getWashBadRequest  %+v", 400, o.Payload)
}

func (o *GetWashBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWashBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWashNotFound creates a GetWashNotFound with default headers values
func NewGetWashNotFound() *GetWashNotFound {
	return &GetWashNotFound{}
}

/*
GetWashNotFound describes a response with status code 404, with default header values.

Wash not exists
*/
type GetWashNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get wash not found response has a 2xx status code
func (o *GetWashNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get wash not found response has a 3xx status code
func (o *GetWashNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get wash not found response has a 4xx status code
func (o *GetWashNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get wash not found response has a 5xx status code
func (o *GetWashNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get wash not found response a status code equal to that given
func (o *GetWashNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get wash not found response
func (o *GetWashNotFound) Code() int {
	return 404
}

func (o *GetWashNotFound) Error() string {
	return fmt.Sprintf("[GET /wash/{id}][%d] getWashNotFound  %+v", 404, o.Payload)
}

func (o *GetWashNotFound) String() string {
	return fmt.Sprintf("[GET /wash/{id}][%d] getWashNotFound  %+v", 404, o.Payload)
}

func (o *GetWashNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWashInternalServerError creates a GetWashInternalServerError with default headers values
func NewGetWashInternalServerError() *GetWashInternalServerError {
	return &GetWashInternalServerError{}
}

/*
GetWashInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type GetWashInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get wash internal server error response has a 2xx status code
func (o *GetWashInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get wash internal server error response has a 3xx status code
func (o *GetWashInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get wash internal server error response has a 4xx status code
func (o *GetWashInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get wash internal server error response has a 5xx status code
func (o *GetWashInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get wash internal server error response a status code equal to that given
func (o *GetWashInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get wash internal server error response
func (o *GetWashInternalServerError) Code() int {
	return 500
}

func (o *GetWashInternalServerError) Error() string {
	return fmt.Sprintf("[GET /wash/{id}][%d] getWashInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWashInternalServerError) String() string {
	return fmt.Sprintf("[GET /wash/{id}][%d] getWashInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWashInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWashInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
