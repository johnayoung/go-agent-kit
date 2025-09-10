# Continuation Prompt: Next Implementation Step

## 🎯 Your Task
Continue implementation by completing the next unchecked commit in `docs/implementation.md`.

## 📋 Pre-Implementation Checklist

**STOP and verify ALL of the following before proceeding:**

### 1. Check Current State
- [ ] Read `docs/implementation.md` to identify the next unchecked commit
- [ ] Verify all previous commits are checked off
- [ ] Confirm you understand the requirements for the next commit

### 2. Stability Verification
Run these checks based on project type:

#### For Go Projects:
```bash
go build ./...           # Compilation check
go test ./...           # Run all tests
go vet ./...           # Static analysis
go fmt ./...           # Format check
```

#### For Python Projects:
```bash
python -m pytest        # Run tests
python -m mypy .       # Type checking
python -m ruff check . # Linting
```

#### For JavaScript/TypeScript:
```bash
npm test               # Run tests
npm run build         # Build check
npm run lint          # Linting
```

### 3. Repository State
- [ ] No uncommitted changes blocking progress
- [ ] Previous commit is complete and working
- [ ] All added files are tracked in git

## 🚀 Implementation Process

### Step 1: Identify Next Commit
Look at `docs/implementation.md` and find the first unchecked item:
- [ ] Commit 1: Project initialization and CLI foundation
- [ ] Commit 2: Init command with file operations
- [ ] Commit 3: Specify and Plan prompt templates
- [ ] Commit 4: Implement and Verify prompt templates
- [ ] Commit 5: Polish, testing, and documentation

### Step 2: Read Requirements
Carefully read the requirements for that commit. Pay attention to:
- **Objective**: What should be working after this commit
- **Requirements**: Specific features to implement
- **Validation**: How to test that it works

### Step 3: Implement
- Follow project conventions and best practices
- Keep changes focused on the commit objectives
- Don't add features from future commits
- Ensure each file follows language idioms

### Step 4: Validate
Run the validation steps listed for the commit:
- Test the specific features added
- Ensure nothing from previous commits broke
- Verify the success criteria

### Step 5: Update Checklist
Once validated and working:
1. Update the checkbox in `docs/implementation.md`
2. Commit your changes with a clear message
3. Report completion status

## 📝 Completion Report Format

After completing a commit, report:

```
✅ Completed: Commit [NUMBER] - [TITLE]

Changes made:
- [Brief list of key changes]

Validation results:
- [What was tested]
- [What's working]

Ready for: Commit [NEXT NUMBER] - [NEXT TITLE]
```

## ⚠️ If You Encounter Issues

If stability checks fail or you find problems:

1. **STOP immediately**
2. **Fix the issues** before proceeding
3. **Re-run stability checks**
4. **Document what was fixed**

Report format for issues:
```
🔧 Fixed issues before continuing:
- [Issue found]: [How it was fixed]
- Tests now passing: [YES/NO]
- Build successful: [YES/NO]
```

## 🎯 Current Focus

**Your next step**: Run stability checks, then implement the next unchecked commit from the implementation document.

**Remember**: 
- Complete only ONE commit
- Validate thoroughly before marking complete
- Update the checklist after success
- Stop and wait for further instructions

---

*Begin by running the stability checks, then proceed with the next commit.*