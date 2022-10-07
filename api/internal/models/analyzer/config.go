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

package analyzer

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer/rule"
	"github.com/cisco-developer/api-insights/api/pkg/utils"
	"text/template"
)

const (
	ConfigScoreConfig           = "score_config"
	ConfigServiceNameID         = "service_name_id"
	ConfigServiceNameIDTemplate = "service_name_id_template"
)

type Config map[string]interface{}

func (c Config) UnmarshalInto(v interface{}) error {
	return utils.UnmarshalMapInto(c, v)
}

func (c Config) GetScoreConfig() *ScoreConfig {
	var scoreCfg *ScoreConfig
	scoreCfgRaw, ok := c[ConfigScoreConfig]
	if ok {
		scoreCfgRawBytes, err := json.Marshal(scoreCfgRaw)
		if err == nil {
			_ = json.Unmarshal(scoreCfgRawBytes, &scoreCfg)
		}
	}
	return scoreCfg
}

func (c Config) ServiceNameID() string {
	var s string
	v, ok := c[ConfigServiceNameID]
	if ok {
		s, _ = v.(string)
	}
	return s
}

func (c Config) ServiceNameIDFromTemplate(serviceNameID string) string {
	var s string
	tmpl := c.ServiceNameIDTemplate()
	if tmpl != "" {
		templ := template.Must(template.New("").Parse(tmpl))
		var buf bytes.Buffer
		_ = templ.Execute(&buf, map[string]interface{}{
			"nameID": serviceNameID,
		})
		return buf.String()
	}
	return s
}

func (c Config) ServiceNameIDTemplate() string {
	var s string
	v, ok := c[ConfigServiceNameIDTemplate]
	if ok {
		s, _ = v.(string)
	}
	return s
}

// Scan implements sql.Scanner interface.
// See https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type.
func (c *Config) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	return json.Unmarshal(bytes, &c)
}

func ListToMap(list []*Analyzer) map[SpecAnalyzer]*Analyzer {
	m := map[SpecAnalyzer]*Analyzer{}
	for _, a := range list {
		m[SpecAnalyzer(a.NameID)] = a
	}
	return m
}

// Value implements driver.Valuer interface.
// See https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type.
func (c Config) Value() (driver.Value, error) { return json.Marshal(c) }

type ScoreConfig struct {
	AnalyzerWeight  *float32                  `json:"analyzer_weight"`
	SeverityWeights map[rule.SeverityName]int `json:"severity_weights"`
}

func NewScoreConfig(setDefaults bool) *ScoreConfig {
	cfg := &ScoreConfig{}
	if setDefaults {
		cfg.AnalyzerWeight = utils.Float32Ptr(0)
	}
	return cfg
}

type AnalyzersScoreConfigs map[SpecAnalyzer]*ScoreConfig

func NewAnalyzersScoreConfigsFrom(analyzers map[SpecAnalyzer]*Analyzer) (AnalyzersScoreConfigs, error) {
	var (
		analyzerNames []SpecAnalyzer
		scoreCfgs     = make(map[SpecAnalyzer]*ScoreConfig, len(analyzers))
	)
	for analyzerName, a := range analyzers {
		analyzerNames = append(analyzerNames, analyzerName)
		if scoreCfg := a.Config.GetScoreConfig(); scoreCfg != nil {
			scoreCfgs[analyzerName] = scoreCfg
		} else {
			scoreCfgs[analyzerName] = NewScoreConfig(true)
		}
	}

	// TODO Validate that the sum of the AnalyzerWeights across all scoreCfgs <= 100.

	analyzersWithoutScoreCfgs := len(analyzerNames) - len(scoreCfgs)

	var defaultAnalyzerWeight float32
	if analyzersWithoutScoreCfgs != 0 {
		defaultAnalyzerWeight = float32(100.0/analyzersWithoutScoreCfgs) / 100.0
	}
	defaultSeverityWeights := rule.DefaultSeverityWeights()

	cfg := AnalyzersScoreConfigs{}

	for _, analyzerName := range analyzerNames {
		analyzerWeight := defaultAnalyzerWeight
		severityWeights := defaultSeverityWeights
		if analyzerScoreCfg, ok := scoreCfgs[analyzerName]; ok {
			if analyzerScoreCfg.AnalyzerWeight != nil {
				analyzerWeight = *analyzerScoreCfg.AnalyzerWeight
			}
			if len(analyzerScoreCfg.SeverityWeights) > 0 {
				severityWeights = analyzerScoreCfg.SeverityWeights
			}
		}
		cfg[analyzerName] = &ScoreConfig{
			AnalyzerWeight:  &analyzerWeight,
			SeverityWeights: severityWeights,
		}
	}

	return cfg, nil
}
