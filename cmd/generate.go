package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"smartcommit/config"
	"smartcommit/diff"
	"smartcommit/llm"
)

var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate & commit a message from your diff",
    Run: func(cmd *cobra.Command, args []string) {
        cfg := config.LoadOrDefault()

        d, err := diff.GetStagedDiff()
        if err != nil || d == "" {
            fmt.Println("❌ No staged changes.")
            return
        }

        provider, err := llm.GetProvider(cfg)
        if err != nil {
            fmt.Println("❌", err)
            return
        }

        prompt := cfg.SystemPrompt + "\n\nDiff:\n" + d
        msg, err := provider.Generate(prompt)
        if err != nil {
            fmt.Println("❌ Generation failed:", err)
            return
        }
        msg = strings.TrimSpace(msg)

        fmt.Println("\n💡 Commit message:\n\n", msg, "\n")
        git := exec.Command("git", "commit", "-m", msg)
        git.Stdout = cmd.OutOrStdout()
        git.Stderr = cmd.OutOrStderr()
        if err := git.Run(); err != nil {
            fmt.Println("❌ Git commit failed:", err)
        }
    },
}

func init() {
    rootCmd.AddCommand(generateCmd)
}
