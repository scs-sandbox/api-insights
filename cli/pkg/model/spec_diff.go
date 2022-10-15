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
	"github.com/getkin/kin-openapi/openapi3"
	"io"
	"time"
)

const (
	DiffOutputJSON     = "json"
	DiffOutputHTML     = "html"
	DiffOutputMarkdown = "markdown"
	DiffOutputText     = "text"
)

// SpecDiff represents a specDiff
type SpecDiff struct {
	ID string `json:"id,omitempty" validate:"required"`
	*SpecDiffRequest
	Result    *SpecDiffResult `json:"result"`
	ServiceID string          `json:"service_id,omitempty"`
	Status    string          `json:"status,omitempty"` // Submitted, Invalid, Diffed
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (m *SpecDiff) HasBreakingChangesInJSON() bool {
	if m.Result == nil || m.Result.JSON == nil {
		return false
	}

	return m.Result.JSON.Breaking
}

type SpecDiffList []*SpecDiff

func (m SpecDiffList) Print(w io.Writer) {
	table := utils.NewTable(w)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"#", "ID", "New Spec ID", "Old Spec ID", "Service ID", "Status"})

	for i, diff := range m {
		row := []string{fmt.Sprintf("%d", i+1), diff.ID, diff.NewSpecID, diff.OldSpecID, diff.ServiceID, diff.Status}
		table.Append(row)
	}

	table.Render()
}

// SpecDiffConfig represents configs for a SpecDiff (SpecDiff.Config)
type SpecDiffConfig struct {
	OutputFormat string `json:"output_format,omitempty"` // json, html, markdown, text
}

type SpecDiffRequest struct {
	NewSpecID string `json:"new_spec_id,omitempty"`
	OldSpecID string `json:"old_spec_id,omitempty"`

	OldSpecDoc SpecDoc `json:"old_spec_doc,omitempty"`
	NewSpecDoc SpecDoc `json:"new_spec_doc,omitempty"`

	Config *SpecDiffConfig `json:"config,omitempty"`
}

// SpecDiffResult represents the result for a models.SpecDiff (SpecDiff.JSONResult)
type SpecDiffResult struct {
	JSON     *JSONResult `json:"json,omitempty"`
	HTML     string      `json:"html,omitempty"`
	Markdown string      `json:"markdown,omitempty"`
	Text     string      `json:"text,omitempty"`
}

func (m *SpecDiffResult) Print(w io.Writer, output string) {
	var res interface{}
	switch output {
	case DiffOutputJSON:
		res = utils.Pretty(m.JSON)
	case DiffOutputMarkdown:
		res = m.Markdown
	case DiffOutputHTML:
		res = m.HTML
	case DiffOutputText:
		res = m.Text
	default:
		res = "unsupported output"
	}

	_, _ = fmt.Fprintf(w, "%v", res)
}

type Action string

const (
	ActionAdded    = Action("added")
	ActionDeleted  = Action("deleted")
	ActionModified = Action("modified")
)

type JSONResult struct {
	Added      []*EndpointSummary `json:"added"`
	Deleted    []*EndpointSummary `json:"deleted"`
	Deprecated []*EndpointSummary `json:"deprecated"`
	Modified   []*ModifiedSummary `json:"modified"`

	Breaking bool   `json:"breaking"`
	Message  string `json:"message"`
}

type EndpointSummary struct {
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`

	Message string `json:"message"`
}

type ModifiedSummary struct {
	Old *OperationSummary `json:"old"`
	New *OperationSummary `json:"new"`

	Path        string `json:"path"`
	Method      string `json:"method"`
	Summary     string `json:"summary"`
	Description string `json:"description"`

	Breaking bool   `json:"breaking"`
	Message  string `json:"message"`

	ParametersSummary  *ParametersSummary  `json:"parameters"`
	RequestBodySummary *RequestBodySummary `json:"requestBody"`
	ResponsesSummary   *ResponsesSummary   `json:"responses"`
	SecuritySummary    *SecuritySummary    `json:"security"`
}

type OperationSummary struct{ openapi3.Operation }

type (
	ParametersSummary struct {
		Breaking bool                `json:"breaking"`
		Message  string              `json:"message"`
		Details  []*ParameterSummary `json:"details"`
	}
	ParameterSummary struct {
		Parameter, OldParameter, NewParameter *openapi3.Parameter `json:"-"`

		Name        string `json:"name"`
		In          string `json:"in"`
		Description string `json:"description"`
		Deprecated  bool   `json:"deprecated"` // TODO
		Breaking    bool   `json:"breaking"`
		Action      Action `json:"action"`
		Message     string `json:"message"`
	}
)

type (
	RequestBodySummary struct {
		Breaking bool   `json:"breaking"`
		Message  string `json:"message"`

		Description string `json:"description"`

		Details []*RequestBodySummaryDetail `json:"details"`
	}
	RequestBodySummaryDetail struct {
		ReqBody, OldReqBody, NewReqBody *openapi3.RequestBody `json:"-"`

		Properties []*PropertiesSummary `json:"properties"`

		Breaking bool   `json:"breaking"`
		Action   Action `json:"action"`
		Message  string `json:"message"`
		Name     string `json:"name"`
	}
)

type (
	ResponsesSummary struct {
		Breaking bool                      `json:"breaking"`
		Message  string                    `json:"message"`
		Details  []*ResponsesSummaryDetail `json:"details"`
	}
	ResponsesSummaryDetail struct {
		Res, OldRes, NewRes *openapi3.Response `json:"-"`

		Details []*ResponseSummaryDetail `json:"details"`

		Name        string `json:"name"` // status code
		Description string `json:"description"`
		Action      Action `json:"action"`
		Breaking    bool   `json:"breaking"`
		Message     string `json:"message"`
	}
	ResponseSummaryDetail struct {
		Res *openapi3.Response `json:"-"`

		Description string               `json:"description"`
		Name        string               `json:"name"`
		Action      Action               `json:"action"`
		Breaking    bool                 `json:"breaking"`
		Message     string               `json:"message"`
		Properties  []*PropertiesSummary `json:"properties"`
	}
)

type (
	SecuritySummary struct {
		Breaking bool                     `json:"breaking"`
		Message  string                   `json:"message"`
		Details  []*SecuritySummaryDetail `json:"details"`
	}
	SecuritySummaryDetail struct {
		SecReq, OldSecReq, NewSecReq *openapi3.SecurityRequirement `json:"-"`

		Breaking bool   `json:"breaking"`
		Name     string `json:"name"`
		Action   Action `json:"action"`
		Message  string `json:"message"`
	}
)

type PropertiesSummary struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Action      Action `json:"action"`
	Breaking    bool   `json:"breaking"`
	Message     string `json:"message"`

	Nested []*PropertiesSummary `json:"properties"`

	Group string `json:"-"` // group that this property belongs to, e.g. items
}
