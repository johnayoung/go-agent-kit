package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestInstallCommand(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "install-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current dir: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to temp dir: %v", err)
	}

	tests := []struct {
		name           string
		expectedError  bool
		checkFiles     []string
		expectedInText []string
	}{
		{
			name:          "install command creates .github directory and files",
			expectedError: false,
			checkFiles: []string{
				".github/copilot-instructions.md",
				".github/prompts/feat.prompt.md",
				".github/prompts/fix.prompt.md",
				".github/prompts/refactor.prompt.md",
			},
			expectedInText: []string{
				"Successfully installed GitHub Copilot integration!",
				"Created: .github/copilot-instructions.md",
				"Created: .github/prompts/feat.prompt.md",
				"Created: .github/prompts/fix.prompt.md",
				"Created: .github/prompts/refactor.prompt.md",
				"/feat add user authentication",
				"/fix null pointer exception",
				"/refactor simplify error handling",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a command for testing
			var output strings.Builder
			cmd := &cobra.Command{
				Use:  "install",
				RunE: runInstall,
			}
			cmd.SetOut(&output)
			cmd.SetErr(&output)

			// Execute the command
			err := runInstall(cmd, []string{})

			// Check error expectation
			if tt.expectedError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectedError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			// If we expected an error, we're done
			if tt.expectedError {
				return
			}

			// Check output content
			outputStr := output.String()
			for _, expectedText := range tt.expectedInText {
				if !strings.Contains(outputStr, expectedText) {
					t.Errorf("Expected to find '%s' in output", expectedText)
				}
			}

			// Check if files were created
			for _, filePath := range tt.checkFiles {
				if _, err := os.Stat(filePath); os.IsNotExist(err) {
					t.Errorf("Expected file %s to be created", filePath)
				}

				// Check file content for prompt files
				if strings.HasSuffix(filePath, ".prompt.md") {
					content, err := os.ReadFile(filePath)
					if err != nil {
						t.Errorf("Failed to read created file %s: %v", filePath, err)
					}

					contentStr := string(content)
					if strings.Contains(filePath, "feat") {
						expectedContent := []string{
							"Feature Implementation Workflow",
							"STAGE 1: CODEBASE ANALYSIS",
							"{{.Description}}",
						}
						for _, expected := range expectedContent {
							if !strings.Contains(contentStr, expected) {
								t.Errorf("Expected to find '%s' in %s", expected, filePath)
							}
						}
					} else if strings.Contains(filePath, "fix") {
						expectedContent := []string{
							"Bug Fix Workflow",
							"STAGE 1: DIAGNOSIS",
							"{{.Description}}",
						}
						for _, expected := range expectedContent {
							if !strings.Contains(contentStr, expected) {
								t.Errorf("Expected to find '%s' in %s", expected, filePath)
							}
						}
					} else if strings.Contains(filePath, "refactor") {
						expectedContent := []string{
							"Code Refactor Workflow",
							"STAGE 1: CODEBASE ANALYSIS",
							"{{.Description}}",
						}
						for _, expected := range expectedContent {
							if !strings.Contains(contentStr, expected) {
								t.Errorf("Expected to find '%s' in %s", expected, filePath)
							}
						}
					}
				}
			}
		})
	}
}

func TestGenerateCopilotInstructions(t *testing.T) {
	instructions := generateCopilotInstructions()

	expectedContent := []string{
		"GitHub Copilot Instructions for go-agent-kit",
		"/feat - Feature Implementation Workflow",
		"/fix - Bug Fix Workflow",
		"/refactor - Code Refactoring Workflow",
		"CODEBASE ANALYSIS",
		"IMPLEMENTATION PLAN",
		"Language-Agnostic Design",
		"Integration with GitHub Copilot",
	}

	for _, expected := range expectedContent {
		if !strings.Contains(instructions, expected) {
			t.Errorf("Expected to find '%s' in generated instructions", expected)
		}
	}

	// Check that instructions are substantial
	if len(instructions) < 1000 {
		t.Error("Generated instructions seem too short")
	}
}
