package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/techygrrrl/timerrr/models"
	"github.com/techygrrrl/timerrr/utils"
)

var (
	winHeight int
	winWidth  int
)

type tableModel struct {
	timers []models.TimerModel
	table  table.Model
	timer  *models.TimerModel
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#15AEEF"))

func (m tableModel) View() string {
	if m.timer != nil {
		fmt.Println("Returning timer view")
		return m.timer.View()
	}

	result := baseStyle.Render(m.table.View()) + "\n"
	textWidth, textHeight := lipgloss.Size(result)

	return lipgloss.NewStyle().Margin((winHeight-textHeight)/2, (winWidth-textWidth)/2).Render(result) + "\n"
}

func (m tableModel) Init() tea.Cmd { return nil }

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		winHeight, winWidth = msg.Height, msg.Width

		if m.timer != nil {
			m.timer.Update(msg)
		}

		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit

		case "del", "delete", "backspace":
			index := m.table.Cursor()

			m.table.Blur()

			// rw_grim: PomoYoloTimerrrs - By techygrrrl
			err := utils.RemoveTimerAtIndex(index)
			if err != nil {
				log.Fatal(err)
			}

			timers := utils.LoadTimersFromFile()
			var timerModels []models.TimerModel

			for _, element := range timers {
				timerModels = append(timerModels, element.AsTimerModel())
			}

			timerTable := createTableForTimers(timerModels)
			m.table = timerTable

			return m, nil

		case "enter":
			index := m.table.Cursor()
			selected := m.timers[index]
			m.timer = &m.timers[index]

			cmd := selected.Init()

			// This helps to center it vertically
			m.timer.Update(tea.WindowSizeMsg{
				Width:  winWidth,
				Height: winHeight,
			})

			// This successfully returns a functioning timer
			return selected, cmd
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

var rootCmd = &cobra.Command{
	Use: "timerrr",

	Run: func(cmd *cobra.Command, args []string) {
		timers := utils.LoadTimersFromFile()

		var timerModels []models.TimerModel
		for _, element := range timers {
			timerModels = append(timerModels, element.AsTimerModel())
		}

		timerTable := createTableForTimers(timerModels)

		model := tableModel{
			table:  timerTable,
			timers: timerModels,
		}
		_, err := tea.NewProgram(model).Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func createTableForTimers(timers []models.TimerModel) table.Model {
	columns := []table.Column{
		{Title: "Duration", Width: 10},
		{Title: "Name", Width: 60},
	}

	var rows []table.Row
	for _, timer := range timers {
		rows = append(rows, timer.TableRowDisplay())
	}

	timerTable := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	styles := table.DefaultStyles()
	styles.Header = styles.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#15AEEF")).
		BorderBottom(true).
		Bold(true)

	styles.Selected = styles.Selected.
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#EF15BF")).
		Bold(false)

	timerTable.SetStyles(styles)

	return timerTable
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
}
