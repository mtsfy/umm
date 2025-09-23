package cmd

import (
	"os"
	"strings"

	"github.com/mtsfy/umm/internal/umm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "umm [question]",
	Short: "AI-powered CLI assistant",
	Long:  "An AI-powered CLI assistant that converts natural language questions into executable commands with explanations.",
	Example: `  # Get a command to list files in a tree structure
  umm list all files like a tree
  
  # Find out how to search within files
  umm search for text in files recursively
  
  # Run the last suggested command
  umm --run`,
	Args: cobra.ArbitraryArgs,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
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
