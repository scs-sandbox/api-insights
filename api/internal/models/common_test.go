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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnalyzerConfigMap_Merge(t *testing.T) {
	type args struct {
		with AnalyzerConfigMap
	}
	tests := []struct {
		name string
		m    AnalyzerConfigMap
		args args
		want AnalyzerConfigMap
	}{
		{
			name: "normal",
			m:    map[analyzer.SpecAnalyzer]analyzer.Config{},
			args: args{with: map[analyzer.SpecAnalyzer]analyzer.Config{
				analyzer.Completeness: analyzer.Config{
					analyzer.ConfigServiceNameID: "test",
				},
			}},
			want: map[analyzer.SpecAnalyzer]analyzer.Config{
				analyzer.Completeness: analyzer.Config{
					analyzer.ConfigServiceNameID: "test",
				},
			},
		},
		{
			name: "normal",
			m: map[analyzer.SpecAnalyzer]analyzer.Config{
				analyzer.Completeness: analyzer.Config{
					analyzer.ConfigServiceNameID: "test",
				}},
			args: args{with: map[analyzer.SpecAnalyzer]analyzer.Config{
				analyzer.Completeness: analyzer.Config{
					"new": "test",
				},
			}},
			want: map[analyzer.SpecAnalyzer]analyzer.Config{
				analyzer.Completeness: analyzer.Config{
					analyzer.ConfigServiceNameID: "test",
					"new":                        "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Merge(tt.args.with)
			assert.Equal(t, tt.want, tt.m)
		})
	}
}

func TestAnalyzerConfigMap_Scan(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		m       AnalyzerConfigMap
		args    args
		wantErr bool
	}{
		{
			name:    "normal",
			m:       map[analyzer.SpecAnalyzer]analyzer.Config{},
			args:    args{value: []byte(`{"completeness":{"service_name_id": "test"}}`)},
			wantErr: false,
		},
		{
			name:    "invalid",
			m:       map[analyzer.SpecAnalyzer]analyzer.Config{},
			args:    args{value: "invalid data type"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAnalyzerConfigMap_Value(t *testing.T) {
	tests := []struct {
		name    string
		m       AnalyzerConfigMap
		want    driver.Value
		wantErr bool
	}{
		{
			name: "normal",
			m: map[analyzer.SpecAnalyzer]analyzer.Config{
				analyzer.Completeness: analyzer.Config{
					analyzer.ConfigServiceNameID: "test",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.m.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
		})
	}
}
