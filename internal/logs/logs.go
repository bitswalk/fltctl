package logs

import (
	"log/slog"
	"os"
)

func init() {
	logType, _ := rootCmd.Flags().GetString("log")

	switch logType {
	case "text":
		handler := slog.NewTextHandler(os.Stdout, nil)
		slog.New(handler)
	case "json":
		handler := slog.NewJSONHandler(os.Stdout, nil)
		slog.New(handler)
	}

}
