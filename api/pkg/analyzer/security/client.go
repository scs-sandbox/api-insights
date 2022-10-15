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

package security

import (
	"context"
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/panoptica"
	panopticaclient "github.com/cisco-developer/api-insights/api/pkg/panoptica/client"
	api_security2 "github.com/cisco-developer/api-insights/api/pkg/panoptica/client/api_security"
	models2 "github.com/cisco-developer/api-insights/api/pkg/panoptica/models"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"github.com/go-openapi/strfmt"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"time"
)

const (
	waitMaxRetries = 300
	waitDelay      = 1 * time.Second
)

var (
	panopticaURL       = "appsecurity.cisco.com"
	panopticaAccessKey = ""
	panopticaSecretKey = ""
)

func Flags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "panoptica-url",
			Usage:       "Panoptica URL (e.g. appsecurity.cisco.com",
			Value:       panopticaURL,
			Destination: &panopticaURL,
			EnvVars:     []string{"PANOPTICA_URL"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "panoptica-access-key",
			Usage:       "Panoptica Access Key",
			Value:       panopticaAccessKey,
			Destination: &panopticaAccessKey,
			EnvVars:     []string{"PANOPTICA_ACCESS_KEY"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "panoptica-secret-key",
			Usage:       "Panoptica Secret Key",
			Value:       panopticaSecretKey,
			Destination: &panopticaSecretKey,
			EnvVars:     []string{"PANOPTICA_SECRET_KEY"},
		}),
	}
}

func NewClient() (models.SpecDocAnalyzer, error) {
	c := &client{
		baseURL: panopticaURL,
		client:  panoptica.New(panopticaURL, panopticaAccessKey, panopticaSecretKey),
	}

	return c, nil
}

type client struct {
	baseURL string
	client  *panopticaclient.SecureApplicationAPI
}

func (c *client) Analyze(doc models.SpecDoc, cfgMap analyzer.Config, serviceNameID *string) (*analyzer.Result, error) {
	if doc == nil || *doc == "" {
		return nil, fmt.Errorf("analyzer.security: doc is nil or empty")
	}
	if serviceNameID == nil || *serviceNameID == "" {
		return nil, fmt.Errorf("analyzer.security: serviceNameID is nil or empty")
	}

	cfg := &analyzer.SecurityConfig{}
	if cfgMap != nil {
		if err := cfgMap.UnmarshalInto(cfg); err != nil {
			return nil, fmt.Errorf("analyzer.security: invalid config: %v", err)
		}
	}

	ctx := context.Background()
	apiName := *serviceNameID

	if api, _ := c.getExternalAPIByName(ctx, apiName); api != nil {
		err := c.deleteExternalAPI(ctx, *api.Identifier)
		if err != nil {
			return nil, err
		}
		shared.LogDebugf("analyzer.security: deleted external api %s", apiName)
	}

	err := c.addExternalAPI(ctx, apiName)
	if err != nil {
		return nil, err
	}
	shared.LogDebugf("analyzer.security: added external api %s", apiName)

	api, err := c.getExternalAPIByName(ctx, apiName)
	if err != nil {
		return nil, err
	}
	apiID := *api.Identifier
	shared.LogDebugf("analyzer.security: fetched external api %s: %s", apiName, apiID)

	err = c.uploadExternalAPISpec(ctx, apiID, *doc)
	if err != nil {
		return nil, err
	}
	shared.LogDebugf("analyzer.security: uploaded spec for external api %s: %s", apiName, apiID)

	shared.LogDebugf("analyzer.security: waiting until scored for external api %s: %s", apiName, apiID)
	scored, err := c.waitUntilScored(ctx, apiID)
	if err != nil || !scored {
		return nil, err
	}

	res, err := c.getExternalAPI(ctx, apiID)
	if err != nil {
		return nil, err
	}
	shared.LogDebugf("analyzer.security: fetched score for external api %s: %s", apiName, apiID)

	return analyzer.GetSecurityResult(*doc, res)
}

func (c *client) addExternalAPI(ctx context.Context, name string) error {
	params := api_security2.NewPostAPISecurityAPIParamsWithContext(ctx).WithBody(&models2.APISecurityAPI{Name: &name})
	_, err := c.client.APISecurity.PostAPISecurityAPI(params)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) uploadExternalAPISpec(ctx context.Context, id strfmt.UUID, spec string) error {
	params := api_security2.NewPutAPISecurityOpenAPISpecsCatalogIDParamsWithContext(ctx).WithCatalogID(id).WithBody(spec)
	_, err := c.client.APISecurity.PutAPISecurityOpenAPISpecsCatalogID(params)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) getExternalAPIByName(ctx context.Context, name string) (*models2.APIServiceExternal, error) {
	params := api_security2.NewGetAPISecurityExternalCatalogParamsWithContext(ctx).WithName(&name)
	res, err := c.client.APISecurity.GetAPISecurityExternalCatalog(params)
	if err != nil {
		return nil, err
	}
	if len(res.GetPayload().Items) == 0 {
		return nil, fmt.Errorf("%s not found", name)
	}

	return res.GetPayload().Items[0], nil
}

func (c *client) waitUntilScored(ctx context.Context, id strfmt.UUID) (scored bool, err error) {
	params := api_security2.NewGetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParamsWithContext(ctx).WithCatalogID(id)

	i := 0
	for !scored && i < waitMaxRetries {
		time.Sleep(waitDelay)

		shared.LogDebugf("analyzer.security: checking score status for external api %s: %d/%d", id.String(), i+1, waitMaxRetries)
		res, err := c.client.APISecurity.GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatus(params)
		if err != nil {
			return false, err
		}

		if status := res.GetPayload(); status == models2.OpenAPISpecScoreStatusSCORED {
			shared.LogDebugf("analyzer.security: completed for external api %s", id.String())
			return true, nil
		}

		i++
	}

	return
}

func (c *client) getExternalAPI(ctx context.Context, id strfmt.UUID) (*models2.APIServiceDrillDownExternal, error) {
	params := api_security2.NewGetAPISecurityExternalCatalogCatalogIDParamsWithContext(ctx).WithCatalogID(id)
	res, err := c.client.APISecurity.GetAPISecurityExternalCatalogCatalogID(params)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

func (c *client) deleteExternalAPI(ctx context.Context, id strfmt.UUID) error {
	params := api_security2.NewDeleteAPISecurityAPICatalogIDParamsWithContext(ctx).WithCatalogID(id)
	_, err := c.client.APISecurity.DeleteAPISecurityAPICatalogID(params)
	if err != nil {
		return err
	}

	return nil
}
