// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiPhotonicMediaOtsiConnectionEndPointSpec tapi photonic media otsi connection end point spec
// swagger:model tapi.photonic.media.OtsiConnectionEndPointSpec
type TapiPhotonicMediaOtsiConnectionEndPointSpec struct {

	// none
	OtsiTermination *TapiPhotonicMediaOtsiTerminationPac `json:"otsi-termination,omitempty"`
}

// Validate validates this tapi photonic media otsi connection end point spec
func (m *TapiPhotonicMediaOtsiConnectionEndPointSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOtsiTermination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiPhotonicMediaOtsiConnectionEndPointSpec) validateOtsiTermination(formats strfmt.Registry) error {

	if swag.IsZero(m.OtsiTermination) { // not required
		return nil
	}

	if m.OtsiTermination != nil {
		if err := m.OtsiTermination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("otsi-termination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiPhotonicMediaOtsiConnectionEndPointSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPhotonicMediaOtsiConnectionEndPointSpec) UnmarshalBinary(b []byte) error {
	var res TapiPhotonicMediaOtsiConnectionEndPointSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
