package logger

import "log/slog"

const runtimeCallerDepth int = 2

const (
	LevelTrace slog.Level = slog.Level(-8)
	LevelDebug slog.Level = slog.LevelDebug
	LevelInfo  slog.Level = slog.LevelInfo
	LevelWarn  slog.Level = slog.LevelWarn
	LevelError slog.Level = slog.LevelError
	LevelFatal slog.Level = slog.Level(12)
)
