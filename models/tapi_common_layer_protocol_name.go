// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TapiCommonLayerProtocolName tapi common layer protocol name
// swagger:model tapi.common.LayerProtocolName
type TapiCommonLayerProtocolName string

const (

	// TapiCommonLayerProtocolNameODU captures enum value "ODU"
	TapiCommonLayerProtocolNameODU TapiCommonLayerProtocolName = "ODU"

	// TapiCommonLayerProtocolNameETH captures enum value "ETH"
	TapiCommonLayerProtocolNameETH TapiCommonLayerProtocolName = "ETH"

	// TapiCommonLayerProtocolNameDSR captures enum value "DSR"
	TapiCommonLayerProtocolNameDSR TapiCommonLayerProtocolName = "DSR"

	// TapiCommonLayerProtocolNamePHOTONICMEDIA captures enum value "PHOTONIC_MEDIA"
	TapiCommonLayerProtocolNamePHOTONICMEDIA TapiCommonLayerProtocolName = "PHOTONIC_MEDIA"
)

// for schema
var tapiCommonLayerProtocolNameEnum []interface{}

func init() {
	var res []TapiCommonLayerProtocolName
	if err := json.Unmarshal([]byte(`["ODU","ETH","DSR","PHOTONIC_MEDIA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiCommonLayerProtocolNameEnum = append(tapiCommonLayerProtocolNameEnum, v)
	}
}

func (m TapiCommonLayerProtocolName) validateTapiCommonLayerProtocolNameEnum(path, location string, value TapiCommonLayerProtocolName) error {
	if err := validate.Enum(path, location, value, tapiCommonLayerProtocolNameEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi common layer protocol name
func (m TapiCommonLayerProtocolName) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiCommonLayerProtocolNameEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
