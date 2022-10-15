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
	AnalyzerRuleTableName = "analyzer_rules"
)

// Rule represents an analyzer rule
type Rule struct {
	ID             string            `json:"id,omitempty" gorm:"column:id;primaryKey"`
	NameID         string            `json:"name_id" gorm:"column:name_id;unique;index"`
	AnalyzerNameID string            `json:"analyzer_name_id" gorm:"column:analyzer_name_id;index"`
	Title          string            `json:"title" gorm:"column:title"`
	Description    string            `json:"description" gorm:"column:description"`
	Severity       string            `json:"severity" gorm:"column:severity"`
	Mitigation     string            `json:"mitigation" gorm:"column:mitigation"`
	Meta           datatypes.JSONMap `json:"meta" gorm:"column:meta"`
	CreatedAt      time.Time         `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      time.Time         `json:"updated_at" gorm:"column:updated_at"`
}

// TableName implements gorm Tabler interface
func (m *Rule) TableName() string {
	return AnalyzerRuleTableName
}

func (m *Rule) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = shared.TimeUUID()
	return
}

// GetID returns the ID of analyzerRule object
func (m *Rule) GetID() string {
	return fmt.Sprintf("%v", m.ID)
}

// GetTags returns all the tags
func (m *Rule) GetTags() []string {
	tags := make([]string, 0, 10)
	tags = append(tags, m.NameID)
	return tags
}

// String returns the text representation of analyzerRule object
func (m *Rule) String() string {
	return fmt.Sprintf("%v", *m)
}

// GetIndex returns an index for specific field
func (m *Rule) GetIndex(field string) string {
	return m.GetIndexes()[field]
}

// GetIndexes returns all the field indexes
func (m *Rule) GetIndexes() map[string]string {
	return map[string]string{
		"name_id":          "idx_name_id",
		"analyzer_name_id": "idx_analyzer_name_id",
	}
}

// GetIndexValue return index value for specified field
func (m *Rule) GetIndexValue(field string) string {
	return m.GetIndexValues()[field]
}

// GetIndexValues return all field index values
func (m *Rule) GetIndexValues() map[string]string {
	return map[string]string{
		"name_id":          m.NameID,
		"analyzer_name_id": m.AnalyzerNameID,
	}
}

// Sortable checks if field is sortable.
func (m *Rule) Sortable(field string) bool {
	_, found := m.SortableFields()[field]
	return found
}

// SortableFields returns all sortable fields
func (m *Rule) SortableFields() map[string]struct{} {
	return map[string]struct{}{
		"created_at": {},
		"updated_at": {},
	}
}

// AnalyzerRuleResponse wrappers service response
//type AnalyzerRuleResponse struct {
//	models.Pagination
//	Data []Analyzer `json:"data"`
//}
