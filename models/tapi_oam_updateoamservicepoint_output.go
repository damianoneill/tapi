// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamUpdateoamservicepointOutput tapi oam updateoamservicepoint output
// swagger:model tapi.oam.updateoamservicepoint.Output
type TapiOamUpdateoamservicepointOutput struct {

	// none
	OamServicePoint *TapiOamOamserviceOamServicePoint `json:"oam-service-point,omitempty"`
}

// Validate validates this tapi oam updateoamservicepoint output
func (m *TapiOamUpdateoamservicepointOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOamServicePoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamUpdateoamservicepointOutput) validateOamServicePoint(formats strfmt.Registry) error {

	if swag.IsZero(m.OamServicePoint) { // not required
		return nil
	}

	if m.OamServicePoint != nil {
		if err := m.OamServicePoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oam-service-point")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamUpdateoamservicepointOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamUpdateoamservicepointOutput) UnmarshalBinary(b []byte) error {
	var res TapiOamUpdateoamservicepointOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
