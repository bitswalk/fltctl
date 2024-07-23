package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a golden image",
	Long:  `Build a golden image of your own based on a flatcar upstream image`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Retrieving upstream image.")
		// checks.Check()
	},
}
