package models

import (
	"fmt"
	"math/rand"
	"os/exec"
	"runtime"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	winHeight   int
	winWidth    int
	quitKeys    = key.NewBinding(key.WithKeys("esc", "q"))
	intKeys     = key.NewBinding(key.WithKeys("ctrl+c"))
	boldStyle   = lipgloss.NewStyle().Bold(true)
	italicStyle = lipgloss.NewStyle().Italic(true)
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
}

func (m TimerModel) Init() tea.Cmd {
	m.progress.Width = winWidth
	return m.timer.Init()
}

func (m TimerModel) TableRowDisplay() []string {
	return []string{
		m.duration.String(),
		m.name,
	}
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

type SavedTimer struct {
	Name        string        `json:"name"`
	Duration    time.Duration `json:"duration"`
	DoneMessage string        `json:"done_message"`
}

func (timer SavedTimer) AsTimerModel() TimerModel {
	return CreateTimer(timer.Duration, timer.Name, timer.DoneMessage)
}

// region TTS

type speakFinishedMsg struct{ err error }

func speak(m TimerModel) tea.Cmd {
	message := m.finishedMessage
	if message == "" {
		message = fmt.Sprintf("The timer %s has completed", m.name)
	}

	sayCmd := ttsCommandForOS(message)

	return tea.ExecProcess(sayCmd, func(err error) tea.Msg {
		fmt.Println("Error: ", err)
		return speakFinishedMsg{err}
	})
}

func ttsCommandForOS(message string) *exec.Cmd {
	switch runtime.GOOS {
	case "darwin":
		return ttsCommandMac(message)
	case "windows":
		return ttsCommandWindows(message)
	case "linux":
		return ttsCommandLinux(message)
	default:
		// Fallback echo command
		return exec.Command("echo", message)
	}
}

func ttsCommandMac(message string) *exec.Cmd {
	voices := []string{"daniel", "samantha", "rishi", "veena", "moira", "fiona", "tessa"}
	voiceIndex := rand.Intn(len(voices))
	voice := voices[voiceIndex]

	return exec.Command("say", "-v", voice, message)
}

// TODO: Implement espeak - https://github.com/techygrrrl/timerrr/issues/1
// TODO: Implement mimic3 - https://github.com/techygrrrl/timerrr/issues/2
func ttsCommandLinux(message string) *exec.Cmd {
	return exec.Command("echo", message)
}

// TODO: Implement - https://github.com/techygrrrl/timerrr/issues/3
func ttsCommandWindows(message string) *exec.Cmd {
	return exec.Command("echo", message)
}

// endregion TTS

func (m TimerModel) View() string {
	if m.quitting || m.interrupting {
		// Returning a line break removes the progress indicator - Do nothing on quit for now
		//return "\n"
	}

	result := boldStyle.Render(m.start.Format(time.Kitchen))
	if m.name != "" {
		result += ": " + italicStyle.Render(m.name)
	}

	// Sets the width of the progress bar to ensure a consistent size when used by another command
	m.progress.Width = winWidth - padding*2 - 4

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
