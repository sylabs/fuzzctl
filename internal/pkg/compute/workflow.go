// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import "fmt"

// NOTE(ian): the json struct tags are needed by github.com/shurcooL/graphql
// in order to correctly parse input objects into the correct format
// for graphql to consume, note these tags are case sensitive.

// TODO: break this out into another package as it's getting large

type JobSpec struct {
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Command []string `json:"command"`
}

type WorkflowSpec struct {
	Name string    `json:"name"`
	Jobs []JobSpec `json:"jobs"`
}

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

type JobEdge struct {
	Node Job
}

type JobConnection struct {
	Edges []JobEdge
}

type Workflow struct {
	Id         string
	Name       string
	CreatedBy  User
	CreatedAt  string
	StartedAt  string
	FinishedAt string
	Status     string
	Jobs       JobConnection
}

type WorkflowEdge struct {
	Node Workflow
}

type WorkflowConnection struct {
	Edges []WorkflowEdge
}

type Viewer struct {
	Workflows WorkflowConnection
}

func (wf Workflow) String() string {
	return fmt.Sprintf("Name: %s, ID: %s", wf.Name, wf.Id)
}
