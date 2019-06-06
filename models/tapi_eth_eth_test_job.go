// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthTestJob tapi eth eth test job
// swagger:model tapi.eth.EthTestJob
type TapiEthEthTestJob struct {

	// none
	EthTestMsg *TapiEthEthOamMsgCommonPac `json:"eth-test-msg,omitempty"`
}

// Validate validates this tapi eth eth test job
func (m *TapiEthEthTestJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthTestMsg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthTestJob) validateEthTestMsg(formats strfmt.Registry) error {

	if swag.IsZero(m.EthTestMsg) { // not required
		return nil
	}

	if m.EthTestMsg != nil {
		if err := m.EthTestMsg.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-test-msg")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthTestJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthTestJob) UnmarshalBinary(b []byte) error {
	var res TapiEthEthTestJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}