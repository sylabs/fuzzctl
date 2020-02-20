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

const listLineFmt = "%s\t%s\t%s\n"
const verboseListLineFmt = "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n"

var (
	verbose bool
)

func init() {
	listCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable detailed workflow information")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list allows you to list workflows within Fuzzball.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		wfs, err := c.List(context.Background())
		if err != nil {
			logrus.Fatal(err)
		}

		// TODO: consider adding behavior to the workflow type to handle this formatting
		// would make it cleaner to supply N/A when no time is available
		tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		if verbose {
			fmt.Fprintf(tw, verboseListLineFmt, "NAME", "ID", "OWNER", "OWNER ID", "STATUS", "CREATED", "STARTED", "FINISHED")
			for _, w := range wfs {
				if w.StartedAt == "" {
					w.StartedAt = "N/A"
				}
				if w.FinishedAt == "" {
					w.FinishedAt = "N/A"
				}
				fmt.Fprintf(tw, verboseListLineFmt, w.Name, w.ID, w.CreatedBy.Login, w.CreatedBy.ID, w.Status, w.CreatedAt, w.StartedAt, w.FinishedAt)
			}
		} else {
			fmt.Fprintf(tw, listLineFmt, "NAME", "ID", "STATUS")
			for _, w := range wfs {
				fmt.Fprintf(tw, listLineFmt, w.Name, w.ID, w.Status)
			}
		}
		tw.Flush()
	},
}
