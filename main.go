package main

import (
	"fmt"
	"os"

	"smartcommit/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "smartcommit",
		Short: "AI-powered Git commit message generator",
		Long: `smartcommit is a CLI tool to generate AI-based Git commit messages
from staged changes using local or remote LLMs (e.g., Ollama, OpenAI).

Examples:
  smartcommit generate          Generate a commit message from staged changes
  smartcommit config edit       Edit the system prompt (tone/style)
  smartcommit config show       View current configuration

Run 'smartcommit [command] --help' for detailed command help.`,
	}
	
	// Override Cobra's help output with our Long text
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, _ []string) {
		fmt.Println(cmd.Long)
	})

	// Register subcommands
	rootCmd.AddCommand(cmd.GenerateCmd)
	rootCmd.AddCommand(cmd.ConfigCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}
}
