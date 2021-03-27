// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DirectoryContents directory contents
//
// swagger:model DirectoryContents
type DirectoryContents struct {

	// directory ID
	DirectoryID int64 `json:"DirectoryID,omitempty"`

	// the contents of the directory
	// Example: ["file1","file2","directory1/"]
	Listings []string `json:"Listings"`
}

// Validate validates this directory contents
func (m *DirectoryContents) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this directory contents based on context it is used
func (m *DirectoryContents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DirectoryContents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectoryContents) UnmarshalBinary(b []byte) error {
	var res DirectoryContents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
