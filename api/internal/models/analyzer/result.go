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
)

// Result represents the result of a models.SpecAnalysis.
type Result struct {
	Summary  *ResultSummary       `json:"summary" gorm:"column:summary"`
	Findings SeverityRuleFindings `json:"findings" gorm:"column:findings"`

	severityRuleCache map[rule.SeverityName]map[rule.NameID]*Rule `gorm:"column:-"`
}

// NewResult constructs a new Result.
func NewResult() *Result {
	return &Result{
		Summary:           NewResultSummary(),
		Findings:          NewSeverityRuleFindings(),
		severityRuleCache: newSeverityRuleCache(),
	}
}

func (r Result) AddFinding(severity rule.SeverityName, ruleNameID rule.NameID, finding *Finding) {
	analyzerRule := r.getRuleFromCache(severity, ruleNameID)
	if analyzerRule != nil {
		severityRuleFindings, ok := r.Findings[severity]
		if !ok {
			r.Findings[severity] = &RuleFindings{
				Rules: map[rule.NameID]*Findings{},
			}
		}
		if severityRuleFindings.Rules == nil {
			severityRuleFindings.Rules = map[rule.NameID]*Findings{}
		}
		if _, ok := severityRuleFindings.Rules[ruleNameID]; !ok {
			severityRuleFindings.Rules[ruleNameID] = &Findings{
				Message:    analyzerRule.Description,
				Mitigation: analyzerRule.Mitigation,
			}
		}
		severityRuleFindings.Rules[ruleNameID].Data = append(severityRuleFindings.Rules[ruleNameID].Data, finding)
	}

	r.updateSummaryStatsAfterAddFinding(severity, ruleNameID)
}

func (r Result) updateSummaryStatsAfterAddFinding(severity rule.SeverityName, ruleNameID rule.NameID) {
	if r.Summary == nil {
		r.Summary = NewResultSummary()
	}
	switch severity {
	case rule.SeverityNameHint:
		r.Summary.Stats.Hint.Occurrences++
		r.Summary.Stats.Hint.Data[ruleNameID]++
		r.Summary.Stats.Hint.Count = len(r.Summary.Stats.Hint.Data)
	case rule.SeverityNameInfo:
		r.Summary.Stats.Info.Occurrences++
		r.Summary.Stats.Info.Data[ruleNameID]++
		r.Summary.Stats.Info.Count = len(r.Summary.Stats.Info.Data)
	case rule.SeverityNameWarning:
		r.Summary.Stats.Warning.Occurrences++
		r.Summary.Stats.Warning.Data[ruleNameID]++
		r.Summary.Stats.Warning.Count = len(r.Summary.Stats.Warning.Data)
	case rule.SeverityNameError:
		r.Summary.Stats.Error.Occurrences++
		r.Summary.Stats.Error.Data[ruleNameID]++
		r.Summary.Stats.Error.Count = len(r.Summary.Stats.Error.Data)
	}
	r.Summary.Stats.Occurrences++
	r.Summary.Stats.Count = r.Summary.Stats.Hint.Count + r.Summary.Stats.Info.Count + r.Summary.Stats.Warning.Count + r.Summary.Stats.Error.Count
}

func (r Result) storeRuleInCache(severity rule.SeverityName, ruleNameID rule.NameID, analyzerRule *Rule) {
	if _, ok := r.severityRuleCache[severity]; !ok {
		r.severityRuleCache[severity] = map[rule.NameID]*Rule{}
	}
	if _, ok := r.severityRuleCache[severity][ruleNameID]; !ok {
		r.severityRuleCache[severity][ruleNameID] = analyzerRule
	}
}

func (r Result) getRuleFromCache(severity rule.SeverityName, ruleNameID rule.NameID) (_ *Rule) {
	rules, ok := r.severityRuleCache[severity]
	if ok {
		return rules[ruleNameID]
	}
	return
}

// NewResultSummary constructs a new ResultSummary with default stats initialized.
func NewResultSummary() *ResultSummary {
	return &ResultSummary{
		Stats: &SeverityRuleFindingsStats{
			Hint:    &RuleFindingsStats{Data: map[rule.NameID]int{}},
			Info:    &RuleFindingsStats{Data: map[rule.NameID]int{}},
			Warning: &RuleFindingsStats{Data: map[rule.NameID]int{}},
			Error:   &RuleFindingsStats{Data: map[rule.NameID]int{}},
		},
	}
}

// NewSeverityRuleFindings constructs a new SeverityRuleFindings with default severities initialized.
func NewSeverityRuleFindings() SeverityRuleFindings {
	return SeverityRuleFindings{
		rule.SeverityNameHint:    &RuleFindings{},
		rule.SeverityNameInfo:    &RuleFindings{},
		rule.SeverityNameWarning: &RuleFindings{},
		rule.SeverityNameError:   &RuleFindings{},
	}
}

func newSeverityRuleCache() map[rule.SeverityName]map[rule.NameID]*Rule {
	return map[rule.SeverityName]map[rule.NameID]*Rule{
		rule.SeverityNameHint:    {},
		rule.SeverityNameInfo:    {},
		rule.SeverityNameWarning: {},
		rule.SeverityNameError:   {},
	}
}

type (
	// ResultSummary represents a summary of Result.Findings.
	ResultSummary struct {
		Stats *SeverityRuleFindingsStats `json:"stats"`
	}
	// SeverityRuleFindingsStats contains stats of SeverityRuleFindings.
	SeverityRuleFindingsStats struct {
		Count       int                `json:"count"`
		Occurrences int                `json:"occurrences"`
		Hint        *RuleFindingsStats `json:"hint"`
		Info        *RuleFindingsStats `json:"info"`
		Warning     *RuleFindingsStats `json:"warning"`
		Error       *RuleFindingsStats `json:"error"`
	}
	// RuleFindingsStats contains stats of RuleFindings.
	RuleFindingsStats struct {
		Count       int                 `json:"count"`
		Occurrences int                 `json:"occurrences"`
		Data        map[rule.NameID]int `json:"rules"`
	}
)

// SeverityRuleFindings defines a grouping of rule findings by their severities.
type SeverityRuleFindings map[rule.SeverityName]*RuleFindings

type RuleFindings struct {
	Rules map[rule.NameID]*Findings `json:"rules"`
}

type Findings struct {
	Message    string     `json:"message"`
	Mitigation string     `json:"mitigation"`
	Data       []*Finding `json:"data"`
}

type Finding struct {
	Type  rule.FindingType      `json:"type"`
	Path  []string              `json:"path"`
	Range *FindingPositionRange `json:"range,omitempty"`
	Diff  *FindingDiff          `json:"diff,omitempty"`
}

type (
	// FindingPositionRange represents
	FindingPositionRange struct {
		Start *FindingPosition `json:"start"`
		End   *FindingPosition `json:"end"`
	}
	FindingPosition struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	}
)

type FindingDiff struct {
	Old string `json:"old"`
	New string `json:"new"`
}
