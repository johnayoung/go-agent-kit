package templates

import (
	"strings"
	"testing"
)

func TestPromptFilesEmbedded(t *testing.T) {
	// Test that we can access the embedded files
	entries, err := PromptFiles.ReadDir("prompts")
	if err != nil {
		t.Fatalf("Failed to read prompts directory: %v", err)
	}

	if len(entries) == 0 {
		t.Error("No files found in embedded prompts directory")
	}

	// Look for feat.md specifically
	var foundFeat bool
	for _, entry := range entries {
		if entry.Name() == "feat.md" {
			foundFeat = true
			break
		}
	}

	if !foundFeat {
		t.Error("feat.md not found in embedded files")
	}
}

func TestReadEmbeddedFile(t *testing.T) {
	// Test that we can read the embedded feat.md file
	content, err := PromptFiles.ReadFile("prompts/feat.md")
	if err != nil {
		t.Fatalf("Failed to read feat.md: %v", err)
	}

	if len(content) == 0 {
		t.Error("feat.md appears to be empty")
	}

	// Check for expected content
	contentStr := string(content)
	expectedStrings := []string{
		"Feature Implementation Workflow",
		"STAGE 1: CODEBASE ANALYSIS",
		"{{.Description}}",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(contentStr, expected) {
			t.Errorf("Expected to find '%s' in feat.md content", expected)
		}
	}
}
