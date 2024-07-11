package organizer

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/belovetech/file-organizer/utils"
)

func Organize(directory string, dryRun bool, verbose bool) error {
	logger := utils.Log()
	logger.Info("Starting organize")

	files, err := getFiles(directory)

	if err != nil {
		logger.Error("Error reading directory", slog.Any("error", err))
		return fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {

		if file.IsDir() {
			// logger.Info("Found directory", slog.String("directory", file.Name()))
			continue
			//Todo: Implement recursive directory search
		}

		fileExt := getExt(file)
		newDir := filepath.Join(directory, fileExt)
		oldFile := filepath.Join(directory, file.Name())
		newFile := filepath.Join(newDir, file.Name())

		if verbose {
			logger.Info("Verbose: Processing file", slog.String("file", file.Name()), slog.String("directory", newDir))
			return nil
		}

		if dryRun {
			logger.Info("Dry run: Would have created directory", slog.String("directory", newDir))
			logger.Info("Dry run: Would have moved file", slog.String("file", file.Name()), slog.String("directory", newDir))
			return nil
		}

		if _, err := os.Stat(newDir); os.IsNotExist(err) {
			err := os.Mkdir(newDir, 0755)
			if err != nil {
				logger.Error("Error creating directory", slog.Any("error", err))
				return fmt.Errorf("error creating directory: %v", err)
			}
			logger.Info("Created directory", slog.String("directory", newDir))
		}

		e := os.Rename(oldFile, newFile)
		if e != nil {
			logger.Error("Error moving file", slog.Any("error", err))
		}
		logger.Info("Moved file", slog.String("file", file.Name()), slog.String("directory", newDir))

	}
	logger.Info("Finished organizing files")
	return nil
}

func getFiles(directory string) ([]fs.DirEntry, error) {
	logger := utils.Log()
	files, err := os.ReadDir(directory)
	if err != nil {
		logger.Error("Error reading directory", slog.Any("error", err.Error()))
		return nil, err
	}
	return files, nil
}

func getExt(file fs.DirEntry) string {
	logger := utils.Log()
	fileExt := filepath.Ext(file.Name())
	if fileExt == "" {
		logger.Info("Found file without extension", slog.String("file", file.Name()))
		fileExt = "no_extension"
	} else {
		logger.Info("Found file", slog.String("file", file.Name()))
		fileExt = fileExt[1:]
	}
	return fileExt
}
