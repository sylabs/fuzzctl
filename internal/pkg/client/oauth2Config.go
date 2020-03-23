// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package client

import (
	"context"
	"errors"
)

// AuthCodePKCEConfig contains OAuth 2.0 configuration for Authorization Code Flow with Proof Key
// for Code Exchange.
type AuthCodePKCEConfig struct {
	// The client identifier to use.
	ClientID string

	// The URL of the authorization server's authorization endpoint.
	AuthorizationEndpoint string

	// The URL of the authorization server's token endpoint.
	TokenEndpoint string

	// The URL of the redirect endpoint.
	RedirectEndpoint string

	// Recommended scope(s) to request.
	Scopes []string
}

// OAuth2AuthCodePKCEConfig retrieves OAuth 2.0 configuration for Authorization Code Flow with
// Proof Key for Code Exchange (if supported).
func (c *Client) OAuth2AuthCodePKCEConfig(ctx context.Context) (AuthCodePKCEConfig, error) {
	q := struct {
		OAuth2Config struct {
			AuthCodePKCE *AuthCodePKCEConfig
		} `graphql:"oauth2Config"`
	}{}

	if err := c.Query(ctx, &q, nil); err != nil {
		return AuthCodePKCEConfig{}, err
	}
	if q.OAuth2Config.AuthCodePKCE == nil {
		return AuthCodePKCEConfig{}, errors.New("interactive login is not supported by the server")
	}
	return *q.OAuth2Config.AuthCodePKCE, nil
}
