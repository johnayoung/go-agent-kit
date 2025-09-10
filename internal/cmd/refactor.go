package cmd

import (
	"fmt"
	"strings"

	"github.com/johnayoung/go-agent-kit/internal/templates"
	"github.com/spf13/cobra"
)

// refactorCmd represents the refactor command
var refactorCmd = &cobra.Command{
	Use:   "refactor [description]",
	Short: "Generate a code refactoring workflow",
	Long: `Generate a structured AI agent workflow for refactoring and improving existing code.

The refactor command creates a comprehensive, language-agnostic workflow that guides
you through analyzing the current code, planning improvements, implementing changes,
testing thoroughly, and documenting the refactoring.

Examples:
  go-agent-kit refactor "simplify user authentication logic"
  go-agent-kit refactor "extract payment processing into separate service"
  go-agent-kit refactor "improve error handling in API layer"
  go-agent-kit refactor "optimize database query performance"`,
	Args: cobra.MinimumNArgs(1),
	RunE: runRefactor,
}

func runRefactor(cmd *cobra.Command, args []string) error {
	// Join all arguments into a description
	description := strings.Join(args, " ")

	// Create context for template rendering
	ctx := templates.Context{
		Description: description,
	}

	// Render the refactor template
	output, err := templates.Render("refactor", ctx)
	if err != nil {
		return fmt.Errorf("failed to render refactor template: %w", err)
	}

	// Print the rendered output to the command's output
	cmd.Print(output)
	return nil
}

func init() {
	// Add refactor command to root command
	rootCmd.AddCommand(refactorCmd)
}
