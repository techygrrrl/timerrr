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
	home := os.Getenv("HOME")
	if home == "" {
		home = "/home/runner"
	}

	t.Logf("🧵 Config dir: %s", path)

	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%s/.config", home), path)
}

func TestUserConfigFilePathLinux(t *testing.T) {
	path, err := ConfigFilePath()
	home := os.Getenv("HOME")
	if home == "" {
		home = "/home/runner"
	}

	t.Logf("🧵 Config file path: %s", path)

	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%s/.config/timerrr/timers.json", home), path)
}
