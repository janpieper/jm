package cmd

import (
	"fmt"
	"os"

	"github.com/janpieper/jm/internal/marker"
	"github.com/janpieper/jm/internal/table"
	"github.com/spf13/cobra"
)

var detailedList bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List jump markers",
	Run: func(cmd *cobra.Command, args []string) {
		matches, err := marker.List(alternativeMarkers)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if detailedList {
			showDetailedList(matches)
		} else {
			showSimpleList(matches)
		}
	},
}

func showDetailedList(matches []marker.MarkerMatch) {
	t := table.New([]string{"Path", "Marker"})

	for _, match := range matches {
		err := t.Append([]string{match.Path, match.Marker})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	err := t.Render()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func showSimpleList(matches []marker.MarkerMatch) {
	for _, match := range matches {
		fmt.Println(match.Path)
	}
}

func init() {
	listCmd.Flags().StringSliceVarP(&alternativeMarkers, "alternate", "a", []string{}, "Alternative markers")
	listCmd.Flags().BoolVarP(&detailedList, "details", "d", false, "Show details")
	rootCmd.AddCommand(listCmd)
}
