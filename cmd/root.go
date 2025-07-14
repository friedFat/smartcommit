package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v1.1.0"

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
    Version: version,

}
var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print SmartCommit version",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("SmartCommit %s\n", version)
    },
}


func Execute() error {
    return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
  rootCmd.SetVersionTemplate("SmartCommit {{.Version}}\n")
  rootCmd.AddCommand(versionCmd)
  rootCmd.SetVersionTemplate("SmartCommit {{.Version}}\n")


}
