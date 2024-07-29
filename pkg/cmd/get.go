package cmd

import (
	"context"
	"net/http"
	"os"

	"github.com/bitswalk/fltctl/internal/checks"
	"github.com/bitswalk/fltctl/internal/helpers"
	"github.com/bitswalk/fltctl/internal/logs"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.Flags().StringP("channel", "c", "stable", "Select the flatcar release channel you want to use.")
	getCmd.Flags().StringP("type", "t", "production", "Select the flatcar type of release you're looking for.")
	getCmd.Flags().StringP("arch", "a", "arm64", "Select the flatcar architecture of your image.")
	getCmd.Flags().StringP("version", "v", "current", "Select the flatcar version you want to use, use the latest available by default.")
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve upstream image",
	Long:  `Retrieve a flatcar image`,
	Run:   getImage,
}

// ref: https://dev.to/douglasmakey/optimizing-large-file-transfers-in-linux-with-go-an-exploration-of-tcp-and-syscall-15eo
func getImage(cmd *cobra.Command, args []string) {

	logType, _ := cmd.Flags().GetString("log")

	ctx := context.WithValue(cmd.Context(), "logger", logType)
	var logger = logs.NewLoggerWithContext(ctx)
	// Get command flags values.
	domain, _ := cmd.Flags().GetString("domain")
	arch, _ := cmd.Flags().GetString("arch")
	channel, _ := cmd.Flags().GetString("channel")
	version, _ := cmd.Flags().GetString("version")
	imgType, _ := cmd.Flags().GetString("type")
	endpoint := scheme + "://" + channel + "." + domain + "/" + arch + "-usr/" + version + "/"
	uri := "flatcar_" + imgType + "_iso_image.iso"
	url := endpoint + uri

	// Check local cache first
	// If file exist don't do anything except telling user.
	// Try to create the local cache path and file.
	checks.Check(ctx, "cache", uri)

	// Verify ressource availability on remote endpoint before calling for the ressource.
	logger.Info("Checking upstream ressource availability.", "url", url)
	res, err := http.Head(url)
	if err != nil {
		logger.Error("An error occured on endpoint availability request.", "error", err)
	}

	if res.StatusCode != 200 {
		logger.Error("Requested ressource doesn't exist on remote endpoint.", "url", url)
	}

	if version == "current" {
		versionValue := helpers.GetVersion(ctx, endpoint)
		logger.Info("Retrieving latest current release.", "version", versionValue["FLATCAR_VERSION"])
		os.Exit(0)
	}

	// Trigger the remote url call
	logger.Info("Retrieving upstream ressource:", "url", url)
	// Store the resulting data onto local cache file path location.

}
