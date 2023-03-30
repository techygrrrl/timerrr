package tts

func TestSpeak(t *testing.T) {
	result := Speak("test")
	t.Logf("🟣 Speak result: %s", result)
	assert.Equal(t, "Hello, from Linux! - test", result)
}
