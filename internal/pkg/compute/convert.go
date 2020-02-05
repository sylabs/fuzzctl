// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import "github.com/sylabs/compute-cli/internal/pkg/schema"

// convertWorkflow returns a populated workflow structure given a
// GraphQL formated workflow structure
func convertWorkflow(sw schema.Workflow) (w Workflow) {
	w.ID = sw.ID
	w.Name = sw.Name
	w.CreatedBy = User(sw.CreatedBy)
	w.CreatedAt = sw.CreatedAt
	w.StartedAt = sw.StartedAt
	w.FinishedAt = sw.FinishedAt
	w.Status = sw.Status
	for _, je := range sw.Jobs.Edges {
		w.Jobs = append(w.Jobs, convertJob(je.Node))
	}

	return w
}

// convertJob returns a populated job structure given a
// GraphQl formated job structure
func convertJob(sj schema.Job) (j Job) {
	j.ID = sj.ID
	j.Name = sj.Name
	j.Image = sj.Image
	j.Command = sj.Command
	j.Status = sj.Status
	j.ExitCode = sj.ExitCode

	return j
}
