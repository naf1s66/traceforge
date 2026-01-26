# ADR 0001: Monorepo layout

## Decision
Use a monorepo:
- `apps/api` (Go)
- `apps/web` (Remix)
- `infra` (docker compose)
- `docs` (PRD, tasks, ADRs)

## Rationale
One repo makes it easy to:
- demo end-to-end locally
- keep docs and code aligned
- show CI/CD skills clearly

## Consequences
Web uses pnpm; API uses Go modules; Makefile orchestrates both.
