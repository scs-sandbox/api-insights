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
	"errors"
	"fmt"
	"gorm.io/datatypes"
	"regexp"
)

const (
	SuperAdminRole        = "api-insights-admin"
	SuperAdminServiceRole = "api-insights-admin-service"

	// RoleTypeOrgAdmin role type represents roles that can manage other users and services of an org.
	RoleTypeOrgAdmin RoleType = "org-admin"
	// RoleTypeOrgDeveloper role type represents roles that can upload specs and view services of an org.
	RoleTypeOrgDeveloper RoleType = "org-developer"

	roleTypeOrgAdminRegexpExpr            = "api-insights-(.*)-admin"
	roleTypeOrgDeveloperRegexpExpr        = "api-insights-(.*)-developer"
	roleTypeOrgAdminServiceRegexpExpr     = "api-insights-(.*)-admin-service"
	roleTypeOrgDeveloperServiceRegexpExpr = "api-insights-(.*)-developer-service"
)

var (
	roleTypeOrgAdminRegexp, _            = regexp.Compile(roleTypeOrgAdminRegexpExpr)
	roleTypeOrgDeveloperRegexp, _        = regexp.Compile(roleTypeOrgDeveloperRegexpExpr)
	roleTypeOrgAdminServiceRegexp, _     = regexp.Compile(roleTypeOrgAdminServiceRegexpExpr)
	roleTypeOrgDeveloperServiceRegexp, _ = regexp.Compile(roleTypeOrgDeveloperServiceRegexpExpr)
	MatchRoleTypeOrgAdmin                = func(s string) bool { return roleTypeOrgAdminRegexp.MatchString(s) }
	MatchRoleTypeOrgDeveloper            = func(s string) bool { return roleTypeOrgDeveloperRegexp.MatchString(s) }
	MatchRoleTypeOrgAdminService         = func(s string) bool { return roleTypeOrgAdminServiceRegexp.MatchString(s) }
	MatchRoleTypeOrgDeveloperService     = func(s string) bool { return roleTypeOrgDeveloperServiceRegexp.MatchString(s) }
)

// RoleType represents an organization role type
type RoleType string

// Roles represents organization roles
type Roles struct {
	Roles datatypes.JSONMap `json:"roles" gorm:"column:roles"`
	roles map[RoleType]map[string]struct{}
}

func NewOrganizationRoles() Roles {
	return Roles{
		Roles: map[string]interface{}{
			string(RoleTypeOrgAdmin):     []string{},
			string(RoleTypeOrgDeveloper): []string{},
		},
		roles: map[RoleType]map[string]struct{}{
			RoleTypeOrgAdmin:     {},
			RoleTypeOrgDeveloper: {},
		},
	}
}

func (r *Roles) RoleTypeSetMap() map[RoleType]map[string]struct{} {
	return r.roles
}

func (r *Roles) OrgAdminRoleSet() map[string]struct{} {
	return r.roleSetBy(RoleTypeOrgAdmin)
}

func (r *Roles) OrgDeveloperRoleSet() map[string]struct{} {
	return r.roleSetBy(RoleTypeOrgDeveloper)
}

func (r *Roles) roleSetBy(roleType RoleType) (set map[string]struct{}) {
	if r.roles != nil {
		set = r.roles[roleType]
		return
	}
	value, ok := r.Roles[string(roleType)]
	if ok {
		set = map[string]struct{}{}
		roles := value.([]string)
		for _, role := range roles {
			set[role] = struct{}{}
		}
	}
	return
}

func (r *Roles) AddOrgAdminRole(role string) {
	r.addRole(RoleTypeOrgAdmin, role)
}

func (r *Roles) AddOrgDeveloperRole(role string) {
	r.addRole(RoleTypeOrgDeveloper, role)
}

func (r *Roles) addRole(roleType RoleType, role string) {
	if r.Roles == nil {
		r.Roles = datatypes.JSONMap{
			string(RoleTypeOrgAdmin):     []string{},
			string(RoleTypeOrgDeveloper): []string{},
		}
	}
	if r.roles == nil {
		r.roles = map[RoleType]map[string]struct{}{
			RoleTypeOrgAdmin:     {},
			RoleTypeOrgDeveloper: {},
		}
	}
	roleTypeStr := string(roleType)
	_, ok := r.Roles[roleTypeStr]
	if !ok {
		r.Roles[roleTypeStr] = []string{}
	}
	_, ok = r.roles[roleType]
	if !ok {
		r.roles[roleType] = map[string]struct{}{}
	}
	_, ok = r.Roles[roleTypeStr].([]string)
	if ok {
		r.Roles[roleTypeStr] = append(r.Roles[roleTypeStr].([]string), role)
	}
	_, ok = r.roles[roleType][role]
	if !ok {
		r.roles[roleType][role] = struct{}{}
	}
}

type AccessDataFilters []interface{}

func NewOrgServiceAccessFilter() *OrgServiceAccessDataFilter {
	return &OrgServiceAccessDataFilter{
		OrgAccessFilter:     &OrgAccessFilter{},
		ServiceAccessFilter: &ServiceAccessFilter{},
	}
}

type OrgServiceAccessDataFilter struct {
	OrgAccessFilter     *OrgAccessFilter
	ServiceAccessFilter *ServiceAccessFilter
}

func (f OrgServiceAccessDataFilter) CanCreateService(orgID string) (bool, error) {
	if f.ServiceAccessFilter == nil {
		return true, nil
	}
	for _, accessibleOrgID := range f.ServiceAccessFilter.AccessibleOrganizationIDs {
		if accessibleOrgID == orgID {
			return true, nil
		}
	}
	return false, &UnauthorizedResourceAccessError{
		ResourceKind:   ResourceKindServices,
		ResourceAction: ResourceActionCreate,
	}
}

func (f OrgServiceAccessDataFilter) CanListServices() (authorized bool, orgIDs []string, err error) {
	if f.ServiceAccessFilter == nil {
		return true, nil, nil
	}
	return true, f.ServiceAccessFilter.AccessibleOrganizationIDs, nil
}

func (f OrgServiceAccessDataFilter) CanCreateOrg(orgID string) (bool, error) {
	if f.OrgAccessFilter == nil {
		return true, nil
	}
	for _, accessibleOrgID := range f.OrgAccessFilter.AccessibleOrganizationIDs {
		if accessibleOrgID == orgID {
			return true, nil
		}
	}
	return false, &UnauthorizedResourceAccessError{
		ResourceKind:   ResourceKindOrganizations,
		ResourceAction: ResourceActionCreate,
	}
}

func (f OrgServiceAccessDataFilter) CanListOrgs() (authorized bool, orgIDs []string, err error) {
	if f.OrgAccessFilter == nil {
		return true, nil, nil
	}
	return true, f.OrgAccessFilter.AccessibleOrganizationIDs, nil
}

type OrgAccessFilter struct {
	// AccessibleOrganizationIDs contains a list of OrganizationIDs that are authorized to the user for all operations.
	AccessibleOrganizationIDs []string
}

type ServiceAccessFilter struct {
	// AccessibleOrganizationIDs contains a list of OrganizationIDs that are authorized to the user for creating & listing.
	AccessibleOrganizationIDs []string
}

type UnauthorizedResourceAccessError struct {
	ResourceKind        ResourceKind   // organization, service
	ResourceAction      ResourceAction // create, list, get, update, delete
	ResourceID          string         // organization/service/spec ID
	UnauthorizedMessage string
}

func (e UnauthorizedResourceAccessError) Error() string {
	msg := fmt.Sprintf("Unauthorized to %s %s", e.ResourceAction, e.ResourceKind)
	if e.ResourceID != "" {
		msg += fmt.Sprintf("(%s)", e.ResourceID)
	}
	if e.UnauthorizedMessage != "" {
		msg += fmt.Sprintf(" (%s)", e.UnauthorizedMessage)
	}
	return msg
}

var (
	ErrResourceIDExpected = errors.New("access: resource ID expected")
)
