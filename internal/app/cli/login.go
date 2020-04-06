// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package cli

import (
	"math/rand"
	"time"

	"github.com/spf13/cobra"
	"github.com/sylabs/fuzzctl/internal/pkg/auth"
	"github.com/sylabs/fuzzctl/internal/pkg/browse"
	"github.com/sylabs/fuzzctl/internal/pkg/client"
	"github.com/sylabs/fuzzctl/internal/pkg/config"
	"golang.org/x/oauth2"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login allows you to authenticate with Fuzzball.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		ar, err := cfg.GetActiveRemote()
		if err != nil {
			return err
		}

		// Determine base URI.
		baseURI := httpAddr
		if baseURI == "" {
			baseURI = ar.GetBaseURI()
		}

		// OAuth 2.0 configuration from server.
		c, err := client.NewClient(baseURI+"/graphql",
			client.OptUserAgent(getUserAgent()))
		if err != nil {
			return err
		}
		oac, err := c.OAuth2AuthCodePKCEConfig(cmd.Context())
		if err != nil {
			return err
		}

		// Do interactive login.
		oc := oauth2.Config{
			ClientID: oac.ClientID,
			Endpoint: oauth2.Endpoint{
				AuthURL:  oac.AuthorizationEndpoint,
				TokenURL: oac.TokenEndpoint,
			},
			RedirectURL: oac.RedirectEndpoint,
			Scopes:      oac.Scopes,
		}
		ts := auth.NewInteractiveTokenSource(cmd.Context(), rand.NewSource(time.Now().UnixNano()), &browse.Browser{}, &oc)
		tok, err := ts.Token()
		if err != nil {
			return err
		}
		ar.SetAuthType(config.AuthConfigTypeAuthCodePKCE)

		// Update token source to specify the newly obtained token.
		tokenSrc = oauth2.StaticTokenSource(tok)
		return nil
	},
}
