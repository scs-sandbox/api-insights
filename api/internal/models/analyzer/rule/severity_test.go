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

package rule

import (
	"reflect"
	"testing"
)

func TestDefaultSeverityWeights(t *testing.T) {
	tests := []struct {
		name string
		want map[SeverityName]int
	}{
		{
			name: "normal",
			want: defaultSeverityWeights,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultSeverityWeights(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultSeverityWeights() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeverityName_Severity(t *testing.T) {
	tests := []struct {
		name string
		n    SeverityName
		want Severity
	}{
		{
			name: "hint",
			n:    SeverityNameHint,
			want: SeverityHint,
		},
		{
			name: "info",
			n:    SeverityNameInfo,
			want: SeverityInfo,
		},
		{
			name: "warning",
			n:    SeverityNameWarning,
			want: SeverityWarning,
		},
		{
			name: "error",
			n:    SeverityNameError,
			want: SeverityError,
		},
		{
			name: "hint - default",
			n:    "unsupported severity name",
			want: SeverityHint,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Severity(); got != tt.want {
				t.Errorf("Severity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeverityName_String(t *testing.T) {
	tests := []struct {
		name string
		n    SeverityName
		want string
	}{
		{
			name: "hint",
			n:    SeverityNameHint,
			want: "hint",
		},
		{
			name: "info",
			n:    SeverityNameInfo,
			want: "info",
		},
		{
			name: "warning",
			n:    SeverityNameWarning,
			want: "warning",
		},
		{
			name: "error",
			n:    SeverityNameError,
			want: "error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeverity_String(t *testing.T) {
	tests := []struct {
		name string
		s    Severity
		want string
	}{
		{
			name: "hint",
			s:    1,
			want: "hint",
		},
		{
			name: "info",
			s:    2,
			want: "info",
		},
		{
			name: "warning",
			s:    3,
			want: "warning",
		},
		{
			name: "error",
			s:    4,
			want: "error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
