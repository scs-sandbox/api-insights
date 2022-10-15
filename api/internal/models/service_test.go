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
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

func TestContact_Scan(t *testing.T) {
	type fields struct {
		Contact openapi3.Contact
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "normal",
			fields:  fields{Contact: openapi3.Contact{Name: "test"}},
			args:    args{value: []byte(`{"name": "test"}`)},
			wantErr: false,
		},
		{
			name:    "invalid",
			fields:  fields{Contact: openapi3.Contact{Name: "test"}},
			args:    args{value: "invalid data type"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Contact{
				Contact: tt.fields.Contact,
			}
			if err := m.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContact_Value(t *testing.T) {
	type fields struct {
		Contact openapi3.Contact
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Contact: openapi3.Contact{Name: "test"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Contact{
				Contact: tt.fields.Contact,
			}
			got, err := m.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
		})
	}
}

func TestServiceSummary_Scan(t *testing.T) {
	type fields struct {
		Score     *int
		Version   string
		Revision  string
		UpdatedAt time.Time
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "normal",
			fields:  fields{},
			args:    args{value: []byte(`{"score": 80, "version": "v1.0", "revision": "1"}`)},
			wantErr: false,
		},
		{
			name:    "invalid",
			fields:  fields{},
			args:    args{value: "invalid data type"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ServiceSummary{
				Score:     tt.fields.Score,
				Version:   tt.fields.Version,
				Revision:  tt.fields.Revision,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := m.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServiceSummary_Value(t *testing.T) {
	type fields struct {
		Score     *int
		Version   string
		Revision  string
		UpdatedAt time.Time
	}
	score := 80
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Score:     &score,
				Version:   "v1.0",
				Revision:  "1",
				UpdatedAt: time.Time{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := ServiceSummary{
				Score:     tt.fields.Score,
				Version:   tt.fields.Version,
				Revision:  tt.fields.Revision,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			got, err := m.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
		})
	}
}

func TestService_BeforeCreate(t *testing.T) {
	type fields struct {
		ID               string
		AdditionalInfo   datatypes.JSONMap
		Contact          *Contact
		Description      string
		NameID           string
		OrganizationID   string
		ProductTag       string
		Title            string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		AnalyzersConfigs AnalyzerConfigMap
		Summary          *ServiceSummary
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
			m := &Service{
				ID:               tt.fields.ID,
				AdditionalInfo:   tt.fields.AdditionalInfo,
				Contact:          tt.fields.Contact,
				Description:      tt.fields.Description,
				NameID:           tt.fields.NameID,
				OrganizationID:   tt.fields.OrganizationID,
				ProductTag:       tt.fields.ProductTag,
				Title:            tt.fields.Title,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				AnalyzersConfigs: tt.fields.AnalyzersConfigs,
				Summary:          tt.fields.Summary,
			}
			if err := m.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NotEmpty(t, m.ID)
		})
	}
}

func TestService_GetID(t *testing.T) {
	type fields struct {
		ID               string
		AdditionalInfo   datatypes.JSONMap
		Contact          *Contact
		Description      string
		NameID           string
		OrganizationID   string
		ProductTag       string
		Title            string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		AnalyzersConfigs AnalyzerConfigMap
		Summary          *ServiceSummary
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
			m := &Service{
				ID:               tt.fields.ID,
				AdditionalInfo:   tt.fields.AdditionalInfo,
				Contact:          tt.fields.Contact,
				Description:      tt.fields.Description,
				NameID:           tt.fields.NameID,
				OrganizationID:   tt.fields.OrganizationID,
				ProductTag:       tt.fields.ProductTag,
				Title:            tt.fields.Title,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				AnalyzersConfigs: tt.fields.AnalyzersConfigs,
				Summary:          tt.fields.Summary,
			}
			if got := m.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetIndex(t *testing.T) {
	type fields struct {
		ID               string
		AdditionalInfo   datatypes.JSONMap
		Contact          *Contact
		Description      string
		NameID           string
		OrganizationID   string
		ProductTag       string
		Title            string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		AnalyzersConfigs AnalyzerConfigMap
		Summary          *ServiceSummary
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
			name:   "index product_tag",
			fields: fields{},
			args:   args{field: "product_tag"},
			want:   "idx_product_tag",
		},
		{
			name:   "index organization_id",
			fields: fields{},
			args:   args{field: "organization_id"},
			want:   "idx_organization_id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Service{
				ID:               tt.fields.ID,
				AdditionalInfo:   tt.fields.AdditionalInfo,
				Contact:          tt.fields.Contact,
				Description:      tt.fields.Description,
				NameID:           tt.fields.NameID,
				OrganizationID:   tt.fields.OrganizationID,
				ProductTag:       tt.fields.ProductTag,
				Title:            tt.fields.Title,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				AnalyzersConfigs: tt.fields.AnalyzersConfigs,
				Summary:          tt.fields.Summary,
			}
			if got := m.GetIndex(tt.args.field); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetIndexValue(t *testing.T) {
	type fields struct {
		ID               string
		AdditionalInfo   datatypes.JSONMap
		Contact          *Contact
		Description      string
		NameID           string
		OrganizationID   string
		ProductTag       string
		Title            string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		AnalyzersConfigs AnalyzerConfigMap
		Summary          *ServiceSummary
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
			name:   "index product_tag",
			fields: fields{ProductTag: "test"},
			args:   args{field: "product_tag"},
			want:   "test",
		},
		{
			name:   "index organization_id",
			fields: fields{OrganizationID: "test"},
			args:   args{field: "organization_id"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Service{
				ID:               tt.fields.ID,
				AdditionalInfo:   tt.fields.AdditionalInfo,
				Contact:          tt.fields.Contact,
				Description:      tt.fields.Description,
				NameID:           tt.fields.NameID,
				OrganizationID:   tt.fields.OrganizationID,
				ProductTag:       tt.fields.ProductTag,
				Title:            tt.fields.Title,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				AnalyzersConfigs: tt.fields.AnalyzersConfigs,
				Summary:          tt.fields.Summary,
			}
			if got := m.GetIndexValue(tt.args.field); got != tt.want {
				t.Errorf("GetIndexValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetNameID(t *testing.T) {
	type fields struct {
		ID               string
		AdditionalInfo   datatypes.JSONMap
		Contact          *Contact
		Description      string
		NameID           string
		OrganizationID   string
		ProductTag       string
		Title            string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		AnalyzersConfigs AnalyzerConfigMap
		Summary          *ServiceSummary
	}
	type args struct {
		name        analyzer.SpecAnalyzer
		analyzerCfg analyzer.Config
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "nameID from analyzersConfigs",
			fields: fields{
				NameID: "test",
				AnalyzersConfigs: map[analyzer.SpecAnalyzer]analyzer.Config{
					analyzer.Drift: analyzer.Config{
						analyzer.ConfigServiceNameID: "cart.sock-shop",
					},
				},
			},
			args: args{
				name:        analyzer.Drift,
				analyzerCfg: nil,
			},
			want: "cart.sock-shop",
		}, {
			name:   "nameID from analyzerCfg",
			fields: fields{NameID: "test"},
			args: args{
				name:        "",
				analyzerCfg: analyzer.Config{analyzer.ConfigServiceNameIDTemplate: "{{ .nameID }}.api.apiregistry"},
			},
			want: "test.api.apiregistry",
		}, {
			name:   "normal",
			fields: fields{NameID: "test"},
			args: args{
				name:        "",
				analyzerCfg: nil,
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Service{
				ID:               tt.fields.ID,
				AdditionalInfo:   tt.fields.AdditionalInfo,
				Contact:          tt.fields.Contact,
				Description:      tt.fields.Description,
				NameID:           tt.fields.NameID,
				OrganizationID:   tt.fields.OrganizationID,
				ProductTag:       tt.fields.ProductTag,
				Title:            tt.fields.Title,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				AnalyzersConfigs: tt.fields.AnalyzersConfigs,
				Summary:          tt.fields.Summary,
			}
			if got := m.GetNameID(tt.args.name, tt.args.analyzerCfg); got != tt.want {
				t.Errorf("GetNameID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetTags(t *testing.T) {
	type fields struct {
		ID               string
		AdditionalInfo   datatypes.JSONMap
		Contact          *Contact
		Description      string
		NameID           string
		OrganizationID   string
		ProductTag       string
		Title            string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		AnalyzersConfigs AnalyzerConfigMap
		Summary          *ServiceSummary
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "normal",
			fields: fields{
				NameID:         "catalogue",
				ProductTag:     "playground",
				OrganizationID: "test",
			},
			want: []string{"catalogue", "playground", "test", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Service{
				ID:               tt.fields.ID,
				AdditionalInfo:   tt.fields.AdditionalInfo,
				Contact:          tt.fields.Contact,
				Description:      tt.fields.Description,
				NameID:           tt.fields.NameID,
				OrganizationID:   tt.fields.OrganizationID,
				ProductTag:       tt.fields.ProductTag,
				Title:            tt.fields.Title,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				AnalyzersConfigs: tt.fields.AnalyzersConfigs,
				Summary:          tt.fields.Summary,
			}
			if got := m.GetTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_SetSummary(t *testing.T) {
	type fields struct {
		ID               string
		AdditionalInfo   datatypes.JSONMap
		Contact          *Contact
		Description      string
		NameID           string
		OrganizationID   string
		ProductTag       string
		Title            string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		AnalyzersConfigs AnalyzerConfigMap
		Summary          *ServiceSummary
	}
	type args struct {
		score     int
		version   string
		revision  string
		updatedAt time.Time
	}
	now := time.Now()
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "normal",
			fields: fields{Summary: nil},
			args: args{
				score:     80,
				version:   "v1.0",
				revision:  "1",
				updatedAt: now,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Service{
				ID:               tt.fields.ID,
				AdditionalInfo:   tt.fields.AdditionalInfo,
				Contact:          tt.fields.Contact,
				Description:      tt.fields.Description,
				NameID:           tt.fields.NameID,
				OrganizationID:   tt.fields.OrganizationID,
				ProductTag:       tt.fields.ProductTag,
				Title:            tt.fields.Title,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				AnalyzersConfigs: tt.fields.AnalyzersConfigs,
				Summary:          tt.fields.Summary,
			}
			m.SetSummary(tt.args.score, tt.args.version, tt.args.revision, tt.args.updatedAt)
			assert.NotNil(t, m.Summary)
			assert.NotNil(t, m.Summary.Score)
			assert.Equal(t, tt.args.score, *m.Summary.Score)
			assert.Equal(t, tt.args.version, m.Summary.Version)
			assert.Equal(t, tt.args.revision, m.Summary.Revision)
			assert.Equal(t, tt.args.updatedAt, m.Summary.UpdatedAt)
		})
	}
}

func TestService_Sortable(t *testing.T) {
	type fields struct {
		ID               string
		AdditionalInfo   datatypes.JSONMap
		Contact          *Contact
		Description      string
		NameID           string
		OrganizationID   string
		ProductTag       string
		Title            string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		AnalyzersConfigs AnalyzerConfigMap
		Summary          *ServiceSummary
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
			m := &Service{
				ID:               tt.fields.ID,
				AdditionalInfo:   tt.fields.AdditionalInfo,
				Contact:          tt.fields.Contact,
				Description:      tt.fields.Description,
				NameID:           tt.fields.NameID,
				OrganizationID:   tt.fields.OrganizationID,
				ProductTag:       tt.fields.ProductTag,
				Title:            tt.fields.Title,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				AnalyzersConfigs: tt.fields.AnalyzersConfigs,
				Summary:          tt.fields.Summary,
			}
			if got := m.Sortable(tt.args.field); got != tt.want {
				t.Errorf("Sortable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_String(t *testing.T) {
	type fields struct {
		ID               string
		AdditionalInfo   datatypes.JSONMap
		Contact          *Contact
		Description      string
		NameID           string
		OrganizationID   string
		ProductTag       string
		Title            string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		AnalyzersConfigs AnalyzerConfigMap
		Summary          *ServiceSummary
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
			m := &Service{
				ID:               tt.fields.ID,
				AdditionalInfo:   tt.fields.AdditionalInfo,
				Contact:          tt.fields.Contact,
				Description:      tt.fields.Description,
				NameID:           tt.fields.NameID,
				OrganizationID:   tt.fields.OrganizationID,
				ProductTag:       tt.fields.ProductTag,
				Title:            tt.fields.Title,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				AnalyzersConfigs: tt.fields.AnalyzersConfigs,
				Summary:          tt.fields.Summary,
			}
			got := m.String()
			assert.NotEmpty(t, got)
		})
	}
}

func TestService_TableName(t *testing.T) {
	type fields struct {
		ID               string
		AdditionalInfo   datatypes.JSONMap
		Contact          *Contact
		Description      string
		NameID           string
		OrganizationID   string
		ProductTag       string
		Title            string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		AnalyzersConfigs AnalyzerConfigMap
		Summary          *ServiceSummary
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal",
			fields: fields{},
			want:   ServiceTableName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Service{
				ID:               tt.fields.ID,
				AdditionalInfo:   tt.fields.AdditionalInfo,
				Contact:          tt.fields.Contact,
				Description:      tt.fields.Description,
				NameID:           tt.fields.NameID,
				OrganizationID:   tt.fields.OrganizationID,
				ProductTag:       tt.fields.ProductTag,
				Title:            tt.fields.Title,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				AnalyzersConfigs: tt.fields.AnalyzersConfigs,
				Summary:          tt.fields.Summary,
			}
			if got := m.TableName(); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
