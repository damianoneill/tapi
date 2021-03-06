// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiEthHistoryDataAugmentation6 tapi eth history data augmentation6
// swagger:model tapi.eth.HistoryDataAugmentation6
type TapiEthHistoryDataAugmentation6 struct {

	// none
	EthOnDemand1LmSourcePerformanceData TapiEthEthOnDemand1LmSourcePerformanceData `json:"eth-on-demand-1-lm-source-performance-data,omitempty"`
}

// Validate validates this tapi eth history data augmentation6
func (m *TapiEthHistoryDataAugmentation6) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthHistoryDataAugmentation6) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthHistoryDataAugmentation6) UnmarshalBinary(b []byte) error {
	var res TapiEthHistoryDataAugmentation6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
