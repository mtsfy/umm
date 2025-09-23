package cmd

import (
	"fmt"
	"strings"

	"github.com/mtsfy/umm/internal/history"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Display and manage your past interactions.",
	Long:  "View, search, and delete your command history.",
	Example: `  # Delete all history entries (with confirmation)
  umm history --delete -1

  # Delete a specific history entry by ID
  umm history --delete 2

  # View history with pagination (page 2, 10 entries per page)
  umm history --page 2 --size 10

  # Search for entries containing "curl"
  umm history --search curl

  # Display all history entries (default behavior)
  umm history`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteID, _ := cmd.Flags().GetInt("delete")
		search, _ := cmd.Flags().GetString("search")
		page, _ := cmd.Flags().GetInt("page")
		size, _ := cmd.Flags().GetInt("size")

		if cmd.Flags().Changed("delete") {
			if deleteID == -1 {
				fmt.Print("Are you sure you want to delete all history? (y/n): ")
				var response string
				fmt.Scanln(&response)
				if strings.ToLower(response) == "y" {
					history.DeleteAllHistory()
					fmt.Println("All history deleted.")
				} else {
					fmt.Println("Deletion canceled.")
				}
			} else if deleteID > 0 {
				history.DeleteHistory(deleteID)
				fmt.Printf("Deleted history entry %d.\n", deleteID)
			} else {
				fmt.Printf("Error: Invalid delete ID '%d'. Use -1 to delete all or a positive number for a specific entry.\n", deleteID)
				return
			}
			return
		}

		if search != "" {
			history.FilterHistory(search)
			return
		}

		if cmd.Flags().Changed("page") || cmd.Flags().Changed("size") {
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
