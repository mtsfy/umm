package cmd

import (
	"github.com/mtsfy/umm/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings",
	Long:  "Configure API settings, model preferences, and other options for umm CLI.",
	Example: `  # Interactive setup
  umm config setup
  
  # Show current configuration
  umm config show`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Interactive configuration setup",
	Long:  "Launch an interactive form to configure API key, model preferences, and other settings.",
	Run: func(cmd *cobra.Command, args []string) {
		config.Setup()
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display current configuration",
	Long:  "Show all current configuration settings including API key and model preferences.",
	Run: func(cmd *cobra.Command, args []string) {
		config.Show()
	},
}

func init() {
	configCmd.AddCommand(setupCmd)
	configCmd.AddCommand(showCmd)
}
