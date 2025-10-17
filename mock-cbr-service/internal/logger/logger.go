package logger

import (
	"log/slog"
	"os"
)

type Logger interface {
	Info(message string, args ...any)
	Error(message string, args ...any)
	Warn(message string, args ...any)
	Debug(message string, args ...any)
}

type log struct {
	log *slog.Logger
}

func New(env string) Logger {
	var l *log

	switch env {
	case "local":
		l.log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		l.log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case "prod":
		l.log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	default:
		l.log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	}

	return l
}

func (l *log) Info(message string, args ...interface{}) {
	slog.Info(message, args...)
}

func (l *log) Error(message string, args ...interface{}) {
	slog.Error(message, args...)
}

func (l *log) Warn(message string, args ...interface{}) {
	slog.Warn(message, args...)
}

func (l *log) Debug(message string, args ...interface{}) {
	slog.Debug(message, args...)
}
