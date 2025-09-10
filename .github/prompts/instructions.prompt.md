# Generate GitHub Copilot Custom Instructions

## TASK
Create a `.github/copilot-instructions.md` file that helps GitHub Copilot understand this project: {{.Description}}

## STEP 1: ANALYZE PROJECT
Examine the codebase to understand:

1. **What it is**: Project type, purpose, and main functionality
2. **Tech stack**: Languages, frameworks, key dependencies  
3. **Structure**: How directories and files are organized
4. **Patterns**: Coding conventions, error handling, testing approaches
5. **Unique aspects**: Custom utilities, business logic, architectural decisions

@workspace analyze the repository structure, examine key files, and identify patterns

**Output your findings before proceeding to Step 2.**

## STEP 2: CREATE INSTRUCTIONS FILE
Based on your analysis, create `.github/copilot-instructions.md` with these sections:

### Required Sections:

```markdown
# [Project Name] - Custom Instructions for GitHub Copilot

## Project Overview
[1-2 paragraphs: What this project does and its main purpose]

## Code Conventions
### Patterns We Follow
[List specific patterns with code examples from this project]

### What to Avoid  
[Anti-patterns and what to use instead]

## Project Structure
[Key directories and what they contain]
[Important files and their purposes]

## Development Workflow
### Common Commands
[Build, test, run commands specific to this project]

### Testing Requirements
[How we write and organize tests]

## Technical Context
[Key architectural decisions and why]
[Performance or security requirements]
[Preferred libraries for specific tasks]

---
*Last updated: [Date]*
```

### Guidelines for Good Instructions:
- **Be specific**: Focus on what's unique to THIS project
- **Show examples**: Include actual code patterns from the codebase
- **Set boundaries**: Clarify what NOT to do
- **Keep it concise**: Only include information that affects code generation
- **Make it actionable**: Every section should help Copilot write better code

## STEP 3: VALIDATE
Ensure the instructions:
- [ ] Contain project-specific information (not generic advice)
- [ ] Include real examples from the codebase
- [ ] Would help a new developer understand unique patterns
- [ ] Are clear and actionable
- [ ] Are under 500 lines total

## SUCCESS CRITERIA
The generated instructions file should:
1. Help Copilot generate code that matches project style
2. Prevent common mistakes through clear anti-patterns
3. Be maintainable as the project evolves
4. Focus on what makes this project unique

---
**Note**: Generate instructions based on the actual codebase analysis, not generic templates. Every project is different - capture what makes this one special.