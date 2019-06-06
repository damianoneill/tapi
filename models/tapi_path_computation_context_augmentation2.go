// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiPathComputationContextAugmentation2 tapi path computation context augmentation2
// swagger:model tapi.path.computation.ContextAugmentation2
type TapiPathComputationContextAugmentation2 struct {

	// Augments the base TAPI Context with PathComputationService information
	PathComputationContext *TapiPathComputationPathComputationContext `json:"path-computation-context,omitempty"`
}

// Validate validates this tapi path computation context augmentation2
func (m *TapiPathComputationContextAugmentation2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePathComputationContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiPathComputationContextAugmentation2) validatePathComputationContext(formats strfmt.Registry) error {

	if swag.IsZero(m.PathComputationContext) { // not required
		return nil
	}

	if m.PathComputationContext != nil {
		if err := m.PathComputationContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("path-computation-context")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiPathComputationContextAugmentation2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPathComputationContextAugmentation2) UnmarshalBinary(b []byte) error {
	var res TapiPathComputationContextAugmentation2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}