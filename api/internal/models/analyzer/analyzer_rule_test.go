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
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

func TestAnalyzerRule_BeforeCreate(t *testing.T) {
	type fields struct {
		ID             string
		NameID         string
		AnalyzerNameID string
		Title          string
		Description    string
		Severity       string
		Mitigation     string
		Meta           datatypes.JSONMap
		CreatedAt      time.Time
		UpdatedAt      time.Time
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
			m := &Rule{
				ID:             tt.fields.ID,
				NameID:         tt.fields.NameID,
				AnalyzerNameID: tt.fields.AnalyzerNameID,
				Title:          tt.fields.Title,
				Description:    tt.fields.Description,
				Severity:       tt.fields.Severity,
				Mitigation:     tt.fields.Mitigation,
				Meta:           tt.fields.Meta,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
			}
			if err := m.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NotEmpty(t, m.ID)
		})
	}
}

func TestAnalyzerRule_GetID(t *testing.T) {
	type fields struct {
		ID             string
		NameID         string
		AnalyzerNameID string
		Title          string
		Description    string
		Severity       string
		Mitigation     string
		Meta           datatypes.JSONMap
		CreatedAt      time.Time
		UpdatedAt      time.Time
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
			m := &Rule{
				ID:             tt.fields.ID,
				NameID:         tt.fields.NameID,
				AnalyzerNameID: tt.fields.AnalyzerNameID,
				Title:          tt.fields.Title,
				Description:    tt.fields.Description,
				Severity:       tt.fields.Severity,
				Mitigation:     tt.fields.Mitigation,
				Meta:           tt.fields.Meta,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
			}
			if got := m.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerRule_GetIndex(t *testing.T) {
	type fields struct {
		ID             string
		NameID         string
		AnalyzerNameID string
		Title          string
		Description    string
		Severity       string
		Mitigation     string
		Meta           datatypes.JSONMap
		CreatedAt      time.Time
		UpdatedAt      time.Time
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
			name:   "index name_id",
			fields: fields{},
			args:   args{field: "name_id"},
			want:   "idx_name_id",
		},
		{
			name:   "index analyzer_name_id",
			fields: fields{},
			args:   args{field: "analyzer_name_id"},
			want:   "idx_analyzer_name_id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Rule{
				ID:             tt.fields.ID,
				NameID:         tt.fields.NameID,
				AnalyzerNameID: tt.fields.AnalyzerNameID,
				Title:          tt.fields.Title,
				Description:    tt.fields.Description,
				Severity:       tt.fields.Severity,
				Mitigation:     tt.fields.Mitigation,
				Meta:           tt.fields.Meta,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
			}
			if got := m.GetIndex(tt.args.field); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerRule_GetIndexValue(t *testing.T) {
	type fields struct {
		ID             string
		NameID         string
		AnalyzerNameID string
		Title          string
		Description    string
		Severity       string
		Mitigation     string
		Meta           datatypes.JSONMap
		CreatedAt      time.Time
		UpdatedAt      time.Time
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
			name:   "index name_id",
			fields: fields{NameID: "test"},
			args:   args{field: "name_id"},
			want:   "test",
		},
		{
			name:   "index analyzer_name_id",
			fields: fields{AnalyzerNameID: "security"},
			args:   args{field: "analyzer_name_id"},
			want:   "security",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Rule{
				ID:             tt.fields.ID,
				NameID:         tt.fields.NameID,
				AnalyzerNameID: tt.fields.AnalyzerNameID,
				Title:          tt.fields.Title,
				Description:    tt.fields.Description,
				Severity:       tt.fields.Severity,
				Mitigation:     tt.fields.Mitigation,
				Meta:           tt.fields.Meta,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
			}
			if got := m.GetIndexValue(tt.args.field); got != tt.want {
				t.Errorf("GetIndexValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerRule_GetTags(t *testing.T) {
	type fields struct {
		ID             string
		NameID         string
		AnalyzerNameID string
		Title          string
		Description    string
		Severity       string
		Mitigation     string
		Meta           datatypes.JSONMap
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "normal",
			fields: fields{NameID: "test"},
			want:   []string{"test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Rule{
				ID:             tt.fields.ID,
				NameID:         tt.fields.NameID,
				AnalyzerNameID: tt.fields.AnalyzerNameID,
				Title:          tt.fields.Title,
				Description:    tt.fields.Description,
				Severity:       tt.fields.Severity,
				Mitigation:     tt.fields.Mitigation,
				Meta:           tt.fields.Meta,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
			}
			if got := m.GetTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerRule_Sortable(t *testing.T) {
	type fields struct {
		ID             string
		NameID         string
		AnalyzerNameID string
		Title          string
		Description    string
		Severity       string
		Mitigation     string
		Meta           datatypes.JSONMap
		CreatedAt      time.Time
		UpdatedAt      time.Time
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
			m := &Rule{
				ID:             tt.fields.ID,
				NameID:         tt.fields.NameID,
				AnalyzerNameID: tt.fields.AnalyzerNameID,
				Title:          tt.fields.Title,
				Description:    tt.fields.Description,
				Severity:       tt.fields.Severity,
				Mitigation:     tt.fields.Mitigation,
				Meta:           tt.fields.Meta,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
			}
			if got := m.Sortable(tt.args.field); got != tt.want {
				t.Errorf("Sortable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerRule_String(t *testing.T) {
	type fields struct {
		ID             string
		NameID         string
		AnalyzerNameID string
		Title          string
		Description    string
		Severity       string
		Mitigation     string
		Meta           datatypes.JSONMap
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal",
			fields: fields{NameID: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Rule{
				ID:             tt.fields.ID,
				NameID:         tt.fields.NameID,
				AnalyzerNameID: tt.fields.AnalyzerNameID,
				Title:          tt.fields.Title,
				Description:    tt.fields.Description,
				Severity:       tt.fields.Severity,
				Mitigation:     tt.fields.Mitigation,
				Meta:           tt.fields.Meta,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
			}
			got := m.String()
			assert.NotEmpty(t, got)
		})
	}
}

func TestAnalyzerRule_TableName(t *testing.T) {
	type fields struct {
		ID             string
		NameID         string
		AnalyzerNameID string
		Title          string
		Description    string
		Severity       string
		Mitigation     string
		Meta           datatypes.JSONMap
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal",
			fields: fields{},
			want:   AnalyzerRuleTableName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Rule{
				ID:             tt.fields.ID,
				NameID:         tt.fields.NameID,
				AnalyzerNameID: tt.fields.AnalyzerNameID,
				Title:          tt.fields.Title,
				Description:    tt.fields.Description,
				Severity:       tt.fields.Severity,
				Mitigation:     tt.fields.Mitigation,
				Meta:           tt.fields.Meta,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
			}
			if got := m.TableName(); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
