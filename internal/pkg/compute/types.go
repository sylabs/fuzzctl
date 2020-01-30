// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import "fmt"

type User struct {
	Id    string
	Login string
}

type Job struct {
	Id       string
	Name     string
	Image    string
	Command  []string
	Status   string
	ExitCode *int32
}

type Workflow struct {
	Id         string
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
	return fmt.Sprintf("Name: %s, ID: %s", wf.Name, wf.Id)
}
