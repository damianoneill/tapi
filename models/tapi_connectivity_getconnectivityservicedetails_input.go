// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiConnectivityGetconnectivityservicedetailsInput tapi connectivity getconnectivityservicedetails input
// swagger:model tapi.connectivity.getconnectivityservicedetails.Input
type TapiConnectivityGetconnectivityservicedetailsInput struct {

	// none
	ServiceIDOrName string `json:"service-id-or-name,omitempty"`
}

// Validate validates this tapi connectivity getconnectivityservicedetails input
func (m *TapiConnectivityGetconnectivityservicedetailsInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityGetconnectivityservicedetailsInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityGetconnectivityservicedetailsInput) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityGetconnectivityservicedetailsInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
