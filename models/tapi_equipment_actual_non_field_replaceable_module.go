// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEquipmentActualNonFieldReplaceableModule tapi equipment actual non field replaceable module
// swagger:model tapi.equipment.ActualNonFieldReplaceableModule
type TapiEquipmentActualNonFieldReplaceableModule struct {

	// none
	CommonActualProperties *TapiEquipmentCommonActualProperties `json:"common-actual-properties,omitempty"`

	// none
	CommonEquipmentProperties *TapiEquipmentCommonEquipmentProperties `json:"common-equipment-properties,omitempty"`
}

// Validate validates this tapi equipment actual non field replaceable module
func (m *TapiEquipmentActualNonFieldReplaceableModule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommonActualProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommonEquipmentProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEquipmentActualNonFieldReplaceableModule) validateCommonActualProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.CommonActualProperties) { // not required
		return nil
	}

	if m.CommonActualProperties != nil {
		if err := m.CommonActualProperties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("common-actual-properties")
			}
			return err
		}
	}

	return nil
}

func (m *TapiEquipmentActualNonFieldReplaceableModule) validateCommonEquipmentProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.CommonEquipmentProperties) { // not required
		return nil
	}

	if m.CommonEquipmentProperties != nil {
		if err := m.CommonEquipmentProperties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("common-equipment-properties")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEquipmentActualNonFieldReplaceableModule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEquipmentActualNonFieldReplaceableModule) UnmarshalBinary(b []byte) error {
	var res TapiEquipmentActualNonFieldReplaceableModule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
