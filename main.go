// Entry point: main.go
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
		Short: "Generate Git commit messages using AI",
		Long: `smartcommit is a CLI tool that uses local or remote LLMs (e.g. Ollama, OpenAI) to generate commit messages
based on staged Git changes. It supports interactive flow, multiple backends, and prompt customization.`,
	}

	rootCmd.AddCommand(cmd.GenerateCmd)
	rootCmd.AddCommand(cmd.ConfigCmd)

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show help message")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}
}
	