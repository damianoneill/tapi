// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiTopologyGettopologydetailsInput tapi topology gettopologydetails input
// swagger:model tapi.topology.gettopologydetails.Input
type TapiTopologyGettopologydetailsInput struct {

	// none
	TopologyIDOrName string `json:"topology-id-or-name,omitempty"`
}

// Validate validates this tapi topology gettopologydetails input
func (m *TapiTopologyGettopologydetailsInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiTopologyGettopologydetailsInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyGettopologydetailsInput) UnmarshalBinary(b []byte) error {
	var res TapiTopologyGettopologydetailsInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
