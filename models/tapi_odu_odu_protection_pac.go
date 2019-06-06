// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiOduOduProtectionPac tapi odu odu protection pac
// swagger:model tapi.odu.OduProtectionPac
type TapiOduOduProtectionPac struct {

	// This attribute is for enabling/disabling the automatic protection switching (APS) capability at the transport adaptation function that is represented by the ODU_ConnectionTerminationPoint object class. It triggers the MI_APS_EN signal to the transport adaptation function.
	ApsEnable *bool `json:"aps-enable,omitempty"`

	// This attribute is for configuring the automatic protection switching (APS) level that should operate at the transport adaptation function that is represented by the ODU_ConnectionTerminationPoint object class. It triggers the MI_APS_LVL signal to the transport adaptation function. The value 0 means path and the values 1 through 6 mean TCM level 1 through 6 respectively.
	ApsLevel int32 `json:"aps-level,omitempty"`
}

// Validate validates this tapi odu odu protection pac
func (m *TapiOduOduProtectionPac) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiOduOduProtectionPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOduOduProtectionPac) UnmarshalBinary(b []byte) error {
	var res TapiOduOduProtectionPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
