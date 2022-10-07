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

// ServiceDAO is the interface to access database
type ServiceDAO interface {
	List(context context.Context, filter *ListFilter, accessFilter *models.OrgServiceAccessDataFilter) ([]*models.Service, error)
	Save(context context.Context, service *models.Service, accessFilter *models.OrgServiceAccessDataFilter) error
	Get(context context.Context, id string) (*models.Service, error)
	Delete(context context.Context, id string) error
}

// NewServiceDAO create ServiceDAO
var NewServiceDAO = func(config *shared.AppConfig) (ServiceDAO, error) {
	client, err := NewDBClient(config)
	if err != nil {
		return nil, err
	}
	err = client.AutoMigrate(models.Service{})
	if err != nil {
		return nil, err
	}

	dao := &blobServiceDAO{client: client, config: config}
	return dao, nil
}

type blobServiceDAO struct {
	client *Client
	config *shared.AppConfig
}

// Save object to database
func (dao *blobServiceDAO) Save(ctx context.Context, service *models.Service, accessFilter *models.OrgServiceAccessDataFilter) error {
	span, ctx := shared.StartSpan(ctx, "service.id", service.GetID())
	defer span.Finish()

	if accessFilter != nil {
		if authorized, err := accessFilter.CanCreateService(service.OrganizationID); !authorized {
			shared.LogErrorf("unauthorized to save service %s: %s", service.GetID(), err.Error())
			return err
		}
	}

	err := dao.client.WithContext(ctx).Save(service).Error
	if err != nil {
		shared.LogErrorf("failed to save service %s: %s", service.GetID(), err.Error())
		return err
	}

	return nil
}

// Get an object with specified id from database
func (dao *blobServiceDAO) Get(ctx context.Context, id string) (*models.Service, error) {
	span, ctx := shared.StartSpan(ctx, "service.id", id)
	defer span.Finish()

	service := &models.Service{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).Or("name_id = ?", id).First(service).Error
	if err != nil {
		shared.LogErrorf("failed to get service %s: %s", id, err.Error())
		return nil, err
	}

	return service, nil
}

// Delete an object with specified id from database
func (dao *blobServiceDAO) Delete(ctx context.Context, id string) error {
	span, ctx := shared.StartSpan(ctx, "service.id", id)
	defer span.Finish()

	service := models.Service{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).Or("name_id = ?", id).Delete(service).Error
	if err != nil {
		shared.LogErrorf("could not find object with %s: %s", id, err.Error())
		return ErrNotFound
	}

	return nil
}

// List all objects in database with specified filter
func (dao *blobServiceDAO) List(ctx context.Context, filter *ListFilter, accessFilter *models.OrgServiceAccessDataFilter) ([]*models.Service, error) {
	span, ctx := shared.StartSpan(ctx)
	defer span.Finish()

	shared.LogDebugf("fetching services: %#v ...", filter)

	var services []*models.Service
	db := dao.client.WithContext(ctx).Table(models.ServiceTableName)
	query := map[string]interface{}{}
	for k, v := range filter.Indexes {
		query[k] = v
	}
	if len(query) != 0 {
		db = db.Where(query)
	}

	if accessFilter != nil {
		authorized, byOrganizationIDs, err := accessFilter.CanListServices()
		if !authorized && err != nil {
			shared.LogErrorf("unauthorized to list services: %s", err.Error())
			return nil, err
		}
		db = db.Where("visibility IN ?", []string{"", models.ServiceVisibilityPublic}).Or("visibility IS NULL")
		if len(byOrganizationIDs) != 0 {
			db = db.Or("organization_id IN ?", byOrganizationIDs)
		}
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
	err := db.Find(&services).Error
	if err != nil {
		return nil, err
	}

	return services, nil
}
