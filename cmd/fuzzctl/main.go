// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package main

import (
	"os"

	"github.com/sylabs/fuzzctl/internal/app/cli"
)

func main() {
	if err := cli.FuzzctlCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
