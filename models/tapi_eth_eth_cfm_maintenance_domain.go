// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiEthEthCfmMaintenanceDomain tapi eth eth cfm maintenance domain
// swagger:model tapi.eth.EthCfmMaintenanceDomain
type TapiEthEthCfmMaintenanceDomain struct {

	// IEEE P802.1Qcx/D0.3:
	//                 MEF 38:
	//                 A reference to the maintenance domain that this maintenance group is associated with.
	MaintenanceDomainName string `json:"maintenance-domain-name,omitempty"`

	// IEEE P802.1Qcx/D0.3:
	//                 MEF 38:
	//                 The Maintenance Domain name format choice.
	MaintenanceDomainNameType string `json:"maintenance-domain-name-type,omitempty"`
}

// Validate validates this tapi eth eth cfm maintenance domain
func (m *TapiEthEthCfmMaintenanceDomain) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthCfmMaintenanceDomain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthCfmMaintenanceDomain) UnmarshalBinary(b []byte) error {
	var res TapiEthEthCfmMaintenanceDomain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
