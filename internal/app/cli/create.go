// Copyright (c) 2020, Sylabs, Inc. All rights reserved.
package cli

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <NAME>",
	Short: "create enables you to submit a workflow to the compute service queue.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		logrus.Debugf("Creating workflow: %s", name)
		wf, err := c.Create(context.TODO(), name)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Infof("Created workflow: %s\n", wf)
	},
}
