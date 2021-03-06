// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiPhotonicMediaEndPointAugmentation2 tapi photonic media end point augmentation2
// swagger:model tapi.photonic.media.EndPointAugmentation2
type TapiPhotonicMediaEndPointAugmentation2 struct {

	// none
	OtsiConnectivityServiceEndPointSpec *TapiPhotonicMediaOtsiConnectivityServiceEndPointSpec `json:"otsi-connectivity-service-end-point-spec,omitempty"`
}

// Validate validates this tapi photonic media end point augmentation2
func (m *TapiPhotonicMediaEndPointAugmentation2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOtsiConnectivityServiceEndPointSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiPhotonicMediaEndPointAugmentation2) validateOtsiConnectivityServiceEndPointSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.OtsiConnectivityServiceEndPointSpec) { // not required
		return nil
	}

	if m.OtsiConnectivityServiceEndPointSpec != nil {
		if err := m.OtsiConnectivityServiceEndPointSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("otsi-connectivity-service-end-point-spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiPhotonicMediaEndPointAugmentation2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPhotonicMediaEndPointAugmentation2) UnmarshalBinary(b []byte) error {
	var res TapiPhotonicMediaEndPointAugmentation2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
