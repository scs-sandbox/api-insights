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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const (
	flagSpecID = "spec_id"
)

var (
	specID string
)

func init() {
	rootCmd.AddCommand(specAnalysisCmd())
}

func specAnalysisCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "spec-analysis",
		Short: "Manage spec analyses",
	}

	// subcommands
	cmd.AddCommand(specAnalysisListCmd())

	return cmd
}

func specAnalysisListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List spec analyses available for a given service",
		Example: `  # List service spec analyses
  api-insights-cli spec-analysis list -s carts --spec_id f7fb047c-c219-11ec-b2ff-5a0acd2870c9
  api-insights-cli spec-analysis ls -s carts --spec_id f7fb047c-c219-11ec-b2ff-5a0acd2870c9`,
		Run: func(cmd *cobra.Command, args []string) {
			serviceID := viper.GetString(flagService)
			if serviceID == "" {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("service id is required, for example: api-insights-cli spec-analysis list -s carts --spec_id f7fb047c-c219-11ec-b2ff-5a0acd2870c9"))
			}
			specID := viper.GetString(flagSpecID)
			if specID == "" {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("spec id is required, for example: api-insights-cli spec-analysis list -s carts --spec_id f7fb047c-c219-11ec-b2ff-5a0acd2870c9"))
			}
			specAnalyses, err := apiInsightsClient.ListSpecAnalyses(cmd.Context(), serviceID, specID)
			if err != nil {
				fmt.Println(err.Error())
				utils.ExitWithCode(utils.ExitError, err)
			}

			specAnalyses.Print(os.Stdout)
		},
	}

	cmd.Flags().StringVarP(&service, flagService, "s", "", "service id or nameId for API spec")
	cmd.Flags().StringVarP(&specID, flagSpecID, "", "", "API spec id")
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		fmt.Println("failed to bind flags", err.Error())
	}

	return cmd
}
