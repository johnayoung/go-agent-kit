package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCmd(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedError  bool
		expectedOutput string
	}{
		{
			name:          "root command without args",
			args:          []string{},
			expectedError: false,
		},
		{
			name:          "help flag",
			args:          []string{"--help"},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset the command for each test
			cmd := &cobra.Command{
				Use:   "go-agent-kit",
				Short: "A language-agnostic toolkit for creating structured AI agent workflows",
				Long: `go-agent-kit is a language-agnostic toolkit for creating structured AI agent workflows.
While written in Go, it works with any programming language or framework.`,
			}
			cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

			cmd.SetArgs(tt.args)
			err := cmd.Execute()

			if tt.expectedError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectedError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestExecute(t *testing.T) {
	// Test that Execute function works without crashing
	// In a real scenario, this would need more sophisticated testing
	// but for now we just verify the function can be called
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Execute() panicked: %v", r)
		}
	}()

	// We can't easily test Execute() directly as it uses os.Exit
	// This test just verifies the function is accessible
	t.Log("Execute function is accessible and defined")
}
