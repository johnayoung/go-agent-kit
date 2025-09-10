package templates

import (
	"bytes"
	"fmt"
	"text/template"
)

// Context holds the data to be passed to templates
type Context struct {
	Description string
	// Language detected at runtime, not hardcoded
}

// Render loads and executes a template with the given context
func Render(templateName string, ctx Context) (string, error) {
	// Construct the file path
	templatePath := fmt.Sprintf("prompts/%s.md", templateName)
	
	// Read the template file from embedded FS
	templateContent, err := PromptFiles.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template %s: %w", templateName, err)
	}
	
	// Parse the template
	tmpl, err := template.New(templateName).Parse(string(templateContent))
	if err != nil {
		return "", fmt.Errorf("failed to parse template %s: %w", templateName, err)
	}
	
	// Execute the template with the context
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, ctx)
	if err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", templateName, err)
	}
	
	return buf.String(), nil
}
