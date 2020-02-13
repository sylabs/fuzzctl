// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package auth

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type clientCredentialsSource struct {
	ctx context.Context
	oc  *clientcredentials.Config
}

// NewClientCredentialsSource returns a token source that uses the supplied client credentials.
//
// This source uses the OAuth 2.0 Client Credentials flow.
func NewClientCredentialsSource(ctx context.Context, oc *clientcredentials.Config) oauth2.TokenSource {
	return &clientCredentialsSource{ctx: ctx, oc: oc}
}

// Token returns a token or an error. The token is obtained via the OAuth 2.0 Client Credentials
// flow.
func (s *clientCredentialsSource) Token() (*oauth2.Token, error) {
	return s.oc.Token(s.ctx)
}
