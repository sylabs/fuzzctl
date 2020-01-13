package compute

import (
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <NAME>",
	Short: "list allows you to see your workflows state within the compute service.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
