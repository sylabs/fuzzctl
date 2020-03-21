// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sylabs/fuzzctl/internal/pkg/client"
)

func buildInfoToJSON(bi client.BuildInfo) (string, error) {
	b := &bytes.Buffer{}
	if err := json.NewEncoder(b).Encode(bi); err != nil {
		return "", err
	}
	return b.String(), nil
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Client version.
		bi := client.BuildInfo{
			GitVersion:   gitVersion,
			GitCommit:    gitCommit,
			GitTreeState: gitTreeState,
			BuiltAt:      builtAt,
			GoVersion:    runtime.Version(),
			Compiler:     runtime.Compiler,
			Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		}
		if s, err := buildInfoToJSON(bi); err == nil {
			fmt.Printf("Client version: %v", s)
		} else {
			logrus.WithError(err).Warn("failed to convert client version")
		}

		// Server version.
		if bi, err := c.ServerBuildInfo(context.Background()); err == nil {
			if s, err := buildInfoToJSON(bi); err == nil {
				fmt.Printf("Server version: %v", s)
			} else {
				logrus.WithError(err).Warn("failed to convert server version")
			}
		} else {
			logrus.WithError(err).Warn("failed to get server version")
		}
	},
}
