// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package schema

// JobSpec is an input so it requires json struct tags to have
// correct capitalization when used with github.com/shurcooL/graphql
type JobSpec struct {
	Name     string              `json:"name"`
	Image    string              `json:"image"`
	Command  []string            `json:"command"`
	Requires []string            `json:"requires"`
	Volumes  []VolumeRequirement `json:"volumes,omitempty"`
}

type VolumeRequirement struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Job struct {
	ID       string
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
