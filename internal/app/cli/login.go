// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package cli

import (
	"context"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
	"github.com/sylabs/compute-cli/internal/pkg/auth"
	"github.com/sylabs/compute-cli/internal/pkg/browse"
	"golang.org/x/oauth2"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login allows you to authenticate with the compute service.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		ar, err := cfg.GetActiveRemote()
		if err != nil {
			return err
		}
		c := ar.GetOAuth2Config()
		c.Scopes = []string{"offline_access"} // TODO: request additional scopes

		// Do interactive login.
		ts := auth.NewInteractiveTokenSource(context.TODO(), rand.NewSource(time.Now().UnixNano()), &browse.Browser{}, c)
		tok, err := ts.Token()
		if err != nil {
			return err
		}

		// Update token source to specify the newly obtained token.
		tokenSrc = oauth2.StaticTokenSource(tok)
		return nil
	},
}
