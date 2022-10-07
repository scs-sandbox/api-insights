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
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

func TestDistinctSpecAnalyses(t *testing.T) {
	type args struct {
		from []*SpecAnalysis
	}
	score := 80
	tests := []struct {
		name   string
		args   args
		wantTo []*SpecAnalysis
	}{
		{
			name: "normal",
			args: args{from: []*SpecAnalysis{
				{Analyzer: analyzer.Completeness, Score: &score},
				{Analyzer: analyzer.Completeness, Score: &score},
				{Analyzer: analyzer.Completeness, Score: &score},
				{Analyzer: analyzer.Security, Score: &score},
			}},
			wantTo: []*SpecAnalysis{
				{Analyzer: analyzer.Completeness, Score: &score},
				{Analyzer: analyzer.Security, Score: &score},
			},
		},
		{
			name:   "empty analysis",
			args:   args{from: []*SpecAnalysis{}},
			wantTo: nil,
		},
		{
			name:   "nil analysis",
			args:   args{from: []*SpecAnalysis{}},
			wantTo: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTo := DistinctSpecAnalyses(tt.args.from); !reflect.DeepEqual(gotTo, tt.wantTo) {
				t.Errorf("DistinctSpecAnalyses() = %v, want %v", gotTo, tt.wantTo)
			}
		})
	}
}

func TestSpecAnalysisRequest_HasSpec(t *testing.T) {
	type fields struct {
		Analyzers        []analyzer.SpecAnalyzer
		AnalyzersConfigs AnalyzerConfigMap
		Spec             *Spec
		Service          *Service
		ActiveAnalyzers  map[analyzer.SpecAnalyzer]*analyzer.Analyzer
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "normal",
			fields: fields{
				Spec: &Spec{Doc: SpecDoc(loadSpec("testdata/petstore-v2.json"))},
			},
			want: true,
		},
		{
			name: "nil spec",
			fields: fields{
				Spec: nil,
			},
			want: false,
		},
		{
			name: "nil doc",
			fields: fields{
				Spec: &Spec{Doc: nil},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysisRequest{
				Analyzers:        tt.fields.Analyzers,
				AnalyzersConfigs: tt.fields.AnalyzersConfigs,
				Spec:             tt.fields.Spec,
				Service:          tt.fields.Service,
				ActiveAnalyzers:  tt.fields.ActiveAnalyzers,
			}
			if got := m.HasSpec(); got != tt.want {
				t.Errorf("HasSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecAnalysis_AfterFind(t *testing.T) {
	type fields struct {
		ID                 string
		Analyzer           analyzer.SpecAnalyzer
		SpecAnalysisConfig SpecAnalysisConfig
		SpecAnalysisResult SpecAnalysisResult
		Score              *int
		ServiceID          string
		SpecID             string
		Status             string
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				SpecAnalysisConfig: SpecAnalysisConfig{
					RawConfig: []byte(`{"score_config":{"analyzer_weight": 0.3}}`),
				},
				SpecAnalysisResult: SpecAnalysisResult{
					RawResult: []byte(`{
					  "summary": {
						"stats": {
						  "count": 0,
						  "occurrences": 0,
						  "hint": {
							"count": 0,
							"occurrences": 0,
							"rules": {}
						  },
						  "info": {
							"count": 0,
							"occurrences": 0,
							"rules": {}
						  },
						  "warning": {
							"count": 0,
							"occurrences": 0,
							"rules": {}
						  },
						  "error": {
							"count": 0,
							"occurrences": 0,
							"rules": {}
						  }
						}
					  },
					  "findings": {
						"error": {
						  "rules": null
						},
						"hint": {
						  "rules": null
						},
						"info": {
						  "rules": null
						},
						"warning": {
						  "rules": null
						}
					  }
					}`),
				},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "nil raw config",
			fields: fields{
				SpecAnalysisConfig: SpecAnalysisConfig{RawConfig: nil},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "nil raw result",
			fields: fields{
				SpecAnalysisResult: SpecAnalysisResult{RawResult: nil},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "invalid raw config",
			fields: fields{
				SpecAnalysisConfig: SpecAnalysisConfig{RawConfig: []byte(`abc`)},
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "invalid raw result",
			fields: fields{
				SpecAnalysisResult: SpecAnalysisResult{RawResult: []byte(`abc`)},
			},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysis{
				ID:                 tt.fields.ID,
				Analyzer:           tt.fields.Analyzer,
				SpecAnalysisConfig: tt.fields.SpecAnalysisConfig,
				SpecAnalysisResult: tt.fields.SpecAnalysisResult,
				Score:              tt.fields.Score,
				ServiceID:          tt.fields.ServiceID,
				SpecID:             tt.fields.SpecID,
				Status:             tt.fields.Status,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
			}
			if err := m.AfterFind(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("AfterFind() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && tt.fields.SpecAnalysisConfig.RawConfig != nil {
				assert.NotNil(t, m.Config)
			}
			if !tt.wantErr && tt.fields.SpecAnalysisResult.RawResult != nil {
				assert.NotNil(t, m.Result)
			}
		})
	}
}

func TestSpecAnalysis_BeforeCreate(t *testing.T) {
	type fields struct {
		ID                 string
		Analyzer           analyzer.SpecAnalyzer
		SpecAnalysisConfig SpecAnalysisConfig
		SpecAnalysisResult SpecAnalysisResult
		Score              *int
		ServiceID          string
		SpecID             string
		Status             string
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "normal - init id",
			fields:  fields{ID: ""},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysis{
				ID:                 tt.fields.ID,
				Analyzer:           tt.fields.Analyzer,
				SpecAnalysisConfig: tt.fields.SpecAnalysisConfig,
				SpecAnalysisResult: tt.fields.SpecAnalysisResult,
				Score:              tt.fields.Score,
				ServiceID:          tt.fields.ServiceID,
				SpecID:             tt.fields.SpecID,
				Status:             tt.fields.Status,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
			}
			if err := m.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NotEmpty(t, m.ID)
		})
	}
}

func TestSpecAnalysis_BeforeSave(t *testing.T) {
	type fields struct {
		ID                 string
		Analyzer           analyzer.SpecAnalyzer
		SpecAnalysisConfig SpecAnalysisConfig
		SpecAnalysisResult SpecAnalysisResult
		Score              *int
		ServiceID          string
		SpecID             string
		Status             string
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				SpecAnalysisConfig: SpecAnalysisConfig{
					Config:    map[string]interface{}{},
					RawConfig: nil,
				},
				SpecAnalysisResult: SpecAnalysisResult{
					Result: &analyzer.Result{
						Summary:  &analyzer.ResultSummary{},
						Findings: analyzer.SeverityRuleFindings{},
					},
					RawResult: nil,
				},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "nil config",
			fields: fields{
				SpecAnalysisConfig: SpecAnalysisConfig{
					Config:    nil,
					RawConfig: nil,
				},
				SpecAnalysisResult: SpecAnalysisResult{
					Result: &analyzer.Result{
						Summary:  &analyzer.ResultSummary{},
						Findings: analyzer.SeverityRuleFindings{},
					},
					RawResult: nil,
				},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "nil result",
			fields: fields{
				SpecAnalysisConfig: SpecAnalysisConfig{
					Config:    map[string]interface{}{},
					RawConfig: nil,
				},
				SpecAnalysisResult: SpecAnalysisResult{
					Result:    nil,
					RawResult: nil,
				},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "invalid config",
			fields: fields{
				SpecAnalysisConfig: SpecAnalysisConfig{
					Config:    map[string]interface{}{"test": make(chan int)},
					RawConfig: nil,
				},
				SpecAnalysisResult: SpecAnalysisResult{
					Result: &analyzer.Result{
						Summary:  &analyzer.ResultSummary{},
						Findings: analyzer.SeverityRuleFindings{},
					},
					RawResult: nil,
				},
			},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysis{
				ID:                 tt.fields.ID,
				Analyzer:           tt.fields.Analyzer,
				SpecAnalysisConfig: tt.fields.SpecAnalysisConfig,
				SpecAnalysisResult: tt.fields.SpecAnalysisResult,
				Score:              tt.fields.Score,
				ServiceID:          tt.fields.ServiceID,
				SpecID:             tt.fields.SpecID,
				Status:             tt.fields.Status,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
			}

			err := m.BeforeSave(tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeforeSave() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && tt.fields.SpecAnalysisConfig.Config != nil {
				assert.NotNil(t, m.RawConfig)
			}
			if err == nil && tt.fields.SpecAnalysisResult.Result != nil {
				assert.NotNil(t, m.RawResult)
			}
		})
	}
}

func TestSpecAnalysis_GetID(t *testing.T) {
	type fields struct {
		ID                 string
		Analyzer           analyzer.SpecAnalyzer
		SpecAnalysisConfig SpecAnalysisConfig
		SpecAnalysisResult SpecAnalysisResult
		Score              *int
		ServiceID          string
		SpecID             string
		Status             string
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal",
			fields: fields{ID: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysis{
				ID:                 tt.fields.ID,
				Analyzer:           tt.fields.Analyzer,
				SpecAnalysisConfig: tt.fields.SpecAnalysisConfig,
				SpecAnalysisResult: tt.fields.SpecAnalysisResult,
				Score:              tt.fields.Score,
				ServiceID:          tt.fields.ServiceID,
				SpecID:             tt.fields.SpecID,
				Status:             tt.fields.Status,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
			}
			if got := m.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecAnalysis_GetIndex(t *testing.T) {
	type fields struct {
		ID                 string
		Analyzer           analyzer.SpecAnalyzer
		SpecAnalysisConfig SpecAnalysisConfig
		SpecAnalysisResult SpecAnalysisResult
		Score              *int
		ServiceID          string
		SpecID             string
		Status             string
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "index service_id",
			fields: fields{},
			args:   args{field: "service_id"},
			want:   "idx_service_id",
		},
		{
			name:   "index spec_id",
			fields: fields{},
			args:   args{field: "spec_id"},
			want:   "idx_spec_id",
		},
		{
			name:   "index analyzer",
			fields: fields{},
			args:   args{field: "analyzer"},
			want:   "idx_analyzer",
		},
		{
			name:   "index status",
			fields: fields{},
			args:   args{field: "status"},
			want:   "idx_status",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysis{
				ID:                 tt.fields.ID,
				Analyzer:           tt.fields.Analyzer,
				SpecAnalysisConfig: tt.fields.SpecAnalysisConfig,
				SpecAnalysisResult: tt.fields.SpecAnalysisResult,
				Score:              tt.fields.Score,
				ServiceID:          tt.fields.ServiceID,
				SpecID:             tt.fields.SpecID,
				Status:             tt.fields.Status,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
			}
			if got := m.GetIndex(tt.args.field); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecAnalysis_GetIndexValue(t *testing.T) {
	type fields struct {
		ID                 string
		Analyzer           analyzer.SpecAnalyzer
		SpecAnalysisConfig SpecAnalysisConfig
		SpecAnalysisResult SpecAnalysisResult
		Score              *int
		ServiceID          string
		SpecID             string
		Status             string
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "index service_id",
			fields: fields{ServiceID: "test"},
			args:   args{field: "service_id"},
			want:   "test",
		},
		{
			name:   "index spec_id",
			fields: fields{SpecID: "test"},
			args:   args{field: "spec_id"},
			want:   "test",
		},
		{
			name:   "index analyzer",
			fields: fields{Analyzer: "security"},
			args:   args{field: "analyzer"},
			want:   "security",
		},
		{
			name:   "index status",
			fields: fields{Status: "Development"},
			args:   args{field: "status"},
			want:   "Development",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysis{
				ID:                 tt.fields.ID,
				Analyzer:           tt.fields.Analyzer,
				SpecAnalysisConfig: tt.fields.SpecAnalysisConfig,
				SpecAnalysisResult: tt.fields.SpecAnalysisResult,
				Score:              tt.fields.Score,
				ServiceID:          tt.fields.ServiceID,
				SpecID:             tt.fields.SpecID,
				Status:             tt.fields.Status,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
			}
			if got := m.GetIndexValue(tt.args.field); got != tt.want {
				t.Errorf("GetIndexValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecAnalysis_GetTags(t *testing.T) {
	type fields struct {
		ID                 string
		Analyzer           analyzer.SpecAnalyzer
		SpecAnalysisConfig SpecAnalysisConfig
		SpecAnalysisResult SpecAnalysisResult
		Score              *int
		ServiceID          string
		SpecID             string
		Status             string
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "normal",
			fields: fields{
				ServiceID: "catalogue",
				SpecID:    "1",
				Analyzer:  "security",
				Status:    "Development",
			},
			want: []string{"catalogue", "1", "security", "Development"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysis{
				ID:                 tt.fields.ID,
				Analyzer:           tt.fields.Analyzer,
				SpecAnalysisConfig: tt.fields.SpecAnalysisConfig,
				SpecAnalysisResult: tt.fields.SpecAnalysisResult,
				Score:              tt.fields.Score,
				ServiceID:          tt.fields.ServiceID,
				SpecID:             tt.fields.SpecID,
				Status:             tt.fields.Status,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
			}
			if got := m.GetTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecAnalysis_SetResult(t *testing.T) {
	type fields struct {
		ID                 string
		Analyzer           analyzer.SpecAnalyzer
		SpecAnalysisConfig SpecAnalysisConfig
		SpecAnalysisResult SpecAnalysisResult
		Score              *int
		ServiceID          string
		SpecID             string
		Status             string
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	type args struct {
		result *analyzer.Result
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{},
			args: args{
				result: &analyzer.Result{},
				status: "active",
			},
			wantErr: false,
		},
		{
			name:   "nil result",
			fields: fields{},
			args: args{
				result: nil,
				status: "active",
			},
			wantErr: true,
		},
		{
			name:   "empty status",
			fields: fields{},
			args: args{
				result: &analyzer.Result{},
				status: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysis{
				ID:                 tt.fields.ID,
				Analyzer:           tt.fields.Analyzer,
				SpecAnalysisConfig: tt.fields.SpecAnalysisConfig,
				SpecAnalysisResult: tt.fields.SpecAnalysisResult,
				Score:              tt.fields.Score,
				ServiceID:          tt.fields.ServiceID,
				SpecID:             tt.fields.SpecID,
				Status:             tt.fields.Status,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
			}
			if err := m.SetResult(tt.args.result, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("SetResult() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				assert.NotNil(t, m.Result)
				assert.Equal(t, tt.args.status, m.Status)
			}
		})
	}
}

func TestSpecAnalysis_SetScore(t *testing.T) {
	type fields struct {
		ID                 string
		Analyzer           analyzer.SpecAnalyzer
		SpecAnalysisConfig SpecAnalysisConfig
		SpecAnalysisResult SpecAnalysisResult
		Score              *int
		ServiceID          string
		SpecID             string
		Status             string
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	type args struct {
		score int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Score: nil,
			},
			args:    args{score: 100},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysis{
				ID:                 tt.fields.ID,
				Analyzer:           tt.fields.Analyzer,
				SpecAnalysisConfig: tt.fields.SpecAnalysisConfig,
				SpecAnalysisResult: tt.fields.SpecAnalysisResult,
				Score:              tt.fields.Score,
				ServiceID:          tt.fields.ServiceID,
				SpecID:             tt.fields.SpecID,
				Status:             tt.fields.Status,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
			}
			if err := m.SetScore(tt.args.score); (err != nil) != tt.wantErr {
				t.Errorf("SetScore() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				assert.NotNil(t, m.Score)
				assert.Equal(t, tt.args.score, *m.Score)
			}
		})
	}
}

func TestSpecAnalysis_Sortable(t *testing.T) {
	type fields struct {
		ID                 string
		Analyzer           analyzer.SpecAnalyzer
		SpecAnalysisConfig SpecAnalysisConfig
		SpecAnalysisResult SpecAnalysisResult
		Score              *int
		ServiceID          string
		SpecID             string
		Status             string
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "sort by created_at",
			fields: fields{},
			args:   args{field: "created_at"},
			want:   true,
		},
		{
			name:   "sort by updated_at",
			fields: fields{},
			args:   args{field: "created_at"},
			want:   true,
		},
		{
			name:   "sort by unsupported field",
			fields: fields{},
			args:   args{field: "unsupported"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysis{
				ID:                 tt.fields.ID,
				Analyzer:           tt.fields.Analyzer,
				SpecAnalysisConfig: tt.fields.SpecAnalysisConfig,
				SpecAnalysisResult: tt.fields.SpecAnalysisResult,
				Score:              tt.fields.Score,
				ServiceID:          tt.fields.ServiceID,
				SpecID:             tt.fields.SpecID,
				Status:             tt.fields.Status,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
			}
			if got := m.Sortable(tt.args.field); got != tt.want {
				t.Errorf("Sortable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecAnalysis_String(t *testing.T) {
	type fields struct {
		ID                 string
		Analyzer           analyzer.SpecAnalyzer
		SpecAnalysisConfig SpecAnalysisConfig
		SpecAnalysisResult SpecAnalysisResult
		Score              *int
		ServiceID          string
		SpecID             string
		Status             string
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal",
			fields: fields{ServiceID: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysis{
				ID:                 tt.fields.ID,
				Analyzer:           tt.fields.Analyzer,
				SpecAnalysisConfig: tt.fields.SpecAnalysisConfig,
				SpecAnalysisResult: tt.fields.SpecAnalysisResult,
				Score:              tt.fields.Score,
				ServiceID:          tt.fields.ServiceID,
				SpecID:             tt.fields.SpecID,
				Status:             tt.fields.Status,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
			}
			got := m.String()
			assert.NotEmpty(t, got)
		})
	}
}

func TestSpecAnalysis_TableName(t *testing.T) {
	type fields struct {
		ID                 string
		Analyzer           analyzer.SpecAnalyzer
		SpecAnalysisConfig SpecAnalysisConfig
		SpecAnalysisResult SpecAnalysisResult
		Score              *int
		ServiceID          string
		SpecID             string
		Status             string
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal",
			fields: fields{},
			want:   SpecAnalysisTableName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecAnalysis{
				ID:                 tt.fields.ID,
				Analyzer:           tt.fields.Analyzer,
				SpecAnalysisConfig: tt.fields.SpecAnalysisConfig,
				SpecAnalysisResult: tt.fields.SpecAnalysisResult,
				Score:              tt.fields.Score,
				ServiceID:          tt.fields.ServiceID,
				SpecID:             tt.fields.SpecID,
				Status:             tt.fields.Status,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
			}
			if got := m.TableName(); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
