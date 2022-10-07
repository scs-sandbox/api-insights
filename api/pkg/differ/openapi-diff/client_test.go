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

package openapidiff

import (
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/internal/models/diff"
	"github.com/cisco-developer/api-insights/api/pkg/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name      string
		want      Differ
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient()
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_cliClient_DiffDocuments(t *testing.T) {
	jarPath := "/tmp/openapi-diff-cli-2.1.0-beta.3-all.jar"
	if p := os.Getenv("OPENAPI_DIFF_JAR_FILE"); p != "" {
		jarPath = p
	}
	javaOpts := "-Xms512m -Xmx1024m"
	if p := os.Getenv("OPENAPI_DIFF_JAVA_OPTS"); p != "" {
		javaOpts = p
	}

	opts := &Options{
		OpenAPIDiffJarFile:  utils.StringPtr(jarPath),
		OpenAPIDiffJavaOpts: utils.StringPtr(javaOpts),
		Format:              utils.StringPtr("json"),
	}
	client := &cliClient{}

	type args struct {
		oldDoc models.SpecDoc
		newDoc models.SpecDoc
		cfg    *diff.Config
		opts   *Options
	}
	tests := []struct {
		name      string
		c         *cliClient
		args      args
		want      *diff.Result
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "test against meraki spec",
			c:    client,
			args: args{
				oldDoc: models.SpecDoc(loadSpec("testdata/meraki/spec-1.17.0.json")),
				newDoc: models.SpecDoc(loadSpec("testdata/meraki/spec-1.18.0.json")),
				cfg: &diff.Config{
					OutputFormat: "json",
				},
				opts: opts,
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's New\n\n##### `GET` /networks/{networkId}/webhooks/payloadTemplates\n\n> List the webhook payload templates for a network\n\n##### `POST` /networks/{networkId}/webhooks/payloadTemplates\n\n> Create a webhook payload template for a network\n\n##### `GET` /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId}\n\n> Get the webhook payload template for a network\n\n##### `PUT` /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId}\n\n> Update a webhook payload template for a network\n\n##### `DELETE` /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId}\n\n> Destroy a webhook payload template for a network\n\n#### What's Modified\n\n##### `GET` /networks/{networkId}/appliance/uplinks/usageHistory\n\n> Get the sent and received bytes for each uplink of a network.\n\n###### Parameters:\n\nModified: `t1` in `query`\n> The end of the timespan for the data. t1 can be a maximum of 31 days after t0.\n\nModified: `timespan` in `query`\n> The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 10 minutes.\n\n\n##### `POST` /networks/{networkId}/webhooks/httpServers\n\n> Add an HTTP server to a network\n\n###### Request:\n\nModified content type: `application/json`\n\n* Modified property `payloadTemplate` (object)\n    > The payload template to use when posting data to the HTTP server.\n\n    * Added property `name` (string)\n        > The name of the payload template.\n\n\n##### `POST` /networks/{networkId}/webhooks/webhookTests\n\n> Send a test webhook for a network\n\n###### Request:\n\nModified content type: `application/json`\n\n* Added property `payloadTemplateName` (string)\n    > The name of the payload template.\n\n\n##### `PUT` /networks/{networkId}/wireless/ssids/{number}/vpn\n\n> Update the VPN settings for the SSID\n\n###### Request:\n\nModified content type: `application/json`\n\n* Added property `failover` (object)\n    > Secondary VPN concentrator settings. This is only used when two VPN concentrators are configured on the SSID.\n\n\n##### `GET` /organizations\n\n> List the organizations that the user has privileges on\n\n###### Response:\n\nModified response: **200 OK**\n> Successful operation\n\n* Modified content type: `application/json`\n    * Modified items (array):\n\n        * Added property `licensing` (object)\n            > Licensing related settings\n\n        * Modified property `id` (string)\n            > Organization ID\n\n        * Modified property `name` (string)\n            > Organization name\n\n        * Modified property `url` (string)\n            > Organization URL\n\n        * Modified property `api` (object)\n            > API related settings\n\n\n##### `GET` /organizations/{organizationId}\n\n> Return an organization\n\n###### Response:\n\nModified response: **200 OK**\n> Successful operation\n\n* Modified content type: `application/json`\n    * Added property `licensing` (object)\n        > Licensing related settings\n\n    * Modified property `id` (string)\n        > Organization ID\n\n    * Modified property `name` (string)\n        > Organization name\n\n    * Modified property `url` (string)\n        > Organization URL\n\n    * Modified property `api` (object)\n        > API related settings\n\n        * Modified property `enabled` (boolean)\n            > Enable API access\n\n\n##### `GET` /organizations/{organizationId}/devices/statuses\n\n> List the status of every Meraki device in the organization\n\n###### Parameters:\n\nDeleted: `components` in `query`\n> components\n\n\n##### `GET` /organizations/{organizationId}/summary/top/clients/byUsage\n\n> Return metrics for organization's top 10 clients by data usage (in mb) over given time range.\n\n###### Response:\n\nModified response: **200 OK**\n> Successful operation\n\n* Deleted content type: `application/json`\n\n##### `GET` /organizations/{organizationId}/webhooks/alertTypes\n\n> Return a list of alert types to be used with managing webhook alerts\n\n###### Parameters:\n\nAdded: `productType` in `query`\n> Filter sample alerts to a specific product type\n\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "test against petstore spec",
			c:    client,
			args: args{
				oldDoc: models.SpecDoc(loadSpec("testdata/petstore/petstore_1.json")),
				newDoc: models.SpecDoc(loadSpec("testdata/petstore/petstore_2.json")),
				cfg: &diff.Config{
					OutputFormat: "json",
				},
				opts: opts,
			},
			want: &diff.Result{
				JSON: &diff.JSONResult{
					Message:  "#### What's New\n\n##### `GET` /pet/{petId}\n\n> Find pet by ID\n\n#### What's Deleted\n\n##### `POST` /pet/{petId}\n\n> Updates a pet in the store with form data\n\n#### What's Deprecated\n\n##### `GET` /user/logout\n\n> Logs out current logged in user session\n\n#### What's Modified\n\n##### `DELETE` /pet/{petId}\n\n> Deletes a pet\n\n###### Parameters:\n\nAdded: `newHeaderParam` in `header`\n\n\n##### `GET` /user/login\n\n> Logs user into the system\n\n###### Parameters:\n\nDeleted: `password` in `query`\n> The password for login in clear text\n\n###### Response:\n\nModified response: **200 OK**\n> successful operation\n\n* Added header: `X-Rate-Limit-New`\n\n* Deleted header: `X-Rate-Limit`\n\n* Modified header: `X-Expires-After`\n\n##### `GET` /user/logout\n\n> Logs out current logged in user session\n\n\n##### `POST` /pet/{petId}/uploadImage\n\n> uploads an image for pet\n\n###### Parameters:\n\nModified: `petId` in `path`\n> ID of pet to update, default false\n\n\n##### `POST` /user\n\n> Create user\n\n###### Request:\n\nModified content type: `application/json`\n\n* Added property `newUserFeild` (integer)\n    > a new user feild demo\n\n* Deleted property `phone` (string)\n\n\n##### `POST` /user/createWithArray\n\n> Creates list of users with given input array\n\n###### Request:\n\nModified content type: `application/json`\n\n* Modified items (array):\n\n    * Added property `newUserFeild` (integer)\n        > a new user feild demo\n\n    * Deleted property `phone` (string)\n\n\n##### `POST` /user/createWithList\n\n> Creates list of users with given input array\n\n###### Request:\n\nModified content type: `application/json`\n\n* Modified items (array):\n\n    * Added property `newUserFeild` (integer)\n        > a new user feild demo\n\n    * Deleted property `phone` (string)\n\n\n##### `GET` /user/{username}\n\n> Get user by user name\n\n###### Response:\n\nModified response: **200 OK**\n> successful operation\n\n* Modified content type: `application/json`\n    * Added property `newUserFeild` (integer)\n        > a new user feild demo\n\n    * Deleted property `phone` (string)\n\n\n* Modified content type: `application/xml`\n    * Added property `newUserFeild` (integer)\n        > a new user feild demo\n\n    * Deleted property `phone` (string)\n\n\n##### `PUT` /user/{username}\n\n> Updated user\n\n###### Request:\n\nModified content type: `application/json`\n\n* Added property `newUserFeild` (integer)\n    > a new user feild demo\n\n* Deleted property `phone` (string)\n\n\n##### `PUT` /pet\n\n> Update an existing pet\n\n###### Request:\n\nDeleted content type: `application/xml`\n\nModified content type: `application/json`\n\n* Added property `newField` (string)\n    > a field demo\n\n* Modified property `category` (object)\n\n    * Added property `newCatFeild` (string)\n\n    * Deleted property `name` (string)\n\n\n##### `POST` /pet\n\n> Add a new pet to the store\n\n###### Parameters:\n\nAdded: `tags` in `query`\n> add new query param demo\n\n###### Request:\n\nModified content type: `application/xml`\n\n* Added property `newField` (string)\n    > a field demo\n\n* Modified property `category` (object)\n\n    * Added property `newCatFeild` (string)\n\n    * Deleted property `name` (string)\n\nModified content type: `application/json`\n\n* Added property `newField` (string)\n    > a field demo\n\n* Modified property `category` (object)\n\n    * Added property `newCatFeild` (string)\n\n    * Deleted property `name` (string)\n\n\n##### `GET` /pet/findByStatus\n\n> Finds Pets by status\n\n###### Parameters:\n\nModified: `status` in `query`\n> Status values that need to be considered for filter\n\n###### Response:\n\nModified response: **200 OK**\n> successful operation\n\n* Modified content type: `application/xml`\n    * Modified items (array):\n\n        * Added property `newField` (string)\n            > a field demo\n\n        * Modified property `category` (object)\n\n\n* Modified content type: `application/json`\n    * Modified items (array):\n\n        * Added property `newField` (string)\n            > a field demo\n\n        * Modified property `category` (object)\n\n\n##### `GET` /pet/findByTags\n\n> Finds Pets by tags\n\n###### Response:\n\nModified response: **200 OK**\n> successful operation\n\n* Modified content type: `application/xml`\n    * Modified items (array):\n\n        * Added property `newField` (string)\n            > a field demo\n\n        * Modified property `category` (object)\n\n\n* Modified content type: `application/json`\n    * Modified items (array):\n\n        * Added property `newField` (string)\n            > a field demo\n\n        * Modified property `category` (object)\n\n\n",
					Breaking: true,
				},
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cliClient{}
			got, err := c.DiffDocuments(tt.args.oldDoc, tt.args.newDoc, tt.args.cfg, tt.args.opts)
			tt.assertion(t, err)
			// The result is not the same because of element in array is not in same order every time. So here use length to compare.
			assert.Equal(t, len(tt.want.JSON.Message), len(got.JSON.Message))
			assert.Equal(t, tt.want.JSON.Breaking, got.JSON.Breaking)
		})
	}
}

func loadSpec(file string) *string {
	content, _ := os.ReadFile(file)
	s := string(content)
	return &s
}
