// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WashServerAdd wash server add
//
// swagger:model WashServerAdd
type WashServerAdd struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// terminal key
	TerminalKey string `json:"terminal_key,omitempty"`

	// terminal password
	TerminalPassword string `json:"terminal_password,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *WashServerAdd) UnmarshalJSON(data []byte) error {
	var props struct {

		// description
		Description string `json:"description,omitempty"`

		// name
		// Required: true
		Name *string `json:"name"`

		// terminal key
		TerminalKey string `json:"terminal_key,omitempty"`

		// terminal password
		TerminalPassword string `json:"terminal_password,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Description = props.Description
	m.Name = props.Name
	m.TerminalKey = props.TerminalKey
	m.TerminalPassword = props.TerminalPassword
	return nil
}

// Validate validates this wash server add
func (m *WashServerAdd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WashServerAdd) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wash server add based on context it is used
func (m *WashServerAdd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WashServerAdd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WashServerAdd) UnmarshalBinary(b []byte) error {
	var res WashServerAdd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
