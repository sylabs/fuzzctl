// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import "fmt"

type User struct {
	Id    string
	Login string
}

type Workflow struct {
	Id         string
	Name       string
	CreatedBy  User
	CreatedAt  string
	StartedAt  string
	FinishedAt string
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
