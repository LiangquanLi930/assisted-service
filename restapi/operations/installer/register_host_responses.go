// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/filanov/bm-inventory/models"
)

// RegisterHostCreatedCode is the HTTP code returned for type RegisterHostCreated
const RegisterHostCreatedCode int = 201

/*RegisterHostCreated Success.

swagger:response registerHostCreated
*/
type RegisterHostCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Host `json:"body,omitempty"`
}

// NewRegisterHostCreated creates RegisterHostCreated with default headers values
func NewRegisterHostCreated() *RegisterHostCreated {

	return &RegisterHostCreated{}
}

// WithPayload adds the payload to the register host created response
func (o *RegisterHostCreated) WithPayload(payload *models.Host) *RegisterHostCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register host created response
func (o *RegisterHostCreated) SetPayload(payload *models.Host) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterHostCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterHostBadRequestCode is the HTTP code returned for type RegisterHostBadRequest
const RegisterHostBadRequestCode int = 400

/*RegisterHostBadRequest Error.

swagger:response registerHostBadRequest
*/
type RegisterHostBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterHostBadRequest creates RegisterHostBadRequest with default headers values
func NewRegisterHostBadRequest() *RegisterHostBadRequest {

	return &RegisterHostBadRequest{}
}

// WithPayload adds the payload to the register host bad request response
func (o *RegisterHostBadRequest) WithPayload(payload *models.Error) *RegisterHostBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register host bad request response
func (o *RegisterHostBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterHostBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterHostForbiddenCode is the HTTP code returned for type RegisterHostForbidden
const RegisterHostForbiddenCode int = 403

/*RegisterHostForbidden Error.

swagger:response registerHostForbidden
*/
type RegisterHostForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterHostForbidden creates RegisterHostForbidden with default headers values
func NewRegisterHostForbidden() *RegisterHostForbidden {

	return &RegisterHostForbidden{}
}

// WithPayload adds the payload to the register host forbidden response
func (o *RegisterHostForbidden) WithPayload(payload *models.Error) *RegisterHostForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register host forbidden response
func (o *RegisterHostForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterHostForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterHostInternalServerErrorCode is the HTTP code returned for type RegisterHostInternalServerError
const RegisterHostInternalServerErrorCode int = 500

/*RegisterHostInternalServerError Error.

swagger:response registerHostInternalServerError
*/
type RegisterHostInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterHostInternalServerError creates RegisterHostInternalServerError with default headers values
func NewRegisterHostInternalServerError() *RegisterHostInternalServerError {

	return &RegisterHostInternalServerError{}
}

// WithPayload adds the payload to the register host internal server error response
func (o *RegisterHostInternalServerError) WithPayload(payload *models.Error) *RegisterHostInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register host internal server error response
func (o *RegisterHostInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterHostInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
