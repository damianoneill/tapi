// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthMegSpec tapi eth eth meg spec
// swagger:model tapi.eth.EthMegSpec
type TapiEthEthMegSpec struct {

	// none
	EthCfmMaintenanceAssociation *TapiEthEthCfmMaintenanceAssociation `json:"eth-cfm-maintenance-association,omitempty"`

	// none
	EthCfmMaintenanceDomain *TapiEthEthCfmMaintenanceDomain `json:"eth-cfm-maintenance-domain,omitempty"`

	// none
	EthMegCommon *TapiEthEthMegCommon `json:"eth-meg-common,omitempty"`
}

// Validate validates this tapi eth eth meg spec
func (m *TapiEthEthMegSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthCfmMaintenanceAssociation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEthCfmMaintenanceDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEthMegCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthMegSpec) validateEthCfmMaintenanceAssociation(formats strfmt.Registry) error {

	if swag.IsZero(m.EthCfmMaintenanceAssociation) { // not required
		return nil
	}

	if m.EthCfmMaintenanceAssociation != nil {
		if err := m.EthCfmMaintenanceAssociation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-cfm-maintenance-association")
			}
			return err
		}
	}

	return nil
}

func (m *TapiEthEthMegSpec) validateEthCfmMaintenanceDomain(formats strfmt.Registry) error {

	if swag.IsZero(m.EthCfmMaintenanceDomain) { // not required
		return nil
	}

	if m.EthCfmMaintenanceDomain != nil {
		if err := m.EthCfmMaintenanceDomain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-cfm-maintenance-domain")
			}
			return err
		}
	}

	return nil
}

func (m *TapiEthEthMegSpec) validateEthMegCommon(formats strfmt.Registry) error {

	if swag.IsZero(m.EthMegCommon) { // not required
		return nil
	}

	if m.EthMegCommon != nil {
		if err := m.EthMegCommon.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-meg-common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthMegSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthMegSpec) UnmarshalBinary(b []byte) error {
	var res TapiEthEthMegSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
