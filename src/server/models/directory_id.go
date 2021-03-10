// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// DirectoryID the id of the directory to list
// Example: d701748f0851
//
// swagger:model DirectoryId
type DirectoryID string

// Validate validates this directory Id
func (m DirectoryID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this directory Id based on context it is used
func (m DirectoryID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}