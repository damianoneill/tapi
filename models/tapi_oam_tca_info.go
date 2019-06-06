// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamTcaInfo tapi oam tca info
// swagger:model tapi.oam.TcaInfo
type TapiOamTcaInfo struct {

	// none
	IsTransient *bool `json:"is-transient,omitempty"`

	// none
	MeasurementInterval string `json:"measurement-interval,omitempty"`

	// MEF 35.1
	//                 Identification of the PM Session for which the TCA Function was configured.
	OamJob *TapiOamOamJobRef `json:"oam-job,omitempty"`

	// none
	PerceivedSeverity TapiOamPerceivedTcaSeverity `json:"perceived-severity,omitempty"`

	// none
	SuspectIntervalFlag *bool `json:"suspect-interval-flag,omitempty"`

	// none
	ThresholdParameter *TapiOamThresholdParameter `json:"threshold-parameter,omitempty"`
}

// Validate validates this tapi oam tca info
func (m *TapiOamTcaInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOamJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerceivedSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThresholdParameter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamTcaInfo) validateOamJob(formats strfmt.Registry) error {

	if swag.IsZero(m.OamJob) { // not required
		return nil
	}

	if m.OamJob != nil {
		if err := m.OamJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oam-job")
			}
			return err
		}
	}

	return nil
}

func (m *TapiOamTcaInfo) validatePerceivedSeverity(formats strfmt.Registry) error {

	if swag.IsZero(m.PerceivedSeverity) { // not required
		return nil
	}

	if err := m.PerceivedSeverity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("perceived-severity")
		}
		return err
	}

	return nil
}

func (m *TapiOamTcaInfo) validateThresholdParameter(formats strfmt.Registry) error {

	if swag.IsZero(m.ThresholdParameter) { // not required
		return nil
	}

	if m.ThresholdParameter != nil {
		if err := m.ThresholdParameter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("threshold-parameter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamTcaInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamTcaInfo) UnmarshalBinary(b []byte) error {
	var res TapiOamTcaInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}