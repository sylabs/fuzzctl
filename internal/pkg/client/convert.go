// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/sylabs/fuzzctl/internal/pkg/schema"
)

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
	j.Output = sj.Output

	return j
}

func convertBuildInfo(sbi schema.BuildInfo) BuildInfo {
	bi := BuildInfo{
		GitVersion:   "unknown",
		GitCommit:    "unknown",
		GitTreeState: "unknown",
		BuiltAt:      "unknown",
		GoVersion:    sbi.GoVersion,
		Compiler:     sbi.Compiler,
		Platform:     sbi.Platform,
	}

	if v := sbi.GitVersion; v != nil {
		b := &strings.Builder{}
		fmt.Fprintf(b, "%v.%v.%v", v.Major, v.Minor, v.Patch)
		if v.PreRelease != nil {
			fmt.Fprintf(b, "-%v", *v.PreRelease)
		}
		if v.BuildMetadata != nil {
			fmt.Fprintf(b, "+%v", *v.BuildMetadata)
		}
		bi.GitVersion = b.String()
	}
	if c := sbi.GitCommit; c != nil {
		bi.GitCommit = *c
	}
	if s := sbi.GitTreeState; s != nil {
		bi.GitTreeState = *s
	}
	if t := sbi.BuiltAt; t != nil {
		bi.BuiltAt = t.Format(time.RFC3339)
	}

	return bi
}
