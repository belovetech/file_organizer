package main

import (
	"os"

	"github.com/belovetech/file-organizer/cmd"
	"github.com/belovetech/file-organizer/organizer"
	"github.com/belovetech/file-organizer/utils"
)

func main() {
	logger := utils.Log()

	logger.Info("Starting organize")

	args := os.Args[1:]

	// check if there are any arguments
	if len(args) == 0 {
		logger.Error("No arguments provided. Please provide a directory to organize")
		os.Exit(1)
	}

	directory, dryRun, verbose := cmd.FlagParser()
	organizer.Organize(*directory, *dryRun, *verbose)

}
