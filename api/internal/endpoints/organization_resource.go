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
	"github.com/cisco-developer/api-insights/api/internal/access"
	"github.com/cisco-developer/api-insights/api/internal/db"
	"github.com/cisco-developer/api-insights/api/internal/middleware"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"time"
)

type organizationResource struct {
	config        *shared.AppConfig
	dao           db.OrganizationDAO
	validate      *validator.Validate
	accessChecker access.Checker
}

// Register the API
// prefix: /v1/apiregistry/organizations
func (r *organizationResource) Register(config *shared.AppConfig, container *restful.Container, prefix string) {
	ws := &restful.WebService{}
	ws.Path(prefix).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON).ApiVersion(config.AppVersion).Doc("APIs for Organization.")
	ws.Filter(middleware.ResourceAccessChecker(r.accessChecker))

	var org models.Organization
	var orgs []models.Organization
	var orgPatch []models.OrganizationPatch
	var id = ws.PathParameter("id", "unique identifier (UUID or Name ID) for organization.").DataType("string")
	var tags = ws.QueryParameter("tags", "tags for getting organizations").DataType("string")
	var q = ws.QueryParameter("q", "searching criteria for organizations").DataType("string")
	var limit = ws.QueryParameter("limit", "max items to return at one time").DataType("string")
	var offset = ws.QueryParameter("offset", "starting offset").DataType("string")

	ws.Route(
		ws.GET("").
			To(r.list).
			Do(shared.RouteReturns(orgs, http.StatusOK)).
			Do(shared.RouteReturns(nil, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites(orgs)).
			Do(shared.RouteParams(q)).
			Do(shared.RouteParams(tags)).
			Do(shared.RouteParams(offset)).
			Do(shared.RouteParams(limit)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"organization"}).
			Doc("List all organizations with specified tags."))

	ws.Route(
		ws.GET("/{id}").
			To(r.get).
			Do(shared.RouteReturns(org, http.StatusOK)).
			Do(shared.RouteReturns(nil, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError)).
			Do(shared.RouteWrites(org)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"organization"}).
			Doc("Get a organization with specified id"))

	ws.Route(
		ws.POST("").
			To(r.save).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(org, http.StatusOK)).
			Do(shared.RouteReturns(org, http.StatusCreated)).
			Do(shared.RouteReturns(nil, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError)).
			Do(shared.RouteReads(org), shared.RouteWrites(org)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"organization"}).
			Doc("Create a new organization"))

	ws.Route(
		ws.PUT("/{id}").
			To(r.save).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(org, http.StatusOK)).
			Do(shared.RouteReturns(nil, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(org), shared.RouteWrites(org)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"organization"}).
			Doc("Update existing organization"))

	ws.Route(
		ws.PATCH("/{id}").
			To(r.patch).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(org, http.StatusOK)).
			Do(shared.RouteReturns(nil, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(orgPatch), shared.RouteWrites(org)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"organization"}).
			Doc("Patch existing organization"))

	ws.Route(
		ws.DELETE("/{id}").
			To(r.delete).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(nil, http.StatusNoContent)).
			Do(shared.RouteReturns(nil, http.StatusNotFound, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites(nil)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"organization"}).
			Doc("Delete existing organization"))

	container.Add(ws)
}

func (r *organizationResource) list(req *restful.Request, res *restful.Response) {
	shared.LogDebugf("get request to list organizations.")

	var filter = &db.ListFilter{Model: &models.Organization{}}
	err := filter.From(req)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	organizations, err := r.dao.List(req.Request.Context(), filter, orgServiceAccessDataFilterFromReq(req))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	shared.LogDebugf("total %v organization(s) returned", len(organizations))
	_ = res.WriteEntity(organizations)
}

func (r *organizationResource) get(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")
	shared.LogDebugf("get request to retrieve organization: %v", id)

	organization, err := r.dao.Get(req.Request.Context(), id)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	_ = res.WriteEntity(organization)
}

func (r *organizationResource) save(req *restful.Request, res *restful.Response) {
	organization := &models.Organization{}
	id := req.PathParameter("id")
	err := req.ReadEntity(organization)

	if err != nil {
		shared.LogErrorf("failed to get organization from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	shared.LogDebugf("get request to save organization with body: %#v", organization)

	if id == "" {
		organization.CreatedAt = time.Now().UTC()
	} else {
		existing, err := r.dao.Get(req.Request.Context(), id)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		organization.ID = existing.ID
		organization.CreatedAt = existing.CreatedAt
	}
	organization.UpdatedAt = time.Now().UTC()

	err = r.validate.Struct(organization)
	if err != nil {
		shared.LogErrorf("failed to validate organization %s - %v", organization.ID, err.Error())
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = r.dao.Save(req.Request.Context(), organization, orgServiceAccessDataFilterFromReq(req))
	if err != nil {
		handleError(res, err)
		return
	}
	if id == "" {
		res.Header().Add("Location", "/v1/apiregistry/organizations/"+organization.ID)
		_ = res.WriteHeaderAndEntity(http.StatusCreated, organization)
	} else {
		_ = res.WriteEntity(organization)
	}
}

func (r *organizationResource) patch(req *restful.Request, res *restful.Response) {
	patch := &models.OrganizationPatch{}
	id := req.PathParameter("id")
	err := req.ReadEntity(patch)

	if err != nil {
		shared.LogErrorf("failed to get organizationPatch from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	shared.LogDebugf("get request to patch organization with body: %#v", patch)

	org, err := r.dao.Get(req.Request.Context(), id)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	org.UpdatedAt = time.Now().UTC()

	if patch.Meta != nil && len(*patch.Meta) > 0 {
		if org.Meta == nil {
			org.Meta = map[string]interface{}{}
		}
		for k, v := range *patch.Meta {
			org.Meta[k] = v
		}
	}
	if patch.Contact != nil {
		org.Contact = patch.Contact
	}
	if patch.Description != nil {
		org.Description = *patch.Description
	}
	if patch.NameID != nil {
		org.NameID = *patch.NameID
	}
	if patch.Title != nil {
		org.Title = *patch.Title
	}
	if patch.Roles != nil {
		org.Roles = *patch.Roles
	}

	err = r.validate.Struct(org)
	if err != nil {
		shared.LogErrorf("failed to validate Organization %s - %v", org.ID, err.Error())
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = r.dao.Save(req.Request.Context(), org, orgServiceAccessDataFilterFromReq(req))
	if err != nil {
		handleError(res, err)
		return
	}
	_ = res.WriteEntity(org)
}

func (r *organizationResource) delete(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")
	shared.LogDebugf("get request to delete organization: %v", id)

	err := r.dao.Delete(req.Request.Context(), id)
	if err == db.ErrNotFound {
		res.WriteHeader(http.StatusNotFound)
	} else if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	} else {
		res.WriteHeader(http.StatusNoContent)
	}
}
