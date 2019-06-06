// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiPhotonicMediaConnectionEndPointAugmentation4 tapi photonic media connection end point augmentation4
// swagger:model tapi.photonic.media.ConnectionEndPointAugmentation4
type TapiPhotonicMediaConnectionEndPointAugmentation4 struct {

	// none
	OtsiAssemblyConnectionEndPointSpec *TapiPhotonicMediaOtsiAssemblyConnectionEndPointSpec `json:"otsi-assembly-connection-end-point-spec,omitempty"`
}

// Validate validates this tapi photonic media connection end point augmentation4
func (m *TapiPhotonicMediaConnectionEndPointAugmentation4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOtsiAssemblyConnectionEndPointSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiPhotonicMediaConnectionEndPointAugmentation4) validateOtsiAssemblyConnectionEndPointSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.OtsiAssemblyConnectionEndPointSpec) { // not required
		return nil
	}

	if m.OtsiAssemblyConnectionEndPointSpec != nil {
		if err := m.OtsiAssemblyConnectionEndPointSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("otsi-assembly-connection-end-point-spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiPhotonicMediaConnectionEndPointAugmentation4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPhotonicMediaConnectionEndPointAugmentation4) UnmarshalBinary(b []byte) error {
	var res TapiPhotonicMediaConnectionEndPointAugmentation4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
