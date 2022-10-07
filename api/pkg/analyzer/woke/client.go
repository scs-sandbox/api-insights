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

package woke

import (
	"encoding/json"
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/get-woke/woke/cmd"
	"github.com/get-woke/woke/pkg/config"
	"github.com/get-woke/woke/pkg/ignore"
	"github.com/get-woke/woke/pkg/parser"
	"github.com/get-woke/woke/pkg/printer"
	"os"
	"time"
)

func NewClient() (models.SpecDocAnalyzer, error) {
	return &client{}, nil
}

type client struct {
}

func (c client) Analyze(doc models.SpecDoc, cfgMap analyzer.Config, serviceNameID *string) (*analyzer.Result, error) {

	if doc == nil || *doc == "" {
		return nil, fmt.Errorf("analyzer.woke: doc is nil or empty")
	}

	cfg := &analyzer.WokeConfig{}
	if cfgMap != nil {
		if err := cfgMap.UnmarshalInto(cfg); err != nil {
			return nil, fmt.Errorf("analyzer.woke: invalid config: %v", err)
		}
	}
	cfg.SetDefaults()

	var (
		timestamp            = fmt.Sprintf("%v", time.Now().UnixNano())
		inputFilenamePattern = "lint-" + timestamp + "-in-*.json"
		outputFilename       = "lint-" + timestamp + "-out.json"
	)

	inputFile, err := os.CreateTemp("/tmp", inputFilenamePattern)
	if err != nil {
		return nil, err
	} else if _, err = inputFile.Write([]byte(*doc)); err != nil {
		return nil, err
	}
	defer func() {
		_ = inputFile.Close()
		_ = os.Remove(inputFile.Name())
	}()

	outputFile, err := os.CreateTemp("/tmp", outputFilename)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = outputFile.Close()
		_ = os.Remove(outputFile.Name())
	}()

	wokeCfg, err := config.NewConfig(cfg.Config, cfg.DisableDefaultRules)
	if err != nil {
		return nil, fmt.Errorf("analyzer.woke: %v", err)
	}

	if len(wokeCfg.Rules) == 0 {
		return nil, fmt.Errorf("analyzer.woke: %v", cmd.ErrNoRulesEnabled)
	}

	var ignorer *ignore.Ignore
	//if !cfg.NoIgnore {
	// TODO
	//}

	wokeParser := parser.NewParser(wokeCfg.Rules, ignorer)

	wokePrinter, err := printer.NewPrinter(cfg.OutputName, outputFile)
	if err != nil {
		return nil, fmt.Errorf("analyzer.woke: %v", err)
	}

	findings := wokeParser.ParsePaths(wokePrinter, inputFile.Name())

	outputData, err := os.ReadFile(outputFile.Name())
	if err != nil {
		return nil, err
	}

	var result *analyzer.WokeResult
	if findings > 0 {
		if err = json.Unmarshal(outputData, &result); err != nil {
			return nil, err
		}
	}

	return result.Result()
}
