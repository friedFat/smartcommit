package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"smartcommit/config"
	"smartcommit/diff"
	"smartcommit/llm"
	"smartcommit/prompt"
	"time"

	promptui "github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a commit message using AI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadOrDefault()

		diffText, err := diff.GetStagedDiff()
		if err != nil || len(diffText) == 0 {
			fmt.Println("❌ No staged changes found.")
			return
		}

		promptText := prompt.Build(diffText, cfg.SystemPrompt)

		provider, err := llm.GetProvider(cfg)
		if err != nil {
			fmt.Println("❌", err)
			return
		}

		message, err := provider.Generate(promptText)
		if err != nil {
			fmt.Println("❌ Generation failed:", err)
			return
		}

		for {
			fmt.Println("\n💡 Generated Commit Message:")
			fmt.Println("----------------------------------")
			fmt.Println(message)
			fmt.Println("----------------------------------")
			fmt.Print("Choose [c]ommit, [e]dit message, [r]egenerate, [q]uit: ")

			var choice string
			fmt.Scanln(&choice)

			switch choice {
			case "c":
				runGitCommit(message)
				return
			case "e":
				message = editMessage(message)
			case "r":
				fmt.Print("\n🔄 Regenerating")
				for i := 0; i < 3; i++ {
					fmt.Print(".")
					time.Sleep(200 * time.Millisecond)
				}
				message, err = provider.Generate(promptText)
				if err != nil {
					fmt.Println("❌ Regeneration failed:", err)
					continue
				}
			case "q":
				return
			default:
				fmt.Println("❓ Invalid choice.")
			}
		}
	},
}

func runGitCommit(msg string) {
	fmt.Println("✅ Committing with:")
	fmt.Println(msg)
	_ = os.WriteFile(".git/COMMIT_EDITMSG", []byte(msg), 0644)
	_ = executeCommand("git", "commit", "-F", ".git/COMMIT_EDITMSG")
}

func executeCommand(name string, args ...string) error {
	c := exec.Command(name, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

func editMessage(current string) string {
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
