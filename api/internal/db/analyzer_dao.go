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
	"encoding/json"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"gorm.io/gorm/clause"
	"os"
)

const (
	analyzersBatchSize = 50
	analyzersDataFile  = "internal/data/analyzers.json"
)

// AnalyzerDAO is the interface to access database
type AnalyzerDAO interface {
	List(context context.Context, filter *ListFilter, withRules bool) ([]*analyzer.Analyzer, error)
	Save(context context.Context, analyzer *analyzer.Analyzer) error
	Get(context context.Context, id string) (*analyzer.Analyzer, error)
	Delete(context context.Context, id string) error
}

// NewAnalyzerDAO create AnalyzerDAO
var NewAnalyzerDAO = func(config *shared.AppConfig) (AnalyzerDAO, error) {
	client, err := NewDBClient(config)
	if err != nil {
		return nil, err
	}
	err = client.AutoMigrate(analyzer.Analyzer{})
	if err != nil {
		return nil, err
	}

	dao := &blobAnalyzerDAO{client: client, config: config}

	dao.preloadDefaultAnalyzersSilently(context.Background(), analyzersDataFile)

	return dao, nil
}

type blobAnalyzerDAO struct {
	client *Client
	config *shared.AppConfig
}

// Save object to database
func (dao *blobAnalyzerDAO) Save(ctx context.Context, analyzer *analyzer.Analyzer) error {
	span, ctx := shared.StartSpan(ctx, "analyzer.id", analyzer.GetID())
	defer span.Finish()

	err := dao.client.WithContext(ctx).Save(analyzer).Error
	if err != nil {
		shared.LogErrorf("failed to save analyzer %s: %s", analyzer.GetID(), err.Error())
		return err
	}

	return nil
}

// Get an object with specified id from database
func (dao *blobAnalyzerDAO) Get(ctx context.Context, id string) (*analyzer.Analyzer, error) {
	span, ctx := shared.StartSpan(ctx, "analyzer.id", id)
	defer span.Finish()

	analyzer := &analyzer.Analyzer{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).Or("name_id = ?", id).First(analyzer).Error
	if err != nil {
		shared.LogErrorf("failed to get analyzer %s: %s", id, err.Error())
		return nil, err
	}

	return analyzer, nil
}

// Delete an object with specified id from database
func (dao *blobAnalyzerDAO) Delete(ctx context.Context, id string) error {
	span, ctx := shared.StartSpan(ctx, "analyzer.id", id)
	defer span.Finish()

	analyzer := analyzer.Analyzer{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).Or("name_id = ?", id).Delete(analyzer).Error
	if err != nil {
		shared.LogErrorf("could not find object with %s: %s", id, err.Error())
		return ErrNotFound
	}

	return nil
}

// List all objects in database with specified filter
func (dao *blobAnalyzerDAO) List(ctx context.Context, filter *ListFilter, withRules bool) ([]*analyzer.Analyzer, error) {
	span, ctx := shared.StartSpan(ctx)
	defer span.Finish()

	shared.LogDebugf("fetching analyzers: %#v ...", filter)

	var analyzers []*analyzer.Analyzer
	db := dao.client.WithContext(ctx).Table(analyzer.AnalyzerTableName)
	query := map[string]interface{}{}
	for k, v := range filter.Indexes {
		query[k] = v
	}
	if len(query) != 0 {
		db = db.Where(query)
	}

	if withRules {
		db.Preload("Rules")
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
	err := db.Find(&analyzers).Error
	if err != nil {
		return nil, err
	}

	return analyzers, nil
}

func (dao *blobAnalyzerDAO) Import(ctx context.Context, analyzers []*analyzer.Analyzer) error {
	span, ctx := shared.StartSpan(ctx)
	defer span.Finish()

	err := dao.client.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name_id"}},
			UpdateAll: true,
		}).
		CreateInBatches(analyzers, analyzersBatchSize).Error
	if err != nil {
		shared.LogErrorf("failed to import %d analyzers: %s", len(analyzers), err.Error())
		return err
	}

	return nil
}

func (dao *blobAnalyzerDAO) preloadDefaultAnalyzersSilently(ctx context.Context, filename string) {
	analyzersData, err := os.ReadFile(filename)
	if err != nil {
		shared.LogWarnf("failed to read file(%s), err: %s", filename, err.Error())
		return
	}

	shared.LogInfof("preloading default analyzers: %v ...", string(analyzersData))

	var analyzers []*analyzer.Analyzer
	err = json.Unmarshal(analyzersData, &analyzers)
	if err != nil {
		shared.LogWarnf("failed to unmarshal data, err: %s", err.Error())
		return
	}

	err = dao.Import(ctx, analyzers)
	if err != nil {
		shared.LogWarnf("failed to preload default analyzer(s): %v", err)
		return
	}
}
