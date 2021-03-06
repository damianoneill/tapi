// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiEthCurrentDataAugmentation11 tapi eth current data augmentation11
// swagger:model tapi.eth.CurrentDataAugmentation11
type TapiEthCurrentDataAugmentation11 struct {

	// none
	EthProActive1DmSourcePerformanceData TapiEthEthProActive1DmSourcePerformanceData `json:"eth-pro-active-1-dm-source-performance-data,omitempty"`
}

// Validate validates this tapi eth current data augmentation11
func (m *TapiEthCurrentDataAugmentation11) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthCurrentDataAugmentation11) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthCurrentDataAugmentation11) UnmarshalBinary(b []byte) error {
	var res TapiEthCurrentDataAugmentation11
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
