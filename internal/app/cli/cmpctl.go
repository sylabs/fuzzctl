package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sylabs/compute-cli/internal/pkg/compute"
	"github.com/sylabs/compute-cli/internal/pkg/config"
	"golang.org/x/oauth2"
)

var (
	c *compute.Client

	debug bool

	httpAddr string

	tokenSrc oauth2.TokenSource
	cfg      *config.Config
)

var CmpctlCmd = &cobra.Command{
	Use:   "cmpctl",
	Short: "cmpctl enables control of workflows for the compute service.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// change log level if debugging
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		// Read configuration.
		f, err := os.Open(".fuzzconf.yaml")
		if err == nil {
			// Pre-existing config.
			defer f.Close()
			c, err := config.Read(f)
			if err != nil {
				return fmt.Errorf("Failed to read config file: %w", err)
			}
			cfg = c
		} else {
			// Write default config.
			c, err := config.Default()
			if err != nil {
				return fmt.Errorf("Failed to create default config: %w", err)
			}
			cfg = c
		}

		ctx := context.TODO()

		// Configure OAuth2 Token Source
		r, err := cfg.GetActiveRemote()
		if err != nil {
			return fmt.Errorf("Failed to get active remote: %w", err)
		}
		tokenSrc = r.GetOAuth2Config().TokenSource(ctx, r.GetToken())

		// initialize global client for subcommands to leverage
		c = compute.NewClient(ctx, tokenSrc, httpAddr)

		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		ar, err := cfg.GetActiveRemote()
		if err != nil {
			return err
		}

		if tokenSrc != nil {
			// Get updated token.
			t, err := tokenSrc.Token()
			if err != nil {
				return err
			}
			ar.SetToken(t)
		} else {
			ar.SetToken(nil)
		}

		// Save config.
		f, err := os.OpenFile(".fuzzconf.yaml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
		defer f.Close()
		return cfg.Write(f)
	},
}

func init() {
	CmpctlCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug output")
	CmpctlCmd.PersistentFlags().StringVar(&httpAddr, "http_addr", "http://localhost:8080", "Address to reach compute server")

	CmpctlCmd.AddCommand(createCmd)
	CmpctlCmd.AddCommand(deleteCmd)
	CmpctlCmd.AddCommand(infoCmd)
	CmpctlCmd.AddCommand(listCmd)
	CmpctlCmd.AddCommand(serverInfoCmd)
	CmpctlCmd.AddCommand(loginCmd)
}
