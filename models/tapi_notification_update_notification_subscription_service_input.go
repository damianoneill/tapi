// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiNotificationUpdateNotificationSubscriptionServiceInput tapi notification update notification subscription service input
// swagger:model tapi.notification.UpdateNotificationSubscriptionServiceInput
type TapiNotificationUpdateNotificationSubscriptionServiceInput struct {

	// input
	Input *TapiNotificationUpdateNotificationSubscriptionServiceInputInput `json:"input,omitempty"`
}

// Validate validates this tapi notification update notification subscription service input
func (m *TapiNotificationUpdateNotificationSubscriptionServiceInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiNotificationUpdateNotificationSubscriptionServiceInput) validateInput(formats strfmt.Registry) error {

	if swag.IsZero(m.Input) { // not required
		return nil
	}

	if m.Input != nil {
		if err := m.Input.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiNotificationUpdateNotificationSubscriptionServiceInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiNotificationUpdateNotificationSubscriptionServiceInput) UnmarshalBinary(b []byte) error {
	var res TapiNotificationUpdateNotificationSubscriptionServiceInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiNotificationUpdateNotificationSubscriptionServiceInputInput tapi notification update notification subscription service input input
// swagger:model TapiNotificationUpdateNotificationSubscriptionServiceInputInput
type TapiNotificationUpdateNotificationSubscriptionServiceInputInput struct {

	// none
	SubscriptionFilter *TapiNotificationSubscriptionFilter `json:"subscription-filter,omitempty"`

	// none
	SubscriptionIDOrName string `json:"subscription-id-or-name,omitempty"`

	// none
	SubscriptionState TapiNotificationSubscriptionState `json:"subscription-state,omitempty"`
}

// Validate validates this tapi notification update notification subscription service input input
func (m *TapiNotificationUpdateNotificationSubscriptionServiceInputInput) Validate(formats strfmt.Registry) error {
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

func (m *TapiNotificationUpdateNotificationSubscriptionServiceInputInput) validateSubscriptionFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionFilter) { // not required
		return nil
	}

	if m.SubscriptionFilter != nil {
		if err := m.SubscriptionFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "subscription-filter")
			}
			return err
		}
	}

	return nil
}

func (m *TapiNotificationUpdateNotificationSubscriptionServiceInputInput) validateSubscriptionState(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionState) { // not required
		return nil
	}

	if err := m.SubscriptionState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("input" + "." + "subscription-state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiNotificationUpdateNotificationSubscriptionServiceInputInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiNotificationUpdateNotificationSubscriptionServiceInputInput) UnmarshalBinary(b []byte) error {
	var res TapiNotificationUpdateNotificationSubscriptionServiceInputInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}