package logging

import (
	"flag"
	"io"
	"log/slog"
	"os"
)

// Logger struct for writing logs with slog.Logger in io.Writer.
type Logger struct {
	Logger *slog.Logger
	Writer io.Writer
}

const (
	flagFile  = "-file"
	flagLocal = "-local"
	flagDev   = "dev"
	flagProd  = "prod"
)

// New create a new dir "logs" and "all.log" file for writing logs using slog.Logger.
func New() *Logger {

	var writer io.Writer

	var logger *slog.Logger

	flag.Parse()

	switch flag.Arg(1) {
	case flagFile:
		writer = getLogFile()

	default:
		writer = os.Stdout
	}

	switch flag.Arg(2) {
	case flagLocal:
		logger = slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case flagDev:
		logger = slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: slog.LevelInfo}))

	case flagProd:
		logger = slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: slog.LevelError}))

	default:
		logger = slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return &Logger{
		Logger: logger,
		Writer: writer,
	}
}

func getLogFile() *os.File {
	err := os.RemoveAll("logs")
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("logs", 0777)
	if err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	return logFile
}
