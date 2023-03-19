package models

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	winHeight, winWidth int
	quitKeys            = key.NewBinding(key.WithKeys("esc", "q"))
	intKeys             = key.NewBinding(key.WithKeys("ctrl+c"))
	boldStyle           = lipgloss.NewStyle().Bold(true)
	italicStyle         = lipgloss.NewStyle().Italic(true)
)

const (
	padding  = 2
	maxWidth = 80
)

type TimerModel struct {
	timer           timer.Model    // bubble component
	progress        progress.Model // bubble component
	passed          time.Duration
	duration        time.Duration
	name            string // The name of the timer
	start           time.Time
	finishedMessage string
	altScreen       bool // Full screen
	quitting        bool
	interrupting    bool
	// TODO: Customize colour?
	// colour string // Hex representation of the timer
}

func (m TimerModel) Init() tea.Cmd {
	return m.timer.Init()
}

func CreateTimer(duration time.Duration, name string, message string) TimerModel {
	return TimerModel{
		timer:           timer.NewWithInterval(duration, time.Second),
		progress:        progress.New(progress.WithScaledGradient("#EF15BF", "#7515EF")),
		passed:          0,
		duration:        duration,
		name:            name,
		start:           time.Now(),
		finishedMessage: message,
		altScreen:       true,
		quitting:        false,
		interrupting:    false,
	}
}

type speakFinishedMsg struct{ err error }

func speak(m TimerModel) tea.Cmd {
	message := m.finishedMessage
	if message == "" {
		message = fmt.Sprintf("The timer %s has completed", m.name)
	}

	sayCmd := exec.Command("say", "-v", "daniel", message)

	return tea.ExecProcess(sayCmd, func(err error) tea.Msg {
		fmt.Println("Error: ", err)
		return speakFinishedMsg{err}
	})
}

func (m TimerModel) View() string {
	if m.quitting || m.interrupting {
		// Returning a line break removes the progress indicator - Do nothing on quit for now
		//return "\n"
	}

	result := boldStyle.Render(m.start.Format(time.Kitchen))
	if m.name != "" {
		result += ": " + italicStyle.Render(m.name)
	}
	result += " - " + boldStyle.Render(m.timer.View()) + "\n" + m.progress.View()
	if m.altScreen {
		textWidth, textHeight := lipgloss.Size(result)
		return lipgloss.NewStyle().Margin((winHeight-textHeight)/2, (winWidth-textWidth)/2).Render(result)
	}

	return result
}

func (m TimerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmds []tea.Cmd
		var cmd tea.Cmd

		m.passed += m.timer.Interval

		pct := m.passed.Milliseconds() * 100 / m.duration.Milliseconds()
		cmds = append(cmds, m.progress.SetPercent(float64(pct)/100))

		m.timer, cmd = m.timer.Update(msg)
		cmds = append(cmds, cmd)

		if msg.Timeout {
			cmds = append(cmds, speak(m))
		}

		return m, tea.Batch(cmds...)

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		winHeight, winWidth = msg.Height, msg.Width
		if !m.altScreen && m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case tea.KeyMsg:
		if key.Matches(msg, quitKeys) {
			m.quitting = true
			return m, tea.Quit
		}
		if key.Matches(msg, intKeys) {
			m.interrupting = true
			return m, tea.Quit
		}

	case speakFinishedMsg:
		m.quitting = true
		if msg.err != nil {
			fmt.Println("Error: ", msg.err)
		}
		return m, tea.Quit
	}

	return m, nil
}
