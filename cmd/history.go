package cmd

import (
	"github.com/mtsfy/umm/internal/history"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:     "history",
	Short:   "Display and search your past interactions.",
	Long:    "View your command history and search through previous AI interactions using keywords. This feature helps you quickly locate past queries and responses.",
	Example: "umm history --search curl",
	Run: func(cmd *cobra.Command, args []string) {
		history.ViewHistory() // simple history view
	},
}
