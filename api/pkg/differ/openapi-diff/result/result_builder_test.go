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

package result

import (
	"fmt"
	"testing"

	"github.com/cisco-developer/api-insights/api/internal/models/diff"
	"github.com/stretchr/testify/assert"
)

func TestNewResultFrom(t *testing.T) {
	type args struct {
		c                 string
		summaryMsgBuilder diff.SummaryMessageBuilder
	}
	tests := []struct {
		name      string
		args      args
		want      *diff.JSONResult
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "result",
			args: args{
				c:                 "./testdata/openapi-diff-raw-result.json",
				summaryMsgBuilder: diff.NewMarkdownSummaryMessageBuilder(),
			},
			want: &diff.JSONResult{
				Message:  "#### What's New\n\n##### `GET` /test-changed\n\n\n#### What's Deleted\n\n##### `GET` /test\n\n\n",
				Breaking: true,
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c, err := NewChangedOpenAPIFromBytes(loadDiff(tt.args.c))
			assert.NoError(t, err)
			fmt.Printf("%#v\n", c)
			got, err := NewResultFrom(c, tt.args.summaryMsgBuilder)
			tt.assertion(t, err)
			// assert.Equal(t, tt.want, got)
			assert.NotNil(t, got)
			assert.Equal(t, tt.want.Message, got.Message)
			// assert.Equal(t, len(tt.want.JSON.Message), len(got.JSON.Message))
		})
	}
}
