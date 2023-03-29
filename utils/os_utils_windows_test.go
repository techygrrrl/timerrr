package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWindows(t *testing.T) {
	fmt.Printf("is Windows? %t", IsWindows())
	assert.True(t, IsWindows())
}

func TestUserConfigDirWindows(t *testing.T) {
	path, err := os.UserConfigDir()

	assert.Nil(t, err)
	assert.Equal(t, "C:\\Users\\runneradmin\\AppData\\Roaming", path)
}

func TestUserConfigFilePathWindows(t *testing.T) {
	path, err := os.UserConfigDir()

	assert.Nil(t, err)
	assert.Equal(t, "C:\\Users\\runneradmin\\AppData\\Roaming\\timerrr\\timers.json", path)
}
