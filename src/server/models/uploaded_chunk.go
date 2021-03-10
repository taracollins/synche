// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UploadedChunk uploaded chunk
//
// swagger:model UploadedChunk
type UploadedChunk struct {

	// The id of the file this chunk is a part of
	CompositeFileID string `json:"compositeFileId,omitempty"`

	// directory Id
	DirectoryID DirectoryID `json:"directoryId,omitempty"`

	// The file hash of the chunk
	Hash string `json:"hash,omitempty"`
}

// Validate validates this uploaded chunk
func (m *UploadedChunk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectoryID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadedChunk) validateDirectoryID(formats strfmt.Registry) error {
	if swag.IsZero(m.DirectoryID) { // not required
		return nil
	}

	if err := m.DirectoryID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("directoryId")
		}
		return err
	}

	return nil
}

// ContextValidate validate this uploaded chunk based on the context it is used
func (m *UploadedChunk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDirectoryID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadedChunk) contextValidateDirectoryID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DirectoryID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("directoryId")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UploadedChunk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadedChunk) UnmarshalBinary(b []byte) error {
	var res UploadedChunk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}