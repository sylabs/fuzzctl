// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package auth

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type result struct {
	err error
	tok *oauth2.Token
}

type server struct {
	conf   *oauth2.Config
	state  string
	cv     *CodeVerifier
	result chan result
}

var (
	// ErrStateMismatch is returned with invalid state is received.
	ErrStateMismatch = errors.New("auth: state mismatch")
)

// error replies to the request with the specified error and status code, and sends the error
// result over the result channel.
//
// If err is nil, the text associated with the status code is written to w.
func (s server) error(w http.ResponseWriter, err error, code int) {
	s.result <- result{err: err}
	if err != nil {
		http.Error(w, err.Error(), code)
	} else {
		http.Error(w, http.StatusText(code), code)
	}
}

// ServeHTTP handles the authorization callback.
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Ignore anything that isn't a GET to the auth endpoint.
	if r.URL.Path != authPath {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Verify the state parameter.
	v, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		s.error(w, err, http.StatusInternalServerError)
		return
	}
	if v.Get("state") != s.state {
		s.error(w, ErrStateMismatch, http.StatusForbidden)
		return
	}

	// Check whether the request was successful.
	if errStr := v.Get("error"); errStr != "" {
		if d := v.Get("error_description"); d != "" {
			errStr += fmt.Sprintf(": %s", d)
		}
		s.error(w, errors.New(errStr), http.StatusBadRequest)
		return
	}

	// Exchange will do the handshake to retrieve the initial access token.
	tok, err := s.conf.Exchange(r.Context(), v.Get("code"), oauth2.SetAuthURLParam("code_verifier", s.cv.String()))
	if err != nil {
		s.error(w, err, http.StatusUnauthorized)
		return
	}

	// Pass the token!
	s.result <- result{tok: tok}

	// Show succes page!
	fmt.Fprintf(w, `
		<p><strong>Success!</strong></p>
		<p>You are authenticated and can now return to the CLI.</p>
	`)
}

// StartServer starts an HTTP server. The server is guaranteed to be listening when the function
// returns.
func (s server) StartServer(address string) (*http.Server, error) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	hsr := http.Server{Handler: s}
	go func() {
		if err := hsr.Serve(ln); err != http.ErrServerClosed {
			logrus.WithError(err).Printf("serve failed")
		}
	}()
	return &hsr, nil
}
