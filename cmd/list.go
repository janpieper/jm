package cmd

import (
	"fmt"
	"os"

	"github.com/janpieper/jm/internal/marker"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List jump markers",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := marker.List(alternativeMarkers)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, file := range files {
			fmt.Println(file)
		}
	},
}

func init() {
	listCmd.Flags().StringSliceVarP(&alternativeMarkers, "alternate", "a", []string{}, "Alternative markers")
	rootCmd.AddCommand(listCmd)
}
