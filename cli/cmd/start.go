package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/techygrrrl/timerrr/main/models"
)

var minutes int
var seconds int
var sayMessage string
var timerName string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a timerrr!",

	Long: `You can start a timer with the flag --minutes 2 (-m 2) and/or --seconds 30 (-s 30) to start a timer for 2 minutes and 30 seconds.

If both are omitted, a 30 second timer will be started.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: collect message from flags
		//fmt.Printf("start called with minutes = %d and seconds = %d \n", minutes, seconds)

		// With message
		//defaultTimer := models.CreateTimer(minutes, seconds, "Hello! Your timer is done!")

		// Without message
		defaultTimer := models.CreateTimer(minutes, seconds, timerName, sayMessage)

		program := tea.NewProgram(defaultTimer)

		_, err := program.Run()
		if err != nil {
			fmt.Printf("There has been error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	// TODO: Add ability to configure default timer
	// TODO: Voice customization
	startCmd.Flags().IntVarP(&minutes, "minutes", "m", 0, "Minutes")
	startCmd.Flags().IntVarP(&seconds, "seconds", "s", 30, "Seconds")
	startCmd.Flags().StringVarP(&timerName, "name", "n", "My Timerrr", "Message to speak after completed")
	startCmd.Flags().StringVar(&sayMessage, "say", "", "Message to speak after completed")

	rootCmd.AddCommand(startCmd)
}
