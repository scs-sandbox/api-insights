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
	"github.com/cisco-developer/api-insights/api/internal/db"
	"github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"time"
)

type analyzerResource struct {
	config          *shared.AppConfig
	dao             db.AnalyzerDAO
	validate        *validator.Validate
	analyzerRuleDao db.AnalyzerRuleDAO
}

// Register the API
// prefix: /v1/apiregistry/analyzers
func (r *analyzerResource) Register(config *shared.AppConfig, container *restful.Container, prefix string) {
	ws := &restful.WebService{}
	ws.Path(prefix).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON).ApiVersion(config.AppVersion).Doc("APIs for Analyzer.")

	var a analyzer.Analyzer
	var as []analyzer.Analyzer
	var ar analyzer.Rule
	var ars []analyzer.Rule
	var id = ws.PathParameter("id", "unique identifier (UUID or Name ID) for analyzer.").DataType("string")
	var ruleID = ws.PathParameter("ruleID", "unique identifier (UUID or Name ID) for analyzer rule.").DataType("string")
	var nameID = ws.PathParameter("name_id", "unique identifier (UUID or Name ID) for analyzer rule.").DataType("string")
	var tags = ws.QueryParameter("tags", "tags for getting analyzers").DataType("string")
	var q = ws.QueryParameter("q", "searching criteria for analyzers").DataType("string")
	var limit = ws.QueryParameter("limit", "max items to return at one time").DataType("string")
	var offset = ws.QueryParameter("offset", "starting offset").DataType("string")
	var status = ws.QueryParameter("status", "analyzer status, e.g. active").DataType("string")
	var withRules = ws.QueryParameter("withRules", "flag indicating if rules should be included, by default false").DataType("boolean")

	ws.Route(
		ws.GET("").
			To(r.list).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteParams(offset)).
			Do(shared.RouteParams(limit, max)).
			Do(shared.RouteReturns(&as, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites(as)).
			Do(shared.RouteParams(q)).
			Do(shared.RouteParams(tags)).
			Do(shared.RouteParams(status)).
			Do(shared.RouteParams(withRules)).
			Do(shared.RouteParams(sort, sortOrder)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"analyzer"}).
			Notes("List all analyzers with specified tags."))

	ws.Route(
		ws.GET("/{id}").
			To(r.get).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(a, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError)).
			Do(shared.RouteWrites(a)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"analyzer"}).
			Notes("Get a analyzer with specified id"))

	ws.Route(
		ws.POST("").
			To(r.save).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(a, http.StatusOK, http.StatusCreated)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError)).
			Do(shared.RouteReads(a, "analyzer"), shared.RouteWrites(a)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"analyzer"}).
			Notes("Create a new analyzer"))

	ws.Route(
		ws.PUT("/{id}").
			To(r.save).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(a, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(a, "analyzer"), shared.RouteWrites(a)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"analyzer"}).
			Notes("Update existing analyzer"))

	ws.Route(
		ws.DELETE("/{id}").
			To(r.delete).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(nil, http.StatusNoContent)).
			Do(shared.RouteReturns(&se, http.StatusNotFound, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites(nil)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"analyzer"}).
			Notes("Delete existing analyzer"))

	//
	// Non-CRUD routes
	//

	ws.Route(
		ws.GET("/{id}/rules").
			To(r.listRules).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteParams(offset)).
			Do(shared.RouteParams(limit, max)).
			Do(shared.RouteReturns(ars, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites(ars)).
			Do(shared.RouteParams(id)).
			Do(shared.RouteParams(q)).
			Do(shared.RouteParams(tags)).
			Do(shared.RouteParams(nameID)).
			Do(shared.RouteParams(sort, sortOrder)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"analyzer"}).
			Notes("List all analyzer rules with specified tags."))

	ws.Route(
		ws.GET("/{id}/rules/{ruleID}").
			To(r.getRule).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(ar, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError)).
			Do(shared.RouteWrites(ar)).
			Do(shared.RouteParams(id)).
			Do(shared.RouteParams(ruleID)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"analyzer"}).
			Notes("Get a analyzer rule with specified id"))

	ws.Route(
		ws.POST("/{id}/rules").
			To(r.saveRule).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(ar, http.StatusOK, http.StatusCreated)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError)).
			Do(shared.RouteReads(ar, "analyzer rule"), shared.RouteWrites(ar)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"analyzer"}).
			Notes("Create a new analyzer rule"))

	ws.Route(
		ws.PUT("/{id}/rules/{ruleID}").
			To(r.saveRule).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(ar, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(ar, "analyzer rule"), shared.RouteWrites(ar)).
			Do(shared.RouteParams(id)).
			Do(shared.RouteParams(ruleID)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"analyzer"}).
			Notes("Update existing analyzer rule"))

	ws.Route(
		ws.DELETE("/{id}/rules/{ruleID}").
			To(r.deleteRule).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(nil, http.StatusNoContent)).
			Do(shared.RouteReturns(&se, http.StatusNotFound, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites(nil)).
			Do(shared.RouteParams(id)).
			Do(shared.RouteParams(ruleID)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"analyzer"}).
			Notes("Delete existing analyzer rule"))

	ws.Route(
		ws.POST("/{id}/rules/import").
			To(r.importRules).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(nil, http.StatusNoContent)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(ars, "analyzer rules"), shared.RouteWrites(nil)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"analyzer"}).
			Notes("Import a list of analyzer rules"))

	container.Add(ws)
}

func (r *analyzerResource) list(req *restful.Request, res *restful.Response) {
	shared.LogDebugf("get request to list analyzers.")

	var filter = &db.ListFilter{Model: &analyzer.Analyzer{}}
	err := filter.From(req)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(filter.Sorters) == 0 {
		filter.Sorters = append(filter.Sorters, &db.Sorter{
			Order: "asc",
			Field: "position",
		})
	}

	analyzers, err := r.dao.List(req.Request.Context(), filter, parseQueryBool(req, "withRules"))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	shared.LogDebugf("total %v analyzer(s) returned", len(analyzers))
	_ = res.WriteEntity(analyzers)
}

func (r *analyzerResource) get(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")
	shared.LogDebugf("get request to retrieve analyzer: %v", id)

	analyzer, err := r.dao.Get(req.Request.Context(), id)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	_ = res.WriteEntity(analyzer)
}

func (r *analyzerResource) save(req *restful.Request, res *restful.Response) {
	analyzer := &analyzer.Analyzer{}
	id := req.PathParameter("id")
	err := req.ReadEntity(analyzer)

	if err != nil {
		shared.LogErrorf("failed to get analyzer from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	shared.LogDebugf("get request to save analyzer with body: %#v", analyzer)

	if id == "" {
		analyzer.CreatedAt = time.Now().UTC()
	} else {
		existing, err := r.dao.Get(req.Request.Context(), id)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		analyzer.ID = existing.ID
		analyzer.CreatedAt = existing.CreatedAt
	}
	analyzer.UpdatedAt = time.Now().UTC()

	err = r.validate.Struct(analyzer)
	if err != nil {
		shared.LogErrorf("failed to validate analyzer %s - %v", analyzer.ID, err.Error())
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = r.dao.Save(req.Request.Context(), analyzer)
	if err != nil {
		handleError(res, err)
		return
	}
	if id == "" {
		res.Header().Add("Location", "/v1/apiregistry/analyzers/"+analyzer.ID)
		_ = res.WriteHeaderAndEntity(http.StatusCreated, analyzer)
	} else {
		_ = res.WriteEntity(analyzer)
	}
}

func (r *analyzerResource) delete(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")
	shared.LogDebugf("get request to delete analyzer: %v", id)

	err := r.dao.Delete(req.Request.Context(), id)
	if err == db.ErrNotFound {
		res.WriteHeader(http.StatusNotFound)
	} else if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	} else {
		res.WriteHeader(http.StatusNoContent)
	}
}

func (r *analyzerResource) listRules(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")

	shared.LogDebugf("get request to list analyzer rules.")

	var filter = &db.ListFilter{Model: &analyzer.Rule{}}
	err := filter.From(req)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	filter.Indexes["analyzer_name_id"] = id
	ars, err := r.analyzerRuleDao.List(req.Request.Context(), filter)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	shared.LogDebugf("total %v analyzer rule(s) returned", len(ars))
	_ = res.WriteEntity(ars)
}

func (r *analyzerResource) getRule(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("ruleID")
	shared.LogDebugf("get request to retrieve analyzer rule: %v", id)

	ar, err := r.analyzerRuleDao.Get(req.Request.Context(), id)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	_ = res.WriteEntity(ar)
}

func (r *analyzerResource) saveRule(req *restful.Request, res *restful.Response) {
	ar := &analyzer.Rule{}
	id := req.PathParameter("ruleID")
	err := req.ReadEntity(ar)

	if err != nil {
		shared.LogErrorf("failed to get analyzer rule from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	shared.LogDebugf("get request to save analyzer rule with body: %#v", ar)

	if _, err := r.dao.Get(req.Request.Context(), ar.AnalyzerNameID); err != nil {
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if id == "" {
		ar.CreatedAt = time.Now().UTC()
	} else {
		existing, err := r.analyzerRuleDao.Get(req.Request.Context(), id)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		ar.ID = existing.ID
		ar.CreatedAt = existing.CreatedAt
	}
	ar.UpdatedAt = time.Now().UTC()

	err = r.validate.Struct(ar)
	if err != nil {
		shared.LogErrorf("failed to validate analyzer rule %s - %v", ar.ID, err.Error())
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = r.analyzerRuleDao.Save(req.Request.Context(), ar)
	if err != nil {
		handleError(res, err)
		return
	}
	if id == "" {
		res.Header().Add("Location", "/v1/apiregistry/analyzers/"+ar.AnalyzerNameID+"/rules/"+ar.ID)
		_ = res.WriteHeaderAndEntity(http.StatusCreated, ar)
	} else {
		_ = res.WriteEntity(ar)
	}
}

func (r *analyzerResource) deleteRule(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("ruleID")
	shared.LogDebugf("get request to delete analyzer rule: %v", id)

	err := r.analyzerRuleDao.Delete(req.Request.Context(), id)
	if err == db.ErrNotFound {
		res.WriteHeader(http.StatusNotFound)
	} else if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	} else {
		res.WriteHeader(http.StatusNoContent)
	}
}

func (r *analyzerResource) importRules(req *restful.Request, res *restful.Response) {
	var ars []*analyzer.Rule
	err := req.ReadEntity(&ars)
	if err != nil {
		shared.LogErrorf("failed to get analyzer rules from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	shared.LogDebugf("get request to import %d analyzer rules", len(ars))

	err = r.analyzerRuleDao.Import(req.Request.Context(), ars)
	if err != nil {
		handleError(res, err)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}
