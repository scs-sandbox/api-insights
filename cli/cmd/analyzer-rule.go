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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cisco-developer/api-insights/cli/pkg/model"
	"github.com/cisco-developer/api-insights/cli/pkg/utils"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(analyzerRuleCmd())
}

func analyzerRuleCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "analyzer-rule",
		Short: "Manage analyzer rules",
	}

	cmd.PersistentFlags().StringVarP(&analyzer, flagAnalyzer, "a", "", "API spec analyzer")

	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		logDebugln("failed to bind flags", err.Error())
	}

	// subcommands
	cmd.AddCommand(analyzerRuleListCmd())
	cmd.AddCommand(analyzerRuleGetCmd())
	cmd.AddCommand(analyzerRuleImportCmd())

	return cmd
}

func analyzerRuleListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List analyzer rules",
		Example: `  # List analyzer rules
  api-insights-cli analyzer-rule list --analyzer completeness
  api-insights-cli analyzer-rule ls --analyzer completeness`,
		Run: func(cmd *cobra.Command, args []string) {
			analyzerID := viper.GetString(flagAnalyzer)
			if analyzerID == "" {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("analyzer id is required, for example: api-insights-cli analyzer-rule list --analyzer completeness"))
			}
			rules, err := apiInsightsClient.ListAnalyzerRules(cmd.Context(), analyzerID, map[string]string{})
			if err != nil {
				utils.ExitWithCode(utils.ExitError, err)
			}

			rules.Print(os.Stdout)
		},
	}

	return cmd
}

func analyzerRuleGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get ID",
		Short: "Get analyzer rule by id or name_id",
		Example: `  # Get analyzer rule by name_id
  api-insights-cli analyzer-rule get oas3-jwt-format --analyzer completeness

  # Get analyzer rule by id
  api-insights-cli analyzer-rule get a80f4e25-bbed-11ec-bcc5-8e8722760159 --analyzer completeness`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("id is required, for example: api-insights-cli analyzer-rule get oas3-jwt-format --analyzer completeness"))
			}

			analyzerID := viper.GetString(flagAnalyzer)
			if analyzerID == "" {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("analyzer id is required, for example: api-insights-cli analyzer-rule get oas3-jwt-format --analyzer completeness"))
			}

			id := args[0]
			rule, err := apiInsightsClient.GetAnalyzerRule(cmd.Context(), analyzerID, id)
			if err != nil {
				utils.ExitWithCode(utils.ExitError, err)
			}

			fmt.Println(utils.Pretty(rule))
		},
	}

	return cmd
}

func analyzerRuleImportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import LOCAL_RULES",
		Short: "Import analyzer rules from a local rules file",
		Example: `  # Import analyzer rules from a local file
  api-insights-cli analyzer-rule import rules-guidelines.json --analyzer completeness`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("API spec file is required, for example: api-insights-cli analyzer-rule import rules-guidelines.json --analyzer completeness"))
			}

			analyzerID := viper.GetString(flagAnalyzer)
			if analyzerID == "" {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("analyzer id is required, for example: api-insights-cli analyzer-rule import rules-guidelines.json --analyzer completeness"))
			}

			filename := args[0]
			logDebugf("loading local rules: %s\n", filename)

			d, err := os.ReadFile(filename)
			if err != nil {
				utils.ExitWithCode(utils.ExitError, fmt.Errorf("failed to load rules: %s", err.Error()))
			}
			logDebugf("loaded local rules: %s\n", filename)

			var rules []*model.AnalyzerRule
			err = json.Unmarshal(d, &rules)
			if err != nil {
				utils.ExitWithCode(utils.ExitError, fmt.Errorf("failed to parse rules: %s", err.Error()))
			}

			err = apiInsightsClient.ImportAnalyzerRules(cmd.Context(), analyzerID, rules)
			if err != nil {
				utils.ExitWithCode(utils.ExitError, err)
			}

			logDebugln("imported")
		},
	}

	return cmd
}
