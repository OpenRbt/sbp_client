// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"sbp/tinkoffapi/models"
)

// InitReader is a Reader for the Init structure.
type InitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewInitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInitOK creates a InitOK with default headers values
func NewInitOK() *InitOK {
	return &InitOK{}
}

/* InitOK describes a response with status code 200, with default header values.

OK
*/
type InitOK struct {
	Payload *models.ResponseInit
}

func (o *InitOK) Error() string {
	return fmt.Sprintf("[POST /Init/][%d] initOK  %+v", 200, o.Payload)
}
func (o *InitOK) GetPayload() *models.ResponseInit {
	return o.Payload
}

func (o *InitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseInit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInitInternalServerError creates a InitInternalServerError with default headers values
func NewInitInternalServerError() *InitInternalServerError {
	return &InitInternalServerError{}
}

/* InitInternalServerError describes a response with status code 500, with default header values.

error
*/
type InitInternalServerError struct {
}

func (o *InitInternalServerError) Error() string {
	return fmt.Sprintf("[POST /Init/][%d] initInternalServerError ", 500)
}

func (o *InitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}