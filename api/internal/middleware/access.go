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

package middleware

import (
	"github.com/cisco-developer/api-insights/api/internal/access"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"github.com/emicklei/go-restful/v3"
	"net/http"
)

const (
	requestAttrAccessCheckerInput = "access_checker_input"
	requestAttrAccessDataFilters  = "access_data_filters"
)

func AccessCheckerInputFromReq(req *restful.Request) *access.Input {
	raw := req.Attribute(requestAttrAccessCheckerInput)
	v, _ := raw.(*access.Input)
	return v
}

func AccessDataFiltersFromReq(req *restful.Request) models.AccessDataFilters {
	raw := req.Attribute(requestAttrAccessDataFilters)
	v, _ := raw.(models.AccessDataFilters)
	return v
}

// ResourceAccessChecker serves as an authorization middleware filter for the service(s) & organization(s) resources (i.e. /services/*, /organizations/*).
// Consumes the access.Input in the request context & sets models.AccessDataFilters in the request context to be consumed by the DAO layer.
func ResourceAccessChecker(accessChecker access.Checker) restful.FilterFunction {
	return func(req *restful.Request, res *restful.Response, chain *restful.FilterChain) {
		log := shared.LoggerFromContext(req.Request.Context())
		accessInput := AccessCheckerInputFromReq(req)
		if accessInput == nil {
			chain.ProcessFilter(req, res)
			return
		}
		accessOutput, err := accessChecker.CheckAccess(req.Request.Context(), accessInput)
		if err != nil {
			log.Errorf("middleware.ResourceAccessChecker: failed to check access: %v", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		if accessOutput.Denied() {
			res.WriteHeader(http.StatusForbidden)
			return
		}
		req.SetAttribute(requestAttrAccessDataFilters, accessOutput.AllowWithFilters)
		chain.ProcessFilter(req, res)
	}
}
