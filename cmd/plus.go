package cmd

import (
	"fmt"
	"strings"

	"github.com/mtsfy/umm/internal/ai"
	"github.com/mtsfy/umm/internal/history"
	"github.com/spf13/cobra"
)

var plusCmd = &cobra.Command{
	Use:     "+ [question]",
	Short:   "Ask a follow-up question",
	Long:    "A follow-up query that builds on the context of your most recent question to expand your previous query.",
	Args:    cobra.ArbitraryArgs,
	Example: "umm + what about with curl?",
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
