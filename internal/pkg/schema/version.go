// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package schema

import "time"

// BuildInfo represents build information about a component.
type BuildInfo struct {
	GitVersion   *Version   `json:"gitVersion"`
	GitCommit    *string    `json:"gitCommit"`
	GitTreeState *string    `json:"gitTreeState"`
	BuiltAt      *time.Time `json:"builtAt"`
	GoVersion    string     `json:"goVersion"`
	Compiler     string     `json:"compiler"`
	Platform     string     `json:"platform"`
}

// Version represents semantic version information.
type Version struct {
	Major         int32   `json:"major"`
	Minor         int32   `json:"minor"`
	Patch         int32   `json:"patch"`
	PreRelease    *string `json:"preRelease"`
	BuildMetadata *string `json:"buildMetadata"`
}
