package cmd

import (
	"log/slog"
	"os"

	payload "github.com/bitswalk/fltctl/internal/cmd"
	"github.com/spf13/cobra"
)

func init() {
	// Flags should be constraints by ENUM whenever possible as there is only a finite amount of options supported upstream.
	getCmd.Flags().StringP("channel", "c", "stable", "Select the flatcar release channel you want to use.")
	getCmd.Flags().StringP("type", "t", "production", "Select the flatcar type of release you're looking for.")
	getCmd.Flags().StringP("arch", "a", "x86_64", "Select the flatcar architecture of your image, can be 'x86_64' or 'aarch64'.")
	getCmd.Flags().StringP("version", "v", "current", "Select the flatcar version you want to use, use the latest available by default.")
	getCmd.Flags().StringP("oem", "o", "", "Select the flatcar oem specific type you want to use, default to none.")
	getCmd.Flags().StringP("image", "i", "flatcar_production_image.bin.bz2", "Select the flatcar image type you want to use, default to flatcar_production_image.bin.bz2.")
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve upstream image",
	Long:  `Retrieve a flatcar image`,
	Run:   get,
}

func get(cmd *cobra.Command, args []string) {
	scheme := ""
	tls, _ := cmd.Flags().GetBool("tls")
	arch, _ := cmd.Flags().GetString("arch")

	// Check architecture type.
	switch arch {
	case "x86_64":
		arch = "amd64"
	case "aarch64":
		arch = "arm64"
	default:
		slog.Error("An error occured while checking image architecture, only x86_64 or aarch64 are valid values.")
		os.Exit(1)
	}

	switch tls {
	case true:
		scheme = "https"
	case false:
		scheme = "http"
	default:
		scheme = "https"
	}

	// Get command flags values.
	domain, _ := cmd.Flags().GetString("domain")
	channel, _ := cmd.Flags().GetString("channel")
	version, _ := cmd.Flags().GetString("version")
	imgType, _ := cmd.Flags().GetString("type")
	endpoint := scheme + "://" + channel + "." + domain + "/" + arch + "-usr/" + version + "/"
	uri := "flatcar_" + imgType + "_iso_image.iso"
	url := endpoint + uri
	// emptyurl := ""

	if _, err := payload.Locate(uri); err != nil {
		slog.Error("An error occured while checking remote image availability.", "msg", err)
		os.Exit(1)
	}
	slog.Info("Image found remotely, will try to download it.", "url", url)

	if _, err := payload.Check(url); err != nil {
		slog.Error("An error occured while checking remote image availability.", "msg", err)
		os.Exit(1)
	}
	slog.Info("Image found remotely, will try to download it.", "url", url)

	if _, err := payload.Retrieve(url); err != nil {
		slog.Error("An error occured while trying to retrieve the image.", "msg", err)
		os.Exit(1)
	}

}
