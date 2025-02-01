package logs

import (
	"context"
	"log/slog"
	"os"
)

func NewLoggerWithContext(ctx context.Context) *slog.Logger {

	logType := ctx.Value("logger")
	logLevel := ctx.Value("debug")
	loglvl := new(slog.LevelVar)

	switch logLevel {
	case "warn":
		loglvl.Set(slog.LevelWarn)
	case "error":
		loglvl.Set(slog.LevelError)
	case "debug":
		loglvl.Set(slog.LevelDebug)
	default:
		loglvl.Set(slog.LevelInfo)
	}

	opts := &slog.HandlerOptions{
		AddSource: false,
		Level:     loglvl,
	}

	switch logType {
	case "text":
		handler := slog.NewTextHandler(os.Stdout, opts)
		return slog.New(handler)
	case "json":
		handler := slog.NewJSONHandler(os.Stdout, opts)
		return slog.New(handler)
	default:
		handler := slog.NewTextHandler(os.Stdout, opts)
		return slog.New(handler)
	}

}
