// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"golang.org/x/oauth2"
)

const goodState = "good"

type mockOAuthServer struct {
	wantGrantType   string
	wantCode        string
	wantRedirectURI string

	code  int
	token oauth2.Token
}

func (s *mockOAuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if got := v.Get("grant_type"); got != s.wantGrantType {
		http.Error(w, fmt.Sprintf("got grantType %v, want %v", got, s.wantGrantType), http.StatusBadRequest)
		return
	}
	if got := v.Get("code"); got != s.wantCode {
		http.Error(w, fmt.Sprintf("got code %v, want %v", got, s.wantCode), http.StatusBadRequest)
		return
	}
	if got := v.Get("redirect_uri"); got != s.wantRedirectURI {
		http.Error(w, fmt.Sprintf("got redirectURI %v, want %v", got, s.wantRedirectURI), http.StatusBadRequest)
		return
	}
	if s.code != http.StatusOK {
		w.WriteHeader(s.code)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s.token)
	}
}

func TestServeHTTP(t *testing.T) {
	goodValues := url.Values{
		"state": []string{goodState},
	}

	tests := []struct {
		name       string
		path       string
		method     string
		values     url.Values
		code       int
		wantResult bool // If true, result expected over result chan
		wantErr    bool // If true, result expected to contain non-nil error
		wantCode   int  // Expected HTTP status code
	}{
		{
			"BadPath",
			"/bad",
			http.MethodGet,
			goodValues,
			http.StatusOK,
			false,
			false,
			http.StatusNotFound,
		},
		{
			"BadMethod",
			authPath,
			"bad",
			goodValues,
			http.StatusOK,
			false,
			false,
			http.StatusMethodNotAllowed,
		},
		{
			"BadState",
			authPath,
			http.MethodGet,
			url.Values{
				"state": []string{"bad"},
			},
			http.StatusOK,
			true,
			true,
			http.StatusForbidden,
		},
		{
			"Error",
			authPath,
			http.MethodGet,
			url.Values{
				"state":             []string{goodState},
				"error":             []string{"An error"},
				"error_description": []string{"An error description"},
			},
			http.StatusOK,
			true,
			true,
			http.StatusBadRequest,
		},
		{
			"OK",
			authPath,
			http.MethodGet,
			goodValues,
			http.StatusOK,
			true,
			false,
			http.StatusOK,
		},
		{
			"Unauthorized",
			authPath,
			http.MethodGet,
			goodValues,
			http.StatusUnauthorized,
			true,
			true,
			http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := mockOAuthServer{
				token: goodToken,
				code:  tt.code,
			}
			authServer := httptest.NewServer(&ms)
			defer authServer.Close()

			s := server{
				conf: &oauth2.Config{
					ClientID: "12345678",
					Endpoint: oauth2.Endpoint{
						AuthURL:  authServer.URL,
						TokenURL: authServer.URL,
					},
				},
				state:  goodState,
				cv:     &CodeVerifier{},
				result: make(chan result),
			}

			// Build up incoming request according to test parameters.
			url := url.URL{
				Path:     tt.path,
				RawQuery: tt.values.Encode(),
			}
			r := httptest.NewRequest(tt.method, url.String(), nil)
			rr := httptest.NewRecorder()

			// Call handler in goroutine. Results are bifricated over HTTP and via a channel.
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.ServeHTTP(rr, r)
			}()

			// Wait until result is received on channel, or channel is closed.
			if tt.wantResult {
				res := <-s.result
				if got := res.err; (got != nil) != tt.wantErr {
					t.Errorf("got err %v, wantErr %v", got, tt.wantErr)
				}
			}

			// Wait until handler goroutine has returned before validating HTTP response.
			wg.Wait()
			if got := rr.Code; got != tt.wantCode {
				t.Errorf("got code %v, want %v", got, tt.wantCode)
			}
		})
	}
}
