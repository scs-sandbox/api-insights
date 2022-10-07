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

package apiclarity

import (
	"context"
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/apiclarity"
	apiclarityclient "github.com/cisco-developer/api-insights/api/pkg/apiclarity/client"
	operations2 "github.com/cisco-developer/api-insights/api/pkg/apiclarity/client/operations"
	models2 "github.com/cisco-developer/api-insights/api/pkg/apiclarity/models"
	"github.com/cisco-developer/api-insights/api/pkg/utils"
	"github.com/go-openapi/strfmt"
	"time"
)

func NewClient() (models.SpecDocAnalyzer, error) {
	c, err := apiclarity.New(nil)
	if err != nil {
		return nil, err
	}
	s := &client{
		client: c,
	}
	return s, nil
}

type client struct {
	client *apiclarityclient.APIClarityAPIs
}

func (c client) Analyze(doc models.SpecDoc, cfgMap analyzer.Config, serviceNameID *string) (*analyzer.Result, error) {

	if doc == nil || *doc == "" {
		return nil, fmt.Errorf("analyzer.apiclarity: doc is nil or empty")
	}
	if serviceNameID == nil || *serviceNameID == "" {
		return nil, fmt.Errorf("analyzer.apiclarity: serviceNameID is nil or empty")
	}

	cfg := &analyzer.APIClarityConfig{}
	if cfgMap != nil {
		if err := cfgMap.UnmarshalInto(cfg); err != nil {
			return nil, fmt.Errorf("analyzer.apiclarity: invalid config: %v", err)
		}
	}

	apiEvents, err := c.GetAPIEventsByAPIName(context.Background(), *serviceNameID)
	if err != nil {
		return nil, err
	}
	rawResult := &analyzer.APIClarityDriftResult{
		Events:                 apiEvents,
		EventProvidedSpecDiffs: make([]*models2.APIEventSpecDiff, len(apiEvents)),
	}

	// TODO Revisit.
	for i, apiEvent := range apiEvents {
		if apiEvent.HasProvidedSpecDiff != nil && *apiEvent.HasProvidedSpecDiff {
			reconstructedSpecDiff, err := c.GetAPIEventProvidedSpecDiff(context.Background(), apiEvent.ID)
			if err != nil {
				return nil, fmt.Errorf("analyzer.apiclarity: %v", err)
			}
			rawResult.EventProvidedSpecDiffs[i] = reconstructedSpecDiff
		}
	}

	return rawResult.Result()
}

func (c client) GetAPIEventProvidedSpecDiff(ctx context.Context, apiEventID uint32) (*models2.APIEventSpecDiff, error) {
	res, err := c.client.Operations.GetAPIEventsEventIDProvidedSpecDiff(&operations2.GetAPIEventsEventIDProvidedSpecDiffParams{
		EventID: apiEventID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	} else if res == nil || res.Payload == nil {
		return nil, fmt.Errorf("apiclarity.GetAPIEventProvidedSpecDiff(%v): unexpected response (null res/res.Payload)", apiEventID)
	}
	return res.Payload, nil
}

func (c client) GetAPIEventReconstructedSpecDiff(ctx context.Context, apiEventID uint32) (*models2.APIEventSpecDiff, error) {
	res, err := c.client.Operations.GetAPIEventsEventIDReconstructedSpecDiff(&operations2.GetAPIEventsEventIDReconstructedSpecDiffParams{
		EventID: apiEventID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	} else if res == nil || res.Payload == nil {
		return nil, fmt.Errorf("apiclarity.GetAPIEventReconstructedSpecDiff(%v): unexpected response (null res/res.Payload)", apiEventID)
	}
	return res.Payload, nil
}

func (c client) GetAPIEventsByAPIName(ctx context.Context, apiName string) ([]*models2.APIEvent, error) {
	var (
		currentTime = time.Now().UTC()
		endTime     = strfmt.DateTime(currentTime)
		startTime   = strfmt.DateTime(currentTime.AddDate(0, -1, 0))
	)

	res, err := c.client.Operations.GetAPIEvents(&operations2.GetAPIEventsParams{
		EndTime:       endTime,
		HasSpecDiffIs: utils.BoolPtr(true),
		Page:          1,
		PageSize:      50,
		ShowNonAPI:    false,
		SortDir:       utils.StringPtr(apiclarity.SortDirDesc),
		SortKey:       apiclarity.APIEventSortKeyTime,
		SpecIs:        []string{apiName},
		StartTime:     startTime,
		Context:       ctx,
	})
	if err != nil {
		return nil, err
	} else if res == nil || res.Payload == nil {
		return nil, fmt.Errorf("apiclarity.GetAPIEventsByAPIName(%s): unexpected response (null res/res.Payload/res.Payload.Items)", apiName)
	}

	return res.Payload.Items, nil
}
