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
	"github.com/cisco-developer/api-insights/api/internal/models/diff"
	"github.com/emicklei/go-restful/v3"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestSpecDiffRequest_Compare(t *testing.T) {
	type fields struct {
		NewSpecID      string
		OldSpecID      string
		OldSpecDoc     SpecDoc
		NewSpecDoc     SpecDoc
		SpecDiffConfig SpecDiffConfig
	}
	type args struct {
		with *SpecDiffRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpecDiffRequest{
				NewSpecID:      tt.fields.NewSpecID,
				OldSpecID:      tt.fields.OldSpecID,
				OldSpecDoc:     tt.fields.OldSpecDoc,
				NewSpecDoc:     tt.fields.NewSpecDoc,
				SpecDiffConfig: tt.fields.SpecDiffConfig,
			}
			if got := r.Compare(tt.args.with); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: update
func TestSpecDiffRequest_From(t *testing.T) {
	type fields struct {
		NewSpecID      string
		OldSpecID      string
		OldSpecDoc     SpecDoc
		NewSpecDoc     SpecDoc
		SpecDiffConfig SpecDiffConfig
	}
	type args struct {
		req         *restful.Request
		specsGetter func(oldSpecID, newSpecID string) (oldSpec *Spec, newSpec *Spec, err error)
	}
	//req := &restful.Request{Request: &http.Request{
	//	MultipartForm: nil,
	//	Method:        http.MethodPost,
	//	Header:        map[string][]string{"Content-Type": {"application/json"}},
	//	Body: ,
	//}}

	req := &restful.Request{
		Request: httptest.NewRequest(http.MethodPost, "/diff", strings.NewReader(`
			{
				"old_spec_doc": "xxx",
				"new_spec_doc": "xxx",
				"config": {
					"output_format": "markdown"
				}
			}`)),
	}
	req.Request.MultipartForm = nil
	req.Request.Header = map[string][]string{"Content-Type": {"application/json"}}
	//req.Request = httptest.NewRequest(http.MethodPost, "/diff", strings.NewReader(`{"name":"Robert"}`))

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal - non multipart form",
			fields: fields{
				NewSpecID:      "",
				OldSpecID:      "",
				OldSpecDoc:     SpecDoc(loadSpec("testdata/petstore-v2.json")),
				NewSpecDoc:     SpecDoc(loadSpec("testdata/petstore-v2.json")),
				SpecDiffConfig: SpecDiffConfig{},
			},
			args: args{
				req:         req,
				specsGetter: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpecDiffRequest{
				NewSpecID:      tt.fields.NewSpecID,
				OldSpecID:      tt.fields.OldSpecID,
				OldSpecDoc:     tt.fields.OldSpecDoc,
				NewSpecDoc:     tt.fields.NewSpecDoc,
				SpecDiffConfig: tt.fields.SpecDiffConfig,
			}
			if err := r.From(tt.args.req, tt.args.specsGetter); (err != nil) != tt.wantErr {
				t.Errorf("From() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpecDiffRequest_HasSpecDocs(t *testing.T) {
	type fields struct {
		NewSpecID      string
		OldSpecID      string
		OldSpecDoc     SpecDoc
		NewSpecDoc     SpecDoc
		SpecDiffConfig SpecDiffConfig
	}
	tests := []struct {
		name              string
		fields            fields
		wantHasOldSpecDoc bool
		wantHasNewSpecDoc bool
	}{
		{
			name: "provided both",
			fields: fields{
				OldSpecDoc: SpecDoc(loadSpec("testdata/petstore-v2.json")),
				NewSpecDoc: SpecDoc(loadSpec("testdata/petstore-v2.json")),
			},
			wantHasOldSpecDoc: true,
			wantHasNewSpecDoc: true,
		},
		{
			name: "provided old spec only",
			fields: fields{
				OldSpecDoc: SpecDoc(loadSpec("testdata/petstore-v2.json")),
				NewSpecDoc: nil,
			},
			wantHasOldSpecDoc: true,
			wantHasNewSpecDoc: false,
		},
		{
			name: "provided new spec only",
			fields: fields{
				OldSpecDoc: nil,
				NewSpecDoc: SpecDoc(loadSpec("testdata/petstore-v2.json")),
			},
			wantHasOldSpecDoc: false,
			wantHasNewSpecDoc: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpecDiffRequest{
				NewSpecID:      tt.fields.NewSpecID,
				OldSpecID:      tt.fields.OldSpecID,
				OldSpecDoc:     tt.fields.OldSpecDoc,
				NewSpecDoc:     tt.fields.NewSpecDoc,
				SpecDiffConfig: tt.fields.SpecDiffConfig,
			}
			gotHasOldSpecDoc, gotHasNewSpecDoc := r.HasSpecDocs()
			if gotHasOldSpecDoc != tt.wantHasOldSpecDoc {
				t.Errorf("HasSpecDocs() gotHasOldSpecDoc = %v, want %v", gotHasOldSpecDoc, tt.wantHasOldSpecDoc)
			}
			if gotHasNewSpecDoc != tt.wantHasNewSpecDoc {
				t.Errorf("HasSpecDocs() gotHasNewSpecDoc = %v, want %v", gotHasNewSpecDoc, tt.wantHasNewSpecDoc)
			}
		})
	}
}

func TestSpecDiffRequest_tryAsMultipartForm(t *testing.T) {
	type fields struct {
		NewSpecID      string
		OldSpecID      string
		OldSpecDoc     SpecDoc
		NewSpecDoc     SpecDoc
		SpecDiffConfig SpecDiffConfig
	}
	type args struct {
		req *restful.Request
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantErr         error
		wantIsMultipart bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpecDiffRequest{
				NewSpecID:      tt.fields.NewSpecID,
				OldSpecID:      tt.fields.OldSpecID,
				OldSpecDoc:     tt.fields.OldSpecDoc,
				NewSpecDoc:     tt.fields.NewSpecDoc,
				SpecDiffConfig: tt.fields.SpecDiffConfig,
			}
			gotIsMultipart, gotErr := r.tryAsMultipartForm(tt.args.req)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("tryAsMultipartForm() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			if gotIsMultipart != tt.wantIsMultipart {
				t.Errorf("tryAsMultipartForm() gotIsMultipart = %v, want %v", gotIsMultipart, tt.wantIsMultipart)
			}
		})
	}
}

func TestSpecDiff_AfterFind(t *testing.T) {
	type fields struct {
		ID              string
		SpecDiffRequest *SpecDiffRequest
		SpecDiffResult  SpecDiffResult
		ServiceID       string
		Status          string
		CreatedAt       time.Time
		UpdatedAt       time.Time
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
				SpecDiffRequest: &SpecDiffRequest{
					SpecDiffConfig: SpecDiffConfig{
						Config:    nil,
						RawConfig: []byte(`{"output_format": "markdown"}`),
					},
				},
				SpecDiffResult: SpecDiffResult{
					Result:    nil,
					RawResult: []byte(`{"markdown": "### What's New"}`),
				},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "nil raw config",
			fields: fields{
				SpecDiffRequest: &SpecDiffRequest{
					SpecDiffConfig: SpecDiffConfig{
						Config:    nil,
						RawConfig: nil,
					},
				},
				SpecDiffResult: SpecDiffResult{},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "nil raw result",
			fields: fields{
				SpecDiffRequest: &SpecDiffRequest{
					SpecDiffConfig: SpecDiffConfig{
						Config:    nil,
						RawConfig: nil,
					},
				},
				SpecDiffResult: SpecDiffResult{
					Result:    nil,
					RawResult: nil,
				},
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecDiff{
				ID:              tt.fields.ID,
				SpecDiffRequest: tt.fields.SpecDiffRequest,
				SpecDiffResult:  tt.fields.SpecDiffResult,
				ServiceID:       tt.fields.ServiceID,
				Status:          tt.fields.Status,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			if err := m.AfterFind(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("AfterFind() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpecDiff_BeforeCreate(t *testing.T) {
	type fields struct {
		ID              string
		SpecDiffRequest *SpecDiffRequest
		SpecDiffResult  SpecDiffResult
		ServiceID       string
		Status          string
		CreatedAt       time.Time
		UpdatedAt       time.Time
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
			m := &SpecDiff{
				ID:              tt.fields.ID,
				SpecDiffRequest: tt.fields.SpecDiffRequest,
				SpecDiffResult:  tt.fields.SpecDiffResult,
				ServiceID:       tt.fields.ServiceID,
				Status:          tt.fields.Status,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			if err := m.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NotEmpty(t, m.ID)
		})
	}
}

func TestSpecDiff_BeforeSave(t *testing.T) {
	type fields struct {
		ID              string
		SpecDiffRequest *SpecDiffRequest
		SpecDiffResult  SpecDiffResult
		ServiceID       string
		Status          string
		CreatedAt       time.Time
		UpdatedAt       time.Time
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
				SpecDiffRequest: &SpecDiffRequest{
					SpecDiffConfig: SpecDiffConfig{
						Config:    &diff.Config{OutputFormat: "markdown"},
						RawConfig: nil,
					},
				},
				SpecDiffResult: SpecDiffResult{
					Result: &diff.Result{
						Markdown: "### What's New",
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
				SpecDiffRequest: &SpecDiffRequest{
					SpecDiffConfig: SpecDiffConfig{
						Config:    nil,
						RawConfig: nil,
					},
				},
				SpecDiffResult: SpecDiffResult{
					Result: &diff.Result{
						Markdown: "### What's New",
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
				SpecDiffRequest: &SpecDiffRequest{
					SpecDiffConfig: SpecDiffConfig{
						Config:    &diff.Config{OutputFormat: "markdown"},
						RawConfig: nil,
					},
				},
				SpecDiffResult: SpecDiffResult{
					Result:    nil,
					RawResult: nil,
				},
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecDiff{
				ID:              tt.fields.ID,
				SpecDiffRequest: tt.fields.SpecDiffRequest,
				SpecDiffResult:  tt.fields.SpecDiffResult,
				ServiceID:       tt.fields.ServiceID,
				Status:          tt.fields.Status,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			if err := m.BeforeSave(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("BeforeSave() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpecDiff_GetID(t *testing.T) {
	type fields struct {
		ID              string
		SpecDiffRequest *SpecDiffRequest
		SpecDiffResult  SpecDiffResult
		ServiceID       string
		Status          string
		CreatedAt       time.Time
		UpdatedAt       time.Time
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
			m := &SpecDiff{
				ID:              tt.fields.ID,
				SpecDiffRequest: tt.fields.SpecDiffRequest,
				SpecDiffResult:  tt.fields.SpecDiffResult,
				ServiceID:       tt.fields.ServiceID,
				Status:          tt.fields.Status,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			if got := m.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecDiff_GetIndex(t *testing.T) {
	type fields struct {
		ID              string
		SpecDiffRequest *SpecDiffRequest
		SpecDiffResult  SpecDiffResult
		ServiceID       string
		Status          string
		CreatedAt       time.Time
		UpdatedAt       time.Time
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
			name:   "index old_spec_id",
			fields: fields{},
			args:   args{field: "old_spec_id"},
			want:   "idx_old_spec_id",
		},
		{
			name:   "index new_spec_id",
			fields: fields{},
			args:   args{field: "new_spec_id"},
			want:   "idx_new_spec_id",
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
			m := &SpecDiff{
				ID:              tt.fields.ID,
				SpecDiffRequest: tt.fields.SpecDiffRequest,
				SpecDiffResult:  tt.fields.SpecDiffResult,
				ServiceID:       tt.fields.ServiceID,
				Status:          tt.fields.Status,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			if got := m.GetIndex(tt.args.field); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecDiff_GetIndexValue(t *testing.T) {
	type fields struct {
		ID              string
		SpecDiffRequest *SpecDiffRequest
		SpecDiffResult  SpecDiffResult
		ServiceID       string
		Status          string
		CreatedAt       time.Time
		UpdatedAt       time.Time
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
			fields: fields{ServiceID: "test", SpecDiffRequest: &SpecDiffRequest{OldSpecID: "test"}},
			args:   args{field: "service_id"},
			want:   "test",
		},
		{
			name:   "index old_spec_id",
			fields: fields{SpecDiffRequest: &SpecDiffRequest{OldSpecID: "test"}},
			args:   args{field: "old_spec_id"},
			want:   "test",
		},
		{
			name:   "index new_spec_id",
			fields: fields{SpecDiffRequest: &SpecDiffRequest{NewSpecID: "test"}},
			args:   args{field: "new_spec_id"},
			want:   "test",
		},
		{
			name:   "index status",
			fields: fields{Status: "Development", SpecDiffRequest: &SpecDiffRequest{OldSpecID: "test"}},
			args:   args{field: "status"},
			want:   "Development",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecDiff{
				ID:              tt.fields.ID,
				SpecDiffRequest: tt.fields.SpecDiffRequest,
				SpecDiffResult:  tt.fields.SpecDiffResult,
				ServiceID:       tt.fields.ServiceID,
				Status:          tt.fields.Status,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			if got := m.GetIndexValue(tt.args.field); got != tt.want {
				t.Errorf("GetIndexValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecDiff_GetTags(t *testing.T) {
	type fields struct {
		ID              string
		SpecDiffRequest *SpecDiffRequest
		SpecDiffResult  SpecDiffResult
		ServiceID       string
		Status          string
		CreatedAt       time.Time
		UpdatedAt       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "normal",
			fields: fields{
				ServiceID:       "catalogue",
				SpecDiffRequest: &SpecDiffRequest{OldSpecID: "1", NewSpecID: "2"},
				Status:          "Development",
			},
			want: []string{"catalogue", "1", "2", "Development"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecDiff{
				ID:              tt.fields.ID,
				SpecDiffRequest: tt.fields.SpecDiffRequest,
				SpecDiffResult:  tt.fields.SpecDiffResult,
				ServiceID:       tt.fields.ServiceID,
				Status:          tt.fields.Status,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			if got := m.GetTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecDiff_SetResult(t *testing.T) {
	type fields struct {
		ID              string
		SpecDiffRequest *SpecDiffRequest
		SpecDiffResult  SpecDiffResult
		ServiceID       string
		Status          string
		CreatedAt       time.Time
		UpdatedAt       time.Time
	}
	type args struct {
		result *diff.Result
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecDiff{
				ID:              tt.fields.ID,
				SpecDiffRequest: tt.fields.SpecDiffRequest,
				SpecDiffResult:  tt.fields.SpecDiffResult,
				ServiceID:       tt.fields.ServiceID,
				Status:          tt.fields.Status,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			if err := m.SetResult(tt.args.result, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("SetResult() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpecDiff_Sortable(t *testing.T) {
	type fields struct {
		ID              string
		SpecDiffRequest *SpecDiffRequest
		SpecDiffResult  SpecDiffResult
		ServiceID       string
		Status          string
		CreatedAt       time.Time
		UpdatedAt       time.Time
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
			m := &SpecDiff{
				ID:              tt.fields.ID,
				SpecDiffRequest: tt.fields.SpecDiffRequest,
				SpecDiffResult:  tt.fields.SpecDiffResult,
				ServiceID:       tt.fields.ServiceID,
				Status:          tt.fields.Status,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			if got := m.Sortable(tt.args.field); got != tt.want {
				t.Errorf("Sortable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecDiff_String(t *testing.T) {
	type fields struct {
		ID              string
		SpecDiffRequest *SpecDiffRequest
		SpecDiffResult  SpecDiffResult
		ServiceID       string
		Status          string
		CreatedAt       time.Time
		UpdatedAt       time.Time
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
			m := &SpecDiff{
				ID:              tt.fields.ID,
				SpecDiffRequest: tt.fields.SpecDiffRequest,
				SpecDiffResult:  tt.fields.SpecDiffResult,
				ServiceID:       tt.fields.ServiceID,
				Status:          tt.fields.Status,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			got := m.String()
			assert.NotEmpty(t, got)
		})
	}
}

func TestSpecDiff_TableName(t *testing.T) {
	type fields struct {
		ID              string
		SpecDiffRequest *SpecDiffRequest
		SpecDiffResult  SpecDiffResult
		ServiceID       string
		Status          string
		CreatedAt       time.Time
		UpdatedAt       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal",
			fields: fields{},
			want:   SpecDiffTableName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SpecDiff{
				ID:              tt.fields.ID,
				SpecDiffRequest: tt.fields.SpecDiffRequest,
				SpecDiffResult:  tt.fields.SpecDiffResult,
				ServiceID:       tt.fields.ServiceID,
				Status:          tt.fields.Status,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			if got := m.TableName(); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
