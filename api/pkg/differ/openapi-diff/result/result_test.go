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
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewChangedOpenApiFromBytes(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "petstore diff",
			args: args{
				data: loadDiff("./testdata/openapi-diff-raw-result.json"),
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewChangedOpenAPIFromBytes(tt.args.data)
			assert.NotNil(t, got)
			tt.assertion(t, err)
		})
	}
}

func loadDiff(file string) []byte {
	content, _ := os.ReadFile(file)
	return content
}
