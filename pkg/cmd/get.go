package cmd

import (
	"context"
	"io"
	"net/http"
	"os"

	"github.com/bitswalk/fltctl/internal/helpers"
	"github.com/bitswalk/fltctl/internal/logs"
	"github.com/spf13/cobra"
)

func init() {
	// Flags should be constraints by ENUM whenever possible as there is only a finite amount of options supported upstream.
	getCmd.Flags().StringP("channel", "c", "stable", "Select the flatcar release channel you want to use.")
	getCmd.Flags().StringP("type", "t", "production", "Select the flatcar type of release you're looking for.")
	getCmd.Flags().StringP("arch", "a", "arm64", "Select the flatcar architecture of your image.")
	getCmd.Flags().StringP("version", "v", "current", "Select the flatcar version you want to use, use the latest available by default.")
	getCmd.Flags().StringP("oem", "o", "", "Select the flatcar oem specific type you want to use, default to none.")
	getCmd.Flags().StringP("image", "i", "image", "Select the flatcar image type you want to use, default to image.")
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
	if _, err := os.Stat(os.TempDir() + uri); err == nil {
		logger.Info("Image already exist locally.", "uri", uri, "path", os.TempDir())
	} else {
		// When file doesn't exist, try to create the local cache path and file.
		// Verify ressource availability on remote endpoint before calling for the ressource.
		logger.Info("Image doesn't exist locally, we'll try to download it.", "uri", uri, "path", os.TempDir())
		logger.Info("Checking upstream ressource availability.", "url", url)
	}

	// Check the release version we download if we request for latest current release.
	if version == "current" {
		versionValue := helpers.GetVersion(ctx, endpoint)
		logger.Info("Retrieving latest current release.", "version", versionValue["FLATCAR_VERSION"])
	}

	resp, err := http.Head(url)
	if err != nil {
		logger.Error("An error occured on endpoint availability request.", "error", err)
		os.Exit(128)
	}

	if resp.StatusCode != 200 {
		logger.Error("Remote endpoint didn't answered properly.", "url", url)
		os.Exit(128)
	}

	// Trigger the remote url call
	logger.Info("Retrieving upstream ressource:", "url", url, "Content-length", resp.Header.Get("Content-length"))
	resp, err = http.Get(url)
	if err != nil {
		logger.Error("An error occured while trying to contact endpoint.", "url", url, "error", err)
		os.Exit(128)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logger.Warn("An HTTP Error occured.", "url", url, "error", err)
	}
	// Store the resulting data onto local cache file path location.
	logger.Info("Writting image locally", "path", os.TempDir(), "uri", uri)
	file, err := os.Create(os.TempDir() + "/" + uri)
	if err != nil {
		logger.Error("An error occured creating image on local filesystem.", "url", url, "path", os.TempDir(), "uri", uri, "error", err)
		os.Exit(128)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		logger.Error("An error occured writting image locally.", "url", url, "path", os.TempDir(), "uri", uri, "error", err)
		os.Exit(128)
	}
	logger.Info("Image successfully retrieved locally!", "path", os.TempDir(), "uri", uri)
}
