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

func TestAnalyzer_BeforeCreate(t *testing.T) {
	type fields struct {
		ID          string
		NameID      string
		Title       string
		Description string
		Status      string
		Meta        datatypes.JSONMap
		Config      Config
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Position    int
		Rules       []*Rule
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
			m := &Analyzer{
				ID:          tt.fields.ID,
				NameID:      tt.fields.NameID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Status:      tt.fields.Status,
				Meta:        tt.fields.Meta,
				Config:      tt.fields.Config,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
				Position:    tt.fields.Position,
				Rules:       tt.fields.Rules,
			}
			if err := m.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NotEmpty(t, m.ID)
		})
	}
}

func TestAnalyzer_BeforeSave(t *testing.T) {
	type fields struct {
		ID          string
		NameID      string
		Title       string
		Description string
		Status      string
		Meta        datatypes.JSONMap
		Config      Config
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Position    int
		Rules       []*Rule
	}
	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus string
		wantErr    bool
	}{
		{
			name:       "normal - init status if empty",
			fields:     fields{ID: "1", Status: ""},
			args:       args{},
			wantStatus: AnalyzerStatusActive,
			wantErr:    false,
		},
		{
			name:       "normal - skip initializing status if not empty",
			fields:     fields{ID: "1", Status: AnalyzerStatusActive},
			args:       args{},
			wantStatus: AnalyzerStatusActive,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Analyzer{
				ID:          tt.fields.ID,
				NameID:      tt.fields.NameID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Status:      tt.fields.Status,
				Meta:        tt.fields.Meta,
				Config:      tt.fields.Config,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
				Position:    tt.fields.Position,
				Rules:       tt.fields.Rules,
			}
			if err := m.BeforeSave(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("BeforeSave() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NotEmpty(t, m.Status)
			assert.Equal(t, tt.wantStatus, m.Status)
		})
	}
}

func TestAnalyzer_GetID(t *testing.T) {
	type fields struct {
		ID          string
		NameID      string
		Title       string
		Description string
		Status      string
		Meta        datatypes.JSONMap
		Config      Config
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Position    int
		Rules       []*Rule
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
			m := &Analyzer{
				ID:          tt.fields.ID,
				NameID:      tt.fields.NameID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Status:      tt.fields.Status,
				Meta:        tt.fields.Meta,
				Config:      tt.fields.Config,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
				Position:    tt.fields.Position,
				Rules:       tt.fields.Rules,
			}
			if got := m.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzer_GetIndex(t *testing.T) {
	type fields struct {
		ID          string
		NameID      string
		Title       string
		Description string
		Status      string
		Meta        datatypes.JSONMap
		Config      Config
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Position    int
		Rules       []*Rule
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
			name:   "index status",
			fields: fields{},
			args:   args{field: "status"},
			want:   "idx_status",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Analyzer{
				ID:          tt.fields.ID,
				NameID:      tt.fields.NameID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Status:      tt.fields.Status,
				Meta:        tt.fields.Meta,
				Config:      tt.fields.Config,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
				Position:    tt.fields.Position,
				Rules:       tt.fields.Rules,
			}
			if got := m.GetIndex(tt.args.field); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzer_GetIndexValue(t *testing.T) {
	type fields struct {
		ID          string
		NameID      string
		Title       string
		Description string
		Status      string
		Meta        datatypes.JSONMap
		Config      Config
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Position    int
		Rules       []*Rule
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
			name:   "index status",
			fields: fields{Status: "active"},
			args:   args{field: "status"},
			want:   "active",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Analyzer{
				ID:          tt.fields.ID,
				NameID:      tt.fields.NameID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Status:      tt.fields.Status,
				Meta:        tt.fields.Meta,
				Config:      tt.fields.Config,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
				Position:    tt.fields.Position,
				Rules:       tt.fields.Rules,
			}
			if got := m.GetIndexValue(tt.args.field); got != tt.want {
				t.Errorf("GetIndexValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzer_GetTags(t *testing.T) {
	type fields struct {
		ID          string
		NameID      string
		Title       string
		Description string
		Status      string
		Meta        datatypes.JSONMap
		Config      Config
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Position    int
		Rules       []*Rule
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
			m := &Analyzer{
				ID:          tt.fields.ID,
				NameID:      tt.fields.NameID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Status:      tt.fields.Status,
				Meta:        tt.fields.Meta,
				Config:      tt.fields.Config,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
				Position:    tt.fields.Position,
				Rules:       tt.fields.Rules,
			}
			if got := m.GetTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzer_Sortable(t *testing.T) {
	type fields struct {
		ID          string
		NameID      string
		Title       string
		Description string
		Status      string
		Meta        datatypes.JSONMap
		Config      Config
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Position    int
		Rules       []*Rule
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
			m := &Analyzer{
				ID:          tt.fields.ID,
				NameID:      tt.fields.NameID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Status:      tt.fields.Status,
				Meta:        tt.fields.Meta,
				Config:      tt.fields.Config,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
				Position:    tt.fields.Position,
				Rules:       tt.fields.Rules,
			}
			if got := m.Sortable(tt.args.field); got != tt.want {
				t.Errorf("Sortable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzer_String(t *testing.T) {
	type fields struct {
		ID          string
		NameID      string
		Title       string
		Description string
		Status      string
		Meta        datatypes.JSONMap
		Config      Config
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Position    int
		Rules       []*Rule
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
			m := &Analyzer{
				ID:          tt.fields.ID,
				NameID:      tt.fields.NameID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Status:      tt.fields.Status,
				Meta:        tt.fields.Meta,
				Config:      tt.fields.Config,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
				Position:    tt.fields.Position,
				Rules:       tt.fields.Rules,
			}
			got := m.String()
			assert.NotEmpty(t, got)
		})
	}
}

func TestAnalyzer_TableName(t *testing.T) {
	type fields struct {
		ID          string
		NameID      string
		Title       string
		Description string
		Status      string
		Meta        datatypes.JSONMap
		Config      Config
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Position    int
		Rules       []*Rule
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal",
			fields: fields{},
			want:   AnalyzerTableName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Analyzer{
				ID:          tt.fields.ID,
				NameID:      tt.fields.NameID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Status:      tt.fields.Status,
				Meta:        tt.fields.Meta,
				Config:      tt.fields.Config,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
				Position:    tt.fields.Position,
				Rules:       tt.fields.Rules,
			}
			if got := m.TableName(); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
