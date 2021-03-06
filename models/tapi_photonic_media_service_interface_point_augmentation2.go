// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiPhotonicMediaServiceInterfacePointAugmentation2 tapi photonic media service interface point augmentation2
// swagger:model tapi.photonic.media.ServiceInterfacePointAugmentation2
type TapiPhotonicMediaServiceInterfacePointAugmentation2 struct {

	// none
	MediaChannelServiceInterfacePointSpec *TapiPhotonicMediaMediaChannelServiceInterfacePointSpec `json:"media-channel-service-interface-point-spec,omitempty"`
}

// Validate validates this tapi photonic media service interface point augmentation2
func (m *TapiPhotonicMediaServiceInterfacePointAugmentation2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMediaChannelServiceInterfacePointSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiPhotonicMediaServiceInterfacePointAugmentation2) validateMediaChannelServiceInterfacePointSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaChannelServiceInterfacePointSpec) { // not required
		return nil
	}

	if m.MediaChannelServiceInterfacePointSpec != nil {
		if err := m.MediaChannelServiceInterfacePointSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("media-channel-service-interface-point-spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiPhotonicMediaServiceInterfacePointAugmentation2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiPhotonicMediaServiceInterfacePointAugmentation2) UnmarshalBinary(b []byte) error {
	var res TapiPhotonicMediaServiceInterfacePointAugmentation2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
