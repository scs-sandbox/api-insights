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

package differ

import (
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/diff"
	openapidiff "github.com/cisco-developer/api-insights/api/pkg/differ/openapi-diff"
)

type Service interface {
	Diff(req *models.SpecDiffRequest) (*diff.Result, error)
}

func NewService() (Service, error) {
	return &service{}, nil
}

type service struct {
}

func (service) Diff(req *models.SpecDiffRequest) (*diff.Result, error) {
	// Select implementation of differ to run. ATM, only openapidiff differ is supported.

	// openapidiff differ implementation
	differClient, err := openapidiff.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create differ(openapidiff): %v", err)
	}
	diffRes, err := differClient.DiffDocuments(req.OldSpecDoc, req.NewSpecDoc, req.Config, nil)
	if err != nil {
		return nil, err
	}

	return diffRes, nil
}
