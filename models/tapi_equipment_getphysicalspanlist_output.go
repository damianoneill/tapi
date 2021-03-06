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

// TapiEquipmentGetphysicalspanlistOutput tapi equipment getphysicalspanlist output
// swagger:model tapi.equipment.getphysicalspanlist.Output
type TapiEquipmentGetphysicalspanlistOutput struct {

	// none
	PhysicalSpan []*TapiEquipmentPhysicalSpan `json:"physical-span"`
}

// Validate validates this tapi equipment getphysicalspanlist output
func (m *TapiEquipmentGetphysicalspanlistOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhysicalSpan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEquipmentGetphysicalspanlistOutput) validatePhysicalSpan(formats strfmt.Registry) error {

	if swag.IsZero(m.PhysicalSpan) { // not required
		return nil
	}

	for i := 0; i < len(m.PhysicalSpan); i++ {
		if swag.IsZero(m.PhysicalSpan[i]) { // not required
			continue
		}

		if m.PhysicalSpan[i] != nil {
			if err := m.PhysicalSpan[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("physical-span" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEquipmentGetphysicalspanlistOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEquipmentGetphysicalspanlistOutput) UnmarshalBinary(b []byte) error {
	var res TapiEquipmentGetphysicalspanlistOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
