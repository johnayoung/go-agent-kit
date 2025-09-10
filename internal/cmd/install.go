package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/johnayoung/go-agent-kit/internal/templates"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install GitHub Copilot integration files",
	Long: `Install creates a .github/copilot-instructions.md file that tells GitHub Copilot
how to use the go-agent-kit workflows in your project.

After running this command, you can use GitHub Copilot commands like:
  /feat add user authentication
  /fix null pointer exception

The install command creates language-agnostic instructions that work with any
programming language or framework.`,
	RunE: runInstall,
}

func runInstall(cmd *cobra.Command, args []string) error {
	// Create .github directory if it doesn't exist
	githubDir := filepath.Join(".", ".github")
	if err := os.MkdirAll(githubDir, 0755); err != nil {
		return fmt.Errorf("failed to create .github directory: %w", err)
	}

	// Create .github/prompts directory if it doesn't exist
	promptsDir := filepath.Join(githubDir, "prompts")
	if err := os.MkdirAll(promptsDir, 0755); err != nil {
		return fmt.Errorf("failed to create .github/prompts directory: %w", err)
	}

	// Install prompt files
	if err := installPromptFiles(promptsDir); err != nil {
		return fmt.Errorf("failed to install prompt files: %w", err)
	}

	// Generate the copilot instructions
	instructions := generateCopilotInstructions()

	// Write the instructions file
	instructionsPath := filepath.Join(githubDir, "copilot-instructions.md")
	if err := os.WriteFile(instructionsPath, []byte(instructions), 0644); err != nil {
		return fmt.Errorf("failed to write copilot instructions: %w", err)
	}

	fmt.Println("✅ GitHub Copilot integration installed successfully!")
	fmt.Println()
	fmt.Println("Available commands in GitHub Copilot Chat:")
	fmt.Println("  /refactor [description] - Systematic code refactoring workflow")
	fmt.Println("  /instructions           - Generate GitHub Copilot instructions")
	fmt.Println()
	fmt.Println("Example usage:")
	fmt.Println("  /refactor simplify user authentication logic")
	fmt.Println("  /instructions")

	return nil
}

func installPromptFiles(promptsDir string) error {
	// Get feat template content
	featContent, err := templates.PromptFiles.ReadFile("prompts/feat.md")
	if err != nil {
		return fmt.Errorf("failed to read feat template: %w", err)
	}

	// Write feat.prompt.md
	featPath := filepath.Join(promptsDir, "feat.prompt.md")
	if err := os.WriteFile(featPath, featContent, 0644); err != nil {
		return fmt.Errorf("failed to write feat prompt file: %w", err)
	}

	// Get fix template content
	fixContent, err := templates.PromptFiles.ReadFile("prompts/fix.md")
	if err != nil {
		return fmt.Errorf("failed to read fix template: %w", err)
	}

	// Write fix.prompt.md
	fixPath := filepath.Join(promptsDir, "fix.prompt.md")
	if err := os.WriteFile(fixPath, fixContent, 0644); err != nil {
		return fmt.Errorf("failed to write fix prompt file: %w", err)
	}

	// Get refactor template content
	refactorContent, err := templates.PromptFiles.ReadFile("prompts/refactor.md")
	if err != nil {
		return fmt.Errorf("failed to read refactor template: %w", err)
	}

	// Write refactor.prompt.md
	refactorPath := filepath.Join(promptsDir, "refactor.prompt.md")
	if err := os.WriteFile(refactorPath, refactorContent, 0644); err != nil {
		return fmt.Errorf("failed to write refactor prompt file: %w", err)
	}

	// Get instructions template content
	instructionsContent, err := templates.PromptFiles.ReadFile("prompts/instructions.md")
	if err != nil {
		return fmt.Errorf("failed to read instructions template: %w", err)
	}

	// Write instructions.prompt.md
	instructionsPath := filepath.Join(promptsDir, "instructions.prompt.md")
	if err := os.WriteFile(instructionsPath, instructionsContent, 0644); err != nil {
		return fmt.Errorf("failed to write instructions prompt file: %w", err)
	}

	return nil
}

func generateCopilotInstructions() string {
	return `# GitHub Copilot Instructions for go-agent-kit

This project uses go-agent-kit for structured AI agent workflows. Use the following commands for systematic development:

## Available Commands

### /feat - Feature Implementation Workflow
Use this command to implement new features with a structured approach.

**Usage:**
` + "```" + `
/feat [description of the feature to implement]
` + "```" + `

**Examples:**
- /feat add user authentication system
- /feat implement REST API with JWT tokens
- /feat add file upload functionality
- /feat create admin dashboard

**What it does:**
Generates a comprehensive 5-stage workflow:
1. **CODEBASE ANALYSIS** - Detect language, examine patterns, find integration points
2. **IMPLEMENTATION PLAN** - Plan files, dependencies, and implementation order
3. **IMPLEMENTATION** - Step-by-step coding with language-specific best practices
4. **TESTING** - Unit tests, integration tests, and edge cases
5. **DOCUMENTATION** - Code comments, README updates, and API docs

### /fix - Bug Fix Workflow
Use this command to systematically diagnose and fix bugs.

**Usage:**
` + "```" + `
/fix [description of the bug or issue]
` + "```" + `

**Examples:**
- /fix null pointer exception in user service
- /fix memory leak in background worker
- /fix authentication not working on mobile
- /fix database connection timeout errors

**What it does:**
Generates a systematic 5-stage debugging workflow:
1. **DIAGNOSIS** - Understand, locate, reproduce, and analyze the issue
2. **FIX STRATEGY** - Plan the fix approach and assess impact
3. **IMPLEMENTATION** - Apply minimal fix with safety checks
4. **TESTING** - Verify fix and run regression tests
5. **DOCUMENTATION** - Document the fix and add preventive measures

### /refactor - Code Refactoring Workflow
Use this command to systematically improve and refactor existing code.

**Usage:**
` + "```" + `
/refactor [description of the refactoring task]
` + "```" + `

**Examples:**
- /refactor simplify user authentication logic
- /refactor extract payment processing into separate service
- /refactor optimize database query performance
- /refactor improve error handling patterns

**What it does:**
Generates a comprehensive 5-stage refactoring workflow:
1. **CODEBASE ANALYSIS** - Understand current implementation and identify improvements
2. **REFACTOR PLAN** - Plan refactoring strategy and assess risks
3. **IMPLEMENTATION** - Apply refactoring techniques systematically
4. **TESTING** - Verify functionality and performance are maintained
5. **DOCUMENTATION** - Update docs to reflect architectural changes

## Language-Agnostic Design

These workflows are designed to work with ANY programming language:
- **Go** - Follows Go conventions, error patterns, and testing practices
- **Python** - Uses PEP 8, type hints, and Python idioms
- **TypeScript/JavaScript** - Proper types, async/await, modern patterns
- **Java** - Java conventions, exception handling, design patterns
- **C#** - .NET patterns, LINQ, async/await
- **Ruby** - Ruby style guide, Rails conventions where applicable
- **And many more...**

## How It Works

1. **Language Detection**: Workflows automatically detect your project's language by examining files like go.mod, package.json, requirements.txt, etc.

2. **Pattern Analysis**: The AI analyzes your existing codebase to understand your specific patterns, architecture, and conventions.

3. **Guided Implementation**: Each stage provides specific guidance while respecting your project's established patterns.

4. **Best Practices**: Language-specific guidelines ensure code follows community standards and best practices.

## Integration with GitHub Copilot

When you use these commands in GitHub Copilot Chat:

1. **Copy the generated workflow** from the command output
2. **Follow each stage systematically** - don't skip ahead
3. **Let Copilot examine your codebase** when prompted with @workspace
4. **Implement step by step** as guided by the workflow

## Benefits

- ✅ **Consistent Quality**: Every feature and fix follows the same systematic approach
- ✅ **Language Agnostic**: Works across all programming languages and frameworks  
- ✅ **Best Practices**: Incorporates language-specific conventions and patterns
- ✅ **Comprehensive**: Covers analysis, implementation, testing, and documentation
- ✅ **AI-Optimized**: Designed specifically for AI agents like GitHub Copilot

## Getting Started

1. Run ` + "`go-agent-kit install`" + ` in your project (already done!)
2. Open GitHub Copilot Chat
3. Try: ` + "`/feat add a simple hello world endpoint`" + `
4. Follow the generated workflow step by step

---

*Generated by go-agent-kit - A language-agnostic toolkit for structured AI agent workflows.*`
}

func init() {
	// Add install command to root command
	rootCmd.AddCommand(installCmd)
}
