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
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"runtime"
	"strings"
)

const APIInsightsTraceID = "api_insights_trace_id"

// InitTracer returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func InitTracer(service string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: false,
		},
		Headers: &jaeger.HeadersConfig{
			TraceContextHeaderName: APIInsightsTraceID,
		},
	}
	cfg.Headers.ApplyDefaults()

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.NullLogger))
	if err != nil {
		LogErrorf("ERROR: cannot init Jaeger: %s\n", err.Error())
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}

func StartSpan(context context.Context, variables ...string) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(context, callerFunc(2))
	i := 0
	for ; i < len(variables)/2; i++ {
		span.SetTag(variables[i*2], variables[i*2+1])
	}
	if 2*i+1 == len(variables) {
		span.SetTag(variables[2*i], "")
	}
	return span, ctx
}

func callerFunc(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	p := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	return p[len(p)-1]
}
