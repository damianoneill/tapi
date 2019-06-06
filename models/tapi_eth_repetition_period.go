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

// TapiEthRepetitionPeriod tapi eth repetition period
// swagger:model tapi.eth.RepetitionPeriod
type TapiEthRepetitionPeriod string

const (

	// TapiEthRepetitionPeriodNr1MIN captures enum value "1MIN"
	TapiEthRepetitionPeriodNr1MIN TapiEthRepetitionPeriod = "1MIN"

	// TapiEthRepetitionPeriodNr1S captures enum value "1S"
	TapiEthRepetitionPeriodNr1S TapiEthRepetitionPeriod = "1S"

	// TapiEthRepetitionPeriodNr10S captures enum value "10S"
	TapiEthRepetitionPeriodNr10S TapiEthRepetitionPeriod = "10S"

	// TapiEthRepetitionPeriodNr0 captures enum value "0"
	TapiEthRepetitionPeriodNr0 TapiEthRepetitionPeriod = "0"
)

// for schema
var tapiEthRepetitionPeriodEnum []interface{}

func init() {
	var res []TapiEthRepetitionPeriod
	if err := json.Unmarshal([]byte(`["1MIN","1S","10S","0"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiEthRepetitionPeriodEnum = append(tapiEthRepetitionPeriodEnum, v)
	}
}

func (m TapiEthRepetitionPeriod) validateTapiEthRepetitionPeriodEnum(path, location string, value TapiEthRepetitionPeriod) error {
	if err := validate.Enum(path, location, value, tapiEthRepetitionPeriodEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi eth repetition period
func (m TapiEthRepetitionPeriod) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiEthRepetitionPeriodEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
