package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestFeatCommand(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedError  bool
		expectedInText []string
	}{
		{
			name: "feat with single word description",
			args: []string{"authentication"},
			expectedError: false,
			expectedInText: []string{
				"Feature Implementation Workflow",
				"authentication",
				"STAGE 1: CODEBASE ANALYSIS",
			},
		},
		{
			name: "feat with multi-word description",
			args: []string{"add", "user", "authentication", "system"},
			expectedError: false,
			expectedInText: []string{
				"Feature Implementation Workflow",
				"add user authentication system",
				"STAGE 1: CODEBASE ANALYSIS",
				"STAGE 2: IMPLEMENTATION PLAN",
			},
		},
		{
			name:          "feat with no arguments",
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
				Use:   "feat [description]",
				Short: "Generate a feature implementation workflow",
				Args:  cobra.MinimumNArgs(1),
				RunE:  runFeat,
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

func TestRunFeat(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
	}{
		{
			name:        "valid single argument",
			args:        []string{"authentication"},
			expectError: false,
		},
		{
			name:        "valid multiple arguments",
			args:        []string{"add", "user", "management"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock command
			cmd := &cobra.Command{}
			
			err := runFeat(cmd, tt.args)
			
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
