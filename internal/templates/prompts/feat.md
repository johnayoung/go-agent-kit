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
Based on your analysis, create a detailed implementation plan:

1. **Files to create/modify** (be specific about paths)
2. **Implementation order** (what to build first, second, etc.)
3. **Dependencies** (what needs to exist before each step)
4. **Integration points** (how this connects to existing code)
5. **Testing strategy** (what to test and how)

## STAGE 3: IMPLEMENTATION
Now implement the feature following your plan:

1. **Start with interfaces/types** if needed
2. **Implement core logic**
3. **Add integration points**
4. **Follow language-specific patterns** you identified
5. **Handle errors appropriately** for the detected language

## STAGE 4: TESTING
Create comprehensive tests:

1. **Unit tests** for core functionality
2. **Integration tests** if applicable
3. **Edge cases** and error conditions
4. **Follow testing patterns** you identified in Stage 1

## STAGE 5: DOCUMENTATION
Add appropriate documentation:

1. **Code comments** following language conventions
2. **README updates** if this affects usage
3. **API documentation** if this is a public interface

## Language-Specific Guidelines

### For Go:
- Use error return pattern (value, error)
- Follow Go module structure
- Use interfaces for abstraction
- Follow effective Go practices

### For Python:
- Follow PEP 8 style guide
- Use type hints where appropriate
- Create __init__.py for packages
- Use docstrings for documentation

### For TypeScript/JavaScript:
- Use proper TypeScript types (avoid 'any')
- Follow ESLint rules if present
- Use async/await over callbacks
- Prefer functional patterns where appropriate

### For Java:
- Follow Java naming conventions
- Use appropriate design patterns
- Handle exceptions properly
- Use proper package structure

### For C#:
- Follow C# naming conventions
- Use LINQ where appropriate
- Implement IDisposable when needed
- Use proper namespace structure

### For Ruby:
- Follow Ruby style guide
- Use modules for namespacing
- Follow Rails conventions if applicable
- Write idiomatic Ruby code

## Success Criteria
- [ ] Feature works as specified
- [ ] Code follows project patterns
- [ ] Tests pass and provide good coverage
- [ ] Documentation is complete
- [ ] No existing functionality is broken
