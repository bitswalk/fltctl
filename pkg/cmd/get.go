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
	Run:   getImage,
}

func getImage(cmd *cobra.Command, args []string) {
	logType, _ := cmd.Flags().GetString("log")
	fmt.Println("Current requested log format:", logType)
	fmt.Println("Retrieving upstream image.")
}
