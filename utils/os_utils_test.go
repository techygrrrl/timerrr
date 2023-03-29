package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMac(t *testing.T) {
	assert.True(t, IsMac())
}
