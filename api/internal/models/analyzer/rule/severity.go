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

package rule

type (
	Severity     int
	SeverityName string
)

const (
	SeverityHint Severity = iota + 1
	SeverityInfo
	SeverityWarning
	SeverityError

	SeverityNameHint    SeverityName = "hint"
	SeverityNameInfo    SeverityName = "info"
	SeverityNameWarning SeverityName = "warning"
	SeverityNameError   SeverityName = "error"
)

func (s Severity) Name() SeverityName { return severityNameBySeverity[s] }

func (s Severity) String() string { return string(s.Name()) }

func (s Severity) Weight() int { return int(s) }

func (n SeverityName) String() string { return string(n) }

func (n SeverityName) Severity() Severity {
	switch n {
	case SeverityHint.Name():
		return SeverityHint
	case SeverityInfo.Name():
		return SeverityInfo
	case SeverityWarning.Name():
		return SeverityWarning
	case SeverityError.Name():
		return SeverityError
	}
	return SeverityHint
}

var severityNameBySeverity = map[Severity]SeverityName{
	SeverityHint:    SeverityNameHint,
	SeverityInfo:    SeverityNameInfo,
	SeverityWarning: SeverityNameWarning,
	SeverityError:   SeverityNameError,
}

var defaultSeverityWeights = map[SeverityName]int{
	SeverityNameHint:    SeverityNameHint.Severity().Weight(),
	SeverityNameInfo:    SeverityNameInfo.Severity().Weight(),
	SeverityNameWarning: SeverityNameWarning.Severity().Weight(),
	SeverityNameError:   SeverityNameError.Severity().Weight(),
}

func DefaultSeverityWeights() map[SeverityName]int { return defaultSeverityWeights }
