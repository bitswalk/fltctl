package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a golden image",
	Long:  `Build a golden image of your own based on a flatcar upstream image`,
	Run:   buildImage,
}

func buildImage(cmd *cobra.Command, args []string) {
}
