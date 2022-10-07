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
	"encoding/json"
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models"
	apiclarityclient "github.com/cisco-developer/api-insights/api/pkg/apiclarity/client"
	operations2 "github.com/cisco-developer/api-insights/api/pkg/apiclarity/client/operations"
	apiclaritymodels "github.com/cisco-developer/api-insights/api/pkg/apiclarity/models"
	"github.com/cisco-developer/api-insights/api/pkg/utils"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"net/url"
)

const (
	SortDirDesc             = "DESC"
	APIEventSortKeyTime     = "time"
	APIInventorySortKeyName = "name"
	APITypeInternal         = "INTERNAL"
)

var ErrNoAPITrafficFound = errors.New("apiclarity: no API traffic found")

type ClientConfig struct {
	apiclarityURL string
}

var (
	defaultClientCfg = &ClientConfig{
		apiclarityURL: "", // populated by Flags
	}
)

func Flags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "apiclarity-url",
			Usage:       "APIClarity URL (e.g. http://localhost:8080)",
			Value:       defaultClientCfg.apiclarityURL,
			Destination: &defaultClientCfg.apiclarityURL,
			EnvVars:     []string{"APICLARITY_URL"},
		}),
	}
}

func New(cfg *ClientConfig) (*apiclarityclient.APIClarityAPIs, error) {
	if cfg == nil {
		cfg = defaultClientCfg
	}
	apiclarityURL, err := url.Parse(cfg.apiclarityURL)
	if err != nil {
		return nil, err
	}

	return apiclarityclient.NewHTTPClientWithConfig(nil, &apiclarityclient.TransportConfig{
		Host:     apiclarityURL.Host,
		BasePath: "/api",
		Schemes:  []string{apiclarityURL.Scheme},
	}), nil
}

func ReconstructSpec(ctx context.Context, client *apiclarityclient.APIClarityAPIs, apiName string) (models.SpecDoc, error) {
	res, err := client.Operations.GetAPIInventory(&operations2.GetAPIInventoryParams{
		NameIs:   []string{apiName},
		Page:     1,
		PageSize: 1,
		SortDir:  utils.StringPtr(SortDirDesc),
		SortKey:  APIInventorySortKeyName,
		Type:     APITypeInternal,
		Context:  ctx,
	})
	if err != nil {
		return nil, err
	} else if res == nil || res.Payload == nil || len(res.Payload.Items) == 0 {
		return nil, fmt.Errorf("apiclarity.ReconstructSpec(%s): unexpected response for GetAPIInventory (null res/res.Payload/res.Payload.Items)", apiName)
	}

	apiID := res.Payload.Items[0].ID

	suggestedResp, err := client.Operations.GetAPIInventoryAPIIDSuggestedReview(&operations2.GetAPIInventoryAPIIDSuggestedReviewParams{
		APIID:   apiID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	} else if suggestedResp == nil || suggestedResp.Payload == nil {
		return nil, fmt.Errorf("apiclarity.ReconstructSpec(%s): unexpected response for GetAPIInventoryAPIIDSuggestedReview (null res/res.Payload/res.Payload.Items)", apiName)
	}

	// suggestedResp.Payload.ReviewPathItems may be empty in case a reconstructed spec already exists & there exists no diff b/w that & the current runtime reconstruction.

	reviewPathItems := suggestedResp.Payload.ReviewPathItems

	if len(reviewPathItems) == 0 {
		return nil, ErrNoAPITrafficFound
	}

	reviewID := suggestedResp.Payload.ID

	approvedResp, err := client.Operations.PostAPIInventoryReviewIDApprovedReview(&operations2.PostAPIInventoryReviewIDApprovedReviewParams{
		Body: &apiclaritymodels.ApprovedReview{
			ReviewPathItems: reviewPathItems,
		},
		ReviewID: reviewID,
		Context:  ctx,
	})
	if err != nil {
		return nil, err
	} else if approvedResp == nil || approvedResp.Payload == nil {
		return nil, fmt.Errorf("apiclarity.ReconstructSpec(%s): unexpected response for PostAPIInventoryReviewIDApprovedReview (null res/res.Payload/res.Payload.Items)", apiName)
	}

	reconstructedResp, err := client.Operations.GetAPIInventoryAPIIDReconstructedSwaggerJSON(&operations2.GetAPIInventoryAPIIDReconstructedSwaggerJSONParams{
		APIID:   apiID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	} else if reconstructedResp == nil || reconstructedResp.Payload == nil {
		return nil, fmt.Errorf("apiclarity.ReconstructSpec(%s): unexpected response for GetAPIInventoryAPIIDReconstructedSwaggerJSON (null res/res.Payload/res.Payload.Items)", apiName)
	}

	data, err := json.MarshalIndent(reconstructedResp.Payload, "", " ")
	if err != nil {
		return nil, fmt.Errorf("apiclarity.ReconstructSpec(%s): invalid response JSON for GetAPIInventoryAPIIDReconstructedSwaggerJSON: %v", apiName, err)
	}

	return models.NewSpecDocFromBytes(data), nil
}
