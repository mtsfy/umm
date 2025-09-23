package cmd

import (
	"fmt"
	"strings"

	"github.com/mtsfy/umm/internal/ai"
	"github.com/mtsfy/umm/internal/history"
	"github.com/spf13/cobra"
)

var plusCmd = &cobra.Command{
	Use:   "+ [follow-up question]",
	Short: "Ask a follow-up question based on your previous query",
	Long:  "Ask a follow-up question that builds on your most recent interaction, allowing you to refine or expand without starting from scratch.",
	Args:  cobra.ArbitraryArgs,
	Example: `  # Follow up on your last query with additional context
  umm + what about with curl?
  
  # Refine your previous question
  umm + but for hidden files only
  
  # Ask for alternatives to the previous suggestion
  umm + is there a simpler way?`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		newQuery := strings.Join(args, " ")

		lastInteraction := history.GetLatest()
		lastQuery := lastInteraction.UserInput

		if lastQuery == "" {
			fmt.Println("no last query found to follow-up")
			return
		}

		err := ai.FollowUp(lastInteraction, newQuery)
		if err != nil {
			panic(err)
		}
	},
}
