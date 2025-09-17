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
		page, err := cmd.Flags().GetInt("page")
		if err != nil {
			panic(err)
		}

		size, err := cmd.Flags().GetInt("size")
		if err != nil {
			panic(err)
		}

		search, err := cmd.Flags().GetString("search")
		if err != nil {
			panic(err)
		}

		if search != "" {
			history.FilterHistory(search)
			return
		}

		if size != -1 || page != -1 {
			if page == -1 {
				page = 1
			}
			if size == -1 {
				size = 10
			}

			history.PaginatedHistory(page, size)
			return
		}

		history.AllHistory()
	},
}
