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

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Different(t *testing.T) {
	assert.Equal(t, []string{"a"}, Different([]string{"a", "b", "c"}, []string{"b", "c"}))
	assert.Equal(t, []string{"a"}, Different([]string{"a", "c"}, []string{"b", "c"}))
	assert.Equal(t, []string{"b"}, Different([]string{"b", "c"}, []string{"a", "c"}))
}

func Test_Intersect(t *testing.T) {
	assert.Equal(t, []string{"b", "c"}, Intersect([]string{"a", "b", "c"}, []string{"b", "c"}))
	assert.Equal(t, []string{"c"}, Intersect([]string{"a", "c"}, []string{"b", "c"}))
}
