// Code generated by go-swagger; DO NOT EDIT.

package d_t_a_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/theovassiliou/doctrans/gen/rest_models"
)

// DTAServerTransformPipeOKCode is the HTTP code returned for type DTAServerTransformPipeOK
const DTAServerTransformPipeOKCode int = 200

/*DTAServerTransformPipeOK A successful response.

swagger:response dTAServerTransformPipeOK
*/
type DTAServerTransformPipeOK struct {

	/*
	  In: Body
	*/
	Payload *rest_models.DtaserviceTransformDocumentResponse `json:"body,omitempty"`
}

// NewDTAServerTransformPipeOK creates DTAServerTransformPipeOK with default headers values
func NewDTAServerTransformPipeOK() *DTAServerTransformPipeOK {

	return &DTAServerTransformPipeOK{}
}

// WithPayload adds the payload to the d t a server transform pipe o k response
func (o *DTAServerTransformPipeOK) WithPayload(payload *rest_models.DtaserviceTransformDocumentResponse) *DTAServerTransformPipeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the d t a server transform pipe o k response
func (o *DTAServerTransformPipeOK) SetPayload(payload *rest_models.DtaserviceTransformDocumentResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DTAServerTransformPipeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DTAServerTransformPipeDefault An unexpected error response

swagger:response dTAServerTransformPipeDefault
*/
type DTAServerTransformPipeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *rest_models.RuntimeError `json:"body,omitempty"`
}

// NewDTAServerTransformPipeDefault creates DTAServerTransformPipeDefault with default headers values
func NewDTAServerTransformPipeDefault(code int) *DTAServerTransformPipeDefault {
	if code <= 0 {
		code = 500
	}

	return &DTAServerTransformPipeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the d t a server transform pipe default response
func (o *DTAServerTransformPipeDefault) WithStatusCode(code int) *DTAServerTransformPipeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the d t a server transform pipe default response
func (o *DTAServerTransformPipeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the d t a server transform pipe default response
func (o *DTAServerTransformPipeDefault) WithPayload(payload *rest_models.RuntimeError) *DTAServerTransformPipeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the d t a server transform pipe default response
func (o *DTAServerTransformPipeDefault) SetPayload(payload *rest_models.RuntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DTAServerTransformPipeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}