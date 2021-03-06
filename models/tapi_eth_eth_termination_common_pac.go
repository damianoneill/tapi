// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthTerminationCommonPac tapi eth eth termination common pac
// swagger:model tapi.eth.EthTerminationCommonPac
type TapiEthEthTerminationCommonPac struct {

	// This attribute models the ETHx/ETH-m _A_Sk_MI_Etype information defined in G.8021.
	EtherType TapiEthVlanType `json:"ether-type,omitempty"`

	// This attribute models the ETHx/ETH-m_A_Sk_MI_Filter_Config information defined in G.8021.
	//                 It indicates the configured filter action for each of the 33 group MAC addresses for control frames.
	//                 The 33 MAC addresses are:
	//                 01-80-C2-00-00-10,
	//                 01-80-C2-00-00-00 to 01-80-C2-00-00-0F, and
	//                 01-80-C2-00-00-20 to 01-80-C2-00-00-2F.
	//                 The filter action is Pass or Block.
	//                 If the destination address of the incoming ETH_CI_D matches one of the above addresses, the filter process shall perform the corresponding configured filter action.
	//                 If none of the above addresses match, the ETH_CI_D is passed.
	//
	FilterConfig1 []string `json:"filter-config-1"`

	// This attribute models the ETHx/ETH-m_A_Sk_MI_Frametype_Config information defined in G.8021.
	FrametypeConfig TapiEthFrameType `json:"frametype-config,omitempty"`

	// This attribute models the ETHx/ETH-m _A_Sk_MI_PVID information defined in G.8021.
	PortVid *string `json:"port-vid,omitempty"`

	// This attribute models the ETHx/ETH-m _A_Sk_MI_PCP_Config information defined in G.8021.
	PriorityCodePointConfig TapiEthPcpCoding `json:"priority-code-point-config,omitempty"`

	// This attribute models the ETHx/ETH-m _A_Sk_MI_P_Regenerate information defined in G.8021.
	PriorityRegenerate *TapiEthPriorityMapping `json:"priority-regenerate,omitempty"`
}

// Validate validates this tapi eth eth termination common pac
func (m *TapiEthEthTerminationCommonPac) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEtherType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrametypeConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriorityCodePointConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriorityRegenerate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthTerminationCommonPac) validateEtherType(formats strfmt.Registry) error {

	if swag.IsZero(m.EtherType) { // not required
		return nil
	}

	if err := m.EtherType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ether-type")
		}
		return err
	}

	return nil
}

func (m *TapiEthEthTerminationCommonPac) validateFrametypeConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.FrametypeConfig) { // not required
		return nil
	}

	if err := m.FrametypeConfig.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frametype-config")
		}
		return err
	}

	return nil
}

func (m *TapiEthEthTerminationCommonPac) validatePriorityCodePointConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.PriorityCodePointConfig) { // not required
		return nil
	}

	if err := m.PriorityCodePointConfig.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("priority-code-point-config")
		}
		return err
	}

	return nil
}

func (m *TapiEthEthTerminationCommonPac) validatePriorityRegenerate(formats strfmt.Registry) error {

	if swag.IsZero(m.PriorityRegenerate) { // not required
		return nil
	}

	if m.PriorityRegenerate != nil {
		if err := m.PriorityRegenerate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priority-regenerate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthTerminationCommonPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthTerminationCommonPac) UnmarshalBinary(b []byte) error {
	var res TapiEthEthTerminationCommonPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
