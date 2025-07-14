package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
    Use:   "smartcommit",
    Short: "AI‑powered Git commit message generator",
    Long: `
SmartCommit reads your staged Git diff and uses any LLM (local or HTTP API) 
to generate a meaningful commit message, then commits it for you.

Examples:
  # Configure your LLM provider and API key:
  smartcommit setup

  # Generate & commit a message:
  smartcommit generate

If you ever need assistance, run:
  smartcommit --help
or see:
  https://github.com/your‑repo/smartcommit#readme
`,
}


func Execute() error {
    return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
