package cmd

import (
	"fmt"
	"os"

	"github.com/janpieper/jm/internal/marker"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create jump marker in current directory",
	Run: func(cmd *cobra.Command, args []string) {
		if !marker.Exists() {
			err := marker.Create()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
