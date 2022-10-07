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

package shared

import (
	"errors"
	"fmt"
	"github.com/gocql/gocql"
	validator "gopkg.in/validator.v2"
	"regexp"
)

// IsUUID is a validator for structs with isUUID validate tag
func isUUIDValidator(v interface{}, param string) error {
	s, ok := v.(string)
	if !ok {
		return errors.New("IsUUID only validates strings")
	}
	if !IsUUID(s) {
		return fmt.Errorf("%v is not a valid UUID", v)
	}
	return nil
}

func IsUUID(s string) bool {
	return UUIDRegex.MatchString(s)
}

func TimeUUID() string {
	return gocql.TimeUUID().String()
}

func init() {
	if err := validator.SetValidationFunc("isUUID", isUUIDValidator); err != nil {
		LogInfof("failed to set the isUUID validator")
	}
}

var (
	// UUIDRegex describes the form of a valid UUID
	UUIDRegex = regexp.MustCompile("(?i)[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}")
)
