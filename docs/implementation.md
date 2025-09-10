# AgentKit CLI - Implementation Commits

## Commit Sequence Checklist

- [x] Commit 1: Project initialization and CLI foundation
- [ ] Commit 2: Init command with file operations
- [ ] Commit 3: Specify and Plan prompt templates
- [ ] Commit 4: Implement and Verify prompt templates  
- [ ] Commit 5: Polish, testing, and documentation

---

## Commit 1: Project initialization and CLI foundation

**Objective**: Set up Go project with CLI framework and basic command structure.

**Requirements**:
- Initialize Go module `github.com/yourusername/agentkit`
- Add Cobra and color dependencies
- Create entry point that executes root command
- Set up root command with name, description, and version
- Add init command stub (just prints "Init command called")
- Add `--force` flag to init command
- Create Makefile with build, install, clean targets
- Ensure project follows Go best practices for structure

**Validation**:
- `go build` succeeds
- `./agentkit --version` shows version
- `./agentkit init --help` shows command help
- `./agentkit init` prints placeholder message

---

## Commit 2: Init command with file operations

**Objective**: Implement the init command to create directory structure and write files.

**Requirements**:
- Create file writer utility that can:
  - Write files with proper permissions (0644)
  - Check if paths exist
  - Create directories recursively
- Implement init command logic:
  - Check if `.github/prompts/` exists
  - Error if exists without `--force` flag
  - Create directory structure
  - Write placeholder files for now (just "TODO: [filename]")
  - Display colored success output
- Use color package for output:
  - Green checkmarks for success
  - Red for errors
  - Yellow for warnings
  - Cyan for command examples

**Files to Create** (with placeholders):
- `.github/copilot-instructions.md`
- `.github/prompts/specify.prompt.md`
- `.github/prompts/plan.prompt.md`
- `.github/prompts/implement.prompt.md`
- `.github/prompts/verify.prompt.md`

**Validation**:
- `agentkit init` creates all files
- `agentkit init` again shows error about existing files
- `agentkit init --force` overwrites files
- Output is colored and informative

---

## Commit 3: Specify and Plan prompt templates

**Objective**: Create the specification and planning prompt templates.

**Requirements**:

### Specify Prompt Template
Must instruct Copilot to:
- Detect if creating project or feature specification
- Create `project/specification.md` for project-level
- Create `specs/[feature-name].md` for features
- Generate comprehensive WHAT and WHY (no technical details)
- Include "Questions Requiring Answers" section
- Generate business logic questions
- Avoid ALL technical implementation details

### Plan Prompt Template  
Must instruct Copilot to:
- Read all specifications first
- Detect existing tech stack from files (package.json, go.mod, etc.)
- Two modes: user-specified stack or LLM-generated
- Create `project/architecture.md`
- Include justifications for all technology choices
- Consider scale, performance, team size when choosing stack

### Copilot Instructions
Global instructions including:
- Overview of the workflow
- How commands work together
- Emphasis on incremental development
- Instructions to detect context

**Validation**:
- Templates are valid markdown
- Templates contain clear instructions
- No placeholder content remains
- Instructions are comprehensive

---

## Commit 4: Implement and Verify prompt templates

**Objective**: Create the implementation and verification prompt templates.

**Requirements**:

### Implement Prompt Template
Must instruct Copilot to:
- Read all specifications and architecture
- Scan existing code to understand what's built
- Three modes: no args (next logical thing), feature name, specific file
- Determine optimal build sequence automatically
- Break work into small increments internally
- Create actual code files
- Update `implementation-log.md` with progress
- Never duplicate existing code
- Respect existing patterns

### Verify Prompt Template
Must instruct Copilot to:
- Detect project language/framework automatically
- Run appropriate checks per language:
  - Python: mypy, pytest, ruff
  - JavaScript/TypeScript: tsc, jest, eslint
  - Go: go build, go test, go vet
- Check for TODO/FIXME comments
- Compare implementation against specifications
- Format output with clear sections (passing, warnings, errors)
- Provide actionable fix suggestions

**Validation**:
- Templates provide complete instructions
- Language detection logic is comprehensive
- All major languages covered
- Output format is clearly specified

---

## Commit 5: Polish, testing, and documentation

**Objective**: Add tests, improve UX, and create documentation.

**Requirements**:
- Add unit tests for:
  - File writing operations
  - Path existence checks
  - Directory creation
- Add integration test for full init flow
- Create comprehensive README.md with:
  - Installation instructions
  - Usage examples for all scenarios
  - Workflow explanation
  - Troubleshooting section
- Improve error messages to be more helpful
- Add examples to success output
- Ensure cross-platform compatibility (Windows paths)
- Add `.gitignore` for Go projects
- Add GitHub Actions workflow for CI (optional)

**Validation**:
- `go test ./...` passes
- README clearly explains the tool
- Error messages guide users to solutions
- Works on Windows, macOS, and Linux
- Binary size is reasonable (< 10MB)

---

## Implementation Notes

- Each commit should be independently functional
- Run `go fmt` and `go vet` before each commit
- Keep commits focused - resist adding extra features
- Use embedded strings for templates (no external files)
- Follow Go conventions for error handling
- Use clear variable and function names
- Add comments only where behavior isn't obvious

## Success Criteria

After all commits:
- [ ] Single binary that works everywhere
- [ ] One command creates complete prompt structure  
- [ ] Prompts are intelligent and self-contained
- [ ] Clear documentation and examples
- [ ] Tested and reliable
- [ ] Total implementation time: 3-4 hours