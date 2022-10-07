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

// SpecAnalysisDAO is the interface to access database
type SpecAnalysisDAO interface {
	List(context context.Context, filter *ListFilter) ([]*models.SpecAnalysis, error)
	Save(context context.Context, specAnalysis *models.SpecAnalysis) error
	Get(context context.Context, id string) (*models.SpecAnalysis, error)
	Delete(context context.Context, id string) error
	BatchDeleteBySpecID(context context.Context, specID string) error
}

// NewSpecAnalysisDAO create SpecAnalysisDAO
var NewSpecAnalysisDAO = func(config *shared.AppConfig) (SpecAnalysisDAO, error) {
	client, err := NewDBClient(config)
	if err != nil {
		return nil, err
	}
	err = client.AutoMigrate(models.SpecAnalysis{})
	if err != nil {
		return nil, err
	}

	dao := &blobSpecAnalysisDAO{client: client, config: config}
	return dao, nil
}

type blobSpecAnalysisDAO struct {
	client *Client
	config *shared.AppConfig
}

// Save object to database
func (dao *blobSpecAnalysisDAO) Save(ctx context.Context, specAnalysis *models.SpecAnalysis) error {
	span, ctx := shared.StartSpan(ctx, "specAnalysis.id", specAnalysis.GetID())
	defer span.Finish()

	err := dao.client.WithContext(ctx).Save(specAnalysis).Error
	if err != nil {
		shared.LogErrorf("failed to save specAnalysis %s: %s", specAnalysis.GetID(), err.Error())
		return err
	}

	return nil
}

// Get an object with specified id from database
func (dao *blobSpecAnalysisDAO) Get(ctx context.Context, id string) (*models.SpecAnalysis, error) {
	span, ctx := shared.StartSpan(ctx, "specAnalysis.id", id)
	defer span.Finish()

	specAnalysis := &models.SpecAnalysis{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).First(specAnalysis).Error
	if err != nil {
		shared.LogErrorf("failed to get specAnalysis %s: %s", id, err.Error())
		return nil, err
	}

	return specAnalysis, nil
}

// Delete an object with specified id from database
func (dao *blobSpecAnalysisDAO) Delete(ctx context.Context, id string) error {
	span, ctx := shared.StartSpan(ctx, "specAnalysis.id", id)
	defer span.Finish()

	specAnalysis := models.SpecAnalysis{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).Delete(specAnalysis).Error
	if err != nil {
		shared.LogErrorf("could not find object with %s: %s", id, err.Error())
		return ErrNotFound
	}

	return nil
}

// List all objects in database with specified filter
func (dao *blobSpecAnalysisDAO) List(ctx context.Context, filter *ListFilter) ([]*models.SpecAnalysis, error) {
	span, ctx := shared.StartSpan(ctx)
	defer span.Finish()

	shared.LogDebugf("fetching specAnalyses: %#v ...", filter)

	var specAnalyses []*models.SpecAnalysis
	db := dao.client.WithContext(ctx).Table(models.SpecAnalysisTableName)
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
	err := db.Find(&specAnalyses).Error
	if err != nil {
		return nil, err
	}

	return specAnalyses, nil
}

// BatchDeleteBySpecID batch deletes all objects by (models.SpecAnalysis).SpecID
func (dao *blobSpecAnalysisDAO) BatchDeleteBySpecID(ctx context.Context, specID string) error {
	span, ctx := shared.StartSpan(ctx, "specAnalysis.specID", specID)
	defer span.Finish()

	specAnalysis := models.SpecAnalysis{}
	err := dao.client.WithContext(ctx).Where("spec_id = ?", specID).Delete(specAnalysis).Error
	if err != nil {
		shared.LogErrorf("failed to batch delete specAnalysis by spec_id %s: %s", specID, err.Error())
		return err
	}

	return nil
}
