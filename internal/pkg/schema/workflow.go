// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package schema

// WorkflowSpec is an input so it requires json struct tags to have
// correct capitalization when used with github.com/shurcooL/graphql
type WorkflowSpec struct {
	Name string    `json:"name"`
	Jobs []JobSpec `json:"jobs"`
}

type User struct {
	ID    string
	Login string
}

type Workflow struct {
	ID         string
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
