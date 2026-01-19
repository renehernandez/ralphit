# Project Context

## Purpose
ralphit is a CLI tool that implements and manages the "Ralph Wiggum" technique for iteratively building software across multiple AI coding assistant sessions (Claude Code, Codex, OpenCode, etc.) with fresh context. It helps developers maintain continuity and structure when working with AI assistants across different sessions.

## Tech Stack
- **Language:** Go (Golang)
- **CLI Framework:** Cobra + Viper (command structure and configuration management)
- **Testing:** Testify (assertions and mocking)
- **License:** MIT

## Project Conventions

### Code Style
- Follow standard Go conventions (gofmt, golint)
- Adhere to Effective Go guidelines
- Use `gofmt` for formatting
- Run `go vet` and `golint` for static analysis

### Architecture Patterns
- Cobra for CLI command structure
- Viper for configuration management
- Clean separation between commands (cmd/), core logic (internal/), and shared utilities (pkg/)
- Prefer composition over inheritance
- Keep functions small and focused

### Testing Strategy
- Use Testify for assertions (`assert`, `require`) and mocking
- Table-driven tests where appropriate
- Aim for meaningful test coverage on core logic
- Test command behavior via integration tests

### Git Workflow
- Trunk-based development
- Short-lived feature branches merged frequently to main
- Conventional commits preferred (feat:, fix:, docs:, etc.)
- PRs for code review before merging

## Domain Context
- "Ralph Wiggum technique" refers to a methodology for iterative software development using AI coding assistants
- The tool helps manage context, state, and continuity across multiple AI sessions
- Users work with various AI coding CLIs (Claude Code, Codex, OpenCode, etc.)

## Important Constraints
- Must work cross-platform (macOS, Linux, Windows)
- Should be installable via `go install`
- Keep dependencies minimal where possible

## External Dependencies
- None planned yet - the tool is self-contained
