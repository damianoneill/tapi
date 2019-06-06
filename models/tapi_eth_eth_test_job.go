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

	// G.8052: This parameter provides the destination address, i.e., the MAC Address of the target MEP or MIP.
	DestinationAddress string `json:"destination-address,omitempty"`

	// none
	EthOamTestLoopbackCommonPac *TapiEthEthOamTestLoopbackCommonPac `json:"eth-oam-test-loopback-common-pac,omitempty"`

	// none
	EthTestJobSinkPoint *TapiEthEthTestJobSinkPoint `json:"eth-test-job-sink-point,omitempty"`

	// This parameter specifies how many TST messages to be sent.
	Number int32 `json:"number,omitempty"`

	// G.8052: This parameter provides the test pattern to be used in the optional Data TLV.
	//                 Examples of test patterns include pseudo-random bit sequence (PRBS) 2^31-1 as specified in clause 5.8 of [ITU-T O.150], all '0' pattern, etc.
	TestPattern string `json:"test-pattern,omitempty"`
}

// Validate validates this tapi eth eth test job
func (m *TapiEthEthTestJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthOamTestLoopbackCommonPac(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEthTestJobSinkPoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthTestJob) validateEthOamTestLoopbackCommonPac(formats strfmt.Registry) error {

	if swag.IsZero(m.EthOamTestLoopbackCommonPac) { // not required
		return nil
	}

	if m.EthOamTestLoopbackCommonPac != nil {
		if err := m.EthOamTestLoopbackCommonPac.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-oam-test-loopback-common-pac")
			}
			return err
		}
	}

	return nil
}

func (m *TapiEthEthTestJob) validateEthTestJobSinkPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.EthTestJobSinkPoint) { // not required
		return nil
	}

	if m.EthTestJobSinkPoint != nil {
		if err := m.EthTestJobSinkPoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-test-job-sink-point")
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