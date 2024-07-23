package checks

import (
	"os"

	"github.com/bitswalk/fltctl/internal/logs"
)

var logger = logs.SetLogger()

func Check(args string, uri string) interface{} {

	switch args {
	case "cache":
		var tmp = os.TempDir()
		if fileInfo, err := os.Stat(tmp + uri); err == nil {
			logger.Info("Image already exist locally", "image", uri, "path", tmp)
			return fileInfo
		} else {
			logger.Info("Image doesn't exist locally", "image", uri, "path", tmp)
			return 0
		}
	}

	return 0
}
