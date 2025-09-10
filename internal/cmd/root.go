package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "agentkit",
	Short: "AgentKit CLI - AI-powered development workflow tool",
	Long: `AgentKit CLI helps you build applications using AI-powered prompts.
It provides a structured workflow for specification, planning, implementation, and verification.`,
	Version: version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.SetVersionTemplate(fmt.Sprintf("agentkit version %s\n", version))
}
