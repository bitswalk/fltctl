package payload

import (
	"errors"
	"log/slog"
)

func Check(url string) (string, error) {
	if url == "" {
		return "", errors.New("empty url")
	}
	
	// Retrieving image hash if image exist before returning the check state.
	return url, nil
}

func Locate(uri string) (string, error) {
	if uri == "" {
		return "", errors.New("no image name was provided")
	}

	return uri, nil
}

// ref: https://dev.to/douglasmakey/optimizing-large-file-transfers-in-linux-with-go-an-exploration-of-tcp-and-syscall-15eo
func Retrieve(url string) (string, error) {
	slog.Info("Retrieving image remotely,", "url", url)

	if url == "" {
		return "", errors.New("empty url")
	}

	return url, nil
	// // Check local cache first
	// // If file exist don't do anything except telling user.
	// if _, err := os.Stat(os.TempDir() + uri); err == nil {
	// 	slog.Info("Image already exist locally.", "uri", uri, "path", os.TempDir())
	// } else {
	// 	// When file doesn't exist, try to create the local cache path and file.
	// 	// Verify ressource availability on remote endpoint before calling for the ressource.
	// 	slog.Info("Image doesn't exist locally, we'll try to download it.", "uri", uri, "path", os.TempDir())
	// 	slog.Info("Checking upstream ressource availability.", "url", url)
	// }

	// // Check the release version we download if we request for latest current release.
	// if version == "current" {
	// 	versionValue := helpers.GetVersion(ctx, endpoint)
	// 	slog.Info("Retrieving latest current release.", "version", versionValue["FLATCAR_VERSION"])
	// }

	// resp, err := http.Head(url)
	// if err != nil {
	// 	slog.Error("An error occured on endpoint availability request.", "error", err)
	// 	os.Exit(128)
	// }

	// if resp.StatusCode != 200 {
	// 	slog.Error("Remote endpoint didn't answered properly.", "url", url)
	// 	os.Exit(128)
	// }

	// // Trigger the remote url call
	// slog.Info("Retrieving upstream ressource:", "url", url, "Content-length", resp.Header.Get("Content-length"))
	// resp, err = http.Get(url)
	// if err != nil {
	// 	slog.Error("An error occured while trying to contact endpoint.", "url", url, "error", err)
	// 	os.Exit(128)
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	slog.Warn("An HTTP Error occured.", "url", url, "error", err)
	// }
	// // Store the resulting data onto local cache file path location.
	// slog.Info("Writting image locally", "path", os.TempDir(), "uri", uri)
	// file, err := os.Create(os.TempDir() + "/" + uri)
	// if err != nil {
	// 	slog.Error("An error occured creating image on local filesystem.", "url", url, "path", os.TempDir(), "uri", uri, "error", err)
	// 	os.Exit(128)
	// }
	// defer file.Close()

	// _, err = io.Copy(file, resp.Body)
	// if err != nil {
	// 	slog.Error("An error occured writting image locally.", "url", url, "path", os.TempDir(), "uri", uri, "error", err)
	// 	os.Exit(128)
	// }
	// slog.Info("Image successfully retrieved locally!", "path", os.TempDir(), "uri", uri)
}
