// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

// +build mage

package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/sirupsen/logrus"
)

// ldFlags returns standard linker flags to pass to various Go commands.
func ldFlags() string {
	pkgPath := "github.com/sylabs/fuzzctl/internal/app/cli"
	vals := []string{
		fmt.Sprintf("-X %v.builtAt=%v", pkgPath, time.Now().UTC().Format(time.RFC3339)),
	}

	// Attempt to get git details.
	d, err := describeHead()
	if err == nil {
		vals = append(vals, fmt.Sprintf("-X %v.gitCommit=%v", pkgPath, d.ref.Hash().String()))

		if d.isClean {
			vals = append(vals, fmt.Sprintf("-X %v.gitTreeState=clean", pkgPath))
		} else {
			vals = append(vals, fmt.Sprintf("-X %v.gitTreeState=dirty", pkgPath))
		}

		if v, err := getVersion(d); err != nil {
			logrus.WithError(err).Warn("failed to get version from git description")
		} else {
			vals = append(vals, fmt.Sprintf("-X %v.gitVersion=%v", pkgPath, v.String()))
		}
	}

	return strings.Join(vals, " ")
}

// Build builds fuzzctl using `go build`.
func Build() error {
	return sh.RunV(mg.GoCmd(), "build", "-ldflags", ldFlags(), "./cmd/fuzzctl")
}

// Install installs fuzzctl using `go install`.
func Install() error {
	return sh.RunV(mg.GoCmd(), "install", "-ldflags", ldFlags(), "./cmd/fuzzctl")
}

// Test runs unit tests using `go test`.
func Test() error {
	return sh.RunV(mg.GoCmd(), "test", "-ldflags", ldFlags(), "-cover", "-race", "./...")
}

// Deb builds a deb package.
func Deb() error {
	mg.Deps(Build)
	return makePackage("deb")
}

// RPM builds a RPM package.
func RPM() error {
	mg.Deps(Build)
	return makePackage("rpm")
}
