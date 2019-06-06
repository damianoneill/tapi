// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiConnectivityGetconnectionendpointdetailsOutputConnectionEndPoint tapi connectivity getconnectionendpointdetails output connection end point
// swagger:model tapi.connectivity.getconnectionendpointdetails.output.ConnectionEndPoint
type TapiConnectivityGetconnectionendpointdetailsOutputConnectionEndPoint struct {
	TapiConnectivityConnectionEndPoint

	TapiPhotonicMediaConnectionEndPointAugmentation1

	TapiPhotonicMediaConnectionEndPointAugmentation2

	TapiPhotonicMediaConnectionEndPointAugmentation3
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiConnectivityGetconnectionendpointdetailsOutputConnectionEndPoint) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiConnectivityConnectionEndPoint
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiConnectivityConnectionEndPoint = aO0

	// AO1
	var aO1 TapiPhotonicMediaConnectionEndPointAugmentation1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiPhotonicMediaConnectionEndPointAugmentation1 = aO1

	// AO2
	var aO2 TapiPhotonicMediaConnectionEndPointAugmentation2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiPhotonicMediaConnectionEndPointAugmentation2 = aO2

	// AO3
	var aO3 TapiPhotonicMediaConnectionEndPointAugmentation3
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.TapiPhotonicMediaConnectionEndPointAugmentation3 = aO3

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiConnectivityGetconnectionendpointdetailsOutputConnectionEndPoint) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 4)

	aO0, err := swag.WriteJSON(m.TapiConnectivityConnectionEndPoint)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiPhotonicMediaConnectionEndPointAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiPhotonicMediaConnectionEndPointAugmentation2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.TapiPhotonicMediaConnectionEndPointAugmentation3)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi connectivity getconnectionendpointdetails output connection end point
func (m *TapiConnectivityGetconnectionendpointdetailsOutputConnectionEndPoint) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiConnectivityConnectionEndPoint
	if err := m.TapiConnectivityConnectionEndPoint.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiPhotonicMediaConnectionEndPointAugmentation1
	if err := m.TapiPhotonicMediaConnectionEndPointAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiPhotonicMediaConnectionEndPointAugmentation2
	if err := m.TapiPhotonicMediaConnectionEndPointAugmentation2.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiPhotonicMediaConnectionEndPointAugmentation3
	if err := m.TapiPhotonicMediaConnectionEndPointAugmentation3.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityGetconnectionendpointdetailsOutputConnectionEndPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityGetconnectionendpointdetailsOutputConnectionEndPoint) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityGetconnectionendpointdetailsOutputConnectionEndPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
