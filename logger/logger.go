package logger

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
)

var loggers = make(map[string]zerolog.Logger) // Initialize the map

func Init(ctx context.Context) (context.Context, *zerolog.Logger) {
	// Initialize main app logger
	return Create(ctx, "app.log")
}

func Create(ctx context.Context, logFileName string) (context.Context, *zerolog.Logger) {
	// Ensure the log file name ends with ".log"
	if !strings.HasSuffix(logFileName, ".log") {
		logFileName += ".log"
	}

	var (
		key = strings.Replace(logFileName, ".log", "", -1)
		log = Get(key)
	)

	if log != nil {
		return ctx, log
	}

	// Ensure the "logs" directory exists
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		panic("Failed to create logs directory: " + err.Error())
	}

	// Create log files
	appLogFile, err := os.OpenFile(filepath.Join("logs", logFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Default level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Create multi-writer for logging to file
	multiWriter := io.MultiWriter(appLogFile, os.Stdout)

	// Set up zerolog with multi-writer
	logger := zerolog.New(multiWriter).With().Timestamp().Logger()

	// Attach the Logger to the context.Context
	ctx = logger.WithContext(ctx)

	loggers[key] = logger

	return ctx, &logger
}

func Get(loggerKey string) *zerolog.Logger {
	if logger, ok := loggers[loggerKey]; ok {
		return &logger
	}
	return nil
}
