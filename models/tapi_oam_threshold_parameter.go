// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamThresholdParameter tapi oam threshold parameter
// swagger:model tapi.oam.ThresholdParameter
type TapiOamThresholdParameter struct {

	// none
	PmParameterAboveThrs *TapiOamPmParameterValue `json:"pm-parameter-above-thrs,omitempty"`

	// none
	PmParameterBelowThrs *TapiOamPmParameterValue `json:"pm-parameter-below-thrs,omitempty"`

	// none
	PmParameterClearThrs *TapiOamPmParameterValue `json:"pm-parameter-clear-thrs,omitempty"`

	// none
	PmParameterName string `json:"pm-parameter-name,omitempty"`

	// none
	ThresholdLocation TapiOamThresholdCrossingQualifier `json:"threshold-location,omitempty"`
}

// Validate validates this tapi oam threshold parameter
func (m *TapiOamThresholdParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePmParameterAboveThrs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmParameterBelowThrs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmParameterClearThrs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThresholdLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamThresholdParameter) validatePmParameterAboveThrs(formats strfmt.Registry) error {

	if swag.IsZero(m.PmParameterAboveThrs) { // not required
		return nil
	}

	if m.PmParameterAboveThrs != nil {
		if err := m.PmParameterAboveThrs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pm-parameter-above-thrs")
			}
			return err
		}
	}

	return nil
}

func (m *TapiOamThresholdParameter) validatePmParameterBelowThrs(formats strfmt.Registry) error {

	if swag.IsZero(m.PmParameterBelowThrs) { // not required
		return nil
	}

	if m.PmParameterBelowThrs != nil {
		if err := m.PmParameterBelowThrs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pm-parameter-below-thrs")
			}
			return err
		}
	}

	return nil
}

func (m *TapiOamThresholdParameter) validatePmParameterClearThrs(formats strfmt.Registry) error {

	if swag.IsZero(m.PmParameterClearThrs) { // not required
		return nil
	}

	if m.PmParameterClearThrs != nil {
		if err := m.PmParameterClearThrs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pm-parameter-clear-thrs")
			}
			return err
		}
	}

	return nil
}

func (m *TapiOamThresholdParameter) validateThresholdLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.ThresholdLocation) { // not required
		return nil
	}

	if err := m.ThresholdLocation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("threshold-location")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamThresholdParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamThresholdParameter) UnmarshalBinary(b []byte) error {
	var res TapiOamThresholdParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
