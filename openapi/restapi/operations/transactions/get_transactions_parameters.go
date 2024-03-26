// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetTransactionsParams creates a new GetTransactionsParams object
// with the default values initialized.
func NewGetTransactionsParams() GetTransactionsParams {

	var (
		// initialize parameters with default values

		pageDefault     = int64(1)
		pageSizeDefault = int64(10)
	)

	return GetTransactionsParams{
		Page: &pageDefault,

		PageSize: &pageSizeDefault,
	}
}

// GetTransactionsParams contains all the bound params for the get transactions operation
// typically these are obtained from a http.Request
//
// swagger:parameters getTransactions
type GetTransactionsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	GroupID *strfmt.UUID
	/*
	  In: query
	*/
	OrganizationID *strfmt.UUID
	/*
	  Minimum: 1
	  In: query
	  Default: 1
	*/
	Page *int64
	/*
	  Maximum: 100
	  Minimum: 1
	  In: query
	  Default: 10
	*/
	PageSize *int64
	/*
	  In: query
	*/
	PostID *int64
	/*
	  In: query
	*/
	Status *string
	/*
	  In: query
	*/
	WashID *strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTransactionsParams() beforehand.
func (o *GetTransactionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qGroupID, qhkGroupID, _ := qs.GetOK("groupId")
	if err := o.bindGroupID(qGroupID, qhkGroupID, route.Formats); err != nil {
		res = append(res, err)
	}

	qOrganizationID, qhkOrganizationID, _ := qs.GetOK("organizationId")
	if err := o.bindOrganizationID(qOrganizationID, qhkOrganizationID, route.Formats); err != nil {
		res = append(res, err)
	}

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	qPostID, qhkPostID, _ := qs.GetOK("postId")
	if err := o.bindPostID(qPostID, qhkPostID, route.Formats); err != nil {
		res = append(res, err)
	}

	qStatus, qhkStatus, _ := qs.GetOK("status")
	if err := o.bindStatus(qStatus, qhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	qWashID, qhkWashID, _ := qs.GetOK("washId")
	if err := o.bindWashID(qWashID, qhkWashID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGroupID binds and validates parameter GroupID from query.
func (o *GetTransactionsParams) bindGroupID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("groupId", "query", "strfmt.UUID", raw)
	}
	o.GroupID = (value.(*strfmt.UUID))

	if err := o.validateGroupID(formats); err != nil {
		return err
	}

	return nil
}

// validateGroupID carries on validations for parameter GroupID
func (o *GetTransactionsParams) validateGroupID(formats strfmt.Registry) error {

	if err := validate.FormatOf("groupId", "query", "uuid", o.GroupID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindOrganizationID binds and validates parameter OrganizationID from query.
func (o *GetTransactionsParams) bindOrganizationID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("organizationId", "query", "strfmt.UUID", raw)
	}
	o.OrganizationID = (value.(*strfmt.UUID))

	if err := o.validateOrganizationID(formats); err != nil {
		return err
	}

	return nil
}

// validateOrganizationID carries on validations for parameter OrganizationID
func (o *GetTransactionsParams) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.FormatOf("organizationId", "query", "uuid", o.OrganizationID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *GetTransactionsParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetTransactionsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = &value

	if err := o.validatePage(formats); err != nil {
		return err
	}

	return nil
}

// validatePage carries on validations for parameter Page
func (o *GetTransactionsParams) validatePage(formats strfmt.Registry) error {

	if err := validate.MinimumInt("page", "query", *o.Page, 1, false); err != nil {
		return err
	}

	return nil
}

// bindPageSize binds and validates parameter PageSize from query.
func (o *GetTransactionsParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetTransactionsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int64", raw)
	}
	o.PageSize = &value

	if err := o.validatePageSize(formats); err != nil {
		return err
	}

	return nil
}

// validatePageSize carries on validations for parameter PageSize
func (o *GetTransactionsParams) validatePageSize(formats strfmt.Registry) error {

	if err := validate.MinimumInt("pageSize", "query", *o.PageSize, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("pageSize", "query", *o.PageSize, 100, false); err != nil {
		return err
	}

	return nil
}

// bindPostID binds and validates parameter PostID from query.
func (o *GetTransactionsParams) bindPostID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("postId", "query", "int64", raw)
	}
	o.PostID = &value

	return nil
}

// bindStatus binds and validates parameter Status from query.
func (o *GetTransactionsParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Status = &raw

	if err := o.validateStatus(formats); err != nil {
		return err
	}

	return nil
}

// validateStatus carries on validations for parameter Status
func (o *GetTransactionsParams) validateStatus(formats strfmt.Registry) error {

	if err := validate.EnumCase("status", "query", *o.Status, []interface{}{"new", "authorized", "confirmed_not_synced", "confirmed", "canceling", "canceled", "refunded", "unknown"}, true); err != nil {
		return err
	}

	return nil
}

// bindWashID binds and validates parameter WashID from query.
func (o *GetTransactionsParams) bindWashID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("washId", "query", "strfmt.UUID", raw)
	}
	o.WashID = (value.(*strfmt.UUID))

	if err := o.validateWashID(formats); err != nil {
		return err
	}

	return nil
}

// validateWashID carries on validations for parameter WashID
func (o *GetTransactionsParams) validateWashID(formats strfmt.Registry) error {

	if err := validate.FormatOf("washId", "query", "uuid", o.WashID.String(), formats); err != nil {
		return err
	}
	return nil
}