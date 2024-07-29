package logs

import (
	"context"
	"log/slog"
	"os"
)

func NewLoggerWithContext(ctx context.Context) *slog.Logger {

	logType := ctx.Value("logger")

	opts := &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelInfo,
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
