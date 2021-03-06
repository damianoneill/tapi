// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiConnectivityRouteRef tapi connectivity route ref
// swagger:model tapi.connectivity.RouteRef
type TapiConnectivityRouteRef struct {
	TapiConnectivityConnectionRef

	// none
	RouteLocalID string `json:"route-local-id,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiConnectivityRouteRef) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiConnectivityConnectionRef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiConnectivityConnectionRef = aO0

	// AO1
	var dataAO1 struct {
		RouteLocalID string `json:"route-local-id,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.RouteLocalID = dataAO1.RouteLocalID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiConnectivityRouteRef) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiConnectivityConnectionRef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		RouteLocalID string `json:"route-local-id,omitempty"`
	}

	dataAO1.RouteLocalID = m.RouteLocalID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi connectivity route ref
func (m *TapiConnectivityRouteRef) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiConnectivityConnectionRef
	if err := m.TapiConnectivityConnectionRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityRouteRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityRouteRef) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityRouteRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
