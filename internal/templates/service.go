package templates

import (
	"embed"
	"strings"
	"text/template"
)

//go:embed prompts
var templatesFS embed.FS

// TemplateService handles loading and executing templates
type TemplateService struct {
	templates *template.Template
}

// New creates a new TemplateService
func New() (*TemplateService, error) {
	tmpl, err := template.ParseFS(templatesFS, "prompts/*.md")
	if err != nil {
		return nil, err
	}

	return &TemplateService{
		templates: tmpl,
	}, nil
}

// GetSpecifyPrompt returns the specify prompt template content
func (ts *TemplateService) GetSpecifyPrompt() (string, error) {
	var buf strings.Builder
	err := ts.templates.ExecuteTemplate(&buf, "specify.prompt.md", nil)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// GetPlanPrompt returns the plan prompt template content
func (ts *TemplateService) GetPlanPrompt() (string, error) {
	var buf strings.Builder
	err := ts.templates.ExecuteTemplate(&buf, "plan.prompt.md", nil)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
