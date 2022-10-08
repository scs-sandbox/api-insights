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

package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/cisco-developer/api-insights/cli/pkg/model"
	"github.com/cisco-developer/api-insights/cli/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	output = model.DiffOutputText

	failOnIncompatible bool
	latest             bool
	specState          string
)

const (
	flagOutput             = "output"
	flagLatest             = "latest"
	flagSpecState          = "state"
	flagFailOnIncompatible = "fail-on-incompatible"
)

func init() {
	rootCmd.AddCommand(diffCmd())
}

func diffCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "diff LOCAL_SPEC",
		Short: "Diff local and remote API specs",
		Example: `  # Diff local test data carts spec and latest remote spec for service carts
  api-insights-cli diff testdata/carts.json -s carts --latest

  # Diff local test data carts spec and specific remote spec and fail if API changes broke backward compatibility
  api-insights-cli diff testdata/carts.json -s carts --latest --fail-on-incompatible

  # Diff local test data carts spec and specific remote spec for service carts
  api-insights-cli diff testdata/carts.json -s carts --version 0.0.1 --state Release
  api-insights-cli diff testdata/carts.json -s carts --version 0.0.1 --revision 1

  # Diff specs with specific output format, defaults to text
  api-insights-cli diff testdata/carts.json -s carts --latest -o text
  api-insights-cli diff testdata/carts.json -s carts --latest -o json
  api-insights-cli diff testdata/carts.json -s carts --latest -o markdown
  api-insights-cli diff testdata/carts.json -s carts --latest -o html`,
		Run:  diffSpecs,
		Args: cobra.MinimumNArgs(1),
	}

	cmd.Flags().StringVarP(&output, flagOutput, "o", model.DiffOutputText, "diff output, one of text (default), json, markdown, html")
	cmd.Flags().StringVarP(&service, flagService, "s", "", "service id or nameId for API spec")
	cmd.Flags().StringVarP(&specVersion, flagVersion, "", "", "API spec version")
	cmd.Flags().StringVarP(&specRevision, flagRevision, "", "", "API spec revision")
	cmd.Flags().StringVarP(&specState, flagSpecState, "", "", "API spec state")
	cmd.Flags().BoolVarP(&latest, flagLatest, "l", false, "use latest remote spec")
	cmd.Flags().BoolVarP(&failOnIncompatible, flagFailOnIncompatible, "", false, "fail only if API changes broke backward compatibility")
	cmd.MarkFlagRequired(flagService)
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		fmt.Println("failed to bind flags", err.Error())
	}

	return cmd
}

func diffSpecs(cmd *cobra.Command, args []string) {
	logDebugln("started")
	if len(args) < 1 {
		utils.ExitWithCode(utils.ExitInvalidInput, errors.New("API spec file is required, for example: api-insights-cli diff api.yaml -s <service_id or service_nameId>"))
	}

	serviceID := viper.GetString(flagService)
	filename := args[0]
	logDebugf("loading local spec for service %s: %s\n", serviceID, filename)

	localSpec, err := os.ReadFile(filename)
	if err != nil {
		utils.ExitWithCode(utils.ExitError, fmt.Errorf("failed to load spec: %s", err.Error()))
	}
	logDebugf("loaded local spec for service %s: %s\n", serviceID, filename)

	var remoteSpec *model.Spec
	logDebugf("loading remote spec for service %s\n", serviceID)
	if latest {
		remoteSpec, err = apiInsightsClient.GetLatestSpec(cmd.Context(), serviceID)
	} else {
		// load remote spec by version / revision or by tag
		queries := map[string]string{
			flagVersion:   viper.GetString(flagVersion),
			flagRevision:  viper.GetString(flagRevision),
			flagSpecState: viper.GetString(flagSpecState),
		}
		remoteSpec, err = apiInsightsClient.GetServiceSpec(cmd.Context(), serviceID, queries)
	}
	if err != nil {
		utils.ExitWithCode(utils.ExitError, fmt.Errorf("failed to load remote spec for service %s: %s", serviceID, err.Error()))
	}

	logDebugf("loaded remote spec for service %s: version(%s), revision(%s), docType(%s), score(%d), state(%s)\n", serviceID, remoteSpec.Version, remoteSpec.Revision, remoteSpec.DocType, remoteSpec.Score, remoteSpec.State)

	outputFormat := viper.GetString(flagOutput)
	req := &model.SpecDiffRequest{
		OldSpecDoc: remoteSpec.Doc,
		NewSpecDoc: model.NewSpecDoc(localSpec),
		Config:     &model.SpecDiffConfig{OutputFormat: outputFormat},
	}

	logDebugf("comparing specs for service %s\n", serviceID)
	res, err := apiInsightsClient.Diff(cmd.Context(), req)
	if err != nil {
		utils.ExitWithCode(utils.ExitError, fmt.Errorf("failed to diff spec: %s", err.Error()))
	}
	logDebugf("compared specs for service %s\n", serviceID)

	res.Result.Print(os.Stdout, outputFormat)

	if failOnIncompatible && hasBreakingChanges(cmd.Context(), req, res) {
		utils.ExitWithCode(utils.ExitIncompatibleAPISpec)
	}
}

func hasBreakingChanges(ctx context.Context, req *model.SpecDiffRequest, res *model.SpecDiff) bool {
	switch req.Config.OutputFormat {
	case model.DiffOutputJSON:
		return res.HasBreakingChangesInJSON()

	//	TODO: check if we can consolidate diff api by adding diff summary.
	//	for now trigger additional json-formatted diff and check if breaking
	case model.DiffOutputText, model.DiffOutputMarkdown, model.DiffOutputHTML:
		req.Config = &model.SpecDiffConfig{OutputFormat: model.DiffOutputJSON}
		if r, err := apiInsightsClient.Diff(ctx, req); err == nil {
			return r.HasBreakingChangesInJSON()
		}

	}

	return false
}
