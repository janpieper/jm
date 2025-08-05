package cmd

import (
	"fmt"
	"os"

	"github.com/janpieper/jm/internal/marker"
	"github.com/janpieper/jm/internal/table"
	"github.com/spf13/cobra"
)

var detailed bool

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show current jump marker",
	Run: func(cmd *cobra.Command, args []string) {
		match, err := marker.Current(alternativeMarkers)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if detailed {
			showDetailed(match)
		} else {
			showSimple(match)
		}

	},
}

func showSimple(match *marker.MarkerMatch) {
	if match != nil {
		fmt.Printf("%s\n", match.Path)
	}
}

func showDetailed(match *marker.MarkerMatch) {
	t := table.New([]string{"Path", "Marker"})

	if match != nil {
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

func init() {
	currentCmd.Flags().StringSliceVarP(&alternativeMarkers, "alternate", "a", []string{}, "Alternative markers")
	currentCmd.Flags().BoolVarP(&detailed, "details", "d", false, "Show details")
	rootCmd.AddCommand(currentCmd)
}
