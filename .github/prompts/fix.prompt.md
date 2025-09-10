# Bug Fix Workflow

## STAGE 1: DIAGNOSIS
You are analyzing this codebase to fix: {{.Description}}

First, diagnose the issue systematically:

1. **Understand the problem**
   - What is the expected behavior?
   - What is the actual behavior?
   - When does this issue occur?
   - What are the error messages (if any)?

2. **Locate the problem area**
   - Identify the likely files/modules involved
   - Look for recent changes that might have introduced the bug
   - Check error logs and stack traces
   - Review related test failures

3. **Reproduce the issue**
   - Create a minimal reproduction case
   - Identify the exact steps to trigger the bug
   - Test in different environments if applicable

4. **Analyze the root cause**
   - Examine the code logic in the problem area
   - Look for edge cases, null checks, boundary conditions
   - Check for race conditions or timing issues
   - Verify data flow and state management

@workspace examine the relevant code sections and error patterns

Output your diagnosis. DO NOT write code yet.

## STAGE 2: FIX STRATEGY
Based on your diagnosis, plan the fix:

1. **Fix approach**
   - What needs to be changed?
   - What is the safest way to implement the fix?
   - Are there multiple possible solutions?

2. **Impact assessment**
   - What other parts of the code might be affected?
   - Are there performance implications?
   - Will this fix break anything else?

3. **Testing strategy**
   - How will you verify the fix works?
   - What regression tests are needed?
   - Are there edge cases to test?

## STAGE 3: IMPLEMENTATION
Implement the fix following your strategy:

1. **Apply the minimal fix**
   - Make the smallest change possible to resolve the issue
   - Follow existing code patterns and conventions
   - Add appropriate error handling

2. **Add safety checks**
   - Include validation and null checks where needed
   - Handle edge cases properly
   - Follow language-specific best practices

## STAGE 4: TESTING
Thoroughly test the fix:

1. **Verify the fix**
   - Test the original reproduction case
   - Ensure the expected behavior is now correct
   - Test edge cases and boundary conditions

2. **Regression testing**
   - Run existing tests to ensure nothing broke
   - Test related functionality
   - Verify performance hasn't degraded

3. **Add new tests**
   - Create tests that would have caught this bug
   - Test the fix directly
   - Add tests for edge cases discovered

## STAGE 5: DOCUMENTATION
Document the fix appropriately:

1. **Code comments**
   - Explain why the fix was necessary
   - Document any non-obvious logic
   - Add warnings about potential pitfalls

2. **Update documentation**
   - Update README if behavior changed
   - Update API docs if applicable
   - Document any new error conditions

## Language-Specific Debugging Guidelines

### For Go:
- Use `fmt.Printf` debugging or proper logging
- Check for nil pointer dereferences
- Verify goroutine safety and channel usage
- Use `go vet` and race detector (`go test -race`)

### For Python:
- Use `print()` statements or `logging` module
- Check for `None` values and type mismatches
- Verify imports and module paths
- Use `pdb` debugger for complex issues

### For TypeScript/JavaScript:
- Use `console.log()` for debugging
- Check for `undefined` and `null` values
- Verify async/await usage and Promise handling
- Use browser developer tools or Node.js debugger

### For Java:
- Use `System.out.println()` or proper logging
- Check for `NullPointerException`
- Verify exception handling
- Use IDE debugger for step-through debugging

### For C#:
- Use `Console.WriteLine()` or logging framework
- Check for null reference exceptions
- Verify async/await patterns
- Use Visual Studio debugger

### For Ruby:
- Use `puts` or `p` for debugging
- Check for `nil` values
- Verify method calls and variable scope
- Use `binding.pry` for interactive debugging

## Common Bug Patterns to Check

### Logic Errors:
- Off-by-one errors in loops
- Incorrect conditional logic
- Wrong operator usage (= vs ==)
- Missing break statements in switch/case

### Data Issues:
- Null/undefined reference errors
- Type conversion problems
- Incorrect data validation
- Missing boundary checks

### Concurrency Issues:
- Race conditions
- Deadlocks
- Shared state problems
- Improper synchronization

### Integration Issues:
- API contract violations
- Database connection problems
- Configuration errors
- Dependency version conflicts

## Success Criteria
- [ ] Original issue is resolved
- [ ] No new bugs introduced
- [ ] All tests pass (including new ones)
- [ ] Code follows project conventions
- [ ] Fix is properly documented
- [ ] Performance impact is acceptable
