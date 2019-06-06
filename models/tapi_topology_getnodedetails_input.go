// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiTopologyGetnodedetailsInput tapi topology getnodedetails input
// swagger:model tapi.topology.getnodedetails.Input
type TapiTopologyGetnodedetailsInput struct {

	// none
	NodeIDOrName string `json:"node-id-or-name,omitempty"`

	// none
	TopologyIDOrName string `json:"topology-id-or-name,omitempty"`
}

// Validate validates this tapi topology getnodedetails input
func (m *TapiTopologyGetnodedetailsInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiTopologyGetnodedetailsInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyGetnodedetailsInput) UnmarshalBinary(b []byte) error {
	var res TapiTopologyGetnodedetailsInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}