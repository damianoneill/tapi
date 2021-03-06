// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthInputAugmentation4 tapi eth input augmentation4
// swagger:model tapi.eth.InputAugmentation4
type TapiEthInputAugmentation4 struct {

	// none
	EthOnDemandSingleEndedMeasurementJob *TapiEthEthOnDemandSingleEndedMeasurementJob `json:"eth-on-demand-single-ended-measurement-job,omitempty"`
}

// Validate validates this tapi eth input augmentation4
func (m *TapiEthInputAugmentation4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthOnDemandSingleEndedMeasurementJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthInputAugmentation4) validateEthOnDemandSingleEndedMeasurementJob(formats strfmt.Registry) error {

	if swag.IsZero(m.EthOnDemandSingleEndedMeasurementJob) { // not required
		return nil
	}

	if m.EthOnDemandSingleEndedMeasurementJob != nil {
		if err := m.EthOnDemandSingleEndedMeasurementJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-on-demand-single-ended-measurement-job")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthInputAugmentation4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthInputAugmentation4) UnmarshalBinary(b []byte) error {
	var res TapiEthInputAugmentation4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
