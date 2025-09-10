package filewriter

import (
	"os"
	"path/filepath"
)

// FileWriter handles writing files with proper permissions and directory creation
type FileWriter struct{}

// New creates a new FileWriter instance
func New() *FileWriter {
	return &FileWriter{}
}

// WriteFile writes content to a file with proper permissions (0644)
// It creates directories recursively if they don't exist
func (fw *FileWriter) WriteFile(filePath, content string) error {
	// Create directories if they don't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Write file with proper permissions
	return os.WriteFile(filePath, []byte(content), 0644)
}

// PathExists checks if a file or directory exists
func (fw *FileWriter) PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDirectory checks if the given path is a directory
func (fw *FileWriter) IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
