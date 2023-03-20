package utils

import (
	"encoding/json"
	"os"

	"github.com/techygrrrl/timerrr/main/models"
)

const (
	timersJsonFileName = "timers.json"
)

func AddTimer(timer models.SavedTimer) error {
	timers := LoadTimersFromFile()
	timers = append(timers, timer)

	return persistTimers(timers)
}

func RemoveTimerAtIndex(index int) error {
	timers := LoadTimersFromFile()

	// Learn: https://go.dev/play/p/M-7bwMAROWB
	timers = append(timers[:index], timers[index+1:]...)

	return persistTimers(timers)
}

func LoadTimersFromFile() []models.SavedTimer {
	data, err := os.ReadFile(timersJsonFileName)
	if err != nil {
		// Empty file, return empty array
		return []models.SavedTimer{}
	}

	var timers []models.SavedTimer
	err = json.Unmarshal(data, &timers)
	if err != nil {
		return []models.SavedTimer{}
	}

	return timers
}

func persistTimers(timers []models.SavedTimer) error {
	data, err := json.MarshalIndent(timers, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(timersJsonFileName, data, 0644)
}
