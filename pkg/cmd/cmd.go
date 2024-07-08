package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fltctl",
	Short: "fltctl is a convenient flatcar image factory.",
	Long: `fltctl is a flatcar factorty, it conveniently
         help you to retrieve a vanilla upstream base image that you can inject with a
         butane config that suits your need, it automatically transpile that config to the ignition format.`,
	Run: getImage,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
