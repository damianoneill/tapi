// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiVirtualNetworkCreateVirtualNetworkServiceOutput tapi virtual network create virtual network service output
// swagger:model tapi.virtual.network.CreateVirtualNetworkServiceOutput
type TapiVirtualNetworkCreateVirtualNetworkServiceOutput struct {

	// output
	Output *TapiVirtualNetworkCreateVirtualNetworkServiceOutputOutput `json:"output,omitempty"`
}

// Validate validates this tapi virtual network create virtual network service output
func (m *TapiVirtualNetworkCreateVirtualNetworkServiceOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiVirtualNetworkCreateVirtualNetworkServiceOutput) validateOutput(formats strfmt.Registry) error {

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
func (m *TapiVirtualNetworkCreateVirtualNetworkServiceOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiVirtualNetworkCreateVirtualNetworkServiceOutput) UnmarshalBinary(b []byte) error {
	var res TapiVirtualNetworkCreateVirtualNetworkServiceOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiVirtualNetworkCreateVirtualNetworkServiceOutputOutput tapi virtual network create virtual network service output output
// swagger:model TapiVirtualNetworkCreateVirtualNetworkServiceOutputOutput
type TapiVirtualNetworkCreateVirtualNetworkServiceOutputOutput struct {

	// none
	Service *TapiVirtualNetworkVirtualNetworkService `json:"service,omitempty"`
}

// Validate validates this tapi virtual network create virtual network service output output
func (m *TapiVirtualNetworkCreateVirtualNetworkServiceOutputOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiVirtualNetworkCreateVirtualNetworkServiceOutputOutput) validateService(formats strfmt.Registry) error {

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
func (m *TapiVirtualNetworkCreateVirtualNetworkServiceOutputOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiVirtualNetworkCreateVirtualNetworkServiceOutputOutput) UnmarshalBinary(b []byte) error {
	var res TapiVirtualNetworkCreateVirtualNetworkServiceOutputOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
