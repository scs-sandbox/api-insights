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
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer/rule"
)

type Scorer interface {
	Score(scoreCfg analyzer.AnalyzersScoreConfigs, analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule) (int, error)
}

type SpecReport struct {
	spec         *models.Spec
	specAnalyses map[analyzer.SpecAnalyzer]*models.SpecAnalysis

	specAnalysisReports map[analyzer.SpecAnalyzer]*SpecAnalysisReport

	score int
}

func NewSpecReport(spec *models.Spec, specAnalyses map[analyzer.SpecAnalyzer]*models.SpecAnalysis) *SpecReport {
	return &SpecReport{
		spec:                spec,
		specAnalyses:        specAnalyses,
		specAnalysisReports: map[analyzer.SpecAnalyzer]*SpecAnalysisReport{},
	}
}

func (r *SpecReport) WithMitigation(analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule) *SpecReport {
	for _, specAnalysis := range r.specAnalyses {
		NewSpecAnalysisReport(specAnalysis, r.spec).WithMitigation(analyzerRules)
	}

	return r
}

func (r *SpecReport) WithScore(scoreCfg analyzer.AnalyzersScoreConfigs, analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule) (*SpecReport, error) {
	if _, err := r.Score(scoreCfg, analyzerRules); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *SpecReport) Score(scoreCfg analyzer.AnalyzersScoreConfigs, analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule) (int, error) {
	var score float32
	var weightSum float32
	for analyzerName, specAnalysis := range r.specAnalyses {
		specAnalysisReport, err := NewSpecAnalysisReport(specAnalysis, r.spec).WithScore(scoreCfg, analyzerRules)
		if err != nil {
			return 0, fmt.Errorf("reporter.GenerateSpecReport: %v", err)
		}
		r.specAnalysisReports[analyzerName] = specAnalysisReport
		weight := *scoreCfg[analyzerName].AnalyzerWeight
		weightSum += weight
		score += float32(specAnalysisReport.score) * weight
	}
	if weightSum != 0 {
		r.score = int((score / weightSum) + 0.5)
	}
	if r.score < 0 {
		r.score = 0
	}
	return r.score, nil
}

// SpecAnalysisReport represents a report of a models.SpecAnalysis.
type SpecAnalysisReport struct {
	specAnalysis *models.SpecAnalysis
	spec         *models.Spec

	score int
}

func NewSpecAnalysisReport(specAnalysis *models.SpecAnalysis, spec *models.Spec) *SpecAnalysisReport {
	return &SpecAnalysisReport{specAnalysis: specAnalysis, spec: spec}
}

func (r *SpecAnalysisReport) WithMitigation(analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule) *SpecAnalysisReport {
	if r.specAnalysis == nil || r.specAnalysis.Result == nil {
		return r
	}

	for _, findings := range r.specAnalysis.Result.Findings {
		for nameID, f := range findings.Rules {
			if rule := findRule(r.specAnalysis.Analyzer, string(nameID), analyzerRules); rule != nil {
				// overwrite mitigation and message if specified in the rule
				if len(rule.Mitigation) > 0 {
					f.Mitigation = rule.Mitigation
				}
				if len(rule.Description) > 0 {
					f.Message = rule.Description
				}
			}
		}
	}

	return r
}

func findRule(analyzer analyzer.SpecAnalyzer, id string, analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule) *analyzer.Rule {
	rulesMap, found := analyzerRules[analyzer]
	if !found {
		return nil
	}

	for _, rules := range rulesMap {
		for _, rule := range rules {
			if id == rule.NameID {
				return rule
			}
		}
	}

	return nil
}

func (r *SpecAnalysisReport) WithScore(scoreCfg analyzer.AnalyzersScoreConfigs, analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule) (*SpecAnalysisReport, error) {
	if _, err := r.Score(scoreCfg, analyzerRules); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *SpecAnalysisReport) Score(scoreCfg analyzer.AnalyzersScoreConfigs, analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule) (int, error) {
	if r.specAnalysis == nil || r.specAnalysis.Result == nil {
		return 0, fmt.Errorf("analyzer: cannot Score w/o SpecAnalysis or SpecAnalysis.Result")
	}

	analyzerScoreCfg, ok := scoreCfg[r.specAnalysis.Analyzer]
	if !ok {
		return 0, fmt.Errorf("analyzer: cannot Score w/o AnalyzersScoreConfigs for analyzer(%s)", r.specAnalysis.Analyzer)
	}
	var scoreLoss int
	var maxScorePossible int
	if r.specAnalysis.Result.Summary != nil && r.specAnalysis.Result.Summary.Stats != nil {
		if r.specAnalysis.Result.Summary.Stats.Hint != nil {
			severityWeight := analyzerScoreCfg.SeverityWeights[rule.SeverityNameHint]
			scoreLoss += severityWeight * len(r.specAnalysis.Result.Summary.Stats.Hint.Data)
			maxScorePossible += severityWeight * len(analyzerRules[r.specAnalysis.Analyzer][rule.SeverityNameHint])
		}
		if r.specAnalysis.Result.Summary.Stats.Info != nil {
			severityWeight := analyzerScoreCfg.SeverityWeights[rule.SeverityNameInfo]
			scoreLoss += severityWeight * len(r.specAnalysis.Result.Summary.Stats.Info.Data)
			maxScorePossible += severityWeight * len(analyzerRules[r.specAnalysis.Analyzer][rule.SeverityNameInfo])
		}
		if r.specAnalysis.Result.Summary.Stats.Warning != nil {
			severityWeight := analyzerScoreCfg.SeverityWeights[rule.SeverityNameWarning]
			scoreLoss += severityWeight * len(r.specAnalysis.Result.Summary.Stats.Warning.Data)
			maxScorePossible += severityWeight * len(analyzerRules[r.specAnalysis.Analyzer][rule.SeverityNameWarning])
		}
		if r.specAnalysis.Result.Summary.Stats.Error != nil {
			severityWeight := analyzerScoreCfg.SeverityWeights[rule.SeverityNameError]
			scoreLoss += severityWeight * len(r.specAnalysis.Result.Summary.Stats.Error.Data)
			maxScorePossible += severityWeight * len(analyzerRules[r.specAnalysis.Analyzer][rule.SeverityNameError])
		}
	}
	if maxScorePossible != 0 {
		r.score = int((float32(maxScorePossible-scoreLoss) / float32(maxScorePossible)) * 100)
	}
	if r.score < 0 {
		r.score = 0
	}
	return r.score, nil
}

type Reporter interface {
	GenerateSpecReport(spec *models.Spec, specAnalyses map[analyzer.SpecAnalyzer]*models.SpecAnalysis) (*SpecReport, error)
	GenerateSpecAnalysisReport(spec *models.Spec, specAnalysis *models.SpecAnalysis) (*SpecAnalysisReport, error)
}

type reporter struct {
	scoreCfg      analyzer.AnalyzersScoreConfigs
	analyzerRules map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule
}

func NewReporter(scoreCfg analyzer.AnalyzersScoreConfigs, analyzers map[analyzer.SpecAnalyzer]*analyzer.Analyzer) (Reporter, error) {
	if scoreCfg == nil || len(analyzers) == 0 {
		return nil, fmt.Errorf("analyzer: cannot create Reporter w/o AnalyzersScoreConfigs or analyzers")
	}
	r := &reporter{
		scoreCfg:      scoreCfg,
		analyzerRules: map[analyzer.SpecAnalyzer]map[rule.SeverityName][]*analyzer.Rule{},
	}
	for analyzerName, a := range analyzers {
		if _, ok := r.analyzerRules[analyzerName]; !ok {
			r.analyzerRules[analyzerName] = map[rule.SeverityName][]*analyzer.Rule{}
		}
		for _, analyzerRule := range a.Rules {
			r.analyzerRules[analyzerName][rule.SeverityName(analyzerRule.Severity)] = append(
				r.analyzerRules[analyzerName][rule.SeverityName(analyzerRule.Severity)],
				analyzerRule,
			)
		}
	}
	return r, nil
}

func (r reporter) GenerateSpecReport(spec *models.Spec, specAnalyses map[analyzer.SpecAnalyzer]*models.SpecAnalysis) (*SpecReport, error) {
	return NewSpecReport(spec, specAnalyses).WithMitigation(r.analyzerRules).WithScore(r.scoreCfg, r.analyzerRules)
}

func (r reporter) GenerateSpecAnalysisReport(spec *models.Spec, specAnalysis *models.SpecAnalysis) (*SpecAnalysisReport, error) {
	return NewSpecAnalysisReport(specAnalysis, spec).WithMitigation(r.analyzerRules).WithScore(r.scoreCfg, r.analyzerRules)
}
