// Copyright (c) 2020, Sylabs, Inc. All rights reserved.
package cli

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverInfoCmd = &cobra.Command{
	Use:   "serverinfo",
	Short: "serverinfo enables you to view the properties of a compute service.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		si, err := c.ServerInfo(context.Background())
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Infof("Server Info: %s\n", si)
	},
}
