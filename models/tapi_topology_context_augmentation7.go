// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiTopologyContextAugmentation7 tapi topology context augmentation7
// swagger:model tapi.topology.ContextAugmentation7
type TapiTopologyContextAugmentation7 struct {

	// Augments the base TAPI Context with TopologyService information
	TopologyContext *TapiTopologyTopologyContext `json:"topology-context,omitempty"`
}

// Validate validates this tapi topology context augmentation7
func (m *TapiTopologyContextAugmentation7) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTopologyContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiTopologyContextAugmentation7) validateTopologyContext(formats strfmt.Registry) error {

	if swag.IsZero(m.TopologyContext) { // not required
		return nil
	}

	if m.TopologyContext != nil {
		if err := m.TopologyContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topology-context")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiTopologyContextAugmentation7) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyContextAugmentation7) UnmarshalBinary(b []byte) error {
	var res TapiTopologyContextAugmentation7
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}