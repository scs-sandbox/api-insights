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
	"context"
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/diff"
	"github.com/cisco-developer/api-insights/api/pkg/differ/openapi-diff/result"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"github.com/getkin/kin-openapi/openapi3"
	"os"
	"os/exec"
	"strings"
	"time"
)

func NewClient() (Differ, error) {
	if err := preRunCheck(); err != nil {
		return nil, err
	}
	return &cliClient{}, nil
}

// cliClient implements Differ.
type cliClient struct {
}

func (c *cliClient) DiffDocuments(oldDoc, newDoc models.SpecDoc, cfg *diff.Config, opts *Options) (*diff.Result, error) {

	if cfg == nil {
		cfg = &diff.Config{}
	}
	if cfg.OutputFormat == "" {
		cfg.OutputFormat = "json"
	}

	if opts == nil {
		opts = &DefaultOpts
		opts.Format = &cfg.OutputFormat
	}

	var (
		timestamp           = fmt.Sprintf("%v", time.Now().UnixNano())
		oldFilenamePattern  = "diff-" + timestamp + "-old-*.json"
		newFilenamePattern  = "diff-" + timestamp + "-new-*.json"
		diffFilenamePattern = "diff-" + timestamp + "-diff-*"
	)

	oldFile, err := os.CreateTemp("/tmp", oldFilenamePattern)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = oldFile.Close()
		_ = os.Remove(oldFile.Name())
	}()
	if oldDoc == nil || *oldDoc == "" {
		return nil, fmt.Errorf("oldDoc is nil or empty")
	}
	_, err = oldFile.Write([]byte(*oldDoc))
	if err != nil {
		return nil, err
	}
	newFile, err := os.CreateTemp("/tmp", newFilenamePattern)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = newFile.Close()
		_ = os.Remove(newFile.Name())
	}()
	if newDoc == nil || *newDoc == "" {
		return nil, fmt.Errorf("newDoc is nil or empty")
	}
	_, err = newFile.Write([]byte(*newDoc))
	if err != nil {
		return nil, err
	}

	diffFile, err := os.CreateTemp("/tmp", diffFilenamePattern)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = diffFile.Close()
		_ = os.Remove(diffFile.Name())
	}()

	cmd := c.commandFromOpts(context.TODO(), opts, oldFile.Name(), newFile.Name(), diffFile.Name())

	cmdString := cmd.String()
	shared.LogInfof("Running CMD[%s]...", cmdString)
	if out, err := cmd.CombinedOutput(); err != nil {
		shared.LogErrorf("Unexpected error returned by CMD[%s] (err=%v): %v", cmd.String(), err, string(out))
		return nil, err
	}
	shared.LogInfof("Successfully ran CMD[%s].", cmdString)

	diffFileData, err := os.ReadFile(diffFile.Name())
	if err != nil {
		return nil, err
	}

	res := &diff.Result{}

	switch cfg.OutputFormat {
	case "json":
		changedOpenAPI, err := result.NewChangedOpenAPIFromBytes(diffFileData)
		if err != nil {
			return nil, err
		}
		changedOpenAPI.OldSpecOpenAPI, err = openapi3.NewLoader().LoadFromData([]byte(*oldDoc))
		if err != nil {
			return nil, err
		}
		changedOpenAPI.NewSpecOpenAPI, err = openapi3.NewLoader().LoadFromData([]byte(*newDoc))
		if err != nil {
			return nil, err
		}
		resJSON, err := result.NewResultFrom(changedOpenAPI, diff.NewMarkdownSummaryMessageBuilder())
		if err != nil {
			return nil, err
		}
		res.JSON = resJSON
	case "html":
		res.HTML = string(diffFileData)
	case "markdown":
		res.Markdown = string(diffFileData)
	case "text":
		res.Text = string(diffFileData)
	}

	return res, nil
}

// commandFromOpts builds the `npx openapi-diff ...` command.
func (c *cliClient) commandFromOpts(ctx context.Context, opts *Options, oldFilename, newFilename, diffFilename string) *exec.Cmd {

	if ctx == nil {
		ctx = context.Background()
	}

	if opts == nil {
		opts = &DefaultOpts
	}

	var args []string

	if opts.OpenAPIDiffJavaOpts != nil {
		// If OpenAPIDiffJavaOpts is "", then there will be an extra space in the command, like java  other-option, this cause main class not found error.
		if *opts.OpenAPIDiffJavaOpts != "" {
			javaOpts := strings.Split(*opts.OpenAPIDiffJavaOpts, " ")
			args = append(args, javaOpts...)
		}
	}

	args = append(args,
		"-jar",
		*opts.OpenAPIDiffJarFile,
		oldFilename,
		newFilename,
	)

	switch *opts.Format {
	case "html":
		args = append(args, "--html")
	case "json":
		args = append(args, "--json")
	case "markdown":
		args = append(args, "--markdown")
	case "text":
		args = append(args, "--text")
	}

	args = append(args, diffFilename)

	return exec.CommandContext(ctx, "java", args...)
}

func preRunCheck() error {
	// TODO Check dependencies like java
	return nil
}
