// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEquipmentAbstractStrand tapi equipment abstract strand
// swagger:model tapi.equipment.AbstractStrand
type TapiEquipmentAbstractStrand struct {
	TapiCommonLocalClass

	// none
	AdjacentStrand []*TapiEquipmentAbstractStrandRef `json:"adjacent-strand"`

	// A strand can end on two or more Pins (usually 2 pins, but a strand my be spliced to split a signal). This model supports only 2 ended strands.
	//                 A abstract strand may be spliced at both ends and hence have no direct relationship to pins or may be connected to pins at one or both ends.
	//                 In the essential model these Pins would be on connectors that plug in to connectors on Equipments.
	//                 The AbstractStrand is extended to the pins of the AccessPort which are the Pins on the Connectors of the Equipment.
	//                 In some cases it may not be relevant to represent the pin detail and hence the reference is to a connector alone.
	ConnectorPin []*TapiEquipmentConnectorPinAddress `json:"connector-pin"`

	// none
	SplicedStrand []*TapiEquipmentAbstractStrandRef `json:"spliced-strand"`

	// Relevant physical properties of the abstract strand.
	StrandMediaCharacteristics []*TapiCommonNameAndValue `json:"strand-media-characteristics"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiEquipmentAbstractStrand) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonLocalClass
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonLocalClass = aO0

	// AO1
	var dataAO1 struct {
		AdjacentStrand []*TapiEquipmentAbstractStrandRef `json:"adjacent-strand"`

		ConnectorPin []*TapiEquipmentConnectorPinAddress `json:"connector-pin"`

		SplicedStrand []*TapiEquipmentAbstractStrandRef `json:"spliced-strand"`

		StrandMediaCharacteristics []*TapiCommonNameAndValue `json:"strand-media-characteristics"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AdjacentStrand = dataAO1.AdjacentStrand

	m.ConnectorPin = dataAO1.ConnectorPin

	m.SplicedStrand = dataAO1.SplicedStrand

	m.StrandMediaCharacteristics = dataAO1.StrandMediaCharacteristics

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiEquipmentAbstractStrand) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiCommonLocalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AdjacentStrand []*TapiEquipmentAbstractStrandRef `json:"adjacent-strand"`

		ConnectorPin []*TapiEquipmentConnectorPinAddress `json:"connector-pin"`

		SplicedStrand []*TapiEquipmentAbstractStrandRef `json:"spliced-strand"`

		StrandMediaCharacteristics []*TapiCommonNameAndValue `json:"strand-media-characteristics"`
	}

	dataAO1.AdjacentStrand = m.AdjacentStrand

	dataAO1.ConnectorPin = m.ConnectorPin

	dataAO1.SplicedStrand = m.SplicedStrand

	dataAO1.StrandMediaCharacteristics = m.StrandMediaCharacteristics

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi equipment abstract strand
func (m *TapiEquipmentAbstractStrand) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonLocalClass
	if err := m.TapiCommonLocalClass.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdjacentStrand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectorPin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSplicedStrand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrandMediaCharacteristics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEquipmentAbstractStrand) validateAdjacentStrand(formats strfmt.Registry) error {

	if swag.IsZero(m.AdjacentStrand) { // not required
		return nil
	}

	for i := 0; i < len(m.AdjacentStrand); i++ {
		if swag.IsZero(m.AdjacentStrand[i]) { // not required
			continue
		}

		if m.AdjacentStrand[i] != nil {
			if err := m.AdjacentStrand[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adjacent-strand" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiEquipmentAbstractStrand) validateConnectorPin(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectorPin) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectorPin); i++ {
		if swag.IsZero(m.ConnectorPin[i]) { // not required
			continue
		}

		if m.ConnectorPin[i] != nil {
			if err := m.ConnectorPin[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connector-pin" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiEquipmentAbstractStrand) validateSplicedStrand(formats strfmt.Registry) error {

	if swag.IsZero(m.SplicedStrand) { // not required
		return nil
	}

	for i := 0; i < len(m.SplicedStrand); i++ {
		if swag.IsZero(m.SplicedStrand[i]) { // not required
			continue
		}

		if m.SplicedStrand[i] != nil {
			if err := m.SplicedStrand[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spliced-strand" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiEquipmentAbstractStrand) validateStrandMediaCharacteristics(formats strfmt.Registry) error {

	if swag.IsZero(m.StrandMediaCharacteristics) { // not required
		return nil
	}

	for i := 0; i < len(m.StrandMediaCharacteristics); i++ {
		if swag.IsZero(m.StrandMediaCharacteristics[i]) { // not required
			continue
		}

		if m.StrandMediaCharacteristics[i] != nil {
			if err := m.StrandMediaCharacteristics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("strand-media-characteristics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEquipmentAbstractStrand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEquipmentAbstractStrand) UnmarshalBinary(b []byte) error {
	var res TapiEquipmentAbstractStrand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
