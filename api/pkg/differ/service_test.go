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

package differ

import (
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/diff"
	openapidiff "github.com/cisco-developer/api-insights/api/pkg/differ/openapi-diff"
	"github.com/cisco-developer/api-insights/api/pkg/utils"
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewService(t *testing.T) {
	tests := []struct {
		name      string
		want      Service
		assertion assert.ErrorAssertionFunc
	}{
		{
			name:      "create service",
			want:      nil,
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewService()
			tt.assertion(t, err)
			// TODO: panic here
			// assert.Implements(t, (Service)(nil), got)
			// assert.Equal(t, tt.want, got)
		})
	}
}

// non-backward compatible API changes
// func Test_service_Diff_non_break(t *testing.T) {
func Test_service_Diff(t *testing.T) {
	jarPath := "/tmp/openapi-diff-cli-2.1.0-beta.3-all.jar"
	if p := os.Getenv("OPENAPI_DIFF_JAR_FILE"); p != "" {
		jarPath = p
	}
	javaOpts := "-Xms512m -Xmx1024m"
	if p := os.Getenv("OPENAPI_DIFF_JAVA_OPTS"); p != "" {
		javaOpts = p
	}
	stubs := gostub.Stub(&openapidiff.DefaultOpts, openapidiff.Options{
		OpenAPIDiffJarFile:  utils.StringPtr(jarPath),
		OpenAPIDiffJavaOpts: utils.StringPtr(javaOpts),
		Format:              utils.StringPtr("json"),
	})
	defer stubs.Reset()

	type args struct {
		req *models.SpecDiffRequest
	}
	s := service{}

	tests := []struct {
		name      string
		s         service
		args      args
		want      *diff.Result
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "Same doc will not break anything",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Breaking: false,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 1. Changes in URL path structure for previously existing resources.",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/1-change-url.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's New\n\n##### `GET` /test-changed\n\n\n#### What's Deleted\n\n##### `GET` /test\n\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 2. Introduction of new required request query parameters.",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/2-new-required-request-query-param.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Parameters:\n\nAdded: `type` in `query`\n> desc\n\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 3. Introduction of new required representation fields. [in-req]",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/3-1-new-required-representation-field-in-req.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `POST` /pets\n\n\n###### Request:\n\nModified content type: `application/json`\n\n* Added property `info` (string)\n\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 3. Introduction of new required representation fields. [in-res]",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/3-2-new-required-representation-field-in-res.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `POST` /pets\n\n\n###### Response:\n\nModified response: **201 Created**\n> Created\n\n* Modified content type: `application/json`\n    * Added property `info` (string)\n\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 4. New resource access restrictions due to authorization policy changes.",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/4-authentication-policy-change.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Security:\n\nAdded authentication: `bearer`\nDeleted authentication: `api_key`\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 5. Removal of previously existing resource operations.",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/5-remove-existing-resource-operation.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Deleted\n\n##### `GET` /test\n\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 6. Removal of fields previously included in response entities (example).",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/6-remove-field-in-response.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Response:\n\nModified response: **200 OK**\n> OK\n\n* Modified content type: `application/json`\n    * Deleted property `tag` (string)\n\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 7. Rejection of previously recognized query parameters and entity fields. [query]",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/7-1-reject-query-param.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Parameters:\n\nDeleted: `limit` in `query`\n> How many items to return at one time (max 100)\n\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 7. Rejection of previously recognized query parameters and entity fields. [req body]",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/7-2-reject-entity-field-in-req.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `POST` /pets\n\n\n###### Request:\n\nModified content type: `application/json`\n\n* Deleted property `tag` (string)\n\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 7. Rejection of previously recognized query parameters and entity fields. [res body]",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/7-3-reject-entity-field-in-res.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Breaking: true,
					Message:  "#### What's Modified\n\n##### `POST` /pets\n\n\n###### Response:\n\nModified response: **201 Created**\n> Created\n\n* Modified content type: `application/json`\n    * Deleted property `tag` (string)\n\n\n",
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 8. Discontinued support of previously recognized media types.",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/8-discontinued-media-type.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Response:\n\nModified response: **200 OK**\n> OK\n\n* Deleted content type: `application/json`\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 9. Redefinition of existing error response keys and codes. [keys]",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/9-1-redefine-error-response-keys.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Response:\n\nModified response: **500 Internal Server Error**\n> unexpected error\n\n* Modified content type: `application/json`\n    * Added property `msg` (string)\n\n    * Deleted property `message` (string)\n\n\n##### `POST` /pets\n\n\n###### Response:\n\nModified response: **default **\n> unexpected error\n\n* Modified content type: `application/json`\n    * Added property `msg` (string)\n\n    * Deleted property `message` (string)\n\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 9. Redefinition of existing error response keys and codes. [codes]",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/9-2-redefine-error-response-code.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Response:\n\nAdded response: **400 Bad Request**\n> unexpected error\nDeleted response: **500 Internal Server Error**\n> unexpected error\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[break] 12. Change in value type or format of previously recognized query parameters or entity data.",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/break/12-change-query-entity-type.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Parameters:\n\nModified: `limit` in `query`\n> How many items to return at one time (max 100)\n\n###### Response:\n\nModified response: **200 OK**\n> OK\n\n* Modified content type: `application/json`\n    * Modified property `id` (integer)\n\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[non-break] 1. Introduction of new URL path structures and operations.",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/non-break/1-new-url-and-operation.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's New\n\n##### `GET` /new-test\n\n\n",
					Breaking: false,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[non-break] 2. Introduction of new optional query parameters, and entity fields.",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/non-break/2-new-optional-query-param-entity-field.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Parameters:\n\nAdded: `type` in `query`\n> type\n\n###### Response:\n\nModified response: **200 OK**\n> OK\n\n* Modified content type: `application/json`\n    * Added property `info` (string)\n\n\n",
					Breaking: false,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[non-break] 3. Disregarding previously recognized query parameters and entity fields (provided the expected behavior remains consistent).",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/non-break/3-disregard-query-param-entity-field.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Parameters:\n\nModified: `min` in `query`\n> min items to return\n\n\n##### `POST` /pets\n\n\n###### Request:\n\nModified content type: `application/json`\n\n\n",
					Breaking: false,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[non-break] 4. Removal of access restrictions on existing resources.",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/non-break/4-remove-access-restriction.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Security:\n\nDeleted authentication: `api_key`\n\n",
					Breaking: false,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[non-break] 5. Introduction of support for new media types.",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/non-break/5-support-new-media-type.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Response:\n\nModified response: **200 OK**\n> OK\n\n* Added content type: `application/xml`\n\n",
					Breaking: false,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[non-break] 6. Introduction of new error response keys and codes (as defined in STATUSCODES).",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/non-break/6-new-error-response-code-keys.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Response:\n\nAdded response: **400 Bad Request**\n> other error\nModified response: **500 Internal Server Error**\n> unexpected error\n\n* Modified content type: `application/json`\n    * Added property `detail` (string)\n\n\n##### `POST` /pets\n\n\n###### Response:\n\nModified response: **default **\n> unexpected error\n\n* Modified content type: `application/json`\n    * Added property `detail` (string)\n\n\n",
					Breaking: false,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[non-break] 7. Deprecation of existing error response keys and codes. [keys]",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/non-break/7-1-deprecate-error-response-keys.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Response:\n\nModified response: **500 Internal Server Error**\n> unexpected error\n\n* Modified content type: `application/json`\n    * Modified property `message` (string)\n\n\n##### `POST` /pets\n\n\n###### Response:\n\nModified response: **default **\n> unexpected error\n\n* Modified content type: `application/json`\n    * Modified property `message` (string)\n\n\n",
					Breaking: false,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "[non-break] 7. Deprecation of existing error response keys and codes. [delete codes]",
			s:    s,
			args: args{
				req: &models.SpecDiffRequest{
					OldSpecDoc: models.SpecDoc(loadSpec("testdata/0-base.yaml")),
					NewSpecDoc: models.SpecDoc(loadSpec("testdata/non-break/7-2-delete-error-response-code.yaml")),
					SpecDiffConfig: models.SpecDiffConfig{
						Config: &diff.Config{
							OutputFormat: "json",
						},
					},
				},
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's Modified\n\n##### `GET` /test\n\n\n###### Response:\n\nDeleted response: **500 Internal Server Error**\n> unexpected error\n\n",
					Breaking: false,
				},
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{}
			got, err := s.Diff(tt.args.req)
			tt.assertion(t, err)
			// assert.Equal(t, tt.want.JSON, got.JSON)
			assert.Equal(t, tt.want.JSON.Message, got.JSON.Message)
			assert.Equal(t, tt.want.JSON.Breaking, got.JSON.Breaking)
		})
	}
}

func loadSpec(file string) *string {
	content, _ := os.ReadFile(file)
	s := string(content)
	return &s
}
