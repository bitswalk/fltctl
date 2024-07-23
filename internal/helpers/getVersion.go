package helpers

import (
	"bufio"
	"log/slog"
	"net/http"
	"strings"

	"github.com/bitswalk/fltctl/internal/logs"
)

var logger = logs.SetLogger()

func GetVersion(endpoint string) map[string]string {
	var versionInfo = make(map[string]string)
	versionExist, _ := http.Head(endpoint + "version.txt")

	if versionExist.StatusCode != 200 {
		logger.Error("Seems like this endpoint doesn't expose a version file:", slog.Int("return code:", versionExist.StatusCode))
	}

	result, _ := http.Get(endpoint + "version.txt")

	scanner := bufio.NewScanner(result.Body)
	for scanner.Scan() {
		key, value, found := strings.Cut(scanner.Text(), "=")
		if !found {
			logger.Warn("No information to retrieve")
		}
		versionInfo[key] = value
	}

	return versionInfo
}
