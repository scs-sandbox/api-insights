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
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSpectralConfig_SetDefaults(t *testing.T) {
	type fields struct {
		Ruleset *string
	}
	tests := []struct {
		name        string
		fields      fields
		wantRuleset string
	}{
		{
			name:        "normal",
			fields:      fields{Ruleset: nil},
			wantRuleset: "cisco",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SpectralConfig{
				Ruleset: tt.fields.Ruleset,
			}
			c.SetDefaults()
			assert.NotNil(t, c.Ruleset)
			assert.Equal(t, tt.wantRuleset, *c.Ruleset)
		})
	}
}

func TestSpectralConfig_SetRuleset(t *testing.T) {
	type fields struct {
		Ruleset *string
	}
	type args struct {
		ruleset string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantRuleset string
	}{
		{
			name:        "normal",
			fields:      fields{Ruleset: nil},
			args:        args{ruleset: "cisco"},
			wantRuleset: "cisco",
		},
		{
			name:        "normal",
			fields:      fields{Ruleset: nil},
			args:        args{ruleset: "test"},
			wantRuleset: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SpectralConfig{
				Ruleset: tt.fields.Ruleset,
			}
			c.SetRuleset(tt.args.ruleset)
			assert.NotNil(t, c.Ruleset)
			assert.Equal(t, tt.wantRuleset, *c.Ruleset)
		})
	}
}

func TestSpectralResult_Result(t *testing.T) {
	tests := []struct {
		name    string
		m       SpectralResult
		want    *Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Result()
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

func Test_spectralSeverityName(t *testing.T) {
	type args struct {
		spectralSeverity int
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
			if got := spectralSeverityName(tt.args.spectralSeverity); got != tt.want {
				t.Errorf("spectralSeverityName() = %v, want %v", got, tt.want)
			}
		})
	}
}
