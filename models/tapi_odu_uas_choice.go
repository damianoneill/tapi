// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiOduUasChoice tapi odu uas choice
// swagger:model tapi.odu.UasChoice
type TapiOduUasChoice struct {

	// none
	Bidirectional *bool `json:"bidirectional,omitempty"`

	// none
	Fuas int32 `json:"fuas,omitempty"`

	// none
	Nuas int32 `json:"nuas,omitempty"`

	// none
	Uas int32 `json:"uas,omitempty"`
}

// Validate validates this tapi odu uas choice
func (m *TapiOduUasChoice) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOduUasChoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOduUasChoice) UnmarshalBinary(b []byte) error {
	var res TapiOduUasChoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}