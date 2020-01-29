// Copyright (c) 2020, Sylabs, Inc. All rights reserved.
package cli

import (
	"context"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sylabs/compute-cli/internal/pkg/compute"
)

var createCmd = &cobra.Command{
	Use:   "create <WORKLOW_CONFIG_PATH>",
	Short: "create enables you to submit a workflow to the compute service queue.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		data, err := ioutil.ReadFile(path)
		if err != nil {
			logrus.Fatal(err)
		}

		spec, err := compute.ParseSpec(data)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Debugf("Creating workflow: %s", spec.Name)
		wf, err := c.Create(context.Background(), spec)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Infof("Created workflow: %s\n", wf)
	},
}
