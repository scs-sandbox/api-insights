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

package main

import (
	"github.com/cisco-developer/api-insights/api/internal/db"
	"github.com/cisco-developer/api-insights/api/internal/endpoints"
	"github.com/cisco-developer/api-insights/api/internal/info"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer/completeness"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer/security"
	"github.com/cisco-developer/api-insights/api/pkg/apiclarity"
	openapidiff "github.com/cisco-developer/api-insights/api/pkg/differ/openapi-diff"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"github.com/emicklei/go-restful/v3"
	"github.com/urfave/cli/v2"
	"os"
)

func init() {
	restful.SetLogger(shared.GlobalLogger())
}

var (
	appVersion = "0.0.1"
	config     = &shared.AppConfig{
		AppName:     "api-insights",
		AppPort:     8081,
		HTTPHandler: endpoints.API,
	}
)

func main() {
	app := App(appVersion)
	if err := app.Run(os.Args); err != nil {
		shared.LogErrorf("start application api-insights error: %s", err.Error())
		os.Exit(-1)
	}
}

// App returns a top level HTTPApp for apiregistry service
func App(version string) *cli.App {
	config.AppVersion = version
	additionalFlags := make([]cli.Flag, 0)

	additionalFlags = shared.MergeFlags(additionalFlags, openapidiff.Flags())
	additionalFlags = shared.MergeFlags(additionalFlags, apiclarity.Flags())
	additionalFlags = shared.MergeFlags(additionalFlags, db.ClientFlags())
	additionalFlags = shared.MergeFlags(additionalFlags, completeness.Flags())
	additionalFlags = shared.MergeFlags(additionalFlags, security.Flags())
	additionalFlags = shared.MergeFlags(additionalFlags, info.Flags())
	additionalFlags = shared.MergeFlags(additionalFlags, models.Flags())

	return shared.HTTPApp(config, additionalFlags)
}
