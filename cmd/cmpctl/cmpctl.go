// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/sylabs/compute-cli/internal/app/compute"
)

var (
	httpAddr string
)

var rootCmd = &cobra.Command{
	Use:   "cmpctl",
	Short: "cmpctl enables control of workflows for the compute service.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&httpAddr, "http_addr", "http://localhost:8080", "Address to reach compute server")
	rootCmd.AddCommand(compute.CreateCmd)
	rootCmd.AddCommand(compute.DeleteCmd)
	rootCmd.AddCommand(compute.ListCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
