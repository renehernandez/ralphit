# Development Tooling

## ADDED Requirements

### Requirement: Mise Configuration
The project SHALL use mise for development environment automation and tool version management.

#### Scenario: Pin Go version
- **WHEN** a developer runs `mise install`
- **THEN** the configured Go version is installed and activated

#### Scenario: Install development tools
- **WHEN** a developer runs `mise install`
- **THEN** golangci-lint, goreleaser, and lefthook are installed at pinned versions

#### Scenario: Automatic tool activation
- **WHEN** a developer enters the project directory with mise configured
- **THEN** the correct tool versions are automatically activated

### Requirement: Lefthook Git Hooks
The project SHALL use lefthook for git hook management to enforce code quality before commits and pushes.

#### Scenario: Pre-commit hook runs formatting check
- **WHEN** a developer attempts to commit
- **THEN** gofmt checks run and block commit if files are not formatted

#### Scenario: Pre-commit hook runs vet
- **WHEN** a developer attempts to commit
- **THEN** go vet runs and blocks commit on errors

#### Scenario: Pre-push hook runs tests
- **WHEN** a developer attempts to push
- **THEN** go test runs and blocks push on test failures

#### Scenario: Install hooks with lefthook
- **WHEN** a developer runs `lefthook install`
- **THEN** git hooks are configured in .git/hooks/

### Requirement: Consistent Developer Experience
The project SHALL provide documentation and tooling for a consistent development setup across machines.

#### Scenario: New developer setup
- **WHEN** a new developer clones the repository
- **THEN** running `mise install && lefthook install` sets up the complete development environment
