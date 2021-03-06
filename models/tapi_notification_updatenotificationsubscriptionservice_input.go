// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiNotificationUpdatenotificationsubscriptionserviceInput tapi notification updatenotificationsubscriptionservice input
// swagger:model tapi.notification.updatenotificationsubscriptionservice.Input
type TapiNotificationUpdatenotificationsubscriptionserviceInput struct {

	// none
	SubscriptionFilter *TapiNotificationSubscriptionFilter `json:"subscription-filter,omitempty"`

	// none
	SubscriptionIDOrName string `json:"subscription-id-or-name,omitempty"`

	// none
	SubscriptionState TapiNotificationSubscriptionState `json:"subscription-state,omitempty"`
}

// Validate validates this tapi notification updatenotificationsubscriptionservice input
func (m *TapiNotificationUpdatenotificationsubscriptionserviceInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscriptionFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiNotificationUpdatenotificationsubscriptionserviceInput) validateSubscriptionFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionFilter) { // not required
		return nil
	}

	if m.SubscriptionFilter != nil {
		if err := m.SubscriptionFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscription-filter")
			}
			return err
		}
	}

	return nil
}

func (m *TapiNotificationUpdatenotificationsubscriptionserviceInput) validateSubscriptionState(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionState) { // not required
		return nil
	}

	if err := m.SubscriptionState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subscription-state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiNotificationUpdatenotificationsubscriptionserviceInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiNotificationUpdatenotificationsubscriptionserviceInput) UnmarshalBinary(b []byte) error {
	var res TapiNotificationUpdatenotificationsubscriptionserviceInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
