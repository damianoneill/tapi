// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamAlarmInfo tapi oam alarm info
// swagger:model tapi.oam.AlarmInfo
type TapiOamAlarmInfo struct {

	// none
	IsTransient *bool `json:"is-transient,omitempty"`

	// none
	PerceivedSeverity TapiOamPerceivedSeverityType `json:"perceived-severity,omitempty"`

	// none
	ProbableCause string `json:"probable-cause,omitempty"`

	// none
	ServiceAffecting TapiOamServiceAffecting `json:"service-affecting,omitempty"`
}

// Validate validates this tapi oam alarm info
func (m *TapiOamAlarmInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePerceivedSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAffecting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamAlarmInfo) validatePerceivedSeverity(formats strfmt.Registry) error {

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

func (m *TapiOamAlarmInfo) validateServiceAffecting(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceAffecting) { // not required
		return nil
	}

	if err := m.ServiceAffecting.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("service-affecting")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamAlarmInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamAlarmInfo) UnmarshalBinary(b []byte) error {
	var res TapiOamAlarmInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
