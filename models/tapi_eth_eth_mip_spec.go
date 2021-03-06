// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthMipSpec tapi eth eth mip spec
// swagger:model tapi.eth.EthMipSpec
type TapiEthEthMipSpec struct {

	// none
	EthMipCommon *TapiEthEthMipCommon `json:"eth-mip-common,omitempty"`

	// This attribute contains the MAC address of the MIP instance.
	MipMac string `json:"mip-mac,omitempty"`
}

// Validate validates this tapi eth eth mip spec
func (m *TapiEthEthMipSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthMipCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthMipSpec) validateEthMipCommon(formats strfmt.Registry) error {

	if swag.IsZero(m.EthMipCommon) { // not required
		return nil
	}

	if m.EthMipCommon != nil {
		if err := m.EthMipCommon.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-mip-common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthMipSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthMipSpec) UnmarshalBinary(b []byte) error {
	var res TapiEthEthMipSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
