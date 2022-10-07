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

// OrganizationDAO is the interface to access database
type OrganizationDAO interface {
	List(context context.Context, filter *ListFilter, accessFilter *models.OrgServiceAccessDataFilter) ([]*models.Organization, error)
	Save(context context.Context, organization *models.Organization, accessFilter *models.OrgServiceAccessDataFilter) error
	Get(context context.Context, id string) (*models.Organization, error)
	Delete(context context.Context, id string) error
}

// NewOrganizationDAO create OrganizationDAO
var NewOrganizationDAO = func(config *shared.AppConfig) (OrganizationDAO, error) {
	client, err := NewDBClient(config)
	if err != nil {
		return nil, err
	}
	err = client.AutoMigrate(models.Organization{})
	if err != nil {
		return nil, err
	}

	dao := &blobOrganizationDAO{client: client, config: config}
	return dao, nil
}

type blobOrganizationDAO struct {
	client *Client
	config *shared.AppConfig
}

// Save object to database
func (dao *blobOrganizationDAO) Save(ctx context.Context, organization *models.Organization, accessFilter *models.OrgServiceAccessDataFilter) error {
	span, ctx := shared.StartSpan(ctx, "organization.id", organization.GetID())
	defer span.Finish()

	if accessFilter != nil {
		if authorized, err := accessFilter.CanCreateOrg(organization.NameID); !authorized {
			shared.LogErrorf("unauthorized to save organization %s: %s", organization.GetID(), err.Error())
			return err
		}
	}

	err := dao.client.WithContext(ctx).Save(organization).Error
	if err != nil {
		shared.LogErrorf("failed to save organization %s: %s", organization.GetID(), err.Error())
		return err
	}

	return nil
}

// Get an object with specified id from database
func (dao *blobOrganizationDAO) Get(ctx context.Context, id string) (*models.Organization, error) {
	span, ctx := shared.StartSpan(ctx, "organization.id", id)
	defer span.Finish()

	organization := &models.Organization{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).Or("name_id = ?", id).First(organization).Error
	if err != nil {
		shared.LogErrorf("failed to get organization %s: %s", id, err.Error())
		return nil, err
	}

	return organization, nil
}

// Delete an object with specified id from database
func (dao *blobOrganizationDAO) Delete(ctx context.Context, id string) error {
	span, ctx := shared.StartSpan(ctx, "organization.id", id)
	defer span.Finish()

	organization := models.Organization{}
	err := dao.client.WithContext(ctx).Where("id = ?", id).Or("name_id = ?", id).Delete(organization).Error
	if err != nil {
		shared.LogErrorf("could not find object with %s: %s", id, err.Error())
		return ErrNotFound
	}

	return nil
}

// List all objects in database with specified filter
func (dao *blobOrganizationDAO) List(ctx context.Context, filter *ListFilter, accessFilter *models.OrgServiceAccessDataFilter) ([]*models.Organization, error) {
	span, ctx := shared.StartSpan(ctx)
	defer span.Finish()

	shared.LogDebugf("fetching organizations: %#v ...", filter)

	var organizations []*models.Organization
	db := dao.client.WithContext(ctx).Table(models.OrganizationTableName)
	query := map[string]interface{}{}
	for k, v := range filter.Indexes {
		query[k] = v
	}
	if len(query) != 0 {
		db = db.Where(query)
	}

	if accessFilter != nil {
		authorized, byOrganizationIDs, err := accessFilter.CanListOrgs()
		if !authorized && err != nil {
			shared.LogErrorf("unauthorized to list organizations: %s", err.Error())
			return nil, err
		}
		if len(byOrganizationIDs) != 0 {
			db = db.Where("name_id IN ?", byOrganizationIDs)
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
	err := db.Find(&organizations).Error
	if err != nil {
		return nil, err
	}

	return organizations, nil
}
