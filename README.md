# go-agent-kit

A language-agnostic toolkit for creating structured AI agent workflows. While written in Go, it works with any programming language or framework.

## Features

- 🌍 **Language Agnostic**: Works with Go, Python, TypeScript, Ruby, Java, C#, and more
- 🤖 **Agent-First**: Designed specifically for AI agents like GitHub Copilot
- 📋 **Structured Workflows**: Enforces best practices through systematic staged processes
- 🔧 **Extensible**: Easy to add new command templates and workflows
- ✅ **Battle-Tested**: Comprehensive test coverage and robust error handling

## Installation

### From Source
```bash
git clone https://github.com/johnayoung/go-agent-kit.git
cd go-agent-kit
go build -o go-agent-kit ./cmd/go-agent-kit/
```

### Future: Go Install (when published)
```bash
go install github.com/johnayoung/go-agent-kit/cmd/go-agent-kit@latest
```

## Quick Start

1. **Install integration in your project** (any language):
   ```bash
   go-agent-kit install
   ```

2. **Use in GitHub Copilot Chat**:
   ```
   /feat add user authentication system
   /fix null pointer exception in service layer
   ```

3. **Follow the generated workflows** step by step for systematic development

## Commands

### `go-agent-kit feat [description]`
Generate a comprehensive feature implementation workflow with 5 stages:
1. **Codebase Analysis** - Detect language, examine patterns, find integration points
2. **Implementation Plan** - Plan files, dependencies, and implementation order  
3. **Implementation** - Step-by-step coding with language-specific best practices
4. **Testing** - Unit tests, integration tests, and edge cases
5. **Documentation** - Code comments, README updates, and API docs

**Examples:**
```bash
go-agent-kit feat "add user authentication"
go-agent-kit feat "implement REST API with JWT"
go-agent-kit feat "add file upload functionality"
```

### `go-agent-kit fix [description]`
Generate a systematic bug fix workflow with 5 stages:
1. **Diagnosis** - Understand, locate, reproduce, and analyze the issue
2. **Fix Strategy** - Plan the fix approach and assess impact
3. **Implementation** - Apply minimal fix with safety checks
4. **Testing** - Verify fix and run regression tests
5. **Documentation** - Document the fix and add preventive measures

**Examples:**
```bash
go-agent-kit fix "null pointer exception in user service"
go-agent-kit fix "memory leak in background worker"
go-agent-kit fix "authentication not working on mobile"
```

### `go-agent-kit install`
Install GitHub Copilot integration files in your project:
- Creates `.github/copilot-instructions.md` with usage instructions
- Enables `/feat` and `/fix` commands in GitHub Copilot Chat
- Works with any programming language or framework

## Language Support

The workflows automatically detect and provide guidance for:

| Language | Detection | Best Practices | Testing Patterns |
|----------|-----------|----------------|------------------|
| **Go** | `go.mod` | Error patterns, interfaces, Go modules | `_test.go` files, table tests |
| **Python** | `requirements.txt`, `pyproject.toml` | PEP 8, type hints, packages | `pytest`, `unittest` |
| **TypeScript/JavaScript** | `package.json`, `tsconfig.json` | ESLint rules, async/await, types | Jest, Mocha, Cypress |
| **Java** | `pom.xml`, `build.gradle` | Design patterns, exceptions | JUnit, TestNG |
| **C#** | `*.csproj`, `*.sln` | LINQ, async/await, IDisposable | MSTest, NUnit, xUnit |
| **Ruby** | `Gemfile`, `*.gemspec` | Ruby style guide, Rails conventions | RSpec, Minitest |

## How It Works

1. **Language Detection**: Workflows examine your project files to detect the primary language and framework
2. **Pattern Analysis**: AI analyzes your existing codebase to understand your specific patterns and architecture
3. **Guided Implementation**: Each stage provides specific guidance while respecting your project's established conventions
4. **Best Practices**: Language-specific guidelines ensure code follows community standards

## Integration with GitHub Copilot

After running `go-agent-kit install`, you can use these commands directly in GitHub Copilot Chat:

```
/feat add user authentication system
/fix database connection timeout errors
```

The AI will:
1. Generate the appropriate workflow for your request
2. Follow each stage systematically  
3. Examine your codebase when prompted with `@workspace`
4. Implement following your project's patterns and language conventions

## Benefits

- ✅ **Consistent Quality**: Every feature and fix follows the same systematic approach
- ✅ **Language Agnostic**: Works across all programming languages and frameworks
- ✅ **Best Practices**: Incorporates language-specific conventions and patterns  
- ✅ **Comprehensive**: Covers analysis, implementation, testing, and documentation
- ✅ **AI-Optimized**: Designed specifically for AI agents like GitHub Copilot
- ✅ **Extensible**: Easy to add new workflows and customize for your needs

## Development

### Prerequisites
- Go 1.22 or later

### Building
```bash
go mod tidy
go build ./cmd/go-agent-kit/
```

### Testing
```bash
go test ./...                    # Run all tests
go test -v ./internal/cmd/       # Run CLI tests with verbose output
go test -v ./internal/templates/ # Run template tests with verbose output
```

### Code Quality
```bash
go vet ./...     # Static analysis
go fmt ./...     # Format code
```

## Project Structure

```
go-agent-kit/
├── cmd/go-agent-kit/           # CLI entry point
├── internal/
│   ├── cmd/                    # CLI commands (feat, fix, install)
│   └── templates/              # Template system
│       ├── prompts/            # Workflow templates
│       ├── embed.go            # Embedded file system
│       └── renderer.go         # Template rendering
├── .github/
│   └── copilot-instructions.md # GitHub Copilot integration
└── docs/                       # Documentation
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes following the existing patterns
4. Add tests for new functionality
5. Ensure all tests pass (`go test ./...`)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the need for systematic AI-assisted development workflows
- Built for seamless integration with GitHub Copilot and other AI coding assistants
- Designed to work across all programming languages and frameworks
