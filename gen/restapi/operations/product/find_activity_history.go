// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"testrunsystem/gen/models"
)

// FindActivityHistoryHandlerFunc turns a function with the right signature into a find activity history handler
type FindActivityHistoryHandlerFunc func(FindActivityHistoryParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn FindActivityHistoryHandlerFunc) Handle(params FindActivityHistoryParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// FindActivityHistoryHandler interface for that can handle valid find activity history params
type FindActivityHistoryHandler interface {
	Handle(FindActivityHistoryParams, *models.Principal) middleware.Responder
}

// NewFindActivityHistory creates a new http.Handler for the find activity history operation
func NewFindActivityHistory(ctx *middleware.Context, handler FindActivityHistoryHandler) *FindActivityHistory {
	return &FindActivityHistory{Context: ctx, Handler: handler}
}

/*
	FindActivityHistory swagger:route GET /v1/product/{product_id}/activity-history product findActivityHistory

# Find Activity History

Find product activity history using pagination
*/
type FindActivityHistory struct {
	Context *middleware.Context
	Handler FindActivityHistoryHandler
}

func (o *FindActivityHistory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewFindActivityHistoryParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// FindActivityHistoryDefaultBody find activity history default body
//
// swagger:model FindActivityHistoryDefaultBody
type FindActivityHistoryDefaultBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	// Example: error
	Message string `json:"message,omitempty"`
}

// Validate validates this find activity history default body
func (o *FindActivityHistoryDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this find activity history default body based on context it is used
func (o *FindActivityHistoryDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *FindActivityHistoryDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FindActivityHistoryDefaultBody) UnmarshalBinary(b []byte) error {
	var res FindActivityHistoryDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// FindActivityHistoryOKBody find activity history o k body
//
// swagger:model FindActivityHistoryOKBody
type FindActivityHistoryOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// data
	Data []*FindActivityHistoryOKBodyDataItems0 `json:"data"`

	// metadata
	Metadata *FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata `json:"metadata,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *FindActivityHistoryOKBody) UnmarshalJSON(raw []byte) error {
	// FindActivityHistoryOKBodyAO0
	var dataFindActivityHistoryOKBodyAO0 struct {
		Message string `json:"message,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataFindActivityHistoryOKBodyAO0); err != nil {
		return err
	}

	o.Message = dataFindActivityHistoryOKBodyAO0.Message

	// FindActivityHistoryOKBodyAO1
	var dataFindActivityHistoryOKBodyAO1 struct {
		Data []*FindActivityHistoryOKBodyDataItems0 `json:"data"`

		Metadata *FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata `json:"metadata,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataFindActivityHistoryOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataFindActivityHistoryOKBodyAO1.Data

	o.Metadata = dataFindActivityHistoryOKBodyAO1.Metadata

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o FindActivityHistoryOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataFindActivityHistoryOKBodyAO0 struct {
		Message string `json:"message,omitempty"`
	}

	dataFindActivityHistoryOKBodyAO0.Message = o.Message

	jsonDataFindActivityHistoryOKBodyAO0, errFindActivityHistoryOKBodyAO0 := swag.WriteJSON(dataFindActivityHistoryOKBodyAO0)
	if errFindActivityHistoryOKBodyAO0 != nil {
		return nil, errFindActivityHistoryOKBodyAO0
	}
	_parts = append(_parts, jsonDataFindActivityHistoryOKBodyAO0)
	var dataFindActivityHistoryOKBodyAO1 struct {
		Data []*FindActivityHistoryOKBodyDataItems0 `json:"data"`

		Metadata *FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata `json:"metadata,omitempty"`
	}

	dataFindActivityHistoryOKBodyAO1.Data = o.Data

	dataFindActivityHistoryOKBodyAO1.Metadata = o.Metadata

	jsonDataFindActivityHistoryOKBodyAO1, errFindActivityHistoryOKBodyAO1 := swag.WriteJSON(dataFindActivityHistoryOKBodyAO1)
	if errFindActivityHistoryOKBodyAO1 != nil {
		return nil, errFindActivityHistoryOKBodyAO1
	}
	_parts = append(_parts, jsonDataFindActivityHistoryOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this find activity history o k body
func (o *FindActivityHistoryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindActivityHistoryOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("findActivityHistoryOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("findActivityHistoryOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *FindActivityHistoryOKBody) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(o.Metadata) { // not required
		return nil
	}

	if o.Metadata != nil {
		if err := o.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("findActivityHistoryOK" + "." + "metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("findActivityHistoryOK" + "." + "metadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this find activity history o k body based on the context it is used
func (o *FindActivityHistoryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindActivityHistoryOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("findActivityHistoryOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("findActivityHistoryOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *FindActivityHistoryOKBody) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if o.Metadata != nil {
		if err := o.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("findActivityHistoryOK" + "." + "metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("findActivityHistoryOK" + "." + "metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *FindActivityHistoryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FindActivityHistoryOKBody) UnmarshalBinary(b []byte) error {
	var res FindActivityHistoryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// FindActivityHistoryOKBodyDataItems0 find activity history o k body data items0
//
// swagger:model FindActivityHistoryOKBodyDataItems0
type FindActivityHistoryOKBodyDataItems0 struct {

	// id
	// Required: true
	ID *uint64 `json:"id"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at"`

	// created by
	CreatedBy string `json:"created_by"`

	// note
	Note string `json:"note,omitempty"`

	// product id
	ProductID uint64 `json:"product_id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *FindActivityHistoryOKBodyDataItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ID *uint64 `json:"id"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	o.ID = dataAO0.ID

	// AO1
	var dataAO1 struct {
		CreatedAt strfmt.DateTime `json:"created_at"`

		CreatedBy string `json:"created_by"`

		Note string `json:"note,omitempty"`

		ProductID uint64 `json:"product_id,omitempty"`

		Type string `json:"type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	o.CreatedAt = dataAO1.CreatedAt

	o.CreatedBy = dataAO1.CreatedBy

	o.Note = dataAO1.Note

	o.ProductID = dataAO1.ProductID

	o.Type = dataAO1.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o FindActivityHistoryOKBodyDataItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ID *uint64 `json:"id"`
	}

	dataAO0.ID = o.ID

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	var dataAO1 struct {
		CreatedAt strfmt.DateTime `json:"created_at"`

		CreatedBy string `json:"created_by"`

		Note string `json:"note,omitempty"`

		ProductID uint64 `json:"product_id,omitempty"`

		Type string `json:"type,omitempty"`
	}

	dataAO1.CreatedAt = o.CreatedAt

	dataAO1.CreatedBy = o.CreatedBy

	dataAO1.Note = o.Note

	dataAO1.ProductID = o.ProductID

	dataAO1.Type = o.Type

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this find activity history o k body data items0
func (o *FindActivityHistoryOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindActivityHistoryOKBodyDataItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *FindActivityHistoryOKBodyDataItems0) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this find activity history o k body data items0 based on context it is used
func (o *FindActivityHistoryOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *FindActivityHistoryOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FindActivityHistoryOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res FindActivityHistoryOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata find activity history o k body find activity history o k body a o1 metadata
//
// swagger:model FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata
type FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata struct {

	// page
	Page int64 `json:"page,omitempty"`

	// per page
	PerPage int64 `json:"per_page,omitempty"`

	// total page
	TotalPage int64 `json:"total_page,omitempty"`

	// total row
	TotalRow int64 `json:"total_row,omitempty"`
}

// Validate validates this find activity history o k body find activity history o k body a o1 metadata
func (o *FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this find activity history o k body find activity history o k body a o1 metadata based on context it is used
func (o *FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata) UnmarshalBinary(b []byte) error {
	var res FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
