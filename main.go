package main

import (
	"github.com/bitswalk/fltctl/pkg/cmd"
)

func main() {
	// COBRA
	// Parse command line flags.
	// Override config options from config file.
	// Execute cmd/COMMAND from parsed command line argument.
	cmd.Execute()
}
