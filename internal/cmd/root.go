package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-agent-kit",
	Short: "Install GitHub Copilot integration for structured AI workflows",
	Long: `go-agent-kit installs GitHub Copilot integration files that enable structured AI agent workflows.
After installation, use /refactor and /instructions commands directly in GitHub Copilot Chat.
Works with any programming language or framework.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Root command doesn't need any flags for our simple use case
}
