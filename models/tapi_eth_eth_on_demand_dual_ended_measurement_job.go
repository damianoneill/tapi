// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthOnDemandDualEndedMeasurementJob tapi eth eth on demand dual ended measurement job
// swagger:model tapi.eth.EthOnDemandDualEndedMeasurementJob
type TapiEthEthOnDemandDualEndedMeasurementJob struct {

	// none
	EthOnDemandMeasurementJobControlSink *TapiEthEthOnDemandMeasurementJobControlSink `json:"eth-on-demand-measurement-job-control-sink,omitempty"`

	// none
	EthOnDemandMeasurementJobControlSource *TapiEthEthOnDemandMeasurementJobControlSource `json:"eth-on-demand-measurement-job-control-source,omitempty"`
}

// Validate validates this tapi eth eth on demand dual ended measurement job
func (m *TapiEthEthOnDemandDualEndedMeasurementJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthOnDemandMeasurementJobControlSink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEthOnDemandMeasurementJobControlSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthOnDemandDualEndedMeasurementJob) validateEthOnDemandMeasurementJobControlSink(formats strfmt.Registry) error {

	if swag.IsZero(m.EthOnDemandMeasurementJobControlSink) { // not required
		return nil
	}

	if m.EthOnDemandMeasurementJobControlSink != nil {
		if err := m.EthOnDemandMeasurementJobControlSink.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-on-demand-measurement-job-control-sink")
			}
			return err
		}
	}

	return nil
}

func (m *TapiEthEthOnDemandDualEndedMeasurementJob) validateEthOnDemandMeasurementJobControlSource(formats strfmt.Registry) error {

	if swag.IsZero(m.EthOnDemandMeasurementJobControlSource) { // not required
		return nil
	}

	if m.EthOnDemandMeasurementJobControlSource != nil {
		if err := m.EthOnDemandMeasurementJobControlSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-on-demand-measurement-job-control-source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthOnDemandDualEndedMeasurementJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthOnDemandDualEndedMeasurementJob) UnmarshalBinary(b []byte) error {
	var res TapiEthEthOnDemandDualEndedMeasurementJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}