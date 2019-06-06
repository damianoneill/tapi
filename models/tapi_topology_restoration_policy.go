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

// TapiTopologyRestorationPolicy tapi topology restoration policy
// swagger:model tapi.topology.RestorationPolicy
type TapiTopologyRestorationPolicy string

const (

	// TapiTopologyRestorationPolicyPERDOMAINRESTORATION captures enum value "PER_DOMAIN_RESTORATION"
	TapiTopologyRestorationPolicyPERDOMAINRESTORATION TapiTopologyRestorationPolicy = "PER_DOMAIN_RESTORATION"

	// TapiTopologyRestorationPolicyENDTOENDRESTORATION captures enum value "END_TO_END_RESTORATION"
	TapiTopologyRestorationPolicyENDTOENDRESTORATION TapiTopologyRestorationPolicy = "END_TO_END_RESTORATION"

	// TapiTopologyRestorationPolicyNA captures enum value "NA"
	TapiTopologyRestorationPolicyNA TapiTopologyRestorationPolicy = "NA"
)

// for schema
var tapiTopologyRestorationPolicyEnum []interface{}

func init() {
	var res []TapiTopologyRestorationPolicy
	if err := json.Unmarshal([]byte(`["PER_DOMAIN_RESTORATION","END_TO_END_RESTORATION","NA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiTopologyRestorationPolicyEnum = append(tapiTopologyRestorationPolicyEnum, v)
	}
}

func (m TapiTopologyRestorationPolicy) validateTapiTopologyRestorationPolicyEnum(path, location string, value TapiTopologyRestorationPolicy) error {
	if err := validate.Enum(path, location, value, tapiTopologyRestorationPolicyEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi topology restoration policy
func (m TapiTopologyRestorationPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiTopologyRestorationPolicyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}