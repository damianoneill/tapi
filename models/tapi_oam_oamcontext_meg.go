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

// TapiOamOamcontextMeg tapi oam oamcontext meg
// swagger:model tapi.oam.oamcontext.Meg
type TapiOamOamcontextMeg struct {
	TapiCommonGlobalClass

	TapiCommonOperationalStatePac

	TapiEthMegAugmentation2

	TapiOduMegAugmentation1

	// none
	LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

	// 1. ME may have 0 MEPs (case of transit domains where at least 1 MIP is present)
	//                 2. ME may have 1 MEP (case of edge domaind, where the peer MEP is ouside the managed domain)
	//                 3. ME may have 2 MEPs
	Mep []*TapiOamMegMep `json:"mep"`

	// ME may 0, 1, or more MIPs
	Mip []*TapiOamMegMip `json:"mip"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiOamOamcontextMeg) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonGlobalClass
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonGlobalClass = aO0

	// AO1
	var aO1 TapiCommonOperationalStatePac
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiCommonOperationalStatePac = aO1

	// AO2
	var aO2 TapiEthMegAugmentation2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiEthMegAugmentation2 = aO2

	// AO3
	var aO3 TapiOduMegAugmentation1
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.TapiOduMegAugmentation1 = aO3

	// AO4
	var dataAO4 struct {
		LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

		Mep []*TapiOamMegMep `json:"mep"`

		Mip []*TapiOamMegMip `json:"mip"`
	}
	if err := swag.ReadJSON(raw, &dataAO4); err != nil {
		return err
	}

	m.LayerProtocolName = dataAO4.LayerProtocolName

	m.Mep = dataAO4.Mep

	m.Mip = dataAO4.Mip

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiOamOamcontextMeg) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 5)

	aO0, err := swag.WriteJSON(m.TapiCommonGlobalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiCommonOperationalStatePac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiEthMegAugmentation2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.TapiOduMegAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	var dataAO4 struct {
		LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

		Mep []*TapiOamMegMep `json:"mep"`

		Mip []*TapiOamMegMip `json:"mip"`
	}

	dataAO4.LayerProtocolName = m.LayerProtocolName

	dataAO4.Mep = m.Mep

	dataAO4.Mip = m.Mip

	jsonDataAO4, errAO4 := swag.WriteJSON(dataAO4)
	if errAO4 != nil {
		return nil, errAO4
	}
	_parts = append(_parts, jsonDataAO4)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi oam oamcontext meg
func (m *TapiOamOamcontextMeg) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonOperationalStatePac
	if err := m.TapiCommonOperationalStatePac.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthMegAugmentation2
	if err := m.TapiEthMegAugmentation2.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiOduMegAugmentation1
	if err := m.TapiOduMegAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLayerProtocolName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMep(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMip(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamOamcontextMeg) validateLayerProtocolName(formats strfmt.Registry) error {

	if swag.IsZero(m.LayerProtocolName) { // not required
		return nil
	}

	if err := m.LayerProtocolName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("layer-protocol-name")
		}
		return err
	}

	return nil
}

func (m *TapiOamOamcontextMeg) validateMep(formats strfmt.Registry) error {

	if swag.IsZero(m.Mep) { // not required
		return nil
	}

	for i := 0; i < len(m.Mep); i++ {
		if swag.IsZero(m.Mep[i]) { // not required
			continue
		}

		if m.Mep[i] != nil {
			if err := m.Mep[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mep" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiOamOamcontextMeg) validateMip(formats strfmt.Registry) error {

	if swag.IsZero(m.Mip) { // not required
		return nil
	}

	for i := 0; i < len(m.Mip); i++ {
		if swag.IsZero(m.Mip[i]) { // not required
			continue
		}

		if m.Mip[i] != nil {
			if err := m.Mip[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mip" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamOamcontextMeg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamOamcontextMeg) UnmarshalBinary(b []byte) error {
	var res TapiOamOamcontextMeg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
