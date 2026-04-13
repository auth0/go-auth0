You have access to a git worktree checked out at: {{WORKTREE_PATH}}
This is the PR's head branch ({{HEAD_BRANCH}}) code. The base branch is {{BASE_BRANCH}}.

**Run all commands from inside the worktree**: `cd {{WORKTREE_PATH}}` before doing anything.

Analyze the actual code structure for breaking changes, SDK/API surface changes, and overall impact.

### 1. Public API / SDK surface analysis
- Look for exported modules, public APIs, SDK entry points (e.g., index.ts, index.js, mod.rs, __init__.py, lib.rs)
- Check if any public exports were removed or renamed
- Check if function/method signatures changed (parameters added/removed/reordered, return types changed)
- Check if types/interfaces/structs were modified in ways that break consumers
- Look at package.json, Cargo.toml, setup.py, go.mod etc. for version or dependency changes

### 2. File structure changes
- Run: git diff --name-status origin/{{BASE_BRANCH}}...HEAD  (from inside the worktree)
- Identify renamed, moved, or deleted files
- Check if deleted/moved files are still imported by other files in the codebase
- Note any new directories or significant structural reorganization

### 3. Configuration & schema changes
- Check for changes to config files, environment variables, or feature flags
- Check for database migration files (schema changes)
- Check for changes to CI/CD configs, Dockerfiles, or build scripts

### 4. Test coverage assessment
- List test files that were added or modified
- Check if new source files have corresponding test files
- Look for test configuration changes

### 5. Dependency changes
- Check if dependencies were added, removed, or had major version bumps
- Flag any security-sensitive dependency changes

Return a structured summary with clear sections. Flag anything that is a **breaking change** prominently at the top.
