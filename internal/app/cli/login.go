// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package cli

import (
	"context"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
	"github.com/sylabs/fuzzctl/internal/pkg/auth"
	"github.com/sylabs/fuzzctl/internal/pkg/browse"
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
		c := ar.GetAuthCodePKCEConfig()
		c.Scopes = []string{"openid", "offline_access"} // TODO: request additional scopes

		// Do interactive login.
		ts := auth.NewInteractiveTokenSource(context.TODO(), rand.NewSource(time.Now().UnixNano()), &browse.Browser{}, c)
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
