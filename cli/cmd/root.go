package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/techygrrrl/timerrr/main/models"
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
	// TODO: FIX: This doesn't work
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
		return m, nil

	// TODO: Remove if not useful??
	//case timer.TickMsg:
	//	fmt.Println("Tick happened")
	//
	//	if m.timer != nil {
	//		return m.Update(msg)
	//	}

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

		// TODO: Swap to new timer view w/ working functionality
		case "enter":
			//if m.timers == nil {
			//	return m, nil
			//}

			index := m.table.Cursor()
			selected := m.timers[index]
			m.timer = &m.timers[index]

			//timer, cmd := m.timer.Update(timer.TickMsg{})
			//m.timer = &timer

			//return m, cmd
			return selected, selected.Init()
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "timerrr",
	//Short: "⏱ Create timerrrs!",
	//	Long: ``,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Get timers from somewhere
		timers := []models.TimerModel{
			models.CreateTimer(time.Second*2, "My First timer", ""),
			models.CreateTimer(time.Minute*5, "Stand Up", "You can sit down now"),
			models.CreateTimer(time.Second*30+(time.Minute*4), "Tea Timerrr", "Your tea is ready"),
		}

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

		model := tableModel{
			table:  timerTable,
			timers: timers,
		}
		_, err := tea.NewProgram(model).Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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
