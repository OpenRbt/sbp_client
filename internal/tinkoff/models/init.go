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

// Init init
//
// swagger:model Init
type Init struct {

	// amount
	// Required: true
	Amount *int64 `json:"Amount"`

	// order Id
	// Required: true
	OrderID *string `json:"OrderId"`

	// redirect due date
	RedirectDueDate string `json:"RedirectDueDate,omitempty"`

	// terminal key
	// Required: true
	TerminalKey *string `json:"TerminalKey"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Init) UnmarshalJSON(data []byte) error {
	var props struct {

		// amount
		// Required: true
		Amount *int64 `json:"Amount"`

		// order Id
		// Required: true
		OrderID *string `json:"OrderId"`

		// redirect due date
		RedirectDueDate string `json:"RedirectDueDate,omitempty"`

		// terminal key
		// Required: true
		TerminalKey *string `json:"TerminalKey"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Amount = props.Amount
	m.OrderID = props.OrderID
	m.RedirectDueDate = props.RedirectDueDate
	m.TerminalKey = props.TerminalKey
	return nil
}

// Validate validates this init
func (m *Init) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminalKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Init) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("Amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *Init) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("OrderId", "body", m.OrderID); err != nil {
		return err
	}

	return nil
}

func (m *Init) validateTerminalKey(formats strfmt.Registry) error {

	if err := validate.Required("TerminalKey", "body", m.TerminalKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this init based on context it is used
func (m *Init) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Init) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Init) UnmarshalBinary(b []byte) error {
	var res Init
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
