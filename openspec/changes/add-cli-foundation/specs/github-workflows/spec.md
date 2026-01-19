# GitHub Workflows

## ADDED Requirements

### Requirement: Continuous Integration Workflow
The project SHALL have a CI workflow that validates code quality on every push and pull request.

#### Scenario: Run CI on push to main
- **WHEN** code is pushed to the main branch
- **THEN** the CI workflow runs build, test, and lint jobs

#### Scenario: Run CI on pull request
- **WHEN** a pull request is opened or updated targeting main
- **THEN** the CI workflow runs and reports status on the PR

#### Scenario: Fail CI on test failures
- **WHEN** any test fails during CI
- **THEN** the workflow fails and blocks merge

#### Scenario: Fail CI on lint errors
- **WHEN** golangci-lint reports errors
- **THEN** the workflow fails and blocks merge

### Requirement: Release Workflow
The project SHALL have a release workflow that builds and publishes binaries when version tags are pushed.

#### Scenario: Trigger release on version tag
- **WHEN** a tag matching v* pattern is pushed (e.g., v1.0.0)
- **THEN** the release workflow is triggered

#### Scenario: Build multi-platform binaries
- **WHEN** the release workflow runs
- **THEN** GoReleaser builds binaries for macOS, Linux, and Windows (amd64 and arm64)

#### Scenario: Publish GitHub release
- **WHEN** GoReleaser completes successfully
- **THEN** a GitHub release is created with compiled binaries attached

### Requirement: GoReleaser Configuration
The project SHALL include a GoReleaser configuration for consistent release builds.

#### Scenario: Configure binary naming
- **WHEN** GoReleaser builds binaries
- **THEN** binaries are named with pattern ralphit_{os}_{arch}

#### Scenario: Generate checksums
- **WHEN** GoReleaser creates a release
- **THEN** SHA256 checksums are generated for all artifacts
