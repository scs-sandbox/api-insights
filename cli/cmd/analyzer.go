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
	"errors"
	"fmt"
	"github.com/cisco-developer/api-insights/cli/pkg/utils"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(analyzerCmd())
}

func analyzerCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "analyzer",
		Short: "Manage analyzers",
	}

	// subcommands
	cmd.AddCommand(analyzerListCmd())
	cmd.AddCommand(analyzerGetCmd())

	return cmd
}

func analyzerListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List analyzers",
		Example: `  # List analyzers
  api-insights-cli analyzer list
  api-insights-cli analyzer ls`,
		Run: func(cmd *cobra.Command, args []string) {
			analyzers, err := apiInsightsClient.ListAnalyzers(cmd.Context(), map[string]string{})
			if err != nil {
				utils.ExitWithCode(utils.ExitError, err)
			}

			analyzers.Print(os.Stdout)
		},
	}

	return cmd
}

func analyzerGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get ID",
		Short: "Get analyzer by id or name_id",
		Example: `  # Get analyzer by name_id
  api-insights-cli analyzer get guidelines

  # Get analyzer by id
  api-insights-cli analyzer get 4c7a4c72-bbed-11ec-bcc4-8e8722760159`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("id is required, for example: api-insights-cli analyzer get guidelines"))
			}

			id := args[0]
			spec, err := apiInsightsClient.GetAnalyzer(cmd.Context(), id)
			if err != nil {
				utils.ExitWithCode(utils.ExitError, err)
			}

			fmt.Println(utils.Pretty(spec))
		},
	}

	return cmd
}
