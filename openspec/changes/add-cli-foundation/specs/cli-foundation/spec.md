# CLI Foundation

## ADDED Requirements

### Requirement: Root Command
The CLI SHALL provide a root command `ralphit` that serves as the entry point for all subcommands.

#### Scenario: Display help when invoked without arguments
- **WHEN** user runs `ralphit` without arguments
- **THEN** display usage information and available commands

#### Scenario: Display help with --help flag
- **WHEN** user runs `ralphit --help`
- **THEN** display detailed help information including all flags and subcommands

### Requirement: Version Command
The CLI SHALL provide a `version` subcommand that displays build information.

#### Scenario: Display version information
- **WHEN** user runs `ralphit version`
- **THEN** display the application version, Go version, and build date

### Requirement: Configuration Management
The CLI SHALL support configuration via Viper with multiple sources.

#### Scenario: Load configuration from file
- **WHEN** a config file exists at ~/.ralphit.yaml or .ralphit.yaml
- **THEN** load configuration values from the file

#### Scenario: Override configuration with environment variables
- **WHEN** environment variables prefixed with RALPHIT_ are set
- **THEN** use environment variable values over file configuration

#### Scenario: Override configuration with flags
- **WHEN** command-line flags are provided
- **THEN** use flag values over environment and file configuration

### Requirement: Cross-Platform Compatibility
The CLI SHALL build and run on macOS, Linux, and Windows without modification.

#### Scenario: Build for multiple platforms
- **WHEN** building with `go build`
- **THEN** produce a working binary for the target OS/architecture
