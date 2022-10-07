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
	wokerule "github.com/get-woke/woke/pkg/rule"
)

var _ Resulter = (*WokeResult)(nil)

type WokeConfig struct {
	Config              string // Config file (default is .woke.yaml in current directory, or $HOME)
	DisableDefaultRules bool   // Disable the default ruleset
	ExitOneOnFailure    bool   // Exit with exit code 1 on failures
	NoIgnore            bool   // Ignored files in .gitignore, .ignore, .wokeignore, .git/info/exclude, and inline ignores are processed
	OutputName          string // Output type [text,simple,github-actions,json,sonarqube]
}

func (c *WokeConfig) SetDefaults() {
	if c.OutputName == "" {
		c.OutputName = "json"
	}
}

type WokeResult struct {
	Filename string `json:"Filename"`
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
	} `json:"Results"`
}

func (m *WokeResult) Result() (*Result, error) {
	result := NewResult()
	if m == nil {
		return result, nil
	}
	for _, r := range m.Results {
		ruleNameID := rule.NameID(r.Finding)
		severity := wokeSeverityName(r.Rule.Severity)
		result.storeRuleInCache(severity, ruleNameID, &Rule{
			NameID:         r.Finding,
			AnalyzerNameID: string(InclusiveLanguage),
			Title:          r.Finding,
			Description:    r.Finding,
			Severity:       severity.String(),
			Mitigation:     r.Reason,
		})
		result.AddFinding(severity, ruleNameID, &Finding{
			Type: rule.FindingTypeRange,
			Path: []string{},
			Range: &FindingPositionRange{
				Start: &FindingPosition{
					Line:   r.StartPosition.Line,
					Column: r.StartPosition.Column,
				},
				End: &FindingPosition{
					Line:   r.EndPosition.Line,
					Column: r.EndPosition.Column,
				},
			},
		})
	}
	return result, nil
}

func wokeSeverityName(s string) rule.SeverityName {
	switch wokerule.NewSeverity(s) {
	case wokerule.SevError:
		return rule.SeverityNameError
	case wokerule.SevWarn:
		return rule.SeverityNameWarning
	case wokerule.SevInfo:
		return rule.SeverityNameInfo
	}
	return rule.SeverityNameHint
}
