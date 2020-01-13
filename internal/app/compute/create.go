package compute

import (
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <NAME> <WORKFLOW_CONFIG>",
	Short: "create enables you to submit a workflow to the compute service queue.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
