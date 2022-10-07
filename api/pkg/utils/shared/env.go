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
	"os"
)

type Environment string

const (
	EnvTest        Environment = "test"
	EnvDevelopment Environment = "development"
	EnvIntegration Environment = "integration"
	EnvStaging     Environment = "staging"
	EnvProduction  Environment = "production"
)

const (
	EnvCorsWhiteListFilePath = "CORS_WHITELIST_FILEPATH"
)

var (
	CorsHandlerEnabled = GetBoolFromEnv("SHARED_CORS_ENABLED", false)
)

func GetBoolFromEnv(key string, defaultValue bool) bool {
	v := os.Getenv(key)
	if v == "true" || v == "1" {
		return true
	}
	if v == "false" || v == "0" {
		return false
	}

	if v == "" {
		return defaultValue
	}
	return !defaultValue
}

func GetStringFromEnv(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}
