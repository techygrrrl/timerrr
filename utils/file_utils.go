package utils

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/techygrrrl/timerrr/models"
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

// LoadTimersFromFile
// Load timers from file. If any error occurs, ignore it and return an empty data structure
func LoadTimersFromFile() []models.SavedTimer {
	configPath, err := ConfigFilePath()
	if err != nil {
		return []models.SavedTimer{}
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
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
	configPath, err := ConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(timers, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// ConfigFilePath
// Returns the desired configuration path for the file.
// on Mac, it's the Application Support, in Linux it's
func ConfigFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	configDir = filepath.Join(configDir, "timerrr")

	if _, err := os.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return filepath.Join(configDir, "timers.json"), nil
}
