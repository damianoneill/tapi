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

// TapiVirtualNetworkCreatevirtualnetworkserviceInput tapi virtual network createvirtualnetworkservice input
// swagger:model tapi.virtual.network.createvirtualnetworkservice.Input
type TapiVirtualNetworkCreatevirtualnetworkserviceInput struct {

	// none
	ConnSchedule string `json:"conn-schedule,omitempty"`

	// none
	Sep []*TapiVirtualNetworkVirtualNetworkServiceEndPoint `json:"sep"`

	// none
	VnwConstraint *TapiVirtualNetworkVirtualNetworkConstraint `json:"vnw-constraint,omitempty"`
}

// Validate validates this tapi virtual network createvirtualnetworkservice input
func (m *TapiVirtualNetworkCreatevirtualnetworkserviceInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSep(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVnwConstraint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiVirtualNetworkCreatevirtualnetworkserviceInput) validateSep(formats strfmt.Registry) error {

	if swag.IsZero(m.Sep) { // not required
		return nil
	}

	for i := 0; i < len(m.Sep); i++ {
		if swag.IsZero(m.Sep[i]) { // not required
			continue
		}

		if m.Sep[i] != nil {
			if err := m.Sep[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sep" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiVirtualNetworkCreatevirtualnetworkserviceInput) validateVnwConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.VnwConstraint) { // not required
		return nil
	}

	if m.VnwConstraint != nil {
		if err := m.VnwConstraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vnw-constraint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiVirtualNetworkCreatevirtualnetworkserviceInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiVirtualNetworkCreatevirtualnetworkserviceInput) UnmarshalBinary(b []byte) error {
	var res TapiVirtualNetworkCreatevirtualnetworkserviceInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
