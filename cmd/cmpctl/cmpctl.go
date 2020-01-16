// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package main

import (
	"os"

	"github.com/sylabs/compute-cli/internal/app/cli"
)

func main() {
	if err := cli.CmpctlCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
