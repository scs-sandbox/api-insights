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
	"encoding/json"
	"fmt"
	"time"

	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	SpecAnalysisTableName = "spec_analyses"
)

// SpecAnalysis represents a specAnalysis
type SpecAnalysis struct {
	ID       string                `json:"id,omitempty" gorm:"column:id;primaryKey"`
	Analyzer analyzer.SpecAnalyzer `json:"analyzer" gorm:"column:analyzer;index"`

	SpecAnalysisConfig
	SpecAnalysisResult

	Score     *int      `json:"score" gorm:"column:score"`
	ServiceID string    `json:"service_id" gorm:"column:service_id;index:svc_spec_created_idx"`
	SpecID    string    `json:"spec_id" gorm:"column:spec_id;index;index:svc_spec_created_idx"`
	Status    string    `json:"status" gorm:"column:status;index"` // Submitted, Invalid, Analyzed
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;index:svc_spec_created_idx"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName implements gorm Tabler interface
func (m *SpecAnalysis) TableName() string {
	return SpecAnalysisTableName
}

func (m *SpecAnalysis) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = shared.TimeUUID()
	return
}

func (m *SpecAnalysis) BeforeSave(tx *gorm.DB) (err error) {
	if m.Config != nil {
		if m.RawConfig, err = json.Marshal(m.Config); err != nil {
			return err
		}
	}
	if m.Result != nil {
		if m.RawResult, err = json.Marshal(m.Result); err != nil {
			return err
		}
	}
	return
}

func (m *SpecAnalysis) AfterFind(tx *gorm.DB) (err error) {
	m.Result = &analyzer.Result{}
	if m.RawConfig != nil {
		m.Config = analyzer.Config{}
		if err = json.Unmarshal(m.RawConfig, &m.Config); err != nil {
			return err
		}
	}
	if m.RawResult != nil && m.Result != nil {
		if err = json.Unmarshal(m.RawResult, m.Result); err != nil {
			return err
		}
	}
	return
}

// GetID returns the ID of specAnalysis object
func (m *SpecAnalysis) GetID() string {
	return fmt.Sprintf("%v", m.ID)
}

// GetTags returns all the tags
func (m *SpecAnalysis) GetTags() []string {
	tags := make([]string, 0, 10)
	tags = append(tags, m.ServiceID)
	tags = append(tags, m.SpecID)
	tags = append(tags, string(m.Analyzer))
	tags = append(tags, m.Status)
	return tags
}

// String returns the text representation of specAnalysis object
func (m *SpecAnalysis) String() string {
	return fmt.Sprintf("%v", *m)
}

// GetIndex returns an index for specific field
func (m *SpecAnalysis) GetIndex(field string) string {
	return m.GetIndexes()[field]
}

// GetIndexes returns all the field indexes
func (m *SpecAnalysis) GetIndexes() map[string]string {
	return map[string]string{
		"service_id": "idx_service_id",
		"spec_id":    "idx_spec_id",
		"analyzer":   "idx_analyzer",
		"status":     "idx_status",
	}
}

// GetIndexValue return index value for specified field
func (m *SpecAnalysis) GetIndexValue(field string) string {
	return m.GetIndexValues()[field]
}

// GetIndexValues return all field index values
func (m *SpecAnalysis) GetIndexValues() map[string]string {
	return map[string]string{
		"service_id": m.ServiceID,
		"spec_id":    m.SpecID,
		"analyzer":   string(m.Analyzer),
		"status":     m.Status,
	}
}

// Sortable checks if field is sortable.
func (m *SpecAnalysis) Sortable(field string) bool {
	_, found := m.SortableFields()[field]
	return found
}

// SortableFields returns all sortable fields
func (m *SpecAnalysis) SortableFields() map[string]struct{} {
	return map[string]struct{}{
		"created_at": {},
		"updated_at": {},
	}
}

func (m *SpecAnalysis) SetResult(result *analyzer.Result, status string) error {
	if result == nil {
		return fmt.Errorf("spec_analysis: cannot set nil result")
	}
	if status == "" {
		return fmt.Errorf("spec_analysis: cannot set empty status")
	}
	m.Result = result
	m.Status = status
	return nil
}

func (m *SpecAnalysis) SetScore(score int) error {
	m.Score = &score
	return nil
}

func (m *SpecAnalysis) OmitResultFindings() { m.Result.Findings = nil }

// SpecAnalysisResponseWrapper wrappers specAnalysis response
type SpecAnalysisResponseWrapper struct {
	Pagination
	Data []SpecAnalysis `json:"data"`
}

type SpecAnalysisConfig struct {
	Config    analyzer.Config `json:"config,omitempty" gorm:"-"`
	RawConfig datatypes.JSON  `json:"-" gorm:"column:config"`
}

type SpecAnalysisResult struct {
	Result    *analyzer.Result `json:"result" gorm:"-"`
	RawResult datatypes.JSON   `json:"-" gorm:"column:result"`
}

// SpecAnalysisRequest represents a request for a SpecAnalysis
type SpecAnalysisRequest struct {
	Analyzers        []analyzer.SpecAnalyzer `json:"analyzers"`
	AnalyzersConfigs AnalyzerConfigMap       `json:"analyzers_configs,omitempty"`

	Spec    *Spec    `json:"spec,omitempty"`
	Service *Service `json:"service,omitempty"`

	ActiveAnalyzers map[analyzer.SpecAnalyzer]*analyzer.Analyzer `json:"-"`
}

func (m *SpecAnalysisRequest) HasSpec() bool {
	return m.Spec != nil && m.Spec.Doc != nil
}

type SpecAnalysisResponse struct {
	Results map[analyzer.SpecAnalyzer]*SpecAnalysis `json:"results,omitempty"`

	SpecScore int `json:"spec_score"`
}

// SpecDocAnalyzer represents the interface for analyzing a SpecDoc using (optional) analyzer.Config.
type SpecDocAnalyzer interface {
	Analyze(doc SpecDoc, cfgMap analyzer.Config, serviceNameID *string) (*analyzer.Result, error)
}

// DistinctSpecAnalyses filters out duplicate (uniqueness based on SpecAnalysis.SpecID & SpecAnalysis.Analyzer) spec analyses.
func DistinctSpecAnalyses(from []*SpecAnalysis) (to []*SpecAnalysis) {
	byDistinctSpecIDAnalyzers := map[string]map[analyzer.SpecAnalyzer]struct{}{}
	for _, f := range from {
		specID := f.SpecID
		if _, ok := byDistinctSpecIDAnalyzers[specID]; ok {
			if _, ok := byDistinctSpecIDAnalyzers[specID][f.Analyzer]; ok {
				continue
			}
			byDistinctSpecIDAnalyzers[specID][f.Analyzer] = struct{}{}
		} else {
			byDistinctSpecIDAnalyzers[specID] = map[analyzer.SpecAnalyzer]struct{}{f.Analyzer: {}}
		}
		to = append(to, f)
	}
	return
}
