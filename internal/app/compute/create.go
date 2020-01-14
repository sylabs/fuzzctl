// Copyright (c) 2020, Sylabs, Inc. All rights reserved.
package compute

import (
	"context"

	"github.com/shurcooL/graphql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sylabs/compute-cli/internal/pkg/model"
)

var createJob struct {
	model.Job `graphql:"createJob(name: $name)"`
}

var CreateCmd = &cobra.Command{
	Use:   "create <NAME>",
	Short: "create enables you to submit a workflow to the compute service queue.",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		// create a client
		client := graphql.NewClient("http://localhost:8080/graphql", nil)

		variables := map[string]interface{}{
			"name": graphql.String(name),
		}

		err := client.Mutate(context.Background(), &createJob, variables)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Infof("Created job: Name: %s, ID: %s\n", createJob.Name, createJob.Id)

	},
}
