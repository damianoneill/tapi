// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOamOamJob tapi oam oam job
// swagger:model tapi.oam.OamJob
type TapiOamOamJob struct {
	TapiCommonAdminStatePac

	TapiCommonGlobalClass

	// none
	CreationTime string `json:"creation-time,omitempty"`

	// Granularity period of the CurrentData identifies the specific CurrentData instance in the scope of this OamJob.
	//                 For example, typically at least
	//                 one 15min and
	//                 one 24hr;
	//                 optionally one additional configurable (< 15min)
	CurrentData []*TapiOamCurrentData `json:"current-data"`

	// none
	OamJobType string `json:"oam-job-type,omitempty"`

	// none
	OamProfile *TapiOamOamProfileRef `json:"oam-profile,omitempty"`

	// none
	OamServicePoint []*TapiOamOamServicePointRef `json:"oam-service-point"`

	// none
	Schedule *TapiCommonTimeRange `json:"schedule,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiOamOamJob) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonAdminStatePac
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonAdminStatePac = aO0

	// AO1
	var aO1 TapiCommonGlobalClass
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiCommonGlobalClass = aO1

	// AO2
	var dataAO2 struct {
		CreationTime string `json:"creation-time,omitempty"`

		CurrentData []*TapiOamCurrentData `json:"current-data"`

		OamJobType string `json:"oam-job-type,omitempty"`

		OamProfile *TapiOamOamProfileRef `json:"oam-profile,omitempty"`

		OamServicePoint []*TapiOamOamServicePointRef `json:"oam-service-point"`

		Schedule *TapiCommonTimeRange `json:"schedule,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.CreationTime = dataAO2.CreationTime

	m.CurrentData = dataAO2.CurrentData

	m.OamJobType = dataAO2.OamJobType

	m.OamProfile = dataAO2.OamProfile

	m.OamServicePoint = dataAO2.OamServicePoint

	m.Schedule = dataAO2.Schedule

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiOamOamJob) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.TapiCommonAdminStatePac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiCommonGlobalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	var dataAO2 struct {
		CreationTime string `json:"creation-time,omitempty"`

		CurrentData []*TapiOamCurrentData `json:"current-data"`

		OamJobType string `json:"oam-job-type,omitempty"`

		OamProfile *TapiOamOamProfileRef `json:"oam-profile,omitempty"`

		OamServicePoint []*TapiOamOamServicePointRef `json:"oam-service-point"`

		Schedule *TapiCommonTimeRange `json:"schedule,omitempty"`
	}

	dataAO2.CreationTime = m.CreationTime

	dataAO2.CurrentData = m.CurrentData

	dataAO2.OamJobType = m.OamJobType

	dataAO2.OamProfile = m.OamProfile

	dataAO2.OamServicePoint = m.OamServicePoint

	dataAO2.Schedule = m.Schedule

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi oam oam job
func (m *TapiOamOamJob) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonAdminStatePac
	if err := m.TapiCommonAdminStatePac.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOamProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOamServicePoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamOamJob) validateCurrentData(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentData) { // not required
		return nil
	}

	for i := 0; i < len(m.CurrentData); i++ {
		if swag.IsZero(m.CurrentData[i]) { // not required
			continue
		}

		if m.CurrentData[i] != nil {
			if err := m.CurrentData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("current-data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiOamOamJob) validateOamProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.OamProfile) { // not required
		return nil
	}

	if m.OamProfile != nil {
		if err := m.OamProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oam-profile")
			}
			return err
		}
	}

	return nil
}

func (m *TapiOamOamJob) validateOamServicePoint(formats strfmt.Registry) error {

	if swag.IsZero(m.OamServicePoint) { // not required
		return nil
	}

	for i := 0; i < len(m.OamServicePoint); i++ {
		if swag.IsZero(m.OamServicePoint[i]) { // not required
			continue
		}

		if m.OamServicePoint[i] != nil {
			if err := m.OamServicePoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("oam-service-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiOamOamJob) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamOamJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamOamJob) UnmarshalBinary(b []byte) error {
	var res TapiOamOamJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
