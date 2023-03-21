package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserConfigDirOnTechygrrrlsMac(t *testing.T) {
	path, err := os.UserConfigDir()

	assert.Nil(t, err)
	assert.Equal(t, "/Users/techygrrrl/Library/Application Support", path)
}