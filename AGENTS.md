# ralphit - AI Agent Instructions

## Project Overview

**ralphit** is a CLI tool that implements the "Ralph Wiggum" technique for iteratively building software across multiple AI coding assistant sessions (Claude Code, Codex, OpenCode, etc.) with fresh context. It helps developers maintain continuity and structure when working with AI assistants across different sessions.

## Tech Stack

- **Language:** Go 1.21+
- **CLI Framework:** Cobra + Viper
- **Testing:** Testify
- **License:** MIT

## Project Structure

```
ralphit/
├── cmd/           # Cobra commands
├── internal/      # Private application code
├── pkg/           # Public reusable packages
├── openspec/      # OpenSpec configuration and proposals
└── main.go        # Entry point
```

## Development Commands

```bash
# Build
go build -o ralphit .

# Run tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...

# Format code
gofmt -w .

# Lint
go vet ./...
golint ./...

# Install locally
go install .
```

## Code Conventions

### Style
- Follow standard Go conventions (gofmt, golint, Effective Go)
- Use `gofmt` for formatting before committing
- Run `go vet` for static analysis

### Naming
- Use camelCase for unexported identifiers
- Use PascalCase for exported identifiers
- Prefer descriptive names over abbreviations

### Error Handling
- Always handle errors explicitly
- Wrap errors with context using `fmt.Errorf("context: %w", err)`
- Return errors rather than panicking

### Testing
- Use Testify for assertions (`assert`, `require`)
- Write table-driven tests where appropriate
- Name test files `*_test.go`
- Name test functions `TestFunctionName_Scenario`

## Git Workflow

- **Trunk-based development** - short-lived branches merged frequently to main
- **Conventional commits** - use prefixes: `feat:`, `fix:`, `docs:`, `test:`, `refactor:`, `chore:`
- **PRs required** - all changes go through pull requests

## Important Notes for AI Agents

1. **Read before editing** - Always read existing code before making changes
2. **Minimal changes** - Make only the changes necessary to complete the task
3. **Run tests** - Verify changes don't break existing functionality
4. **Cross-platform** - Code must work on macOS, Linux, and Windows
5. **Dependencies** - Keep external dependencies minimal; prefer standard library when practical

## OpenSpec Integration

<!-- OPENSPEC:START -->
# OpenSpec Instructions

These instructions are for AI assistants working in this project.

Always open `@/openspec/AGENTS.md` when the request:
- Mentions planning or proposals (words like proposal, spec, change, plan)
- Introduces new capabilities, breaking changes, architecture shifts, or big performance/security work
- Sounds ambiguous and you need the authoritative spec before coding

Use `@/openspec/AGENTS.md` to learn:
- How to create and apply change proposals
- Spec format and conventions
- Project structure and guidelines

Keep this managed block so 'openspec update' can refresh the instructions.

<!-- OPENSPEC:END -->
