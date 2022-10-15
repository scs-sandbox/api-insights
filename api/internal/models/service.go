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
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

const (
	ServiceTableName         = "services"
	ServiceVisibilityPublic  = "public"
	ServiceVisibilityPrivate = "private"
)

// Service represents a service
type Service struct {
	ID             string            `json:"id,omitempty" gorm:"column:id;primaryKey"`
	AdditionalInfo datatypes.JSONMap `json:"additional_info" gorm:"column:additional_info"`
	Contact        *Contact          `json:"contact" gorm:"column:contact"`
	Description    string            `json:"description" gorm:"column:description"`
	NameID         string            `json:"name_id" gorm:"column:name_id;unique;index"`
	OrganizationID string            `json:"organization_id" gorm:"column:organization_id;index"`
	ProductTag     string            `json:"product_tag" gorm:"column:product_tag;index"`
	Title          string            `json:"title" gorm:"column:title"`
	Visibility     string            `json:"visibility" gorm:"column:visibility;index"`
	CreatedAt      time.Time         `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      time.Time         `json:"updated_at" gorm:"column:updated_at"`

	AnalyzersConfigs AnalyzerConfigMap `json:"analyzers_configs,omitempty" gorm:"column:analyzers_configs"`

	Summary *ServiceSummary `json:"summary" gorm:"column:summary"`
}

// TableName implements gorm Tabler interface
func (m *Service) TableName() string {
	return ServiceTableName
}

func (m *Service) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = shared.TimeUUID()
	return
}

func (m *Service) BeforeSave(tx *gorm.DB) (err error) {
	if len(m.Visibility) == 0 {
		m.Visibility = ServiceVisibilityPublic
	}
	return
}

// GetID returns the ID of service object
func (m *Service) GetID() string {
	return fmt.Sprintf("%v", m.ID)
}

// GetTags returns all the tags
func (m *Service) GetTags() []string {
	tags := make([]string, 0, 10)
	tags = append(tags, m.NameID)
	tags = append(tags, m.ProductTag)
	tags = append(tags, m.OrganizationID)
	tags = append(tags, m.Visibility)
	return tags
}

// String returns the text representation of service object
func (m *Service) String() string {
	return fmt.Sprintf("%v", *m)
}

// GetIndex returns an index for specific field
func (m *Service) GetIndex(field string) string {
	return m.GetIndexes()[field]
}

// GetIndexes returns all the field indexes
func (m *Service) GetIndexes() map[string]string {
	return map[string]string{
		"name_id":         "idx_name_id",
		"product_tag":     "idx_product_tag",
		"organization_id": "idx_organization_id",
		"visibility":      "idx_visibility",
	}
}

// GetIndexValue return index value for specified field
func (m *Service) GetIndexValue(field string) string {
	return m.GetIndexValues()[field]
}

// GetIndexValues return all field index values
func (m *Service) GetIndexValues() map[string]string {
	return map[string]string{
		"name_id":         m.NameID,
		"product_tag":     m.ProductTag,
		"organization_id": m.OrganizationID,
		"visibility":      m.Visibility,
	}
}

// Sortable checks if field is sortable.
func (m *Service) Sortable(field string) bool {
	_, found := m.SortableFields()[field]
	return found
}

// SortableFields returns all sortable fields
func (m *Service) SortableFields() map[string]struct{} {
	return map[string]struct{}{
		"created_at": {},
		"updated_at": {},
	}
}

func (m *Service) SetSummary(score int, version string, revision string, updatedAt time.Time) {
	if m.Summary == nil {
		m.Summary = &ServiceSummary{}
	}
	m.Summary.Score = &score
	m.Summary.Version = version
	m.Summary.Revision = revision
	m.Summary.UpdatedAt = updatedAt
}

// GetNameID retrieves the service name ID, using the following precedence, from:
//   - Service.AnalyzersConfigs[name][analyzer.ConfigServiceNameID], or
//   - analyzerCfg[analyzer.ConfigServiceNameID], or
//   - Service.NameID.
func (m *Service) GetNameID(name analyzer.SpecAnalyzer, analyzerCfg analyzer.Config) string {
	if analyzerCfg, ok := m.AnalyzersConfigs[name]; ok {
		id := analyzerCfg.ServiceNameID()
		if id != "" {
			return id
		}
	}
	if len(analyzerCfg) > 0 {
		id := analyzerCfg.ServiceNameIDFromTemplate(m.NameID)
		if id != "" {
			return id
		}
	}
	return m.NameID
}

// ServiceResponse wrappers service response
type ServiceResponse struct {
	Pagination
	Data []Service `json:"data"`
}

type ServiceSummary struct {
	Score     *int      `json:"score" gorm:"column:score"`
	Version   string    `json:"version" gorm:"column:version"`
	Revision  string    `json:"revision" gorm:"column:revision"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// Scan implements sql.Scanner interface.
// See https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type.
func (m *ServiceSummary) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, &m)
}

// Value implements driver.Valuer interface.
// See https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type.
func (m ServiceSummary) Value() (driver.Value, error) { return json.Marshal(m) }

// ServicePatch represents a service with patchable fields.
type ServicePatch struct {
	AdditionalInfo   *datatypes.JSONMap `json:"additional_info"`
	Contact          *Contact           `json:"contact"`
	Description      *string            `json:"description"`
	NameID           *string            `json:"name_id"`
	OrganizationID   *string            `json:"organization_id"`
	ProductTag       *string            `json:"product_tag"`
	Title            *string            `json:"title"`
	AnalyzersConfigs *AnalyzerConfigMap `json:"analyzers_configs"`
	Visibility       *string            `json:"visibility"`
}
