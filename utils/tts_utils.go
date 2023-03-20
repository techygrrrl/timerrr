package utils

import (
	"math/rand"
)

func GetRandomTTSVoice() string {
	voices := []string{"daniel", "samantha", "rishi", "veena", "moira", "fiona", "tessa"}
	voiceIndex := rand.Intn(len(voices))
	return voices[voiceIndex]
}
