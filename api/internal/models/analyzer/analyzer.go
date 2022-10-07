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

package analyzer

import (
	"fmt"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

const (
	AnalyzerTableName = "analyzers"

	AnalyzerStatusActive = "active"
)

// Analyzer represents an analyzer
type Analyzer struct {
	ID          string            `json:"id,omitempty" gorm:"column:id;primaryKey"`
	NameID      string            `json:"name_id" gorm:"column:name_id;unique;index"`
	Title       string            `json:"title" gorm:"column:title"`
	Description string            `json:"description" gorm:"column:description"`
	Status      string            `json:"status" gorm:"column:status;index"`
	Meta        datatypes.JSONMap `json:"meta" gorm:"column:meta"`
	Config      Config            `json:"config" gorm:"column:config"`
	CreatedAt   time.Time         `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time         `json:"updated_at" gorm:"column:updated_at"`
	Position    int               `json:"position" gorm:"column:position"`

	Rules []*Rule `json:"rules" gorm:"foreignKey:AnalyzerNameID;references:NameID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// TableName implements gorm Tabler interface
func (m *Analyzer) TableName() string {
	return AnalyzerTableName
}

func (m *Analyzer) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = shared.TimeUUID()
	return
}

func (m *Analyzer) BeforeSave(tx *gorm.DB) (err error) {
	if len(m.Status) == 0 {
		m.Status = AnalyzerStatusActive
	}
	return
}

// GetID returns the ID of analyzer object
func (m *Analyzer) GetID() string {
	return fmt.Sprintf("%v", m.ID)
}

// GetTags returns all the tags
func (m *Analyzer) GetTags() []string {
	tags := make([]string, 0, 10)
	tags = append(tags, m.NameID)
	return tags
}

// String returns the text representation of analyzer object
func (m *Analyzer) String() string {
	return fmt.Sprintf("%v", *m)
}

// GetIndex returns an index for specific field
func (m *Analyzer) GetIndex(field string) string {
	return m.GetIndexes()[field]
}

// GetIndexes returns all the field indexes
func (m *Analyzer) GetIndexes() map[string]string {
	return map[string]string{
		"name_id": "idx_name_id",
		"status":  "idx_status",
	}
}

// GetIndexValue return index value for specified field
func (m *Analyzer) GetIndexValue(field string) string {
	return m.GetIndexValues()[field]
}

// GetIndexValues return all field index values
func (m *Analyzer) GetIndexValues() map[string]string {
	return map[string]string{
		"name_id": m.NameID,
		"status":  m.Status,
	}
}

// Sortable checks if field is sortable.
func (m *Analyzer) Sortable(field string) bool {
	_, found := m.SortableFields()[field]
	return found
}

// SortableFields returns all sortable fields
func (m *Analyzer) SortableFields() map[string]struct{} {
	return map[string]struct{}{
		"created_at": {},
		"updated_at": {},
	}
}

// AnalyzerResponse wrappers analyzer response
//type AnalyzerResponse struct {
//	models.Pagination
//	Data []Analyzer `json:"data"`
//}
