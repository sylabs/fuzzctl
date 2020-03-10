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
}

// Default returns a defalt Config.
func Default() (*Config, error) {
	return &Config{
		raw: rawConfig{
			Remotes: map[string]*remote{
				"default": {
					BaseURI: DefaultBaseURI,
					AuthConfig: authConfig{
						Type:             AuthConfigTypeAuthCodePKCE,
						ClientID:         "0oa24wwhwBWYa1T804x6",
						AuthURL:          "https://dev-930666.okta.com/oauth2/default/v1/authorize",
						TokenURL:         "https://dev-930666.okta.com/oauth2/default/v1/token",
						LoginRedirectURL: "http://localhost:9876/authorization/callback",
					},
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
	// TODO: read the active config from the config
	r, ok := c.raw.Remotes["default"]
	if !ok {
		return nil, ErrRemoteNotFound
	}
	return r, nil
}
