// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package config

import (
	"errors"
	"io"

	"gopkg.in/yaml.v3"
)

const (
	// DefaultBaseURI is the default URI for the fuzzball-service
	DefaultBaseURI = "http://localhost:8080"
)

var (
	// ErrRemoteNotFound is returned with the specified remote is not found.
	ErrRemoteNotFound = errors.New("remote not found")
)

type rawConfig struct {
	Remotes map[string]*remote `yaml:"remotes,omitempty"` // List of remotes
}

// Config represents a configuration.
type Config struct {
	raw rawConfig
	// ar contains the current remote profile
	ar Remote
}

// Default returns a defalt Config.
func Default() (*Config, error) {
	return &Config{
		raw: rawConfig{
			Remotes: map[string]*remote{
				"default": {
					BaseURI: DefaultBaseURI,
				},
			},
		},
	}, nil
}

// Read reads a config from the secified reader.
func Read(r io.Reader) (*Config, error) {
	var c Config
	if err := yaml.NewDecoder(r).Decode(&c.raw); err != nil {
		return nil, err
	}
	return &c, nil
}

// Write writes c to the specified writer.
func (c *Config) Write(w io.Writer) error {
	return yaml.NewEncoder(w).Encode(c.raw)
}

// GetActiveRemote returns the active remote.
func (c *Config) GetActiveRemote() (Remote, error) {
	if c.ar != nil {
		return c.ar, nil
	}
	// TODO: read the active config from the config
	r, ok := c.raw.Remotes["default"]
	if !ok {
		return nil, ErrRemoteNotFound
	}
	c.ar = r
	return r, nil
}
