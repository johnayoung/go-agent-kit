package templates

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	service, err := New()
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}

	if service == nil {
		t.Fatal("New() returned nil service")
	}

	if service.templates == nil {
		t.Fatal("New() returned service with nil templates")
	}
}

func TestGetSpecifyPrompt(t *testing.T) {
	service, err := New()
	if err != nil {
		t.Fatalf("Failed to create service: %v", err)
	}

	content, err := service.GetSpecifyPrompt()
	if err != nil {
		t.Fatalf("GetSpecifyPrompt() failed: %v", err)
	}

	if content == "" {
		t.Fatal("GetSpecifyPrompt() returned empty content")
	}

	// Verify it's not placeholder content
	if strings.Contains(content, "TODO") {
		t.Error("GetSpecifyPrompt() returned placeholder content")
	}

	// Verify it contains expected content
	expectedPhrases := []string{
		"Specification Assistant",
		"GitHub Copilot Instructions",
		"/specify",
		"project/specification.md",
	}

	for _, phrase := range expectedPhrases {
		if !strings.Contains(content, phrase) {
			t.Errorf("GetSpecifyPrompt() missing expected phrase: %s", phrase)
		}
	}
}

func TestGetPlanPrompt(t *testing.T) {
	service, err := New()
	if err != nil {
		t.Fatalf("Failed to create service: %v", err)
	}

	content, err := service.GetPlanPrompt()
	if err != nil {
		t.Fatalf("GetPlanPrompt() failed: %v", err)
	}

	if content == "" {
		t.Fatal("GetPlanPrompt() returned empty content")
	}

	// Verify it's not placeholder content
	if strings.Contains(content, "TODO") {
		t.Error("GetPlanPrompt() returned placeholder content")
	}

	// Verify it contains expected content
	expectedPhrases := []string{
		"Technical Planning Assistant",
		"GitHub Copilot Instructions",
		"/plan",
		"project/architecture.md",
		"Pre-flight Checks",
	}

	for _, phrase := range expectedPhrases {
		if !strings.Contains(content, phrase) {
			t.Errorf("GetPlanPrompt() missing expected phrase: %s", phrase)
		}
	}
}

func TestTemplateEmbedding(t *testing.T) {
	// Test that template files are properly embedded
	files, err := templatesFS.ReadDir("prompts")
	if err != nil {
		t.Fatalf("Failed to read embedded templates: %v", err)
	}

	expectedFiles := []string{"specify.prompt.md", "plan.prompt.md"}
	foundFiles := make(map[string]bool)

	for _, file := range files {
		foundFiles[file.Name()] = true
	}

	for _, expectedFile := range expectedFiles {
		if !foundFiles[expectedFile] {
			t.Errorf("Expected embedded file not found: %s", expectedFile)
		}
	}
}
