// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type User struct {
	ID    string
	Login string
}

type Job struct {
	ID       string
	Name     string
	Image    string
	Command  []string
	Status   string
	ExitCode *int32
	Output   string
	Requires []Job
}

type Workflow struct {
	ID         string
	Name       string
	CreatedBy  User
	CreatedAt  string
	StartedAt  string
	FinishedAt string
	Status     string
	Jobs       []Job
}

type Viewer struct {
	Workflows []Workflow
}

func (wf Workflow) String() string {
	return fmt.Sprintf("Name: %s, ID: %s", wf.Name, wf.ID)
}

// BuildInfo represents build information about a component.
type BuildInfo struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuiltAt      string `json:"builtAt"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String returns the string representation of bi.
func (bi BuildInfo) String() string {
	b := &bytes.Buffer{}
	if err := json.NewEncoder(b).Encode(bi); err != nil {
		return ""
	}
	return b.String()
}
