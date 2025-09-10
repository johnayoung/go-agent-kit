# AgentKit Implementation - Detailed Requirements

## Project Context
AgentKit is a CLI tool that creates intelligent prompt templates for GitHub Copilot. Users run `agentkit init` once to set up their project, then use commands like `@workspace /specify` in VS Code's Copilot Chat to build software through specification-driven development.

## Current Status
- ✅ Commit 1: Basic CLI structure exists
- ✅ Commit 2: Init command creates files with placeholders
- ✅ Commit 3: Specify and Plan Prompt Templates with comprehensive content
- 📍 **NEXT: Commit 4** - Implement and Verify prompt templates

---

# Commit 3: Specify and Plan Prompt Templates

## Background Understanding

### What These Prompts Do
The prompts are instructions that tell GitHub Copilot how to behave when users type specific commands. Think of them as "teaching" Copilot to be a specification expert, architect, developer, and QA engineer.

### How They're Used
1. User types `@workspace /specify` in VS Code Copilot Chat
2. Copilot reads the prompt template we created
3. Copilot follows those instructions to help the user
4. Copilot creates actual files in the project (not just chat responses)

## Requirements for the Specify Prompt

### Core Purpose
The `/specify` command helps users define WHAT they want to build without any technical details. It's like a business analyst gathering requirements.

### Key Behaviors Required

#### Context Detection
- The prompt must teach Copilot to detect whether the user wants a project-level specification or a feature specification
- If user types just `/specify` → create project specification
- If user types `/specify user-authentication` → create feature specification

#### File Creation
- Must instruct Copilot to create actual markdown files, not just output text
- Project specs go in `project/specification.md`
- Feature specs go in `specs/[feature-name].md`

#### Content Generation
The prompt must guide Copilot to generate:

**For Project Specifications:**
- What problem does this solve?
- Who are the users?
- What are the main features (high-level)?
- How do we measure success?
- What are we NOT building?

**Critical: Question Generation**
- Must generate 10-20 questions that need answers before building
- Questions like: "Can users have multiple accounts?", "What happens when X fails?", "Is there a limit to Y?"
- These questions prevent incomplete specifications

**For Feature Specifications:**
- What is this specific feature?
- User stories with acceptance criteria
- Business rules and logic
- How it connects to other features
- Feature-specific questions

### Strict Rules to Enforce
- ZERO technical details (no mention of databases, APIs, frameworks)
- Everything from the user's perspective
- Must be explicit and detailed about behavior
- Must create actual files, not just chat output

## Requirements for the Plan Prompt

### Core Purpose
The `/plan` command helps users define HOW to build what was specified. It's like a technical architect designing the system.

### Key Behaviors Required

#### Pre-flight Checks
The prompt must teach Copilot to:
1. Check if `project/specification.md` exists (can't plan without specs)
2. Look for unanswered questions in the specification
3. Refuse to proceed if questions aren't answered
4. Detect existing tech stack by looking for config files

#### Tech Stack Detection
Must check for these files to understand existing technology:
- `package.json` → Node.js project
- `go.mod` → Go project
- `requirements.txt` → Python project
- `Gemfile` → Ruby project
- And others...

#### Two Modes of Operation

**Mode 1: User Specifies Stack**
- If user types `/plan React, Node.js, PostgreSQL`
- Use exactly what they specified

**Mode 2: AI Chooses Stack**
- If user just types `/plan`
- Analyze the specification to understand needs
- Choose appropriate modern technologies
- Justify every choice

#### Architecture Document Creation
Must create `project/architecture.md` containing:
- Technology choices with justifications
- System design (monolith vs microservices vs serverless)
- Database strategy
- API approach
- Security architecture
- Deployment strategy
- Development workflow

### Decision Criteria for Tech Selection
When AI chooses, consider:
- Project complexity from specification
- Performance requirements
- Team size (if mentioned)
- Time to market
- Maintenance needs

## Requirements for Copilot Instructions

### Core Purpose
This is the "master document" that teaches Copilot about the entire workflow.

### Must Include
- Clear explanation of all 4 commands
- The proper sequence (specify → plan → implement → verify)
- How commands work together
- Project structure that gets created
- Key principles like "specification first" and "incremental development"

### Important Context
- Commands can be run multiple times
- Each command adapts to project state
- Commands create real files, not just chat output

---

# Commit 4: Implement and Verify Prompt Templates

## Requirements for the Implement Prompt

### Core Purpose
The `/implement` command is the builder - it reads specifications and creates actual code incrementally.

### Key Behaviors Required

#### Intelligent Context Understanding
Must teach Copilot to:
1. Read all specifications to understand what to build
2. Scan existing code to see what's already built
3. Determine the next logical piece to build
4. Avoid duplicating existing code

#### Three Usage Modes

**Mode 1: Automatic** (`/implement`)
- Figure out what to build next
- Usually: auth first, then core models, then APIs, then UI

**Mode 2: Feature-specific** (`/implement shopping-cart`)
- Build the specified feature

**Mode 3: File-specific** (`/implement specs/payment.md`)
- Build from a specific specification file

#### Implementation Strategy
Must internally break work into small chunks:
1. Data models first (database schemas, classes)
2. Business logic (core functionality)
3. API endpoints (routes, controllers)
4. UI components (if applicable)
5. Tests for everything
6. Documentation

#### Progress Tracking
Must create/update `implementation-log.md` showing:
- What's completed (with file paths)
- What's in progress
- What's next
- Files created or modified

#### Language-Specific Patterns
Must adapt to the detected language:
- Go: Follow standard project layout, proper error handling
- Node.js: Use modern ES6+, async/await, proper middleware
- Python: Follow PEP 8, use type hints, write docstrings
- React: Functional components, hooks, TypeScript

### Critical Rules
- Always read specs first
- Never duplicate existing code
- Build in small, working increments
- Create real files
- Include error handling
- Write tests

## Requirements for the Verify Prompt

### Core Purpose
The `/verify` command is the quality checker - it ensures the code is healthy and matches specifications.

### Key Behaviors Required

#### Language Detection
Must automatically detect project type by looking for:
- Config files (package.json, go.mod, etc.)
- Source file extensions
- Project structure

#### Language-Specific Checks
Must run appropriate checks for each language:

**For Go:**
- Build verification (does it compile?)
- Test execution
- Static analysis
- Format checking

**For JavaScript/TypeScript:**
- Build/compilation
- Test execution
- Linting
- Type checking

**For Python:**
- Test execution
- Type checking
- Format checking
- Linting

#### Universal Checks
Regardless of language:
- Search for TODO/FIXME/HACK comments
- Check documentation exists
- Compare implementation against specification
- Basic security checks (hardcoded secrets, obvious vulnerabilities)

#### Report Generation
Must create a clear report with:
- ✅ What's passing
- ⚠️ Warnings (non-blocking issues)
- ❌ Errors (must-fix issues)
- 📊 Coverage metrics
- 📝 Recommendations for improvement
- Specification compliance status

### Actionable Output
- Every issue must include how to fix it
- Prioritize issues by severity
- Suggest next steps based on specification

---

# Commit 5: Polish, Testing, and Documentation

## Testing Requirements

### Unit Tests Needed
Create tests that verify:
- All prompt templates are non-empty
- No TODO placeholders remain
- Templates are sufficient length (>500 characters)
- File writing operations work correctly
- Directory creation handles errors

### Integration Tests Needed
- Full init command flow works
- Force flag overwrites existing files
- Proper error when files exist without force flag
- Cross-platform path handling

### Test Coverage Goals
- Core functionality 100% covered
- Edge cases handled
- Error conditions tested

## Documentation Requirements

### README.md Must Include

#### Installation Section
- How to install from source
- How to install pre-built binaries
- System requirements

#### Usage Section
- Basic usage example
- Complete workflow walkthrough
- Example for new projects
- Example for existing projects

#### How It Works Section
- Explain the specification-driven approach
- Describe what each command does
- Show the project structure created

#### Troubleshooting Section
- Common issues and solutions
- FAQ
- How to report bugs

### In-Code Documentation
- Add comments for complex logic
- Document public functions
- Explain non-obvious decisions

## UX Improvements

### Better Error Messages
- Instead of "file exists", say "AgentKit already initialized. Use --force to reinitialize"
- Include suggestions for fixing problems
- Use color coding consistently

### Better Success Output
- Show example commands users can try next
- Include tips for effective usage
- Make output encouraging and helpful

### Cross-Platform Compatibility
- Handle Windows path separators
- Test on Windows, macOS, and Linux
- Handle different terminal color support

## Performance Considerations
- Binary size should be <10MB
- Init command should complete in <1 second
- Templates should be embedded in binary (no external files needed)

---

# Success Criteria for Complete Implementation

## Functional Requirements
- [ ] `agentkit init` creates all required files
- [ ] All 4 prompts contain comprehensive instructions
- [ ] Prompts guide Copilot to create actual files
- [ ] Each prompt adapts to project context
- [ ] No placeholder content remains

## Quality Requirements
- [ ] All tests pass
- [ ] README is clear and helpful
- [ ] Error messages guide users to solutions
- [ ] Works on all major operating systems
- [ ] Code follows Go best practices

## User Experience
- [ ] User can go from zero to building in <5 minutes
- [ ] Workflow is intuitive and natural
- [ ] Output is helpful and encouraging
- [ ] Tool feels professional and polished

---

# Implementation Notes for the Agent

## Key Principles
1. **The prompts are teachers** - They teach Copilot how to behave
2. **Be comprehensive** - Each prompt needs complete instructions
3. **Create files, not chat** - Prompts must instruct Copilot to create actual files
4. **Context awareness** - Prompts must adapt to existing projects
5. **No technical details in specs** - Maintain separation between WHAT and HOW

## Common Pitfalls to Avoid
- Don't make prompts too short - they need detailed instructions
- Don't forget the question generation in /specify
- Don't let /plan proceed without answered questions
- Don't let /implement duplicate existing code
- Don't make /verify output without actionable fixes

## Testing Your Work
After each commit:
1. Build the project
2. Run the init command
3. Check the generated files contain real instructions
4. Test in VS Code if possible
5. Verify no placeholders remain