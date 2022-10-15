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

package shared

import (
	"encoding/json"
	"fmt"
	restful "github.com/emicklei/go-restful/v3"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"os"
	"path"
	"strings"
)

type AppConfig struct {
	// AppVersion is the version of the microservice.
	AppVersion string `json:"appVersion"`

	// AppName is the name of the microservice.
	AppName string `json:"appName"`

	// AppPort the port for the App microservice web server to listen on.
	AppPort int `json:"appPort"`

	// AppHost is the host name for the App microservice.  This is used
	// for the OAuth callback url.
	AppHost string `json:"appHost"`

	// AppEnvironment signals the environment this config is intended
	// to be. Possible values include 'production', 'development', etc.
	AppEnvironment Environment `json:"appEnvironment"`

	// ConfigDir is the directory to search for configuration JSON files. Defaults to working directory
	ConfigDir string `json:"configDir"`

	// HTTPHandler is the func that registers the microservice and
	// returns the restful.Container.
	HTTPHandler func(*AppConfig) (*restful.Container, error) `json:"-"`
}

func (c *AppConfig) Env(suffix string) string {
	envStr := fmt.Sprintf("%s_%s", strings.ToUpper(c.AppName), strings.ToUpper(suffix))
	return strings.Replace(envStr, "-", "_", -1)
}

func (c *AppConfig) String() string {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (c *AppConfig) LoadConfigFile(relPath string) ([]byte, error) {
	configPath := path.Join(c.ConfigDir, relPath)
	return os.ReadFile(configPath)
}

func (c *AppConfig) LoadConfigJSON(relPath string, dest interface{}) error {
	raw, err := c.LoadConfigFile(relPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, dest)
}

func wrapCommand(config *AppConfig, cmd *cli.Command) {
	action := cmd.Action
	if action != nil {
		cmd.Action = func(ctx *cli.Context) error {
			return action(ctx)
		}
		cmd.Before = flagJSON(cmd.Flags)
	}
	cmd.Flags = append(serveFlags(config), cmd.Flags...)

	for _, subcmd := range cmd.Subcommands {
		wrapCommand(config, subcmd)
	}
}

func ServeCommand(config *AppConfig, customFlags []cli.Flag, commands ...*cli.Command) *cli.Command {
	for _, cmd := range commands {
		wrapCommand(config, cmd)
	}

	return &cli.Command{
		Name:        config.AppName,
		Description: fmt.Sprintf("Run the %s microservice tool", config.AppName),
		Subcommands: serveCommands(config, customFlags, commands...),
	}
}

func loadConfigJSON(relPath string) (map[string]interface{}, error) {
	raw, err := os.ReadFile(relPath)
	if err != nil {
		return nil, err
	}
	var dest map[string]interface{}
	err = json.Unmarshal(raw, &dest)
	return dest, err
}

func flagJSON(flags []cli.Flag) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		configFile := path.Join(ctx.String("config-dir"), "config.json")
		userConfigFile := path.Join(ctx.String("config-dir"), ".config.json")
		LogDebugf("CONFIG: %s\n", configFile)

		data, err := loadConfigJSON(configFile)
		if err != nil {
			return err
		}

		userData, _ := loadConfigJSON(userConfigFile)
		for k, v := range userData {
			data[k] = v
		}

		b, err := json.Marshal(data)
		if err != nil {
			return err
		}
		inputSourceContext, err := altsrc.NewJSONSource(b)
		if err != nil {
			return err
		}
		return altsrc.ApplyInputSourceValues(ctx, inputSourceContext, flags)
	}
}

func serveFlags(config *AppConfig) []cli.Flag {
	return []cli.Flag{
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "port",
			Usage:       "the http listener port for the service",
			Value:       config.AppPort,
			Destination: &config.AppPort,
			EnvVars:     []string{config.Env("PORT")},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "host",
			Usage:       "the ip address to bind the http listener to for the service",
			Value:       config.AppHost,
			Destination: &config.AppHost,
			EnvVars:     []string{config.Env("HOST")},
		}),
		&cli.StringFlag{
			Name:        "config-dir",
			Usage:       "directory to search for configuration JSON files. Defaults to working directory",
			Value:       "",
			Destination: &config.ConfigDir,
			EnvVars:     []string{config.Env("CONFIG_DIR")},
		},
	}
}
