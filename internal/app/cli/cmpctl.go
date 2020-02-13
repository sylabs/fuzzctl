package cli

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sylabs/compute-cli/internal/pkg/compute"
	"golang.org/x/oauth2"
)

var (
	c *compute.Client

	debug bool

	httpAddr string
)

var CmpctlCmd = &cobra.Command{
	Use:   "cmpctl",
	Short: "cmpctl enables control of workflows for the compute service.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// change log level if debugging
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		// TODO: OAuth2 token source.
		ts := oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: "accesstoken",
		})

		// initialize global client for subcommands to leverage
		c = compute.NewClient(context.TODO(), ts, httpAddr)
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
