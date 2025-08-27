package cmd

import (
	"os"
	"strings"

	"github.com/mtsfy/umm/internal/umm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "umm [question]",
	Short: "Command-Line AI assistant",
	Long:  "A command-line AI assistant tool that helps answer questions using natural language.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			q := strings.Join(args[0:], " ")
			umm.Query(q)
		}

	},
	Example: "umm curl a website but print the headers only",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
