// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package config

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	// AuthConfigTypeAuthCodePKCE represents the OAuth 2.0 Authorization Code Grant with Proof Key
	// for Code Exchange (PKCE) flow.
	AuthConfigTypeAuthCodePKCE = "AuthCodePKCE"
	// AuthConfigTypeClientCredentials represents the OAuth 2.0 Client Credentials flow.
	AuthConfigTypeClientCredentials = "ClientCredentials"
)

// authConfig describes authentication and authorization configuration.
type authConfig struct {
	Type             string `yaml:"type"`                       // Auth type ("AuthCodePKCE" or "ClientCredentials")
	ClientID         string `yaml:"clientId,omitempty"`         // OAuth 2.0 client ID
	ClientSecret     string `yaml:"clientSecret,omitempty"`     // OAuth 2.0 client secret
	AuthURL          string `yaml:"authUrl,omitempty"`          // OAuth 2.0 authorization endpoint
	TokenURL         string `yaml:"tokenUrl,omitempty"`         // OAuth 2.0 token endpoint
	LoginRedirectURL string `yaml:"loginRedirectUrl,omitempty"` // OAuth 2.0 login redirect endpoint
}

// GetAuthType returns the auth type.
func (r *remote) GetAuthType() string {
	switch r.AuthConfig.Type {
	case AuthConfigTypeClientCredentials:
		return AuthConfigTypeClientCredentials
	case AuthConfigTypeAuthCodePKCE:
		fallthrough
	default:
		return AuthConfigTypeAuthCodePKCE
	}
}

// SetAuthType sets the auth type.
func (r *remote) SetAuthType(s string) error {
	switch s {
	case AuthConfigTypeClientCredentials:
		fallthrough
	case AuthConfigTypeAuthCodePKCE:
		r.AuthConfig.Type = s
	default:
		return fmt.Errorf("Unknown auth type: %v", s)
	}
	return nil
}

// GetAuthCodePKCEConfig gets the OAuth 2.0 configuration for remote r.
func (r *remote) GetAuthCodePKCEConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID: r.AuthConfig.ClientID,
		Endpoint: oauth2.Endpoint{
			AuthURL:  r.AuthConfig.AuthURL,
			TokenURL: r.AuthConfig.TokenURL,
		},
		RedirectURL: r.AuthConfig.LoginRedirectURL,
	}
}

// GetClientCredentialsConfig gets the OAuth 2.0 configuration for remote r.
func (r *remote) GetClientCredentialsConfig() *clientcredentials.Config {
	return &clientcredentials.Config{
		ClientID:     r.AuthConfig.ClientID,
		ClientSecret: r.AuthConfig.ClientSecret,
		TokenURL:     r.AuthConfig.TokenURL,
	}
}

func (r *remote) GetBaseURI() string {
	return r.BaseURI
}
