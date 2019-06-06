// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamOamServicePointRef tapi oam oam service point ref
// swagger:model tapi.oam.OamServicePointRef
type TapiOamOamServicePointRef struct {
	TapiOamOamServiceRef

	// none
	OamServicePointLocalID string `json:"oam-service-point-local-id,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiOamOamServicePointRef) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiOamOamServiceRef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiOamOamServiceRef = aO0

	// AO1
	var dataAO1 struct {
		OamServicePointLocalID string `json:"oam-service-point-local-id,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.OamServicePointLocalID = dataAO1.OamServicePointLocalID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiOamOamServicePointRef) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiOamOamServiceRef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		OamServicePointLocalID string `json:"oam-service-point-local-id,omitempty"`
	}

	dataAO1.OamServicePointLocalID = m.OamServicePointLocalID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi oam oam service point ref
func (m *TapiOamOamServicePointRef) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiOamOamServiceRef
	if err := m.TapiOamOamServiceRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamOamServicePointRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamOamServicePointRef) UnmarshalBinary(b []byte) error {
	var res TapiOamOamServicePointRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}