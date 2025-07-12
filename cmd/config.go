package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"smartcommit/config"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "View or change smartcommit configuration",
	Long: `Manage smartcommit configuration

Examples:
  smartcommit config edit       # Edit the system prompt (tone/style)
  smartcommit config show       # View current configuration`,
}

var EditConfigCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit system prompt in your preferred editor (e.g., vim, nano, code)",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadOrDefault()

		tmpFile, err := os.CreateTemp("", "smartcommit-prompt-*.txt")
		if err != nil {
			fmt.Println("❌ Failed to create temp file:", err)
			return
		}
		defer os.Remove(tmpFile.Name())

		// Write current prompt to temp file
		_, _ = tmpFile.WriteString(cfg.SystemPrompt)
		tmpFile.Close()

		// Open in default editor
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "vim"
		}
		cmdEdit := exec.Command(editor, tmpFile.Name())
		cmdEdit.Stdin = os.Stdin
		cmdEdit.Stdout = os.Stdout
		cmdEdit.Stderr = os.Stderr

		if err := cmdEdit.Run(); err != nil {
			fmt.Println("❌ Failed to open editor:", err)
			return
		}

		// Read edited prompt
		editedBytes, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			fmt.Println("❌ Failed to read edited prompt:", err)
			return
		}

		edited := string(editedBytes)
		if edited == "" {
			fmt.Println("⚠️ Prompt left empty. Aborting.")
			return
		}

		cfg.SystemPrompt = edited
		if err := config.Save(cfg); err != nil {
			fmt.Println("❌ Failed to save config:", err)
		} else {
			fmt.Println("✅ System prompt updated successfully.")
		}
	},
}


var ShowConfigCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current config",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadOrDefault()
		fmt.Printf("\nSystem Prompt: %s\n", cfg.SystemPrompt)
		fmt.Printf("Provider: %s\n", cfg.Provider)
		fmt.Printf("Model: %s\n", cfg.Model)
		if cfg.APIKey != "" {
			fmt.Println("API Key: [set]")
		} else {
			fmt.Println("API Key: [not set]")
		}
	},
}

func init() {
	ConfigCmd.AddCommand(EditConfigCmd)
	ConfigCmd.AddCommand(ShowConfigCmd)
} 
