// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package config

import (
	"golang.org/x/oauth2"
)

// authConfig describes authentication and authorization configuration.
type authConfig struct {
	ClientID         string `yaml:"clientId,omitempty"`         // OAuth 2.0 Client ID
	AuthURL          string `yaml:"authUrl,omitempty"`          // OAuth 2.0 authorization endpoint
	TokenURL         string `yaml:"tokenUrl,omitempty"`         // OAuth 2.0 token endpoint
	LoginRedirectURL string `yaml:"loginRedirectUrl,omitempty"` // OAuth 2.0 login redirect endpoint
}

// GetOAuth2Config gets the OAuth 2.0 configuration for remote r.
func (r *remote) GetOAuth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID: r.AuthConfig.ClientID,
		Endpoint: oauth2.Endpoint{
			AuthURL:  r.AuthConfig.AuthURL,
			TokenURL: r.AuthConfig.TokenURL,
		},
		RedirectURL: r.AuthConfig.LoginRedirectURL,
	}
}
