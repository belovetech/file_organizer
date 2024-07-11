package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/belovetech/file-organizer/cmd"
	"github.com/belovetech/file-organizer/organizer"
	"github.com/belovetech/file-organizer/utils"
)

func main() {
	logger := utils.Log()

	start := time.Now()

	args := os.Args[1:]

	// check if there are any arguments
	if len(args) == 0 {
		logger.Error("No arguments provided. Please provide a directory to organize")
		os.Exit(1)
	}

	directory, dryRun, verbose := cmd.FlagParser()
	err := organizer.Organize(*directory, *dryRun, *verbose)

	if err != nil {
		logger.Error("Error organizing directory", slog.Any("error", err))
		os.Exit(1)
	}

	duration := time.Since(start)
	logger.Info("Organize completed", slog.String("duration", duration.String()))

}
