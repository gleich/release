package new_release

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Check for an update. Takes in the current version and GitHub repo URL.
// Returns true or false if there is an update or not as well as the version
// value. Will return false if there is no network connection.
func Check(localVersion string, repoURL string) (bool, string, error) {
	hasConnection := checkConnection()
	if !hasConnection {
		return false, "", nil
	}
	requestURL := convertURL(repoURL)
	currentVersion, err := getVersion(requestURL)
	if err != nil {
		return false, "", err
	}
	if localVersion != currentVersion {
		return true, currentVersion, nil
	}
	return false, currentVersion, nil
}

// Check for a network connection
func checkConnection() bool {
	resp, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil || resp.StatusCode != 204 {
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

// Make the actual get request to get the version
func getVersion(requestURL string) (string, error) {
	resp, err := http.Get(requestURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	version := fmt.Sprintf("%v", data["tag_name"])
	if version == "" {
		return "", errors.New("Version number for repo is blank")
	}
	return version, nil
}
