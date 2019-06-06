// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamCreateoamprofileOutput tapi oam createoamprofile output
// swagger:model tapi.oam.createoamprofile.Output
type TapiOamCreateoamprofileOutput struct {

	// none
	OamProfile *TapiOamOamProfile `json:"oam-profile,omitempty"`
}

// Validate validates this tapi oam createoamprofile output
func (m *TapiOamCreateoamprofileOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOamProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamCreateoamprofileOutput) validateOamProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.OamProfile) { // not required
		return nil
	}

	if m.OamProfile != nil {
		if err := m.OamProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oam-profile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamCreateoamprofileOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamCreateoamprofileOutput) UnmarshalBinary(b []byte) error {
	var res TapiOamCreateoamprofileOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}