// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetQr get qr
//
// swagger:model GetQr
type GetQr struct {

	// payment Id
	PaymentID string `json:"PaymentId,omitempty"`

	// terminal key
	TerminalKey string `json:"TerminalKey,omitempty"`

	// token
	Token string `json:"Token,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *GetQr) UnmarshalJSON(data []byte) error {
	var props struct {

		// payment Id
		PaymentID string `json:"PaymentId,omitempty"`

		// terminal key
		TerminalKey string `json:"TerminalKey,omitempty"`

		// token
		Token string `json:"Token,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.PaymentID = props.PaymentID
	m.TerminalKey = props.TerminalKey
	m.Token = props.Token
	return nil
}

// Validate validates this get qr
func (m *GetQr) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get qr based on context it is used
func (m *GetQr) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetQr) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetQr) UnmarshalBinary(b []byte) error {
	var res GetQr
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}