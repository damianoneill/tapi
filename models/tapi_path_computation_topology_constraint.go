// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiPathComputationTopologyConstraint tapi path computation topology constraint
// swagger:model tapi.path.computation.TopologyConstraint
type TapiPathComputationTopologyConstraint struct {
	TapiCommonLocalClass

	// none
	AvoidTopology string `json:"avoid-topology,omitempty"`

	// Zero and positive values: zero means 'strongly required to be included', +1 means 'less strongly required to be included', etc.
	//                 For example the work/intended route will be calculated considering the topologies which weights are lowest (but not negative).
	//                 Negative values: -1 means 'strongly required to be excluded', -2 means 'less strongly required to be excluded', etc.
	ConstraintWeight int32 `json:"constraint-weight,omitempty"`

	// none
	ExcludeLink string `json:"exclude-link,omitempty"`

	// none
	ExcludeNode string `json:"exclude-node,omitempty"`

	// none
	ExcludeNodeEdgePoint string `json:"exclude-node-edge-point,omitempty"`

	// none
	ExcludePath string `json:"exclude-path,omitempty"`

	// This is a loose constraint - that is it is unordered and could be a partial list
	IncludeLink string `json:"include-link,omitempty"`

	// This is a loose constraint - that is it is unordered and could be a partial list
	IncludeNode string `json:"include-node,omitempty"`

	// none
	IncludeNodeEdgePoint string `json:"include-node-edge-point,omitempty"`

	// none
	IncludePath string `json:"include-path,omitempty"`

	// none
	IncludeTopology string `json:"include-topology,omitempty"`

	// soft constraint requested by client to indicate the layer(s) of transport connection that it prefers to carry the service. This could be same as the service layer or one of the supported server layers
	PreferredTransportLayer TapiCommonLayerProtocolName `json:"preferred-transport-layer,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiPathComputationTopologyConstraint) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonLocalClass
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonLocalClass = aO0

	// AO1
	var dataAO1 struct {
		AvoidTopology string `json:"avoid-topology,omitempty"`

		ConstraintWeight int32 `json:"constraint-weight,omitempty"`

		ExcludeLink string `json:"exclude-link,omitempty"`

		ExcludeNode string `json:"exclude-node,omitempty"`

		ExcludeNodeEdgePoint string `json:"exclude-node-edge-point,omitempty"`

		ExcludePath string `json:"exclude-path,omitempty"`

		IncludeLink string `json:"include-link,omitempty"`

		IncludeNode string `json:"include-node,omitempty"`

		IncludeNodeEdgePoint string `json:"include-node-edge-point,omitempty"`

		IncludePath string `json:"include-path,omitempty"`

		IncludeTopology string `json:"include-topology,omitempty"`

		PreferredTransportLayer TapiCommonLayerProtocolName `json:"preferred-transport-layer,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AvoidTopology = dataAO1.AvoidTopology

	m.ConstraintWeight = dataAO1.ConstraintWeight

	m.ExcludeLink = dataAO1.ExcludeLink

	m.ExcludeNode = dataAO1.ExcludeNode

	m.ExcludeNodeEdgePoint = dataAO1.ExcludeNodeEdgePoint

	m.ExcludePath = dataAO1.ExcludePath

	m.IncludeLink = dataAO1.IncludeLink

	m.IncludeNode = dataAO1.IncludeNode

	m.IncludeNodeEdgePoint = dataAO1.IncludeNodeEdgePoint

	m.IncludePath = dataAO1.IncludePath

	m.IncludeTopology = dataAO1.IncludeTopology

	m.PreferredTransportLayer = dataAO1.PreferredTransportLayer

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiPathComputationTopologyConstraint) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiCommonLocalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AvoidTopology string `json:"avoid-topology,omitempty"`

		ConstraintWeight int32 `json:"constraint-weight,omitempty"`

		ExcludeLink string `json:"exclude-link,omitempty"`

		ExcludeNode string `json:"exclude-node,omitempty"`

		ExcludeNodeEdgePoint string `json:"exclude-node-edge-point,omitempty"`

		ExcludePath string `json:"exclude-path,omitempty"`

		IncludeLink string `json:"include-link,omitempty"`

		IncludeNode string `json:"include-node,omitempty"`

		IncludeNodeEdgePoint string `json:"include-node-edge-point,omitempty"`

		IncludePath string `json:"include-path,omitempty"`

		IncludeTopology string `json:"include-topology,omitempty"`

		PreferredTransportLayer TapiCommonLayerProtocolName `json:"preferred-transport-layer,omitempty"`
	}

	dataAO1.AvoidTopology = m.AvoidTopology

	dataAO1.ConstraintWeight = m.ConstraintWeight

	dataAO1.ExcludeLink = m.ExcludeLink

	dataAO1.ExcludeNode = m.ExcludeNode

	dataAO1.ExcludeNodeEdgePoint = m.ExcludeNodeEdgePoint

	dataAO1.ExcludePath = m.ExcludePath

	dataAO1.IncludeLink = m.IncludeLink

	dataAO1.IncludeNode = m.IncludeNode

	dataAO1.IncludeNodeEdgePoint = m.IncludeNodeEdgePoint

	dataAO1.IncludePath = m.IncludePath

	dataAO1.IncludeTopology = m.IncludeTopology

	dataAO1.PreferredTransportLayer = m.PreferredTransportLayer

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi path computation topology constraint
func (m *TapiPathComputationTopologyConstraint) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonLocalClass
	if err := m.TapiCommonLocalClass.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferredTransportLayer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiPathComputationTopologyConstraint) validatePreferredTransportLayer(formats strfmt.Registry) error {

	if swag.IsZero(m.PreferredTransportLayer) { // not required
		return nil
	}

	if err := m.PreferredTransportLayer.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("preferred-transport-layer")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiPathComputationTopologyConstraint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPathComputationTopologyConstraint) UnmarshalBinary(b []byte) error {
	var res TapiPathComputationTopologyConstraint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
