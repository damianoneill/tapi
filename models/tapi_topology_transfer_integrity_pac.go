// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TapiTopologyTransferIntegrityPac tapi topology transfer integrity pac
// swagger:model tapi.topology.TransferIntegrityPac
type TapiTopologyTransferIntegrityPac struct {

	// Describes the degree to which packets will be delivered out of sequence.
	//                 Does not apply to TDM as the TDM protocols maintain strict order.
	DeliveryOrderCharacteristic string `json:"delivery-order-characteristic,omitempty"`

	// Describes the degree to which the signal propagated can be errored.
	//                 Applies to TDM systems as the errored signal will be propagated and not packet as errored packets will be discarded.
	ErrorCharacteristic string `json:"error-characteristic,omitempty"`

	// Describes the acceptable characteristic of lost packets where loss may result from discard due to errors or overflow.
	//                 Applies to packet systems and not TDM (as for TDM errored signals are propagated unless grossly errored and overflow/underflow turns into timing slips).
	LossCharacteristic string `json:"loss-characteristic,omitempty"`

	// Primarily applies to packet systems where a packet may be delivered more than once (in fault recovery for example).
	//                 It can also apply to TDM where several frames may be received twice due to switching in a system with a large differential propagation delay.
	RepeatDeliveryCharacteristic string `json:"repeat-delivery-characteristic,omitempty"`

	// Describes the effect of any server integrity enhancement process on the characteristics of the TopologicalEntity.
	ServerIntegrityProcessCharacteristic string `json:"server-integrity-process-characteristic,omitempty"`

	// Describes the duration for which there may be no valid signal propagated.
	UnavailableTimeCharacteristic string `json:"unavailable-time-characteristic,omitempty"`
}

// Validate validates this tapi topology transfer integrity pac
func (m *TapiTopologyTransferIntegrityPac) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiTopologyTransferIntegrityPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyTransferIntegrityPac) UnmarshalBinary(b []byte) error {
	var res TapiTopologyTransferIntegrityPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
