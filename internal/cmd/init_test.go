package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/johnayoung/go-agent-kit/internal/filewriter"
)

func TestRunInitCreatesExpectedFiles(t *testing.T) {
	// Create temporary directory for test
	tempDir := t.TempDir()
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	// Change to temp directory
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}
	defer func() {
		os.Chdir(originalWd)
	}()

	// Run init command
	err = runInit(initCmd, []string{})
	if err != nil {
		t.Fatalf("runInit failed: %v", err)
	}

	// Verify expected files were created
	expectedFiles := []string{
		".github/copilot-instructions.md",
		".github/prompts/specify.prompt.md",
		".github/prompts/plan.prompt.md",
		".github/prompts/implement.prompt.md",
		".github/prompts/verify.prompt.md",
	}

	writer := filewriter.New()
	for _, file := range expectedFiles {
		if !writer.PathExists(file) {
			t.Errorf("Expected file not created: %s", file)
		}
	}
}

func TestRunInitWithForceFlag(t *testing.T) {
	// Create temporary directory for test
	tempDir := t.TempDir()
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	// Change to temp directory
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}
	defer func() {
		os.Chdir(originalWd)
	}()

	// Create .github/prompts directory to simulate existing structure
	writer := filewriter.New()
	if err := writer.WriteFile(".github/prompts/existing.md", "existing content"); err != nil {
		t.Fatalf("Failed to create existing file: %v", err)
	}

	// Set force flag and run init
	forceFlag = true
	err = runInit(initCmd, []string{})
	if err != nil {
		t.Fatalf("runInit with force flag failed: %v", err)
	}
	forceFlag = false // Reset for other tests

	// Verify new files were created even with existing directory
	expectedFiles := []string{
		".github/copilot-instructions.md",
		".github/prompts/specify.prompt.md",
		".github/prompts/plan.prompt.md",
		".github/prompts/implement.prompt.md",
		".github/prompts/verify.prompt.md",
	}

	for _, file := range expectedFiles {
		if !writer.PathExists(file) {
			t.Errorf("Expected file not created with force flag: %s", file)
		}
	}
}

func TestRunInitTemplateContent(t *testing.T) {
	// Create temporary directory for test
	tempDir := t.TempDir()
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	// Change to temp directory
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}
	defer func() {
		os.Chdir(originalWd)
	}()

	// Run init command
	err = runInit(initCmd, []string{})
	if err != nil {
		t.Fatalf("runInit failed: %v", err)
	}

	// Test specify prompt content
	specifyContent, err := os.ReadFile(".github/prompts/specify.prompt.md")
	if err != nil {
		t.Fatalf("Failed to read specify prompt: %v", err)
	}

	specifyStr := string(specifyContent)
	if strings.Contains(specifyStr, "TODO") {
		t.Error("Specify prompt contains TODO placeholder")
	}
	if !strings.Contains(specifyStr, "Specification Assistant") {
		t.Error("Specify prompt missing expected content")
	}

	// Test plan prompt content
	planContent, err := os.ReadFile(".github/prompts/plan.prompt.md")
	if err != nil {
		t.Fatalf("Failed to read plan prompt: %v", err)
	}

	planStr := string(planContent)
	if strings.Contains(planStr, "TODO") {
		t.Error("Plan prompt contains TODO placeholder")
	}
	if !strings.Contains(planStr, "Technical Planning Assistant") {
		t.Error("Plan prompt missing expected content")
	}

	// Test that other files have placeholder content (as expected)
	implementContent, err := os.ReadFile(".github/prompts/implement.prompt.md")
	if err != nil {
		t.Fatalf("Failed to read implement prompt: %v", err)
	}

	if !strings.Contains(string(implementContent), "TODO") {
		t.Error("Implement prompt should contain TODO placeholder")
	}
}

func TestInitCmdConfiguration(t *testing.T) {
	if initCmd.Use != "init" {
		t.Errorf("Expected Use to be 'init', got '%s'", initCmd.Use)
	}

	if initCmd.Short == "" {
		t.Error("Short description should not be empty")
	}

	if initCmd.Long == "" {
		t.Error("Long description should not be empty")
	}

	if initCmd.RunE == nil {
		t.Error("RunE function should not be nil")
	}
}
