package templates

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	tests := []struct {
		name           string
		templateName   string
		context        Context
		expectedError  bool
		expectedInText []string
	}{
		{
			name:         "feat template with description",
			templateName: "feat",
			context: Context{
				Description: "add user authentication",
			},
			expectedError: false,
			expectedInText: []string{
				"Feature Implementation Workflow",
				"add user authentication",
				"STAGE 1: CODEBASE ANALYSIS",
				"STAGE 2: IMPLEMENTATION PLAN",
			},
		},
		{
			name:         "feat template with complex description",
			templateName: "feat",
			context: Context{
				Description: "implement REST API with JWT authentication and rate limiting",
			},
			expectedError: false,
			expectedInText: []string{
				"implement REST API with JWT authentication and rate limiting",
				"Feature Implementation Workflow",
			},
		},
		{
			name:          "nonexistent template",
			templateName:  "nonexistent",
			context:       Context{Description: "test"},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Render(tt.templateName, tt.context)

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

			// Check that expected text is present
			for _, expectedText := range tt.expectedInText {
				if !strings.Contains(result, expectedText) {
					t.Errorf("Expected to find '%s' in rendered output", expectedText)
				}
			}

			// Basic sanity checks
			if len(result) == 0 {
				t.Error("Rendered output should not be empty")
			}
		})
	}
}

func TestContext(t *testing.T) {
	// Test that Context struct works as expected
	ctx := Context{
		Description: "test feature",
	}

	if ctx.Description != "test feature" {
		t.Errorf("Expected Description to be 'test feature', got '%s'", ctx.Description)
	}
}

func TestRenderEmptyDescription(t *testing.T) {
	// Test rendering with empty description
	ctx := Context{
		Description: "",
	}

	result, err := Render("feat", ctx)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should still contain the template structure even with empty description
	expectedStrings := []string{
		"Feature Implementation Workflow",
		"STAGE 1: CODEBASE ANALYSIS",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(result, expected) {
			t.Errorf("Expected to find '%s' in rendered output", expected)
		}
	}
}
