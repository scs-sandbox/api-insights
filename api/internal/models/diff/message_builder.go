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

package diff

type SummaryMessageBuilder interface {
	resultSummaryMessageBuilder

	modifiedSummaryMessageBuilder

	parametersSummaryMessageBuilder
	parameterSummaryMessageBuilder

	requestBodySummaryMessageBuilder
	requestBodySummaryDetailMessageBuilder

	responsesSummaryMessageBuilder
	responsesSummaryDetailMessageBuilder
	responseSummaryDetailMessageBuilder

	securitySummaryMessageBuilder
	securitySummaryDetailMessageBuilder

	propertiesSummaryMessageBuilder
}

type resultSummaryMessageBuilder interface {
	BuildResultSummaryMessage(result *JSONResult) string
}

type modifiedSummaryMessageBuilder interface {
	BuildModifiedSummaryMessage(s *ModifiedSummary) string
}

type (
	parametersSummaryMessageBuilder interface {
		BuildParametersSummaryMessage(s *ParametersSummary) string
	}
	parameterSummaryMessageBuilder interface {
		BuildParameterSummaryMessage(s *ParameterSummary) string
	}
)

type (
	requestBodySummaryMessageBuilder interface {
		BuildRequestBodySummaryMessage(s *RequestBodySummary) string
	}
	requestBodySummaryDetailMessageBuilder interface {
		BuildRequestBodySummaryDetailMessage(d *RequestBodySummaryDetail, contentType string) string
	}
)

type (
	responsesSummaryMessageBuilder interface {
		BuildResponsesSummaryMessage(s *ResponsesSummary) string
	}
	responsesSummaryDetailMessageBuilder interface {
		BuildResponsesSummaryDetailMessage(d *ResponsesSummaryDetail, statusCode string) string
	}
	responseSummaryDetailMessageBuilder interface {
		BuildResponseSummaryDetailMessage(d *ResponseSummaryDetail, key, value string) string
	}
)

type (
	securitySummaryDetailMessageBuilder interface {
		BuildSecuritySummaryDetailMessage(d *SecuritySummaryDetail) string
	}
	securitySummaryMessageBuilder interface {
		BuildSecuritySummaryMessage(s *SecuritySummary) string
	}
)

type propertiesSummaryMessageBuilder interface {
	BuildPropertiesSummaryMessage(s *PropertiesSummary, indentLevel int) string
}
