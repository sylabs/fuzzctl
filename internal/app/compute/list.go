// Copyright (c) 2020, Sylabs, Inc. All rights reserved.
package compute

import (
	"context"

	"github.com/shurcooL/graphql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sylabs/compute-cli/internal/pkg/model"
)

var job struct {
	model.Job `graphql:"job(id: $id)"`
}

var ListCmd = &cobra.Command{
	Use:   "list <ID>",
	Short: "list allows you to see your workflows state within the compute service.",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		// create a client
		client := graphql.NewClient("http://localhost:8080/graphql", nil)

		variables := map[string]interface{}{
			"id": graphql.ID(id),
		}

		err := client.Query(context.Background(), &job, variables)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Infof("Job ID: %s info: Name: %s, ID: %s\n", id, job.Name, job.Id)

	},
}
