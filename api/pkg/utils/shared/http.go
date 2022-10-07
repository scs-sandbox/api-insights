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
	"encoding/json"
	"fmt"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
	"github.com/urfave/cli/v2"
	"net/http"
)

func ServeJSON(rw http.ResponseWriter, httpStatus int, data interface{}) error {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpStatus)
	body, err := json.Marshal(data)
	if err == nil {
		_, err = rw.Write(body)
	}
	return err
}

func HTTPApp(config *AppConfig, customFlags []cli.Flag, commands ...*cli.Command) *cli.App {
	for _, cmd := range commands {
		wrapCommand(config, cmd)
	}

	app := &cli.App{
		Version:  config.AppVersion,
		Usage:    fmt.Sprintf("Run the %s microservice tool", config.AppName),
		Commands: serveCommands(config, customFlags, commands...),
	}
	return app
}

func MergeFlags(flagGroups ...[]cli.Flag) []cli.Flag {
	// TODO: doesn't actually merge, just appends for now
	all := []cli.Flag{}
	for _, group := range flagGroups {
		all = append(all, group...)
	}
	return all
}

func serveCommands(config *AppConfig, customFlags []cli.Flag, commands ...*cli.Command) []*cli.Command {
	flags := append(serveFlags(config), customFlags...)
	return append(commands, &cli.Command{
		Name:   "serve",
		Action: httpServerAction(config),
		Before: flagJSON(flags),
		Flags:  flags,
	})
}

func httpServerAction(config *AppConfig) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		return startHTTPListener(c.Int("port"), config)
	}
}

func startHTTPListener(port int, config *AppConfig) error {
	handler, err := config.HTTPHandler(config)
	if err != nil {
		return err
	}
	swaggerConfig := restfulspec.Config{
		WebServices:                   handler.RegisteredWebServices(), // you control what services are visible
		WebServicesURL:                fmt.Sprintf("http://%s:%d", config.AppHost, config.AppPort),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: buildSwaggerInfo(config),
	}

	handler.Add(restfulspec.NewOpenAPIService(swaggerConfig))

	if CorsHandlerEnabled {
		addCorsHandler(handler)
	}

	serverPort := fmt.Sprintf(":%d", port)
	LogInfof("starting %s (%s) webserver on %s", config.AppName, config.AppVersion, serverPort)

	return http.ListenAndServe(serverPort, handler)
}

func addCorsHandler(handler *restful.Container) {
	corsList := GetCORSWhitelist()

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-Set-Authorization"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization", "Origin", "Token", "csrf", "timestamp", "x-xsrf-token"},
		AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"},
		AllowedDomains: corsList,
		CookiesAllowed: true,
		Container:      handler}
	handler.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	handler.Filter(handler.OPTIONSFilter)
}

func buildSwaggerInfo(config *AppConfig) func(*spec.Swagger) {
	return func(s *spec.Swagger) {
		s.Info = &spec.Info{
			InfoProps: spec.InfoProps{
				Contact: &spec.ContactInfo{
					ContactInfoProps: spec.ContactInfoProps{
						Name:  "Ask API Insights",
						Email: "ask-api-insights@cisco.com",
					},
				},
				License: &spec.License{
					LicenseProps: spec.LicenseProps{
						Name: "Apache 2.0",
						URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
					},
				},
				Title:       "API Insights",
				Description: "API Insights is an open-source tool that helps developers improve API quality and security.",
			},
		}
		s.Host = fmt.Sprintf("%s:%d", config.AppHost, config.AppPort)
		s.ExternalDocs = &spec.ExternalDocumentation{
			Description: "Visit API Insights",
			URL:         "https://developer.cisco.com/site/api-insights/",
		}
		s.Consumes = []string{restful.MIME_JSON}
		s.Produces = []string{restful.MIME_JSON}
		s.BasePath = "/v1"

		jwtSecurityScheme := spec.APIKeyAuth("Authorization", "header")
		jwtSecurityScheme.Description = "jwt token"
		s.SecurityDefinitions = map[string]*spec.SecurityScheme{
			"jwt": jwtSecurityScheme,
		}
	}
}
