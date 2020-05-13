// Code generated by go-swagger; DO NOT EDIT.

package d_t_a_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/theovassiliou/doctrans/gen/rest_models"
)

// DTAServerTransformDocumentOKCode is the HTTP code returned for type DTAServerTransformDocumentOK
const DTAServerTransformDocumentOKCode int = 200

/*DTAServerTransformDocumentOK A successful response.

swagger:response dTAServerTransformDocumentOK
*/
type DTAServerTransformDocumentOK struct {

	/*
	  In: Body
	*/
	Payload *rest_models.DtaserviceTransformDocumentResponse `json:"body,omitempty"`
}

// NewDTAServerTransformDocumentOK creates DTAServerTransformDocumentOK with default headers values
func NewDTAServerTransformDocumentOK() *DTAServerTransformDocumentOK {

	return &DTAServerTransformDocumentOK{}
}

// WithPayload adds the payload to the d t a server transform document o k response
func (o *DTAServerTransformDocumentOK) WithPayload(payload *rest_models.DtaserviceTransformDocumentResponse) *DTAServerTransformDocumentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the d t a server transform document o k response
func (o *DTAServerTransformDocumentOK) SetPayload(payload *rest_models.DtaserviceTransformDocumentResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DTAServerTransformDocumentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DTAServerTransformDocumentDefault An unexpected error response

swagger:response dTAServerTransformDocumentDefault
*/
type DTAServerTransformDocumentDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *rest_models.RuntimeError `json:"body,omitempty"`
}

// NewDTAServerTransformDocumentDefault creates DTAServerTransformDocumentDefault with default headers values
func NewDTAServerTransformDocumentDefault(code int) *DTAServerTransformDocumentDefault {
	if code <= 0 {
		code = 500
	}

	return &DTAServerTransformDocumentDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the d t a server transform document default response
func (o *DTAServerTransformDocumentDefault) WithStatusCode(code int) *DTAServerTransformDocumentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the d t a server transform document default response
func (o *DTAServerTransformDocumentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the d t a server transform document default response
func (o *DTAServerTransformDocumentDefault) WithPayload(payload *rest_models.RuntimeError) *DTAServerTransformDocumentDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the d t a server transform document default response
func (o *DTAServerTransformDocumentDefault) SetPayload(payload *rest_models.RuntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DTAServerTransformDocumentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
