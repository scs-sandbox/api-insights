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
	"crypto/tls"
	panopticaclient "github.com/cisco-developer/api-insights/api/pkg/panoptica/client"
	"github.com/go-openapi/runtime/client"
	"net/http"
)

const defaultCredentialScope = "global/services/portshift_request"

func New(host, accessKey, secretKey string) *panopticaclient.SecureApplicationAPI {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
	}

	cfg := panopticaclient.DefaultTransportConfig().WithHost(host)
	runtime := client.NewWithClient(cfg.Host, cfg.BasePath, cfg.Schemes, httpClient)
	runtime.DefaultAuthentication = NewAuth(accessKey, cfg.Host, cfg.BasePath, secretKey, defaultCredentialScope)

	return panopticaclient.New(runtime, nil)
}
