// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/checkr/flagr/swagger_gen/models"
)

// GetFlagByNameOKCode is the HTTP code returned for type GetFlagByNameOK
const GetFlagByNameOKCode int = 200

/*GetFlagByNameOK returns the flag

swagger:response getFlagByNameOK
*/
type GetFlagByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.Flag `json:"body,omitempty"`
}

// NewGetFlagByNameOK creates GetFlagByNameOK with default headers values
func NewGetFlagByNameOK() *GetFlagByNameOK {

	return &GetFlagByNameOK{}
}

// WithPayload adds the payload to the get flag by name o k response
func (o *GetFlagByNameOK) WithPayload(payload *models.Flag) *GetFlagByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get flag by name o k response
func (o *GetFlagByNameOK) SetPayload(payload *models.Flag) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFlagByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetFlagByNameDefault generic error response

swagger:response getFlagByNameDefault
*/
type GetFlagByNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFlagByNameDefault creates GetFlagByNameDefault with default headers values
func NewGetFlagByNameDefault(code int) *GetFlagByNameDefault {
	if code <= 0 {
		code = 500
	}

	return &GetFlagByNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get flag by name default response
func (o *GetFlagByNameDefault) WithStatusCode(code int) *GetFlagByNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get flag by name default response
func (o *GetFlagByNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get flag by name default response
func (o *GetFlagByNameDefault) WithPayload(payload *models.Error) *GetFlagByNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get flag by name default response
func (o *GetFlagByNameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFlagByNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
