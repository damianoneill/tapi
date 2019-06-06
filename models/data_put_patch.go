// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DataPutPatch data put patch
// swagger:model data-put-patch
type DataPutPatch struct {

	//
	// This module contains TAPI Network Element Model definitions.
	// Source: TapiNetworkElement.uml
	// Copyright (c) 2018 Open Networking Foundation (ONF). All rights reserved.
	// License: This module is distributed under the Apache License 2.0
	IetfRestconfData interface{} `json:"ietf-restconf:data,omitempty"`
}

// Validate validates this data put patch
func (m *DataPutPatch) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataPutPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataPutPatch) UnmarshalBinary(b []byte) error {
	var res DataPutPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
