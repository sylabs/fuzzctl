// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import "fmt"

type Workflow struct {
	Id   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
}

func (wf Workflow) String() string {
	return fmt.Sprintf("Name: %s, ID: %s", wf.Name, wf.Id)
}
