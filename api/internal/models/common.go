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
	"github.com/getkin/kin-openapi/openapi3"
)

// ModelObject defines the interface of a model
type ModelObject interface {
	GetID() string
	String() string
	GetIndex(field string) string
	GetIndexes() map[string]string
	GetIndexValue(field string) string
	GetIndexValues() map[string]string

	Sortable(field string) bool
	SortableFields() map[string]struct{}
}

// Pagination represents pagination arguments
type Pagination struct {
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	PageSize   int `json:"page_size"`
	PageNum    int `json:"current_page"`
}

type AnalyzerConfigMap map[analyzer.SpecAnalyzer]analyzer.Config

// Scan implements sql.Scanner interface.
// See https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type.
func (m *AnalyzerConfigMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	return json.Unmarshal(bytes, &m)
}

// Value implements driver.Valuer interface.
// See https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type.
func (m AnalyzerConfigMap) Value() (driver.Value, error) { return json.Marshal(m) }

func (m *AnalyzerConfigMap) Merge(with AnalyzerConfigMap) {
	for rightAnalyzerName, rightCfgMap := range with {
		_, ok := (*m)[rightAnalyzerName]
		if !ok {
			(*m)[rightAnalyzerName] = map[string]interface{}{}
		}
		for k, v := range rightCfgMap {
			(*m)[rightAnalyzerName][k] = v
		}
	}
}

type Contact struct{ openapi3.Contact }

// Scan implements sql.Scanner interface.
// See https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type.
func (m *Contact) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, &m)
}

// Value implements driver.Valuer interface.
// See https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type.
func (m *Contact) Value() (driver.Value, error) { return json.Marshal(m) }
