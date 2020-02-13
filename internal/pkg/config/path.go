// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package config

import (
	"os"
	"path"
)

const (
	appName  = "fuzzbomb"
	fileName = "config.yaml"
)

// GetPath returns the path to the configuration file.
func GetPath() (string, error) {
	ucd, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return path.Join(ucd, appName, fileName), nil
}
