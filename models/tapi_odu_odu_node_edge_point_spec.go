// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOduOduNodeEdgePointSpec tapi odu odu node edge point spec
// swagger:model tapi.odu.OduNodeEdgePointSpec
type TapiOduOduNodeEdgePointSpec struct {

	// none
	OduPool *TapiOduOduPoolPac `json:"odu-pool,omitempty"`
}

// Validate validates this tapi odu odu node edge point spec
func (m *TapiOduOduNodeEdgePointSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOduPool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOduOduNodeEdgePointSpec) validateOduPool(formats strfmt.Registry) error {

	if swag.IsZero(m.OduPool) { // not required
		return nil
	}

	if m.OduPool != nil {
		if err := m.OduPool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("odu-pool")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOduOduNodeEdgePointSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOduOduNodeEdgePointSpec) UnmarshalBinary(b []byte) error {
	var res TapiOduOduNodeEdgePointSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}