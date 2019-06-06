// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiPhotonicMediaMediaChannelConnectivityServiceEndPointSpec tapi photonic media media channel connectivity service end point spec
// swagger:model tapi.photonic.media.MediaChannelConnectivityServiceEndPointSpec
type TapiPhotonicMediaMediaChannelConnectivityServiceEndPointSpec struct {

	// none
	McConfig *TapiPhotonicMediaMediaChannelConfigPac `json:"mc-config,omitempty"`

	// This parameters shall be used to configure the expected
	//                 and intended (desired) power levels at the endpoints of the media
	//                 Channel connectivity service. These parameters are dependent of the
	//                 related OTSi power-management capabilities exposed at the SIPs
	PowerManagementConfig *TapiPhotonicMediaPowerManagementConfigPac `json:"power-management-config,omitempty"`
}

// Validate validates this tapi photonic media media channel connectivity service end point spec
func (m *TapiPhotonicMediaMediaChannelConnectivityServiceEndPointSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMcConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerManagementConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiPhotonicMediaMediaChannelConnectivityServiceEndPointSpec) validateMcConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.McConfig) { // not required
		return nil
	}

	if m.McConfig != nil {
		if err := m.McConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mc-config")
			}
			return err
		}
	}

	return nil
}

func (m *TapiPhotonicMediaMediaChannelConnectivityServiceEndPointSpec) validatePowerManagementConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerManagementConfig) { // not required
		return nil
	}

	if m.PowerManagementConfig != nil {
		if err := m.PowerManagementConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power-management-config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiPhotonicMediaMediaChannelConnectivityServiceEndPointSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPhotonicMediaMediaChannelConnectivityServiceEndPointSpec) UnmarshalBinary(b []byte) error {
	var res TapiPhotonicMediaMediaChannelConnectivityServiceEndPointSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}