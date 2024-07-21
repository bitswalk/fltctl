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

var scheme = "https"

// ref: https://dev.to/douglasmakey/optimizing-large-file-transfers-in-linux-with-go-an-exploration-of-tcp-and-syscall-15eo
func getImage(cmd *cobra.Command, args []string) {
	// Check local cache first
	// If file exist don't do anything except telling user.
	// Try to create the local cache path and file.
	// Trigger the remote url call
	// Store the resulting data onto local cache file path location.

	domain, _ := cmd.Flags().GetString("domain")
	arch, _ := cmd.Flags().GetString("arch")
	channel, _ := cmd.Flags().GetString("channel")
	imgVersion, _ := cmd.Flags().GetString("version")
	logType, _ := cmd.Flags().GetString("log")
	url := scheme + "://" + channel + "." + domain + "/" + arch + "-usr/" + imgVersion + "/"

	fmt.Println("Current requested log format:", logType)
	fmt.Println("Retrieving upstream image:", url)
}
