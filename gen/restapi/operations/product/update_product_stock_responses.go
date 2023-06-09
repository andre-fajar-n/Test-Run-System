// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateProductStockCreatedCode is the HTTP code returned for type UpdateProductStockCreated
const UpdateProductStockCreatedCode int = 201

/*
UpdateProductStockCreated Success update

swagger:response updateProductStockCreated
*/
type UpdateProductStockCreated struct {

	/*
	  In: Body
	*/
	Payload *UpdateProductStockCreatedBody `json:"body,omitempty"`
}

// NewUpdateProductStockCreated creates UpdateProductStockCreated with default headers values
func NewUpdateProductStockCreated() *UpdateProductStockCreated {

	return &UpdateProductStockCreated{}
}

// WithPayload adds the payload to the update product stock created response
func (o *UpdateProductStockCreated) WithPayload(payload *UpdateProductStockCreatedBody) *UpdateProductStockCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update product stock created response
func (o *UpdateProductStockCreated) SetPayload(payload *UpdateProductStockCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProductStockCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
UpdateProductStockDefault Server Error

swagger:response updateProductStockDefault
*/
type UpdateProductStockDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *UpdateProductStockDefaultBody `json:"body,omitempty"`
}

// NewUpdateProductStockDefault creates UpdateProductStockDefault with default headers values
func NewUpdateProductStockDefault(code int) *UpdateProductStockDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateProductStockDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update product stock default response
func (o *UpdateProductStockDefault) WithStatusCode(code int) *UpdateProductStockDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update product stock default response
func (o *UpdateProductStockDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update product stock default response
func (o *UpdateProductStockDefault) WithPayload(payload *UpdateProductStockDefaultBody) *UpdateProductStockDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update product stock default response
func (o *UpdateProductStockDefault) SetPayload(payload *UpdateProductStockDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProductStockDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
