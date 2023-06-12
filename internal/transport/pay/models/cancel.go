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

// Cancel cancel
//
// swagger:model Cancel
type Cancel struct {

	// payment Id
	PaymentID string `json:"PaymentId,omitempty"`

	// terminal key
	TerminalKey string `json:"TerminalKey,omitempty"`

	// token
	Token string `json:"Token,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Cancel) UnmarshalJSON(data []byte) error {
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

// Validate validates this cancel
func (m *Cancel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cancel based on context it is used
func (m *Cancel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Cancel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cancel) UnmarshalBinary(b []byte) error {
	var res Cancel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
