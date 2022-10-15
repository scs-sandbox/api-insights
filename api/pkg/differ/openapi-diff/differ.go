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

package openapidiff

import (
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/diff"
	"github.com/cisco-developer/api-insights/api/pkg/utils"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

type Differ interface {
	DiffDocuments(oldDoc, newDoc models.SpecDoc, cfg *diff.Config, opts *Options) (*diff.Result, error)
}

type Options struct {
	OpenAPIDiffJarFile  *string // file path of the OpenAPI JAR file, set by Flags
	OpenAPIDiffJavaOpts *string // Java options for running openapi-diff, e.g. -Xms512m -Xmx1024m, set by Flags
	Format              *string // format to use for outputting results [string] [choices: "html", "json", "markdown", "html", "text"] [default: "json"]
}

var (
	DefaultOpts = Options{
		// OpenAPIDiffJarFile:  utils.StringPtr("/tmp/openapi-diff-cli-2.0.1-all.jar"), // populated by Flags
		OpenAPIDiffJarFile:  utils.StringPtr(""), // populated by Flags
		OpenAPIDiffJavaOpts: utils.StringPtr(""), // populated by Flags
		Format:              utils.StringPtr("json"),
	}
)

func Flags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "openapi-diff-jar-file",
			Usage:       "File path of the openapi-diff JAR file",
			Value:       *DefaultOpts.OpenAPIDiffJarFile,
			Destination: DefaultOpts.OpenAPIDiffJarFile,
			EnvVars:     []string{"OPENAPI_DIFF_JAR_FILE"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "openapi-diff-java-opts",
			Usage:       "Java options for running openapi-diff, e.g. -Xms512m -Xmx1024m",
			Value:       *DefaultOpts.OpenAPIDiffJavaOpts,
			Destination: DefaultOpts.OpenAPIDiffJavaOpts,
			EnvVars:     []string{"OPENAPI_DIFF_JAVA_OPTS"},
		}),
	}
}
