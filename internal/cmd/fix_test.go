package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestFixCommand(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedError  bool
		expectedInText []string
	}{
		{
			name:          "fix with single word description",
			args:          []string{"crash"},
			expectedError: false,
			expectedInText: []string{
				"Bug Fix Workflow",
				"crash",
				"STAGE 1: DIAGNOSIS",
			},
		},
		{
			name:          "fix with multi-word description",
			args:          []string{"null", "pointer", "exception", "in", "service"},
			expectedError: false,
			expectedInText: []string{
				"Bug Fix Workflow",
				"null pointer exception in service",
				"STAGE 1: DIAGNOSIS",
				"STAGE 2: FIX STRATEGY",
			},
		},
		{
			name:          "fix with no arguments",
			args:          []string{},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			var buf bytes.Buffer

			// Create a new command instance for testing
			cmd := &cobra.Command{
				Use:   "fix [description]",
				Short: "Generate a bug fix workflow",
				Args:  cobra.MinimumNArgs(1),
				RunE:  runFix,
			}

			// Set output to our buffer
			cmd.SetOut(&buf)
			cmd.SetErr(&buf)
			cmd.SetArgs(tt.args)

			// Execute the command
			err := cmd.Execute()

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
			output := buf.String()
			for _, expectedText := range tt.expectedInText {
				if !strings.Contains(output, expectedText) {
					t.Errorf("Expected to find '%s' in output", expectedText)
				}
			}

			// Basic sanity checks
			if len(output) == 0 {
				t.Error("Command output should not be empty")
			}
		})
	}
}

func TestRunFix(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
	}{
		{
			name:        "valid single argument",
			args:        []string{"crash"},
			expectError: false,
		},
		{
			name:        "valid multiple arguments",
			args:        []string{"memory", "leak", "issue"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock command
			cmd := &cobra.Command{}
			
			err := runFix(cmd, tt.args)
			
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
