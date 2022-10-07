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
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer/rule"
	"reflect"
	"testing"
)

func TestWokeConfig_SetDefaults(t *testing.T) {
	type fields struct {
		Config              string
		DisableDefaultRules bool
		ExitOneOnFailure    bool
		NoIgnore            bool
		OutputName          string
	}
	tests := []struct {
		name           string
		fields         fields
		wantOutputName string
	}{
		{
			name:           "empty",
			fields:         fields{OutputName: ""},
			wantOutputName: "json",
		},
		{
			name:           "json",
			fields:         fields{OutputName: "json"},
			wantOutputName: "json",
		},
		{
			name:           "text",
			fields:         fields{OutputName: "text"},
			wantOutputName: "text",
		},
		{
			name:           "simple",
			fields:         fields{OutputName: "simple"},
			wantOutputName: "simple",
		},
		{
			name:           "github-actions",
			fields:         fields{OutputName: "github-actions"},
			wantOutputName: "github-actions",
		},
		{
			name:           "sonarqube",
			fields:         fields{OutputName: "sonarqube"},
			wantOutputName: "sonarqube",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WokeConfig{
				Config:              tt.fields.Config,
				DisableDefaultRules: tt.fields.DisableDefaultRules,
				ExitOneOnFailure:    tt.fields.ExitOneOnFailure,
				NoIgnore:            tt.fields.NoIgnore,
				OutputName:          tt.fields.OutputName,
			}
			c.SetDefaults()
		})
	}
}

func TestWokeResult_Result(t *testing.T) {
	type fields struct {
		Filename string
		Results  []struct {
			Rule struct {
				Name         string   `json:"Name"`
				Terms        []string `json:"Terms"`
				Alternatives []string `json:"Alternatives"`
				Note         string   `json:"Note"`
				Severity     string   `json:"Severity"`
				Options      struct {
					WordBoundary      bool        `json:"WordBoundary"`
					WordBoundaryStart bool        `json:"WordBoundaryStart"`
					WordBoundaryEnd   bool        `json:"WordBoundaryEnd"`
					IncludeNote       bool        `json:"IncludeNote"`
					Categories        interface{} `json:"Categories"`
				} `json:"Options"`
			} `json:"Rule"`
			Finding       string `json:"Finding"`
			Line          string `json:"Line"`
			StartPosition struct {
				Filename string `json:"Filename"`
				Offset   int    `json:"Offset"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			} `json:"StartPosition"`
			EndPosition struct {
				Filename string `json:"Filename"`
				Offset   int    `json:"Offset"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			} `json:"EndPosition"`
			Reason string `json:"Reason"`
		}
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &WokeResult{
				Filename: tt.fields.Filename,
				Results:  tt.fields.Results,
			}
			got, err := m.Result()
			if (err != nil) != tt.wantErr {
				t.Errorf("Result() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Result() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wokeSeverityName(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want rule.SeverityName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wokeSeverityName(tt.args.s); got != tt.want {
				t.Errorf("wokeSeverityName() = %v, want %v", got, tt.want)
			}
		})
	}
}
