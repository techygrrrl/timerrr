package utils

import (
	"encoding/json"
	"os"
	"time"
)

type SavedTimer struct {
	Name        string        `json:"name"`
	Duration    time.Duration `json:"duration"`
	DoneMessage string        `json:"done_message"`
}

const (
	timersJsonFileName = "timers.json"
)

func AddTimer(timer SavedTimer) error {
	timers := loadTimersFromFile()
	timers = append(timers, timer)

	return persistTimers(timers)
}

func loadTimersFromFile() []SavedTimer {
	data, err := os.ReadFile(timersJsonFileName)
	if err != nil {
		// Empty file, return empty array
		return []SavedTimer{}
	}

	var timers []SavedTimer
	err = json.Unmarshal(data, &timers)
	if err != nil {
		return []SavedTimer{}
	}

	return timers
}

func persistTimers(timers []SavedTimer) error {
	data, err := json.MarshalIndent(timers, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(timersJsonFileName, data, 0644)
}
