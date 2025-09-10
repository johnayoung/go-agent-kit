# MVP Execution Plan: Language-Agnostic Agent Kit

## Starting Point
Project already initialized with:
- Standard Go project structure
- `cmd/go-agent-kit/main.go`
- Makefile
- README.md
- `.gitignore`

## Phase 1: Core Template System (Day 1)

### Commit 1: Create language-agnostic feat template
**Files:**
- `/prompts/feat.md`
```markdown
# Feature Implementation Workflow

## STAGE 1: CODEBASE ANALYSIS
You are analyzing this codebase to implement: {{.Description}}

First, examine the codebase and report:

1. **Detect the project language and framework**
   - Look for: go.mod, package.json, requirements.txt, Gemfile, pom.xml, etc.
   - Identify the primary language and any frameworks

2. **Read and list all relevant files** for this feature
   
3. **Identify existing patterns**:
   - Architecture style (look at folder structure)
   - How similar features are implemented
   - Error handling patterns
   - Testing patterns

4. **Find integration points** where this feature will connect

@workspace examine the project structure and main entry points

Output a summary of what you found. DO NOT write code yet.

## STAGE 2: IMPLEMENTATION PLAN
[Rest of template...]
```

**Commit message:** `feat: add language-agnostic feat prompt template`

### Commit 2: Add template embedding system
**Files:**
- `/internal/templates/embed.go`
```go
package templates

import (
    "embed"
)

//go:embed prompts/*.md
var PromptFiles embed.FS
```

**Commit message:** `feat: add embedded template system`

### Commit 3: Create template renderer
**Files:**
- `/internal/templates/renderer.go`
```go
package templates

import (
    "bytes"
    "text/template"
)

type Context struct {
    Description string
    // Language detected at runtime, not hardcoded
}

func Render(templateName string, ctx Context) (string, error) {
    // Load and execute template
}
```

**Commit message:** `feat: add template rendering engine`

## Phase 2: CLI Command Structure (Day 1)

### Commit 4: Implement main CLI with feat command
**Files:**
- `/cmd/go-agent-kit/main.go` (update existing)
```go
package main

import (
    "fmt"
    "os"
    "github.com/yourusername/go-agent-kit/internal/commands"
)

func main() {
    if len(os.Args) < 2 {
        printUsage()
        os.Exit(1)
    }
    
    switch os.Args[1] {
    case "feat":
        commands.Feat(os.Args[2:])
    case "help":
        printHelp()
    default:
        fmt.Printf("Unknown command: %s\n", os.Args[1])
        printUsage()
        os.Exit(1)
    }
}
```

**Commit message:** `feat: implement CLI with feat command`

### Commit 5: Create feat command handler
**Files:**
- `/internal/commands/feat.go`
```go
package commands

import (
    "fmt"
    "strings"
    "github.com/yourusername/go-agent-kit/internal/templates"
)

func Feat(args []string) error {
    if len(args) == 0 {
        return fmt.Errorf("feat requires a description")
    }
    
    ctx := templates.Context{
        Description: strings.Join(args, " "),
    }
    
    output, err := templates.Render("feat", ctx)
    if err != nil {
        return err
    }
    
    fmt.Println(output)
    return nil
}
```

**Commit message:** `feat: implement feat command with template output`

## Phase 3: Multi-Language Support Templates (Day 2)

### Commit 6: Add language-specific sections to template
**Files:**
- `/prompts/feat.md` (updated)
- `/prompts/partials/go-patterns.md`
- `/prompts/partials/python-patterns.md`
- `/prompts/partials/typescript-patterns.md`

```markdown
## Language-Specific Patterns

Based on the detected language, follow these additional guidelines:

### If Go:
- Use error return pattern (value, error)
- Follow Go module structure
- Use interfaces for abstraction

### If Python:
- Follow PEP 8 style guide
- Use type hints where appropriate
- Create __init__.py for packages

### If TypeScript/JavaScript:
- Use proper TypeScript types (no 'any')
- Follow ESLint rules if present
- Use async/await over callbacks
```

**Commit message:** `feat: add multi-language pattern guidelines`

### Commit 7: Add fix command template
**Files:**
- `/prompts/fix.md`
```markdown
# Bug Fix Workflow

## STAGE 1: DIAGNOSIS
Analyze the reported issue: {{.Description}}

1. Identify the symptoms
2. Locate relevant code sections
3. Form hypothesis about root cause
4. Check for similar patterns that might have the same bug

[Template continues...]
```

**Commit message:** `feat: add language-agnostic fix command template`

## Phase 4: GitHub Integration (Day 2)

### Commit 8: Add install command
**Files:**
- `/internal/commands/install.go`
```go
package commands

import (
    "os"
    "path/filepath"
)

func Install() error {
    // Creates .github/copilot-instructions.md
    instructions := generateInstructions()
    
    githubDir := filepath.Join(".", ".github")
    if err := os.MkdirAll(githubDir, 0755); err != nil {
        return err
    }
    
    return os.WriteFile(
        filepath.Join(githubDir, "copilot-instructions.md"),
        []byte(instructions),
        0644,
    )
}
```

**Commit message:** `feat: add install command for GitHub Copilot integration`

### Commit 9: Wire up install command
**Files:**
- `/cmd/go-agent-kit/main.go` (add install case)

**Commit message:** `feat: connect install command to CLI`

## Phase 5: Testing & Documentation (Day 3)

### Commit 10: Add example outputs
**Files:**
- `/examples/go-project-feat.md`
- `/examples/python-project-feat.md`
- `/examples/typescript-project-feat.md`

**Commit message:** `docs: add language-specific example outputs`

### Commit 11: Update README with language-agnostic focus
**Files:**
- `/README.md`
```markdown
# go-agent-kit

A language-agnostic toolkit for creating structured AI agent workflows. While written in Go, it works with any programming language or framework.

## Features

- 🌍 **Language Agnostic**: Works with Go, Python, TypeScript, Ruby, Java, etc.
- 🤖 **Agent-First**: Designed for AI agents like GitHub Copilot
- 📋 **Structured Workflows**: Enforces best practices through staged processes
- 🔧 **Extensible**: Easy to add new command templates

## Installation

```bash
go install github.com/yourusername/go-agent-kit/cmd/go-agent-kit@latest
```

## Usage

```bash
# In any project (Go, Python, TypeScript, etc.)
go-agent-kit install

# Then in GitHub Copilot:
/feat add user authentication
/fix null pointer in user service
```
```

**Commit message:** `docs: update README for language-agnostic usage`

### Commit 12: Add integration tests
**Files:**
- `/test/integration/feat_test.go`
```go
package integration

import (
    "testing"
    "strings"
)

func TestFeatCommand(t *testing.T) {
    // Test that feat command generates proper workflow
    output := runCommand("feat", "add authentication")
    
    if !strings.Contains(output, "STAGE 1: CODEBASE ANALYSIS") {
        t.Error("Missing Stage 1")
    }
}
```

**Commit message:** `test: add integration tests for feat command`

## Phase 6: Self-Improvement Demo (Day 3-4)

### Commit 13: Use /feat to add /refactor command
**Description:** Actually use the tool's `/feat` command to generate the refactor command
**Files:**
- `/prompts/refactor.md` (generated using feat!)
- `/internal/commands/refactor.go`

**Commit message:** `feat: add refactor command using feat workflow [dogfooding]`

### Commit 14: Add test command template
**Files:**
- `/prompts/test.md`
- `/internal/commands/test.go`

**Commit message:** `feat: add test command for generating comprehensive tests`

## Phase 7: Polish & Release (Day 4)

### Commit 15: Add version command with build info
**Files:**
- `/internal/version/version.go`
```go
package version

var (
    Version = "0.1.0"
    Commit  = "unknown"
    Date    = "unknown"
)
```
- Update Makefile with ldflags

**Commit message:** `feat: add version command with build information`

### Commit 16: Add CI/CD pipeline
**Files:**
- `/.github/workflows/ci.yml`
- `/.github/workflows/release.yml`
- `/.goreleaser.yml`

**Commit message:** `ci: add GitHub Actions for testing and releases`

### Commit 17: Tag initial release
```bash
git tag -a v0.1.0 -m "Initial release: language-agnostic agent workflows"
git push origin v0.1.0
```

**Commit message:** `release: v0.1.0 - language-agnostic agent kit`

---

## Success Metrics

- [ ] Works with Go, Python, and TypeScript projects
- [ ] `/feat` command generates appropriate workflow for any language
- [ ] GitHub Copilot integration works seamlessly
- [ ] Tool can improve itself using its own commands
- [ ] Clear documentation showing multi-language support

## Key Differences from Original Plan

1. **No language detection in CLI** - Let Copilot detect the language
2. **Language-agnostic templates** - Templates work for any language
3. **Focus on patterns, not syntax** - Templates describe what to do, not how
4. **Examples for multiple languages** - Show it works everywhere