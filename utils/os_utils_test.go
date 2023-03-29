package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMac(t *testing.T) {
	fmt.Printf("is Mac? %t", IsMac())
	assert.True(t, IsMac())
}

func TestIsWindows(t *testing.T) {
	fmt.Printf("is Windows? %t", IsWindows())
	assert.True(t, IsWindows())
}

func TestIsLinux(t *testing.T) {
	fmt.Printf("is Linux? %t", IsLinux())
	assert.True(t, IsLinux())
}
