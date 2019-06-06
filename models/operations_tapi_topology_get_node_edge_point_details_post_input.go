// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OperationsTapiTopologyGetNodeEdgePointDetailsPostInput operations tapi topology get node edge point details post input
// swagger:model operations_tapi-topology_get-node-edge-point-details-post-input
type OperationsTapiTopologyGetNodeEdgePointDetailsPostInput struct {

	// none (leaf)
	TapiTopologyEpIDOrName string `json:"tapi-topology:ep-id-or-name,omitempty"`

	// none (leaf)
	TapiTopologyNodeIDOrName string `json:"tapi-topology:node-id-or-name,omitempty"`

	// none (leaf)
	TapiTopologyTopologyIDOrName string `json:"tapi-topology:topology-id-or-name,omitempty"`
}

// Validate validates this operations tapi topology get node edge point details post input
func (m *OperationsTapiTopologyGetNodeEdgePointDetailsPostInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperationsTapiTopologyGetNodeEdgePointDetailsPostInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsTapiTopologyGetNodeEdgePointDetailsPostInput) UnmarshalBinary(b []byte) error {
	var res OperationsTapiTopologyGetNodeEdgePointDetailsPostInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
