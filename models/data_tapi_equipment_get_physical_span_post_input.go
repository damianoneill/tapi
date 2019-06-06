// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DataTapiEquipmentGetPhysicalSpanPostInput data tapi equipment get physical span post input
// swagger:model data_tapi-equipment_get-physical-span-post-input
type DataTapiEquipmentGetPhysicalSpanPostInput struct {

	// none (leaf)
	TapiEquipmentSpanID string `json:"tapi-equipment:span-id,omitempty"`
}

// Validate validates this data tapi equipment get physical span post input
func (m *DataTapiEquipmentGetPhysicalSpanPostInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostInput) UnmarshalBinary(b []byte) error {
	var res DataTapiEquipmentGetPhysicalSpanPostInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}