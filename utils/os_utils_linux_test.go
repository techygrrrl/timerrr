package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLinux(t *testing.T) {
	fmt.Printf("is Linux? %t", IsLinux())
	assert.True(t, IsLinux())
}

func TestUserConfigDirLinux(t *testing.T) {
	path, err := os.UserConfigDir()

	t.Logf("🧵 Config dir: %s", path)

	assert.Nil(t, err)
	assert.Equal(t, "/home/runner/.config", path)
}

func TestUserConfigFilePathLinux(t *testing.T) {
	path, err := ConfigFilePath()

	t.Logf("🧵 Config file path: %s", path)

	assert.Nil(t, err)
	assert.Equal(t, "/home/runner/.config/timerrr/timers.json", path)
}
