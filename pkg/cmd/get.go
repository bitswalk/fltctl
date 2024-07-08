package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve upstream image",
	Long:  `Retrieve a flatcar image`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Retrieving upstream image.")
	},
}
