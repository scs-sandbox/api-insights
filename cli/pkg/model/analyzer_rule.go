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
	"fmt"
	"github.com/cisco-developer/api-insights/cli/pkg/utils"
	"io"
	"time"
)

// AnalyzerRule represents an analyzer rule
type AnalyzerRule struct {
	ID             string                 `json:"id,omitempty"`
	NameID         string                 `json:"name_id"`
	AnalyzerNameID string                 `json:"analyzer_name_id"`
	Title          string                 `json:"title"`
	Description    string                 `json:"description"`
	Severity       string                 `json:"severity"`
	Mitigation     string                 `json:"mitigation"`
	Meta           map[string]interface{} `json:"meta"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
}

type AnalyzerRuleList []*AnalyzerRule

func (m AnalyzerRuleList) Print(w io.Writer) {
	table := utils.NewTable(w)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"#", "ID", "Name ID", "Analyzer Name ID", "Severity", "Title", "Description", "Mitigation"})

	for i, rule := range m {
		row := []string{fmt.Sprintf("%d", i+1), rule.ID, rule.NameID, rule.AnalyzerNameID, rule.Severity, rule.Title, rule.Description, rule.Mitigation}
		table.Append(row)
	}

	table.Render()
}
