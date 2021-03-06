// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiOamGetoamjobInput tapi oam getoamjob input
// swagger:model tapi.oam.getoamjob.Input
type TapiOamGetoamjobInput struct {

	// none
	OamJobID string `json:"oam-job-id,omitempty"`
}

// Validate validates this tapi oam getoamjob input
func (m *TapiOamGetoamjobInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamGetoamjobInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamGetoamjobInput) UnmarshalBinary(b []byte) error {
	var res TapiOamGetoamjobInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
