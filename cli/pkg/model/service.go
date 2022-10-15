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

// Service represents a service
type Service struct {
	ID             string                 `json:"id,omitempty"`
	AdditionalInfo map[string]interface{} `json:"additional_info"`
	Contact        interface{}            `json:"contact"`
	Description    string                 `json:"description"`
	NameID         string                 `json:"name_id"`
	OrganizationID string                 `json:"organization_id"`
	ProductTag     string                 `json:"product_tag"`
	Title          string                 `json:"title"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`

	AnalyzersConfigs AnalyzerConfigMap `json:"analyzers_configs,omitempty"`
	Summary          *ServiceSummary   `json:"summary"`
}

type Config map[string]interface{}

func (c *Config) String() string {
	return utils.Pretty(c)
}

type AnalyzerConfigMap map[SpecAnalyzer]Config

type ServiceSummary struct {
	Score     *int      `json:"score"`
	Version   string    `json:"version"`
	Revision  string    `json:"revision"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ServiceList []*Service

func (m ServiceList) Print(w io.Writer) {
	table := utils.NewTable(w)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"#", "ID", "Name ID", "Name", "Description", "Organization", "Product"})

	for i, service := range m {
		row := []string{fmt.Sprintf("%d", i+1), service.ID, service.NameID, service.Title, service.Description, service.OrganizationID, service.ProductTag}
		table.Append(row)
	}

	table.Render()
}
