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

package shared

import (
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

// TODO: there should be no hard-coded superuser credentials. This is
// here explicitly because this capability should stop working in prod
const superuserAuthHeader = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7InRpdGxlIjoiIiwiZGlzcGxheU5hbWUiOiIiLCJmaXJzdE5hbWUiOiIiLCJsYXN0TmFtZSI6IiIsImZ1bGxOYW1lIjoiIiwiZW1haWwiOiIiLCJyb2xlIjoiIn0sInBlcm1zIjpbeyJjb21wYW55IjoiKiIsImRvbWFpbiI6IioiLCJyb2xlIjoiKiIsInVzZXIiOiIqIn1dfQ.IUa6OHlYquw-8k5YPYEVNc9US-LPl9Rt7Fms960MDTk"

func RouteAuthHeader(ws *restful.WebService) func(*restful.RouteBuilder) {
	param := ws.
		HeaderParameter("Authorization", "authorization header").
		Required(true).
		DataType("string").
		DefaultValue(superuserAuthHeader)

	return RouteParams(param)
}

func RouteParams(params ...*restful.Parameter) func(*restful.RouteBuilder) {
	return func(b *restful.RouteBuilder) {
		for _, param := range params {
			b.Param(param)
		}
	}
}

func RouteReturns(sample interface{}, responseCodes ...int) func(*restful.RouteBuilder) {
	headers := map[string]restful.Header{
		"TrackingID": NewHeader("TrackingID"),
		"Date":       NewHeader("Date"),
	}

	return func(b *restful.RouteBuilder) {
		for _, responseCode := range responseCodes {
			if responseCode == http.StatusCreated {
				headers["Location"] = NewHeader("Location")
			}

			if p := b.ParameterNamed("max"); p != nil && responseCode == http.StatusOK {
				headers["Link"] = NewHeader("Link")
			}

			b.ReturnsWithHeaders(responseCode, http.StatusText(responseCode), sample, headers)
		}
	}
}

func RouteReads(sample interface{}, description ...string) func(*restful.RouteBuilder) {
	return func(b *restful.RouteBuilder) { b.Reads(sample, description...) }
}

func RouteWrites(sample interface{}) func(*restful.RouteBuilder) {
	return func(b *restful.RouteBuilder) { b.Writes(sample) }
}

func NewHeader(description string) restful.Header {
	return restful.Header{Description: description, Items: &restful.Items{Type: "string"}}
}
