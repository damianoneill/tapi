// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiEthEthProActiveMeasurementJobControlSink tapi eth eth pro active measurement job control sink
// swagger:model tapi.eth.EthProActiveMeasurementJobControlSink
type TapiEthEthProActiveMeasurementJobControlSink struct {

	// This attribute identifies the state of the measurement job. If set to TRUE, the MEP performs proactive Performance Measurement.
	IsEnabled *bool `json:"is-enabled,omitempty"`

	// none
	ResponderMepID string `json:"responder-mep-id,omitempty"`

	// This attribute contains the MAC address of the peer MEP. See G.8013 for details.
	SourceAddress string `json:"source-address,omitempty"`

	// This attribute is used to distinguish each measurement session if multiple measurement sessions are simultaneously activated towards a peer MEP including concurrent on-demand and proactive tests. It must be unique at least within the context of any measurement type for the MEG and initiating MEP.
	//                     range of type : 0..(2^32) - 1
	TestIdentifier int32 `json:"test-identifier,omitempty"`
}

// Validate validates this tapi eth eth pro active measurement job control sink
func (m *TapiEthEthProActiveMeasurementJobControlSink) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthProActiveMeasurementJobControlSink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthProActiveMeasurementJobControlSink) UnmarshalBinary(b []byte) error {
	var res TapiEthEthProActiveMeasurementJobControlSink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
