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

package model

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Result represents the result of a models.SpecAnalysis.
type Result struct {
	Summary  *ResultSummary       `json:"summary" gorm:"column:summary"`
	Findings SeverityRuleFindings `json:"findings" gorm:"column:findings"`

	severityRuleCache map[SeverityName]map[NameID]*AnalyzerRule `gorm:"column:-"`
}

// NewResult constructs a new Result.
func NewResult() *Result {
	return &Result{
		Summary:           NewResultSummary(),
		Findings:          NewSeverityRuleFindings(),
		severityRuleCache: newSeverityRuleCache(),
	}
}

// NewResultSummary constructs a new ResultSummary with default stats initialized.
func NewResultSummary() *ResultSummary {
	return &ResultSummary{
		Stats: &SeverityRuleFindingsStats{
			Hint:    &RuleFindingsStats{Data: map[NameID]int{}},
			Info:    &RuleFindingsStats{Data: map[NameID]int{}},
			Warning: &RuleFindingsStats{Data: map[NameID]int{}},
			Error:   &RuleFindingsStats{Data: map[NameID]int{}},
		},
	}
}

// NewSeverityRuleFindings constructs a new SeverityRuleFindings with default severities initialized.
func NewSeverityRuleFindings() SeverityRuleFindings {
	return SeverityRuleFindings{
		SeverityNameHint:    &RuleFindings{},
		SeverityNameInfo:    &RuleFindings{},
		SeverityNameWarning: &RuleFindings{},
		SeverityNameError:   &RuleFindings{},
	}
}

func newSeverityRuleCache() map[SeverityName]map[NameID]*AnalyzerRule {
	return map[SeverityName]map[NameID]*AnalyzerRule{
		SeverityNameHint:    {},
		SeverityNameInfo:    {},
		SeverityNameWarning: {},
		SeverityNameError:   {},
	}
}

// ResultSummary represents a summary of Result.Findings.
type ResultSummary struct {
	Stats *SeverityRuleFindingsStats `json:"stats"`
}

func (m *ResultSummary) String() string {
	s := m.Stats
	if s == nil {
		return ""
	}

	return fmt.Sprintf("%d Findings (%d Error, %d Warning, %d Info, %d Hint)",
		s.Count, s.TotalError(), s.TotalWarning(), s.TotalInfo(), s.TotalHint())
}

// SeverityRuleFindingsStats contains stats of SeverityRuleFindings.
type SeverityRuleFindingsStats struct {
	Count   int                `json:"count"`
	Hint    *RuleFindingsStats `json:"hint"`
	Info    *RuleFindingsStats `json:"info"`
	Warning *RuleFindingsStats `json:"warning"`
	Error   *RuleFindingsStats `json:"error"`
}

func (s *SeverityRuleFindingsStats) TotalError() int {
	if s.Error == nil {
		return 0
	}

	return s.Error.Count
}

func (s *SeverityRuleFindingsStats) TotalWarning() int {
	if s.Warning == nil {
		return 0
	}

	return s.Warning.Count
}

func (s *SeverityRuleFindingsStats) TotalInfo() int {
	if s.Info == nil {
		return 0
	}

	return s.Info.Count
}

func (s *SeverityRuleFindingsStats) TotalHint() int {
	if s.Hint == nil {
		return 0
	}

	return s.Hint.Count
}

// RuleFindingsStats contains stats of RuleFindings.
type RuleFindingsStats struct {
	Count int            `json:"count"`
	Data  map[NameID]int `json:"rules"`
}

// SeverityRuleFindings defines a grouping of rule findings by their severities.
type SeverityRuleFindings map[SeverityName]*RuleFindings

type RuleFindings struct {
	Rules map[NameID]*Findings `json:"rules"`
}

type Findings struct {
	Message    string     `json:"message"`
	Mitigation string     `json:"mitigation"`
	Data       []*Finding `json:"data"`
}

type Finding struct {
	Type  FindingType           `json:"type"`
	Path  []string              `json:"path"`
	Range *FindingPositionRange `json:"range,omitempty"`
	Diff  *FindingDiff          `json:"diff,omitempty"`
}

func (f *Finding) Start() string {
	return fmt.Sprintf("%v:%v", f.Range.Start.Line, f.Range.Start.Column)
}

func (f *Finding) End() string {
	return fmt.Sprintf("%v:%v", f.Range.End.Line, f.Range.End.Column)
}

func (f *Finding) JSONPath() string {
	if len(f.Path) < 2 {
		return strings.Join(f.Path, ".")
	}

	var b bytes.Buffer
	b.WriteString(f.Path[0])
	for j := 1; j < len(f.Path); j++ {
		if n, err := strconv.Atoi(f.Path[j]); err == nil {
			b.WriteString(fmt.Sprintf("[%v]", n))
		} else {
			b.WriteString(fmt.Sprintf(".%v", f.Path[j]))
		}
	}

	return b.String()
}

// FindingPositionRange represents
type FindingPositionRange struct {
	Start *FindingPosition `json:"start"`
	End   *FindingPosition `json:"end"`
}

type FindingPosition struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type FindingDiff struct {
	Old string `json:"old"`
	New string `json:"new"`
}

type NameID string

type FindingType string

const (
	FindingTypeRange FindingType = "range"
	FindingTypeDiff              = "diff"
)

type (
	Severity     int
	SeverityName string
)

const (
	SeverityHint Severity = iota + 1
	SeverityInfo
	SeverityWarning
	SeverityError

	SeverityNameHint    SeverityName = "hint"
	SeverityNameInfo    SeverityName = "info"
	SeverityNameWarning SeverityName = "warning"
	SeverityNameError   SeverityName = "error"
)

func (s Severity) Name() SeverityName { return severityNameBySeverity[s] }

func (s Severity) String() string { return string(s.Name()) }

func (s Severity) Weight() int { return int(s) }

func (n SeverityName) String() string { return string(n) }

func (n SeverityName) Severity() Severity {
	switch n {
	case SeverityHint.Name():
		return SeverityHint
	case SeverityInfo.Name():
		return SeverityInfo
	case SeverityWarning.Name():
		return SeverityWarning
	case SeverityError.Name():
		return SeverityError
	}
	return SeverityHint
}

var severityNameBySeverity = map[Severity]SeverityName{
	SeverityHint:    SeverityNameHint,
	SeverityInfo:    SeverityNameInfo,
	SeverityWarning: SeverityNameWarning,
	SeverityError:   SeverityNameError,
}
