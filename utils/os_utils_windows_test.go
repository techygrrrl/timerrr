package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWindows(t *testing.T) {
	fmt.Printf("is Windows? %t", IsWindows())
	assert.True(t, IsWindows())
}

func TestUserConfigDirWindows(t *testing.T) {
	path, err := os.UserConfigDir()

	t.Logf("🧵 Config dir: %s", path)

	assert.Nil(t, err)
	assert.Equal(t, "C:\\Users\\runneradmin\\AppData\\Roaming", path)
}

func TestUserConfigFilePathWindows(t *testing.T) {
	path, err := ConfigFilePath()

	t.Logf("🧵 Config file path: %s", path)

	assert.Nil(t, err)
	assert.Equal(t, "C:\\Users\\runneradmin\\AppData\\Roaming\\timerrr\\timers.json", path)
}
