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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const (
	flagService  = "service"
	flagVersion  = "version"
	flagRevision = "revision"
	flagFile     = "file"
	flagData     = "data"
)

var (
	service      string
	specVersion  string
	specRevision string
	file         string
	data         string
)

func init() {
	rootCmd.AddCommand(serviceCmd())
}

func serviceCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "service",
		Short: "Manage services and related specs",
	}

	// subcommands
	cmd.AddCommand(serviceListCmd())
	cmd.AddCommand(serviceGetCmd())
	cmd.AddCommand(serviceCreateCmd())
	cmd.AddCommand(serviceDeleteCmd())
	cmd.AddCommand(serviceUploadSpecCmd())

	return cmd
}

func serviceListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List services registered for spec analysis",
		Example: `  # List services
  api-insights-cli service list
  api-insights-cli service ls`,
		Run: func(cmd *cobra.Command, args []string) {
			services, err := apiInsightsClient.ListServices(cmd.Context())
			if err != nil {
				utils.ExitWithCode(utils.ExitError, err)
			}

			services.Print(os.Stdout)
		},
	}

	return cmd
}

func serviceGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get ID",
		Short: "Get service by id or name_id",
		Example: `  # Get service by name_id
  api-insights-cli service get carts

  # Get service by id
  api-insights-cli service get 8d7b0619-b97d-11ec-8967-86748ae98d9b`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("id is required, for example: api-insights-cli service get carts"))
			}

			id := args[0]
			service, err := apiInsightsClient.GetService(cmd.Context(), id)
			if err != nil {
				utils.ExitWithCode(utils.ExitError, err)
			}

			fmt.Println(utils.Pretty(service))
		},
	}

	return cmd
}

func serviceCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a service for API spec analysis",
		Example: `  # Create a service from a file
  api-insights-cli service create -f carts-service.json

  # Create a service from raw data
  api-insights-cli service create --data '{
    "organization_id": "Carts",
    "product_tag": "Carts",
    "name_id": "carts",
    "title": "Carts APIs",
    "description": "Carts APIs",
    "contact": {
    "name": "Carts Management Team",
      "url": "https://carts.example.com",
      "email": "carts@example.com"
    }
  }'`,
		Run: func(cmd *cobra.Command, args []string) {
			s := &model.Service{}

			filename := viper.GetString(flagFile)
			data := viper.GetString(flagData)
			if len(filename) > 0 {
				logDebugf("loading file: %s\n", filename)

				b, err := os.ReadFile(filename)
				if err != nil {
					utils.ExitWithCode(utils.ExitError, fmt.Errorf("failed to load file: %s", err.Error()))
				}
				logDebugf("loaded file: %s\n", filename)

				if err = json.Unmarshal(b, &s); err != nil {
					utils.ExitWithCode(utils.ExitError, fmt.Errorf("invalid service object: %s", err.Error()))
				}
			} else if len(data) > 0 {
				if err := json.Unmarshal([]byte(data), &s); err != nil {
					utils.ExitWithCode(utils.ExitError, fmt.Errorf("invalid service object: %s", err.Error()))
				}
			} else {
				utils.ExitWithCode(utils.ExitError, fmt.Errorf("use -f or -d to specify service object"))
			}

			logDebugf("creating service: %s\n", s.NameID)
			res, err := apiInsightsClient.CreateService(cmd.Context(), s)
			if err != nil {
				utils.ExitWithCode(utils.ExitError, err)
			}
			logDebugf("created service: %s - %s\n", res.NameID, res.ID)

			fmt.Println(utils.Pretty(res))
		},
	}

	cmd.Flags().StringVarP(&file, flagFile, "f", "", "file name")
	cmd.Flags().StringVarP(&data, flagData, "d", "", "raw data")
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		fmt.Println("failed to bind flags", err.Error())
	}

	return cmd
}

func serviceDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete ID",
		Short: "Delete a service by id or name_id",
		Example: `  # Delete a service by name_id
  api-insights-cli service delete carts

  # Delete a service by id
  api-insights-cli service delete 8d7b0619-b97d-11ec-8967-86748ae98d9b`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("id is required, for example: api-insights-cli service delete carts"))
			}

			id := args[0]
			err := apiInsightsClient.DeleteService(cmd.Context(), id)
			if err != nil {
				utils.ExitWithCode(utils.ExitError, err)
			}

			fmt.Println(utils.Pretty(service))
		},
	}

	return cmd
}

func serviceUploadSpecCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uploadspec LOCAL_SPEC",
		Short: "Upload local spec for a service under analysis",
		Example: `  # Upload local spec with specific service name_id, spec version, revision
  api-insights-cli service uploadspec testdata/carts.json -s carts -v 1.0.0 -r 1

  # Upload local spec with specific service id, spec version, revision
  api-insights-cli service uploadspec testdata/carts.json -s 1555b762-b9d3-11ec-af7b-a6db741213e2 -v 1.0.0 -r 1

  # Upload local spec with specific service name id and spec revision (spec version will be derived from spec)
  api-insights-cli service uploadspec testdata/carts.json -s carts -r 1

  # Upload local spec with specific service id and spec revision (spec version will be derived from spec)
  api-insights-cli service uploadspec testdata/carts.json -s 1555b762-b9d3-11ec-af7b-a6db741213e2 -r 1`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logDebugln("started")
			if len(args) < 1 {
				utils.ExitWithCode(utils.ExitInvalidInput, errors.New("API spec file is required, for example: api-insights-cli service uploadspec sample-apis.yaml"))
			}

			serviceID := viper.GetString(flagService)
			version := viper.GetString(flagVersion)
			revision := viper.GetString(flagRevision)

			filename := args[0]
			logDebugf("loading spec: %s\n", filename)

			data, err := os.ReadFile(filename)
			if err != nil {
				utils.ExitWithCode(utils.ExitError, fmt.Errorf("failed to load spec: %s", err.Error()))
			}
			logDebugf("loaded spec: %s\n", filename)

			logDebugln("uploading spec")
			spec := &model.Spec{
				Doc:       model.NewSpecDoc(data),
				Revision:  revision,
				ServiceID: serviceID,
				Version:   version,
			}
			res, err := apiInsightsClient.UploadSpec(cmd.Context(), serviceID, spec)
			if err != nil {
				utils.ExitWithCode(utils.ExitError, err)
			}
			logDebugf("uploaded spec %s: %s\n", filename, res.ID)
			logDebugln("completed")
		},
	}

	cmd.Flags().StringVarP(&service, flagService, "s", "", "service id or nameId for API spec")
	cmd.Flags().StringVarP(&specVersion, flagVersion, "v", "", "API spec version, optional if version value is provided in spec")
	cmd.Flags().StringVarP(&specRevision, flagRevision, "r", "", "API spec revision")
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		fmt.Println("failed to bind flags", err.Error())
	}

	return cmd
}
