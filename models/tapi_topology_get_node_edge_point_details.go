// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiTopologyGetNodeEdgePointDetails tapi topology get node edge point details
// swagger:model tapi.topology.GetNodeEdgePointDetails
type TapiTopologyGetNodeEdgePointDetails struct {

	// output
	Output *TapiTopologyGetnodeedgepointdetailsOutput `json:"output,omitempty"`
}

// Validate validates this tapi topology get node edge point details
func (m *TapiTopologyGetNodeEdgePointDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiTopologyGetNodeEdgePointDetails) validateOutput(formats strfmt.Registry) error {

	if swag.IsZero(m.Output) { // not required
		return nil
	}

	if m.Output != nil {
		if err := m.Output.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiTopologyGetNodeEdgePointDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyGetNodeEdgePointDetails) UnmarshalBinary(b []byte) error {
	var res TapiTopologyGetNodeEdgePointDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
