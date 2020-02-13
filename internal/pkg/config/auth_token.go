// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package config

import (
	"time"

	"golang.org/x/oauth2"
)

// authToken describes authentication and authorization tokens.
type authToken struct {
	AccessToken  string    `yaml:"accessToken,omitempty"`  // OAuth 2.0 access token
	TokenType    string    `yaml:"tokenType,omitempty"`    // OAuth 2.0 token type
	RefreshToken string    `yaml:"refreshToken,omitempty"` // OAuth 2.0 refresh token
	Expiry       time.Time `yaml:"expiry,omitempty"`       // OAuth 2.0 access token expiry
}

// SetToken sets the token for remote r to t. If t is nil, the token for the remote is cleared.
func (r *remote) SetToken(t *oauth2.Token) error {
	if t != nil {
		r.AuthToken.AccessToken = t.AccessToken
		r.AuthToken.TokenType = t.TokenType
		r.AuthToken.RefreshToken = t.RefreshToken
		r.AuthToken.Expiry = t.Expiry
	} else {
		r.AuthToken.AccessToken = ""
		r.AuthToken.TokenType = ""
		r.AuthToken.RefreshToken = ""
		r.AuthToken.Expiry = time.Time{}
	}
	return nil
}

// GetToken gets the token for remote r.
func (r *remote) GetToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  r.AuthToken.AccessToken,
		TokenType:    r.AuthToken.TokenType,
		RefreshToken: r.AuthToken.RefreshToken,
		Expiry:       r.AuthToken.Expiry,
	}
}
