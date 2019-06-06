// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DataTapiConnectivityGetConnectionDetailsPostInput data tapi connectivity get connection details post input
// swagger:model data_tapi-connectivity_get-connection-details-post-input
type DataTapiConnectivityGetConnectionDetailsPostInput struct {

	// none (leaf)
	TapiConnectivityConnectionIDOrName string `json:"tapi-connectivity:connection-id-or-name,omitempty"`
}

// Validate validates this data tapi connectivity get connection details post input
func (m *DataTapiConnectivityGetConnectionDetailsPostInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiConnectivityGetConnectionDetailsPostInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiConnectivityGetConnectionDetailsPostInput) UnmarshalBinary(b []byte) error {
	var res DataTapiConnectivityGetConnectionDetailsPostInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}