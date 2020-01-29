// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import "gopkg.in/yaml.v2"

import "fmt"

const specVersion = "0.1"

// NOTE(ian): This intermediate struct is needed due to the current
// yaml specification accepting an map of structs while the
// graphQL api expects a list
type WorkflowSpecIntermediate struct {
	Name string
	Jobs map[string]*JobSpec
}

// ParseSpec converts a yaml specification for a workflow
// into a structure that can be fed to a GraphQL client
// as an input object. Due to the differing formats of the
// yaml specification and required GraphQL client structure
// an intermediate is required.
func ParseSpec(b []byte) (*WorkflowSpec, error) {
	var w WorkflowSpec
	s := struct {
		Version  string
		Workflow WorkflowSpecIntermediate
	}{}

	if err := yaml.Unmarshal(b, &s); err != nil {
		return nil, err
	}

	if s.Version != specVersion {
		return nil, fmt.Errorf("unknown spec version: %s. Expected: %s", s.Version, specVersion)
	}

	// convert intermediate structure
	w.Name = s.Workflow.Name
	for n, j := range s.Workflow.Jobs {
		j.Name = n
		w.Jobs = append(w.Jobs, *j)
	}

	return &w, nil
}
