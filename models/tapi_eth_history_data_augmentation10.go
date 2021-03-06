// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiEthHistoryDataAugmentation10 tapi eth history data augmentation10
// swagger:model tapi.eth.HistoryDataAugmentation10
type TapiEthHistoryDataAugmentation10 struct {

	// none
	EthOnDemand1DmSourcePerformanceData TapiEthEthOnDemand1DmSourcePerformanceData `json:"eth-on-demand-1-dm-source-performance-data,omitempty"`
}

// Validate validates this tapi eth history data augmentation10
func (m *TapiEthHistoryDataAugmentation10) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthHistoryDataAugmentation10) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthHistoryDataAugmentation10) UnmarshalBinary(b []byte) error {
	var res TapiEthHistoryDataAugmentation10
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
