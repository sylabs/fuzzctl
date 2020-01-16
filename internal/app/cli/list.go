// Copyright (c) 2020, Sylabs, Inc. All rights reserved.
package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list allows you to list workflows within the compute service.",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Currently not implemented")
	},
}
