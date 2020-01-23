package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sylabs/compute-cli/internal/pkg/compute"
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

		// initialize global client for subcommands to leverage
		c = compute.NewClient(httpAddr)
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
}
