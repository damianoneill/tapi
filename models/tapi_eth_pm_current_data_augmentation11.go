// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthPmCurrentDataAugmentation11 tapi eth pm current data augmentation11
// swagger:model tapi.eth.PmCurrentDataAugmentation11
type TapiEthPmCurrentDataAugmentation11 struct {

	// none
	EthLinkTraceResultData *TapiEthEthLinkTraceResultData `json:"eth-link-trace-result-data,omitempty"`
}

// Validate validates this tapi eth pm current data augmentation11
func (m *TapiEthPmCurrentDataAugmentation11) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthLinkTraceResultData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthPmCurrentDataAugmentation11) validateEthLinkTraceResultData(formats strfmt.Registry) error {

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
func (m *TapiEthPmCurrentDataAugmentation11) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthPmCurrentDataAugmentation11) UnmarshalBinary(b []byte) error {
	var res TapiEthPmCurrentDataAugmentation11
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}