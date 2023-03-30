package tts

import "fmt"

func Speak(message string) string {
	return fmt.Sprintf("Hello, from Linux! - %s", message)
}
