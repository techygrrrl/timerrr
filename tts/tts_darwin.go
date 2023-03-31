package tts


import (
    "math/rand"
    "os/exec"
)



func TtsCommand(message string) *(exec.Cmd) {
	voices := []string{"daniel", "samantha", "rishi", "veena", "moira", "fiona", "tessa"}
	voiceIndex := rand.Intn(len(voices))
	voice := voices[voiceIndex]

	return exec.Command("say", "-v", voice, message)
}
