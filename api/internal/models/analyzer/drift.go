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
	models2 "github.com/cisco-developer/api-insights/api/pkg/apiclarity/models"
	"strings"
)

var _ Resulter = (*APIClarityDriftResult)(nil)

type APIClarityConfig struct {
}

type APIClarityDriftResult struct {
	Events []*models2.APIEvent

	EventProvidedSpecDiffs []*models2.APIEventSpecDiff
}

func (m *APIClarityDriftResult) Result() (*Result, error) {
	result := NewResult()
	if m == nil {
		return result, nil
	}
	for i, r := range m.Events {
		ruleNameID := rule.NameID(*r.SpecDiffType)
		severity := apiClarityDiffTypeToSeverity(*r.SpecDiffType)
		result.storeRuleInCache(severity, ruleNameID, &Rule{
			NameID:         string(*r.SpecDiffType),
			AnalyzerNameID: string(InclusiveLanguage),
			Title:          string(*r.SpecDiffType),
			Description:    apiClarityDiffTypeToMessage(*r.SpecDiffType),
			Severity:       severity.String(),
			Mitigation:     "",
		})
		var (
			oldSpec, newSpec string
		)
		if m.EventProvidedSpecDiffs[i] != nil {
			if m.EventProvidedSpecDiffs[i].OldSpec != nil {
				oldSpec = *m.EventProvidedSpecDiffs[i].OldSpec
			}
			if m.EventProvidedSpecDiffs[i].NewSpec != nil {
				newSpec = *m.EventProvidedSpecDiffs[i].NewSpec
			}
		}
		result.AddFinding(severity, ruleNameID, &Finding{
			Type: rule.FindingTypeDiff,
			Path: apiclarityAPIEventToPath(r),
			Diff: &FindingDiff{
				Old: oldSpec,
				New: newSpec,
			},
		})
	}
	return result, nil
}

func apiClarityDiffTypeToSeverity(diffType models2.DiffType) rule.SeverityName {
	switch diffType {
	case models2.DiffTypeGENERALDIFF:
		return rule.SeverityNameError
	case models2.DiffTypeSHADOWDIFF:
		return rule.SeverityNameError
	case models2.DiffTypeZOMBIEDIFF:
		return rule.SeverityNameError
	}
	return rule.SeverityNameHint
}

func apiClarityDiffTypeToMessage(diffType models2.DiffType) string {
	switch diffType {
	case models2.DiffTypeGENERALDIFF:
		return "General diff: a general diff has been detected"
	case models2.DiffTypeSHADOWDIFF:
		return "Shadow: an undocumented API has been detected"
	case models2.DiffTypeZOMBIEDIFF:
		return "Zombie: a deprecated API has been detected"
	}
	return ""
}

func apiclarityAPIEventToPath(apiEvent *models2.APIEvent) []string {
	return []string{
		"paths",
		apiEvent.Path,
		strings.ToLower(string(apiEvent.Method)),
	}
}
