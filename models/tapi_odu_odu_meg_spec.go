// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiOduOduMegSpec tapi odu odu meg spec
// swagger:model tapi.odu.OduMegSpec
type TapiOduOduMegSpec struct {

	// none
	MegLevel int32 `json:"meg-level,omitempty"`
}

// Validate validates this tapi odu odu meg spec
func (m *TapiOduOduMegSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOduOduMegSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOduOduMegSpec) UnmarshalBinary(b []byte) error {
	var res TapiOduOduMegSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}