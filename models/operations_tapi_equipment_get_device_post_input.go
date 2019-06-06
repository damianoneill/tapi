// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OperationsTapiEquipmentGetDevicePostInput operations tapi equipment get device post input
// swagger:model operations_tapi-equipment_get-device-post-input
type OperationsTapiEquipmentGetDevicePostInput struct {

	// none (leaf)
	TapiEquipmentDeviceID string `json:"tapi-equipment:device-id,omitempty"`
}

// Validate validates this operations tapi equipment get device post input
func (m *OperationsTapiEquipmentGetDevicePostInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetDevicePostInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetDevicePostInput) UnmarshalBinary(b []byte) error {
	var res OperationsTapiEquipmentGetDevicePostInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
