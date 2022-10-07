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

package panoptica

import (
	"fmt"
	"net/http"

	"github.com/EscherAuth/escher/config"
	escher_request "github.com/EscherAuth/escher/request"
	"github.com/EscherAuth/escher/signer"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type HmacSha2Auth struct {
	Host            string
	BasePath        string
	Credentials     *credentials.Credentials
	CredentialScope string
}

func NewAuth(agentID, host, basePath, sharedKey, credentialScope string) HmacSha2Auth {
	staticCreds := credentials.NewStaticCredentials(agentID, sharedKey, "")
	return HmacSha2Auth{
		Credentials:     staticCreds,
		CredentialScope: credentialScope,
		Host:            host,
		BasePath:        basePath,
	}
}

// copied from https://github.com/EscherAuth/escher/blob/master/request/httprequest.go#L74
func (auth HmacSha2Auth) createEscherHeadersFromHTTPHeaders(h http.Header) escher_request.Headers {
	headers := escher_request.Headers{}

	for key, values := range h {
		for _, value := range values {
			headers = append(headers, [2]string{key, value})
		}
	}

	if h.Get("host") == "" {
		// https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.23
		headers = append(headers, [2]string{"host", auth.Host})
	}

	return headers
}

func (auth HmacSha2Auth) AuthenticateRequest(r runtime.ClientRequest, _ strfmt.Registry) error {
	c := config.Config{}
	config.SetDefaults(&c)
	c.CredentialScope = auth.CredentialScope

	value, err := auth.Credentials.Get()
	if err != nil {
		return err
	}

	c.AccessKeyId = value.AccessKeyID
	c.ApiSecret = value.SecretAccessKey

	signerObj := signer.New(c)
	escherReq := escher_request.New(r.GetMethod(), auth.buildURLString(r), auth.createEscherHeadersFromHTTPHeaders(r.GetHeaderParams()), string(r.GetBody()), 36000)

	signReq, err := signerObj.SignRequest(escherReq, []string{})
	if err != nil {
		return err
	}

	err = setHeader(r, signReq, c.GetAuthHeaderName())
	if err != nil {
		return err
	}

	err = setHeader(r, signReq, c.GetDateHeaderName())
	if err != nil {
		return err
	}

	return nil
}

func (auth HmacSha2Auth) buildURLString(r runtime.ClientRequest) string {
	urlString := ""
	if auth.BasePath != "" {
		urlString += auth.BasePath
	}
	urlString += r.GetPath()
	if queryParams := r.GetQueryParams().Encode(); queryParams != "" {
		urlString += "?" + queryParams
	}
	return urlString
}

func setHeader(r runtime.ClientRequest, signReq *escher_request.Request, header string) error {
	authValue, success := signReq.Headers().Get(header)
	if !success {
		return fmt.Errorf("could not get auth header")
	}

	err := r.SetHeaderParam(header, authValue)
	if err != nil {
		return fmt.Errorf("failed to set header: %v", err)
	}

	return nil
}
