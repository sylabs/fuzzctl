// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import (
	"fmt"
	"sort"

	"github.com/sylabs/compute-cli/internal/pkg/schema"
	"gopkg.in/yaml.v2"
)

const specVersion = "0.1"

// NOTE(ian): This intermediate struct is needed due to the current
// yaml specification accepting an map of structs while the
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

	if err := yaml.UnmarshalStrict(b, &s); err != nil {
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
