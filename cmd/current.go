package cmd

import (
	"fmt"
	"os"

	"github.com/janpieper/jm/internal/marker"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show current jump marker",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := marker.Current(alternativeMarkers)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if file != nil {
			fmt.Printf("%s\n", *file)
		}
	},
}

func init() {
	currentCmd.Flags().StringSliceVarP(&alternativeMarkers, "alternate", "a", []string{}, "Alternative markers")
	rootCmd.AddCommand(currentCmd)
}
