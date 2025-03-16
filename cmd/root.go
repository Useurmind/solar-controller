package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "solar-controller",
	Short: "solar controller is an app that can controll consumers in your household depending on the current solar power",
}

func initRoot() {
	initServer(rootCmd)
}

func Execute() {
	initRoot()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
