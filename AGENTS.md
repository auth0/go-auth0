# AGENTS.md

## Project Overview

This is the **go-auth0** SDK, a Go client library for Auth0's Management API and Authentication API. The SDK is primarily **auto-generated** from Auth0's OpenAPI specifications, with some hand-written code in the `authentication` package.

**Key Points:**
- Most management API code is auto-generated and should **NOT** be manually modified
- Manual changes are only allowed in the `authentication` package
- The SDK follows the same support policy as Go (last two major releases)
- Requires Go 1.24+

## Setup Commands

```sh
# Download dependencies
go mod download

# Or download to vendor folder
make deps

# Install development tools (linter, vulnerability checker)
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
go install golang.org/x/vuln/cmd/govulncheck@latest
```

## Development Environment

- **Package Manager**: Standard Go modules (`go.mod`)
- **Vendor Directory**: Dependencies can be vendored with `make deps`
- **Go Version**: 1.24 or newer required
- **GOPATH**: Must be properly configured and exported

## Testing Instructions

### Run Tests with HTTP Recordings (Recommended for Development)

```sh
# Run all tests with recordings
make test

# Run specific test with recordings
make test FILTER="TestResourceServer_Read"
```

### Run Tests Against Real Auth0 Tenant (E2E)

```sh
# Run all E2E tests
make test-e2e

# Run specific E2E test
make test-e2e FILTER="TestResourceServer_Read"
```

### Record New HTTP Interactions

```sh
# Record new HTTP test interactions
make test-record

# Record specific test
make test-record FILTER="TestResourceServer_Read"
```

**Important**: When recording new HTTP interactions, delete the old recording file first if regenerating.

### Environment Setup for E2E Tests

Create a `.env` file in the root with:

```
AUTH0_DOMAIN=<your-tenant-domain>
AUTH0_CLIENT_ID=<management-m2m-client-id>
AUTH0_CLIENT_SECRET=<management-m2m-client-secret>
AUTH0_AUTH_CLIENT_ID=<authentication-m2m-client-id>
AUTH0_AUTH_CLIENT_SECRET=<authentication-m2m-client-secret>
AUTH0_DEBUG=true  # Optional: enables debug mode
```

## Code Style and Conventions

### General Guidelines

- Follow standard Go conventions and idioms
- Use functional patterns where appropriate
- Leverage Go's built-in formatting with `gofmt`

### Safe Getters Pattern

The SDK provides safe getter methods to avoid nil pointer panics:

```go
// GetLogoURI() returns "" if client.LogoURI is nil
logoURI := client.GetLogoURI()
```

These getters are auto-generated as part of the SDK generation process.

### Code Modification Rules

- **Management API (`management/` package)**: Do NOT manually modify - code is auto-generated
- **Authentication API (`authentication/` package)**: Manual modifications allowed
- **Internal packages**: Modifications may be allowed depending on use case
- **Tests**: Can be modified and added as needed

## Linting and Quality Checks

### Run Linter

```sh
# Run golangci-lint with auto-fix
make lint
```

This runs `golangci-lint` across the codebase to identify stylistic issues and suspicious constructs.

### Check for Vulnerabilities

```sh
# Run vulnerability scanner
make check-vuln
```

This uses `govulncheck` to scan for known security issues in dependencies.

## Project Structure

```
├── authentication/          # Hand-written authentication API code (modifiable)
├── management/             # Auto-generated management API code (do NOT modify)
├── internal/               # Internal utilities and helpers
├── test/                   # Test files and test data
│   └── data/recordings/    # HTTP test recordings
├── vendor/                 # Vendored dependencies
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
└── Makefile               # Build and test automation
```

## Before Submitting Changes

1. **Run linter**: `make lint`
2. **Check for vulnerabilities**: `make check-vuln`
3. **Run tests**: `make test`
4. **Verify code coverage**: Check that `coverage.out` is generated
5. **Verify you haven't modified auto-generated code** (unless in `authentication/` package)

## Testing Coverage

- Tests use HTTP recordings stored in `test/data/recordings/`
- Coverage reports are generated in `coverage.out`
- Coverage mode is set to `atomic`
- E2E tests have a 20-minute timeout

## Common Make Targets

```sh
make help          # Show all available commands
make deps          # Download dependencies to vendor folder
make lint          # Run linting with auto-fix
make check-vuln    # Check for vulnerabilities
make test          # Run tests with HTTP recordings
make test-record   # Record new HTTP interactions
make test-e2e      # Run tests against real Auth0 tenant
```

## Beta Track Releases

> Applies only when working on the `beta` branch.

This SDK ships two tracks from one module path (`github.com/auth0/go-auth0/v2`):

- **Stable** (`main`): EA/GA endpoints only. Released by a maintainer via a `release/*` branch.
- **Beta** (`beta`): a superset of stable plus beta-only endpoints. Released automatically when a PR is merged into `beta`.

The `beta` branch is **regenerated** from the stable spec plus the beta-only spec files; it is never produced by merging `main` into `beta`. It receives one combined regeneration PR (stable + beta) as a single squash commit. The beta-only versus stable-mirrored split cannot be detected from code or file paths, so it must be recorded in the squash commit message.

### Versioning

Beta = the next stable minor + `-beta.N`, derived from git tags. Latest stable `v2.13.0` -> beta `v2.14.0-beta.N`; `-beta.N` starts at 0 per base and auto-increments; once stable `v2.14.0` ships, beta rolls to `v2.15.0-beta.0`. No state file.

### When merging a beta regeneration PR

Squash and merge with a commit message that marks each group:

```
Regenerate SDK (stable + beta) (#<pr-number>)

<!-- BETA -->
- feat: add `Management.Sandbox` preview API (Beta)
<!-- /BETA -->

<!-- STABLE -->
- feat: add tenant security headers configuration
<!-- /STABLE -->
```

- Beta-only changes go inside the `BETA` markers; stable-mirrored changes go inside the `STABLE` markers.
- The `<!-- ... -->` markers are HTML comments and stay invisible in GitHub's rendered view.
- Omitting a section renders a "No ... changes in this release." note; omitting both falls back to the raw commit subject. Always prefer the structured form.
- Everything reaches `beta` through a PR. Never push directly to `beta` (a direct push will not trigger a release).

### Do not hand-edit release files on `beta`

The Beta Auto-Release workflow (`.github/workflows/beta-autorelease.yml`) owns versioning. When a PR is merged into `beta` it computes the next `vX.Y.0-beta.N`, aborts if that tag already exists, stamps `.version` and `meta.go`, prepends a `CHANGELOG.md` entry with separate **Beta** and **Stable (from main)** sections, creates the release commit **through the GitHub API** (so it is signed/Verified, no GPG key), tags it, and publishes a prerelease. Never manually bump `.version`, `meta.go`, or `CHANGELOG.md` on `beta`.

### Hand-written code

The `authentication/` package is hand-written and is **not** regenerated on `beta`. A fix on `main` does not reach `beta` automatically, so hand-written changes must be PR'd to **both** `main` and `beta`.

## Important Notes for Agents

- **Never modify auto-generated code** in the `management/` package unless explicitly asked
- Always run `make lint` before committing changes
- Use HTTP recordings for tests when possible to avoid rate limiting
- When adding new tests, consider adding HTTP recordings with `make test-record`
- Check that tests pass with `make test` before finalizing changes
- The SDK uses pointer fields extensively - use safe getters to avoid nil panics
- Authentication package is hand-written and can be modified carefully
- Follow existing code patterns, especially in the authentication package
- On the `beta` branch, never hand-edit `.version`, `meta.go`, or `CHANGELOG.md`; mark beta vs. stable changes in the squash commit message (see "Beta Track Releases")
