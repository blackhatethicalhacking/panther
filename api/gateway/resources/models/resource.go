// Code generated by go-swagger; DO NOT EDIT.

// Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
// Copyright (C) 2020 Panther Labs Inc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Resource resource
//
// swagger:model Resource
type Resource struct {

	// attributes
	// Required: true
	Attributes Attributes `json:"attributes"`

	// compliance status
	// Required: true
	ComplianceStatus ComplianceStatus `json:"complianceStatus"`

	// deleted
	// Required: true
	Deleted Deleted `json:"deleted"`

	// id
	// Required: true
	ID ResourceID `json:"id"`

	// integration Id
	// Required: true
	IntegrationID IntegrationID `json:"integrationId"`

	// integration type
	// Required: true
	IntegrationType IntegrationType `json:"integrationType"`

	// last modified
	// Required: true
	// Format: date-time
	LastModified LastModified `json:"lastModified"`

	// type
	// Required: true
	Type ResourceType `json:"type"`
}

// Validate validates this resource
func (m *Resource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComplianceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resource) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateComplianceStatus(formats strfmt.Registry) error {

	if err := m.ComplianceStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("complianceStatus")
		}
		return err
	}

	return nil
}

func (m *Resource) validateDeleted(formats strfmt.Registry) error {

	if err := m.Deleted.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deleted")
		}
		return err
	}

	return nil
}

func (m *Resource) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Resource) validateIntegrationID(formats strfmt.Registry) error {

	if err := m.IntegrationID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("integrationId")
		}
		return err
	}

	return nil
}

func (m *Resource) validateIntegrationType(formats strfmt.Registry) error {

	if err := m.IntegrationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("integrationType")
		}
		return err
	}

	return nil
}

func (m *Resource) validateLastModified(formats strfmt.Registry) error {

	if err := m.LastModified.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastModified")
		}
		return err
	}

	return nil
}

func (m *Resource) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Resource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resource) UnmarshalBinary(b []byte) error {
	var res Resource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
