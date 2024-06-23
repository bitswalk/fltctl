package main

import (
	"github.com/bitswalk/fltctl/pkg/cmd"
)

func main() {
	// VIPER
	// Check available config file.
	// if available then load it.
	// Set config options.

	// COBRA
	// Parse command line flags.
	// Override config options from config file.
	// Execute cmd/COMMAND from parsed command line argument.
	cmd.Execute()
}
