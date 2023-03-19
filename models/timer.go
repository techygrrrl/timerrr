package models

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TimerModel struct {
	duration time.Duration // The millis for the timer
	name     string        // The name of the timer
	// endMessage string // The message to display and say at the end of the timer
	// colour string // Hex representation of the timer
}

func (m TimerModel) Init() tea.Cmd {
	return nil
}

func DefaultTimer() TimerModel {
	return TimerModel{
		duration: time.Second * 30,
		name:     "Timerrr",
	}
}

func (m TimerModel) View() string {
	// TODO: format duration better
	// Progress bar??????????
	return fmt.Sprintf("[%s]: %s", m.name, m.duration)
}

func (m TimerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
