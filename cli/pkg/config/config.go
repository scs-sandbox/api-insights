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

package config

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/clientcredentials"
)

// Environment Variables
const (
	Host               = "API_INSIGHTS_HOST"
	BasePath           = "API_INSIGHTS_BASE_PATH"
	AuthType           = "API_INSIGHTS_AUTH_TYPE"
	AuthUsername       = "API_INSIGHTS_USERNAME"
	AuthPassword       = "API_INSIGHTS_PASSWORD"
	AuthBearerToken    = "API_INSIGHTS_BEARER_TOKEN"
	OAuth2GrantType    = "API_INSIGHTS_OAUTH2_GRANT_TYPE"
	OAuth2TokenURL     = "API_INSIGHTS_OAUTH2_TOKEN_URL"
	OAuth2ClientID     = "API_INSIGHTS_OAUTH2_CLIENT_ID"
	OAuth2ClientSecret = "API_INSIGHTS_OAUTH2_CLIENT_SECRET"
)

const (
	authTypeBasic                    = "basic"
	authTypeBearerToken              = "bearer-token"
	authTypeOAuth2                   = "oauth2"
	oAuth2GrantTypeClientCredentials = "client_credentials"
)

const (
	hostDefault     = "https://host.example.com"
	basePathDefault = "/v1/apiregistry"
)

type Config struct {
	APIInsightsHost     string
	APIInsightsBasePath string
	Headers             map[string]string
	AuthConfig          *AuthConfig
}

type AuthConfig struct {
	Type               string // basic, bearer-token, oauth2
	Username           string // Type: basic
	Password           string // Type: basic
	BearerToken        string // Type: bearer-token
	OAuth2GrantType    string // Type: oauth2
	OAuth2TokenURL     string // Type: oauth2, OAuth2GrantType: client_credentials
	OAuth2ClientID     string // Type: oauth2, OAuth2GrantType: client_credentials
	OAuth2ClientSecret string // Type: oauth2, OAuth2GrantType: client_credentials
}

func (cfg AuthConfig) AuthSetter(ctx context.Context, rc *resty.Client) error {
	switch cfg.Type {
	case authTypeBasic:
		if cfg.Username == "" || cfg.Password == "" {
			return fmt.Errorf("config: invalid AuthConfig - missing required field(s) for type(%s): Username, Password", cfg.Type)
		}
		rc.SetBasicAuth(cfg.Username, cfg.Password)
	case authTypeBearerToken:
		if cfg.BearerToken == "" {
			return fmt.Errorf("config: invalid AuthConfig - missing required field(s) for type(%s): BearerToken", cfg.Type)
		}
		rc.SetAuthToken(cfg.BearerToken)
	case authTypeOAuth2:
		switch cfg.OAuth2GrantType {
		case oAuth2GrantTypeClientCredentials:
			if cfg.OAuth2TokenURL == "" || cfg.OAuth2ClientID == "" || cfg.OAuth2ClientSecret == "" {
				return fmt.Errorf("config: invalid AuthConfig - missing required field(s) for type(%s): OAuth2TokenURL, OAuth2ClientID, OAuth2ClientSecret", cfg.Type)
			}
			client := (&clientcredentials.Config{
				ClientID:     cfg.OAuth2ClientID,
				ClientSecret: cfg.OAuth2ClientSecret,
				TokenURL:     cfg.OAuth2TokenURL,
			}).Client(ctx)
			rc.SetTransport(client.Transport)
		default:
			return fmt.Errorf("config: invalid AuthConfig - invalid OAuth2GrantType(%s) for type(%s)", cfg.OAuth2GrantType, cfg.Type)
		}
	}
	return nil
}

// LoadConfig returns config
func LoadConfig() *Config {
	setConfigDefaults()
	c := &Config{
		APIInsightsHost:     viper.GetString(Host),
		APIInsightsBasePath: viper.GetString(BasePath),
		AuthConfig: &AuthConfig{
			Type:               viper.GetString(AuthType),
			Username:           viper.GetString(AuthUsername),
			Password:           viper.GetString(AuthPassword),
			BearerToken:        viper.GetString(AuthBearerToken),
			OAuth2GrantType:    viper.GetString(OAuth2GrantType),
			OAuth2TokenURL:     viper.GetString(OAuth2TokenURL),
			OAuth2ClientID:     viper.GetString(OAuth2ClientID),
			OAuth2ClientSecret: viper.GetString(OAuth2ClientSecret),
		},
	}

	return c
}

func setConfigDefaults() {
	viper.SetDefault(Host, hostDefault)
	viper.SetDefault(BasePath, basePathDefault)
}
