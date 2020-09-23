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

// PaginatedQueryTxs paginated query txs
//
// swagger:model PaginatedQueryTxs
type PaginatedQueryTxs struct {

	// count
	Count string `json:"count,omitempty"`

	// limit
	Limit string `json:"limit,omitempty"`

	// page number
	PageNumber string `json:"page_number,omitempty"`

	// page total
	PageTotal string `json:"page_total,omitempty"`

	// total count
	TotalCount string `json:"total_count,omitempty"`

	// txs
	Txs []*TxQuery `json:"txs"`
}

// Validate validates this paginated query txs
func (m *PaginatedQueryTxs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTxs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginatedQueryTxs) validateTxs(formats strfmt.Registry) error {

	if swag.IsZero(m.Txs) { // not required
		return nil
	}

	for i := 0; i < len(m.Txs); i++ {
		if swag.IsZero(m.Txs[i]) { // not required
			continue
		}

		if m.Txs[i] != nil {
			if err := m.Txs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("txs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaginatedQueryTxs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginatedQueryTxs) UnmarshalBinary(b []byte) error {
	var res PaginatedQueryTxs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
