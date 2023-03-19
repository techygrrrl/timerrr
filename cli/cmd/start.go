package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var minutes int
var seconds int

// TODO: Add ability to configure default timer
// TODO: Add ability to save timers (TUI)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a timerrr!",

	Long: `You can start a timer with the flag --minutes 2 (-m 2) and/or --seconds 30 (-s 30) to start a timer for 2 minutes and 30 seconds.

If both are omitted, a 30 second timer will be started.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("start called with minutes = %d and seconds = %d \n", minutes, seconds)

		// TODO: bubbletea things
	},
}

func init() {
	startCmd.Flags().IntVarP(&minutes, "minutes", "m", 0, "Minutes")
	startCmd.Flags().IntVarP(&seconds, "seconds", "s", 30, "Seconds")

	rootCmd.AddCommand(startCmd)
}
