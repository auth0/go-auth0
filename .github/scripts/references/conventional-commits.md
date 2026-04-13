# Conventional Commits — Title Format

## Rules
- Max 70 characters
- Format: `type(scope): description` or `type: description`
- Use lowercase
- No period at the end
- Append `!` before `:` for breaking changes

## Types
- `feat` — new feature
- `fix` — bug fix
- `refactor` — code restructuring (no behavior change)
- `docs` — documentation only
- `test` — tests only
- `chore` — maintenance / dependencies
- `ci` — CI/CD changes
- `perf` — performance improvement
- `style` — formatting (no logic change)
- `build` — build system changes

## Examples
- `feat: add user authentication`
- `feat(auth): add OAuth2 support`
- `fix: resolve null pointer in user resolver`
- `fix(api): handle timeout in health check`
- `refactor: extract validation logic`
- `docs: update API migration guide`
- `test: add integration tests for payments`
- `chore: upgrade dependencies`
- `ci: add GitHub Actions workflow`
- `perf: optimize database queries`
- `feat(api)!: remove deprecated endpoints`
