package cmd

import (
	"os"
	"path/filepath"
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
		checkFile      bool
		expectedInText []string
	}{
		{
			name:          "install command creates .github directory and file",
			expectedError: false,
			checkFile:     true,
			expectedInText: []string{
				"Successfully installed GitHub Copilot integration!",
				"Created: .github/copilot-instructions.md",
				"/feat add user authentication",
				"/fix null pointer exception",
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

			// Check if file was created
			if tt.checkFile {
				instructionsPath := filepath.Join(".github", "copilot-instructions.md")
				if _, err := os.Stat(instructionsPath); os.IsNotExist(err) {
					t.Errorf("Expected file %s to be created", instructionsPath)
				}

				// Check file content
				content, err := os.ReadFile(instructionsPath)
				if err != nil {
					t.Errorf("Failed to read created file: %v", err)
				}

				contentStr := string(content)
				expectedFileContent := []string{
					"GitHub Copilot Instructions for go-agent-kit",
					"/feat - Feature Implementation Workflow",
					"/fix - Bug Fix Workflow",
					"Language-Agnostic Design",
				}

				for _, expected := range expectedFileContent {
					if !strings.Contains(contentStr, expected) {
						t.Errorf("Expected to find '%s' in file content", expected)
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
