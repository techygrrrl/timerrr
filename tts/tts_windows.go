package tts

import (
	"fmt"
	"os/exec"
)

func TtsCommand(message string) *exec.Cmd {
	speakCmd := fmt.Sprintf("(New-Object System.Speech.Synthesis.SpeechSynthesizer).Speak('%s');", message)
	return exec.Command("PowerShell", "-Command", "Add-Type", "-AssemblyName", "System.Speech;", speakCmd)
}
