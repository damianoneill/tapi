// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiCommonCapacityValue tapi common capacity value
// swagger:model tapi.common.CapacityValue
type TapiCommonCapacityValue struct {

	// none
	Unit TapiCommonCapacityUnit `json:"unit,omitempty"`

	// none
	Value int32 `json:"value,omitempty"`
}

// Validate validates this tapi common capacity value
func (m *TapiCommonCapacityValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiCommonCapacityValue) validateUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	if err := m.Unit.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("unit")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiCommonCapacityValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiCommonCapacityValue) UnmarshalBinary(b []byte) error {
	var res TapiCommonCapacityValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}