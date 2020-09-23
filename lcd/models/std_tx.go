// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StdTx std tx
//
// swagger:model StdTx
type StdTx struct {

	// fee
	Fee *StdTxFee `json:"fee,omitempty"`

	// memo
	Memo string `json:"memo,omitempty"`

	// msg
	Msg []*Msg `json:"msg"`

	// signatures
	Signatures []*TxSignature `json:"signatures"`
}

// Validate validates this std tx
func (m *StdTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StdTx) validateFee(formats strfmt.Registry) error {

	if swag.IsZero(m.Fee) { // not required
		return nil
	}

	if m.Fee != nil {
		if err := m.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fee")
			}
			return err
		}
	}

	return nil
}

func (m *StdTx) validateMsg(formats strfmt.Registry) error {

	if swag.IsZero(m.Msg) { // not required
		return nil
	}

	for i := 0; i < len(m.Msg); i++ {
		if swag.IsZero(m.Msg[i]) { // not required
			continue
		}

		if m.Msg[i] != nil {
			if err := m.Msg[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("msg" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StdTx) validateSignatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Signatures) { // not required
		return nil
	}

	for i := 0; i < len(m.Signatures); i++ {
		if swag.IsZero(m.Signatures[i]) { // not required
			continue
		}

		if m.Signatures[i] != nil {
			if err := m.Signatures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StdTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StdTx) UnmarshalBinary(b []byte) error {
	var res StdTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StdTxFee std tx fee
//
// swagger:model StdTxFee
type StdTxFee struct {

	// amount
	Amount []*Coin `json:"amount"`

	// gas
	Gas string `json:"gas,omitempty"`
}

// Validate validates this std tx fee
func (m *StdTxFee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StdTxFee) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	for i := 0; i < len(m.Amount); i++ {
		if swag.IsZero(m.Amount[i]) { // not required
			continue
		}

		if m.Amount[i] != nil {
			if err := m.Amount[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StdTxFee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StdTxFee) UnmarshalBinary(b []byte) error {
	var res StdTxFee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
