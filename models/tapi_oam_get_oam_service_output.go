// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamGetOamServiceOutput tapi oam get oam service output
// swagger:model tapi.oam.GetOamServiceOutput
type TapiOamGetOamServiceOutput struct {

	// output
	Output *TapiOamGetOamServiceOutputOutput `json:"output,omitempty"`
}

// Validate validates this tapi oam get oam service output
func (m *TapiOamGetOamServiceOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamGetOamServiceOutput) validateOutput(formats strfmt.Registry) error {

	if swag.IsZero(m.Output) { // not required
		return nil
	}

	if m.Output != nil {
		if err := m.Output.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamGetOamServiceOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamGetOamServiceOutput) UnmarshalBinary(b []byte) error {
	var res TapiOamGetOamServiceOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiOamGetOamServiceOutputOutput tapi oam get oam service output output
// swagger:model TapiOamGetOamServiceOutputOutput
type TapiOamGetOamServiceOutputOutput struct {

	// none
	Service *TapiOamOamService `json:"service,omitempty"`
}

// Validate validates this tapi oam get oam service output output
func (m *TapiOamGetOamServiceOutputOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamGetOamServiceOutputOutput) validateService(formats strfmt.Registry) error {

	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output" + "." + "service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamGetOamServiceOutputOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamGetOamServiceOutputOutput) UnmarshalBinary(b []byte) error {
	var res TapiOamGetOamServiceOutputOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
