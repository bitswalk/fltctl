package checks

import (
	"fmt"
	"os"

	"github.com/bitswalk/fltctl/internal/logs"
)

var logger = logs.SetLogger()

func Check(args string, uri string) interface{} {

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
