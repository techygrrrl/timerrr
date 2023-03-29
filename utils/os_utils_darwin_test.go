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

	assert.Nil(t, err)
	assert.Equal(t, "/Users/runner/Library/Application Support", path)
}

func TestUserConfigFilePathMac(t *testing.T) {
	path, err := os.UserConfigDir()

	assert.Nil(t, err)
	assert.Equal(t, "/Users/runner/Library/Application Support/timers.json", path)
}
