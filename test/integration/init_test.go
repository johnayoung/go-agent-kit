package integration

import (
	"os"
	"strings"
	"testing"
)

// TestInitCommand tests the init command functionality in isolation
func TestInitCommand(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "agentkit_init_test")
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

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change to temp dir: %v", err)
	}

	t.Run("InitCreatesFiles", func(t *testing.T) {
		expectedFiles := []string{
			".github/copilot-instructions.md",
			".github/prompts/specify.prompt.md",
			".github/prompts/plan.prompt.md",
			".github/prompts/implement.prompt.md",
			".github/prompts/verify.prompt.md",
		}

		// Test that files don't exist initially
		for _, file := range expectedFiles {
			if _, err := os.Stat(file); !os.IsNotExist(err) {
				t.Errorf("File %s should not exist initially", file)
			}
		}

		// TODO: Refactor cmd package to allow dependency injection for testing
		t.Skip("Skipping until cmd package is refactored for testability")
	})

	t.Run("InitFailsWithoutForce", func(t *testing.T) {
		// Create the .github/prompts directory
		err := os.MkdirAll(".github/prompts", 0755)
		if err != nil {
			t.Fatalf("Failed to create .github/prompts: %v", err)
		}

		// TODO: Test that init command fails when directory exists and --force is not used
		t.Skip("Skipping until cmd package is refactored for testability")
	})

	t.Run("InitSucceedsWithForce", func(t *testing.T) {
		// Create the .github/prompts directory with existing files
		err := os.MkdirAll(".github/prompts", 0755)
		if err != nil {
			t.Fatalf("Failed to create .github/prompts: %v", err)
		}

		existingFile := ".github/prompts/specify.prompt.md"
		err = os.WriteFile(existingFile, []byte("existing content"), 0644)
		if err != nil {
			t.Fatalf("Failed to create existing file: %v", err)
		}

		// TODO: Test that init command succeeds with --force flag
		t.Skip("Skipping until cmd package is refactored for testability")
	})
}

// TestInitCommandEndToEnd tests the init command through the CLI
func TestInitCommandEndToEnd(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "agentkit_e2e_test")
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

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change to temp dir: %v", err)
	}

	t.Run("CLIInitCommand", func(t *testing.T) {
		// This would require building and executing the CLI binary
		// For now, we'll focus on unit tests and defer this
		t.Skip("End-to-end CLI testing deferred for now")
	})
}

// Helper function to check if file exists and has expected content
func checkFileContent(t *testing.T, filePath, expectedContent string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Failed to read file %s: %v", filePath, err)
		return
	}

	if strings.TrimSpace(string(content)) != strings.TrimSpace(expectedContent) {
		t.Errorf("File %s content mismatch.\nExpected: %q\nGot: %q",
			filePath, expectedContent, string(content))
	}
}
