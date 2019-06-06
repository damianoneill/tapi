// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamCreateOamServiceEndPointOutput tapi oam create oam service end point output
// swagger:model tapi.oam.CreateOamServiceEndPointOutput
type TapiOamCreateOamServiceEndPointOutput struct {

	// output
	Output *TapiOamCreateOamServiceEndPointOutputOutput `json:"output,omitempty"`
}

// Validate validates this tapi oam create oam service end point output
func (m *TapiOamCreateOamServiceEndPointOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamCreateOamServiceEndPointOutput) validateOutput(formats strfmt.Registry) error {

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
func (m *TapiOamCreateOamServiceEndPointOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamCreateOamServiceEndPointOutput) UnmarshalBinary(b []byte) error {
	var res TapiOamCreateOamServiceEndPointOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiOamCreateOamServiceEndPointOutputOutput tapi oam create oam service end point output output
// swagger:model TapiOamCreateOamServiceEndPointOutputOutput
type TapiOamCreateOamServiceEndPointOutputOutput struct {

	// none
	EndPoint *TapiOamOamServiceEndPoint `json:"end-point,omitempty"`
}

// Validate validates this tapi oam create oam service end point output output
func (m *TapiOamCreateOamServiceEndPointOutputOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndPoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamCreateOamServiceEndPointOutputOutput) validateEndPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.EndPoint) { // not required
		return nil
	}

	if m.EndPoint != nil {
		if err := m.EndPoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output" + "." + "end-point")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamCreateOamServiceEndPointOutputOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamCreateOamServiceEndPointOutputOutput) UnmarshalBinary(b []byte) error {
	var res TapiOamCreateOamServiceEndPointOutputOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}