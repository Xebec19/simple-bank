package logger

import "log/slog"

func LogInfo(msg string, args interface{}) {
	slog.Info(msg, args)
}

func LogError(msg string, args interface{}) {
	slog.Error(msg, args)
}

func LogWarn(msg string, args interface{}) {
	slog.Warn(msg, args)
}
