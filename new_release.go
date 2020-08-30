package new_release

import (
	"net/http"
	"strings"
)

// Check for an update. Takes in the current version and GitHub repo URL.
// Returns true or false if there is an update or not. Will return false if
// there is no network connection
func Check(currentVersion string, repoURL string) (bool, error) {
	hasConnection := checkConnection()
	if !hasConnection {
		return false, nil
	}
	return true, nil
}

// Check for a network connection
func checkConnection() bool {
	resp, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil || resp.StatusCode != 200 && resp.StatusCode != 204 {
		return false
	}
	return true
}

// Convert repo url to api url
// From: https://github.com/repos/Matt-Gleich/nuke
// To:   https://api.github.com/repos/Matt-Gleich/nuke/releases/latest
func convertURL(repoURL string) string {
	var fixedURL string
	fixedURL = strings.Replace(repoURL, "https://github.com", "https://api.github.com", 1)
	if fixedURL[len(fixedURL)-1:] == "/" {
		fixedURL = fixedURL + "releases/latest"
	} else {
		fixedURL = fixedURL + "/releases/latest"
	}
	return fixedURL
}
