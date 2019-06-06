// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiPathComputationDeletep2ppathInput tapi path computation deletep2ppath input
// swagger:model tapi.path.computation.deletep2ppath.Input
type TapiPathComputationDeletep2ppathInput struct {

	// none
	PathIDOrName string `json:"path-id-or-name,omitempty"`
}

// Validate validates this tapi path computation deletep2ppath input
func (m *TapiPathComputationDeletep2ppathInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiPathComputationDeletep2ppathInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPathComputationDeletep2ppathInput) UnmarshalBinary(b []byte) error {
	var res TapiPathComputationDeletep2ppathInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}