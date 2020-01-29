// Copyright (c) 2020, Sylabs, Inc. All rights reserved.
package cli

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <ID>",
	Short: "info allows you to see a workflow's state within the compute service.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		wf, err := c.Info(context.Background(), id)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Infof("Workflow: %s\n", wf)
	},
}
