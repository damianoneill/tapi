// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DataTapiTopologyGetNodeDetailsPostInput data tapi topology get node details post input
// swagger:model data_tapi-topology_get-node-details-post-input
type DataTapiTopologyGetNodeDetailsPostInput struct {

	// none (leaf)
	TapiTopologyNodeIDOrName string `json:"tapi-topology:node-id-or-name,omitempty"`

	// none (leaf)
	TapiTopologyTopologyIDOrName string `json:"tapi-topology:topology-id-or-name,omitempty"`
}

// Validate validates this data tapi topology get node details post input
func (m *DataTapiTopologyGetNodeDetailsPostInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiTopologyGetNodeDetailsPostInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiTopologyGetNodeDetailsPostInput) UnmarshalBinary(b []byte) error {
	var res DataTapiTopologyGetNodeDetailsPostInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
