// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package client

import "fmt"

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
