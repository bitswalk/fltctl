package logs

import (
	"log/slog"
	"os"
)

func SetLogger() *slog.Logger {

	// logType, _ := cmd.Flags().GetString("log")
	logType := "text"

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
