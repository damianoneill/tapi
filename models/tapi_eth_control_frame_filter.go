// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiEthControlFrameFilter tapi eth control frame filter
// swagger:model tapi.eth.ControlFrameFilter
type TapiEthControlFrameFilter struct {

	// Reserved for future standardization.
	C200000A *bool `json:"c-2-00-00-0-a,omitempty"`

	// Reserved for future standardization.
	C200000B *bool `json:"c-2-00-00-0-b,omitempty"`

	// Reserved for future standardization.
	C200000C *bool `json:"c-2-00-00-0-c,omitempty"`

	// This attribute identifies the Provider Bridge MVRP address.
	C200000D *bool `json:"c-2-00-00-0-d,omitempty"`

	// This attribute identifies the Individual LAN Scope group address, Nearest Bridge group address (LLDP protocol).
	C200000E *bool `json:"c-2-00-00-0-e,omitempty"`

	// Reserved for future standardization.
	C200000F *bool `json:"c-2-00-00-0-f,omitempty"`

	// This attribute identifies the STP/RSTP/MSTP protocol address.
	C2000000 *bool `json:"c-2-00-00-00,omitempty"`

	// This attribute identifies the IEEE MAC-specific Control Protocols group address (PAUSE protocol).
	C2000001 *bool `json:"c-2-00-00-01,omitempty"`

	// This attribute identifies the IEEE 802.3 Slow_Protocols_Multicast address (LACP/LAMP or Link OAM protocols).
	C2000002 *bool `json:"c-2-00-00-02,omitempty"`

	// This attribute identifies the Nearest non-TPMR Bridge group address (Port Authentication protocol).
	C2000003 *bool `json:"c-2-00-00-03,omitempty"`

	// This attribute identifies the IEEE MAC-specific Control Protocols group address.
	C2000004 *bool `json:"c-2-00-00-04,omitempty"`

	// Reserved for future standardization.
	C2000005 *bool `json:"c-2-00-00-05,omitempty"`

	// Reserved for future standardization.
	C2000006 *bool `json:"c-2-00-00-06,omitempty"`

	// This attribute identifies the Metro Ethernet Forum E-LMI protocol group address.
	C2000007 *bool `json:"c-2-00-00-07,omitempty"`

	// This attribute identifies the Provider Bridge Group address.
	C2000008 *bool `json:"c-2-00-00-08,omitempty"`

	// Reserved for future standardization.
	C2000009 *bool `json:"c-2-00-00-09,omitempty"`

	// This attribute identifies the 'All LANs Bridge Management Group Address'.
	C2000010 *bool `json:"c-2-00-00-10,omitempty"`

	// Reserved for future standardization.
	C200002A *bool `json:"c-2-00-00-2-a,omitempty"`

	// Reserved for future standardization.
	C200002B *bool `json:"c-2-00-00-2-b,omitempty"`

	// Reserved for future standardization.
	C200002C *bool `json:"c-2-00-00-2-c,omitempty"`

	// Reserved for future standardization.
	C200002D *bool `json:"c-2-00-00-2-d,omitempty"`

	// Reserved for future standardization.
	C200002E *bool `json:"c-2-00-00-2-e,omitempty"`

	// Reserved for future standardization.
	C200002F *bool `json:"c-2-00-00-2-f,omitempty"`

	// This attribute identifies the Customer and Provider Bridge MMRP address.
	C2000020 *bool `json:"c-2-00-00-20,omitempty"`

	// This attribute identifies the Customer Bridge MVRP address.
	C2000021 *bool `json:"c-2-00-00-21,omitempty"`

	// Reserved for future standardization.
	C2000022 *bool `json:"c-2-00-00-22,omitempty"`

	// Reserved for future standardization.
	C2000023 *bool `json:"c-2-00-00-23,omitempty"`

	// Reserved for future standardization.
	C2000024 *bool `json:"c-2-00-00-24,omitempty"`

	// Reserved for future standardization.
	C2000025 *bool `json:"c-2-00-00-25,omitempty"`

	// Reserved for future standardization.
	C2000026 *bool `json:"c-2-00-00-26,omitempty"`

	// Reserved for future standardization.
	C2000027 *bool `json:"c-2-00-00-27,omitempty"`

	// Reserved for future standardization.
	C2000028 *bool `json:"c-2-00-00-28,omitempty"`

	// Reserved for future standardization.
	C2000029 *bool `json:"c-2-00-00-29,omitempty"`
}

// Validate validates this tapi eth control frame filter
func (m *TapiEthControlFrameFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthControlFrameFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthControlFrameFilter) UnmarshalBinary(b []byte) error {
	var res TapiEthControlFrameFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
