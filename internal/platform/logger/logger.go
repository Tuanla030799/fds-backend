package logger

import (
	"log/slog"
	"os"
)

type Logger struct{ *slog.Logger }

func New(env string) *Logger {
	level := slog.LevelInfo
	if env == "development" {
		level = slog.LevelDebug
	}
	return &Logger{Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))}
}
