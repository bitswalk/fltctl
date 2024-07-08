package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print fltctl release version.",
	Long:  `Show the fltctl release version currently called`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the function that gave you fltctl release version.")
	},
}
