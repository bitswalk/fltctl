package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "fltctl",
	Version: version,
	Short:   "fltctl is a convenient flatcar image factory.",
	Long: `fltctl is a flatcar factory, it conveniently help you to retrieve
a vanilla upstream base image that you can inject with a butane config
that suits your need, it automatically transpile that config onto an ignition file,
mount the appropriate upstream vanilla upstream system partition, push it on and
umount the partition, letting you with a golden iso image.`,
}

func Execute() {
	rootCmd.PersistentFlags().StringP("log", "l", "text", "Log format to be used, can be 'text' or 'json'.")
	rootCmd.PersistentFlags().StringP("debug", "V", "info", "Log level to be used, can be one of 'info' 'warn' 'error' or 'debug'.")
	rootCmd.PersistentFlags().StringP("config", "f", "~/.config/fltctl/config.yaml", "fltctl config file to use.")
	rootCmd.PersistentFlags().StringP("domain", "d", "release.flatcar-linux.net", "fltctl default domain to construct images from.")
	rootCmd.PersistentFlags().BoolP("tls", "s", true, "Try to retrieve images from a TLS termination by default, can fallback to non-tls.")
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
