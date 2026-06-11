# Contributing

We appreciate feedback and contribution to this repo.
Before you submit a pull request, there are a couple requirements to satisfy.

## Prerequisites

Ensure you have the following prerequisites met:

- **Go Version**: You must have Go version 1.24 or newer installed on your system. You can check your Go version by running `go version` in your terminal.
- **GOPATH Set Up**: Make sure you have set up your GOPATH.
- **Export GOPATH**: Ensure you have exported your GOPATH to your system's environment variables. This allows tools and libraries to locate your Go workspace.
- **Modify bashrc or zshrc**: To automatically set your GOPATH and export it every time you open a terminal, you can add the following lines to your `~/.bashrc` or `~/.zshrc` file on Linux and macOS systems:
- **Environment File**: Ensure you have a `.env` file in the root folder of the repository or that the environment variables have been exported.

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

After adding these lines, run source `~/.bashrc` or source `~/.zshrc` to apply the changes.

## About This SDK

This SDK is auto-generated from Auth0's OpenAPI specifications. Most of the management API code is automatically generated and should **not** be manually modified.

**Manual modifications are only allowed in the `authentication` package**, which contains hand-written code for authentication flows.

## Beta Track Releases

> This section applies only on the `beta` branch.

The SDK ships two parallel tracks from one Go module path
(`github.com/auth0/go-auth0/v2`):

- **Stable** (`main` branch): EA/GA endpoints only. Released the usual way, by a maintainer, through a `release/*` branch.
- **Beta** (`beta` branch): a superset that contains everything in stable **plus** beta-only endpoints. Released automatically whenever a PR is merged into `beta`.

Consumers opt into beta explicitly; a plain `go get` never selects it:

```sh
# Stable (default): always resolves to the latest stable release.
go get github.com/auth0/go-auth0/v2

# Beta: explicit prerelease pin.
go get github.com/auth0/go-auth0/v2@v2.14.0-beta.1
```

### How beta stays a superset of stable

The `beta` branch is **regenerated** from the stable spec plus the beta-only
spec files. It is never produced by `git merge main -> beta`. When endpoints
change upstream, fern-bot opens a regeneration PR into `beta` (combined stable +
beta) and a separate PR into `main` (stable only). Do **not** merge `main` into
`beta`.

**Hand-written code** (the `authentication/` package) is not regenerated, so a
fix made on `main` does not reach `beta` automatically. Hand-written changes must
be opened as separate PRs to **both** `main` and `beta`.

### How beta versions are numbered

Beta always carries the **next** stable minor with a `-beta.N` suffix, derived
entirely from the git tags (no state file):

- Latest stable `v2.13.0` -> beta is `v2.14.0-beta.N`.
- `-beta.N` auto-increments on each beta release.
- The moment stable `v2.14.0` is tagged, the next beta automatically becomes
  `v2.15.0-beta.1`.

Because `v2.14.0-beta.3` sorts before `v2.14.0` in semver, beta is never picked
up by `go get` unless pinned.

### Merging a beta regeneration PR

The combined regeneration PR lands as a **single squash commit**, and the
beta-only versus stable-mirrored split cannot be recovered from the code or file
paths. The person merging records the split in the squash commit message, and
the Beta Auto-Release workflow parses it into the changelog and release notes.

When you squash and merge a beta PR, format the commit message like this:

```
Regenerate SDK (stable + beta) (#<pr-number>)

<!-- BETA -->
- feat: add `Management.Sandbox` preview API (Beta)
- feat: add device-flow polling endpoint (Beta)
<!-- /BETA -->

<!-- STABLE -->
- feat: add tenant security headers configuration
- feat: add minute-based session lifetime fields
<!-- /STABLE -->
```

Guidelines:

- Put **beta-only** changes inside the `BETA` markers and **stable-mirrored** changes inside the `STABLE` markers.
- The `<!-- ... -->` markers are HTML comments, so they stay invisible in GitHub's rendered view while remaining easy to parse.
- If a release has no beta-only changes, leave the `BETA` section empty or omit it. The workflow records "No beta-only changes in this release." (and likewise for a missing `STABLE` section).
- If you omit both sections entirely, the workflow falls back to the raw commit subject so the release is never empty. Always prefer the structured form.
- All changes must reach `beta` through a PR. Do not push commits directly to `beta`; a direct push will not trigger a release.

### What happens automatically after merge

Do **not** hand-edit `.version`, `meta.go`, or `CHANGELOG.md` on the `beta`
branch. When a PR is merged into `beta`, the Beta Auto-Release workflow
(`.github/workflows/beta-autorelease.yml`):

1. Computes the next beta version from the existing tags (`vX.Y.0-beta.N`).
2. Aborts if that tag already exists, so a re-run can never overwrite a published release.
3. Builds the release notes by parsing the `BETA` / `STABLE` sections of the merge commit message.
4. Stamps `.version` and `meta.go`, and prepends a self-contained `CHANGELOG.md` entry with a **Beta** section and a **Stable (from main)** section.
5. Creates the `Release vX.Y.0-beta.N` commit **through the GitHub API**, so the commit is signed by GitHub and shows as **Verified** (the same way stable releases are signed). No GPG key is involved.
6. Tags that commit and publishes a GitHub release marked as a prerelease.

The release fires on PR merge (not on raw pushes), so the workflow's own release
commit cannot re-trigger it, and concurrent merges are serialized and each get
their own version.

### CI on the beta branch

Beta is gated by the same checks as `main`: the build/lint/test workflow
(`main.yml`), Go vulnerability scanning (`govulncheck.yml`), and the SCA scan
(`sca_scan.yml`) all run on `beta` pushes and PRs.

## Safe Getters

The SDK offers safe getters to access pointer fields and avoid panics on nil pointers.

```go
// Example
client, err := mgmt.Client.Read(context.Background(), "EXAMPLE_16L9d34h0qe4NVE6SaHxZEid")
if err != nil {
    return err
}

// GetLogoURI() is safe to use as it will return "" if client.LogoURI is nil.
fmt.Printf("Logo URI: %v", client.GetLogoURI())
```

The getter methods are automatically generated as part of the SDK generation process.

## Code Quality and Security

### Linting

Before submitting your pull request, ensure your code follows the project's coding standards by running the linting task:

```sh
make lint
```

This will run golangci-lint across the codebase, identifying potential stylistic issues and suspicious constructs.

### Vulnerability Checking

To check for any known vulnerabilities in the dependencies, run:

```sh
make check-vuln
```

This uses govulncheck to scan the codebase for known security issues in the Go ecosystem.

### Code Validation

For the auto-generated management API code, validation is handled during the generation process. For manual changes to the authentication package, ensure your code follows Go best practices and doesn't break existing functionality by running the standard Go tools and tests.

## Running the Tests

There are two ways of running the tests:

- `make test` - runs the tests with http recordings. To run a specific test pass the `FILTER` var. Usage `make test FILTER="TestResourceServer_Read"`.
- `make test-e2e` - runs the tests against a real Auth0 tenant. To run a specific test pass the `FILTER` var. Usage `make test-record FILTER="TestResourceServer_Read"`.

### Running against an Auth0 tenant

To run the tests against an Auth0 tenant start by creating an M2M app using `auth0 apps create --name go-auth0-mgmt-tests --description "App used for go-auth0 management tests" --type m2m`, then
run `auth0 apps open <CLIENT ID>`. Authorize the Management API in the `APIs` tab and enable all permissions.

Then create a `.env` file in the root folder of the repository with the following details :

- `AUTH0_DOMAIN`: The **Domain** of the Auth0 tenant
- `AUTH0_CLIENT_ID`: The **Client ID** of the M2M app
- `AUTH0_CLIENT_SECRET`: The **Client Secret** of the M2M app
- `AUTH0_DEBUG`: Set to `true` to call the Management API in debug mode, which dumps the HTTP requests and responses to the output

Now for the Authentication tests create another M2M app using `auth0 apps create --name go-auth0-auth-tests --description "App used for go-auth0 authentication tests" --type m2m`, then run
`auth0 apps open <CLIENT ID>`. Ensure all `Grant Types` except `Client Credentials` are enabled in `Advanced Settings`, then set the `Authentication Method` to `None` in the `Credentials` tab.

Then create a `.env` file in the root folder of the repository with the following details :

- `AUTH0_DOMAIN`: The **Domain** of the Auth0 tenant
- `AUTH0_CLIENT_ID`: The **Client ID** of the management M2M app
- `AUTH0_CLIENT_SECRET`: The **Client Secret** of the management M2M app
- `AUTH0_AUTH_CLIENT_ID`: The **Client ID** of the authentication M2M app
- `AUTH0_AUTH_CLIENT_SECRET`: The **Client Secret** of the authentication M2M app
- `AUTH0_DEBUG`: Set to `true` to call the Management API in debug mode, which dumps the HTTP requests and responses to the output

> **Note**
> The http test recordings can be found in the [recordings](./test/data/recordings) folder.

## Adding New HTTP Test Recordings

To add new http test recordings or to regenerate old ones you can make use of the `make test-record` command.

> **Warning**
> If you need to regenerate an old recording, make sure to delete the corresponding recording file first.

To add only one specific http test recording pass the `FILTER` var, e.g. `make test-record FILTER="TestResourceServer_Read"`.

> **Warning**
> Recording a new http test interaction will make use of a real Auth0 test tenant.
