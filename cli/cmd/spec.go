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

func init() {
	rootCmd.AddCommand(specCmd())
}

func specCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "spec",
		Short: "Manage specs",
	}

	cmd.PersistentFlags().StringVarP(&service, flagService, "s", "", "service id or nameId for api spec")

	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		logDebugln("failed to bind flags", err.Error())
	}

	// subcommands
	cmd.AddCommand(specListCmd())
	cmd.AddCommand(specGetCmd())

	return cmd
}

func specListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List service specs",
		Example: `  # List service specs
  api-insights-cli spec list -s carts
  api-insights-cli spec ls -s carts`,
		Run: func(cmd *cobra.Command, args []string) {
			serviceID := viper.GetString(flagService)
			if serviceID == "" {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("service id is required, for example: api-insights-cli spec get 0ca61204-b97e-11ec-896d-86748ae98d9b -s carts"))
			}
			specs, err := apiInsightsClient.ListSpecs(cmd.Context(), serviceID)
			if err != nil {
				utils.ExitWithCode(utils.ExitError, err)
			}

			specs.Print(os.Stdout)
		},
	}

	return cmd
}

func specGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get ID",
		Short: "Get service spec by id",
		Example: `  # Get service spec by id
  api-insights-cli spec get 0ca61204-b97e-11ec-896d-86748ae98d9b -s carts`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("id is required, for example: api-insights-cli spec get 0ca61204-b97e-11ec-896d-86748ae98d9b"))
			}

			serviceID := viper.GetString(flagService)
			if serviceID == "" {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("service id is required, for example: api-insights-cli spec get 0ca61204-b97e-11ec-896d-86748ae98d9b -s carts"))
			}

			id := args[0]
			spec, err := apiInsightsClient.GetSpec(cmd.Context(), serviceID, id)
			if err != nil {
				utils.ExitWithCode(utils.ExitError, err)
			}

			fmt.Println(utils.Pretty(spec))
		},
	}

	return cmd
}
