package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/techygrrrl/timerrr/main/utils"
)

var (
	focusedStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#EF15BF"))
	blurredButtonStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#15AEEF"))
	cursorStyle        = focusedStyle.Copy()
	defaultStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredButtonStyle.Render("Submit"))
)

type AddModel struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

func (m AddModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AddModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				// Validate the duration
				maybeDuration := m.inputs[1].Value()
				duration, err := time.ParseDuration(maybeDuration)
				if err != nil {
					log.Fatal(err)
				}

				// Add the new timer
				err = utils.AddTimer(utils.SavedTimer{
					Name:        m.inputs[0].Value(),
					Duration:    duration,
					DoneMessage: m.inputs[2].Value(),
				})
				if err != nil {
					log.Fatal(err)
				}

				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}

				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = defaultStyle
				m.inputs[i].TextStyle = defaultStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m AddModel) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}

// TODO: Fix warning?
//		Struct AddModel has methods on both value and pointer receivers. Such usage is not recommended by the Go Documentation.
func (m *AddModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func createAddModel() AddModel {
	model := AddModel{
		focusIndex: 0,
		inputs:     make([]textinput.Model, 3),
		cursorMode: cursor.CursorBlink,
	}

	var input textinput.Model
	for i := range model.inputs {
		input = textinput.New()
		input.CursorStyle = cursorStyle

		switch i {
		case 0:
			input.Placeholder = "Timer name"
			input.Focus()
			input.PromptStyle = focusedStyle
			input.TextStyle = focusedStyle
			input.CharLimit = 15

		case 1:
			input.Placeholder = "Duration, e.g. 5m30s"
			input.CharLimit = 7

		case 2:
			input.Placeholder = "Text-to-speech done message"
			input.CharLimit = 300
		}

		model.inputs[i] = input
	}

	return model
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new timer",
	Run: func(cmd *cobra.Command, args []string) {
		model := createAddModel()

		_, err := tea.NewProgram(model).Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
