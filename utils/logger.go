package utils

import (
	"log/slog"
	"os"
)

// Custom logger function
// Returns a logger object
func Log() *slog.Logger {
	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	mylogger := slog.New(logHandler)
	return mylogger
}
