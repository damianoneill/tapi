// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiVirtualNetworkCreateVirtualNetworkService tapi virtual network create virtual network service
// swagger:model tapi.virtual.network.CreateVirtualNetworkService
type TapiVirtualNetworkCreateVirtualNetworkService struct {

	// output
	Output *TapiVirtualNetworkCreatevirtualnetworkserviceOutput `json:"output,omitempty"`
}

// Validate validates this tapi virtual network create virtual network service
func (m *TapiVirtualNetworkCreateVirtualNetworkService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiVirtualNetworkCreateVirtualNetworkService) validateOutput(formats strfmt.Registry) error {

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
func (m *TapiVirtualNetworkCreateVirtualNetworkService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiVirtualNetworkCreateVirtualNetworkService) UnmarshalBinary(b []byte) error {
	var res TapiVirtualNetworkCreateVirtualNetworkService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}