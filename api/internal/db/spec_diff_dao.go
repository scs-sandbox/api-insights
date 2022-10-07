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

// SpecDiffDAO is the interface to access database
type SpecDiffDAO interface {
	List(context context.Context, filter *ListFilter) ([]*models.SpecDiff, error)
	Save(context context.Context, specDiff *models.SpecDiff) error
	Get(context context.Context, id string) (*models.SpecDiff, error)
	Delete(context context.Context, id string) error
	BatchDeleteBySpecID(context context.Context, specID string) error
}

// NewSpecDiffDAO create SpecDiffDAO
var NewSpecDiffDAO = func(config *shared.AppConfig) (SpecDiffDAO, error) {
	client, err := NewDBClient(config)
	if err != nil {
		return nil, err
	}
	err = client.AutoMigrate(models.SpecDiff{})
	if err != nil {
		return nil, err
	}

	dao := &blobSpecDiffDAO{client: client, config: config}
	return dao, nil
}

type blobSpecDiffDAO struct {
	client *Client
	config *shared.AppConfig
}

// Save object to database
func (dao *blobSpecDiffDAO) Save(ctx context.Context, specDiff *models.SpecDiff) error {
	span, ctx := shared.StartSpan(ctx, "specDiff.id", specDiff.GetID())
	defer span.Finish()

	err := dao.client.WithContext(ctx).Save(specDiff).Error
	if err != nil {
		shared.LogErrorf("failed to save specDiff %s: %s", specDiff.GetID(), err.Error())
		return err
	}

	return nil
}

// Get an object with specified id from database
func (dao *blobSpecDiffDAO) Get(ctx context.Context, id string) (*models.SpecDiff, error) {
	span, ctx := shared.StartSpan(ctx, "specDiff.id", id)
	defer span.Finish()

	specDiff := &models.SpecDiff{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).First(specDiff).Error
	if err != nil {
		shared.LogErrorf("failed to get specDiff %s: %s", id, err.Error())
		return nil, err
	}

	return specDiff, nil
}

// Delete an object with specified id from database
func (dao *blobSpecDiffDAO) Delete(ctx context.Context, id string) error {
	span, ctx := shared.StartSpan(ctx, "specDiff.id", id)
	defer span.Finish()

	specDiff := &models.SpecDiff{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).Delete(specDiff).Error
	if err != nil {
		shared.LogErrorf("could not find object with %s: %s", id, err.Error())
		return ErrNotFound
	}

	return nil
}

// List all objects in database with specified filter
func (dao *blobSpecDiffDAO) List(ctx context.Context, filter *ListFilter) ([]*models.SpecDiff, error) {
	span, ctx := shared.StartSpan(ctx)
	defer span.Finish()

	shared.LogDebugf("fetching specDiffs: %#v ...", filter)

	var specDiffs []*models.SpecDiff
	db := dao.client.WithContext(ctx).Table(models.SpecDiffTableName)
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
	err := db.Find(&specDiffs).Error
	if err != nil {
		return nil, err
	}

	return specDiffs, nil
}

// BatchDeleteBySpecID batch deletes all objects by (models.SpecAnalysis).OldSpecID or (models.SpecAnalysis).NewSpecID
func (dao *blobSpecDiffDAO) BatchDeleteBySpecID(ctx context.Context, specID string) error {
	span, ctx := shared.StartSpan(ctx, "specDiff.specID", specID)
	defer span.Finish()

	specDiff := &models.SpecDiff{}
	err := dao.client.WithContext(ctx).Where("old_spec_id = ?", specID).Or("new_spec_id = ?", specID).Delete(specDiff).Error
	if err != nil {
		shared.LogErrorf("failed to batch delete specDiff by spec_id %s: %s", specID, err.Error())
		return err
	}

	return nil
}
