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
	"fmt"
	"github.com/cisco-developer/api-insights/cli/pkg/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd())
}

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version information",
		Example: `  # Print the version information
  api-insights-cli version`,
		Run: printVersion,
	}

	return cmd
}

func printVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("Version: %s\n", version.Version)

	if len(version.CommitHash) > 0 {
		fmt.Printf("Commit: %s\n", version.CommitHash)
	}

	if len(version.BuildTimestamp) > 0 {
		fmt.Printf("Build Time: %s\n", version.BuildTimestamp)
	}
}
