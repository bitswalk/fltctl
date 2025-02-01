package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var scheme = "https"

func checkImage(cmd *cobra.Command, args []string) {
	// Get command flags values.
	domain, _ := cmd.Flags().GetString("domain")
	arch, _ := cmd.Flags().GetString("arch")
	channel, _ := cmd.Flags().GetString("channel")
	version, _ := cmd.Flags().GetString("version")
	imgType, _ := cmd.Flags().GetString("type")
	endpoint := scheme + "://" + channel + "." + domain + "/" + arch + "-usr/" + version + "/"
	uri := "flatcar_" + imgType + "_iso_image.iso"
	url := endpoint + uri
	slog.Info("Checking if image exist remotely.", "url:", url)
}

// ref: https://dev.to/douglasmakey/optimizing-large-file-transfers-in-linux-with-go-an-exploration-of-tcp-and-syscall-15eo
func getImage(cmd *cobra.Command, args []string) {
	// Get command flags values.
	domain, _ := cmd.Flags().GetString("domain")
	arch, _ := cmd.Flags().GetString("arch")
	channel, _ := cmd.Flags().GetString("channel")
	version, _ := cmd.Flags().GetString("version")
	imgType, _ := cmd.Flags().GetString("type")
	endpoint := scheme + "://" + channel + "." + domain + "/" + arch + "-usr/" + version + "/"
	uri := "flatcar_" + imgType + "_iso_image.iso"
	url := endpoint + uri
	slog.Info("Retrieving image remotely.", "url:", url)

	// // Check local cache first
	// // If file exist don't do anything except telling user.
	// if _, err := os.Stat(os.TempDir() + uri); err == nil {
	// 	slog.Info("Image already exist locally.", "uri", uri, "path", os.TempDir())
	// } else {
	// 	// When file doesn't exist, try to create the local cache path and file.
	// 	// Verify ressource availability on remote endpoint before calling for the ressource.
	// 	slog.Info("Image doesn't exist locally, we'll try to download it.", "uri", uri, "path", os.TempDir())
	// 	slog.Info("Checking upstream ressource availability.", "url", url)
	// }

	// // Check the release version we download if we request for latest current release.
	// if version == "current" {
	// 	versionValue := helpers.GetVersion(ctx, endpoint)
	// 	slog.Info("Retrieving latest current release.", "version", versionValue["FLATCAR_VERSION"])
	// }

	// resp, err := http.Head(url)
	// if err != nil {
	// 	slog.Error("An error occured on endpoint availability request.", "error", err)
	// 	os.Exit(128)
	// }

	// if resp.StatusCode != 200 {
	// 	slog.Error("Remote endpoint didn't answered properly.", "url", url)
	// 	os.Exit(128)
	// }

	// // Trigger the remote url call
	// slog.Info("Retrieving upstream ressource:", "url", url, "Content-length", resp.Header.Get("Content-length"))
	// resp, err = http.Get(url)
	// if err != nil {
	// 	slog.Error("An error occured while trying to contact endpoint.", "url", url, "error", err)
	// 	os.Exit(128)
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	slog.Warn("An HTTP Error occured.", "url", url, "error", err)
	// }
	// // Store the resulting data onto local cache file path location.
	// slog.Info("Writting image locally", "path", os.TempDir(), "uri", uri)
	// file, err := os.Create(os.TempDir() + "/" + uri)
	// if err != nil {
	// 	slog.Error("An error occured creating image on local filesystem.", "url", url, "path", os.TempDir(), "uri", uri, "error", err)
	// 	os.Exit(128)
	// }
	// defer file.Close()

	// _, err = io.Copy(file, resp.Body)
	// if err != nil {
	// 	slog.Error("An error occured writting image locally.", "url", url, "path", os.TempDir(), "uri", uri, "error", err)
	// 	os.Exit(128)
	// }
	// slog.Info("Image successfully retrieved locally!", "path", os.TempDir(), "uri", uri)
}
