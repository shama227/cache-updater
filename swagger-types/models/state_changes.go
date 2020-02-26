// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StateChanges state changes
// swagger:model StateChanges
type StateChanges struct {

	// height
	// Required: true
	Height *int32 `json:"height"`

	// id
	// Required: true
	ID *string `json:"id"`

	// state changes
	// Required: true
	StateChanges *StateChangesStateChanges `json:"stateChanges"`
}

// Validate validates this state changes
func (m *StateChanges) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateChanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateChanges) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

func (m *StateChanges) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *StateChanges) validateStateChanges(formats strfmt.Registry) error {

	if err := validate.Required("stateChanges", "body", m.StateChanges); err != nil {
		return err
	}

	if m.StateChanges != nil {
		if err := m.StateChanges.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stateChanges")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StateChanges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateChanges) UnmarshalBinary(b []byte) error {
	var res StateChanges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StateChangesStateChanges state changes state changes
// swagger:model StateChangesStateChanges
type StateChangesStateChanges struct {

	// data
	// Required: true
	Data []*StateChangesStateChangesDataItems0 `json:"data"`

	// transfers
	// Required: true
	Transfers []*StateChangesStateChangesTransfersItems0 `json:"transfers"`
}

// Validate validates this state changes state changes
func (m *StateChangesStateChanges) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransfers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateChangesStateChanges) validateData(formats strfmt.Registry) error {

	if err := validate.Required("stateChanges"+"."+"data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stateChanges" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StateChangesStateChanges) validateTransfers(formats strfmt.Registry) error {

	if err := validate.Required("stateChanges"+"."+"transfers", "body", m.Transfers); err != nil {
		return err
	}

	for i := 0; i < len(m.Transfers); i++ {
		if swag.IsZero(m.Transfers[i]) { // not required
			continue
		}

		if m.Transfers[i] != nil {
			if err := m.Transfers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stateChanges" + "." + "transfers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StateChangesStateChanges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateChangesStateChanges) UnmarshalBinary(b []byte) error {
	var res StateChangesStateChanges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StateChangesStateChangesDataItems0 state changes state changes data items0
// swagger:model StateChangesStateChangesDataItems0
type StateChangesStateChangesDataItems0 struct {

	// key
	// Required: true
	Key *string `json:"key"`

	// integer
	// Required: true
	// Enum: [integer boolean binary string]
	Type *string `json:"type"`

	// value
	// Required: true
	Value interface{} `json:"value"`
}

// Validate validates this state changes state changes data items0
func (m *StateChangesStateChangesDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateChangesStateChangesDataItems0) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

var stateChangesStateChangesDataItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["integer","boolean","binary","string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stateChangesStateChangesDataItems0TypeTypePropEnum = append(stateChangesStateChangesDataItems0TypeTypePropEnum, v)
	}
}

const (

	// StateChangesStateChangesDataItems0TypeInteger captures enum value "integer"
	StateChangesStateChangesDataItems0TypeInteger string = "integer"

	// StateChangesStateChangesDataItems0TypeBoolean captures enum value "boolean"
	StateChangesStateChangesDataItems0TypeBoolean string = "boolean"

	// StateChangesStateChangesDataItems0TypeBinary captures enum value "binary"
	StateChangesStateChangesDataItems0TypeBinary string = "binary"

	// StateChangesStateChangesDataItems0TypeString captures enum value "string"
	StateChangesStateChangesDataItems0TypeString string = "string"
)

// prop value enum
func (m *StateChangesStateChangesDataItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stateChangesStateChangesDataItems0TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StateChangesStateChangesDataItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *StateChangesStateChangesDataItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StateChangesStateChangesDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateChangesStateChangesDataItems0) UnmarshalBinary(b []byte) error {
	var res StateChangesStateChangesDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StateChangesStateChangesTransfersItems0 state changes state changes transfers items0
// swagger:model StateChangesStateChangesTransfersItems0
type StateChangesStateChangesTransfersItems0 struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// amount
	// Required: true
	Amount *int64 `json:"amount"`

	// asset
	// Required: true
	Asset *string `json:"asset"`
}

// Validate validates this state changes state changes transfers items0
func (m *StateChangesStateChangesTransfersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAsset(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateChangesStateChangesTransfersItems0) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *StateChangesStateChangesTransfersItems0) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *StateChangesStateChangesTransfersItems0) validateAsset(formats strfmt.Registry) error {

	if err := validate.Required("asset", "body", m.Asset); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StateChangesStateChangesTransfersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateChangesStateChangesTransfersItems0) UnmarshalBinary(b []byte) error {
	var res StateChangesStateChangesTransfersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
