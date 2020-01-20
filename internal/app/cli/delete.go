// Copyright (c) 2020, Sylabs, Inc. All rights reserved.
package cli

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <ID>",
	Short: "delete allows you to remove a workflow from the compute service queue.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		wf, err := c.Delete(context.TODO(), id)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Infof("Deleted workflow: %s\n", wf)
	},
}
