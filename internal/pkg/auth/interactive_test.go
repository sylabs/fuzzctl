// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package auth

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"testing"

	"golang.org/x/oauth2"
)

const goodClientID = "AGoodClientID"

var goodScopes = []string{"one", "two"}
var goodToken = oauth2.Token{
	AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ." +
		"SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	TokenType:    "Bearer",
	RefreshToken: "491GzL1WlDbMJ4-LmZokOes7lts-v89KMZ2k2f2SQFA",
}

type mockBrowserOpener struct {
	wantValues url.Values
	err        error
}

func (bo *mockBrowserOpener) Open(rawurl string) error {
	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}
	if got, want := u.Query(), bo.wantValues; !reflect.DeepEqual(got, want) {
		return fmt.Errorf("got query values %#v, want %#v", got, want)
	}
	return bo.err
}

func TestInteractiveSourceToken(t *testing.T) {
	wantValues := url.Values{
		"access_type":           []string{"offline"},
		"client_id":             []string{goodClientID},
		"code_challenge":        []string{"TWCwB8gObw2_ul7zTk_D4iDuRmRp8ODbC2MkqjIQyTM"},
		"code_challenge_method": []string{"S256"},
		"redirect_uri":          []string{"http://localhost:/authorization/callback"},
		"response_type":         []string{"code"},
		"scope":                 []string{strings.Join(goodScopes, " ")},
		"state":                 []string{"AZT9wvov_MBB0_8SBFtzyG5P-V_2YqXu6Cq99EotC3U"},
	}
	expired, cancel := context.WithCancel(context.Background())
	cancel()

	tests := []struct {
		name       string
		ctx        context.Context
		browserErr error
		wantValues url.Values
		wantErr    bool
		wantResult *result
	}{
		{"OK", context.Background(), nil, wantValues, false, &result{tok: &goodToken}},
		{"BrowserError", context.Background(), errors.New("oops"), wantValues, true, nil},
		{"ExpiredContext", expired, nil, wantValues, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := rand.NewSource(0)
			bo := mockBrowserOpener{
				wantValues: tt.wantValues,
				err:        tt.browserErr,
			}
			oc := &oauth2.Config{
				ClientID: goodClientID,
				Endpoint: oauth2.Endpoint{
					AuthURL: "http://test/authorization/callback",
				},
				RedirectURL: "http://localhost:/authorization/callback",
				Scopes:      goodScopes,
			}

			s := NewInteractiveTokenSource(tt.ctx, rs, &bo, oc)
			is := s.(*interactiveSource)
			is.testChan = make(chan result) // side channel to inject result during test
			defer close(is.testChan)

			// If a result is expected, spin off a goroutine to send that over the channel. This
			// simulates the result coming back from the callback.
			wg := sync.WaitGroup{}
			if tt.wantResult != nil {
				wg.Add(1)
				go func() {
					defer wg.Done()
					is.testChan <- *tt.wantResult
				}()
			}
			defer wg.Wait() // Ensure goroutine exits before testChan is closed

			// Call the token endpoint.
			tok, err := s.Token()
			if got := err; (got != nil) != tt.wantErr {
				t.Fatalf("got err %v, wantErr %v", got, tt.wantErr)
			}

			// Verify token result.
			if tt.wantResult != nil {
				if got, want := tok, tt.wantResult.tok; !reflect.DeepEqual(got, want) {
					t.Errorf("got result %v, want %v", got, want)
				}
			}
		})
	}
}
