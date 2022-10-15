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
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer/rule"
	"reflect"
	"testing"
	"time"
)

var (
	serviceID         = "test"
	specID            = "test"
	weight    float32 = 25
	spec              = &models.Spec{}
	sa                = &models.SpecAnalysis{}

	saSecurity = &models.SpecAnalysis{
		ID:                 serviceID,
		Analyzer:           analyzer.Security,
		SpecAnalysisConfig: models.SpecAnalysisConfig{},
		SpecAnalysisResult: models.SpecAnalysisResult{
			Result: &analyzer.Result{
				Summary: &analyzer.ResultSummary{Stats: &analyzer.SeverityRuleFindingsStats{
					Count:       0,
					Occurrences: 0,
					Hint:        nil,
					Info:        nil,
					Warning:     nil,
					Error:       nil,
				}},
				Findings: nil,
			},
			RawResult: nil,
		},
		Score:     nil,
		ServiceID: serviceID,
		SpecID:    specID,
		Status:    "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	sas = map[analyzer.SpecAnalyzer]*models.SpecAnalysis{
		analyzer.Completeness: &models.SpecAnalysis{
			ID:                 "",
			Analyzer:           analyzer.Completeness,
			SpecAnalysisConfig: models.SpecAnalysisConfig{},
			SpecAnalysisResult: models.SpecAnalysisResult{},
			Score:              nil,
			ServiceID:          "",
			SpecID:             "",
			Status:             "",
			CreatedAt:          time.Time{},
			UpdatedAt:          time.Time{},
		},
	}

	newRule = func(sa analyzer.SpecAnalyzer, nameID string, severity rule.Severity) *analyzer.Rule {
		return &analyzer.Rule{
			NameID:         nameID,
			AnalyzerNameID: string(sa),
			Severity:       severity.String(),
		}
	}

	sampleRule1      = newRule(analyzer.Completeness, "sample-rule-1", rule.SeverityError)
	sampleRule2      = newRule(analyzer.Completeness, "sample-rule-2", rule.SeverityWarning)
	sampleRule3      = newRule(analyzer.Completeness, "sample-rule-3", rule.SeverityInfo)
	sampleRule4      = newRule(analyzer.Completeness, "sample-rule-4", rule.SeverityHint)
	sampleRuleError1 = newRule(analyzer.Completeness, "sample-rule-error-1", rule.SeverityHint)

	analyzers = map[analyzer.SpecAnalyzer]*analyzer.Analyzer{
		analyzer.Completeness: &analyzer.Analyzer{
			ID:          "",
			NameID:      string(analyzer.Completeness),
			Title:       "",
			Description: "",
			Status:      "",
			Meta:        nil,
			Config:      nil,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			Position:    0,
			Rules:       []*analyzer.Rule{sampleRule1, sampleRule2, sampleRule3, sampleRule4, sampleRuleError1},
		},
	}

	analyzerRules = map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule{
		analyzer.Completeness: {
			rule.SeverityNameError:   []*analyzer.Rule{sampleRule1, sampleRuleError1},
			rule.SeverityNameWarning: []*analyzer.Rule{sampleRule2},
			rule.SeverityNameInfo:    []*analyzer.Rule{sampleRule3},
			rule.SeverityNameHint:    []*analyzer.Rule{sampleRule4},
		},
	}

	scoreCfgsOnlyCompleteness = analyzer.AnalyzersScoreConfigs{
		analyzer.Completeness: &analyzer.ScoreConfig{
			AnalyzerWeight: &weight,
			SeverityWeights: map[rule.SeverityName]int{
				rule.SeverityNameError:   10,
				rule.SeverityNameWarning: 20,
				rule.SeverityNameInfo:    30,
				rule.SeverityNameHint:    40,
			},
		},
	}

	scoreCfgs = analyzer.AnalyzersScoreConfigs{
		analyzer.Completeness: &analyzer.ScoreConfig{
			AnalyzerWeight: &weight,
			SeverityWeights: map[rule.SeverityName]int{
				rule.SeverityNameError:   10,
				rule.SeverityNameWarning: 20,
				rule.SeverityNameInfo:    30,
				rule.SeverityNameHint:    40,
			},
		},
		analyzer.Security: &analyzer.ScoreConfig{
			AnalyzerWeight: &weight,
			SeverityWeights: map[rule.SeverityName]int{
				rule.SeverityNameError:   10,
				rule.SeverityNameWarning: 20,
				rule.SeverityNameInfo:    30,
				rule.SeverityNameHint:    40,
			},
		},
		analyzer.CiscoAPIGuidelines: &analyzer.ScoreConfig{
			AnalyzerWeight: &weight,
			SeverityWeights: map[rule.SeverityName]int{
				rule.SeverityNameError:   10,
				rule.SeverityNameWarning: 20,
				rule.SeverityNameInfo:    30,
				rule.SeverityNameHint:    40,
			},
		},
		analyzer.Drift: &analyzer.ScoreConfig{
			AnalyzerWeight: &weight,
			SeverityWeights: map[rule.SeverityName]int{
				rule.SeverityNameError:   10,
				rule.SeverityNameWarning: 20,
				rule.SeverityNameInfo:    30,
				rule.SeverityNameHint:    40,
			},
		},
	}

	nofindings = &analyzer.RuleFindingsStats{
		Count:       0,
		Occurrences: 0,
		Data:        map[rule.NameID]int{},
	}
)

func TestNewReporter(t *testing.T) {
	type args struct {
		scoreCfg  analyzer.AnalyzersScoreConfigs
		analyzers map[analyzer.SpecAnalyzer]*analyzer.Analyzer
	}

	tests := []struct {
		name    string
		args    args
		want    *reporter
		wantErr bool
	}{
		//{
		//	name: "normal",
		//	args: args{
		//		scoreCfg:  scoreCfgs,
		//		analyzers: analyzers,
		//	},
		//	want: &reporter{
		//		scoreCfg:      scoreCfgs,
		//		analyzerRules: analyzerRules,
		//	},
		//},
		{
			name: "nil - empty config",
			args: args{
				scoreCfg:  nil,
				analyzers: analyzers,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nil - empty analyzers",
			args: args{
				scoreCfg:  scoreCfgs,
				analyzers: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReporter(tt.args.scoreCfg, tt.args.analyzers)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewReporter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReporter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSpecAnalysisReport(t *testing.T) {
	type args struct {
		specAnalysis *models.SpecAnalysis
		spec         *models.Spec
	}

	tests := []struct {
		name string
		args args
		want *SpecAnalysisReport
	}{
		{
			name: "normal",
			args: args{
				specAnalysis: sa,
				spec:         spec,
			},
			want: &SpecAnalysisReport{
				specAnalysis: sa,
				spec:         spec,
				score:        0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpecAnalysisReport(tt.args.specAnalysis, tt.args.spec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpecAnalysisReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSpecReport(t *testing.T) {
	type args struct {
		spec         *models.Spec
		specAnalyses map[analyzer.SpecAnalyzer]*models.SpecAnalysis
	}

	tests := []struct {
		name string
		args args
		want *SpecReport
	}{
		{
			name: "normal",
			args: args{
				spec:         spec,
				specAnalyses: sas,
			},
			want: &SpecReport{
				spec:                spec,
				specAnalyses:        sas,
				specAnalysisReports: map[analyzer.SpecAnalyzer]*SpecAnalysisReport{},
				score:               0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpecReport(tt.args.spec, tt.args.specAnalyses); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpecReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecAnalysisReport_Score(t *testing.T) {
	type fields struct {
		specAnalysis *models.SpecAnalysis
		spec         *models.Spec
		score        int
	}
	type args struct {
		scoreCfg      analyzer.AnalyzersScoreConfigs
		analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "fail - nil specAnalysis",
			fields: fields{
				specAnalysis: sa,
				spec:         spec,
				score:        0,
			},
			args: args{
				scoreCfg:      scoreCfgs,
				analyzerRules: analyzerRules,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "fail - missing analyzer score config",
			fields: fields{
				specAnalysis: saSecurity,
				spec:         spec,
				score:        0,
			},
			args: args{
				scoreCfg:      scoreCfgsOnlyCompleteness,
				analyzerRules: analyzerRules,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "normal - 0 finding based on 5 rules - score 100",
			fields: fields{
				specAnalysis: &models.SpecAnalysis{
					Analyzer:           analyzer.Completeness,
					SpecAnalysisConfig: models.SpecAnalysisConfig{},
					SpecAnalysisResult: models.SpecAnalysisResult{
						Result: &analyzer.Result{
							Summary: &analyzer.ResultSummary{
								Stats: &analyzer.SeverityRuleFindingsStats{
									Count:       0,
									Occurrences: 0,
									Error: &analyzer.RuleFindingsStats{
										Count:       0,
										Occurrences: 0,
										Data:        map[rule.NameID]int{},
									},
									Warning: nofindings,
									Info:    nofindings,
									Hint:    nofindings,
								}},
						},
					},
				},
			},
			args: args{
				scoreCfg: scoreCfgs,
				analyzerRules: map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule{
					analyzer.Completeness: {
						rule.SeverityNameError:   []*analyzer.Rule{sampleRule1, sampleRuleError1},
						rule.SeverityNameWarning: []*analyzer.Rule{sampleRule2},
						rule.SeverityNameInfo:    []*analyzer.Rule{sampleRule3},
						rule.SeverityNameHint:    []*analyzer.Rule{sampleRule4},
					},
				},
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "normal - 1 finding based on 2 rules - score 50",
			fields: fields{
				specAnalysis: &models.SpecAnalysis{
					Analyzer:           analyzer.Completeness,
					SpecAnalysisConfig: models.SpecAnalysisConfig{},
					SpecAnalysisResult: models.SpecAnalysisResult{
						Result: &analyzer.Result{
							Summary: &analyzer.ResultSummary{
								Stats: &analyzer.SeverityRuleFindingsStats{
									Count:       1,
									Occurrences: 1,
									Error: &analyzer.RuleFindingsStats{
										Count:       1,
										Occurrences: 1,
										Data: map[rule.NameID]int{
											rule.NameID(sampleRule1.NameID): 1,
										},
									},
									Warning: nofindings,
									Info:    nofindings,
									Hint:    nofindings,
								}},
						},
					},
				},
			},
			args: args{
				scoreCfg: scoreCfgs,
				analyzerRules: map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule{
					analyzer.Completeness: {
						rule.SeverityNameError: []*analyzer.Rule{sampleRule1, sampleRuleError1},
					},
				},
			},
			want:    50,
			wantErr: false,
		},
		{
			name: "normal - 2 finding based on 2 rules - score 0",
			fields: fields{
				specAnalysis: &models.SpecAnalysis{
					Analyzer:           analyzer.Completeness,
					SpecAnalysisConfig: models.SpecAnalysisConfig{},
					SpecAnalysisResult: models.SpecAnalysisResult{
						Result: &analyzer.Result{
							Summary: &analyzer.ResultSummary{
								Stats: &analyzer.SeverityRuleFindingsStats{
									Count:       2,
									Occurrences: 2,
									Error: &analyzer.RuleFindingsStats{
										Count:       2,
										Occurrences: 2,
										Data: map[rule.NameID]int{
											rule.NameID(sampleRule1.NameID):      1,
											rule.NameID(sampleRuleError1.NameID): 1,
										},
									},
									Warning: nofindings,
									Info:    nofindings,
									Hint:    nofindings,
								}},
						},
					},
				},
			},
			args: args{
				scoreCfg: scoreCfgs,
				analyzerRules: map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule{
					analyzer.Completeness: {
						rule.SeverityNameError: []*analyzer.Rule{sampleRule1, sampleRuleError1},
					},
				},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpecAnalysisReport{
				specAnalysis: tt.fields.specAnalysis,
				spec:         tt.fields.spec,
				score:        tt.fields.score,
			}
			got, err := r.Score(tt.args.scoreCfg, tt.args.analyzerRules)
			if (err != nil) != tt.wantErr {
				t.Errorf("Score() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Score() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecAnalysisReport_WithMitigation(t *testing.T) {
	type fields struct {
		specAnalysis *models.SpecAnalysis
		spec         *models.Spec
		score        int
	}
	type args struct {
		analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SpecAnalysisReport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpecAnalysisReport{
				specAnalysis: tt.fields.specAnalysis,
				spec:         tt.fields.spec,
				score:        tt.fields.score,
			}
			if got := r.WithMitigation(tt.args.analyzerRules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMitigation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecAnalysisReport_WithScore(t *testing.T) {
	type fields struct {
		specAnalysis *models.SpecAnalysis
		spec         *models.Spec
		score        int
	}
	type args struct {
		scoreCfg      analyzer.AnalyzersScoreConfigs
		analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SpecAnalysisReport
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpecAnalysisReport{
				specAnalysis: tt.fields.specAnalysis,
				spec:         tt.fields.spec,
				score:        tt.fields.score,
			}
			got, err := r.WithScore(tt.args.scoreCfg, tt.args.analyzerRules)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithScore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecReport_Score(t *testing.T) {
	type fields struct {
		spec                *models.Spec
		specAnalyses        map[analyzer.SpecAnalyzer]*models.SpecAnalysis
		specAnalysisReports map[analyzer.SpecAnalyzer]*SpecAnalysisReport
		score               int
	}
	type args struct {
		scoreCfg      analyzer.AnalyzersScoreConfigs
		analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpecReport{
				spec:                tt.fields.spec,
				specAnalyses:        tt.fields.specAnalyses,
				specAnalysisReports: tt.fields.specAnalysisReports,
				score:               tt.fields.score,
			}
			got, err := r.Score(tt.args.scoreCfg, tt.args.analyzerRules)
			if (err != nil) != tt.wantErr {
				t.Errorf("Score() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Score() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecReport_WithMitigation(t *testing.T) {
	type fields struct {
		spec                *models.Spec
		specAnalyses        map[analyzer.SpecAnalyzer]*models.SpecAnalysis
		specAnalysisReports map[analyzer.SpecAnalyzer]*SpecAnalysisReport
		score               int
	}
	type args struct {
		analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SpecReport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpecReport{
				spec:                tt.fields.spec,
				specAnalyses:        tt.fields.specAnalyses,
				specAnalysisReports: tt.fields.specAnalysisReports,
				score:               tt.fields.score,
			}
			if got := r.WithMitigation(tt.args.analyzerRules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMitigation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecReport_WithScore(t *testing.T) {
	type fields struct {
		spec                *models.Spec
		specAnalyses        map[analyzer.SpecAnalyzer]*models.SpecAnalysis
		specAnalysisReports map[analyzer.SpecAnalyzer]*SpecAnalysisReport
		score               int
	}
	type args struct {
		scoreCfg      analyzer.AnalyzersScoreConfigs
		analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SpecReport
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpecReport{
				spec:                tt.fields.spec,
				specAnalyses:        tt.fields.specAnalyses,
				specAnalysisReports: tt.fields.specAnalysisReports,
				score:               tt.fields.score,
			}
			got, err := r.WithScore(tt.args.scoreCfg, tt.args.analyzerRules)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithScore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRule(t *testing.T) {
	type args struct {
		analyzer      analyzer.SpecAnalyzer
		id            string
		analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule
	}

	tests := []struct {
		name string
		args args
		want *analyzer.Rule
	}{
		{
			name: "normal",
			args: args{
				analyzer:      analyzer.Completeness,
				id:            "sample-rule-1",
				analyzerRules: analyzerRules,
			},
			want: sampleRule1,
		},
		{
			name: "unsupported analyzer",
			args: args{
				analyzer:      "unsupported-analyzer",
				id:            "sample",
				analyzerRules: analyzerRules,
			},
			want: nil,
		},
		{
			name: "unsupported analyzer rule",
			args: args{
				analyzer:      analyzer.Completeness,
				id:            "unsupported-rule",
				analyzerRules: analyzerRules,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRule(tt.args.analyzer, tt.args.id, tt.args.analyzerRules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reporter_GenerateSpecAnalysisReport(t *testing.T) {
	type fields struct {
		scoreCfg      analyzer.AnalyzersScoreConfigs
		analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule
	}
	type args struct {
		spec         *models.Spec
		specAnalysis *models.SpecAnalysis
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SpecAnalysisReport
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := reporter{
				scoreCfg:      tt.fields.scoreCfg,
				analyzerRules: tt.fields.analyzerRules,
			}
			got, err := r.GenerateSpecAnalysisReport(tt.args.spec, tt.args.specAnalysis)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateSpecAnalysisReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateSpecAnalysisReport() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reporter_GenerateSpecReport(t *testing.T) {
	type fields struct {
		scoreCfg      analyzer.AnalyzersScoreConfigs
		analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule
	}
	type args struct {
		spec         *models.Spec
		specAnalyses map[analyzer.SpecAnalyzer]*models.SpecAnalysis
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SpecReport
		wantErr bool
	}{
		{
			name: "fail",
			fields: fields{
				scoreCfg:      scoreCfgs,
				analyzerRules: analyzerRules,
			},
			args: args{
				spec:         spec,
				specAnalyses: sas,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := reporter{
				scoreCfg:      tt.fields.scoreCfg,
				analyzerRules: tt.fields.analyzerRules,
			}
			got, err := r.GenerateSpecReport(tt.args.spec, tt.args.specAnalyses)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateSpecReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateSpecReport() got = %v, want %v", got, tt.want)
			}
		})
	}
}
