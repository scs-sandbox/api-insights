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

package diff

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"strconv"
	"strings"
)

const (
	h4         = "#### "
	h5         = "##### "
	h6         = "###### "
	blockquote = "> "
	code       = "`"
	indent     = "    "
)

var (
	markdownHeading4   = func(s string) string { return h4 + s + "\n" }
	markdownHeading5   = func(s string) string { return h5 + s + "\n" }
	markdownHeading6   = func(s string) string { return h6 + s + "\n" }
	markdownCode       = func(s string) string { return code + s + code }
	markdownBlockquote = func(prefix, s string) string {
		if s == "" {
			return ""
		}
		blockquote := prefix + blockquote
		return blockquote + strings.ReplaceAll(strings.TrimSpace(s), "\n", "\n"+blockquote) + "\n"
	}
	markdownIndent = func(depth int) string {
		var sb strings.Builder
		for i := 0; i < depth; i++ {
			sb.WriteString(indent)
		}
		return sb.String()
	}
)

func NewMarkdownSummaryMessageBuilder() SummaryMessageBuilder {
	return &markdownSummaryMessageBuilder{}
}

type markdownSummaryMessageBuilder struct{}

func (m markdownSummaryMessageBuilder) BuildResultSummaryMessage(result *JSONResult) string {
	var sb strings.Builder

	var endpointHeading = func(method, path, description string) string {
		return markdownHeading5(fmt.Sprintf("%s %s\n\n%s", markdownCode(method), path, markdownBlockquote("", description)))
	}

	if len(result.Added) > 0 {
		sb.WriteString(markdownHeading4("What's New"))
		sb.WriteString("\n")
		for _, e := range result.Added {
			sb.WriteString(endpointHeading(e.Method, e.Path, e.Description))
		}
	}

	if len(result.Deleted) > 0 {
		sb.WriteString(markdownHeading4("What's Deleted"))
		sb.WriteString("\n")
		for _, e := range result.Deleted {
			sb.WriteString(endpointHeading(e.Method, e.Path, e.Description))
		}
	}

	if len(result.Deprecated) > 0 {
		sb.WriteString(markdownHeading4("What's Deprecated"))
		sb.WriteString("\n")
		for _, e := range result.Deprecated {
			sb.WriteString(endpointHeading(e.Method, e.Path, e.Description))
		}
	}

	if len(result.Modified) > 0 {
		sb.WriteString(markdownHeading4("What's Modified"))
		sb.WriteString("\n")
		for _, m := range result.Modified {
			sb.WriteString(endpointHeading(m.Method, m.Path, m.Summary))
			sb.WriteString(m.Message)
		}
	}

	return sb.String()
}

func (m markdownSummaryMessageBuilder) BuildModifiedSummaryMessage(s *ModifiedSummary) string {
	var sb strings.Builder

	if s.ParametersSummary != nil {
		sb.WriteString(s.ParametersSummary.Message)
	}

	if s.RequestBodySummary != nil {
		sb.WriteString(s.RequestBodySummary.Message)
	}

	if s.ResponsesSummary != nil {
		sb.WriteString(s.ResponsesSummary.Message)
	}

	if s.SecuritySummary != nil {
		sb.WriteString(s.SecuritySummary.Message)
	}

	sb.WriteString("\n")

	return sb.String()
}

func (m markdownSummaryMessageBuilder) BuildParametersSummaryMessage(s *ParametersSummary) string {
	var sb strings.Builder
	sb.WriteString(markdownHeading6("Parameters:"))
	sb.WriteString("\n")
	for _, detail := range s.Details {
		sb.WriteString(detail.Message)
	}
	return sb.String()
}

func (m markdownSummaryMessageBuilder) BuildParameterSummaryMessage(s *ParameterSummary) string {
	return fmt.Sprintf("%s: %s in %s\n%s\n",
		cases.Title(language.Und, cases.NoLower).String(string(s.Action)),
		markdownCode(s.Name),
		markdownCode(s.In),
		markdownBlockquote("", s.Description),
	)
}

func (m markdownSummaryMessageBuilder) BuildRequestBodySummaryMessage(s *RequestBodySummary) string {
	var sb strings.Builder
	sb.WriteString(markdownHeading6("Request:"))
	sb.WriteString("\n")
	for _, detail := range s.Details {
		sb.WriteString(detail.Message)
	}
	return sb.String()
}

func (m markdownSummaryMessageBuilder) BuildRequestBodySummaryDetailMessage(d *RequestBodySummaryDetail, contentType string) string {
	var sb strings.Builder
	switch d.Action {
	case ActionAdded:
		sb.WriteString(fmt.Sprintf("Added content type: `%s`\n\n", contentType))
	case ActionDeleted:
		sb.WriteString(fmt.Sprintf("Deleted content type: `%s`\n\n", contentType))
	case ActionModified:
		sb.WriteString(fmt.Sprintf("Modified content type: `%s`\n\n", contentType))
		for _, p := range d.Properties {
			if p.Group == "items" {
				sb.WriteString(fmt.Sprintf("* Modified %s (%s):\n\n", p.Group, p.Type))
			}
			sb.WriteString(p.Message)
			for _, nested := range p.Nested {
				sb.WriteString(nested.Message)
			}
		}
	}
	return sb.String()
}

func (m markdownSummaryMessageBuilder) BuildResponsesSummaryMessage(s *ResponsesSummary) string {
	var sb strings.Builder
	sb.WriteString(markdownHeading6("Response:"))
	sb.WriteString("\n")
	for _, detail := range s.Details {
		sb.WriteString(detail.Message)
	}
	return sb.String()
}

func (m markdownSummaryMessageBuilder) BuildResponsesSummaryDetailMessage(d *ResponsesSummaryDetail, statusCode string) string {
	statusCodeInt, _ := strconv.Atoi(statusCode)
	var sb strings.Builder
	switch d.Action {
	case ActionAdded:
		sb.WriteString(fmt.Sprintf("Added response: **%s %s**\n%s", statusCode, http.StatusText(statusCodeInt), markdownBlockquote("", d.Description)))
	case ActionDeleted:
		sb.WriteString(fmt.Sprintf("Deleted response: **%s %s**\n%s", statusCode, http.StatusText(statusCodeInt), markdownBlockquote("", d.Description)))
	case ActionModified:
		sb.WriteString(fmt.Sprintf("Modified response: **%s %s**\n%s", statusCode, http.StatusText(statusCodeInt), markdownBlockquote("", d.Description)))
		for _, d := range d.Details {
			sb.WriteString(d.Message)
		}
	}
	return sb.String()
}

func (m markdownSummaryMessageBuilder) BuildResponseSummaryDetailMessage(s *ResponseSummaryDetail, key, value string) string {
	var sb strings.Builder
	sb.WriteString("\n")
	switch s.Action {
	case ActionAdded:
		sb.WriteString(fmt.Sprintf("* Added %s: `%s`\n", key, value))
	case ActionDeleted:
		sb.WriteString(fmt.Sprintf("* Deleted %s: `%s`\n", key, value))
	case ActionModified:
		sb.WriteString(fmt.Sprintf("* Modified %s: `%s`\n", key, value))
		for _, p := range s.Properties {
			if p.Group == "items" {
				sb.WriteString(fmt.Sprintf("%s* Modified %s (%s):\n\n", markdownIndent(1), p.Group, p.Type))
			}
			sb.WriteString(p.Message)
			for _, nested := range p.Nested {
				sb.WriteString(nested.Message)
			}
		}
	}
	return sb.String()
}

func (m markdownSummaryMessageBuilder) BuildSecuritySummaryMessage(s *SecuritySummary) string {
	var sb strings.Builder
	sb.WriteString(markdownHeading6("Security:"))
	sb.WriteString("\n")
	for _, detail := range s.Details {
		sb.WriteString(detail.Message)
	}
	return sb.String()
}

func (m markdownSummaryMessageBuilder) BuildSecuritySummaryDetailMessage(d *SecuritySummaryDetail) string {
	return fmt.Sprintf("%s authentication: %s\n",
		cases.Title(language.Und, cases.NoLower).String(string(d.Action)),
		markdownCode(d.Name),
	)
}

func (m markdownSummaryMessageBuilder) BuildPropertiesSummaryMessage(s *PropertiesSummary, indentLevel int) string {
	switch s.Action {
	case ActionAdded:
		return fmt.Sprintf("%s* Added property `%s` (%s)\n%s\n",
			markdownIndent(indentLevel),
			s.Name,
			s.Type,
			markdownBlockquote(markdownIndent(1+indentLevel), s.Description),
		)
	case ActionDeleted:
		return fmt.Sprintf("%s* Deleted property `%s` (%s)\n%s\n",
			markdownIndent(indentLevel),
			s.Name,
			s.Type,
			markdownBlockquote(markdownIndent(1+indentLevel), s.Description),
		)
	case ActionModified:
		return fmt.Sprintf("%s* Modified property `%s` (%s)\n%s\n",
			markdownIndent(indentLevel),
			s.Name,
			s.Type,
			markdownBlockquote(markdownIndent(1+indentLevel), s.Description),
		)
	}
	return ""
}
