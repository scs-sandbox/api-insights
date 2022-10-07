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
	"reflect"
	"testing"
)

func TestNewMarkdownSummaryMessageBuilder(t *testing.T) {
	tests := []struct {
		name string
		want SummaryMessageBuilder
	}{
		{
			name: "normal",
			want: &markdownSummaryMessageBuilder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMarkdownSummaryMessageBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMarkdownSummaryMessageBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markdownSummaryMessageBuilder_BuildModifiedSummaryMessage(t *testing.T) {
	type args struct {
		s *ModifiedSummary
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{s: &ModifiedSummary{
				ParametersSummary:  &ParametersSummary{Message: "s1"},
				RequestBodySummary: &RequestBodySummary{Message: "s2"},
				ResponsesSummary:   &ResponsesSummary{Message: "s3"},
				SecuritySummary:    &SecuritySummary{Message: "s4"},
			}},
			want: "s1s2s3s4\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := markdownSummaryMessageBuilder{}
			if got := m.BuildModifiedSummaryMessage(tt.args.s); got != tt.want {
				t.Errorf("BuildModifiedSummaryMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markdownSummaryMessageBuilder_BuildParameterSummaryMessage(t *testing.T) {
	type args struct {
		s *ParameterSummary
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "added",
			args: args{s: &ParameterSummary{
				Name:        "max",
				In:          "query",
				Description: "query",
				Deprecated:  false,
				Breaking:    false,
				Action:      "added",
			}},
			want: "Added: `max` in `query`\n> query\n\n",
		},
		{
			name: "modified",
			args: args{s: &ParameterSummary{
				Name:        "id",
				In:          "path",
				Description: "catalogue id",
				Deprecated:  false,
				Breaking:    false,
				Action:      "modified",
			}},
			want: "Modified: `id` in `path`\n> catalogue id\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := markdownSummaryMessageBuilder{}
			if got := m.BuildParameterSummaryMessage(tt.args.s); got != tt.want {
				t.Errorf("BuildParameterSummaryMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markdownSummaryMessageBuilder_BuildParametersSummaryMessage(t *testing.T) {
	type args struct {
		s *ParametersSummary
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{s: &ParametersSummary{
				Breaking: false,
				Details: []*ParameterSummary{
					{Message: "Added: `max` in `query`\n> query\n\n"},
					{Message: "Modified: `id` in `path`\n> catalogue id\n\n"},
				},
			}},
			want: "###### Parameters:\n\nAdded: `max` in `query`\n> query\n\nModified: `id` in `path`\n> catalogue id\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := markdownSummaryMessageBuilder{}
			if got := m.BuildParametersSummaryMessage(tt.args.s); got != tt.want {
				t.Errorf("BuildParametersSummaryMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markdownSummaryMessageBuilder_BuildPropertiesSummaryMessage(t *testing.T) {
	type args struct {
		s           *PropertiesSummary
		indentLevel int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "added - indentLevel 0",
			args: args{
				s: &PropertiesSummary{
					Name:        "tags",
					Type:        "array",
					Description: "",
					Action:      "added",
					Breaking:    false,
				},
				indentLevel: 0,
			},
			want: "* Added property `tags` (array)\n\n",
		},
		{
			name: "added - indentLevel 1",
			args: args{
				s: &PropertiesSummary{
					Name:        "tags",
					Type:        "array",
					Description: "",
					Action:      "added",
					Breaking:    false,
				},
				indentLevel: 1,
			},
			want: "    * Added property `tags` (array)\n\n",
		},
		{
			name: "deleted - indentLevel 0",
			args: args{
				s: &PropertiesSummary{
					Name:        "tag",
					Type:        "array",
					Description: "",
					Action:      "deleted",
					Breaking:    false,
				},
				indentLevel: 0,
			},
			want: "* Deleted property `tag` (array)\n\n",
		},
		{
			name: "deleted - indentLevel 1",
			args: args{
				s: &PropertiesSummary{
					Name:        "tag",
					Type:        "array",
					Description: "",
					Action:      "deleted",
					Breaking:    false,
				},
				indentLevel: 1,
			},
			want: "    * Deleted property `tag` (array)\n\n",
		},
		{
			name: "modified - indentLevel 0",
			args: args{
				s: &PropertiesSummary{
					Name:        "id",
					Type:        "string",
					Description: "",
					Action:      "modified",
					Breaking:    false,
				},
				indentLevel: 0,
			},
			want: "* Modified property `id` (string)\n\n",
		},
		{
			name: "modified - indentLevel 1",
			args: args{
				s: &PropertiesSummary{
					Name:        "id",
					Type:        "string",
					Description: "",
					Action:      "modified",
					Breaking:    false,
				},
				indentLevel: 1,
			},
			want: "    * Modified property `id` (string)\n\n",
		},
		{
			name: "unsupported action",
			args: args{
				s: &PropertiesSummary{
					Name:        "id",
					Type:        "string",
					Description: "",
					Action:      "unsupported-action",
					Breaking:    false,
				},
				indentLevel: 1,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := markdownSummaryMessageBuilder{}
			if got := m.BuildPropertiesSummaryMessage(tt.args.s, tt.args.indentLevel); got != tt.want {
				t.Errorf("BuildPropertiesSummaryMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markdownSummaryMessageBuilder_BuildRequestBodySummaryDetailMessage(t *testing.T) {
	type args struct {
		d           *RequestBodySummaryDetail
		contentType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "added",
			args: args{
				d: &RequestBodySummaryDetail{
					Breaking:   false,
					Action:     "added",
					Name:       "application/json;charset=UTF-8",
					Properties: nil,
				},
				contentType: "application/json;charset=UTF-8",
			},
			want: "Added content type: `application/json;charset=UTF-8`\n\n",
		},
		{
			name: "deleted",
			args: args{
				d: &RequestBodySummaryDetail{
					Breaking:   false,
					Action:     "deleted",
					Name:       "application/json;charset=UTF-8",
					Properties: nil,
				},
				contentType: "application/json;charset=UTF-8",
			},
			want: "Deleted content type: `application/json;charset=UTF-8`\n\n",
		},
		{
			name: "modified",
			args: args{
				d: &RequestBodySummaryDetail{
					Breaking: false,
					Action:   "modified",
					Name:     "application/json;charset=UTF-8",
					Properties: []*PropertiesSummary{
						{
							Group: "items",
							Type:  "array",
							Nested: []*PropertiesSummary{
								{
									Message: "",
								},
							},
						},
					},
				},
				contentType: "application/json;charset=UTF-8",
			},
			want: "Modified content type: `application/json;charset=UTF-8`\n\n* Modified items (array):\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := markdownSummaryMessageBuilder{}
			if got := m.BuildRequestBodySummaryDetailMessage(tt.args.d, tt.args.contentType); got != tt.want {
				t.Errorf("BuildRequestBodySummaryDetailMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markdownSummaryMessageBuilder_BuildRequestBodySummaryMessage(t *testing.T) {
	type args struct {
		s *RequestBodySummary
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{s: &RequestBodySummary{
				Breaking:    false,
				Message:     "",
				Description: "",
				Details: []*RequestBodySummaryDetail{
					{
						Name:    "",
						Message: "test1",
					},
					{
						Name:    "",
						Message: "test2",
					},
				},
			}},
			want: "###### Request:\n\ntest1test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := markdownSummaryMessageBuilder{}
			if got := m.BuildRequestBodySummaryMessage(tt.args.s); got != tt.want {
				t.Errorf("BuildRequestBodySummaryMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markdownSummaryMessageBuilder_BuildResponseSummaryDetailMessage(t *testing.T) {
	type args struct {
		s     *ResponseSummaryDetail
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "added",
			args: args{
				s: &ResponseSummaryDetail{
					Name:        "Date",
					Description: "",
					Action:      "added",
					Breaking:    false,
				},
				key:   "header",
				value: "Date",
			},
			want: "\n* Added header: `Date`\n",
		},
		{
			name: "deleted",
			args: args{
				s: &ResponseSummaryDetail{
					Name:        "Date",
					Description: "",
					Action:      "deleted",
					Breaking:    false,
				},
				key:   "header",
				value: "Date",
			},
			want: "\n* Deleted header: `Date`\n",
		},
		{
			name: "modified",
			args: args{
				s: &ResponseSummaryDetail{
					Name:        "test",
					Description: "",
					Action:      "modified",
					Breaking:    false,
					Properties: []*PropertiesSummary{
						{
							Group: "items",
							Type:  "array",
							Nested: []*PropertiesSummary{
								{
									Message: "",
								},
							},
						},
					},
				},
				key:   "header",
				value: "test",
			},
			want: "\n* Modified header: `test`\n    * Modified items (array):\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := markdownSummaryMessageBuilder{}
			if got := m.BuildResponseSummaryDetailMessage(tt.args.s, tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("BuildResponseSummaryDetailMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markdownSummaryMessageBuilder_BuildResponsesSummaryDetailMessage(t *testing.T) {
	type args struct {
		d          *ResponsesSummaryDetail
		statusCode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "added",
			args: args{
				d: &ResponsesSummaryDetail{
					Name:        "401",
					Description: "Unauthorized",
					Action:      "added",
					Breaking:    false,
				},
				statusCode: "401",
			},
			want: "Added response: **401 Unauthorized**\n> Unauthorized\n",
		},
		{
			name: "modified",
			args: args{
				d: &ResponsesSummaryDetail{
					Name:        "200",
					Description: "OK",
					Action:      "modified",
					Breaking:    true,
					Details: []*ResponseSummaryDetail{
						{
							Name:        "If-None-Match",
							Description: "",
							Action:      "added",
							Breaking:    false,
							Message:     "\n* Added header: `If-None-Match`\n",
						},
						{
							Name:        "If-Match",
							Description: "",
							Action:      "added",
							Breaking:    false,
							Message:     "\n* Added header: `If-Match`\n",
						},
						{
							Name:        "TRACKINGID",
							Description: "",
							Action:      "added",
							Breaking:    false,
							Message:     "\n* Added header: `TRACKINGID`\n",
						},
						{
							Name:        "ETag",
							Description: "",
							Action:      "added",
							Breaking:    false,
							Message:     "\n* Added header: `ETag`\n",
						},
						{
							Name:        "Date",
							Description: "",
							Action:      "added",
							Breaking:    false,
							Message:     "\n* Added header: `Date`\n",
						},
						{
							Name:        "Link",
							Description: "",
							Action:      "added",
							Breaking:    false,
							Message:     "\n* Added header: `Link`\n",
						},
					},
				},
				statusCode: "200",
			},
			want: "Modified response: **200 OK**\n> OK\n\n* Added header: `If-None-Match`\n\n* Added header: `If-Match`\n\n* Added header: `TRACKINGID`\n\n* Added header: `ETag`\n\n* Added header: `Date`\n\n* Added header: `Link`\n",
		},
		{
			name: "deleted",
			args: args{
				d: &ResponsesSummaryDetail{
					Name:        "401",
					Description: "Unauthorized",
					Action:      "deleted",
					Breaking:    false,
				},
				statusCode: "401",
			},
			want: "Deleted response: **401 Unauthorized**\n> Unauthorized\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := markdownSummaryMessageBuilder{}
			if got := m.BuildResponsesSummaryDetailMessage(tt.args.d, tt.args.statusCode); got != tt.want {
				t.Errorf("BuildResponsesSummaryDetailMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markdownSummaryMessageBuilder_BuildResponsesSummaryMessage(t *testing.T) {
	type args struct {
		s *ResponsesSummary
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{s: &ResponsesSummary{
				Breaking: false,
				Message:  "",
				Details: []*ResponsesSummaryDetail{
					{
						Name:        "401",
						Description: "Unauthorized",
						Action:      "added",
						Breaking:    false,
						Message:     "Added response: **401 Unauthorized**\n> Unauthorized\n",
					},
					{
						Name:        "403",
						Description: "Forbidden",
						Action:      "added",
						Breaking:    false,
						Message:     "Added response: **403 Forbidden**\n> Forbidden\n",
					},
				},
			}},
			want: "###### Response:\n\nAdded response: **401 Unauthorized**\n> Unauthorized\nAdded response: **403 Forbidden**\n> Forbidden\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := markdownSummaryMessageBuilder{}
			if got := m.BuildResponsesSummaryMessage(tt.args.s); got != tt.want {
				t.Errorf("BuildResponsesSummaryMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markdownSummaryMessageBuilder_BuildResultSummaryMessage(t *testing.T) {
	type args struct {
		result *JSONResult
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{result: &JSONResult{
				Added: []*EndpointSummary{
					{
						Method:      "GET",
						Path:        "/tag",
						Description: "test",
					},
				},
				Deleted: []*EndpointSummary{
					{
						Method:      "GET",
						Path:        "/tags",
						Description: "test",
					},
				},
				Deprecated: []*EndpointSummary{
					{
						Method:      "GET",
						Path:        "/catalogue/size",
						Description: "test",
					},
				},
				Modified: []*ModifiedSummary{
					{
						Method:      "GET",
						Path:        "/catalogue",
						Description: "test",
						Summary:     "summary",
						Message:     "message",
					},
				},
				Breaking: false,
				Message:  "",
			},
			},
			want: "#### What's New\n\n##### `GET` /tag\n\n> test\n\n" +
				"#### What's Deleted\n\n##### `GET` /tags\n\n> test\n\n" +
				"#### What's Deprecated\n\n##### `GET` /catalogue/size\n\n> test\n\n" +
				"#### What's Modified\n\n##### `GET` /catalogue\n\n> summary\n\nmessage",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := markdownSummaryMessageBuilder{}
			if got := m.BuildResultSummaryMessage(tt.args.result); got != tt.want {
				t.Errorf("BuildResultSummaryMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markdownSummaryMessageBuilder_BuildSecuritySummaryDetailMessage(t *testing.T) {
	type args struct {
		d *SecuritySummaryDetail
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{d: &SecuritySummaryDetail{
				Breaking: false,
				Name:     "BearerAuth",
				Action:   "added",
				Message:  "Added authentication: `BearerAuth`\n",
			}},
			want: "Added authentication: `BearerAuth`\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := markdownSummaryMessageBuilder{}
			if got := m.BuildSecuritySummaryDetailMessage(tt.args.d); got != tt.want {
				t.Errorf("BuildSecuritySummaryDetailMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markdownSummaryMessageBuilder_BuildSecuritySummaryMessage(t *testing.T) {
	type args struct {
		s *SecuritySummary
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{s: &SecuritySummary{
				Details: []*SecuritySummaryDetail{
					{
						Name:     "BearerAuth",
						Action:   "added",
						Breaking: false,
						Message:  "Added authentication: `BearerAuth`\n",
					},
				},
			}},
			want: "###### Security:\n\nAdded authentication: `BearerAuth`\n",
		},
		{
			name: "empty",
			args: args{s: &SecuritySummary{
				Details: []*SecuritySummaryDetail{},
			}},
			want: "###### Security:\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := markdownSummaryMessageBuilder{}
			if got := m.BuildSecuritySummaryMessage(tt.args.s); got != tt.want {
				t.Errorf("BuildSecuritySummaryMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
