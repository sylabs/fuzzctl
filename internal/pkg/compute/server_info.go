// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import "fmt"

type ServerInfo struct {
	Hostname        string
	CPUArchitecture string
	OSPlatform      string
	Memory          int
	Capabilities    []Capability
}

type Capability struct {
	Key   string
	Value string
}

func (si ServerInfo) String() string {
	s := fmt.Sprintf("Hostname: %s, CPU Arch: %s, OS Platform: %s, Memory: %dMB, Capabilities: ",
		si.Hostname,
		si.CPUArchitecture,
		si.OSPlatform,
		si.Memory)

	if len(si.Capabilities) == 0 {
		// no capabilities listed by server
		return s + "none"
	}

	for _, c := range si.Capabilities {
		s += fmt.Sprintf("%s:%s ", c.Key, c.Value)
	}

	return s
}
