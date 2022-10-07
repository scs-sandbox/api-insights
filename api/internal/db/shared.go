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

package db

import (
	"errors"
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"github.com/emicklei/go-restful/v3"
	"strconv"
	"strings"
)

const (
	OrderAsc  = "asc"
	OrderDesc = "desc"
)

// ErrNotFound represents a not found error
var ErrNotFound = errors.New("object not found")

// ListFilter represents request for List operation
type ListFilter struct {
	Model   models.ModelObject
	Indexes map[string]string
	Tags    []string
	Query   string
	Offset  int
	Limit   int

	Sorters []*Sorter
}

// From collects info from http request
func (l *ListFilter) From(req *restful.Request) error {
	// filters
	indexes := map[string]string{}
	for k, v := range req.Request.URL.Query() {
		idx := l.Model.GetIndex(k)
		if idx != "" && v[0] != "" {
			indexes[k] = v[0]
		}
	}
	l.Indexes = indexes

	// sorters
	l.updateSorters(req)

	// pagination
	limitQuery := req.QueryParameter("limit")
	if max := req.QueryParameter("max"); len(max) > 0 {
		limitQuery = max
	}
	limit, err := intParameter(limitQuery, 100, 1000)
	if err != nil {
		return err
	}
	l.Limit = limit

	if i, err := strconv.Atoi(req.QueryParameter("offset")); err == nil {
		l.Offset = i
	}

	return nil
}

func (l *ListFilter) updateSorters(req *restful.Request) {
	sort := req.QueryParameter("sort")
	if len(sort) == 0 {
		return
	}

	order := req.QueryParameter("order")
	if order != OrderDesc {
		order = OrderAsc
	}

	fields := strings.Split(sort, ",")
	for _, field := range fields {
		if !l.Model.Sortable(field) {
			continue
		}
		s := &Sorter{
			Order: order,
			Field: field,
		}

		l.Sorters = append(l.Sorters, s)
	}
}

// AddIndex adds new filter index
func (l *ListFilter) AddIndex(name, value string) *ListFilter {
	idx := l.Model.GetIndex(name)
	if idx == "" {
		idx = name
	}
	l.Indexes[idx] = value
	return l
}

// Pagination calculates pagination by offset, limit and total
func (l *ListFilter) Pagination(total int) models.Pagination {
	return models.Pagination{
		Total:      total,
		PageSize:   l.Limit,
		PageNum:    l.getPageNum(l.Offset, l.Limit),
		TotalPages: l.getTotalPages(total, l.Limit),
	}
}

func (l *ListFilter) getPageNum(offset, limit int) int {
	if limit == 0 {
		return 1
	}

	return 1 + offset/limit
}

func (l *ListFilter) getTotalPages(total, limit int) int {
	if limit == 0 {
		return 1
	}

	return (total + limit - 1) / limit
}

type Sorter struct {
	Order string
	Field string
}

// OrderBy returns MySQL order by value.
func (m *Sorter) OrderBy() string {
	return fmt.Sprintf("%s %s", m.Field, m.Order)
}

// Init initializes service clients
func Init(cfg *shared.AppConfig) {
	if cfg.AppEnvironment == shared.EnvTest {
		return
	}

}

func intParameter(value string, defaultValue int, maxValue int) (int, error) {
	if value != "" {
		ret, err := strconv.Atoi(value)
		if err != nil {
			shared.LogErrorf("invalid value: %v: %v", value, err.Error())
			return defaultValue, err
		}
		if ret > maxValue {
			return maxValue, nil
		}
		return ret, nil
	}
	return defaultValue, nil
}
