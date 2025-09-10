package integration

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestAgentKitInitIntegration(t *testing.T) {
	// Build the binary first
	binaryPath := buildAgentKitBinary(t)
	defer os.Remove(binaryPath)

	// Create temporary directory for test
	tempDir := t.TempDir()

	// Run agentkit init in temp directory
	cmd := exec.Command(binaryPath, "init")
	cmd.Dir = tempDir

	startTime := time.Now()
	output, err := cmd.CombinedOutput()
	duration := time.Since(startTime)

	if err != nil {
		t.Fatalf("agentkit init failed: %v\nOutput: %s", err, output)
	}

	// Test performance requirement (should complete in <1 second)
	if duration > time.Second {
		t.Errorf("agentkit init took too long: %v (should be <1s)", duration)
	}

	// Verify output messages
	outputStr := string(output)
	expectedMessages := []string{
		"Creating AgentKit prompt structure...",
		"✓ Created .github/copilot-instructions.md",
		"✓ Created .github/prompts/specify.prompt.md",
		"✓ Created .github/prompts/plan.prompt.md",
		"✓ Created .github/prompts/implement.prompt.md",
		"✓ Created .github/prompts/verify.prompt.md",
		"✓ AgentKit structure initialized successfully!",
		"@workspace /specify",
		"@workspace /plan",
	}

	for _, msg := range expectedMessages {
		if !strings.Contains(outputStr, msg) {
			t.Errorf("Output missing expected message: %s", msg)
		}
	}

	// Verify all expected files were created
	expectedFiles := []string{
		".github/copilot-instructions.md",
		".github/prompts/specify.prompt.md",
		".github/prompts/plan.prompt.md",
		".github/prompts/implement.prompt.md",
		".github/prompts/verify.prompt.md",
	}

	for _, file := range expectedFiles {
		fullPath := filepath.Join(tempDir, file)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			t.Errorf("Expected file not created: %s", file)
		}
	}

	// Verify template content quality
	verifyTemplateContent(t, tempDir)
}

func TestAgentKitInitForceFlag(t *testing.T) {
	// Build the binary first
	binaryPath := buildAgentKitBinary(t)
	defer os.Remove(binaryPath)

	// Create temporary directory for test
	tempDir := t.TempDir()

	// Create existing .github/prompts directory
	promptsDir := filepath.Join(tempDir, ".github", "prompts")
	if err := os.MkdirAll(promptsDir, 0755); err != nil {
		t.Fatalf("Failed to create prompts directory: %v", err)
	}

	// Create existing file
	existingFile := filepath.Join(promptsDir, "existing.md")
	if err := os.WriteFile(existingFile, []byte("existing content"), 0644); err != nil {
		t.Fatalf("Failed to create existing file: %v", err)
	}

	// Run agentkit init without force (should fail)
	cmd := exec.Command(binaryPath, "init")
	cmd.Dir = tempDir
	output, err := cmd.CombinedOutput()

	// Should exit with non-zero code since directory exists
	if err == nil {
		t.Error("agentkit init should have failed without --force flag")
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "already exists") {
		t.Errorf("Error message should mention directory already exists, got: %s", outputStr)
	}

	// Run agentkit init with force flag (should succeed)
	cmd = exec.Command(binaryPath, "init", "--force")
	cmd.Dir = tempDir
	output, err = cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("agentkit init --force failed: %v\nOutput: %s", err, output)
	}

	if !strings.Contains(string(output), "Overwriting existing files") {
		t.Error("Should show overwrite warning with --force flag")
	}
}

func TestAgentKitVersion(t *testing.T) {
	// Build the binary first
	binaryPath := buildAgentKitBinary(t)
	defer os.Remove(binaryPath)

	// Test version flag (not subcommand)
	cmd := exec.Command(binaryPath, "--version")
	output, err := cmd.Output()

	if err != nil {
		t.Fatalf("agentkit --version failed: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "agentkit version") {
		t.Error("Version output should contain 'agentkit version'")
	}

	if !strings.Contains(outputStr, "0.1.0") {
		t.Error("Version output should contain version number")
	}
}

func buildAgentKitBinary(t *testing.T) string {
	// Get the project root (2 levels up from test/integration)
	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		t.Fatalf("Failed to get project root: %v", err)
	}

	// Create temporary binary
	binaryPath := filepath.Join(t.TempDir(), "agentkit")

	// Build the binary
	cmd := exec.Command("go", "build", "-o", binaryPath, "./cmd/go-agent-kit")
	cmd.Dir = projectRoot

	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build agentkit binary: %v", err)
	}

	return binaryPath
}

func verifyTemplateContent(t *testing.T, tempDir string) {
	// Verify specify prompt content
	specifyPath := filepath.Join(tempDir, ".github", "prompts", "specify.prompt.md")
	specifyContent, err := os.ReadFile(specifyPath)
	if err != nil {
		t.Fatalf("Failed to read specify prompt: %v", err)
	}

	specifyStr := string(specifyContent)

	// Should not be placeholder content
	if strings.Contains(specifyStr, "TODO") {
		t.Error("Specify prompt should not contain TODO placeholders")
	}

	// Should contain expected content
	specifyExpected := []string{
		"Specification Assistant",
		"GitHub Copilot Instructions",
		"/specify",
		"project/specification.md",
		"specs/[feature-name].md",
	}

	for _, expected := range specifyExpected {
		if !strings.Contains(specifyStr, expected) {
			t.Errorf("Specify prompt missing expected content: %s", expected)
		}
	}

	// Verify plan prompt content
	planPath := filepath.Join(tempDir, ".github", "prompts", "plan.prompt.md")
	planContent, err := os.ReadFile(planPath)
	if err != nil {
		t.Fatalf("Failed to read plan prompt: %v", err)
	}

	planStr := string(planContent)

	// Should not be placeholder content
	if strings.Contains(planStr, "TODO") {
		t.Error("Plan prompt should not contain TODO placeholders")
	}

	// Should contain expected content
	planExpected := []string{
		"Technical Planning Assistant",
		"GitHub Copilot Instructions",
		"/plan",
		"project/architecture.md",
		"Pre-flight Checks",
		"package.json",
		"go.mod",
	}

	for _, expected := range planExpected {
		if !strings.Contains(planStr, expected) {
			t.Errorf("Plan prompt missing expected content: %s", expected)
		}
	}

	// Verify placeholder files still have TODOs (as expected for Commit 3)
	implementPath := filepath.Join(tempDir, ".github", "prompts", "implement.prompt.md")
	implementContent, err := os.ReadFile(implementPath)
	if err != nil {
		t.Fatalf("Failed to read implement prompt: %v", err)
	}

	if !strings.Contains(string(implementContent), "TODO") {
		t.Error("Implement prompt should still contain TODO placeholder")
	}
}
