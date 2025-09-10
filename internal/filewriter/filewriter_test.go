package filewriter

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNew(t *testing.T) {
	fw := New()
	if fw == nil {
		t.Error("New() should return a non-nil FileWriter")
	}
}

func TestWriteFile(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "filewriter_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	fw := New()

	// Test writing a file
	testFile := filepath.Join(tempDir, "test", "subdir", "file.txt")
	testContent := "test content"

	err = fw.WriteFile(testFile, testContent)
	if err != nil {
		t.Errorf("WriteFile() failed: %v", err)
	}

	// Verify file exists and has correct content
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("Failed to read created file: %v", err)
	}

	if string(content) != testContent {
		t.Errorf("File content mismatch. Expected: %q, Got: %q", testContent, string(content))
	}

	// Verify file permissions
	info, err := os.Stat(testFile)
	if err != nil {
		t.Errorf("Failed to stat file: %v", err)
	}

	expectedMode := os.FileMode(0644)
	if info.Mode().Perm() != expectedMode {
		t.Errorf("File permissions mismatch. Expected: %v, Got: %v", expectedMode, info.Mode().Perm())
	}
}

func TestPathExists(t *testing.T) {
	fw := New()

	// Test with existing file
	tempFile, err := os.CreateTemp("", "test_exists")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	if !fw.PathExists(tempFile.Name()) {
		t.Error("PathExists() should return true for existing file")
	}

	// Test with non-existing file
	nonExistentPath := filepath.Join(os.TempDir(), "non_existent_file_12345")
	if fw.PathExists(nonExistentPath) {
		t.Error("PathExists() should return false for non-existing file")
	}
}

func TestIsDirectory(t *testing.T) {
	fw := New()

	// Test with existing directory
	tempDir, err := os.MkdirTemp("", "test_dir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	if !fw.IsDirectory(tempDir) {
		t.Error("IsDirectory() should return true for existing directory")
	}

	// Test with file
	tempFile, err := os.CreateTemp(tempDir, "test_file")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tempFile.Close()

	if fw.IsDirectory(tempFile.Name()) {
		t.Error("IsDirectory() should return false for file")
	}

	// Test with non-existing path
	nonExistentPath := filepath.Join(os.TempDir(), "non_existent_dir_12345")
	if fw.IsDirectory(nonExistentPath) {
		t.Error("IsDirectory() should return false for non-existing path")
	}
}
