// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthCurrentDataAugmentation15 tapi eth current data augmentation15
// swagger:model tapi.eth.CurrentDataAugmentation15
type TapiEthCurrentDataAugmentation15 struct {

	// none
	EthLinkTraceResultData *TapiEthEthLinkTraceResultData `json:"eth-link-trace-result-data,omitempty"`
}

// Validate validates this tapi eth current data augmentation15
func (m *TapiEthCurrentDataAugmentation15) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthLinkTraceResultData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthCurrentDataAugmentation15) validateEthLinkTraceResultData(formats strfmt.Registry) error {

	if swag.IsZero(m.EthLinkTraceResultData) { // not required
		return nil
	}

	if m.EthLinkTraceResultData != nil {
		if err := m.EthLinkTraceResultData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-link-trace-result-data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthCurrentDataAugmentation15) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthCurrentDataAugmentation15) UnmarshalBinary(b []byte) error {
	var res TapiEthCurrentDataAugmentation15
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
