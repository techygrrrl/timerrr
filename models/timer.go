package models

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	name                string
	altScreen           bool
	winHeight, winWidth int
	version             = "dev"
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
	timer        timer.Model    // bubble component
	progress     progress.Model // bubble component
	passed       time.Duration
	duration     time.Duration
	name         string // The name of the timer
	start        time.Time
	altScreen    bool
	quitting     bool
	interrupting bool
	// endMessage string // The message to display and say at the end of the timer
	// colour string // Hex representation of the timer
}

func (m TimerModel) Init() tea.Cmd {
	return m.timer.Init()
}

func CreateTimer(minutes int, seconds int, running bool) TimerModel {
	minutesDuration := time.Duration(minutes) * time.Minute
	secondsDuration := time.Duration(seconds) * time.Second

	duration := minutesDuration + secondsDuration

	return TimerModel{
		timer:        timer.NewWithInterval(duration, time.Second),
		progress:     progress.New(progress.WithDefaultGradient()),
		passed:       0,
		duration:     duration,
		name:         "Timerrr",
		start:        time.Now(),
		altScreen:    altScreen,
		quitting:     false,
		interrupting: false,
	}
}

func (m TimerModel) View() string {
	if m.quitting || m.interrupting {
		return "\n"
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
	// TODO: format duration better
	// Progress bar??????????
	//return fmt.Sprintf("[%s]: %s ... %s", m.name, m.duration, m.passed)
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

	case timer.TimeoutMsg:
		m.quitting = true
		return m, tea.Quit

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
	}

	return m, nil
}
