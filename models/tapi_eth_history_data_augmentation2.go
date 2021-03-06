// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthHistoryDataAugmentation2 tapi eth history data augmentation2
// swagger:model tapi.eth.HistoryDataAugmentation2
type TapiEthHistoryDataAugmentation2 struct {

	// none
	EthOnDemandLmPerformanceData *TapiEthEthOnDemandLmPerformanceData `json:"eth-on-demand-lm-performance-data,omitempty"`
}

// Validate validates this tapi eth history data augmentation2
func (m *TapiEthHistoryDataAugmentation2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthOnDemandLmPerformanceData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthHistoryDataAugmentation2) validateEthOnDemandLmPerformanceData(formats strfmt.Registry) error {

	if swag.IsZero(m.EthOnDemandLmPerformanceData) { // not required
		return nil
	}

	if m.EthOnDemandLmPerformanceData != nil {
		if err := m.EthOnDemandLmPerformanceData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-on-demand-lm-performance-data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthHistoryDataAugmentation2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthHistoryDataAugmentation2) UnmarshalBinary(b []byte) error {
	var res TapiEthHistoryDataAugmentation2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
