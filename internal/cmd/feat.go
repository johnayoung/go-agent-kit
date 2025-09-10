package cmd

import (
	"fmt"
	"strings"

	"github.com/johnayoung/go-agent-kit/internal/templates"
	"github.com/spf13/cobra"
)

// featCmd represents the feat command
var featCmd = &cobra.Command{
	Use:   "feat [description]",
	Short: "Generate a feature implementation workflow",
	Long: `Generate a structured AI agent workflow for implementing a new feature.

The feat command creates a comprehensive, language-agnostic workflow that guides
you through analyzing the codebase, planning the implementation, writing code,
testing, and documenting the feature.

Examples:
  go-agent-kit feat "add user authentication"
  go-agent-kit feat "implement REST API with JWT"
  go-agent-kit feat "add file upload functionality"`,
	Args: cobra.MinimumNArgs(1),
	RunE: runFeat,
}

func runFeat(cmd *cobra.Command, args []string) error {
	// Join all arguments into a description
	description := strings.Join(args, " ")

	// Create context for template rendering
	ctx := templates.Context{
		Description: description,
	}

	// Render the feat template
	output, err := templates.Render("feat", ctx)
	if err != nil {
		return fmt.Errorf("failed to render feat template: %w", err)
	}

	// Print the rendered output to the command's output
	cmd.Print(output)
	return nil
}

func init() {
	// Add feat command to root command
	rootCmd.AddCommand(featCmd)
}
