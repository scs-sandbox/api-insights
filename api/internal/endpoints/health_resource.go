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

package endpoints

import (
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
	"net/http"
)

type healthCheck struct {
	Status string `json:"Status"`
}

type HealthCheckFunction func(*restful.Request, *restful.Response) string

type HealthCheckResource struct {
	HealthCheckFunc HealthCheckFunction
}

func (ref *HealthCheckResource) Register(container *restful.Container, prefix string) {
	ws := &restful.WebService{}
	ws.Path(prefix).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(
		ws.GET("").
			To(ref.getHealth).
			Do(shared.RouteReturns(healthCheck{}, http.StatusOK)).
			Do(shared.RouteReturns(restful.ServiceError{}, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites(healthCheck{})).
			Metadata(restfulspec.KeyOpenAPITags, []string{"health"}).
			Notes("Returns the health status of the service"))

	container.Add(ws)
}

func (ref *HealthCheckResource) getHealth(request *restful.Request, response *restful.Response) {
	if ref.HealthCheckFunc != nil {
		_ = response.WriteEntity(healthCheck{Status: ref.HealthCheckFunc(request, response)})
	} else {
		_ = response.WriteEntity(healthCheck{Status: "OK"})
	}
}
