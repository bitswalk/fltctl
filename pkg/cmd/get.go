package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// Flags should be constraints by ENUM whenever possible as there is only a finite amount of options supported upstream.
	getCmd.Flags().StringP("channel", "c", "stable", "Select the flatcar release channel you want to use.")
	getCmd.Flags().StringP("type", "t", "production", "Select the flatcar type of release you're looking for.")
	getCmd.Flags().StringP("arch", "a", "arm64", "Select the flatcar architecture of your image.")
	getCmd.Flags().StringP("version", "v", "current", "Select the flatcar version you want to use, use the latest available by default.")
	getCmd.Flags().StringP("oem", "o", "", "Select the flatcar oem specific type you want to use, default to none.")
	getCmd.Flags().StringP("image", "i", "flatcar_production_image.bin.bz2", "Select the flatcar image type you want to use, default to flatcar_production_image.bin.bz2.")
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve upstream image",
	Long:  `Retrieve a flatcar image`,
	Run:   getImage,
}
