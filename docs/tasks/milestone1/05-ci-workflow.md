# Task: Add CI for API and web builds

## Summary
- Configure GitHub Actions to run Go tests/builds and web builds on main PRs.
- Ensure CI provisions Postgres for API tests.

**Status:** Completed.

## Acceptance Criteria
- [x] `.github/workflows/ci.yml` runs on push/pull_request for `main`.
- [x] Go tests and builds run successfully in CI.
- [x] Web lint/build steps run with env placeholders.

## Notes
- Keep CI fast and deterministic; no secrets in workflow envs.
