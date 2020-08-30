package new_release

import (
	"testing"

	"github.com/tj/assert"
)

func TestCheckConnection(t *testing.T) {
	instance := checkConnection()
	assert.True(t, instance)
}
