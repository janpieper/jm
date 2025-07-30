package cmd

import (
	"fmt"

	"github.com/janpieper/jm/common"
	"github.com/spf13/cobra"
)

var short bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		version := common.GetVersion()

		fmt.Printf("%s: %s\n", "go-example", version)

		if short {
			return
		}

		fmt.Printf("  BuildDate:      %s\n", version.BuildDate)
		fmt.Printf("  GitCommit:      %s\n", version.GitCommit)
		fmt.Printf("  GitTreeState:   %s\n", version.GitTreeState)

		if version.GitTag != "" {
			fmt.Printf("  GitTag:         %s\n", version.GitTag)
		}

		fmt.Printf("  GoVersion:      %s\n", version.GoVersion)
		fmt.Printf("  Compiler:       %s\n", version.Compiler)
		fmt.Printf("  Platform:       %s\n", version.Platform)
	},
}

func init() {
	versionCmd.Flags().BoolVar(&short, "short", false, "Print only the version number")
	rootCmd.AddCommand(versionCmd)
}
