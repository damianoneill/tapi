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

// DataTapiEquipmentGetPhysicalSpanPostOutput data tapi equipment get physical span post output
// swagger:model data_tapi-equipment_get-physical-span-post-output
type DataTapiEquipmentGetPhysicalSpanPostOutput struct {

	// tapi equipment physical span
	TapiEquipmentPhysicalSpan *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpan `json:"tapi-equipment:physical-span,omitempty"`
}

// Validate validates this data tapi equipment get physical span post output
func (m *DataTapiEquipmentGetPhysicalSpanPostOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTapiEquipmentPhysicalSpan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataTapiEquipmentGetPhysicalSpanPostOutput) validateTapiEquipmentPhysicalSpan(formats strfmt.Registry) error {

	if swag.IsZero(m.TapiEquipmentPhysicalSpan) { // not required
		return nil
	}

	if m.TapiEquipmentPhysicalSpan != nil {
		if err := m.TapiEquipmentPhysicalSpan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tapi-equipment:physical-span")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutput) UnmarshalBinary(b []byte) error {
	var res DataTapiEquipmentGetPhysicalSpanPostOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpan none (non-presence)
// swagger:model DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpan
type DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpan struct {

	// Both the serial segments that form an end-end strand and the parallel end-end strands. (list)
	AbstractStrand []*DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0 `json:"abstract-strand"`

	// none (list)
	AccessPort []*DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAccessPortItems0 `json:"access-port"`

	// List of names. A property of an entity with a value that is unique in some namespace but may change during the life of the entity. A name carries no semantics with respect to the purpose of the entity. (list)
	Name []*DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanNameItems0 `json:"name"`

	// UUID: An identifier that is universally unique within an identifier space, where the identifier space is itself globally unique, and immutable. An UUID carries no semantics with respect to the purpose or state of the entity.
	// UUID here uses string representation as defined in RFC 4122.  The canonical representation uses lowercase characters.
	// Pattern: [0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-' + '[0-9a-fA-F]{4}-[0-9a-fA-F]{12}
	// Example of a UUID in string representation: f81d4fae-7dec-11d0-a765-00a0c91e6bf6 (leaf)
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this data tapi equipment get physical span post output tapi equipment physical span
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbstractStrand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpan) validateAbstractStrand(formats strfmt.Registry) error {

	if swag.IsZero(m.AbstractStrand) { // not required
		return nil
	}

	for i := 0; i < len(m.AbstractStrand); i++ {
		if swag.IsZero(m.AbstractStrand[i]) { // not required
			continue
		}

		if m.AbstractStrand[i] != nil {
			if err := m.AbstractStrand[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tapi-equipment:physical-span" + "." + "abstract-strand" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpan) validateAccessPort(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessPort) { // not required
		return nil
	}

	for i := 0; i < len(m.AccessPort); i++ {
		if swag.IsZero(m.AccessPort[i]) { // not required
			continue
		}

		if m.AccessPort[i] != nil {
			if err := m.AccessPort[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tapi-equipment:physical-span" + "." + "access-port" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpan) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	for i := 0; i < len(m.Name); i++ {
		if swag.IsZero(m.Name[i]) { // not required
			continue
		}

		if m.Name[i] != nil {
			if err := m.Name[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tapi-equipment:physical-span" + "." + "name" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpan) UnmarshalBinary(b []byte) error {
	var res DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0 data tapi equipment get physical span post output tapi equipment physical span abstract strand items0
// swagger:model DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0
type DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0 struct {

	// none (list)
	AdjacentStrand []*DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0AdjacentStrandItems0 `json:"adjacent-strand"`

	// A strand can end on two or more Pins (usually 2 pins, but a strand my be spliced to split a signal). This model supports only 2 ended strands.
	// A abstract strand may be spliced at both ends and hence have no direct relationship to pins or may be connected to pins at one or both ends.
	// In the essential model these Pins would be on connectors that plug in to connectors on Equipments.
	// The AbstractStrand is extended to the pins of the AccessPort which are the Pins on the Connectors of the Equipment.
	// In some cases it may not be relevant to represent the pin detail and hence the reference is to a connector alone. (list)
	ConnectorPin []*DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0ConnectorPinItems0 `json:"connector-pin"`

	// none (leaf)
	LocalID string `json:"local-id,omitempty"`

	// List of names. A property of an entity with a value that is unique in some namespace but may change during the life of the entity. A name carries no semantics with respect to the purpose of the entity. (list)
	Name []*DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0NameItems0 `json:"name"`

	// none (list)
	SplicedStrand []*DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0SplicedStrandItems0 `json:"spliced-strand"`

	// Relevant physical properties of the abstract strand. (list)
	StrandMediaCharacteristics []*DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0StrandMediaCharacteristicsItems0 `json:"strand-media-characteristics"`
}

// Validate validates this data tapi equipment get physical span post output tapi equipment physical span abstract strand items0
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdjacentStrand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectorPin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSplicedStrand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrandMediaCharacteristics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0) validateAdjacentStrand(formats strfmt.Registry) error {

	if swag.IsZero(m.AdjacentStrand) { // not required
		return nil
	}

	for i := 0; i < len(m.AdjacentStrand); i++ {
		if swag.IsZero(m.AdjacentStrand[i]) { // not required
			continue
		}

		if m.AdjacentStrand[i] != nil {
			if err := m.AdjacentStrand[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adjacent-strand" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0) validateConnectorPin(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectorPin) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectorPin); i++ {
		if swag.IsZero(m.ConnectorPin[i]) { // not required
			continue
		}

		if m.ConnectorPin[i] != nil {
			if err := m.ConnectorPin[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connector-pin" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	for i := 0; i < len(m.Name); i++ {
		if swag.IsZero(m.Name[i]) { // not required
			continue
		}

		if m.Name[i] != nil {
			if err := m.Name[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("name" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0) validateSplicedStrand(formats strfmt.Registry) error {

	if swag.IsZero(m.SplicedStrand) { // not required
		return nil
	}

	for i := 0; i < len(m.SplicedStrand); i++ {
		if swag.IsZero(m.SplicedStrand[i]) { // not required
			continue
		}

		if m.SplicedStrand[i] != nil {
			if err := m.SplicedStrand[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spliced-strand" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0) validateStrandMediaCharacteristics(formats strfmt.Registry) error {

	if swag.IsZero(m.StrandMediaCharacteristics) { // not required
		return nil
	}

	for i := 0; i < len(m.StrandMediaCharacteristics); i++ {
		if swag.IsZero(m.StrandMediaCharacteristics[i]) { // not required
			continue
		}

		if m.StrandMediaCharacteristics[i] != nil {
			if err := m.StrandMediaCharacteristics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("strand-media-characteristics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0) UnmarshalBinary(b []byte) error {
	var res DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0AdjacentStrandItems0 data tapi equipment get physical span post output tapi equipment physical span abstract strand items0 adjacent strand items0
// swagger:model DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0AdjacentStrandItems0
type DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0AdjacentStrandItems0 struct {

	// none (leaf)
	AbstractStrandLocalID string `json:"abstract-strand-local-id,omitempty"`

	// none (leaf)
	PhysicalSpanUUID string `json:"physical-span-uuid,omitempty"`
}

// Validate validates this data tapi equipment get physical span post output tapi equipment physical span abstract strand items0 adjacent strand items0
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0AdjacentStrandItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0AdjacentStrandItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0AdjacentStrandItems0) UnmarshalBinary(b []byte) error {
	var res DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0AdjacentStrandItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0ConnectorPinItems0 data tapi equipment get physical span post output tapi equipment physical span abstract strand items0 connector pin items0
// swagger:model DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0ConnectorPinItems0
type DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0ConnectorPinItems0 struct {

	// Identification of the Connector in the conetxt of the referenced Equipment. (leaf)
	ConnectorIdentification string `json:"connector-identification,omitempty"`

	// Reference to the Equipment that is fitted with the Connector/Pin. (leaf)
	EquipmentUUID string `json:"equipment-uuid,omitempty"`

	// Where relevant, identification of the Pin in the contect of the connector.
	// Where the whole connector is used, then individual Pins need not be identified. (leaf)
	PinIdentification string `json:"pin-identification,omitempty"`
}

// Validate validates this data tapi equipment get physical span post output tapi equipment physical span abstract strand items0 connector pin items0
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0ConnectorPinItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0ConnectorPinItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0ConnectorPinItems0) UnmarshalBinary(b []byte) error {
	var res DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0ConnectorPinItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0NameItems0 data tapi equipment get physical span post output tapi equipment physical span abstract strand items0 name items0
// swagger:model DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0NameItems0
type DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0NameItems0 struct {

	// The value (leaf)
	Value string `json:"value,omitempty"`

	// The name of the value. The value need not have a name. (leaf)
	ValueName string `json:"value-name,omitempty"`
}

// Validate validates this data tapi equipment get physical span post output tapi equipment physical span abstract strand items0 name items0
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0NameItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0NameItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0NameItems0) UnmarshalBinary(b []byte) error {
	var res DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0NameItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0SplicedStrandItems0 data tapi equipment get physical span post output tapi equipment physical span abstract strand items0 spliced strand items0
// swagger:model DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0SplicedStrandItems0
type DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0SplicedStrandItems0 struct {

	// none (leaf)
	AbstractStrandLocalID string `json:"abstract-strand-local-id,omitempty"`

	// none (leaf)
	PhysicalSpanUUID string `json:"physical-span-uuid,omitempty"`
}

// Validate validates this data tapi equipment get physical span post output tapi equipment physical span abstract strand items0 spliced strand items0
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0SplicedStrandItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0SplicedStrandItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0SplicedStrandItems0) UnmarshalBinary(b []byte) error {
	var res DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0SplicedStrandItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0StrandMediaCharacteristicsItems0 data tapi equipment get physical span post output tapi equipment physical span abstract strand items0 strand media characteristics items0
// swagger:model DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0StrandMediaCharacteristicsItems0
type DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0StrandMediaCharacteristicsItems0 struct {

	// The value (leaf)
	Value string `json:"value,omitempty"`

	// The name of the value. The value need not have a name. (leaf)
	ValueName string `json:"value-name,omitempty"`
}

// Validate validates this data tapi equipment get physical span post output tapi equipment physical span abstract strand items0 strand media characteristics items0
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0StrandMediaCharacteristicsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0StrandMediaCharacteristicsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0StrandMediaCharacteristicsItems0) UnmarshalBinary(b []byte) error {
	var res DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAbstractStrandItems0StrandMediaCharacteristicsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAccessPortItems0 data tapi equipment get physical span post output tapi equipment physical span access port items0
// swagger:model DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAccessPortItems0
type DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAccessPortItems0 struct {

	// none (leaf)
	AccessPortUUID string `json:"access-port-uuid,omitempty"`

	// none (leaf)
	DeviceUUID string `json:"device-uuid,omitempty"`
}

// Validate validates this data tapi equipment get physical span post output tapi equipment physical span access port items0
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAccessPortItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAccessPortItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAccessPortItems0) UnmarshalBinary(b []byte) error {
	var res DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanAccessPortItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanNameItems0 data tapi equipment get physical span post output tapi equipment physical span name items0
// swagger:model DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanNameItems0
type DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanNameItems0 struct {

	// The value (leaf)
	Value string `json:"value,omitempty"`

	// The name of the value. The value need not have a name. (leaf)
	ValueName string `json:"value-name,omitempty"`
}

// Validate validates this data tapi equipment get physical span post output tapi equipment physical span name items0
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanNameItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanNameItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanNameItems0) UnmarshalBinary(b []byte) error {
	var res DataTapiEquipmentGetPhysicalSpanPostOutputTapiEquipmentPhysicalSpanNameItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}