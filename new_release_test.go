package new_release

import (
	"testing"

	"github.com/tj/assert"
)

func TestCheckConnection(t *testing.T) {
	instance := checkConnection()
	assert.True(t, instance)
}

func TestConvertURL(t *testing.T) {
	instance1 := convertURL("https://github.com/repos/Matt-Gleich/nuke")
	assert.Equal(t, "https://api.github.com/repos/Matt-Gleich/nuke/releases/latest", instance1)

	instance2 := convertURL("https://github.com/repos/Matt-Gleich/nuke/")
	assert.Equal(t, "https://api.github.com/repos/Matt-Gleich/nuke/releases/latest", instance2)
}
