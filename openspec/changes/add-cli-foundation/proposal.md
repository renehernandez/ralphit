# Change: Add CLI Foundation

## Why
ralphit needs a foundational CLI structure before any features can be implemented. This establishes the project skeleton using Cobra and Viper, sets up CI/CD pipelines for automated testing and releases, and configures development tooling for a consistent developer experience.

## What Changes
- Initialize Go module with proper module path
- Add Cobra CLI framework with root command
- Integrate Viper for configuration management
- Create standard project directory structure (cmd/, internal/, pkg/)
- Implement version command for release tracking
- Add GitHub Actions workflow for CI (build, test, lint)
- Add GitHub Actions workflow for releases (GoReleaser)
- Configure mise for development environment automation
- Configure lefthook for git hooks (pre-commit, pre-push)

## Impact
- Affected specs: cli-foundation (new), github-workflows (new), dev-tooling (new)
- Affected code: main.go, cmd/, internal/config/, .github/workflows/, .mise.toml, lefthook.yml
