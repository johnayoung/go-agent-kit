# go-agent-kit

Install GitHub Copilot integration for structured AI agent workflows. Works with any programming language or framework.

## What it does

go-agent-kit installs GitHub Copilot prompt files that enable you to use structured AI workflows directly in GitHub Copilot Chat:

- � **`/feat [description]`** - Generate feature implementation workflows
- 🔧 **`/fix [description]`** - Generate systematic bug fix workflows  
- ♻️ **`/refactor [description]`** - Generate code improvement workflows

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

## Usage

### 1. Install GitHub Copilot Integration
Run this command in any project directory:

```bash
go-agent-kit install
```

This creates:
- `.github/copilot-instructions.md` - Instructions for your team
- `.github/prompts/feat.prompt.md` - Feature workflow template
- `.github/prompts/fix.prompt.md` - Bug fix workflow template  
- `.github/prompts/refactor.prompt.md` - Refactoring workflow template

### 2. Use in GitHub Copilot Chat

Open GitHub Copilot Chat and use the workflows:

```
/feat add user authentication system
/fix null pointer exception in service layer
/refactor simplify payment processing logic
```

Each command generates a comprehensive 5-stage workflow:
1. **Analysis** - Understand the codebase and requirements
2. **Planning** - Create detailed implementation strategy
3. **Implementation** - Step-by-step coding with best practices
4. **Testing** - Comprehensive testing and validation
5. **Documentation** - Proper documentation and comments

## Language Support

The workflows automatically detect and provide guidance for:

| Language | Framework Support | Best Practices |
|----------|------------------|----------------|
| **Go** | Gin, Echo, Fiber | Error handling, interfaces, testing |
| **Python** | Django, Flask, FastAPI | PEP 8, type hints, testing |
| **TypeScript** | React, Next.js, Express | ESLint, async/await, types |
| **JavaScript** | Vue, Angular, Node.js | Modern ES6+, testing |
| **Java** | Spring, Spring Boot | Design patterns, testing |
| **C#** | ASP.NET, .NET Core | LINQ, async/await, testing |
| **Ruby** | Rails, Sinatra | Style guide, conventions |

## Example Workflows

### Feature Implementation
```
/feat add JWT authentication middleware
```
Generates a workflow that:
- Analyzes your project structure and patterns
- Plans the implementation with security best practices
- Guides through step-by-step implementation
- Creates comprehensive tests
- Documents the new functionality

### Bug Fixing
```
/fix memory leak in background worker
```
Generates a workflow that:
- Systematically diagnoses the issue
- Plans the safest fix approach
- Implements minimal, targeted changes
- Validates the fix with testing
- Documents the solution

### Code Refactoring
```
/refactor extract database layer into repository pattern
```
Generates a workflow that:
- Analyzes current code structure
- Plans refactoring strategy and risks
- Applies changes systematically
- Ensures all tests still pass
- Updates documentation

## How It Works

1. **Language Detection**: Automatically detects your project's language and framework
2. **Pattern Analysis**: Understands your existing code patterns and conventions  
3. **Systematic Guidance**: Provides step-by-step instructions following best practices
4. **Quality Assurance**: Ensures proper testing and documentation

## Benefits

✅ **Consistent Quality** - Every change follows the same systematic approach  
✅ **Language Agnostic** - Works across all programming languages and frameworks  
✅ **Best Practices** - Incorporates community standards and conventions  
✅ **Comprehensive** - Covers analysis, implementation, testing, and documentation  
✅ **AI-Optimized** - Designed specifically for GitHub Copilot integration  
✅ **Team-Friendly** - Share consistent workflows across your entire team

## Project Structure

```
go-agent-kit/
├── cmd/go-agent-kit/           # CLI entry point
├── internal/
│   ├── cmd/                    # Install command
│   └── templates/              # Workflow templates
└── .github/
    ├── copilot-instructions.md # GitHub Copilot integration
    └── prompts/                # Workflow prompt files
```

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
go test ./...     # Run all tests
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
