// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package parse

import (
	"fmt"
	"sort"

	"github.com/sylabs/compute-cli/internal/pkg/schema"
)

const specVersion = "0.1"

type Decoder interface {
	Decode(v interface{}) error
}

type Parser struct {
	d Decoder
}

func New(d Decoder) Parser {
	return Parser{d}
}

// NOTE(ian): This intermediate struct is needed due to the current
// specification accepting an map of structs while the
// graphQL api expects a list
type workflowSpecIntermediate struct {
	Name    string
	Jobs    map[string]*jobSpecIntermediate
	Volumes map[string]*schema.VolumeSpec
}

type jobSpecIntermediate struct {
	Name     string
	Image    string
	Command  []string
	Requires []string
	Volumes  map[string]*schema.VolumeRequirement
}

// ParseWorkflowSpec converts a specification for a workflow
// into a structure that can be fed to a GraphQL client
// as an input object. Due to the differing formats of the
// specification and required GraphQL client structure
// an intermediate is required for conversion.
func (p Parser) ParseWorkflowSpec() (schema.WorkflowSpec, error) {
	s := struct {
		Version  string
		Workflow workflowSpecIntermediate
	}{}

	if err := p.d.Decode(&s); err != nil {
		return schema.WorkflowSpec{}, err
	}

	if s.Version != specVersion {
		return schema.WorkflowSpec{}, fmt.Errorf("unknown spec version: %s. Expected: %s", s.Version, specVersion)
	}

	// convert intermediate structure
	var w schema.WorkflowSpec
	w.Name = s.Workflow.Name
	for n, ji := range s.Workflow.Jobs {
		var j schema.JobSpec
		j.Name = n
		j.Image = ji.Image
		j.Command = ji.Command
		j.Requires = ji.Requires
		for n, vr := range ji.Volumes {
			vr.Name = n
			j.Volumes = append(j.Volumes, *vr)
		}

		sort.Slice(j.Volumes, func(i, k int) bool { return j.Volumes[i].Name < j.Volumes[k].Name })
		w.Jobs = append(w.Jobs, j)
	}
	for n, v := range s.Workflow.Volumes {
		v.Name = n
		w.Volumes = append(w.Volumes, *v)
	}

	// sort slices by name to ensure deterministic specs
	sort.Slice(w.Jobs, func(i, j int) bool { return w.Jobs[i].Name < w.Jobs[j].Name })
	sort.Slice(w.Volumes, func(i, j int) bool { return w.Volumes[i].Name < w.Volumes[j].Name })

	return w, nil
}
