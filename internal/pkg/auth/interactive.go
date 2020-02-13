// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package auth

import (
	"context"
	"math/rand"
	"net/url"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

const (
	authPath = "/authorization/callback"
)

// BrowserOpener describes the interface to open an URL in a browser.
type BrowserOpener interface {
	Open(url string) error
}

type interactiveSource struct {
	ctx      context.Context
	rs       rand.Source
	bo       BrowserOpener
	oc       *oauth2.Config
	testChan chan result // used to inject results during testing only
}

// NewInteractiveTokenSource returns a token source that allows a user to interactively log in.
//
// This source uses the OAuth 2.0 Authorization Code with Proof Key Code Exchange flow.
func NewInteractiveTokenSource(ctx context.Context, rs rand.Source, bo BrowserOpener, oc *oauth2.Config) oauth2.TokenSource {
	return &interactiveSource{
		ctx:      ctx,
		rs:       rs,
		bo:       bo,
		oc:       oc,
		testChan: nil,
	}
}

// Token returns a token or an error. The token is obtained via interactive login using a browser
// window.
func (s *interactiveSource) Token() (*oauth2.Token, error) {
	resultChan := s.testChan
	if resultChan == nil {
		resultChan = make(chan result)
		defer close(resultChan)
	}
	state, err := newState(s.rs)
	if err != nil {
		return nil, err
	}
	cv, err := NewCodeVerifier(s.rs)
	if err != nil {
		return nil, err
	}
	sr := server{
		conf:   s.oc,
		state:  state,
		cv:     cv,
		result: resultChan,
	}

	u, err := url.Parse(s.oc.RedirectURL)
	if err != nil {
		return nil, err
	}

	// Start listening for incoming connection before we open the URL to avoid a race condition.
	hsr, err := sr.StartServer(u.Host)
	if err != nil {
		return nil, err
	}
	defer hsr.Close()

	// Open the URL to begin the OAuth2 flow.
	url := sr.conf.AuthCodeURL(
		state,
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam("code_challenge", cv.Challenge()),
		oauth2.SetAuthURLParam("code_challenge_method", cv.ChallengeMethod()),
	)
	if err := s.bo.Open(url); err != nil {
		return nil, err
	}

	// Wait until a result is received, or context cancellation.
	select {
	case r := <-resultChan:
		// Need to do clean shutdown here, to ensure HTTP response has been served.
		if err := hsr.Shutdown(s.ctx); err != nil {
			logrus.WithError(err).Print("shutdown failed")
		}
		return r.tok, r.err
	case <-s.ctx.Done():
		return nil, s.ctx.Err()
	}
}
