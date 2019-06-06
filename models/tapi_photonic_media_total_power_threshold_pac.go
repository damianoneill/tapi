// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiPhotonicMediaTotalPowerThresholdPac tapi photonic media total power threshold pac
// swagger:model tapi.photonic.media.TotalPowerThresholdPac
type TapiPhotonicMediaTotalPowerThresholdPac struct {

	// Can read the value of the default  threshold that was set
	TotalPowerLowerWarnThresholdDefault string `json:"total-power-lower-warn-threshold-default,omitempty"`

	// Can  read the value of the upper threshold that was set
	TotalPowerLowerWarnThresholdMax string `json:"total-power-lower-warn-threshold-max,omitempty"`

	// Can read the value of the lower threshold that was set
	TotalPowerLowerWarnThresholdMin string `json:"total-power-lower-warn-threshold-min,omitempty"`

	// Can read the value of the default  threshold that was set
	TotalPowerUpperWarnThresholdDefault string `json:"total-power-upper-warn-threshold-default,omitempty"`

	// Can  read the value of the upper threshold that was set
	TotalPowerUpperWarnThresholdMax string `json:"total-power-upper-warn-threshold-max,omitempty"`

	// Can read the value of the lower threshold that was set
	TotalPowerUpperWarnThresholdMin string `json:"total-power-upper-warn-threshold-min,omitempty"`
}

// Validate validates this tapi photonic media total power threshold pac
func (m *TapiPhotonicMediaTotalPowerThresholdPac) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiPhotonicMediaTotalPowerThresholdPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPhotonicMediaTotalPowerThresholdPac) UnmarshalBinary(b []byte) error {
	var res TapiPhotonicMediaTotalPowerThresholdPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
