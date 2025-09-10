package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-agent-kit",
	Short: "A language-agnostic toolkit for creating structured AI agent workflows",
	Long: `go-agent-kit is a language-agnostic toolkit for creating structured AI agent workflows.
While written in Go, it works with any programming language or framework.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
