# go-agent-kit - Custom Instructions for GitHub Copilot

## Project Overview

go-agent-kit is a CLI tool that installs GitHub Copilot integration files for structured AI workflows. It creates language-agnostic prompt templates that enable systematic development approaches using `/feat`, `/fix`, and `/refactor` commands in GitHub Copilot Chat. The tool works across any programming language by detecting project patterns and applying appropriate conventions.

The core philosophy is to provide structured, 5-stage workflows that guide AI assistants through comprehensive development processes, from analysis to documentation, while respecting each project's unique patterns and conventions.

## Code Conventions

### Patterns We Follow

**Error Handling - Go Idiomatic Wrapping:**
```go
if err := os.MkdirAll(githubDir, 0755); err != nil {
    return fmt.Errorf("failed to create .github directory: %w", err)
}
```

**Table-Driven Tests:**
```go
tests := []struct {
    name           string
    expectedError  bool
    checkFiles     []string
    expectedInText []string
}{
    {
        name:          "install command creates .github directory and files",
        expectedError: false,
        checkFiles: []string{
            ".github/copilot-instructions.md",
            ".github/prompts/feat.prompt.md",
        },
    },
}
```

**Cobra Command Registration:**
```go
var installCmd = &cobra.Command{
    Use:   "install",
    Short: "Install GitHub Copilot integration files",
    Long:  `Detailed description...`,
    RunE:  runInstall,
}

func init() {
    rootCmd.AddCommand(installCmd)
}
```

**Template Embedding with embed.FS:**
```go
//go:embed prompts/*.md
var PromptFiles embed.FS

content, err := templates.PromptFiles.ReadFile("prompts/feat.md")
```

### What to Avoid

- **Don't use os.Exit() in commands** - Use error returns instead: `return fmt.Errorf("...")`
- **Don't hardcode file paths** - Use `filepath.Join()`: `filepath.Join(githubDir, "prompts")`
- **Don't skip error wrapping** - Always provide context: `fmt.Errorf("failed to X: %w", err)`
- **Don't create files without proper permissions** - Use `0755` for directories, `0644` for files
- **Don't write unit tests without table-driven structure** - Follow the established pattern

## Project Structure

```
cmd/go-agent-kit/          # Main CLI entry point
internal/
  cmd/                     # Command implementations (install, root)
  templates/               # Embedded template system
    embed.go              # Template embedding with embed.FS
    prompts/              # Markdown workflow templates
      feat.md             # Feature implementation workflow
      fix.md              # Bug fix workflow  
      refactor.md         # Code refactoring workflow
      instructions.md     # Instructions generation workflow
```

**Key Files:**
- `internal/cmd/install.go` - Core installation logic, file creation, instruction generation
- `internal/templates/embed.go` - Template embedding system using Go's embed.FS
- `cmd/go-agent-kit/main.go` - Simple main entry point calling `cmd.Execute()`

## Development Workflow

### Common Commands
```bash
# Build the CLI
make build

# Run tests
make test

# Clean build artifacts
make clean

# Build and run
make run

# Tidy dependencies
go mod tidy

# Cross-compile for Linux
make build-linux
```

### Testing Requirements

**Test Structure - Follow Table-Driven Pattern:**
```go
func TestInstallCommand(t *testing.T) {
    tests := []struct {
        name           string
        expectedError  bool
        checkFiles     []string
        expectedInText []string
    }{
        // Test cases here
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

**Test Organization:**
- One test file per source file: `install.go` → `install_test.go`
- Use temporary directories for file system tests
- Clean up resources with `defer os.RemoveAll(tempDir)`
- Test both success and error cases

## Technical Context

### Template System Architecture
Templates use Go's `text/template` syntax with `{{.Description}}` placeholders. The embed.FS system allows templates to be compiled into the binary while remaining editable during development.

### Language-Agnostic Design Philosophy
The workflow templates are designed to work with ANY programming language by:
1. **Language Detection** - Examining project files (go.mod, package.json, etc.)
2. **Pattern Analysis** - Understanding existing codebase conventions
3. **Adaptive Implementation** - Applying language-specific best practices

### File System Operations
- Always use `filepath.Join()` for cross-platform compatibility
- Create directories with `os.MkdirAll(dir, 0755)`
- Write files with `os.WriteFile(path, content, 0644)`
- Handle permissions explicitly for security

### CLI Design Principles
- Simple command structure: `go-agent-kit install`
- Rich help text with examples and context
- Informative success messages with next steps
- Error messages that guide users toward solutions

### Performance Considerations
- Templates are embedded at compile-time for fast access
- File operations are batched in single `install` command
- No external dependencies beyond standard library + Cobra

### Preferred Libraries
- **CLI**: `github.com/spf13/cobra` for command structure
- **Colors**: `github.com/fatih/color` for terminal styling
- **File Operations**: Standard library `os`, `filepath` packages
- **Templates**: Standard library `embed` and `text/template`

---
*Last updated: September 10, 2025*
