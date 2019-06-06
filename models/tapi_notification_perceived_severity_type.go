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

// TapiNotificationPerceivedSeverityType tapi notification perceived severity type
// swagger:model tapi.notification.PerceivedSeverityType
type TapiNotificationPerceivedSeverityType string

const (

	// TapiNotificationPerceivedSeverityTypeCRITICAL captures enum value "CRITICAL"
	TapiNotificationPerceivedSeverityTypeCRITICAL TapiNotificationPerceivedSeverityType = "CRITICAL"

	// TapiNotificationPerceivedSeverityTypeMAJOR captures enum value "MAJOR"
	TapiNotificationPerceivedSeverityTypeMAJOR TapiNotificationPerceivedSeverityType = "MAJOR"

	// TapiNotificationPerceivedSeverityTypeMINOR captures enum value "MINOR"
	TapiNotificationPerceivedSeverityTypeMINOR TapiNotificationPerceivedSeverityType = "MINOR"

	// TapiNotificationPerceivedSeverityTypeWARNING captures enum value "WARNING"
	TapiNotificationPerceivedSeverityTypeWARNING TapiNotificationPerceivedSeverityType = "WARNING"

	// TapiNotificationPerceivedSeverityTypeCLEARED captures enum value "CLEARED"
	TapiNotificationPerceivedSeverityTypeCLEARED TapiNotificationPerceivedSeverityType = "CLEARED"
)

// for schema
var tapiNotificationPerceivedSeverityTypeEnum []interface{}

func init() {
	var res []TapiNotificationPerceivedSeverityType
	if err := json.Unmarshal([]byte(`["CRITICAL","MAJOR","MINOR","WARNING","CLEARED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiNotificationPerceivedSeverityTypeEnum = append(tapiNotificationPerceivedSeverityTypeEnum, v)
	}
}

func (m TapiNotificationPerceivedSeverityType) validateTapiNotificationPerceivedSeverityTypeEnum(path, location string, value TapiNotificationPerceivedSeverityType) error {
	if err := validate.Enum(path, location, value, tapiNotificationPerceivedSeverityTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi notification perceived severity type
func (m TapiNotificationPerceivedSeverityType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiNotificationPerceivedSeverityTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}