package cmd

import (
	"fmt"
	"os"

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
			fmt.Println(args[0:])
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
