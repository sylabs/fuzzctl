// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package cli

import (
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout logs you out of Fuzzball.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tokenSrc = nil
	},
}
