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

package access

import (
	"context"
	"github.com/cisco-developer/api-insights/api/internal/models"
)

// Checker serves as an access checker interface.
type Checker interface {
	CheckAccess(ctx context.Context, input *Input) (*Output, error)
}

type Input struct {
	UserRoles models.Roles
	Resources []*models.Resource
}

type Output struct {
	Allow            bool
	DenyError        error
	AllowWithFilters models.AccessDataFilters
}

func (o *Output) Deny(err error) *Output {
	o.Allow = false
	o.DenyError = err
	return o
}

func (o *Output) Denied() bool { return !o.Allow || o.DenyError != nil }
