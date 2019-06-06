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

// OperationsTapiEquipmentGetPhysicalSpanListPostOutput operations tapi equipment get physical span list post output
// swagger:model operations_tapi-equipment_get-physical-span-list-post-output
type OperationsTapiEquipmentGetPhysicalSpanListPostOutput struct {

	// none (list)
	TapiEquipmentPhysicalSpan []*OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0 `json:"tapi-equipment:physical-span"`
}

// Validate validates this operations tapi equipment get physical span list post output
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTapiEquipmentPhysicalSpan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutput) validateTapiEquipmentPhysicalSpan(formats strfmt.Registry) error {

	if swag.IsZero(m.TapiEquipmentPhysicalSpan) { // not required
		return nil
	}

	for i := 0; i < len(m.TapiEquipmentPhysicalSpan); i++ {
		if swag.IsZero(m.TapiEquipmentPhysicalSpan[i]) { // not required
			continue
		}

		if m.TapiEquipmentPhysicalSpan[i] != nil {
			if err := m.TapiEquipmentPhysicalSpan[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tapi-equipment:physical-span" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutput) UnmarshalBinary(b []byte) error {
	var res OperationsTapiEquipmentGetPhysicalSpanListPostOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0 operations tapi equipment get physical span list post output tapi equipment physical span items0
// swagger:model OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0
type OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0 struct {

	// Both the serial segments that form an end-end strand and the parallel end-end strands. (list)
	AbstractStrand []*OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0 `json:"abstract-strand"`

	// none (list)
	AccessPort []*OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AccessPortItems0 `json:"access-port"`

	// List of names. A property of an entity with a value that is unique in some namespace but may change during the life of the entity. A name carries no semantics with respect to the purpose of the entity. (list)
	Name []*OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0NameItems0 `json:"name"`

	// UUID: An identifier that is universally unique within an identifier space, where the identifier space is itself globally unique, and immutable. An UUID carries no semantics with respect to the purpose or state of the entity.
	// UUID here uses string representation as defined in RFC 4122.  The canonical representation uses lowercase characters.
	// Pattern: [0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-' + '[0-9a-fA-F]{4}-[0-9a-fA-F]{12}
	// Example of a UUID in string representation: f81d4fae-7dec-11d0-a765-00a0c91e6bf6 (leaf)
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this operations tapi equipment get physical span list post output tapi equipment physical span items0
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0) Validate(formats strfmt.Registry) error {
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

func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0) validateAbstractStrand(formats strfmt.Registry) error {

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
					return ve.ValidateName("abstract-strand" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0) validateAccessPort(formats strfmt.Registry) error {

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
					return ve.ValidateName("access-port" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0) validateName(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0) UnmarshalBinary(b []byte) error {
	var res OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0 operations tapi equipment get physical span list post output tapi equipment physical span items0 abstract strand items0
// swagger:model OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0
type OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0 struct {

	// none (list)
	AdjacentStrand []*OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0AdjacentStrandItems0 `json:"adjacent-strand"`

	// A strand can end on two or more Pins (usually 2 pins, but a strand my be spliced to split a signal). This model supports only 2 ended strands.
	// A abstract strand may be spliced at both ends and hence have no direct relationship to pins or may be connected to pins at one or both ends.
	// In the essential model these Pins would be on connectors that plug in to connectors on Equipments.
	// The AbstractStrand is extended to the pins of the AccessPort which are the Pins on the Connectors of the Equipment.
	// In some cases it may not be relevant to represent the pin detail and hence the reference is to a connector alone. (list)
	ConnectorPin []*OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0ConnectorPinItems0 `json:"connector-pin"`

	// none (leaf)
	LocalID string `json:"local-id,omitempty"`

	// List of names. A property of an entity with a value that is unique in some namespace but may change during the life of the entity. A name carries no semantics with respect to the purpose of the entity. (list)
	Name []*OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0NameItems0 `json:"name"`

	// none (list)
	SplicedStrand []*OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0SplicedStrandItems0 `json:"spliced-strand"`

	// Relevant physical properties of the abstract strand. (list)
	StrandMediaCharacteristics []*OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0StrandMediaCharacteristicsItems0 `json:"strand-media-characteristics"`
}

// Validate validates this operations tapi equipment get physical span list post output tapi equipment physical span items0 abstract strand items0
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0) Validate(formats strfmt.Registry) error {
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

func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0) validateAdjacentStrand(formats strfmt.Registry) error {

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

func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0) validateConnectorPin(formats strfmt.Registry) error {

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

func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0) validateName(formats strfmt.Registry) error {

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

func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0) validateSplicedStrand(formats strfmt.Registry) error {

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

func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0) validateStrandMediaCharacteristics(formats strfmt.Registry) error {

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
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0) UnmarshalBinary(b []byte) error {
	var res OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0AdjacentStrandItems0 operations tapi equipment get physical span list post output tapi equipment physical span items0 abstract strand items0 adjacent strand items0
// swagger:model OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0AdjacentStrandItems0
type OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0AdjacentStrandItems0 struct {

	// none (leaf)
	AbstractStrandLocalID string `json:"abstract-strand-local-id,omitempty"`

	// none (leaf)
	PhysicalSpanUUID string `json:"physical-span-uuid,omitempty"`
}

// Validate validates this operations tapi equipment get physical span list post output tapi equipment physical span items0 abstract strand items0 adjacent strand items0
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0AdjacentStrandItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0AdjacentStrandItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0AdjacentStrandItems0) UnmarshalBinary(b []byte) error {
	var res OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0AdjacentStrandItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0ConnectorPinItems0 operations tapi equipment get physical span list post output tapi equipment physical span items0 abstract strand items0 connector pin items0
// swagger:model OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0ConnectorPinItems0
type OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0ConnectorPinItems0 struct {

	// Identification of the Connector in the conetxt of the referenced Equipment. (leaf)
	ConnectorIdentification string `json:"connector-identification,omitempty"`

	// Reference to the Equipment that is fitted with the Connector/Pin. (leaf)
	EquipmentUUID string `json:"equipment-uuid,omitempty"`

	// Where relevant, identification of the Pin in the contect of the connector.
	// Where the whole connector is used, then individual Pins need not be identified. (leaf)
	PinIdentification string `json:"pin-identification,omitempty"`
}

// Validate validates this operations tapi equipment get physical span list post output tapi equipment physical span items0 abstract strand items0 connector pin items0
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0ConnectorPinItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0ConnectorPinItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0ConnectorPinItems0) UnmarshalBinary(b []byte) error {
	var res OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0ConnectorPinItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0NameItems0 operations tapi equipment get physical span list post output tapi equipment physical span items0 abstract strand items0 name items0
// swagger:model OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0NameItems0
type OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0NameItems0 struct {

	// The value (leaf)
	Value string `json:"value,omitempty"`

	// The name of the value. The value need not have a name. (leaf)
	ValueName string `json:"value-name,omitempty"`
}

// Validate validates this operations tapi equipment get physical span list post output tapi equipment physical span items0 abstract strand items0 name items0
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0NameItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0NameItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0NameItems0) UnmarshalBinary(b []byte) error {
	var res OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0NameItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0SplicedStrandItems0 operations tapi equipment get physical span list post output tapi equipment physical span items0 abstract strand items0 spliced strand items0
// swagger:model OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0SplicedStrandItems0
type OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0SplicedStrandItems0 struct {

	// none (leaf)
	AbstractStrandLocalID string `json:"abstract-strand-local-id,omitempty"`

	// none (leaf)
	PhysicalSpanUUID string `json:"physical-span-uuid,omitempty"`
}

// Validate validates this operations tapi equipment get physical span list post output tapi equipment physical span items0 abstract strand items0 spliced strand items0
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0SplicedStrandItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0SplicedStrandItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0SplicedStrandItems0) UnmarshalBinary(b []byte) error {
	var res OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0SplicedStrandItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0StrandMediaCharacteristicsItems0 operations tapi equipment get physical span list post output tapi equipment physical span items0 abstract strand items0 strand media characteristics items0
// swagger:model OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0StrandMediaCharacteristicsItems0
type OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0StrandMediaCharacteristicsItems0 struct {

	// The value (leaf)
	Value string `json:"value,omitempty"`

	// The name of the value. The value need not have a name. (leaf)
	ValueName string `json:"value-name,omitempty"`
}

// Validate validates this operations tapi equipment get physical span list post output tapi equipment physical span items0 abstract strand items0 strand media characteristics items0
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0StrandMediaCharacteristicsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0StrandMediaCharacteristicsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0StrandMediaCharacteristicsItems0) UnmarshalBinary(b []byte) error {
	var res OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AbstractStrandItems0StrandMediaCharacteristicsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AccessPortItems0 operations tapi equipment get physical span list post output tapi equipment physical span items0 access port items0
// swagger:model OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AccessPortItems0
type OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AccessPortItems0 struct {

	// none (leaf)
	AccessPortUUID string `json:"access-port-uuid,omitempty"`

	// none (leaf)
	DeviceUUID string `json:"device-uuid,omitempty"`
}

// Validate validates this operations tapi equipment get physical span list post output tapi equipment physical span items0 access port items0
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AccessPortItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AccessPortItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AccessPortItems0) UnmarshalBinary(b []byte) error {
	var res OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0AccessPortItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0NameItems0 operations tapi equipment get physical span list post output tapi equipment physical span items0 name items0
// swagger:model OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0NameItems0
type OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0NameItems0 struct {

	// The value (leaf)
	Value string `json:"value,omitempty"`

	// The name of the value. The value need not have a name. (leaf)
	ValueName string `json:"value-name,omitempty"`
}

// Validate validates this operations tapi equipment get physical span list post output tapi equipment physical span items0 name items0
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0NameItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0NameItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0NameItems0) UnmarshalBinary(b []byte) error {
	var res OperationsTapiEquipmentGetPhysicalSpanListPostOutputTapiEquipmentPhysicalSpanItems0NameItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}