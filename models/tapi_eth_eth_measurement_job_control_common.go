// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthMeasurementJobControlCommon tapi eth eth measurement job control common
// swagger:model tapi.eth.EthMeasurementJobControlCommon
type TapiEthEthMeasurementJobControlCommon struct {

	// Time length over which each Availability Frame Loss Ratio value is calculated.
	//                 MEF 35.1:
	//                 [R78]/[CR58] [O8] A SOAM PM Implementation MUST support a configurable parameter for the length of time over which each Availability flr value is calculated, with a range of 1s – 300s.  This parameter is equivalent to delta-t as specified by MEF 10.3.
	//                 [R79]/[CR59] [O8] The length of time over which each Availability flr value is calculated (delta-t) MUST be an integer multiple of the interval between each SLM/1SL frame transmission.
	//                 [D31]/[CD16] [O8] The default length of time over which each Availability flr value is calculated SHOULD be 1s.
	FlrAvailabilityDeltaTime *int32 `json:"flr-availability-delta-time,omitempty"`

	// Number of consecutive Availability Frame Loss Ratio measurements to be used to determine Available/Unavailable state transitions.
	//                 MEF 35.1:
	//                 [R80]/[CR60] [O8] The number range of 1 through 10 MUST be supported for the configurable number of consecutive Availability flr measurements to be used to determine Available/Unavailable state transitions.
	//                 This parameter is equivalent to the Availability parameter of n as specified by MEF 10.3.
	//                 [D32]/[CD17] [O8] The default number of n for Availability SHOULD be 10.
	FlrAvailabilitySamples *int32 `json:"flr-availability-samples,omitempty"`

	// Frame loss ratio threshold to be used in evaluating the Available/Unavailable state of each time interval (as specified by Availability Delta Time).
	//                 MEF 35.1:
	//                 [R81]/[CR61] A SOAM PM Implementation MUST support a configurable Availability frame loss ratio threshold to be used in evaluating the Available/Unavailable state of each delta-t interval per MEF 10.3
	//                 [R82]/[CR62] The Availability frame loss ratio threshold range of 0.00 through 1.00 MUST be supported in increments of 0.01.
	//                 [D33]/[CD18] [O8] The default Availability frame loss ratio threshold SHOULD be 0.1.
	FlrAvailabilityThreshold *string `json:"flr-availability-threshold,omitempty"`

	// This attribute contains the discrete non overlapping periods of time (in seconds) during which measurements are performed
	//                 (i.e., OAM messages are generated) and reports are gathered at the end of the measurement intervals.
	//                 Note that the value 0 means a degenerated measurement interval with a single OAM message and the report is sent as immediately as possible.
	MeasurementInterval int32 `json:"measurement-interval,omitempty"`

	// This attribute indicates the period (frequency) of the measurement frame transmission.
	//                 Note that the value 0 means that only one OAM message per measurement interval is generated.
	//                 Unit is milliseconds.
	//                 range of type : 100ms, 1s, 10s
	MessagePeriod *int32 `json:"message-period,omitempty"`

	// MEF 35.1:
	//                 [D8] A SOAM PM Implementation SHOULD support a configurable (in minutes) offset from ToD time for alignment of the start of Measurement Intervals other than the first Measurement Interval.
	OffsetFromTimeOfTheDay int32 `json:"offset-from-time-of-the-day,omitempty"`

	// This attribute contains the priority value on which the MEP performs the measurement.
	//                 When the measurement is enabled, the MEP should use this value to encode the priority of generated measurement frames (OAM PDU frames.).
	//                 The EMF usese this value to assign the P parameter of the measurement operation.
	Priority *int32 `json:"priority,omitempty"`

	// This attribute contains the time between the start of two measurement intervals. This IS applicable for the repetitive instance type and MAY be applicable for the repetitive series type.
	//                 Note that a value of 0 means not applicable (NA), which is for the cases of single instance, single series, or repetitive series without extra gap in between the measurement intervals (i.e., also as known as continuous series).
	RepetitionPeriod TapiEthRepetitionPeriod `json:"repetition-period,omitempty"`

	// This attribute is used to distinguish each measurement session if multiple measurement sessions are simultaneously activated towards a peer MEP including concurrent on-demand and proactive tests.
	//                 It must be unique at least within the context of any measurement type for the MEG and initiating MEP.
	//                 Note: The attribute is not used in case of 2-way loss measurement.
	//                 range of type : 0..(2^32) - 1
	TestIdentifier int32 `json:"test-identifier,omitempty"`

	// MEF 35.1:
	//                 [D7] A SOAM PM Implementation SHOULD allow for no alignment to the time-of-day clock.
	TimeOfTheDayAlignment *bool `json:"time-of-the-day-alignment,omitempty"`
}

// Validate validates this tapi eth eth measurement job control common
func (m *TapiEthEthMeasurementJobControlCommon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepetitionPeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthMeasurementJobControlCommon) validateRepetitionPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.RepetitionPeriod) { // not required
		return nil
	}

	if err := m.RepetitionPeriod.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("repetition-period")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthMeasurementJobControlCommon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthMeasurementJobControlCommon) UnmarshalBinary(b []byte) error {
	var res TapiEthEthMeasurementJobControlCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}