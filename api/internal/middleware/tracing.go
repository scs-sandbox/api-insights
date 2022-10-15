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
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	restful "github.com/emicklei/go-restful/v3"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// NewTracingMiddleware - Add tracing capability to HTTP Request Object
func NewTracingMiddleware() restful.FilterFunction {
	tracer := opentracing.GlobalTracer()
	return func(req *restful.Request, res *restful.Response, chain *restful.FilterChain) {
		shared.LogDebugf("enter tracing middleware for request: %s", req.Request.URL.Path)
		httpRequest := req.Request
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(httpRequest.Header))
		span := tracer.StartSpan(req.SelectedRoutePath(), ext.RPCServerOption(spanCtx))
		defer span.Finish()
		for k, v := range req.PathParameters() {
			span.SetTag(k, v)
		}
		ctx := opentracing.ContextWithSpan(httpRequest.Context(), span)
		req.Request = httpRequest.WithContext(ctx)
		chain.ProcessFilter(req, res)
	}
}
