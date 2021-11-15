// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewV2ResetHostValidationParams creates a new V2ResetHostValidationParams object
//
// There are no default values defined in the spec.
func NewV2ResetHostValidationParams() V2ResetHostValidationParams {

	return V2ResetHostValidationParams{}
}

// V2ResetHostValidationParams contains all the bound params for the v2 reset host validation operation
// typically these are obtained from a http.Request
//
// swagger:parameters v2ResetHostValidation
type V2ResetHostValidationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The host that its validation is being reset.
	  Required: true
	  In: path
	*/
	HostID strfmt.UUID
	/*The infra-env of the host that its validation is being reset.
	  Required: true
	  In: path
	*/
	InfraEnvID strfmt.UUID
	/*The id of the validation being reset.
	  Required: true
	  In: path
	*/
	ValidationID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewV2ResetHostValidationParams() beforehand.
func (o *V2ResetHostValidationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rHostID, rhkHostID, _ := route.Params.GetOK("host_id")
	if err := o.bindHostID(rHostID, rhkHostID, route.Formats); err != nil {
		res = append(res, err)
	}

	rInfraEnvID, rhkInfraEnvID, _ := route.Params.GetOK("infra_env_id")
	if err := o.bindInfraEnvID(rInfraEnvID, rhkInfraEnvID, route.Formats); err != nil {
		res = append(res, err)
	}

	rValidationID, rhkValidationID, _ := route.Params.GetOK("validation_id")
	if err := o.bindValidationID(rValidationID, rhkValidationID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindHostID binds and validates parameter HostID from path.
func (o *V2ResetHostValidationParams) bindHostID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("host_id", "path", "strfmt.UUID", raw)
	}
	o.HostID = *(value.(*strfmt.UUID))

	if err := o.validateHostID(formats); err != nil {
		return err
	}

	return nil
}

// validateHostID carries on validations for parameter HostID
func (o *V2ResetHostValidationParams) validateHostID(formats strfmt.Registry) error {

	if err := validate.FormatOf("host_id", "path", "uuid", o.HostID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindInfraEnvID binds and validates parameter InfraEnvID from path.
func (o *V2ResetHostValidationParams) bindInfraEnvID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("infra_env_id", "path", "strfmt.UUID", raw)
	}
	o.InfraEnvID = *(value.(*strfmt.UUID))

	if err := o.validateInfraEnvID(formats); err != nil {
		return err
	}

	return nil
}

// validateInfraEnvID carries on validations for parameter InfraEnvID
func (o *V2ResetHostValidationParams) validateInfraEnvID(formats strfmt.Registry) error {

	if err := validate.FormatOf("infra_env_id", "path", "uuid", o.InfraEnvID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindValidationID binds and validates parameter ValidationID from path.
func (o *V2ResetHostValidationParams) bindValidationID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ValidationID = raw

	return nil
}
