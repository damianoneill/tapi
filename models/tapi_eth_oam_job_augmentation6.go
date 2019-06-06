// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthOamJobAugmentation6 tapi eth oam job augmentation6
// swagger:model tapi.eth.OamJobAugmentation6
type TapiEthOamJobAugmentation6 struct {

	// none
	EthProActive1wayMeasurementJob *TapiEthEthProActive1wayMeasurementJob `json:"eth-pro-active-1way-measurement-job,omitempty"`
}

// Validate validates this tapi eth oam job augmentation6
func (m *TapiEthOamJobAugmentation6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthProActive1wayMeasurementJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthOamJobAugmentation6) validateEthProActive1wayMeasurementJob(formats strfmt.Registry) error {

	if swag.IsZero(m.EthProActive1wayMeasurementJob) { // not required
		return nil
	}

	if m.EthProActive1wayMeasurementJob != nil {
		if err := m.EthProActive1wayMeasurementJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-pro-active-1way-measurement-job")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthOamJobAugmentation6) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthOamJobAugmentation6) UnmarshalBinary(b []byte) error {
	var res TapiEthOamJobAugmentation6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}