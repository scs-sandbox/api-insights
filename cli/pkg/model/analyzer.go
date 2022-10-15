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

// Analyzer represents an analyzer
type Analyzer struct {
	ID          string                 `json:"id,omitempty"`
	NameID      string                 `json:"name_id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Status      string                 `json:"status"`
	Meta        map[string]interface{} `json:"meta"`
	Position    uint8                  `json:"position"`
	Config      Config                 `json:"config"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`

	Rules []AnalyzerRule `json:"rules"`
}

type AnalyzerList []*Analyzer

func (m AnalyzerList) Print(w io.Writer) {
	table := utils.NewTable(w)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"#", "ID", "Name ID", "Title", "Description", "Position", "Status"})

	for i, analyzer := range m {
		row := []string{fmt.Sprintf("%d", i+1), analyzer.ID, analyzer.NameID, analyzer.Title, analyzer.Description, fmt.Sprintf("%v", analyzer.Position), analyzer.Status}
		table.Append(row)
	}

	table.Render()
}
