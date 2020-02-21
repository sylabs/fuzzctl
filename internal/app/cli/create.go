// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package cli

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sylabs/compute-cli/internal/pkg/parse"
	"github.com/sylabs/compute-cli/internal/pkg/schema"
	"gopkg.in/yaml.v3"
)

var createCmd = &cobra.Command{
	Use:   "create <WORKLOW_CONFIG_PATH>",
	Short: "create enables you to submit a workflow to the Fuzzball queue.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		spec, err := parseWorkflowSpec(path)
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

func parseWorkflowSpec(path string) (schema.WorkflowSpec, error) {
	specFile, err := os.Open(path)
	if err != nil {
		return schema.WorkflowSpec{}, err
	}
	defer specFile.Close()

	// Since YAML is a superset of JSON, we can actually accept both
	// with this decoder.
	d := yaml.NewDecoder(specFile)
	d.KnownFields(true)

	p := parse.New(d)
	spec, err := p.ParseWorkflowSpec()
	if err != nil {
		return schema.WorkflowSpec{}, err
	}
	return spec, nil
}
