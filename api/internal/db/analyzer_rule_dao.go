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
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"gorm.io/gorm/clause"
	"log"
	"os"
	"strings"
)

const (
	analyzerRulesBatchSize = 100
)

// AnalyzerRuleDAO is the interface to access database
type AnalyzerRuleDAO interface {
	List(ctx context.Context, filter *ListFilter) ([]*analyzer.Rule, error)
	Save(ctx context.Context, ar *analyzer.Rule) error
	Get(ctx context.Context, id string) (*analyzer.Rule, error)
	Delete(ctx context.Context, id string) error
	Import(ctx context.Context, ars []*analyzer.Rule) error
}

// NewAnalyzerRuleDAO create AnalyzerRuleDAO
var NewAnalyzerRuleDAO = func(config *shared.AppConfig) (AnalyzerRuleDAO, error) {
	shared.LogInfof("init NewAnalyzerRuleDAO")
	client, err := NewDBClient(config)
	if err != nil {
		return nil, err
	}
	err = client.AutoMigrate(analyzer.Rule{})
	if err != nil {
		return nil, err
	}

	dao := &blobAnalyzerRuleDAO{client: client, config: config}

	err = dao.preloadDefaultAnalyzerRulesSilently(context.Background(), "internal/data", "rules")
	shared.LogInfof("preloading rules")
	if err != nil {
		shared.LogWarnf("failed to preload rules: %s", err.Error())
	} else {
		shared.LogInfof("preloaded rules")
	}

	return dao, nil
}

type blobAnalyzerRuleDAO struct {
	client *Client
	config *shared.AppConfig
}

// Save object to database
func (dao *blobAnalyzerRuleDAO) Save(ctx context.Context, ar *analyzer.Rule) error {
	span, ctx := shared.StartSpan(ctx, "analyzerRule.id", ar.GetID())
	defer span.Finish()

	err := dao.client.WithContext(ctx).Save(ar).Error
	if err != nil {
		shared.LogErrorf("failed to save analyzer rule %s: %s", ar.GetID(), err.Error())
		return err
	}

	return nil
}

// Get an object with specified id from database
func (dao *blobAnalyzerRuleDAO) Get(ctx context.Context, id string) (*analyzer.Rule, error) {
	span, ctx := shared.StartSpan(ctx, "analyzerRule.id", id)
	defer span.Finish()

	ar := &analyzer.Rule{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).Or("name_id = ?", id).First(ar).Error
	if err != nil {
		shared.LogErrorf("failed to get analyzer rule %s: %s", id, err.Error())
		return nil, err
	}

	return ar, nil
}

// Delete an object with specified id from database
func (dao *blobAnalyzerRuleDAO) Delete(ctx context.Context, id string) error {
	span, ctx := shared.StartSpan(ctx, "analyzerRule.id", id)
	defer span.Finish()

	ar := analyzer.Rule{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).Or("name_id = ?", id).Delete(ar).Error
	if err != nil {
		shared.LogErrorf("could not find object with %s: %s", id, err.Error())
		return ErrNotFound
	}

	return nil
}

// List all objects in database with specified filter
func (dao *blobAnalyzerRuleDAO) List(ctx context.Context, filter *ListFilter) ([]*analyzer.Rule, error) {
	span, ctx := shared.StartSpan(ctx)
	defer span.Finish()

	shared.LogDebugf("fetching analyzer rules: %#v ...", filter)

	var ars []*analyzer.Rule
	db := dao.client.WithContext(ctx).Table(analyzer.AnalyzerRuleTableName)
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
	err := db.Find(&ars).Error
	if err != nil {
		return nil, err
	}

	return ars, nil
}

func (dao *blobAnalyzerRuleDAO) Import(ctx context.Context, ars []*analyzer.Rule) error {
	span, ctx := shared.StartSpan(ctx)
	defer span.Finish()

	err := dao.client.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name_id"}},
			UpdateAll: true,
		}).
		CreateInBatches(ars, analyzerRulesBatchSize).Error
	if err != nil {
		shared.LogErrorf("failed to import %d analyzer rules: %s", len(ars), err.Error())
		return err
	}

	return nil
}

func (dao *blobAnalyzerRuleDAO) preloadDefaultAnalyzerRulesSilently(ctx context.Context, dir, keyword string) error {
	var paths []string

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() || !strings.Contains(f.Name(), keyword) {
			continue
		}

		paths = append(paths, fmt.Sprintf("%s/%s", strings.TrimSuffix(dir, "/"), f.Name()))
	}

	var allRules []*analyzer.Rule
	for _, path := range paths {
		var rules []*analyzer.Rule
		data, err := os.ReadFile(path)
		if err != nil {
			shared.LogWarnf("failed to read file(%s), err: %s", path, err.Error())
			continue
		}

		err = json.Unmarshal(data, &rules)
		if err != nil {
			shared.LogWarnf("failed to unmarshal data, err: %s", err.Error())
			continue
		}

		allRules = append(allRules, rules...)
	}

	return dao.Import(ctx, allRules)
}
