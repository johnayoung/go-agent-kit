# Code Refactor Workflow

## STAGE 1: CODEBASE ANALYSIS
You are analyzing this codebase to refactor: {{.Description}}

First, examine the current state and identify improvement opportunities:

1. **Understand the current implementation**
   - What does the existing code do?
   - What are the current pain points or limitations?
   - What specific aspects need improvement?
   - Are there performance, maintainability, or readability issues?

2. **Identify code smells and areas for improvement**
   - Look for duplicated code patterns
   - Find overly complex functions or classes
   - Identify poor naming conventions
   - Check for tight coupling and low cohesion
   - Look for violation of SOLID principles

3. **Analyze the current architecture**
   - Understanding existing patterns and conventions
   - Identify dependencies and relationships
   - Map out the current data flow
   - Find integration points that will be affected

4. **Assess impact and risk**
   - What other parts of the codebase depend on this code?
   - Are there existing tests that need to be updated?
   - What are the potential breaking changes?

@workspace examine the code sections that need refactoring

Output your analysis. DO NOT write code yet.

## STAGE 2: REFACTOR PLAN
Based on your analysis, create a detailed refactoring plan:

1. **Refactoring goals**
   - What specific improvements are you trying to achieve?
   - How will this make the code better (performance, readability, maintainability)?
   - What design patterns or principles will guide the refactor?

2. **Refactoring strategy**
   - What refactoring techniques will you use? (Extract Method, Extract Class, Move Method, etc.)
   - In what order will you apply the changes?
   - How will you maintain backward compatibility if needed?

3. **Implementation approach**
   - What files will need to be modified or created?
   - What are the key interfaces or abstractions to introduce?
   - How will you handle data migration or state transitions?

4. **Risk mitigation**
   - How will you ensure existing functionality remains intact?
   - What safeguards will you put in place?
   - How will you handle rollback if needed?

## STAGE 3: IMPLEMENTATION
Implement the refactoring following your plan:

1. **Start with the safest changes**
   - Begin with low-risk refactoring (renaming, extracting constants)
   - Apply one refactoring technique at a time
   - Test after each significant change

2. **Follow language-specific best practices**
   - Use appropriate design patterns for your language
   - Follow naming conventions and coding standards
   - Leverage language-specific features and idioms

3. **Maintain existing behavior**
   - Ensure all existing functionality still works
   - Preserve public interfaces unless explicitly changing them
   - Handle edge cases that the original code handled

4. **Improve code quality**
   - Make code more readable and self-documenting
   - Reduce complexity and improve maintainability
   - Eliminate code duplication and improve reusability

## STAGE 4: TESTING
Thoroughly test the refactored code:

1. **Verify existing functionality**
   - Run all existing tests to ensure they still pass
   - Test all code paths and edge cases
   - Verify performance hasn't degraded

2. **Update tests as needed**
   - Modify tests that were affected by the refactoring
   - Add new tests for any new abstractions or interfaces
   - Ensure test coverage remains high or improves

3. **Integration testing**
   - Test how the refactored code interacts with other components
   - Verify that dependent systems still work correctly
   - Test in realistic scenarios and environments

4. **Performance validation**
   - Measure and compare performance before and after
   - Ensure memory usage and execution time are acceptable
   - Load test if the refactor affects performance-critical paths

## STAGE 5: DOCUMENTATION
Update documentation to reflect the changes:

1. **Code documentation**
   - Update comments to reflect new structure and logic
   - Document new interfaces, classes, or methods
   - Explain design decisions and trade-offs made

2. **API documentation**
   - Update public API documentation if interfaces changed
   - Add examples showing how to use new or modified APIs
   - Document any breaking changes and migration paths

3. **Architecture documentation**
   - Update design documents to reflect new structure
   - Document new patterns or principles introduced
   - Explain how the refactor improves the overall architecture

## Language-Specific Refactoring Guidelines

### For Go:
- Use interfaces for abstraction and testability
- Extract functions to improve readability and reusability
- Leverage Go's composition over inheritance
- Use channels and goroutines appropriately for concurrency
- Follow effective Go practices and idioms

### For Python:
- Use classes and modules for better organization
- Leverage Python's duck typing and dynamic features
- Apply PEP 8 style guidelines consistently
- Use decorators and context managers where appropriate
- Consider using dataclasses for data structures

### For TypeScript/JavaScript:
- Use proper TypeScript types to improve type safety
- Extract modules and classes for better organization
- Leverage modern ES6+ features (arrow functions, destructuring, etc.)
- Use async/await for better asynchronous code handling
- Apply functional programming concepts where appropriate

### For Java:
- Use design patterns (Strategy, Factory, Observer, etc.)
- Leverage Java 8+ features (streams, lambdas, optional)
- Apply SOLID principles consistently
- Use proper exception handling patterns
- Consider using builder pattern for complex object creation

### For C#:
- Use LINQ for data manipulation and querying
- Leverage C# features (properties, events, extension methods)
- Apply dependency injection for better testability
- Use async/await for asynchronous operations
- Follow .NET conventions and best practices

### For Ruby:
- Use modules for mixins and namespacing
- Leverage Ruby's metaprogramming features judiciously
- Follow Ruby style guide and idiomatic patterns
- Use blocks and yield for flexible APIs
- Apply Rails conventions if working in Rails

## Common Refactoring Techniques

### Code Organization:
- **Extract Method**: Break large methods into smaller, focused ones
- **Extract Class**: Split classes with multiple responsibilities
- **Move Method**: Relocate methods to more appropriate classes
- **Inline Method**: Remove unnecessary method indirection

### Data Structure Improvements:
- **Replace Magic Numbers with Constants**: Improve readability
- **Encapsulate Field**: Provide proper access to data
- **Replace Array with Object**: Use proper data structures
- **Replace Type Code with Class**: Use polymorphism instead of conditionals

### Simplification:
- **Simplify Conditional Expressions**: Make logic clearer
- **Remove Dead Code**: Eliminate unused code
- **Consolidate Duplicate Code**: Extract common functionality
- **Replace Nested Conditional with Guard Clauses**: Reduce complexity

### Architecture Improvements:
- **Introduce Design Patterns**: Apply appropriate patterns
- **Separate Concerns**: Ensure single responsibility
- **Dependency Inversion**: Depend on abstractions, not concretions
- **Interface Segregation**: Create focused, cohesive interfaces

## Success Criteria
- [ ] Code is more readable and maintainable
- [ ] Performance is maintained or improved
- [ ] All existing tests pass
- [ ] Test coverage is maintained or improved
- [ ] Code follows language and project conventions
- [ ] Documentation accurately reflects changes
- [ ] No regressions in functionality
- [ ] Refactoring goals are achieved
