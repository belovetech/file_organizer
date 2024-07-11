package cmd

import (
	"flag"
	"log/slog"
	"os"

	"github.com/belovetech/file-organizer/utils"
)

func FlagParser() (*string, *bool, *bool) {
	logger := utils.Log()
	logger.Info("Parsing flags")

	directory := flag.String("dir", ".", "Directory to organize")
	dryRun := flag.Bool("dry-run", false, "Dry run")
	verbose := flag.Bool("verbose", false, "Verbose output")

	flag.Parse()

	if _, err := os.Stat(*directory); os.IsNotExist(err) {
		logger.Error("Directory does not exist", slog.String("directory", *directory))
		os.Exit(1)
	}

	logger.Info("Directory exists", slog.String("directory", *directory))

	return directory, dryRun, verbose
}
