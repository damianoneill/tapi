// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthOamMsgCommonPac tapi eth eth oam msg common pac
// swagger:model tapi.eth.EthOamMsgCommonPac
type TapiEthEthOamMsgCommonPac struct {
	TapiEthEthOamOperationCommonPac

	// G.8052: This parameter provides the length (in number of octet) of the optional Data TLV to be included in the TST frame.
	DataTlvLength int32 `json:"data-tlv-length,omitempty"`

	// G.8052: This parameter provides the eligibility of frames with unicast ETH-TST information to be discarded when congestion conditions are encountered.
	DropEligibility *bool `json:"drop-eligibility,omitempty"`

	// G.8052: This parameter provides the periodicity of the TST OAM messages.
	Period TapiEthOamPeriod `json:"period,omitempty"`

	// G.8052: This parameter provides the test pattern to be used in the optional Data TLV.
	//                     Examples of test patterns include pseudo-random bit sequence (PRBS) 2^31-1 as specified in clause 5.8 of [ITU-T O.150], all '0' pattern, etc.
	//                     The following values of pattern types are defined:
	//                     0: 'Null signal without CRC-32'
	//                     1: 'Null signal with CRC-32'
	//                     2: 'PRBS 2^31-1 without CRC-32'
	//                     3: 'PRBS 2^31-1 with CRC-32'.
	TestPattern int32 `json:"test-pattern,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiEthEthOamMsgCommonPac) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiEthEthOamOperationCommonPac
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiEthEthOamOperationCommonPac = aO0

	// AO1
	var dataAO1 struct {
		DataTlvLength int32 `json:"data-tlv-length,omitempty"`

		DropEligibility *bool `json:"drop-eligibility,omitempty"`

		Period TapiEthOamPeriod `json:"period,omitempty"`

		TestPattern int32 `json:"test-pattern,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.DataTlvLength = dataAO1.DataTlvLength

	m.DropEligibility = dataAO1.DropEligibility

	m.Period = dataAO1.Period

	m.TestPattern = dataAO1.TestPattern

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiEthEthOamMsgCommonPac) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiEthEthOamOperationCommonPac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		DataTlvLength int32 `json:"data-tlv-length,omitempty"`

		DropEligibility *bool `json:"drop-eligibility,omitempty"`

		Period TapiEthOamPeriod `json:"period,omitempty"`

		TestPattern int32 `json:"test-pattern,omitempty"`
	}

	dataAO1.DataTlvLength = m.DataTlvLength

	dataAO1.DropEligibility = m.DropEligibility

	dataAO1.Period = m.Period

	dataAO1.TestPattern = m.TestPattern

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi eth eth oam msg common pac
func (m *TapiEthEthOamMsgCommonPac) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiEthEthOamOperationCommonPac
	if err := m.TapiEthEthOamOperationCommonPac.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthOamMsgCommonPac) validatePeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.Period) { // not required
		return nil
	}

	if err := m.Period.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("period")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthOamMsgCommonPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthOamMsgCommonPac) UnmarshalBinary(b []byte) error {
	var res TapiEthEthOamMsgCommonPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
