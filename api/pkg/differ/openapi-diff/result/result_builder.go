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

package result

import (
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models/diff"
	"github.com/getkin/kin-openapi/openapi3"
)

func NewResultFrom(c *ChangedOpenAPI, summaryMsgBuilder diff.SummaryMessageBuilder) (*diff.JSONResult, error) {
	result := &diff.JSONResult{
		Breaking: c.Incompatible,
	}

	for _, e := range c.NewEndpoints {
		result.Added = append(result.Added, &diff.EndpointSummary{Path: e.PathURL, Method: e.Method, Description: e.Summary, Message: fmt.Sprintf("%s - Added", e.PathURL)})
	}

	for _, e := range c.MissingEndpoints {
		result.Deleted = append(result.Deleted, &diff.EndpointSummary{Path: e.PathURL, Method: e.Method, Description: e.Summary, Message: fmt.Sprintf("%s - Deleted", e.PathURL)})
	}

	for _, e := range c.DeprecatedEndpoints {
		result.Deprecated = append(result.Deprecated, &diff.EndpointSummary{Path: e.PathURL, Method: e.Method, Description: e.Summary, Message: fmt.Sprintf("%s - Deprecated", e.PathURL)})
	}

	for _, o := range c.ChangedOperations {
		changedOperationSummary := &diff.ModifiedSummary{
			Path:        o.PathURL,
			Method:      o.HTTPMethod,
			Summary:     o.NewOperation.Summary,
			Description: o.NewOperation.Description,
			Breaking:    o.Incompatible,
		}

		oldOpSummary, newOpSummary, err := buildOperationSummariesFrom(o.HTTPMethod, o.PathURL, c)
		if err != nil {
			return nil, err
		}

		changedOperationSummary.Old, changedOperationSummary.New = oldOpSummary, newOpSummary

		// Build parameters s report.
		if o.Parameters != nil {
			changedOperationSummary.ParametersSummary = buildParametersSummary(o.Parameters, summaryMsgBuilder)
		}
		// Build request body s report.
		if o.RequestBody != nil {
			changedOperationSummary.RequestBodySummary = buildRequestBodySummary(o.RequestBody, summaryMsgBuilder)
		}
		// Build responses s report.
		if o.APIResponses != nil {
			changedOperationSummary.ResponsesSummary = buildResponsesSummary(o.APIResponses, summaryMsgBuilder)
		}
		// Build security s report.
		if o.SecurityRequirements != nil {
			changedOperationSummary.SecuritySummary = buildSecuritySummary(o.SecurityRequirements, summaryMsgBuilder)
		}

		changedOperationSummary.Message = summaryMsgBuilder.BuildModifiedSummaryMessage(changedOperationSummary)

		result.Modified = append(result.Modified, changedOperationSummary)
	}

	result.Message = summaryMsgBuilder.BuildResultSummaryMessage(result)

	return result, nil
}

func buildOperationSummariesFrom(httpMethod, pathURL string, c *ChangedOpenAPI) (oldOpSummary, newOpSummary *diff.OperationSummary, err error) {
	var (
		oldSpecOpenAPI, newSpecOpenAPI = c.OldSpecOpenAPI, c.NewSpecOpenAPI
	)
	if oldSpecOpenAPI == nil || newSpecOpenAPI == nil {
		return nil, nil, fmt.Errorf("oldSpecOpenAPI or newSpecOpenAPI not found in ChangedOpenAPI")
	}

	pathItem := oldSpecOpenAPI.Paths.Find(pathURL)
	if pathItem == nil {
		return nil, nil, fmt.Errorf("path %s not found in oldSpecOpenAPI", pathURL)
	}

	oldOp := pathItem.GetOperation(httpMethod)
	if oldOp == nil {
		return nil, nil, fmt.Errorf("operation %s %s not found in oldSpecOpenAPI", httpMethod, pathURL)
	}
	oldOpSummary = &diff.OperationSummary{Operation: *oldOp}

	pathItem = newSpecOpenAPI.Paths.Find(pathURL)
	if pathItem == nil {
		return nil, nil, fmt.Errorf("path %s not found in newSpecOpenAPI", pathURL)
	}

	op := pathItem.GetOperation(httpMethod)
	if op == nil {
		return nil, nil, fmt.Errorf("operation %s %s not found in newSpecOpenAPI", httpMethod, pathURL)
	}
	newOpSummary = &diff.OperationSummary{Operation: *op}

	return oldOpSummary, newOpSummary, nil
}

func buildParametersSummary(c *ChangedParameters, msgBuilder diff.SummaryMessageBuilder) *diff.ParametersSummary {
	s := &diff.ParametersSummary{
		Breaking: c.Incompatible,
	}

	for _, p := range c.Increased {
		pSummary := &diff.ParameterSummary{
			Parameter:   p,
			Name:        p.Name,
			In:          p.In,
			Description: p.Description,
			Action:      diff.ActionAdded,
		}
		pSummary.Message = msgBuilder.BuildParameterSummaryMessage(pSummary)

		s.Details = append(s.Details, pSummary)
	}

	for _, p := range c.Missing {
		pSummary := &diff.ParameterSummary{
			Parameter:   p,
			Name:        p.Name,
			In:          p.In,
			Description: p.Description,
			Action:      diff.ActionDeleted,
		}
		pSummary.Message = msgBuilder.BuildParameterSummaryMessage(pSummary)
		s.Details = append(s.Details, pSummary)
	}

	for _, p := range c.Changed {
		pSummary := &diff.ParameterSummary{
			OldParameter: p.OldParameter,
			NewParameter: p.NewParameter,
			Name:         p.Name,
			In:           p.In,
			Description:  p.NewParameter.Description,
			Breaking:     p.Incompatible,
			Action:       diff.ActionModified,
		}
		pSummary.Message = msgBuilder.BuildParameterSummaryMessage(pSummary)
		s.Details = append(s.Details, pSummary)
	}

	s.Message = msgBuilder.BuildParametersSummaryMessage(s)

	return s
}

func buildRequestBodySummary(c *ChangedRequestBody, msgBuilder diff.SummaryMessageBuilder) *diff.RequestBodySummary {
	if c.Content == nil {
		return nil
	}

	s := &diff.RequestBodySummary{
		Breaking: c.Incompatible,
	}

	for mediaTypeKey, mediaType := range c.Content.Increased {
		reqBodySummaryDetail := &diff.RequestBodySummaryDetail{
			ReqBody: &openapi3.RequestBody{
				Content: openapi3.Content{
					mediaTypeKey: mediaType,
				},
			},
			Name:   mediaTypeKey,
			Action: diff.ActionAdded,
		}
		reqBodySummaryDetail.Message = msgBuilder.BuildRequestBodySummaryDetailMessage(reqBodySummaryDetail, mediaTypeKey)
		s.Details = append(s.Details, reqBodySummaryDetail)
	}

	for mediaTypeKey, mediaType := range c.Content.Missing {
		reqBodySummaryDetail := &diff.RequestBodySummaryDetail{
			ReqBody: &openapi3.RequestBody{
				Content: openapi3.Content{
					mediaTypeKey: mediaType,
				},
			},
			Name:   mediaTypeKey,
			Action: diff.ActionDeleted,
		}
		reqBodySummaryDetail.Message = msgBuilder.BuildRequestBodySummaryDetailMessage(reqBodySummaryDetail, mediaTypeKey)
		s.Details = append(s.Details, reqBodySummaryDetail)
	}

	for mediaType, changedMediaType := range c.Content.Changed {
		if changedMediaType.Schema != nil {
			reqBodySummaryDetail := &diff.RequestBodySummaryDetail{
				OldReqBody: c.OldRequestBody,
				NewReqBody: c.NewRequestBody,
				Name:       mediaType,
				Action:     diff.ActionModified,
				Breaking:   changedMediaType.Incompatible,
			}

			for k, p := range changedMediaType.Schema.IncreasedProperties {
				pSummary := &diff.PropertiesSummary{
					Name:        k,
					Type:        p.Type,
					Description: p.Description,
					Action:      diff.ActionAdded,
				}
				pSummary.Message = msgBuilder.BuildPropertiesSummaryMessage(pSummary, 0)
				reqBodySummaryDetail.Properties = append(reqBodySummaryDetail.Properties, pSummary)
			}

			for k, p := range changedMediaType.Schema.MissingProperties {
				pSummary := &diff.PropertiesSummary{
					Name:        k,
					Type:        p.Type,
					Description: p.Description,
					Action:      diff.ActionDeleted,
				}
				pSummary.Message = msgBuilder.BuildPropertiesSummaryMessage(pSummary, 0)
				reqBodySummaryDetail.Properties = append(reqBodySummaryDetail.Properties, pSummary)
			}

			for k, changedProperty := range changedMediaType.Schema.ChangedProperties {
				pSummary := &diff.PropertiesSummary{
					Name:        k,
					Type:        changedProperty.NewSchema.Type,
					Description: changedProperty.NewSchema.Description,
					Action:      diff.ActionModified,
					Breaking:    changedProperty.Incompatible,
				}
				pSummary.Nested = buildNestedPropertiesSummary(changedProperty, msgBuilder, 1)
				pSummary.Message = msgBuilder.BuildPropertiesSummaryMessage(pSummary, 0)
				reqBodySummaryDetail.Properties = append(reqBodySummaryDetail.Properties, pSummary)
			}

			if changedMediaType.Schema.Type == "array" && changedMediaType.Schema.Items != nil {
				pSummary := &diff.PropertiesSummary{
					Action: diff.ActionModified,
					Name:   "items",
					Type:   changedMediaType.Schema.Type,
					Group:  "items",
				}
				pSummary.Nested = buildNestedPropertiesSummary(changedMediaType.Schema.Items, msgBuilder, 1)
				reqBodySummaryDetail.Properties = append(reqBodySummaryDetail.Properties, pSummary)
			}

			reqBodySummaryDetail.Message = msgBuilder.BuildRequestBodySummaryDetailMessage(reqBodySummaryDetail, mediaType)
			s.Details = append(s.Details, reqBodySummaryDetail)
		}
	}

	s.Message = msgBuilder.BuildRequestBodySummaryMessage(s)

	return s
}

func buildResponsesSummary(c *ChangedAPIResponse, msgBuilder diff.SummaryMessageBuilder) *diff.ResponsesSummary {
	s := &diff.ResponsesSummary{
		Breaking: c.Incompatible,
	}

	for statusCode, response := range c.Increased {
		resSummaryDetail := &diff.ResponsesSummaryDetail{
			Res:         response,
			Name:        statusCode,
			Description: getStringSafely(response.Description),
			Action:      diff.ActionAdded,
		}
		resSummaryDetail.Message = msgBuilder.BuildResponsesSummaryDetailMessage(resSummaryDetail, statusCode)
		s.Details = append(s.Details, resSummaryDetail)
	}

	for statusCode, response := range c.Missing {
		resSummaryDetail := &diff.ResponsesSummaryDetail{
			Res:         response,
			Name:        statusCode,
			Description: getStringSafely(response.Description),
			Action:      diff.ActionDeleted,
		}
		resSummaryDetail.Message = msgBuilder.BuildResponsesSummaryDetailMessage(resSummaryDetail, statusCode)
		s.Details = append(s.Details, resSummaryDetail)
	}

	for statusCode, changedResponse := range c.Changed {
		responsesSummaryDetail := &diff.ResponsesSummaryDetail{
			OldRes:      changedResponse.OldAPIResponse,
			NewRes:      changedResponse.NewAPIResponse,
			Name:        statusCode,
			Description: getStringSafely(changedResponse.NewAPIResponse.Description),
			Action:      diff.ActionModified,
			Breaking:    changedResponse.Incompatible,
		}

		if changedResponse.Headers != nil {
			for name, header := range changedResponse.Headers.Increased {
				responseSummaryDetail := &diff.ResponseSummaryDetail{
					Action:      diff.ActionAdded,
					Name:        name,
					Description: header.Description,
				}
				responseSummaryDetail.Message = msgBuilder.BuildResponseSummaryDetailMessage(responseSummaryDetail, "header", name)
				responsesSummaryDetail.Details = append(responsesSummaryDetail.Details, responseSummaryDetail)
			}
			for name, header := range changedResponse.Headers.Missing {
				responseSummaryDetail := &diff.ResponseSummaryDetail{
					Action:      diff.ActionDeleted,
					Name:        name,
					Description: header.Description,
				}
				responseSummaryDetail.Message = msgBuilder.BuildResponseSummaryDetailMessage(responseSummaryDetail, "header", name)
				responsesSummaryDetail.Details = append(responsesSummaryDetail.Details, responseSummaryDetail)
			}
			for name, changedHeader := range changedResponse.Headers.Changed {
				responseSummaryDetail := &diff.ResponseSummaryDetail{
					Action:      diff.ActionModified,
					Name:        name,
					Description: changedHeader.NewHeader.Description,
				}
				responseSummaryDetail.Message = msgBuilder.BuildResponseSummaryDetailMessage(responseSummaryDetail, "header", name)
				responsesSummaryDetail.Details = append(responsesSummaryDetail.Details, responseSummaryDetail)
			}
		}

		if changedResponse.Content != nil {
			for mediaTypeKey, mediaType := range changedResponse.Content.Increased {
				responsesSummaryDetailItem := &diff.ResponseSummaryDetail{
					Res: &openapi3.Response{
						Content: openapi3.Content{
							mediaTypeKey: mediaType,
						},
					},
					Name:   mediaTypeKey,
					Action: diff.ActionAdded,
				}
				responsesSummaryDetailItem.Message = msgBuilder.BuildResponseSummaryDetailMessage(responsesSummaryDetailItem, "content type", mediaTypeKey)
				responsesSummaryDetail.Details = append(responsesSummaryDetail.Details, responsesSummaryDetailItem)
			}
			for mediaTypeKey, mediaType := range changedResponse.Content.Missing {
				responsesSummaryDetailItem := &diff.ResponseSummaryDetail{
					Res: &openapi3.Response{
						Content: openapi3.Content{
							mediaTypeKey: mediaType,
						},
					},
					Name:   mediaTypeKey,
					Action: diff.ActionDeleted,
				}
				responsesSummaryDetailItem.Message = msgBuilder.BuildResponseSummaryDetailMessage(responsesSummaryDetailItem, "content type", mediaTypeKey)
				responsesSummaryDetail.Details = append(responsesSummaryDetail.Details, responsesSummaryDetailItem)
			}

			for mediaTypeKey, changedMediaType := range changedResponse.Content.Changed {
				if changedMediaType.Schema != nil {

					resSummaryDetail := &diff.ResponseSummaryDetail{
						Name:     mediaTypeKey,
						Action:   diff.ActionModified,
						Breaking: changedMediaType.Incompatible,
					}

					for k, p := range changedMediaType.Schema.IncreasedProperties {
						pSummary := &diff.PropertiesSummary{
							Name:        k,
							Type:        p.Type,
							Description: p.Description,
							Action:      diff.ActionAdded,
						}
						pSummary.Message = msgBuilder.BuildPropertiesSummaryMessage(pSummary, 1)
						resSummaryDetail.Properties = append(resSummaryDetail.Properties, pSummary)
					}

					for k, p := range changedMediaType.Schema.MissingProperties {
						pSummary := &diff.PropertiesSummary{
							Name:        k,
							Type:        p.Type,
							Description: p.Description,
							Action:      diff.ActionDeleted,
						}
						pSummary.Message = msgBuilder.BuildPropertiesSummaryMessage(pSummary, 1)
						resSummaryDetail.Properties = append(resSummaryDetail.Properties, pSummary)
					}

					for k, changedProperty := range changedMediaType.Schema.ChangedProperties {
						pSummary := &diff.PropertiesSummary{
							Name:        k,
							Type:        changedProperty.NewSchema.Type,
							Description: changedProperty.NewSchema.Description,
							Action:      diff.ActionModified,
						}
						pSummary.Nested = buildNestedPropertiesSummary(changedProperty, msgBuilder, 2)
						pSummary.Message = msgBuilder.BuildPropertiesSummaryMessage(pSummary, 1)
						resSummaryDetail.Properties = append(resSummaryDetail.Properties, pSummary)
					}

					if changedMediaType.Schema.Type == "array" && changedMediaType.Schema.Items != nil {
						pSummary := &diff.PropertiesSummary{
							Action: diff.ActionModified,
							Name:   "items",
							Type:   changedMediaType.Schema.Type,
							Group:  "items",
						}
						pSummary.Nested = buildNestedPropertiesSummary(changedMediaType.Schema.Items, msgBuilder, 2)
						resSummaryDetail.Properties = append(resSummaryDetail.Properties, pSummary)
					}

					resSummaryDetail.Message = msgBuilder.BuildResponseSummaryDetailMessage(resSummaryDetail, "content type", mediaTypeKey)
					responsesSummaryDetail.Details = append(responsesSummaryDetail.Details, resSummaryDetail)
				}
			}
		}

		responsesSummaryDetail.Message = msgBuilder.BuildResponsesSummaryDetailMessage(responsesSummaryDetail, statusCode)
		s.Details = append(s.Details, responsesSummaryDetail)
	}

	s.Message = msgBuilder.BuildResponsesSummaryMessage(s)

	return s
}

func buildSecuritySummary(c *ChangedSecurityRequirements, msgBuilder diff.SummaryMessageBuilder) *diff.SecuritySummary {
	s := &diff.SecuritySummary{}

	for _, secReq := range c.Increased {
		if secReq != nil {
			for secReqKey, v := range *secReq {
				_ = v
				d := &diff.SecuritySummaryDetail{
					SecReq: secReq,
					Action: diff.ActionAdded,
					Name:   secReqKey,
				}
				d.Message = msgBuilder.BuildSecuritySummaryDetailMessage(d)
				s.Details = append(s.Details, d)
			}
		}
	}

	for _, secReq := range c.Missing {
		if secReq != nil {
			for secReqKey, v := range *secReq {
				_ = v
				d := &diff.SecuritySummaryDetail{
					SecReq: secReq,
					Action: diff.ActionDeleted,
					Name:   secReqKey,
				}
				d.Message = msgBuilder.BuildSecuritySummaryDetailMessage(d)
				s.Details = append(s.Details, d)
			}
		}
	}

	for _, changedSecReq := range c.Changed {
		if changedSecReq.NewSecurityRequirement != nil {
			for secReqKey, v := range *changedSecReq.NewSecurityRequirement {
				_ = v
				d := &diff.SecuritySummaryDetail{
					OldSecReq: changedSecReq.OldSecurityRequirement,
					NewSecReq: changedSecReq.NewSecurityRequirement,
					Action:    diff.ActionModified,
					Name:      secReqKey,
					Breaking:  changedSecReq.Incompatible,
				}
				d.Message = msgBuilder.BuildSecuritySummaryDetailMessage(d)
				s.Details = append(s.Details, d)
			}
		}
	}

	s.Message = msgBuilder.BuildSecuritySummaryMessage(s)
	return s
}

func buildNestedPropertiesSummary(schema *ChangedSchema, msgBuilder diff.SummaryMessageBuilder, indentLevel int) []*diff.PropertiesSummary {
	var nested []*diff.PropertiesSummary
	for k, p := range schema.IncreasedProperties {
		pSummaryNested := &diff.PropertiesSummary{
			Name:        k,
			Type:        p.Type,
			Description: p.Description,
			Action:      diff.ActionAdded,
		}
		pSummaryNested.Message = msgBuilder.BuildPropertiesSummaryMessage(pSummaryNested, indentLevel)
		nested = append(nested, pSummaryNested)
	}
	for k, p := range schema.MissingProperties {
		pSummaryNested := &diff.PropertiesSummary{
			Name:        k,
			Type:        p.Type,
			Description: p.Description,
			Action:      diff.ActionDeleted,
		}
		pSummaryNested.Message = msgBuilder.BuildPropertiesSummaryMessage(pSummaryNested, indentLevel)
		nested = append(nested, pSummaryNested)
	}
	for k, changedProperty := range schema.ChangedProperties {
		pSummaryNested := &diff.PropertiesSummary{
			Name:        k,
			Type:        changedProperty.NewSchema.Type,
			Description: changedProperty.NewSchema.Description,
			Action:      diff.ActionModified,
		}
		pSummaryNested.Message = msgBuilder.BuildPropertiesSummaryMessage(pSummaryNested, indentLevel)
		nested = append(nested, pSummaryNested)
	}
	return nested
}

func getStringSafely(sPtr *string) (s string) {
	if sPtr != nil {
		return *sPtr
	}
	return
}
