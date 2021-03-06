// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiPhotonicMediaPowerProperties tapi photonic media power properties
// swagger:model tapi.photonic.media.PowerProperties
type TapiPhotonicMediaPowerProperties struct {

	// This describes how power of a signal  is distributed over frequency specified in nW/MHz
	PowerSpectralDensity string `json:"power-spectral-density,omitempty"`

	// The total power at any point in a channel specified in dBm.
	TotalPower string `json:"total-power,omitempty"`
}

// Validate validates this tapi photonic media power properties
func (m *TapiPhotonicMediaPowerProperties) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiPhotonicMediaPowerProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPhotonicMediaPowerProperties) UnmarshalBinary(b []byte) error {
	var res TapiPhotonicMediaPowerProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
