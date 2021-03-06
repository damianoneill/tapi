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

// TapiOamServiceAffecting tapi oam service affecting
// swagger:model tapi.oam.ServiceAffecting
type TapiOamServiceAffecting string

const (

	// TapiOamServiceAffectingSERVICEAFFECTING captures enum value "SERVICE_AFFECTING"
	TapiOamServiceAffectingSERVICEAFFECTING TapiOamServiceAffecting = "SERVICE_AFFECTING"

	// TapiOamServiceAffectingNOTSERVICEAFFECTING captures enum value "NOT_SERVICE_AFFECTING"
	TapiOamServiceAffectingNOTSERVICEAFFECTING TapiOamServiceAffecting = "NOT_SERVICE_AFFECTING"

	// TapiOamServiceAffectingUNKNOWN captures enum value "UNKNOWN"
	TapiOamServiceAffectingUNKNOWN TapiOamServiceAffecting = "UNKNOWN"
)

// for schema
var tapiOamServiceAffectingEnum []interface{}

func init() {
	var res []TapiOamServiceAffecting
	if err := json.Unmarshal([]byte(`["SERVICE_AFFECTING","NOT_SERVICE_AFFECTING","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiOamServiceAffectingEnum = append(tapiOamServiceAffectingEnum, v)
	}
}

func (m TapiOamServiceAffecting) validateTapiOamServiceAffectingEnum(path, location string, value TapiOamServiceAffecting) error {
	if err := validate.Enum(path, location, value, tapiOamServiceAffectingEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi oam service affecting
func (m TapiOamServiceAffecting) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiOamServiceAffectingEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
