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

// TapiNotificationThresholdCrossingType tapi notification threshold crossing type
// swagger:model tapi.notification.ThresholdCrossingType
type TapiNotificationThresholdCrossingType string

const (

	// TapiNotificationThresholdCrossingTypeTHRESHOLDABOVE captures enum value "THRESHOLD_ABOVE"
	TapiNotificationThresholdCrossingTypeTHRESHOLDABOVE TapiNotificationThresholdCrossingType = "THRESHOLD_ABOVE"

	// TapiNotificationThresholdCrossingTypeTHRESHOLDBELOW captures enum value "THRESHOLD_BELOW"
	TapiNotificationThresholdCrossingTypeTHRESHOLDBELOW TapiNotificationThresholdCrossingType = "THRESHOLD_BELOW"

	// TapiNotificationThresholdCrossingTypeCLEARED captures enum value "CLEARED"
	TapiNotificationThresholdCrossingTypeCLEARED TapiNotificationThresholdCrossingType = "CLEARED"
)

// for schema
var tapiNotificationThresholdCrossingTypeEnum []interface{}

func init() {
	var res []TapiNotificationThresholdCrossingType
	if err := json.Unmarshal([]byte(`["THRESHOLD_ABOVE","THRESHOLD_BELOW","CLEARED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiNotificationThresholdCrossingTypeEnum = append(tapiNotificationThresholdCrossingTypeEnum, v)
	}
}

func (m TapiNotificationThresholdCrossingType) validateTapiNotificationThresholdCrossingTypeEnum(path, location string, value TapiNotificationThresholdCrossingType) error {
	if err := validate.Enum(path, location, value, tapiNotificationThresholdCrossingTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi notification threshold crossing type
func (m TapiNotificationThresholdCrossingType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiNotificationThresholdCrossingTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}