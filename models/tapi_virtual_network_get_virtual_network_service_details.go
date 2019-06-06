// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiVirtualNetworkGetVirtualNetworkServiceDetails tapi virtual network get virtual network service details
// swagger:model tapi.virtual.network.GetVirtualNetworkServiceDetails
type TapiVirtualNetworkGetVirtualNetworkServiceDetails struct {

	// output
	Output *TapiVirtualNetworkGetvirtualnetworkservicedetailsOutput `json:"output,omitempty"`
}

// Validate validates this tapi virtual network get virtual network service details
func (m *TapiVirtualNetworkGetVirtualNetworkServiceDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiVirtualNetworkGetVirtualNetworkServiceDetails) validateOutput(formats strfmt.Registry) error {

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
func (m *TapiVirtualNetworkGetVirtualNetworkServiceDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiVirtualNetworkGetVirtualNetworkServiceDetails) UnmarshalBinary(b []byte) error {
	var res TapiVirtualNetworkGetVirtualNetworkServiceDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
