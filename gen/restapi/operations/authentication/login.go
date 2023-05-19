// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LoginHandlerFunc turns a function with the right signature into a login handler
type LoginHandlerFunc func(LoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LoginHandlerFunc) Handle(params LoginParams) middleware.Responder {
	return fn(params)
}

// LoginHandler interface for that can handle valid login params
type LoginHandler interface {
	Handle(LoginParams) middleware.Responder
}

// NewLogin creates a new http.Handler for the login operation
func NewLogin(ctx *middleware.Context, handler LoginHandler) *Login {
	return &Login{Context: ctx, Handler: handler}
}

/*
	Login swagger:route POST /v1/login authentication login

# Login

Login
*/
type Login struct {
	Context *middleware.Context
	Handler LoginHandler
}

func (o *Login) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewLoginParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// LoginBody login body
//
// swagger:model LoginBody
type LoginBody struct {

	// password
	// Required: true
	Password *string `json:"password"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this login body
func (o *LoginBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	return nil
}

func (o *LoginBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this login body based on context it is used
func (o *LoginBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LoginBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginBody) UnmarshalBinary(b []byte) error {
	var res LoginBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// LoginCreatedBody login created body
//
// swagger:model LoginCreatedBody
type LoginCreatedBody struct {

	// message
	Message string `json:"message,omitempty"`

	// expired at
	ExpiredAt string `json:"expired_at,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *LoginCreatedBody) UnmarshalJSON(raw []byte) error {
	// LoginCreatedBodyAO0
	var dataLoginCreatedBodyAO0 struct {
		Message string `json:"message,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataLoginCreatedBodyAO0); err != nil {
		return err
	}

	o.Message = dataLoginCreatedBodyAO0.Message

	// LoginCreatedBodyAO1
	var dataLoginCreatedBodyAO1 struct {
		ExpiredAt string `json:"expired_at,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataLoginCreatedBodyAO1); err != nil {
		return err
	}

	o.ExpiredAt = dataLoginCreatedBodyAO1.ExpiredAt

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o LoginCreatedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataLoginCreatedBodyAO0 struct {
		Message string `json:"message,omitempty"`
	}

	dataLoginCreatedBodyAO0.Message = o.Message

	jsonDataLoginCreatedBodyAO0, errLoginCreatedBodyAO0 := swag.WriteJSON(dataLoginCreatedBodyAO0)
	if errLoginCreatedBodyAO0 != nil {
		return nil, errLoginCreatedBodyAO0
	}
	_parts = append(_parts, jsonDataLoginCreatedBodyAO0)
	var dataLoginCreatedBodyAO1 struct {
		ExpiredAt string `json:"expired_at,omitempty"`
	}

	dataLoginCreatedBodyAO1.ExpiredAt = o.ExpiredAt

	jsonDataLoginCreatedBodyAO1, errLoginCreatedBodyAO1 := swag.WriteJSON(dataLoginCreatedBodyAO1)
	if errLoginCreatedBodyAO1 != nil {
		return nil, errLoginCreatedBodyAO1
	}
	_parts = append(_parts, jsonDataLoginCreatedBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this login created body
func (o *LoginCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this login created body based on context it is used
func (o *LoginCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LoginCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginCreatedBody) UnmarshalBinary(b []byte) error {
	var res LoginCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// LoginDefaultBody login default body
//
// swagger:model LoginDefaultBody
type LoginDefaultBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	// Example: error
	Message string `json:"message,omitempty"`
}

// Validate validates this login default body
func (o *LoginDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this login default body based on context it is used
func (o *LoginDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LoginDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginDefaultBody) UnmarshalBinary(b []byte) error {
	var res LoginDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}