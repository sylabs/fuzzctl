// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import "fmt"

type Workflow struct {
	Id   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
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
