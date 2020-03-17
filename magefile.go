// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Build builds fuzzctl using `go build`.
func Build() error {
	return sh.RunV(mg.GoCmd(), "build", "./cmd/fuzzctl")
}

// Install installs fuzzctl using `go install`.
func Install() error {
	return sh.RunV(mg.GoCmd(), "install", "./cmd/fuzzctl")
}

// Test runs unit tests using `go test`.
func Test() error {
	return sh.RunV(mg.GoCmd(), "test", "-cover", "-race", "./...")
}
