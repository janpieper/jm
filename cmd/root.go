package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var alternativeMarkers []string

var rootCmd = &cobra.Command{
	Use:   "jm",
	Short: "Manage jump markers (jm).",
	Long:  "Jump marker (jm) lookup and management.",
}

func Execute() {
	// Disable the `completion` command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
