package new_release

import (
	"errors"
	"strings"
	"testing"

	"github.com/tj/assert"
)

func TestCheckConnection(t *testing.T) {
	instance := checkConnection()
	assert.True(t, instance)
}

func TestConvertURL(t *testing.T) {
	instance1 := convertURL("https://github.com/Matt-Gleich/nuke")
	assert.Equal(t, "https://api.github.com/repos/Matt-Gleich/nuke/releases/latest", instance1)

	instance2 := convertURL("https://github.com/Matt-Gleich/nuke/")
	assert.Equal(t, "https://api.github.com/repos/Matt-Gleich/nuke/releases/latest", instance2)
}

func TestGetVersion(t *testing.T) {
	instance, err := getVersion("https://api.github.com/repos/Matt-Gleich/nuke/releases/latest")
	checkTestingErr(t, err)
	if instance[:1] != "v" || !strings.Contains(instance, ".") {
		t.Error("instance looks like this:", instance)
	}
}

func TestCheck(t *testing.T) {
	successInstance, successVersion, successErr := Check("v1.0.0", "https://github.com/Matt-Gleich/nuke")
	checkTestingErr(t, successErr)
	assert.True(t, successInstance)
	if successVersion[:1] != "v" || !strings.Contains(successVersion, ".") {
		t.Error("instance looks like this:", successVersion)
	}

	failedInstance, failedVersion, failedErr := Check("", "https://github.com/repos/Matt-Gleich/nuke")
	assert.False(t, failedInstance)
	assert.Equal(t, "", failedVersion)
	assert.Equal(t, errors.New("Latest release not found for given repo URL"), failedErr)
}

// Check for a error in one line
func checkTestingErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
