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

// TapiOamPerceivedSeverityType tapi oam perceived severity type
// swagger:model tapi.oam.PerceivedSeverityType
type TapiOamPerceivedSeverityType string

const (

	// TapiOamPerceivedSeverityTypeCRITICAL captures enum value "CRITICAL"
	TapiOamPerceivedSeverityTypeCRITICAL TapiOamPerceivedSeverityType = "CRITICAL"

	// TapiOamPerceivedSeverityTypeMAJOR captures enum value "MAJOR"
	TapiOamPerceivedSeverityTypeMAJOR TapiOamPerceivedSeverityType = "MAJOR"

	// TapiOamPerceivedSeverityTypeMINOR captures enum value "MINOR"
	TapiOamPerceivedSeverityTypeMINOR TapiOamPerceivedSeverityType = "MINOR"

	// TapiOamPerceivedSeverityTypeWARNING captures enum value "WARNING"
	TapiOamPerceivedSeverityTypeWARNING TapiOamPerceivedSeverityType = "WARNING"

	// TapiOamPerceivedSeverityTypeCLEARED captures enum value "CLEARED"
	TapiOamPerceivedSeverityTypeCLEARED TapiOamPerceivedSeverityType = "CLEARED"
)

// for schema
var tapiOamPerceivedSeverityTypeEnum []interface{}

func init() {
	var res []TapiOamPerceivedSeverityType
	if err := json.Unmarshal([]byte(`["CRITICAL","MAJOR","MINOR","WARNING","CLEARED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiOamPerceivedSeverityTypeEnum = append(tapiOamPerceivedSeverityTypeEnum, v)
	}
}

func (m TapiOamPerceivedSeverityType) validateTapiOamPerceivedSeverityTypeEnum(path, location string, value TapiOamPerceivedSeverityType) error {
	if err := validate.Enum(path, location, value, tapiOamPerceivedSeverityTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi oam perceived severity type
func (m TapiOamPerceivedSeverityType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiOamPerceivedSeverityTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
