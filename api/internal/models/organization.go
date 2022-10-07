// Copyright 2022 Cisco Systems, Inc. and its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"fmt"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

const (
	OrganizationTableName = "organizations"
)

// Organization represents an organization
type Organization struct {
	ID          string            `json:"id,omitempty" gorm:"column:id;primaryKey"`
	NameID      string            `json:"name_id" gorm:"column:name_id;unique;index"`
	Title       string            `json:"title" gorm:"column:title"`
	Description string            `json:"description" gorm:"column:description"`
	Meta        datatypes.JSONMap `json:"meta" gorm:"column:meta"`
	Contact     *Contact          `json:"contact" gorm:"column:contact"`
	CreatedAt   time.Time         `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time         `json:"updated_at" gorm:"column:updated_at"`

	Roles
}

// TableName implements gorm Tabler interface
func (m *Organization) TableName() string {
	return OrganizationTableName
}

func (m *Organization) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = shared.TimeUUID()
	return
}

func (m *Organization) BeforeSave(tx *gorm.DB) (err error) {
	return
}

// GetID returns the ID of analyzer object
func (m *Organization) GetID() string {
	return fmt.Sprintf("%v", m.ID)
}

// GetTags returns all the tags
func (m *Organization) GetTags() []string {
	tags := make([]string, 0, 10)
	tags = append(tags, m.NameID)
	return tags
}

// String returns the text representation of analyzer object
func (m *Organization) String() string {
	return fmt.Sprintf("%v", *m)
}

// GetIndex returns an index for specific field
func (m *Organization) GetIndex(field string) string {
	return m.GetIndexes()[field]
}

// GetIndexes returns all the field indexes
func (m *Organization) GetIndexes() map[string]string {
	return map[string]string{
		"name_id": "idx_name_id",
	}
}

// GetIndexValue return index value for specified field
func (m *Organization) GetIndexValue(field string) string {
	return m.GetIndexValues()[field]
}

// GetIndexValues return all field index values
func (m *Organization) GetIndexValues() map[string]string {
	return map[string]string{
		"name_id": m.NameID,
	}
}

// Sortable checks if field is sortable.
func (m *Organization) Sortable(field string) bool {
	_, found := m.SortableFields()[field]
	return found
}

// SortableFields returns all sortable fields
func (m *Organization) SortableFields() map[string]struct{} {
	return map[string]struct{}{
		"created_at": {},
		"updated_at": {},
	}
}

// OrganizationPatch represents a service with patchable fields.
type OrganizationPatch struct {
	NameID      *string            `json:"name_id"`
	Title       *string            `json:"title"`
	Description *string            `json:"description"`
	Meta        *datatypes.JSONMap `json:"meta"`
	Contact     *Contact           `json:"contact"`
	*Roles
}
