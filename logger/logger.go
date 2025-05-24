package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
)

var Logger *slog.Logger

func init() {
	// Open the log file or create it if it doesn't exist.
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}

	// Create a multi-writer to output to both stdout and file
	multiWriter := io.MultiWriter(os.Stdout, file)

	// Create a new logger with a custom handler
	handler := slog.NewJSONHandler(multiWriter, nil)
	Logger = slog.New(handler)
}
