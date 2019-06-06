// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiOduOduTcmMipPac tapi odu odu tcm mip pac
// swagger:model tapi.odu.OduTcmMipPac
type TapiOduOduTcmMipPac struct {

	// This attribute indicates the tandem connection monitoring field of the ODU OH.
	TcmField int32 `json:"tcm-field,omitempty"`
}

// Validate validates this tapi odu odu tcm mip pac
func (m *TapiOduOduTcmMipPac) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOduOduTcmMipPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOduOduTcmMipPac) UnmarshalBinary(b []byte) error {
	var res TapiOduOduTcmMipPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}