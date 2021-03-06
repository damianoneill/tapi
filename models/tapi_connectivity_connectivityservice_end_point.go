// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiConnectivityConnectivityserviceEndPoint tapi connectivity connectivityservice end point
// swagger:model tapi.connectivity.connectivityservice.EndPoint
type TapiConnectivityConnectivityserviceEndPoint struct {
	TapiConnectivityConnectivityServiceEndPoint

	TapiEthEndPointAugmentation2

	TapiOduEndPointAugmentation1

	TapiPhotonicMediaEndPointAugmentation3

	TapiPhotonicMediaEndPointAugmentation4
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiConnectivityConnectivityserviceEndPoint) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiConnectivityConnectivityServiceEndPoint
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiConnectivityConnectivityServiceEndPoint = aO0

	// AO1
	var aO1 TapiEthEndPointAugmentation2
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiEthEndPointAugmentation2 = aO1

	// AO2
	var aO2 TapiOduEndPointAugmentation1
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiOduEndPointAugmentation1 = aO2

	// AO3
	var aO3 TapiPhotonicMediaEndPointAugmentation3
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.TapiPhotonicMediaEndPointAugmentation3 = aO3

	// AO4
	var aO4 TapiPhotonicMediaEndPointAugmentation4
	if err := swag.ReadJSON(raw, &aO4); err != nil {
		return err
	}
	m.TapiPhotonicMediaEndPointAugmentation4 = aO4

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiConnectivityConnectivityserviceEndPoint) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 5)

	aO0, err := swag.WriteJSON(m.TapiConnectivityConnectivityServiceEndPoint)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiEthEndPointAugmentation2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiOduEndPointAugmentation1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.TapiPhotonicMediaEndPointAugmentation3)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	aO4, err := swag.WriteJSON(m.TapiPhotonicMediaEndPointAugmentation4)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO4)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi connectivity connectivityservice end point
func (m *TapiConnectivityConnectivityserviceEndPoint) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiConnectivityConnectivityServiceEndPoint
	if err := m.TapiConnectivityConnectivityServiceEndPoint.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiEthEndPointAugmentation2
	if err := m.TapiEthEndPointAugmentation2.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiOduEndPointAugmentation1
	if err := m.TapiOduEndPointAugmentation1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiPhotonicMediaEndPointAugmentation3
	if err := m.TapiPhotonicMediaEndPointAugmentation3.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiPhotonicMediaEndPointAugmentation4
	if err := m.TapiPhotonicMediaEndPointAugmentation4.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityConnectivityserviceEndPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityConnectivityserviceEndPoint) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityConnectivityserviceEndPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
