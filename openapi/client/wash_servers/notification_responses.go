// Code generated by go-swagger; DO NOT EDIT.

package wash_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// NotificationReader is a Reader for the Notification structure.
type NotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationOK creates a NotificationOK with default headers values
func NewNotificationOK() *NotificationOK {
	return &NotificationOK{}
}

/*
NotificationOK describes a response with status code 200, with default header values.

OK
*/
type NotificationOK struct {
	Payload string
}

// IsSuccess returns true when this notification o k response has a 2xx status code
func (o *NotificationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notification o k response has a 3xx status code
func (o *NotificationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification o k response has a 4xx status code
func (o *NotificationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notification o k response has a 5xx status code
func (o *NotificationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notification o k response a status code equal to that given
func (o *NotificationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the notification o k response
func (o *NotificationOK) Code() int {
	return 200
}

func (o *NotificationOK) Error() string {
	return fmt.Sprintf("[POST /Notification][%d] notificationOK  %+v", 200, o.Payload)
}

func (o *NotificationOK) String() string {
	return fmt.Sprintf("[POST /Notification][%d] notificationOK  %+v", 200, o.Payload)
}

func (o *NotificationOK) GetPayload() string {
	return o.Payload
}

func (o *NotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
