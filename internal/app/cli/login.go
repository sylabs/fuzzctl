// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
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
	Run: func(cmd *cobra.Command, args []string) {
		ts := auth.NewInteractiveTokenSource(
			context.TODO(),
			rand.NewSource(time.Now().UnixNano()),
			&browse.Browser{},
			"0oa24wwhwBWYa1T804x6", // TODO: discover from .well-known?
			oauth2.Endpoint{
				AuthURL:  "https://dev-930666.okta.com/oauth2/default/v1/authorize", // TODO: discover from .well-known
				TokenURL: "https://dev-930666.okta.com/oauth2/default/v1/token",     // TODO: discover from .well-known
			},
			"offline_access", // TODO: request additional scopes
		)
		tok, err := ts.Token()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Login failed: %v", err)
		}

		// TODO: save token instead of displaying it!
		if b, err := json.MarshalIndent(tok, "", "  "); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", err)
		} else {
			fmt.Printf("token: %s\n", string(b))
		}
	},
}
