package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/johnayoung/go-agent-kit/internal/filewriter"
	"github.com/spf13/cobra"
)

var forceFlag bool

// File structure to create
var filesToCreate = map[string]string{
	".github/copilot-instructions.md":     "TODO: copilot-instructions.md",
	".github/prompts/specify.prompt.md":   "TODO: specify.prompt.md",
	".github/prompts/plan.prompt.md":      "TODO: plan.prompt.md",
	".github/prompts/implement.prompt.md": "TODO: implement.prompt.md",
	".github/prompts/verify.prompt.md":    "TODO: verify.prompt.md",
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize AgentKit prompt structure",
	Long: `Initialize creates the .github directory structure with copilot-instructions.md
and prompt templates for the AgentKit workflow.`,
	Run: runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	writer := filewriter.New()

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
		return
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
			return
		}
		green.Printf("✓ ")
		fmt.Printf("Created %s\n", filePath)
	}

	fmt.Println()
	green.Printf("✓ AgentKit structure initialized successfully!\n")
	fmt.Println()
	fmt.Println("Next steps:")
	cyan.Printf("  agentkit specify")
	fmt.Printf(" - Create project specification\n")
	cyan.Printf("  agentkit plan")
	fmt.Printf("    - Generate implementation plan\n")
	cyan.Printf("  agentkit implement")
	fmt.Printf(" - Execute implementation\n")
	cyan.Printf("  agentkit verify")
	fmt.Printf("   - Verify implementation\n")
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force overwrite existing files")
}
