package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRefactorCommand(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedError  bool
		expectedInText []string
	}{
		{
			name:          "refactor with single word description",
			args:          []string{"authentication"},
			expectedError: false,
			expectedInText: []string{
				"Code Refactor Workflow",
				"authentication",
				"STAGE 1: CODEBASE ANALYSIS",
			},
		},
		{
			name:          "refactor with multi-word description",
			args:          []string{"simplify", "user", "authentication", "logic"},
			expectedError: false,
			expectedInText: []string{
				"Code Refactor Workflow",
				"simplify user authentication logic",
				"STAGE 1: CODEBASE ANALYSIS",
				"STAGE 2: REFACTOR PLAN",
			},
		},
		{
			name:          "refactor with complex description",
			args:          []string{"extract", "payment", "processing", "into", "separate", "service"},
			expectedError: false,
			expectedInText: []string{
				"Code Refactor Workflow",
				"extract payment processing into separate service",
				"STAGE 1: CODEBASE ANALYSIS",
				"STAGE 2: REFACTOR PLAN",
				"STAGE 3: IMPLEMENTATION",
			},
		},
		{
			name:          "refactor with no arguments",
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
				Use:   "refactor [description]",
				Short: "Generate a code refactoring workflow",
				Args:  cobra.MinimumNArgs(1),
				RunE:  runRefactor,
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

			// Check that all 5 stages are present in refactor workflow
			stages := []string{
				"STAGE 1: CODEBASE ANALYSIS",
				"STAGE 2: REFACTOR PLAN",
				"STAGE 3: IMPLEMENTATION",
				"STAGE 4: TESTING",
				"STAGE 5: DOCUMENTATION",
			}
			for _, stage := range stages {
				if !strings.Contains(output, stage) {
					t.Errorf("Expected to find '%s' in refactor workflow output", stage)
				}
			}
		})
	}
}

func TestRunRefactor(t *testing.T) {
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
			args:        []string{"simplify", "error", "handling"},
			expectError: false,
		},
		{
			name:        "valid complex refactoring description",
			args:        []string{"extract", "database", "layer", "into", "repository", "pattern"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock command
			cmd := &cobra.Command{}

			err := runRefactor(cmd, tt.args)

			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
