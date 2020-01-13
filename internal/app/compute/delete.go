package compute

import (
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <NAME>",
	Short: "delete allows you to remove a workflow from the compute service queue.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
