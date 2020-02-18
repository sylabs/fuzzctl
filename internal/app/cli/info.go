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

const infoLineFmt = "%s\t%s\t%s\t%s\n"

var infoCmd = &cobra.Command{
	Use:   "info <ID>",
	Short: "info allows you to see a workflow's state within the compute service.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		wf, err := c.Info(context.Background(), id)
		if err != nil {
			logrus.Fatal(err)
		}

		// TODO: verbose output
		fmt.Printf("ID: %s\nNAME: %s\nJOBS:\n", wf.ID, wf.Name)
		tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(tw, infoLineFmt, "NAME", "ID", "STATUS", "EXITCODE")
		for _, j := range wf.Jobs {
			exitCode := "N/A"
			if j.ExitCode != nil {
				exitCode = fmt.Sprintf("%d", *j.ExitCode)
			}
			fmt.Fprintf(tw, infoLineFmt, j.Name, j.ID, j.Status, exitCode)
		}
		tw.Flush()

		printHeader := true
		for _, j := range wf.Jobs {
			if j.Output != "" {
				if printHeader {
					printHeader = false
					fmt.Printf("\nJOB OUTPUT:\n")
				}
				fmt.Printf("=== %v ===:\n", j.Name)
				fmt.Printf("%v", j.Output)
			}
		}
	},
}
