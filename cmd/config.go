// File: cmd/config.go
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"smartcommit/config"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "View or change smartcommit configuration",
}

var setCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key, value := args[0], args[1]
		cfg := config.LoadOrDefault()
		cfg.Set(key, value)
		if err := config.Save(cfg); err != nil {
			fmt.Println("❌ Failed to save config:", err)
			return
		}
		fmt.Println("✅ Config updated")
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadOrDefault()
		cfg.PrettyPrint()
	},
}

var editCmd = &cobra.Command{
	Use:   "edit system_prompt",
	Short: "Edit the system prompt using your default editor",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "system_prompt" {
			fmt.Println("❌ Only 'system_prompt' can be edited via editor for now.")
			return
		}

		cfg := config.LoadOrDefault()
		tmpfile := filepath.Join(os.TempDir(), "smartcommit_system_prompt.txt")
		os.WriteFile(tmpfile, []byte(cfg.SystemPrompt), 0644)

		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "vim"
		}

		execCmd := exec.Command(editor, tmpfile)
		execCmd.Stdin = os.Stdin
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		execCmd.Run()

		updated, _ := os.ReadFile(tmpfile)
		cfg.SystemPrompt = strings.TrimSpace(string(updated))
		if err := config.Save(cfg); err != nil {
			fmt.Println("❌ Failed to save:", err)
			return
		}
		fmt.Println("✅ Updated system prompt")
	},
}

func init() {
	ConfigCmd.AddCommand(setCmd, showCmd, editCmd)
}
