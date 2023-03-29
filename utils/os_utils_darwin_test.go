package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMac(t *testing.T) {
	fmt.Printf("is Mac? %t", IsMac())
	assert.True(t, IsMac())
}

func TestUserConfigDirMac(t *testing.T) {
	path, err := os.UserConfigDir()

	t.Logf("🧵 Config dir: %s", path)

	assert.Nil(t, err)
	assert.Equal(t, "/Users/runner/Library/Application Support", path)
}

func TestUserConfigFilePathMac(t *testing.T) {
	path, err := ConfigFilePath()

	t.Logf("🧵 Config file path: %s", path)

	assert.Nil(t, err)
	assert.Equal(t, "/Users/runner/Library/Application Support/timerrr/timers.json", path)
}
