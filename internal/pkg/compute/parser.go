// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import (
	"fmt"

	"github.com/sylabs/compute-cli/internal/pkg/schema"
	"gopkg.in/yaml.v2"
)

const specVersion = "0.1"

// NOTE(ian): This intermediate struct is needed due to the current
// yaml specification accepting an map of structs while the
// graphQL api expects a list
type workflowSpecIntermediate struct {
	Name string
	Jobs map[string]*schema.JobSpec
}

// ParseSpec converts a yaml specification for a workflow
// into a structure that can be fed to a GraphQL client
// as an input object. Due to the differing formats of the
// yaml specification and required GraphQL client structure
// an intermediate is required.
func ParseSpec(b []byte) (schema.WorkflowSpec, error) {
	s := struct {
		Version  string
		Workflow workflowSpecIntermediate
	}{}

	if err := yaml.Unmarshal(b, &s); err != nil {
		return schema.WorkflowSpec{}, err
	}

	if s.Version != specVersion {
		return schema.WorkflowSpec{}, fmt.Errorf("unknown spec version: %s. Expected: %s", s.Version, specVersion)
	}

	// convert intermediate structure
	var w schema.WorkflowSpec
	w.Name = s.Workflow.Name
	for n, j := range s.Workflow.Jobs {
		j.Name = n
		w.Jobs = append(w.Jobs, *j)
	}

	return w, nil
}
