package cmd

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/techygrrrl/timerrr/models"
)

var timerDuration time.Duration
var sayMessage string
var timerName string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a timerrr!",

	Long: `You can start a timer with the flag --minutes 2 (-m 2) and/or --seconds 30 (-s 30) to start a timer for 2 minutes and 30 seconds.

If both are omitted, a 30 second timer will be started.`,

	Run: func(cmd *cobra.Command, args []string) {
		defaultTimer := models.CreateTimer(timerDuration, timerName, sayMessage)
		program := tea.NewProgram(defaultTimer)

		_, err := program.Run()
		if err != nil {
			fmt.Printf("There has been error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	// TODO: Voice customization
	startCmd.Flags().DurationVarP(&timerDuration, "duration", "d", 30*time.Second, "Duration to run the timer")
	startCmd.Flags().StringVarP(&timerName, "name", "n", "My Timerrr", "Message to speak after completed")
	startCmd.Flags().StringVar(&sayMessage, "say", "", "Message to speak after completed")

	rootCmd.AddCommand(startCmd)
}
