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

package models

import (
	"github.com/cisco-developer/api-insights/api/pkg/utils"
	"github.com/emicklei/go-restful/v3"
	"net/http"
	"strings"
)

type (
	// ResourceAction specifies a performable action on a ResourceKind, e.g. create, update, delete, get, list.
	ResourceAction string
	// ResourceKind specifies a resource object, e.g. organization, service.
	ResourceKind string
)

const (
	ResourceActionCreate ResourceAction = "create"
	ResourceActionUpdate ResourceAction = "update"
	ResourceActionDelete ResourceAction = "delete"
	ResourceActionGet    ResourceAction = "get"
	ResourceActionList   ResourceAction = "list"

	ResourceKindOrganizations ResourceKind = "organizations"
	ResourceKindOrganization  ResourceKind = "organization"
	ResourceKindServices      ResourceKind = "services"
	ResourceKindService       ResourceKind = "service"
	ResourceKindSpecs         ResourceKind = "specs"
	ResourceKindSpec          ResourceKind = "spec"
)

type Resource struct {
	Kind           ResourceKind
	Action         ResourceAction
	ID             *string
	Subresource    *Resource
	IsRootResource bool
	IsSubresource  bool
}

func NewResourceFromReq(req *restful.Request) *Resource {
	resource := &Resource{
		IsRootResource: true,
	}

	selectedRouteRelativePathParts := strings.Split(strings.Trim(req.SelectedRoutePath(), "/"), "/")[2:]
	switch selectedRouteRelativePathParts[0] { // services, organizations
	case string(ResourceKindOrganizations):
		resource.Kind = ResourceKindOrganizations
	case string(ResourceKindServices):
		resource.Kind = ResourceKindServices
	}

	// Build resource.Subresource.
	if len(selectedRouteRelativePathParts) > 1 {
		selectedRouteRelativePath := "/" + strings.Join(selectedRouteRelativePathParts, "/")
		switch resource.Kind {
		case ResourceKindOrganizations:
			if strings.HasPrefix(selectedRouteRelativePath, "/organizations/{id}") {
				resource.Subresource = &Resource{
					Kind:          ResourceKindOrganization,
					ID:            utils.StringPtr(req.PathParameter("id")),
					IsSubresource: true,
				}
			}
		case ResourceKindServices:
			if strings.HasPrefix(selectedRouteRelativePath, "/services/{id}") {
				resource.Subresource = &Resource{
					Kind:          ResourceKindService,
					ID:            utils.StringPtr(req.PathParameter("id")),
					IsSubresource: true,
				}
			}
			if strings.HasPrefix(selectedRouteRelativePath, "/services/{id}/specs") {
				resource.Subresource.Subresource = &Resource{
					Kind:          ResourceKindSpecs,
					IsSubresource: true,
				}
			}
			if strings.HasPrefix(selectedRouteRelativePath, "/services/{id}/specs/{specID}") {
				resource.Subresource.Subresource.Subresource = &Resource{
					Kind:          ResourceKindSpec,
					ID:            utils.StringPtr(req.PathParameter("specID")),
					IsSubresource: true,
				}
			}
		}
	}

	switch req.Request.Method {
	case http.MethodPost:
		resource.Action = ResourceActionCreate
	case http.MethodPut, http.MethodPatch:
		resource.Action = ResourceActionUpdate
	case http.MethodDelete:
		resource.Action = ResourceActionDelete
	case http.MethodGet:
		// HTTP method GET translates to either ResourceActionGet or ResourceActionList.
		if strings.HasSuffix(req.SelectedRoutePath(), "}") {
			resource.Action = ResourceActionGet
		} else {
			resource.Action = ResourceActionList
		}
	}
	return resource
}

func (r Resource) IsCreateOrganization() bool {
	return r.IsRootResource &&
		r.Kind == ResourceKindOrganizations &&
		r.Action == ResourceActionCreate &&
		r.Subresource == nil
}

func (r Resource) IsUpdateOrganization() bool {
	return r.IsRootResource &&
		r.Kind == ResourceKindOrganizations &&
		r.Action == ResourceActionUpdate &&
		r.Subresource != nil &&
		r.Subresource.Kind == ResourceKindOrganization &&
		r.Subresource.Subresource == nil
}

func (r Resource) IsDeleteOrganization() bool {
	return r.IsRootResource &&
		r.Kind == ResourceKindOrganizations &&
		r.Action == ResourceActionDelete &&
		r.Subresource != nil &&
		r.Subresource.Kind == ResourceKindOrganization &&
		r.Subresource.Subresource == nil
}

func (r Resource) IsGetOrganization() bool {
	return r.IsRootResource &&
		r.Kind == ResourceKindOrganizations &&
		r.Action == ResourceActionGet &&
		r.Subresource != nil &&
		r.Subresource.Kind == ResourceKindOrganization &&
		r.Subresource.Subresource == nil
}

func (r Resource) IsListOrganizations() bool {
	return r.IsRootResource &&
		r.Kind == ResourceKindOrganization &&
		r.Action == ResourceActionList &&
		r.Subresource == nil
}

func (r Resource) IsCreateService() bool {
	return r.IsRootResource &&
		r.Kind == ResourceKindServices &&
		r.Action == ResourceActionCreate &&
		r.Subresource == nil
}

func (r Resource) IsUpdateService() bool {
	return r.IsRootResource &&
		r.Kind == ResourceKindServices &&
		r.Action == ResourceActionUpdate &&
		r.Subresource != nil &&
		r.Subresource.Kind == ResourceKindService &&
		r.Subresource.Subresource == nil
}

func (r Resource) IsDeleteService() bool {
	return r.IsRootResource &&
		r.Kind == ResourceKindServices &&
		r.Action == ResourceActionDelete &&
		r.Subresource != nil &&
		r.Subresource.Kind == ResourceKindService &&
		r.Subresource.Subresource == nil
}

func (r Resource) IsGetService() bool {
	return r.IsRootResource &&
		r.Kind == ResourceKindServices &&
		r.Action == ResourceActionGet &&
		r.Subresource != nil &&
		r.Subresource.Kind == ResourceKindService &&
		r.Subresource.Subresource == nil
}

func (r Resource) IsListServices() bool {
	return r.IsRootResource &&
		r.Kind == ResourceKindServices &&
		r.Action == ResourceActionList &&
		r.Subresource == nil
}

func (r Resource) IsReadOnlyAction() bool {
	switch r.Action {
	case ResourceActionList, ResourceActionGet:
		return true
	}
	return false
}

func (r Resource) IsSpecs() bool {
	return r.IsRootResource &&
		r.Kind == ResourceKindServices &&
		r.Subresource != nil &&
		r.Subresource.Kind == ResourceKindSpecs
}

func (r Resource) IsSpec() bool {
	return r.IsRootResource &&
		r.Kind == ResourceKindServices &&
		r.Subresource != nil &&
		r.Subresource.Kind == ResourceKindSpecs &&
		r.Subresource.Subresource != nil &&
		r.Subresource.Subresource.Kind == ResourceKindSpec
}
