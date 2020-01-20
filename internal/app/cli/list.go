// Copyright (c) 2020, Sylabs, Inc. All rights reserved.
package cli

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const lineFmt = "%s\t%s\t%s\n"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list allows you to list workflows within the compute service.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		wfs, err := c.List(context.TODO())
		if err != nil {
			logrus.Fatal(err)
		}

		tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(tw, lineFmt, "NAME", "ID", "STATUS")
		for _, w := range wfs {
			fmt.Fprintf(tw, lineFmt, w.Name, w.Id, "QUEUED")
		}
		tw.Flush()
	},
}
