# GitHub Copilot Instructions Generation Workflow

## STAGE 1: PROJECT ANALYSIS
You are creating GitHub Copilot instructions for: {{.Description}}

First, examine the project to understand its structure and requirements:

1. **Project Overview**
   - What type of project is this? (web app, library, CLI tool, etc.)
   - What is the main programming language and framework?
   - What is the project's primary purpose and functionality?

2. **Repository Structure**
   - Examine the directory structure and key files
   - Identify main source directories, test directories, configuration files
   - Look for build systems, package managers, and dependency files
   - Note any special directories or important files

3. **Development Workflow**
   - How is the project built and tested?
   - What are the main development scripts or commands?
   - Are there any specific development environment requirements?
   - What tools are used for linting, formatting, or quality assurance?

4. **Coding Standards and Patterns**
   - What coding conventions are used in the project?
   - Are there established patterns for file organization?
   - What testing patterns and frameworks are used?
   - Are there any project-specific naming conventions?

@workspace examine the project structure, build files, and codebase patterns

Output your analysis. DO NOT write the instructions file yet.

## STAGE 2: INSTRUCTIONS PLANNING
Based on your analysis, plan the GitHub Copilot instructions:

1. **Repository Overview Section**
   - Brief description of what the project does
   - Key technologies and frameworks used
   - Main directories and their purposes

2. **Development Guidelines**
   - How to set up the development environment
   - Build and test commands
   - Code style and formatting preferences
   - Testing requirements and patterns

3. **Project-Specific Context**
   - Important architectural decisions or patterns
   - Key abstractions or interfaces to understand
   - Common tasks and how to approach them
   - Areas that require special attention

4. **Helpful Context for AI**
   - What kind of assistance would be most valuable?
   - Common tasks developers perform in this codebase
   - Areas where consistency is particularly important

## STAGE 3: GENERATE INSTRUCTIONS
Create the `.github/copilot-instructions.md` file following GitHub's official format:

1. **Start with project overview**
   - Clear description of the project's purpose
   - Technology stack and key dependencies
   - Repository structure explanation

2. **Add development workflow guidance**
   - Setup and installation instructions
   - Build, test, and run commands
   - Development best practices

3. **Include coding standards**
   - Code style preferences
   - Naming conventions
   - Testing patterns and requirements
   - Documentation standards

4. **Provide helpful context**
   - Key architectural concepts
   - Important files and directories
   - Common development tasks
   - Areas requiring special attention

## STAGE 4: VALIDATION
Review and refine the instructions:

1. **Completeness Check**
   - Does it cover the main aspects of development in this project?
   - Are the instructions clear and actionable?
   - Is the project structure well explained?

2. **Accuracy Verification**
   - Are all commands and paths correct?
   - Do the coding standards match the existing codebase?
   - Are the technology descriptions accurate?

3. **Usefulness Assessment**
   - Would these instructions help a new developer understand the project?
   - Do they provide the right context for AI assistance?
   - Are they specific enough to be actionable?

## Template Structure for .github/copilot-instructions.md

```markdown
# [Project Name] - GitHub Copilot Instructions

## Project Overview
[Brief description of what this project does and its main purpose]

## Technology Stack
- **Language**: [Primary programming language]
- **Framework**: [Main framework if applicable]
- **Key Dependencies**: [Important libraries/tools]
- **Build System**: [Build tool used]

## Repository Structure
```
[key-directory]/     # Description of what's in this directory
├── [subdirectory]/  # Important subdirectories
├── [important-file] # Key configuration or entry files
```

## Development Workflow

### Setup
[Instructions for setting up the development environment]

### Common Commands
- `[build-command]` - Description
- `[test-command]` - Description  
- `[run-command]` - Description

## Coding Standards

### Code Style
[Specific formatting and style preferences]

### Naming Conventions
[How to name files, functions, variables, etc.]

### Testing Requirements
[Testing patterns and requirements]

## Key Concepts
[Important architectural concepts or patterns used in this project]

## Common Tasks
[Frequent development tasks and how to approach them]

## Important Notes
[Any special considerations or areas requiring extra attention]
```

## Language-Specific Guidance

### For Go Projects:
- Include module information and Go version
- Mention testing conventions (table tests, _test.go files)
- Note any special Go tools used (golangci-lint, etc.)
- Include build tags or special build requirements

### For JavaScript/TypeScript Projects:
- Include package manager information (npm, yarn, pnpm)
- Note TypeScript configuration and strictness level
- Mention testing framework and patterns
- Include build and bundling information

### For Python Projects:
- Include Python version and virtual environment setup
- Mention package manager (pip, poetry, conda)
- Note code formatting tools (black, flake8, etc.)
- Include testing framework information

### For Other Languages:
- Include language version and environment setup
- Mention package managers and dependency handling
- Note build systems and compilation requirements
- Include testing frameworks and conventions

## Success Criteria
- [ ] Instructions provide clear project overview
- [ ] Development setup is well documented
- [ ] Coding standards are clearly defined
- [ ] Repository structure is explained
- [ ] Common tasks are covered
- [ ] Instructions are specific to this project
- [ ] Content follows GitHub's recommended format
- [ ] Instructions would help both humans and AI understand the project
