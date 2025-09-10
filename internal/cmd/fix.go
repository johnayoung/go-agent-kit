package cmd

import (
	"fmt"
	"strings"

	"github.com/johnayoung/go-agent-kit/internal/templates"
	"github.com/spf13/cobra"
)

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix [description]",
	Short: "Generate a bug fix workflow",
	Long: `Generate a structured AI agent workflow for fixing bugs and issues.

The fix command creates a comprehensive, language-agnostic workflow that guides
you through diagnosing the problem, planning the fix, implementing the solution,
testing thoroughly, and documenting the changes.

Examples:
  go-agent-kit fix "null pointer exception in user service"
  go-agent-kit fix "memory leak in background worker"
  go-agent-kit fix "authentication not working on mobile"`,
	Args: cobra.MinimumNArgs(1),
	RunE: runFix,
}

func runFix(cmd *cobra.Command, args []string) error {
	// Join all arguments into a description
	description := strings.Join(args, " ")

	// Create context for template rendering
	ctx := templates.Context{
		Description: description,
	}

	// Render the fix template
	output, err := templates.Render("fix", ctx)
	if err != nil {
		return fmt.Errorf("failed to render fix template: %w", err)
	}

	// Print the rendered output to the command's output
	cmd.Print(output)
	return nil
}

func init() {
	// Add fix command to root command
	rootCmd.AddCommand(fixCmd)
}
