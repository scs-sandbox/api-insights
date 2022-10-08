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

package model

import (
	"fmt"
	"github.com/cisco-developer/api-insights/cli/pkg/utils"
	"io"
	"time"
)

type SpecDoc *string

func NewSpecDoc(b []byte) SpecDoc {
	s := string(b)
	return &s
}

// Spec represents a spec
type Spec struct {
	ID        string    `json:"id,omitempty" validate:"required"`
	Doc       SpecDoc   `json:"doc"`
	DocType   string    `json:"doc_type"`
	Revision  string    `json:"revision"`
	Score     int       `json:"score"`
	ServiceID string    `json:"service_id"`
	State     string    `json:"state"` // Archive, Releases, Development, Latest
	Valid     string    `json:"valid"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SpecList []*Spec

func (m SpecList) Print(w io.Writer) {
	table := utils.NewTable(w)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"#", "ID", "Version", "Revision", "Score", "Valid", "State", "Service ID"})

	for i, spec := range m {
		row := []string{fmt.Sprintf("%d", i+1), spec.ID, spec.Version, spec.Revision, fmt.Sprint(spec.Score), spec.Valid, spec.State, spec.ServiceID}
		table.Append(row)
	}

	table.Render()
}
