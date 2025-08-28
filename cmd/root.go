package cmd

import (
	"os"
	"strings"

	"github.com/mtsfy/umm/internal/umm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "umm [question]",
	Short:   "Command-Line AI assistant",
	Long:    "A command-line AI assistant tool that helps answer questions using natural language.",
	Example: "umm curl a website but print the headers only",
	Run: func(cmd *cobra.Command, args []string) {
		run, err := cmd.Flags().GetBool("run")
		if err != nil {
			panic(err)
		}

		if run {
			umm.Execute()
			return
		}

		if len(args) == 0 {
			cmd.Help()
			return
		}

		q := strings.Join(args[0:], " ")
		umm.Query(q)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("run", "r", false, "run the last suggested command")
}
