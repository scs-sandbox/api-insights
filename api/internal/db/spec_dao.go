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
	"context"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
)

// SpecDAO is the interface to access database
type SpecDAO interface {
	List(context context.Context, filter *ListFilter, withDoc bool) ([]*models.Spec, error)
	Save(context context.Context, spec *models.Spec) error
	Get(context context.Context, id string, withDoc bool) (*models.Spec, error)
	Delete(context context.Context, id string) error
}

// NewSpecDAO create SpecDAO
var NewSpecDAO = func(config *shared.AppConfig) (SpecDAO, error) {
	client, err := NewDBClient(config)
	if err != nil {
		return nil, err
	}
	err = client.AutoMigrate(models.Spec{})
	if err != nil {
		return nil, err
	}

	dao := &blobSpecDAO{client: client, config: config}
	return dao, nil
}

type blobSpecDAO struct {
	client *Client
	config *shared.AppConfig
}

// Save object to database
func (dao *blobSpecDAO) Save(ctx context.Context, spec *models.Spec) error {
	span, ctx := shared.StartSpan(ctx, "spec.id", spec.GetID())
	defer span.Finish()

	err := dao.client.WithContext(ctx).Save(spec).Error
	if err != nil {
		shared.LogErrorf("failed to save spec %s: %s", spec.GetID(), err.Error())
		return err
	}

	return nil
}

// Get an object with specified id from database
func (dao *blobSpecDAO) Get(ctx context.Context, id string, withDoc bool) (*models.Spec, error) {
	span, ctx := shared.StartSpan(ctx, "spec.id", id)
	defer span.Finish()

	spec := &models.Spec{}
	db := dao.client.WithContext(ctx).Table(models.SpecTableName)

	if !withDoc {
		db.Omit("doc")
	}

	err := db.Where("id = ?", id).First(spec).Error
	if err != nil {
		shared.LogErrorf("failed to get spec %s: %s", id, err.Error())
		return nil, err
	}

	return spec, nil
}

// Delete an object with specified id from database
func (dao *blobSpecDAO) Delete(ctx context.Context, id string) error {
	span, ctx := shared.StartSpan(ctx, "spec.id", id)
	defer span.Finish()

	spec := models.Spec{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).Delete(spec).Error
	if err != nil {
		shared.LogErrorf("could not find object with %s: %s", id, err.Error())
		return ErrNotFound
	}

	return nil
}

// List all objects in database with specified filter
func (dao *blobSpecDAO) List(ctx context.Context, filter *ListFilter, withDoc bool) ([]*models.Spec, error) {
	span, ctx := shared.StartSpan(ctx)
	defer span.Finish()

	shared.LogDebugf("fetching specs: %#v ...", filter)

	var specs []*models.Spec
	db := dao.client.WithContext(ctx).Table(models.SpecTableName)

	if !withDoc {
		db.Omit("doc", "doc_compressed")
	}

	query := map[string]interface{}{}
	for k, v := range filter.Indexes {
		query[k] = v
	}
	if len(query) != 0 {
		db = db.Where(query)
	}

	for _, sorter := range filter.Sorters {
		db = db.Order(sorter.OrderBy())
	}

	if filter.Limit > 0 {
		db = db.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		db = db.Offset(filter.Offset)
	}
	err := db.Find(&specs).Error
	if err != nil {
		return nil, err
	}

	return specs, nil
}
