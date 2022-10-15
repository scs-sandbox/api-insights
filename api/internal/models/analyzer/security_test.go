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
	"encoding/json"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer/rule"
	panopticamodels "github.com/cisco-developer/api-insights/api/pkg/panoptica/models"
	"os"
	"reflect"
	"testing"
)

func TestGetSecurityResult(t *testing.T) {
	type args struct {
		spec string
		in   *panopticamodels.APIServiceDrillDownExternal
	}
	var in *panopticamodels.APIServiceDrillDownExternal
	var r *Result
	loadJSON("testdata/petstore-v2-panoptica-findings.json", &in)
	loadJSON("testdata/petstore-v2-security-result.json", &r)

	tests := []struct {
		name    string
		args    args
		want    *Result
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				spec: loadSpec("../testdata/petstore-v2.json"),
				in:   in,
			},
			want:    r,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSecurityResult(tt.args.spec, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSecurityResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Summary, tt.want.Summary) || !reflect.DeepEqual(got.Findings, tt.want.Findings) {
				t.Errorf("GetSecurityResult() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSecurityFindingData(t *testing.T) {
	//type args struct {
	//	sf *SecurityFinding
	//}
	//tests := []struct {
	//	name     string
	//	args     args
	//	wantData []*SecurityFinding
	//	wantErr  bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		gotData, err := NewSecurityFindingData(tt.args.sf)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("NewSecurityFindingData() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(gotData, tt.wantData) {
	//			t.Errorf("NewSecurityFindingData() gotData = %v, want %v", gotData, tt.wantData)
	//		}
	//	})
	//}
}

func TestSecurityFinding_JSONPaths(t *testing.T) {
	type fields struct {
		Severity           string
		Kind               string
		Type               string
		Code               string
		Message            string
		Location           []interface{}
		CrRawFindingID     string
		CrFindingIndex     int
		AffectedEndpoints  []interface{}
		Source             string
		SeverityCategory   string
		CrankshaftClassID  string
		CrankshaftSeverity string
		CrankshaftCategory string
		CrankshaftJsonpath string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SecurityFinding{
				Severity:           tt.fields.Severity,
				Kind:               tt.fields.Kind,
				Type:               tt.fields.Type,
				Code:               tt.fields.Code,
				Message:            tt.fields.Message,
				Location:           tt.fields.Location,
				CrRawFindingID:     tt.fields.CrRawFindingID,
				CrFindingIndex:     tt.fields.CrFindingIndex,
				AffectedEndpoints:  tt.fields.AffectedEndpoints,
				Source:             tt.fields.Source,
				SeverityCategory:   tt.fields.SeverityCategory,
				CrankshaftClassID:  tt.fields.CrankshaftClassID,
				CrankshaftSeverity: tt.fields.CrankshaftSeverity,
				CrankshaftCategory: tt.fields.CrankshaftCategory,
				CrankshaftJsonpath: tt.fields.CrankshaftJsonpath,
			}
			if got := m.JSONPaths(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_securityResultSeverityName(t *testing.T) {
	type args struct {
		severity string
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
			if got := securityResultSeverityName(tt.args.severity); got != tt.want {
				t.Errorf("securityResultSeverityName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func loadSpec(file string) string {
	content, _ := os.ReadFile(file)
	return string(content)
}

func loadJSON(file string, v interface{}) {
	data, _ := os.ReadFile(file)
	_ = json.Unmarshal(data, &v)
}
