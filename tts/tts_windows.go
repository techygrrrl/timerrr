package tts


func TtsCommand(message string) *exec.Cmd {
    return exec.Command("echo", message)
}
