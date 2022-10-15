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

package result

import (
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi3"
)

// ChangedOpenAPI represents the (JSON) result returned by openapi-diff.
// Based on https://github.com/OpenAPITools/openapi-diff/blob/fd29e3ed0dde25055c7a360fe84e32107ae75ccf/core/src/main/java/org/openapitools/openapidiff/core/model/ChangedOpenApi.java.
type ChangedOpenAPI struct {
	OldSpecOpenAPI      *openapi3.T         `json:"-"`
	NewSpecOpenAPI      *openapi3.T         `json:"-"`
	NewEndpoints        []*Endpoint         `json:"newEndpoints,omitempty"`
	MissingEndpoints    []*Endpoint         `json:"missingEndpoints,omitempty"`
	DeprecatedEndpoints []*Endpoint         `json:"deprecatedEndpoints,omitempty"`
	ChangedOperations   []*ChangedOperation `json:"changedOperations,omitempty"`
	ChangedExtensions   *ChangedExtensions  `json:"changedExtensions,omitempty"`
	*DiffResult
}
type Endpoint struct {
	PathURL   string              `json:"pathUrl,omitempty"`
	Method    string              `json:"method,omitempty"`
	Summary   string              `json:"summary,omitempty"`
	Path      *openapi3.PathItem  `json:"path,omitempty"`
	Operation *openapi3.Operation `json:"operation,omitempty"`
}
type ChangedOperation struct {
	OldOperation         *openapi3.Operation          `json:"oldOperation,omitempty"`
	NewOperation         *openapi3.Operation          `json:"newOperation,omitempty"`
	PathURL              string                       `json:"pathUrl,omitempty"`
	HTTPMethod           string                       `json:"httpMethod,omitempty"`
	Summary              *ChangedMetadata             `json:"summary,omitempty"`
	Description          *ChangedMetadata             `json:"description,omitempty"`
	OperationID          *ChangedMetadata             `json:"operationId,omitempty"`
	Deprecated           bool                         `json:"deprecated,omitempty"`
	Parameters           *ChangedParameters           `json:"parameters,omitempty"`
	RequestBody          *ChangedRequestBody          `json:"requestBody,omitempty"`
	APIResponses         *ChangedAPIResponse          `json:"apiResponses,omitempty"`
	SecurityRequirements *ChangedSecurityRequirements `json:"securityRequirements,omitempty"`
	Extensions           *ChangedExtensions           `json:"extensions,omitempty"`
	*DiffResult
}

type ChangedParameters struct {
	//OldParameterList []*openapi3.Parameter `json:"oldParameterList,omitempty"`
	//NewParameterList []*openapi3.Parameter `json:"newParameterList,omitempty"`
	Increased []*openapi3.Parameter `json:"increased,omitempty"`
	Missing   []*openapi3.Parameter `json:"missing,omitempty"`
	Changed   []*ChangedParameter   `json:"changed,omitempty"`
	*DiffResult
}
type ChangedParameter struct {
	OldParameter          *openapi3.Parameter `json:"oldParameter,omitempty"`
	NewParameter          *openapi3.Parameter `json:"newParameter,omitempty"`
	Name                  string              `json:"name,omitempty"`
	In                    string              `json:"in,omitempty"`
	Description           *ChangedMetadata    `json:"description,omitempty"`
	ChangeStyle           bool                `json:"changeStyle,omitempty"`
	ChangeExplode         bool                `json:"changeExplode,omitempty"`
	ChangeAllowEmptyValue bool                `json:"changeAllowEmptyValue,omitempty"`
	Deprecated            bool                `json:"deprecated,omitempty"`
	ChangeRequired        bool                `json:"changeRequired,omitempty"`
	Schema                *ChangedSchema      `json:"schema,omitempty"`
	Content               *ChangedContent     `json:"content,omitempty"`
	Extensions            *ChangedExtensions  `json:"extensions,omitempty"`
	*DiffResult
}
type ChangedContent struct {
	//OldContent *openapi3.Content              `json:"oldContent,omitempty"`
	//NewContent *openapi3.Content              `json:"newContent,omitempty"`
	Increased map[string]*openapi3.MediaType `json:"increased,omitempty"`
	Missing   map[string]*openapi3.MediaType `json:"missing,omitempty"`
	Changed   map[string]*ChangedMediaType   `json:"changed,omitempty"`
	*DiffResult
}
type ChangedMediaType struct {
	//OldSchema *openapi3.Schema `json:"oldSchema,omitempty"`
	//NewSchema *openapi3.Schema `json:"newSchema,omitempty"`
	Schema *ChangedSchema `json:"schema,omitempty"`
	*DiffResult
}
type ChangedRequestBody struct {
	OldRequestBody *openapi3.RequestBody `json:"oldRequestBody,omitempty"`
	NewRequestBody *openapi3.RequestBody `json:"newRequestBody,omitempty"`
	Description    *ChangedMetadata      `json:"description,omitempty"`
	ChangeRequired bool                  `json:"changeRequired,omitempty"`
	Content        *ChangedContent       `json:"content,omitempty"`
	Extensions     *ChangedExtensions    `json:"extensions,omitempty"`
	*DiffResult
}
type ChangedAPIResponse struct {
	//OldApiResponses map[string]*openapi3.Response `json:"oldApiResponses,omitempty"`
	//NewApiResponses map[string]*openapi3.Response `json:"newApiResponses,omitempty"`
	Increased  map[string]*openapi3.Response `json:"increased,omitempty"`
	Missing    map[string]*openapi3.Response `json:"missing,omitempty"`
	Changed    map[string]*ChangedResponse   `json:"changed,omitempty"`
	Extensions *ChangedExtensions            `json:"extensions,omitempty"`
	*DiffResult
}
type ChangedSecurityRequirements struct {
	//OldSecurityRequirements []*openapi3.SecurityRequirement `json:"oldSecurityRequirements,omitempty"`
	//NewSecurityRequirements []*openapi3.SecurityRequirement `json:"newSecurityRequirements,omitempty"`
	Increased []*openapi3.SecurityRequirement `json:"increased,omitempty"`
	Missing   []*openapi3.SecurityRequirement `json:"missing,omitempty"`
	Changed   []*ChangedSecurityRequirement   `json:"changed,omitempty"`
	*DiffResult
}
type ChangedSecurityRequirement struct {
	OldSecurityRequirement *openapi3.SecurityRequirement `json:"oldSecurityRequirement,omitempty"`
	NewSecurityRequirement *openapi3.SecurityRequirement `json:"newSecurityRequirement,omitempty"`
	Increased              *openapi3.SecurityRequirement `json:"increased,omitempty"`
	Missing                *openapi3.SecurityRequirement `json:"missing,omitempty"`
	Changed                []*ChangedSecurityScheme      `json:"changed,omitempty"`
	*DiffResult
}
type ChangedSecurityScheme struct {
	//OldSecurityRequirement      *openapi3.SecurityScheme     `json:"oldSecurityScheme,omitempty"`
	//NewSecurityRequirement      *openapi3.SecurityScheme     `json:"newSecurityScheme,omitempty"`
	ChangedType                 bool                         `json:"changedType,omitempty"`
	ChangedIn                   bool                         `json:"changedIn,omitempty"`
	ChangedScheme               bool                         `json:"changedScheme,omitempty"`
	ChangedBearerFormat         bool                         `json:"changedBearerFormat,omitempty"`
	ChangedOpenIDConnectURL     bool                         `json:"changedOpenIdConnectUrl,omitempty"`
	ChangedSecuritySchemeScopes *ChangedSecuritySchemeScopes `json:"changedScopes,omitempty"`
	Description                 *ChangedMetadata             `json:"description,omitempty"`
	ChangedOAuthFlows           *ChangedOAuthFlows           `json:"oAuthFlows,omitempty"`
	Extensions                  *ChangedExtensions           `json:"extensions,omitempty"`
	*DiffResult
}
type ChangedOAuthFlows struct {
	//OldOAuthFlows              *openapi3.OAuthFlows `json:"oldOAuthFlows,omitempty"`
	//NewOAuthFlows              *openapi3.OAuthFlows `json:"newOAuthFlows,omitempty"`
	ImplicitOAuthFlow          *ChangedOAuthFlow  `json:"implicitOAuthFlow,omitempty"`
	PasswordOAuthFlow          *ChangedOAuthFlow  `json:"passwordOAuthFlow,omitempty"`
	ClientCredentialOAuthFlow  *ChangedOAuthFlow  `json:"clientCredentialOAuthFlow,omitempty"`
	AuthorizationCodeOAuthFlow *ChangedOAuthFlow  `json:"authorizationCodeOAuthFlow,omitempty"`
	Extensions                 *ChangedExtensions `json:"extensions,omitempty"`
}
type ChangedOAuthFlow struct {
	//OldOAuthFlow     *openapi3.OAuthFlow `json:"oldOAuthFlow,omitempty"`
	//NewOAuthFlow     *openapi3.OAuthFlow `json:"newOAuthFlow,omitempty"`
	AuthorizationURL bool               `json:"authorizationUrl,omitempty"`
	TokenURL         bool               `json:"tokenUrl,omitempty"`
	RefreshURL       bool               `json:"refreshUrl,omitempty"`
	Extensions       *ChangedExtensions `json:"extensions,omitempty"`
}
type ChangedResponse struct {
	OldAPIResponse *openapi3.Response `json:"oldApiResponse,omitempty"`
	NewAPIResponse *openapi3.Response `json:"newApiResponse,omitempty"`
	Description    *ChangedMetadata   `json:"description,omitempty"`
	Headers        *ChangedHeaders    `json:"headers,omitempty"`
	Content        *ChangedContent    `json:"content,omitempty"`
	Extensions     *ChangedExtensions `json:"extensions,omitempty"`
	*DiffResult
}
type ChangedHeaders struct {
	//OldHeaders map[string]*openapi3.Header `json:"oldHeaders,omitempty"`
	//NewHeaders map[string]*openapi3.Header `json:"newHeaders,omitempty"`

	Increased map[string]*openapi3.Header `json:"increased,omitempty"`
	Missing   map[string]*openapi3.Header `json:"missing,omitempty"`
	Changed   map[string]*ChangedHeader   `json:"changed,omitempty"`
	*DiffResult
}
type ChangedHeader struct {
	OldHeader   *openapi3.Header   `json:"oldHeader,omitempty"`
	NewHeader   *openapi3.Header   `json:"newHeader,omitempty"`
	Required    bool               `json:"required,omitempty"`
	Deprecated  bool               `json:"deprecated,omitempty"`
	Style       bool               `json:"style,omitempty"`
	Explode     bool               `json:"explode,omitempty"`
	Description *ChangedMetadata   `json:"description,omitempty"`
	Schema      *ChangedSchema     `json:"schema,omitempty"`
	Content     *ChangedContent    `json:"content,omitempty"`
	Extensions  *ChangedExtensions `json:"extensions,omitempty"`
	*DiffResult
}
type ChangedSchema struct {
	OldSchema                    *openapi3.Schema            `json:"oldSchema,omitempty"`
	NewSchema                    *openapi3.Schema            `json:"newSchema,omitempty"`
	Type                         string                      `json:"type,omitempty"`
	ChangedProperties            map[string]*ChangedSchema   `json:"changedProperties,omitempty"`
	IncreasedProperties          map[string]*openapi3.Schema `json:"increasedProperties,omitempty"`
	MissingProperties            map[string]*openapi3.Schema `json:"missingProperties,omitempty"`
	ChangeDeprecated             bool                        `json:"changeDeprecated,omitempty"`
	Description                  *ChangedMetadata            `json:"description,omitempty"`
	ChangeTitle                  bool                        `json:"changeTitle,omitempty"`
	Required                     *ChangedRequired            `json:"required,omitempty"`
	ChangeDefault                bool                        `json:"changeDefault,omitempty"`
	Enumeration                  *ChangedEnum                `json:"enumeration,omitempty"`
	ChangeFormat                 bool                        `json:"changeFormat,omitempty"`
	ReadOnly                     *ChangedReadOnly            `json:"readOnly,omitempty"`
	WriteOnly                    *ChangedWriteOnly           `json:"writeOnly,omitempty"`
	ChangedType                  bool                        `json:"changedType,omitempty"`
	MaxLength                    *ChangedMaxLength           `json:"maxLength,omitempty"`
	DiscriminatorPropertyChanged bool                        `json:"discriminatorPropertyChanged,omitempty"`
	Items                        *ChangedSchema              `json:"items,omitempty"`
	OneOfSchema                  *ChangedOneOfSchema         `json:"oneOfSchema,omitempty"`
	AddProp                      *ChangedSchema              `json:"addProp,omitempty"`
	Extensions                   *ChangedExtensions          `json:"extensions,omitempty"`
	*DiffResult
}

func (c *ChangedSchema) MarshalJSON() ([]byte, error) {
	if c != nil {
		if c.Required != nil && c.Required.DiffResult != nil && (!c.Required.Different && !c.Required.Incompatible) {
			c.Required = nil
		}
		if c.Enumeration != nil && c.Enumeration.DiffResult != nil && (!c.Enumeration.Different && !c.Enumeration.Incompatible) {
			c.Enumeration = nil
		}
		if c.ReadOnly != nil && c.ReadOnly.DiffResult != nil && (!c.ReadOnly.Different && !c.ReadOnly.Incompatible) {
			c.ReadOnly = nil
		}
		if c.WriteOnly != nil && c.WriteOnly.DiffResult != nil && (!c.WriteOnly.Different && !c.WriteOnly.Incompatible) {
			c.WriteOnly = nil
		}
		if c.MaxLength != nil && c.MaxLength.DiffResult != nil && (!c.MaxLength.Different && !c.MaxLength.Incompatible) {
			c.MaxLength = nil
		}
	}
	type changedSchemaType ChangedSchema
	return json.Marshal((*changedSchemaType)(c))
}

type ChangedOneOfSchema struct {
	//OldMapping map[string]string           `json:"oldMapping,omitempty"`
	//NewMapping map[string]string           `json:"newMapping,omitempty"`
	Increased map[string]*openapi3.Schema `json:"increased,omitempty"`
	Missing   map[string]*openapi3.Schema `json:"missing,omitempty"`
	Changed   map[string]*ChangedSchema   `json:"changed,omitempty"`
	*DiffResult
}
type ChangedMaxLength struct {
	OldValue *int `json:"oldValue,omitempty"`
	NewValue *int `json:"newValue,omitempty"`
	*DiffResult
}
type ChangedWriteOnly struct {
	*DiffResult
}
type ChangedReadOnly struct {
	*DiffResult
}
type ChangedSecuritySchemeScopes struct {
	*ChangedList
}
type ChangedEnum struct {
	*ChangedList
}
type ChangedRequired struct {
	*ChangedList
}
type ChangedList struct {
	OldValue  []string      `json:"oldValue,omitempty"`
	NewValue  []string      `json:"newValue,omitempty"`
	Increased []interface{} `json:"increased,omitempty"`
	Missing   []interface{} `json:"missing,omitempty"`
	Shared    []interface{} `json:"shared,omitempty"`
	*DiffResult
}
type ChangedExtensions struct {
	OldExtensions map[string]interface{} `json:"oldExtensions,omitempty"`
	NewExtensions map[string]interface{} `json:"newExtensions,omitempty"`
	Increased     map[string]interface{} `json:"increased,omitempty"`
	Missing       map[string]interface{} `json:"missing,omitempty"`
	Changed       map[string]interface{} `json:"changed,omitempty"`
	*DiffResult
}
type ChangedMetadata struct {
	Left  string `json:"left,omitempty"`
	Right string `json:"right,omitempty"`
	*DiffResult
}
type DiffResult struct {
	Incompatible bool `json:"incompatible"`
	Different    bool `json:"different"`
}

func NewChangedOpenAPIFromBytes(data []byte) (*ChangedOpenAPI, error) {
	var changedOpenAPI *ChangedOpenAPI
	err := json.Unmarshal(data, &changedOpenAPI)
	if err != nil {
		return nil, err
	}
	return changedOpenAPI, nil
}
