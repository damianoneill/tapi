// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthHistoryDataAugmentation4 tapi eth history data augmentation4
// swagger:model tapi.eth.HistoryDataAugmentation4
type TapiEthHistoryDataAugmentation4 struct {

	// none
	EthProActiveDmPerformanceData *TapiEthEthProActiveDmPerformanceData `json:"eth-pro-active-dm-performance-data,omitempty"`
}

// Validate validates this tapi eth history data augmentation4
func (m *TapiEthHistoryDataAugmentation4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthProActiveDmPerformanceData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthHistoryDataAugmentation4) validateEthProActiveDmPerformanceData(formats strfmt.Registry) error {

	if swag.IsZero(m.EthProActiveDmPerformanceData) { // not required
		return nil
	}

	if m.EthProActiveDmPerformanceData != nil {
		if err := m.EthProActiveDmPerformanceData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-pro-active-dm-performance-data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthHistoryDataAugmentation4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthHistoryDataAugmentation4) UnmarshalBinary(b []byte) error {
	var res TapiEthHistoryDataAugmentation4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
