package cmd

import (
	"fmt"
	"os"

	"github.com/janpieper/jm/internal/marker"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove jump marker from current directory",
	Run: func(cmd *cobra.Command, args []string) {
		err := marker.Remove()
		if err != nil && !os.IsNotExist(err) {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
