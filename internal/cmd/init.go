package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/johnayoung/go-agent-kit/internal/filewriter"
	"github.com/johnayoung/go-agent-kit/internal/templates"
	"github.com/spf13/cobra"
)

var forceFlag bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize AgentKit prompt structure",
	Long: `Initialize creates the .github directory structure with copilot-instructions.md
and prompt templates for the AgentKit workflow.`,
	RunE: runInit,
}

func runInit(cmd *cobra.Command, args []string) error {
	writer := filewriter.New()

	// Initialize template service
	templateService, err := templates.New()
	if err != nil {
		red := color.New(color.FgRed, color.Bold)
		red.Printf("✗ Error: ")
		fmt.Printf("Failed to initialize templates: %v\n", err)
		return err
	}

	// Get template content
	specifyPrompt, err := templateService.GetSpecifyPrompt()
	if err != nil {
		red := color.New(color.FgRed, color.Bold)
		red.Printf("✗ Error: ")
		fmt.Printf("Failed to load specify template: %v\n", err)
		return err
	}

	planPrompt, err := templateService.GetPlanPrompt()
	if err != nil {
		red := color.New(color.FgRed, color.Bold)
		red.Printf("✗ Error: ")
		fmt.Printf("Failed to load plan template: %v\n", err)
		return err
	}

	// File structure to create
	filesToCreate := map[string]string{
		".github/copilot-instructions.md":     "TODO: copilot-instructions.md",
		".github/prompts/specify.prompt.md":   specifyPrompt,
		".github/prompts/plan.prompt.md":      planPrompt,
		".github/prompts/implement.prompt.md": "TODO: implement.prompt.md",
		".github/prompts/verify.prompt.md":    "TODO: verify.prompt.md",
	}

	// Color setup
	green := color.New(color.FgGreen, color.Bold)
	red := color.New(color.FgRed, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	cyan := color.New(color.FgCyan)

	// Check if .github/prompts/ already exists
	promptsDir := ".github/prompts"
	if writer.PathExists(promptsDir) && !forceFlag {
		red.Printf("✗ Error: ")
		fmt.Printf("Directory %s already exists.\n", promptsDir)
		yellow.Printf("  Use ")
		cyan.Printf("agentkit init --force")
		yellow.Printf(" to overwrite existing files.\n")
		return fmt.Errorf("directory %s already exists", promptsDir)
	}

	if writer.PathExists(promptsDir) && forceFlag {
		yellow.Printf("⚠ Warning: ")
		fmt.Printf("Overwriting existing files in %s\n", promptsDir)
	}

	// Create all files
	fmt.Println("Creating AgentKit prompt structure...")
	for filePath, content := range filesToCreate {
		if err := writer.WriteFile(filePath, content); err != nil {
			red.Printf("✗ Failed to create %s: %v\n", filePath, err)
			return err
		}
		green.Printf("✓ ")
		fmt.Printf("Created %s\n", filePath)
	}

	fmt.Println()
	green.Printf("✓ AgentKit structure initialized successfully!\n")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("Open VS Code and use GitHub Copilot Chat with these commands:")
	cyan.Printf("  @workspace /specify")
	fmt.Printf("   - Create project specification\n")
	cyan.Printf("  @workspace /plan")
	fmt.Printf("      - Generate implementation plan\n")
	cyan.Printf("  @workspace /implement")
	fmt.Printf(" - Execute implementation\n")
	cyan.Printf("  @workspace /verify")
	fmt.Printf("     - Verify implementation\n")

	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force overwrite existing files")
}
