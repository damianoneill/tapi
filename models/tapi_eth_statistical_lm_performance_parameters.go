// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiEthStatisticalLmPerformanceParameters tapi eth statistical lm performance parameters
// swagger:model tapi.eth.StatisticalLmPerformanceParameters
type TapiEthStatisticalLmPerformanceParameters struct {

	// This attribute contains the average frame loss ratio calculated over a period of time.
	AverageFrameLossRatio string `json:"average-frame-loss-ratio,omitempty"`

	// A generalized SES.
	//                 MEF 10.3: The Resiliency attributes are similar to the definitions of Severely Errored Seconds (SES) and Consecutive SES in section 9 and Annex B (respectively) of Y.1563 [6], when delta-t = 1 second.
	//                 MEF 35.1: Count of High Loss Intervals during the Measurement Interval.
	//                 range of type : 0..900 for 15min interval or 0..86400 for 24 hr interval.
	HliCount int32 `json:"hli-count,omitempty"`

	// This attribute contains the maximum frame loss ratio calculated over a period of time.
	MaximumFrameLossRatio string `json:"maximum-frame-loss-ratio,omitempty"`

	// This attribute contains the minimum frame loss ratio calculated over a period of time.
	MinimumFrameLossRatio string `json:"minimum-frame-loss-ratio,omitempty"`

	// A generalized UAS.
	//                 MEF 35.1: A 32-bit counter reflecting the number of delta-t intervals evaluated as Unavailable (i.e., for which A<Controller, Responder>(delta-t) = 0).
	//                 range of type : 0..900 for 15min interval or 0..86400 for 24 hr interval.
	UnavailableIntervals int32 `json:"unavailable-intervals,omitempty"`
}

// Validate validates this tapi eth statistical lm performance parameters
func (m *TapiEthStatisticalLmPerformanceParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthStatisticalLmPerformanceParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthStatisticalLmPerformanceParameters) UnmarshalBinary(b []byte) error {
	var res TapiEthStatisticalLmPerformanceParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
