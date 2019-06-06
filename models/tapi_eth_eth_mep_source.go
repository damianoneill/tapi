// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthMepSource tapi eth eth mep source
// swagger:model tapi.eth.EthMepSource
type TapiEthEthMepSource struct {

	// This attribute specifies the priority of the APS messages.
	//                     See section 8.1.5    APS insert process in G.8021.
	ApsPriority *int32 `json:"aps-priority,omitempty"`

	// This attribute models the combination of all CSF related MI signals (MI_CSF_Enable, MI_CSFrdifdi_Enable, MI_CSFdci_Enable) as defined in G.8021.
	CsfConfig TapiEthCsfConfig `json:"csf-config,omitempty"`

	// This attribute models the MI_CSF_Period signal defined in G.8021 and configured as specified in G8051. It is the period at which the CSF messages should be sent.
	//                     range of type : 1s, 1min
	CsfPeriod TapiEthOamPeriod `json:"csf-period,omitempty"`

	// This attribute models the MI_CSF_Pri signal defined in G.8021 and configured as specified in G8051. It is the priority at which the CSF messages should be sent
	CsfPriority *int32 `json:"csf-priority,omitempty"`
}

// Validate validates this tapi eth eth mep source
func (m *TapiEthEthMepSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCsfConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsfPeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthMepSource) validateCsfConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.CsfConfig) { // not required
		return nil
	}

	if err := m.CsfConfig.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("csf-config")
		}
		return err
	}

	return nil
}

func (m *TapiEthEthMepSource) validateCsfPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.CsfPeriod) { // not required
		return nil
	}

	if err := m.CsfPeriod.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("csf-period")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthMepSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthMepSource) UnmarshalBinary(b []byte) error {
	var res TapiEthEthMepSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
