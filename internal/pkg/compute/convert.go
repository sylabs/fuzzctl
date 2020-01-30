// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import "github.com/sylabs/compute-cli/internal/pkg/schema"

// convertWorkflow returns a populates workflow structure given a
// GraphQL formated workflow structure
func convertWorkflow(sw schema.Workflow) (w Workflow) {
	w.Id = sw.Id
	w.Name = sw.Name
	w.CreatedBy = User(sw.CreatedBy)
	w.CreatedAt = sw.CreatedAt
	w.StartedAt = sw.StartedAt
	w.FinishedAt = sw.FinishedAt
	w.Status = sw.Status
	for _, je := range sw.Jobs.Edges {
		w.Jobs = append(w.Jobs, Job(je.Node))
	}

	return w
}
