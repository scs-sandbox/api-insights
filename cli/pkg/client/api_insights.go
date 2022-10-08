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

package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/cisco-developer/api-insights/cli/pkg/config"
	"github.com/cisco-developer/api-insights/cli/pkg/model"
	"github.com/go-resty/resty/v2"
)

type APIInsightsClient interface {
	AnalyzeAPISpec(ctx context.Context, req *model.SpecAnalysisRequest) (*model.SpecAnalysisResponse, error)

	ListServices(ctx context.Context) (model.ServiceList, error)
	GetService(ctx context.Context, id string) (*model.Service, error)
	CreateService(ctx context.Context, s *model.Service) (*model.Service, error)
	DeleteService(ctx context.Context, id string) error
	GetLatestSpec(ctx context.Context, serviceID string) (*model.Spec, error)
	GetServiceSpec(ctx context.Context, serviceID string, queries map[string]string) (*model.Spec, error)
	UploadSpec(ctx context.Context, serviceID string, spec *model.Spec) (*model.Spec, error)

	ListSpecs(ctx context.Context, serviceID string) (model.SpecList, error)
	GetSpec(ctx context.Context, serviceID, id string) (*model.Spec, error)

	ListSpecAnalyses(ctx context.Context, serviceID, specID string) (model.SpecAnalysisList, error)

	Diff(ctx context.Context, req *model.SpecDiffRequest) (*model.SpecDiff, error)

	ListAnalyzers(ctx context.Context, queries map[string]string) (model.AnalyzerList, error)
	GetAnalyzer(ctx context.Context, id string) (*model.Analyzer, error)

	ListAnalyzerRules(ctx context.Context, analyzerID string, queries map[string]string) (model.AnalyzerRuleList, error)
	GetAnalyzerRule(ctx context.Context, analyzerID string, id string) (*model.AnalyzerRule, error)
	ImportAnalyzerRules(ctx context.Context, analyzerID string, ars []*model.AnalyzerRule) error
}

type apiInsightsClient struct {
	baseURL    string
	basePath   string
	headers    map[string]string
	authSetter func(ctx context.Context, rc *resty.Client) error
}

// NewAPIInsightsClient creates a new client for API Insights service.
func NewAPIInsightsClient(conf *config.Config) APIInsightsClient {
	return &apiInsightsClient{
		baseURL:    conf.APIInsightsHost,
		basePath:   conf.APIInsightsBasePath,
		headers:    conf.Headers,
		authSetter: conf.AuthConfig.AuthSetter,
	}
}

func (c *apiInsightsClient) newRestyClient(ctx context.Context) (*resty.Client, error) {
	client := resty.New()
	client.SetBaseURL(c.baseURL)
	err := c.authSetter(ctx, client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *apiInsightsClient) AnalyzeAPISpec(ctx context.Context, req *model.SpecAnalysisRequest) (*model.SpecAnalysisResponse, error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	var result *model.SpecAnalysisResponse
	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s/specs/analyses/analyze", c.basePath))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return result, nil
}

func (c *apiInsightsClient) ListServices(ctx context.Context) (services model.ServiceList, err error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetResult(&services).
		Get(fmt.Sprintf("%s/services", c.basePath))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return
}

func (c *apiInsightsClient) GetService(ctx context.Context, id string) (s *model.Service, err error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetResult(&s).
		Get(fmt.Sprintf("%s/services/%s", c.basePath, id))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return
}

func (c *apiInsightsClient) CreateService(ctx context.Context, s *model.Service) (sOut *model.Service, err error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetBody(s).
		SetResult(&sOut).
		Post(fmt.Sprintf("%s/services", c.basePath))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return sOut, nil
}

func (c *apiInsightsClient) DeleteService(ctx context.Context, id string) error {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return err
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		Delete(fmt.Sprintf("%s/services/%s", c.basePath, id))
	if err != nil {
		return err
	}
	if !res.IsSuccess() {
		return errors.New(res.Status())
	}

	return nil
}

func (c *apiInsightsClient) GetLatestSpec(ctx context.Context, serviceID string) (*model.Spec, error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	var specs model.SpecList
	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetResult(&specs).
		Get(fmt.Sprintf("%s/services/%s/specs?sort=created_at&order=desc&limit=1&withDoc=true", c.basePath, serviceID))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}
	if len(specs) == 0 {
		return nil, errors.New("not found")
	}

	return specs[0], nil
}

func (c *apiInsightsClient) GetServiceSpec(ctx context.Context, serviceID string, queries map[string]string) (*model.Spec, error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	var specs model.SpecList
	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetQueryParams(queries).
		SetResult(&specs).
		Get(fmt.Sprintf("%s/services/%s/specs?sort=created_at&order=desc&limit=1", c.basePath, serviceID))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}
	if len(specs) == 0 {
		return nil, errors.New("not found")
	}

	return specs[0], nil
}

func (c *apiInsightsClient) UploadSpec(ctx context.Context, serviceID string, spec *model.Spec) (*model.Spec, error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	var s *model.Spec
	url := fmt.Sprintf("%s/specs", c.basePath)
	if len(serviceID) > 0 {
		url = fmt.Sprintf("%s/services/%s/specs", c.basePath, serviceID)
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetBody(spec).
		SetResult(&s).
		Post(url)
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return s, nil
}

func (c *apiInsightsClient) ListSpecs(ctx context.Context, serviceID string) (specs model.SpecList, err error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetResult(&specs).
		Get(fmt.Sprintf("%s/services/%s/specs?withDoc=true", c.basePath, serviceID))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return
}

func (c *apiInsightsClient) GetSpec(ctx context.Context, serviceID, id string) (s *model.Spec, err error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetResult(&s).
		Get(fmt.Sprintf("%s/services/%s/specs/%s", c.basePath, serviceID, id))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return
}

func (c *apiInsightsClient) ListSpecAnalyses(ctx context.Context, serviceID, specID string) (analyses model.SpecAnalysisList, err error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetResult(&analyses).
		Get(fmt.Sprintf("%s/services/%s/specs/%s/analyses", c.basePath, serviceID, specID))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return
}

func (c *apiInsightsClient) Diff(ctx context.Context, req *model.SpecDiffRequest) (*model.SpecDiff, error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	var result *model.SpecDiff
	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetBody(req).
		SetResult(&result).
		Post(fmt.Sprintf("%s/specs/diffs/diff", c.basePath))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return result, nil
}

func (c *apiInsightsClient) ListAnalyzers(ctx context.Context, queries map[string]string) (analyzers model.AnalyzerList, err error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetQueryParams(queries).
		SetResult(&analyzers).
		Get(fmt.Sprintf("%s/analyzers", c.basePath))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return
}

func (c *apiInsightsClient) GetAnalyzer(ctx context.Context, id string) (a *model.Analyzer, err error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetResult(&a).
		Get(fmt.Sprintf("%s/analyzers/%s", c.basePath, id))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return
}

func (c *apiInsightsClient) ListAnalyzerRules(ctx context.Context, analyzerID string, queries map[string]string) (rules model.AnalyzerRuleList, err error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetQueryParams(queries).
		SetResult(&rules).
		Get(fmt.Sprintf("%s/analyzers/%s/rules", c.basePath, analyzerID))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return
}

func (c *apiInsightsClient) GetAnalyzerRule(ctx context.Context, analyzerID string, id string) (ar *model.AnalyzerRule, err error) {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return nil, err
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetResult(&ar).
		Get(fmt.Sprintf("%s/analyzers/%s/rules/%s", c.basePath, analyzerID, id))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, errors.New(res.Status())
	}

	return
}

func (c *apiInsightsClient) ImportAnalyzerRules(ctx context.Context, analyzerID string, ars []*model.AnalyzerRule) error {
	client, err := c.newRestyClient(ctx)
	if err != nil {
		return err
	}

	res, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeaders(c.headers).
		SetBody(ars).
		Post(fmt.Sprintf("%s/analyzers/%s/rules/import", c.basePath, analyzerID))
	if err != nil {
		return err
	}
	if !res.IsSuccess() {
		return errors.New(res.Status())
	}

	return nil
}
