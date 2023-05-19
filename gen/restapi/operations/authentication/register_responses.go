// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// RegisterCreatedCode is the HTTP code returned for type RegisterCreated
const RegisterCreatedCode int = 201

/*
RegisterCreated Success register

swagger:response registerCreated
*/
type RegisterCreated struct {

	/*
	  In: Body
	*/
	Payload *RegisterCreatedBody `json:"body,omitempty"`
}

// NewRegisterCreated creates RegisterCreated with default headers values
func NewRegisterCreated() *RegisterCreated {

	return &RegisterCreated{}
}

// WithPayload adds the payload to the register created response
func (o *RegisterCreated) WithPayload(payload *RegisterCreatedBody) *RegisterCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register created response
func (o *RegisterCreated) SetPayload(payload *RegisterCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
RegisterDefault Server Error

swagger:response registerDefault
*/
type RegisterDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *RegisterDefaultBody `json:"body,omitempty"`
}

// NewRegisterDefault creates RegisterDefault with default headers values
func NewRegisterDefault(code int) *RegisterDefault {
	if code <= 0 {
		code = 500
	}

	return &RegisterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the register default response
func (o *RegisterDefault) WithStatusCode(code int) *RegisterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the register default response
func (o *RegisterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the register default response
func (o *RegisterDefault) WithPayload(payload *RegisterDefaultBody) *RegisterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register default response
func (o *RegisterDefault) SetPayload(payload *RegisterDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}