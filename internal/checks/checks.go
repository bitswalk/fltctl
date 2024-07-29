package checks

import (
	"context"
	"fmt"
	"os"

	"github.com/bitswalk/fltctl/internal/logs"
)

func Check(ctx context.Context, args string, uri string) interface{} {

	logger := logs.NewLoggerWithContext(ctx)
	switch args {
	case "cache":
		var tmp = os.TempDir()
		if fileInfo, err := os.Stat(tmp + uri); err == nil {
			logger.Info("Image already exist locally", "uri", uri, "path", tmp)
			return fileInfo
		} else {
			logger.Info("Image doesn't exist locally", "uri", uri, "path", tmp)
			return 0
		}
	default:
		fmt.Println("Nothing to check, bye bye.")
		logger.Warn("Nothing to match.")
	}

	return 0
}
