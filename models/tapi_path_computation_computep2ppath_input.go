// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiPathComputationComputep2ppathInput tapi path computation computep2ppath input
// swagger:model tapi.path.computation.computep2ppath.Input
type TapiPathComputationComputep2ppathInput struct {

	// none
	ObjectiveFunction *TapiPathComputationPathObjectiveFunction `json:"objective-function,omitempty"`

	// none
	RoutingConstraint *TapiPathComputationRoutingConstraint `json:"routing-constraint,omitempty"`

	// none
	Sep []*TapiPathComputationPathServiceEndPoint `json:"sep"`

	// none
	TopologyConstraint *TapiPathComputationTopologyConstraint `json:"topology-constraint,omitempty"`
}

// Validate validates this tapi path computation computep2ppath input
func (m *TapiPathComputationComputep2ppathInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectiveFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingConstraint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSep(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopologyConstraint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiPathComputationComputep2ppathInput) validateObjectiveFunction(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectiveFunction) { // not required
		return nil
	}

	if m.ObjectiveFunction != nil {
		if err := m.ObjectiveFunction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objective-function")
			}
			return err
		}
	}

	return nil
}

func (m *TapiPathComputationComputep2ppathInput) validateRoutingConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.RoutingConstraint) { // not required
		return nil
	}

	if m.RoutingConstraint != nil {
		if err := m.RoutingConstraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routing-constraint")
			}
			return err
		}
	}

	return nil
}

func (m *TapiPathComputationComputep2ppathInput) validateSep(formats strfmt.Registry) error {

	if swag.IsZero(m.Sep) { // not required
		return nil
	}

	for i := 0; i < len(m.Sep); i++ {
		if swag.IsZero(m.Sep[i]) { // not required
			continue
		}

		if m.Sep[i] != nil {
			if err := m.Sep[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sep" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiPathComputationComputep2ppathInput) validateTopologyConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.TopologyConstraint) { // not required
		return nil
	}

	if m.TopologyConstraint != nil {
		if err := m.TopologyConstraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topology-constraint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiPathComputationComputep2ppathInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPathComputationComputep2ppathInput) UnmarshalBinary(b []byte) error {
	var res TapiPathComputationComputep2ppathInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}