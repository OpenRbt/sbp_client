// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseInit response init
//
// swagger:model ResponseInit
type ResponseInit struct {

	// amount
	Amount int64 `json:"Amount,omitempty"`

	// details
	Details string `json:"Details,omitempty"`

	// error code
	ErrorCode string `json:"ErrorCode,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// order Id
	OrderID string `json:"OrderId,omitempty"`

	// payment Id
	PaymentID string `json:"PaymentId,omitempty"`

	// payment URL
	PaymentURL string `json:"PaymentURL,omitempty"`

	// status
	Status string `json:"Status,omitempty"`

	// success
	Success bool `json:"Success,omitempty"`

	// terminal key
	TerminalKey string `json:"TerminalKey,omitempty"`
}

// Validate validates this response init
func (m *ResponseInit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this response init based on context it is used
func (m *ResponseInit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseInit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseInit) UnmarshalBinary(b []byte) error {
	var res ResponseInit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
