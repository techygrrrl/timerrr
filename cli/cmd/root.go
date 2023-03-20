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

type tableModel struct {
	table table.Model
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#15AEEF"))

func (m tableModel) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func (m tableModel) Init() tea.Cmd { return nil }

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
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
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
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

	// TODO: Do the TUI stuff here
	Run: func(cmd *cobra.Command, args []string) {
		staticTimers := []models.TimerModel{
			models.CreateTimer(time.Second*2, "My First timer", ""),
			models.CreateTimer(time.Minute*5, "Stand Up", "You can sit down now"),
			models.CreateTimer(time.Second*30+(time.Minute*4), "Tea Timerrr", "Your tea is ready"),
		}

		columns := []table.Column{
			{Title: "Duration", Width: 10},
			{Title: "Name", Width: 60},
		}

		var rows []table.Row
		for _, timer := range staticTimers {
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

		// TODO: Full screen table

		model := tableModel{timerTable}
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
