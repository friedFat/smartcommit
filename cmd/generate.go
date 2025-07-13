// File: cmd/generate.go
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"smartcommit/config"
	"smartcommit/diff"
	"smartcommit/llm"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var yesFlag bool

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a commit message using AI",
	Long: `Generate a commit message from staged Git changes using a local or remote LLM.

By default, it launches an interactive flow.
Use --yes or -y to skip the prompt and commit directly.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadOrDefault()

		d, err := diff.GetStagedDiff()
		if err != nil || strings.TrimSpace(d) == "" {
			fmt.Println("❌ No staged changes found.")
			return
		}

		prompt := cfg.SystemPrompt + "\n\nDiff:\n" + d

		provider, err := llm.GetProvider(cfg)
		if err != nil {
			fmt.Println("❌", err)
			return
		}

		message, err := provider.Generate(prompt)
		if err != nil {
			fmt.Println("❌ Generation failed:", err)
			return
		}
		message = strings.TrimSpace(message)

		if yesFlag {
			fmt.Println("💡 Generated Commit Message:")
			fmt.Println("----------------------------------")
			fmt.Println(message)
			fmt.Println("----------------------------------")
			commit(message)
			return
		}

		// Interactive loop
		for {
			fmt.Println("\n💡 Generated Commit Message:")
			fmt.Println("----------------------------------")
			fmt.Println(message)
			fmt.Println("----------------------------------")
			fmt.Print("Choose [c]ommit, [e]dit, [r]egenerate, [q]uit: ")

			var choice string
			fmt.Scanln(&choice)

			switch choice {
			case "c":
				commit(message)
				return
			case "e":
				message = edit(message)
			case "r":
				fmt.Print("\n🔄 Regenerating")
				for i := 0; i < 3; i++ {
					fmt.Print(".")
					time.Sleep(200 * time.Millisecond)
				}
				msg, err := provider.Generate(prompt)
				if err != nil {
					fmt.Println("\n❌ Regeneration failed:", err)
				} else {
					message = strings.TrimSpace(msg)
				}
			case "q":
				fmt.Println("👋 Aborted.")
				return
			default:
				fmt.Println("❓ Invalid choice.")
			}
		}
	},
}

func init() {
	generateCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "Autogenerate and commit without prompting")
	rootCmd.AddCommand(generateCmd)
}

func commit(msg string) {
	fmt.Println("✅ Committing with:")
	fmt.Println(msg)
	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("❌ Git commit failed:", err)
	}
}

func edit(current string) string {
	prompt := promptui.Prompt{
		Label:     "Edit Commit Message",
		Default:   current,
		AllowEdit: true,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println("⚠️ Edit cancelled.")
		return current
	}
	return result
}
