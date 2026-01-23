# Contributing to You

Thank you for your interest in contributing to You! This document provides guidelines for contributing to the project.

## Code of Conduct

- Be respectful and inclusive
- Welcome newcomers and help them get started
- Focus on constructive feedback
- Respect differing opinions and experiences

## How to Contribute

### Reporting Bugs

1. **Check existing issues** to see if the bug has already been reported
2. **Create a new issue** with:
   - Clear, descriptive title
   - Steps to reproduce
   - Expected vs actual behavior
   - Your environment (OS, Go version, OpenCode version)
   - Relevant logs or screenshots

### Suggesting Features

1. **Check existing feature requests** to avoid duplicates
2. **Create a new issue** with:
   - Clear description of the feature
   - Use cases and examples
   - Why this feature would be valuable
   - Possible implementation approaches

### Contributing Code

#### 1. Fork and Clone

```bash
# Fork the repository on GitHub
# Then clone your fork
git clone https://github.com/your-username/you.git
cd you/push-to-github
```

#### 2. Create a Branch

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/your-bug-fix
```

#### 3. Set Up Development Environment

```bash
# Install dependencies
go mod download

# Build the project
go build -o you.exe .

# Run tests (when available)
go test ./...
```

#### 4. Make Your Changes

Follow these guidelines:

**Code Style:**
- Follow Go idioms and conventions
- Use `gofmt` to format your code
- Run `golint` and address warnings
- Add comments for exported functions and types

**Commit Messages:**
```
<type>: <subject>

<body>

<footer>
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Example:
```
feat: add custom agent creation command

Implement 'you.exe --create-agent <name>' command that allows
users to create custom agent definitions interactively.

Closes #42
```

**Testing:**
- Add tests for new features
- Ensure existing tests pass
- Aim for >80% code coverage for new code

**Documentation:**
- Update README.md if needed
- Add/update comments in code
- Update QUICKSTART.md for user-facing changes
- Update agent prompts if agent behavior changes

#### 5. Test Your Changes

```bash
# Build
go build -o you.exe .

# Test manually
mkdir test-project
cd test-project
../you.exe --presets
# Edit USER_INPUT.md
../you.exe --orchestrate

# Clean up
cd ..
rm -rf test-project
```

#### 6. Push and Create Pull Request

```bash
git push origin feature/your-feature-name
```

Then create a Pull Request on GitHub with:
- Clear title and description
- Reference related issues
- Screenshots/examples if applicable
- Checklist of changes

## Development Guidelines

### Project Structure

```
you/
├── cmd/                    # CLI commands (future)
├── internal/
│   ├── agents/            # Agent templates and prompts
│   ├── models/            # Data models
│   ├── orchestrator/      # Core logic
│   └── state/             # State management (SCR)
├── main.go                # Entry point
└── go.mod
```

### Adding a New Agent

1. Add agent role constant in `internal/models/models.go`:
```go
const (
    // ... existing roles
    AgentRoleNewRole AgentRole = "NEW_ROLE"
)
```

2. Create agent template function in `internal/agents/templates.go`:
```go
func getNewRoleAgentTemplate() AgentTemplate {
    return AgentTemplate{
        Name: "new-role",
        Role: models.AgentRoleNewRole,
        Description: "Does something specific",
        // ... configuration
        Prompt: generateNewRolePrompt(),
    }
}
```

3. Create prompt function in `internal/agents/prompts.go`:
```go
func generateNewRolePrompt() string {
    return `# SYSTEM IDENTITY
<role>
You are the NewRole Agent...
</role>
// ... rest of prompt
`
}
```

4. Add to agent list in `GetAllAgentTemplates()`:
```go
func GetAllAgentTemplates() []AgentTemplate {
    return []AgentTemplate{
        // ... existing agents
        getNewRoleAgentTemplate(),
    }
}
```

5. Update `orchestrator.go` to include the agent in `opencode.json` generation

6. Update documentation (README, QUICKSTART, etc.)

### Adding a New CLI Command

1. Add command handling in `main.go`:
```go
case "--new-command":
    handleNewCommand(projectPath)
```

2. Implement handler function:
```go
func handleNewCommand(projectPath string) {
    // Implementation
}
```

3. Update `printUsage()` with new command documentation

### Modifying State Management

When adding new state fields:

1. Update models in `internal/models/models.go`
2. Add save/load methods in `internal/state/scr.go`
3. Update orchestrator logic if needed
4. Document state structure changes

## Testing

### Manual Testing Checklist

Before submitting a PR, test:

- [ ] `you.exe --help` displays correctly
- [ ] `you.exe --version` shows version
- [ ] `you.exe --presets` creates all required files
- [ ] `you.exe --orchestrate` generates orchestration guide
- [ ] Agent markdown files have valid frontmatter
- [ ] OpenCode config is valid JSON
- [ ] State files are created in `.you/` directory
- [ ] Build succeeds without errors or warnings

### Future: Automated Tests

We plan to add:
- Unit tests for all packages
- Integration tests for CLI commands
- End-to-end tests with mock OpenCode
- CI/CD pipeline with GitHub Actions

## Documentation Standards

### Code Comments

```go
// PackageName provides functionality for X
package packagename

// FunctionName does something specific
// It takes parameter x which represents Y
// Returns Z or an error if something fails
func FunctionName(x Type) (Type, error) {
    // Implementation
}
```

### Markdown

- Use clear headings hierarchy (H1 → H2 → H3)
- Include code examples with syntax highlighting
- Add links to relevant documentation
- Use tables for structured data
- Include TOC for long documents

## Release Process

(For maintainers)

1. Update version in `main.go`
2. Update CHANGELOG.md
3. Create git tag: `git tag v0.x.0`
4. Push tag: `git push origin v0.x.0`
5. Create GitHub release with binary attachments
6. Update documentation if needed

## Questions?

- Open an issue for general questions
- Tag maintainers for urgent matters
- Join Discord (coming soon) for real-time chat

## Recognition

Contributors will be:
- Listed in README.md
- Mentioned in release notes
- Given credit in commit history

Thank you for contributing to You! 🎉
