// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Cancel cancel
//
// swagger:model cancel
type Cancel struct {

	// order ID
	OrderID string `json:"orderID,omitempty"`

	// post ID
	PostID string `json:"postID,omitempty"`

	// wash ID
	WashID string `json:"washID,omitempty"`
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
