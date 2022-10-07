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

func TestNewResult(t *testing.T) {
	tests := []struct {
		name string
		want *Result
	}{
		{
			name: "normal",
			want: &Result{
				Summary: &ResultSummary{
					Stats: &SeverityRuleFindingsStats{
						Hint:    &RuleFindingsStats{Data: map[rule.NameID]int{}},
						Info:    &RuleFindingsStats{Data: map[rule.NameID]int{}},
						Warning: &RuleFindingsStats{Data: map[rule.NameID]int{}},
						Error:   &RuleFindingsStats{Data: map[rule.NameID]int{}},
					},
				},
				Findings: SeverityRuleFindings{
					rule.SeverityNameHint:    &RuleFindings{},
					rule.SeverityNameInfo:    &RuleFindings{},
					rule.SeverityNameWarning: &RuleFindings{},
					rule.SeverityNameError:   &RuleFindings{},
				},
				severityRuleCache: map[rule.SeverityName]map[rule.NameID]*Rule{
					rule.SeverityNameHint:    {},
					rule.SeverityNameInfo:    {},
					rule.SeverityNameWarning: {},
					rule.SeverityNameError:   {},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResult(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewResultSummary(t *testing.T) {
	tests := []struct {
		name string
		want *ResultSummary
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResultSummary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResultSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSeverityRuleFindings(t *testing.T) {
	tests := []struct {
		name string
		want SeverityRuleFindings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSeverityRuleFindings(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSeverityRuleFindings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_AddFinding(t *testing.T) {
	type fields struct {
		Summary           *ResultSummary
		Findings          SeverityRuleFindings
		severityRuleCache map[rule.SeverityName]map[rule.NameID]*Rule
	}
	type args struct {
		severity   rule.SeverityName
		ruleNameID rule.NameID
		finding    *Finding
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantFindings *RuleFindings
	}{
		{
			name: "normal - error",
			fields: fields{
				Summary: nil,
				Findings: SeverityRuleFindings{
					rule.SeverityNameHint:    &RuleFindings{},
					rule.SeverityNameInfo:    &RuleFindings{},
					rule.SeverityNameWarning: &RuleFindings{},
					rule.SeverityNameError:   &RuleFindings{},
				},
				severityRuleCache: map[rule.SeverityName]map[rule.NameID]*Rule{},
			},
			args: args{
				severity:   rule.SeverityNameError,
				ruleNameID: "sample-error-rule",
				finding: &Finding{
					Type: "test",
					Path: []string{"paths", "/catalogue"},
					Range: &FindingPositionRange{
						Start: &FindingPosition{},
						End:   &FindingPosition{},
					},
				},
			},
			wantFindings: &RuleFindings{
				Rules: map[rule.NameID]*Findings{
					"sample-error-rule": &Findings{
						Message:    "",
						Mitigation: "",
						Data: []*Finding{{
							Type: "test",
							Path: []string{"paths", "/catalogue"},
							Range: &FindingPositionRange{
								Start: &FindingPosition{},
								End:   &FindingPosition{},
							},
						}},
					}},
			},
		},
		{
			name: "normal - warning",
			fields: fields{
				Summary: nil,
				Findings: SeverityRuleFindings{
					rule.SeverityNameHint:    &RuleFindings{},
					rule.SeverityNameInfo:    &RuleFindings{},
					rule.SeverityNameWarning: &RuleFindings{},
					rule.SeverityNameError:   &RuleFindings{},
				},
				severityRuleCache: map[rule.SeverityName]map[rule.NameID]*Rule{},
			},
			args: args{
				severity:   rule.SeverityNameWarning,
				ruleNameID: "sample-warning-rule",
				finding: &Finding{
					Type: "test",
					Path: []string{"paths", "/catalogue"},
					Range: &FindingPositionRange{
						Start: &FindingPosition{},
						End:   &FindingPosition{},
					},
				},
			},
			wantFindings: &RuleFindings{
				Rules: map[rule.NameID]*Findings{
					"sample-warning-rule": &Findings{
						Message:    "",
						Mitigation: "",
						Data: []*Finding{{
							Type: "test",
							Path: []string{"paths", "/catalogue"},
							Range: &FindingPositionRange{
								Start: &FindingPosition{},
								End:   &FindingPosition{},
							},
						}},
					}},
			},
		},
		{
			name: "normal - info",
			fields: fields{
				Summary: nil,
				Findings: SeverityRuleFindings{
					rule.SeverityNameHint:    &RuleFindings{},
					rule.SeverityNameInfo:    &RuleFindings{},
					rule.SeverityNameWarning: &RuleFindings{},
					rule.SeverityNameError:   &RuleFindings{},
				},
				severityRuleCache: map[rule.SeverityName]map[rule.NameID]*Rule{},
			},
			args: args{
				severity:   rule.SeverityNameInfo,
				ruleNameID: "sample-info-rule",
				finding: &Finding{
					Type: "test",
					Path: []string{"paths", "/catalogue"},
					Range: &FindingPositionRange{
						Start: &FindingPosition{},
						End:   &FindingPosition{},
					},
				},
			},
			wantFindings: &RuleFindings{
				Rules: map[rule.NameID]*Findings{
					"sample-info-rule": &Findings{
						Message:    "",
						Mitigation: "",
						Data: []*Finding{{
							Type: "test",
							Path: []string{"paths", "/catalogue"},
							Range: &FindingPositionRange{
								Start: &FindingPosition{},
								End:   &FindingPosition{},
							},
						}},
					}},
			},
		},
		{
			name: "normal - hint",
			fields: fields{
				Summary: nil,
				Findings: SeverityRuleFindings{
					rule.SeverityNameHint:    &RuleFindings{},
					rule.SeverityNameInfo:    &RuleFindings{},
					rule.SeverityNameWarning: &RuleFindings{},
					rule.SeverityNameError:   &RuleFindings{},
				},
				severityRuleCache: map[rule.SeverityName]map[rule.NameID]*Rule{},
			},
			args: args{
				severity:   rule.SeverityNameHint,
				ruleNameID: "sample-hint-rule",
				finding: &Finding{
					Type: "test",
					Path: []string{"paths", "/catalogue"},
					Range: &FindingPositionRange{
						Start: &FindingPosition{},
						End:   &FindingPosition{},
					},
				},
			},
			wantFindings: &RuleFindings{
				Rules: map[rule.NameID]*Findings{
					"sample-hint-rule": &Findings{
						Message:    "",
						Mitigation: "",
						Data: []*Finding{{
							Type: "test",
							Path: []string{"paths", "/catalogue"},
							Range: &FindingPositionRange{
								Start: &FindingPosition{},
								End:   &FindingPosition{},
							},
						}},
					}},
			},
		},
		{
			name: "normal - hint with cached findings - different id",
			fields: fields{
				Summary: nil,
				Findings: SeverityRuleFindings{
					rule.SeverityNameHint:    &RuleFindings{},
					rule.SeverityNameInfo:    &RuleFindings{},
					rule.SeverityNameWarning: &RuleFindings{},
					rule.SeverityNameError:   &RuleFindings{},
				},
				severityRuleCache: map[rule.SeverityName]map[rule.NameID]*Rule{
					rule.SeverityNameHint: {
						"sample-hint-rule-01": &Rule{
							NameID:         "sample-hint-rule-01",
							AnalyzerNameID: "completeness",
							Severity:       rule.SeverityNameHint.String(),
						},
					},
					rule.SeverityNameInfo:    {},
					rule.SeverityNameWarning: {},
					rule.SeverityNameError:   {},
				},
			},
			args: args{
				severity:   rule.SeverityNameHint,
				ruleNameID: "sample-hint-rule-02",
				finding: &Finding{
					Type: "test",
					Path: []string{"paths", "/catalogue"},
					Range: &FindingPositionRange{
						Start: &FindingPosition{},
						End:   &FindingPosition{},
					},
				},
			},
			wantFindings: &RuleFindings{
				Rules: map[rule.NameID]*Findings{
					"sample-hint-rule-01": &Findings{
						Message:    "",
						Mitigation: "",
						Data: []*Finding{{
							Type: "test",
							Path: []string{"paths", "/catalogue"},
							Range: &FindingPositionRange{
								Start: &FindingPosition{},
								End:   &FindingPosition{},
							},
						}},
					},
					"sample-hint-rule-02": &Findings{
						Message:    "",
						Mitigation: "",
						Data: []*Finding{{
							Type: "test",
							Path: []string{"paths", "/catalogue"},
							Range: &FindingPositionRange{
								Start: &FindingPosition{},
								End:   &FindingPosition{},
							},
						}},
					},
				},
			},
		},
		{
			name: "normal - hint with cached findings - same id",
			fields: fields{
				Summary: nil,
				Findings: SeverityRuleFindings{
					rule.SeverityNameHint:    &RuleFindings{},
					rule.SeverityNameInfo:    &RuleFindings{},
					rule.SeverityNameWarning: &RuleFindings{},
					rule.SeverityNameError:   &RuleFindings{},
				},
				severityRuleCache: map[rule.SeverityName]map[rule.NameID]*Rule{
					rule.SeverityNameHint: {
						"sample-hint-rule-01": &Rule{
							NameID:         "sample-hint-rule-01",
							AnalyzerNameID: "completeness",
							Severity:       rule.SeverityNameHint.String(),
						},
					},
					rule.SeverityNameInfo:    {},
					rule.SeverityNameWarning: {},
					rule.SeverityNameError:   {},
				},
			},
			args: args{
				severity:   rule.SeverityNameHint,
				ruleNameID: "sample-hint-rule-01",
				finding: &Finding{
					Type: "test",
					Path: []string{"paths", "/catalogue/{id}"},
					Range: &FindingPositionRange{
						Start: &FindingPosition{},
						End:   &FindingPosition{},
					},
				},
			},
			wantFindings: &RuleFindings{
				Rules: map[rule.NameID]*Findings{
					"sample-hint-rule-01": &Findings{
						Message:    "",
						Mitigation: "",
						Data: []*Finding{
							{
								Type: "test",
								Path: []string{"paths", "/catalogue"},
								Range: &FindingPositionRange{
									Start: &FindingPosition{},
									End:   &FindingPosition{},
								},
							},
							{
								Type: "test",
								Path: []string{"paths", "/catalogue/{id}"},
								Range: &FindingPositionRange{
									Start: &FindingPosition{},
									End:   &FindingPosition{},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Result{
				Summary:           tt.fields.Summary,
				Findings:          tt.fields.Findings,
				severityRuleCache: tt.fields.severityRuleCache,
			}
			r.AddFinding(tt.args.severity, tt.args.ruleNameID, tt.args.finding)
		})
	}
}

func TestResult_getRuleFromCache(t *testing.T) {
	type fields struct {
		Summary           *ResultSummary
		Findings          SeverityRuleFindings
		severityRuleCache map[rule.SeverityName]map[rule.NameID]*Rule
	}
	type args struct {
		severity   rule.SeverityName
		ruleNameID rule.NameID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Rule
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Result{
				Summary:           tt.fields.Summary,
				Findings:          tt.fields.Findings,
				severityRuleCache: tt.fields.severityRuleCache,
			}
			if got := r.getRuleFromCache(tt.args.severity, tt.args.ruleNameID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRuleFromCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_storeRuleInCache(t *testing.T) {
	type fields struct {
		Summary           *ResultSummary
		Findings          SeverityRuleFindings
		severityRuleCache map[rule.SeverityName]map[rule.NameID]*Rule
	}
	type args struct {
		severity     rule.SeverityName
		ruleNameID   rule.NameID
		analyzerRule *Rule
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[rule.SeverityName]map[rule.NameID]*Rule
	}{
		{
			name: "normal",
			fields: fields{
				severityRuleCache: map[rule.SeverityName]map[rule.NameID]*Rule{},
			},
			args: args{
				severity:   rule.SeverityNameError,
				ruleNameID: "sample-rule",
				analyzerRule: &Rule{
					NameID:         "sample-rule",
					AnalyzerNameID: "completeness",
					Severity:       rule.SeverityNameError.String(),
				},
			},
			want: map[rule.SeverityName]map[rule.NameID]*Rule{
				rule.SeverityNameError: {
					"sample-rule": &Rule{
						NameID:         "sample-rule",
						AnalyzerNameID: "completeness",
						Severity:       rule.SeverityNameError.String(),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Result{
				Summary:           tt.fields.Summary,
				Findings:          tt.fields.Findings,
				severityRuleCache: tt.fields.severityRuleCache,
			}
			r.storeRuleInCache(tt.args.severity, tt.args.ruleNameID, tt.args.analyzerRule)
			assert.Equal(t, r.severityRuleCache, tt.want)
		})
	}
}

func TestResult_updateSummaryStatsAfterAddFinding(t *testing.T) {
	type fields struct {
		Summary           *ResultSummary
		Findings          SeverityRuleFindings
		severityRuleCache map[rule.SeverityName]map[rule.NameID]*Rule
	}
	type args struct {
		severity   rule.SeverityName
		ruleNameID rule.NameID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Result{
				Summary:           tt.fields.Summary,
				Findings:          tt.fields.Findings,
				severityRuleCache: tt.fields.severityRuleCache,
			}
			r.updateSummaryStatsAfterAddFinding(tt.args.severity, tt.args.ruleNameID)
		})
	}
}

func Test_newSeverityRuleCache(t *testing.T) {
	tests := []struct {
		name string
		want map[rule.SeverityName]map[rule.NameID]*Rule
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSeverityRuleCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSeverityRuleCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
