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
	"github.com/cisco-developer/api-insights/api/pkg/utils"
)

var _ Resulter = (*SpectralResult)(nil)

type SpectralConfig struct {
	Ruleset *string // choose a ruleset [string] [choices: "cisco", "cisco-without-oas", "cx", "cx-without-oas", "cx-dictionary"] [default: "cisco-without-oas"]
}

func (c *SpectralConfig) SetDefaults() {
	if c.Ruleset == nil || *c.Ruleset == "" {
		c.Ruleset = utils.StringPtr("cisco-without-oas")
	}
}

func (c *SpectralConfig) SetRuleset(ruleset string) {
	if len(ruleset) > 0 {
		c.Ruleset = utils.StringPtr(ruleset)
	}
}

type (
	SpectralResult     []*SpectralResultItem
	SpectralResultItem struct {
		Code     string   `json:"code"`
		Path     []string `json:"path"`
		Message  string   `json:"message"`
		Severity int      `json:"severity"`
		Range    struct {
			Start struct {
				Line      int `json:"line"`
				Character int `json:"character"`
			} `json:"start"`
			End struct {
				Line      int `json:"line"`
				Character int `json:"character"`
			} `json:"end"`
		} `json:"range"`
		Source string `json:"source"`
	}
)

func (m SpectralResult) Result() (*Result, error) {
	result := NewResult()
	if m == nil {
		return result, nil
	}
	for _, r := range m {
		ruleNameID := rule.NameID(r.Code)
		severity := spectralSeverityName(r.Severity)
		result.storeRuleInCache(severity, ruleNameID, &Rule{
			NameID:         r.Code,
			AnalyzerNameID: string(CiscoAPIGuidelines),
			Title:          r.Code,
			Description:    r.Message,
			Severity:       severity.String(),
			Mitigation:     "",
		})
		result.AddFinding(severity, ruleNameID, &Finding{
			Type: rule.FindingTypeRange,
			Path: r.Path,
			Range: &FindingPositionRange{
				Start: &FindingPosition{
					Line:   r.Range.Start.Line + 1,
					Column: r.Range.Start.Character + 1,
				},
				End: &FindingPosition{
					Line:   r.Range.Start.Line + 1,
					Column: r.Range.Start.Character + 1,
				},
			},
		})
	}
	return result, nil
}

func spectralSeverityName(spectralSeverity int) rule.SeverityName {
	switch spectralSeverity {
	case 0: // Error
		return rule.SeverityNameError
	case 1: // Warn
		return rule.SeverityNameWarning
	case 2: // Info
		return rule.SeverityNameInfo
	case 3: // Hint
		return rule.SeverityNameHint
	}
	return rule.SeverityNameHint
}
