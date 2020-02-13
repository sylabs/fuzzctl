// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package config

import "golang.org/x/oauth2"

// Remote defines the operations that can be performed on a remote config.
type Remote interface {
	GetOAuth2Config() *oauth2.Config

	GetToken() *oauth2.Token
	SetToken(t *oauth2.Token) error
}

// remote describes an endpoint.
type remote struct {
	BaseURI    string     `yaml:"baseUri"`              // Base URI of service
	AuthConfig authConfig `yaml:"authConfig,omitempty"` // Authentication and authorization configuration
	AuthToken  authToken  `yaml:"authToken,omitempty"`  // Authentication and authorization token
}
