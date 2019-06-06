// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamOamprofilePmThresholdData tapi oam oamprofile pm threshold data
// swagger:model tapi.oam.oamprofile.PmThresholdData
type TapiOamOamprofilePmThresholdData struct {
	TapiEthPmThresholdDataAugmentation1

	TapiEthPmThresholdDataAugmentation2

	TapiEthPmThresholdDataAugmentation3

	TapiEthPmThresholdDataAugmentation4

	TapiOamPmThresholdData
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiOamOamprofilePmThresholdData) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiEthPmThresholdDataAugmentation1
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiEthPmThresholdDataAugmentation1 = aO0

	// AO1
	var aO1 TapiEthPmThresholdDataAugmentation2
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiEthPmThresholdDataAugmentation2 = aO1

	// AO2
	var aO2 TapiEthPmThresholdDataAugmentation3
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiEthPmThresholdDataAugmentation3 = aO2

	// AO3
	var aO3 TapiEthPmThresholdDataAugmentation4
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.TapiEthPmThresholdDataAugmentation4 = aO3

	// AO4
	var aO4 TapiOamPmThresholdData
	if err := swag.ReadJSON(raw, &aO4); err != nil {
		return err
	}
	m.TapiOamPmThresholdData = aO4

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiOamOamprofilePmThresholdData) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 5)

	aO0, err := swag.WriteJSON(m.TapiEthPmThresholdDataAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiEthPmThresholdDataAugmentation2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiEthPmThresholdDataAugmentation3)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.TapiEthPmThresholdDataAugmentation4)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	aO4, err := swag.WriteJSON(m.TapiOamPmThresholdData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO4)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi oam oamprofile pm threshold data
func (m *TapiOamOamprofilePmThresholdData) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiEthPmThresholdDataAugmentation1
	if err := m.TapiEthPmThresholdDataAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthPmThresholdDataAugmentation2
	if err := m.TapiEthPmThresholdDataAugmentation2.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthPmThresholdDataAugmentation3
	if err := m.TapiEthPmThresholdDataAugmentation3.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthPmThresholdDataAugmentation4
	if err := m.TapiEthPmThresholdDataAugmentation4.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiOamPmThresholdData
	if err := m.TapiOamPmThresholdData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamOamprofilePmThresholdData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamOamprofilePmThresholdData) UnmarshalBinary(b []byte) error {
	var res TapiOamOamprofilePmThresholdData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}