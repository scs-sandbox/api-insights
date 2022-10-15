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
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer/rule"
	"github.com/cisco-developer/api-insights/api/pkg/panoptica/models"
	"github.com/cisco-developer/api-insights/api/pkg/utils/speciterator"
	"strings"
)

const (
	ScoreCategoryAPISecurity = "api-specification"

	SACSeverityCritical = "CRITICAL"
	SACSeverityHigh     = "HIGH"
	SACSeverityMedium   = "MEDIUM"
	SACSeverityLow      = "LOW"
	SACSeverityUnknown  = "UNKNOWN"
)

func securityResultSeverityName(severity string) rule.SeverityName {
	severity = strings.ToUpper(strings.TrimSpace(severity))
	switch severity {
	case SACSeverityCritical:
		return rule.SeverityNameError
	case SACSeverityHigh:
		return rule.SeverityNameError
	case SACSeverityMedium:
		return rule.SeverityNameWarning
	case SACSeverityLow:
		return rule.SeverityNameInfo
	}
	return rule.SeverityNameHint
}

type SecurityConfig struct {
	Name string // composed sac external api name, e.g. carts.api.apiregistry
}

type SecurityFinding struct {
	Severity           string        `json:"severity"`
	Kind               string        `json:"kind"`
	Type               string        `json:"type"`
	Code               string        `json:"code"`
	Message            string        `json:"message"`
	Location           []interface{} `json:"location"`
	CrRawFindingID     string        `json:"cr_raw_finding_id"`
	CrFindingIndex     int           `json:"cr_finding_index"`
	AffectedEndpoints  []interface{} `json:"affected_endpoints"`
	Source             string        `json:"source"`
	SeverityCategory   string        `json:"severity_category"`
	CrankshaftClassID  string        `json:"crankshaft_class_id"`
	CrankshaftSeverity string        `json:"crankshaft_severity"`
	CrankshaftCategory string        `json:"crankshaft_category"`
	CrankshaftJsonpath string        `json:"crankshaft_jsonpath"`
}

func (m *SecurityFinding) JSONPaths() []string {
	s := make([]string, 0, len(m.Location))

	for _, l := range m.Location {
		s = append(s, fmt.Sprintf("%v", l))
	}

	return s
}

func NewSecurityFindingData(sf *models.ScoreFinding) (data []*SecurityFinding, err error) {
	if sf == nil {
		return
	}

	b, err := json.Marshal(sf.Data)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return
	}

	return
}

func GetSecurityResult(spec string, in *models.APIServiceDrillDownExternal) (*Result, error) {
	result := NewResult()
	if in == nil || in.Score == nil || in.Score.API == nil || in.Score.API.Categories == nil {
		return result, nil
	}

	c := in.Score.API.Categories[ScoreCategoryAPISecurity]
	fgs := []*models.ScoreFindingGroup{c.Critical, c.High, c.Medium, c.Low, c.Unclassified}

	var possByPaths = map[string]*speciterator.Pos{}
	var itemVisited = func(path *speciterator.Path, pos *speciterator.Pos) {
		possByPaths[path.String()] = pos
	}
	if len(fgs) > 0 {
		si := speciterator.NewSpecIterator([]byte(spec))
		_ = si.Iterate(itemVisited)
	}

	for _, fg := range fgs {
		if fg == nil {
			continue
		}
		for _, finding := range fg.Findings {
			items, err := NewSecurityFindingData(finding)
			if err != nil {
				return nil, err
			}

			for _, item := range items {
				// TODO: remove
				if isStringLoosePattern := item.Code == "string-loosepattern" || item.Code == "TDT008"; isStringLoosePattern {
					continue
				}

				ruleNameID := rule.NameID(item.Code)
				severity := securityResultSeverityName(item.CrankshaftSeverity)

				d := finding.Description
				if finding.Name != nil {
					d = fmt.Sprintf("%s\n%s", *finding.Name, d)
				}

				result.storeRuleInCache(severity, ruleNameID, &Rule{
					NameID:         item.Code,
					AnalyzerNameID: string(Security),
					Title:          item.Code,
					Description:    d,
					Severity:       severity.String(),
					Mitigation:     finding.Mitigation,
				})

				p := item.JSONPaths()

				line, column := 1, 1
				if pos, found := possByPaths[strings.Join(p, "|")]; found && pos != nil {
					line = pos.Line
					column = pos.Column
				}

				result.AddFinding(severity, ruleNameID, &Finding{
					Type: rule.FindingTypeRange,
					Path: p,
					Range: &FindingPositionRange{
						Start: &FindingPosition{Line: line, Column: column},
						End:   &FindingPosition{Line: line, Column: column},
					},
				})
			}
		}
	}

	return result, nil
}
