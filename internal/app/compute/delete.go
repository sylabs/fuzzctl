// Copyright (c) 2020, Sylabs, Inc. All rights reserved.
package compute

import (
	"context"

	"github.com/shurcooL/graphql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sylabs/compute-cli/internal/pkg/model"
)

var deleteJob struct {
	model.Job `graphql:"deleteJob(id: $id)"`
}

var DeleteCmd = &cobra.Command{
	Use:   "delete <ID>",
	Short: "delete allows you to remove a workflow from the compute service queue.",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		// create a client
		client := graphql.NewClient("http://localhost:8080/graphql", nil)

		variables := map[string]interface{}{
			"id": graphql.ID(id),
		}

		err := client.Mutate(context.Background(), &deleteJob, variables)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Infof("Deleted job: Name: %s, ID: %s\n", deleteJob.Name, deleteJob.Id)
	},
}
