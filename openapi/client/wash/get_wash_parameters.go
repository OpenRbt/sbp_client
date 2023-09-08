// Code generated by go-swagger; DO NOT EDIT.

package wash

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetWashParams creates a new GetWashParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWashParams() *GetWashParams {
	return &GetWashParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWashParamsWithTimeout creates a new GetWashParams object
// with the ability to set a timeout on a request.
func NewGetWashParamsWithTimeout(timeout time.Duration) *GetWashParams {
	return &GetWashParams{
		timeout: timeout,
	}
}

// NewGetWashParamsWithContext creates a new GetWashParams object
// with the ability to set a context for a request.
func NewGetWashParamsWithContext(ctx context.Context) *GetWashParams {
	return &GetWashParams{
		Context: ctx,
	}
}

// NewGetWashParamsWithHTTPClient creates a new GetWashParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWashParamsWithHTTPClient(client *http.Client) *GetWashParams {
	return &GetWashParams{
		HTTPClient: client,
	}
}

/*
GetWashParams contains all the parameters to send to the API endpoint

	for the get wash operation.

	Typically these are written to a http.Request.
*/
type GetWashParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get wash params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWashParams) WithDefaults() *GetWashParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get wash params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWashParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get wash params
func (o *GetWashParams) WithTimeout(timeout time.Duration) *GetWashParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wash params
func (o *GetWashParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wash params
func (o *GetWashParams) WithContext(ctx context.Context) *GetWashParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wash params
func (o *GetWashParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wash params
func (o *GetWashParams) WithHTTPClient(client *http.Client) *GetWashParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wash params
func (o *GetWashParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get wash params
func (o *GetWashParams) WithID(id string) *GetWashParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get wash params
func (o *GetWashParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetWashParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}