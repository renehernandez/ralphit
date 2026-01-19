# Tasks: Add CLI Foundation

## 1. Project Setup
- [x] 1.1 Initialize Go module (`go mod init github.com/ralphit/ralphit`)
- [x] 1.2 Create directory structure (cmd/, internal/config/, pkg/)

## 2. CLI Framework
- [x] 2.1 Add Cobra dependency
- [x] 2.2 Create root command (cmd/root.go)
- [x] 2.3 Create main.go entry point
- [x] 2.4 Add version command (cmd/version.go)

## 3. Configuration
- [x] 3.1 Add Viper dependency
- [x] 3.2 Create config initialization (internal/config/config.go)
- [x] 3.3 Support config file locations (~/.ralphit.yaml, .ralphit.yaml)

## 4. GitHub Workflows
- [x] 4.1 Create CI workflow (.github/workflows/ci.yml)
  - Build on push/PR to main
  - Run tests with coverage
  - Run golangci-lint
- [x] 4.2 Create release workflow (.github/workflows/release.yml)
  - Trigger on version tags (v*)
  - Use GoReleaser for multi-platform builds
- [x] 4.3 Add .goreleaser.yml configuration

## 5. Development Tooling
- [x] 5.1 Create .mise.toml configuration
  - Pin Go version
  - Add golangci-lint tool
  - Add goreleaser tool
  - Add lefthook tool
- [x] 5.2 Create lefthook.yml configuration
  - pre-commit: gofmt, go vet
  - pre-push: go test

## 6. Validation
- [x] 6.1 Verify `go build` succeeds
- [x] 6.2 Verify `ralphit --help` displays usage
- [x] 6.3 Verify `ralphit version` displays version info
- [ ] 6.4 Verify `mise install` sets up tools
- [ ] 6.5 Verify `lefthook install` configures hooks
- [x] 6.6 Run `go vet` with no errors
