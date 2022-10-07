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
	"database/sql/driver"
	"reflect"
	"testing"
)

func TestConfig_GetScoreConfig(t *testing.T) {
	tests := []struct {
		name string
		c    Config
		want *ScoreConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetScoreConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetScoreConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Scan(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		c       Config
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_ServiceNameID(t *testing.T) {
	tests := []struct {
		name string
		c    Config
		want string
	}{
		{
			name: "normal",
			c:    Config{ConfigServiceNameID: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ServiceNameID(); got != tt.want {
				t.Errorf("ServiceNameID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_ServiceNameIDFromTemplate(t *testing.T) {
	type args struct {
		serviceNameID string
	}
	tests := []struct {
		name string
		c    Config
		args args
		want string
	}{
		{
			name: "normal",
			c:    Config{ConfigServiceNameIDTemplate: "{{ .nameID }}.api.apiregistry"},
			args: args{serviceNameID: "test"},
			want: "test.api.apiregistry",
		},
		{
			name: "empty template",
			c:    Config{ConfigServiceNameIDTemplate: ""},
			args: args{serviceNameID: "test"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ServiceNameIDFromTemplate(tt.args.serviceNameID); got != tt.want {
				t.Errorf("ServiceNameIDFromTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_UnmarshalInto(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		c       Config
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnmarshalInto(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalInto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_Value(t *testing.T) {
	tests := []struct {
		name    string
		c       Config
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListToMap(t *testing.T) {
	type args struct {
		list []*Analyzer
	}
	tests := []struct {
		name string
		args args
		want map[SpecAnalyzer]*Analyzer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListToMap(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAnalyzersScoreConfigsFrom(t *testing.T) {
	type args struct {
		analyzers map[SpecAnalyzer]*Analyzer
	}
	tests := []struct {
		name    string
		args    args
		want    AnalyzersScoreConfigs
		wantErr bool
	}{
		//{
		//	name:    "",
		//	args:    args{},
		//	want:    nil,
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAnalyzersScoreConfigsFrom(tt.args.analyzers)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAnalyzersScoreConfigsFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnalyzersScoreConfigsFrom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewScoreConfig(t *testing.T) {
	type args struct {
		setDefaults bool
	}
	tests := []struct {
		name string
		args args
		want *ScoreConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScoreConfig(tt.args.setDefaults); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScoreConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
