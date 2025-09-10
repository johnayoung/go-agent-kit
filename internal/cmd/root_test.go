package cmd

import (
	"strings"
	"testing"
)

func TestRootCmdConfiguration(t *testing.T) {
	if rootCmd.Use != "agentkit" {
		t.Errorf("Expected Use to be 'agentkit', got '%s'", rootCmd.Use)
	}

	if rootCmd.Short == "" {
		t.Error("Short description should not be empty")
	}

	if rootCmd.Long == "" {
		t.Error("Long description should not be empty")
	}

	if rootCmd.Version != version {
		t.Errorf("Expected Version to be '%s', got '%s'", version, rootCmd.Version)
	}
}

func TestExecute(t *testing.T) {
	// Test that Execute function exists and is callable
	// We can't easily test execution without complex setup,
	// but we can verify the function doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Execute() panicked: %v", r)
		}
	}()

	// This will fail because we don't have args, but shouldn't panic
	Execute()
}

func TestVersionTemplate(t *testing.T) {
	// Test that version template is set correctly
	versionTemplate := rootCmd.VersionTemplate()
	expectedVersion := "agentkit version " + version + "\n"

	if !strings.Contains(versionTemplate, version) {
		t.Errorf("Version template should contain version '%s', got '%s'", version, versionTemplate)
	}

	if versionTemplate != expectedVersion {
		t.Errorf("Expected version template '%s', got '%s'", expectedVersion, versionTemplate)
	}
}

func TestVersionConstant(t *testing.T) {
	if version == "" {
		t.Error("Version constant should not be empty")
	}

	// Version should follow semantic versioning pattern
	if !strings.Contains(version, ".") {
		t.Error("Version should contain dots (semantic versioning)")
	}
}
