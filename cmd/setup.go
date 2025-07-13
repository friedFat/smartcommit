// File: cmd/setup.go
package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"smartcommit/config"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure LLM provider and settings",
	Long: `Interactive setup for SmartCommit.

This will prompt you to configure:
  • Provider (ollama or http)
  • Model name (e.g. llama3, gpt-3.5-turbo, gemini)
  • API Key (for HTTP-based providers)
  • Base URL (your LLM endpoint)

Need help? Run 'smartcommit --help' or see:
https://github.com/your-repo/smartcommit#readme`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔧 SmartCommit Setup — press Ctrl+C to abort, or run 'smartcommit --help' for guidance.")

		cfg := config.LoadOrDefault()

		// 1) Provider selection
		providers := []string{"ollama", "http"}
		defaultIdx := 0
		for i, p := range providers {
			if p == cfg.Provider {
				defaultIdx = i
				break
			}
		}
		providerSelect := promptui.Select{
			Label:     "Choose LLM provider",
			Items:     providers,
			CursorPos: defaultIdx,
			Stdout:    os.Stdout,
		}
		_, provider, err := providerSelect.Run()
		if err != nil {
			fmt.Println("Setup aborted.")
			os.Exit(1)
		}
		cfg.Provider = provider

		// 2) Model name
		modelPrompt := promptui.Prompt{
			Label:   "Model name",
			Default: cfg.Model,
			Stdout:  os.Stdout,
		}
		if result, err := modelPrompt.Run(); err == nil {
			cfg.Model = result
		}

		// 3) API Key (only for http)
		if cfg.Provider == "http" {
			keyPrompt := promptui.Prompt{
				Label:   "API Key (for HTTP, leave blank to skip)",
				Default: cfg.APIKey,
				Mask:    '*',
				Stdout:  os.Stdout,
			}
			if result, err := keyPrompt.Run(); err == nil {
				cfg.APIKey = result
			}
		}

		// 4) Base URL
		urlPrompt := promptui.Prompt{
			Label:   "Base URL",
			Default: cfg.BaseURL,
			Stdout:  os.Stdout,
		}
		if result, err := urlPrompt.Run(); err == nil {
			cfg.BaseURL = result
		}

		// 5) Save configuration
		if err := config.Save(cfg); err != nil {
			fmt.Println("❌ Failed to save config:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Configuration saved.")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
