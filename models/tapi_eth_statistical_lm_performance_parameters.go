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
	//                     The accuracy of the value is for further study.
	AverageFrameLossRatio string `json:"average-frame-loss-ratio,omitempty"`

	// This attribute contains the maximum frame loss ratio calculated over a period of time.
	//                     The accuracy of the value is for further study.
	MaximumFrameLossRatio string `json:"maximum-frame-loss-ratio,omitempty"`

	// This attribute contains the minimum frame loss ratio calculated over a period of time.
	//                     The accuracy of the value is for further study.
	MinimumFrameLossRatio string `json:"minimum-frame-loss-ratio,omitempty"`

	// This attribute contains the SES detected in the monitoring interval.
	//                     range of type : 0..900 for 15min interval or 0..86400 for 24 hr interval.
	Ses int32 `json:"ses,omitempty"`

	// This attribute contains UAS (unavailable seconds) detected in the monitoring interval.
	//                     range of type : 0..900 for 15min interval or 0..86400 for 24 hr interval.
	Uas int32 `json:"uas,omitempty"`
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