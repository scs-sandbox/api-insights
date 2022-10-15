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
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type infoResource struct {
	config   *shared.AppConfig
	validate *validator.Validate
	info     *models.Info
}

// Register the API
// prefix: /v1/apiregistry/info
func (r *infoResource) Register(config *shared.AppConfig, container *restful.Container, prefix string) {
	ws := &restful.WebService{}
	ws.Path(prefix).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON).ApiVersion(config.AppVersion).Doc("APIs for Info.")

	var info models.Info

	ws.Route(
		ws.GET("").
			To(r.getInfo).
			Do(shared.RouteReturns(info, http.StatusOK)).
			Do(shared.RouteReturns(nil, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites(info)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"info"}).
			Doc("Get server info"))

	container.Add(ws)
}

func (r *infoResource) getInfo(req *restful.Request, res *restful.Response) {
	shared.LogDebugf("get request to get info: %+v", r.info)
	_ = res.WriteEntity(r.info)
}
