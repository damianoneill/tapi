// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthProActiveDualEndedMeasurementJob tapi eth eth pro active dual ended measurement job
// swagger:model tapi.eth.EthProActiveDualEndedMeasurementJob
type TapiEthEthProActiveDualEndedMeasurementJob struct {

	// none
	EthProActiveMeasurementJobControlSink *TapiEthEthProActiveMeasurementJobControlSink `json:"eth-pro-active-measurement-job-control-sink,omitempty"`

	// none
	EthProActiveMeasurementJobControlSource *TapiEthEthProActiveMeasurementJobControlSource `json:"eth-pro-active-measurement-job-control-source,omitempty"`
}

// Validate validates this tapi eth eth pro active dual ended measurement job
func (m *TapiEthEthProActiveDualEndedMeasurementJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthProActiveMeasurementJobControlSink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEthProActiveMeasurementJobControlSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthProActiveDualEndedMeasurementJob) validateEthProActiveMeasurementJobControlSink(formats strfmt.Registry) error {

	if swag.IsZero(m.EthProActiveMeasurementJobControlSink) { // not required
		return nil
	}

	if m.EthProActiveMeasurementJobControlSink != nil {
		if err := m.EthProActiveMeasurementJobControlSink.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-pro-active-measurement-job-control-sink")
			}
			return err
		}
	}

	return nil
}

func (m *TapiEthEthProActiveDualEndedMeasurementJob) validateEthProActiveMeasurementJobControlSource(formats strfmt.Registry) error {

	if swag.IsZero(m.EthProActiveMeasurementJobControlSource) { // not required
		return nil
	}

	if m.EthProActiveMeasurementJobControlSource != nil {
		if err := m.EthProActiveMeasurementJobControlSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-pro-active-measurement-job-control-source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthProActiveDualEndedMeasurementJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthProActiveDualEndedMeasurementJob) UnmarshalBinary(b []byte) error {
	var res TapiEthEthProActiveDualEndedMeasurementJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}