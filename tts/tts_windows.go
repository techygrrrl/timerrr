package tts

import (
	"os/exec"
)

func TtsCommand(message string) *exec.Cmd {
	return exec.Command("echo", message)
}
