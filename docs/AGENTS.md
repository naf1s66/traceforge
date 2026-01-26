# TraceForge — AGENTS (Codex)

This repo is designed for small, parallel PRs with clear acceptance criteria.

## Branching
- Base branch for work: `main`
- Feature branches: `milestone-<n>/<slug>`
- PRs must target: `main`
- Keep PRs small (aim ≤ 300 lines net where possible).

## Quality gates
- `make ci` must pass locally when feasible.
- No secrets in commits (use env examples).
- Add/Update OpenAPI when endpoints change.
- Each PR must include:
  - Summary
  - How to test
  - Screenshots (if web)

## Milestone task docs
- See `docs/tasks/` for day-by-day scope.

## System constraints
- Free-tier safe by default:
  - Strict rate limiting on public read endpoints
  - Hosted demo read-only (write disabled unless explicitly enabled)
- Multi-tenant (`workspace_id`) exists in the data model to signal SaaS readiness.
- Idempotency is required for write ingestion.

## Coding standards
- Go: keep handlers small; validate input; return typed errors.
- Web: Remix loaders/actions; shadcn/ui; dark default; minimal and sleek.
