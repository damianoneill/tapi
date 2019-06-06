// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthCfmMaintenanceAssociation tapi eth eth cfm maintenance association
// swagger:model tapi.eth.EthCfmMaintenanceAssociation
type TapiEthEthCfmMaintenanceAssociation struct {

	// IEEE P802.1Qcx/D0.3:
	//                 MEF 38:
	//                 This parameter indicates what, if anything, is to be included in the Sender ID TLV transmitted by Maintenance Points configured in this MA.
	//                 A value of 'defer' means that the contents of the Sender ID TLV are determined by the enclosing Maintenance Domain instance.
	IDPermission string `json:"id-permission,omitempty"`

	// IEEE P802.1Qcx/D0.3:
	//                 MEF 38:
	//                 The Maintenance Association name and name format choice.
	MaintenanceAssociationName *TapiEthMaintenanceAssociationName `json:"maintenance-association-name,omitempty"`
}

// Validate validates this tapi eth eth cfm maintenance association
func (m *TapiEthEthCfmMaintenanceAssociation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaintenanceAssociationName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthCfmMaintenanceAssociation) validateMaintenanceAssociationName(formats strfmt.Registry) error {

	if swag.IsZero(m.MaintenanceAssociationName) { // not required
		return nil
	}

	if m.MaintenanceAssociationName != nil {
		if err := m.MaintenanceAssociationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance-association-name")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthCfmMaintenanceAssociation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthCfmMaintenanceAssociation) UnmarshalBinary(b []byte) error {
	var res TapiEthEthCfmMaintenanceAssociation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
