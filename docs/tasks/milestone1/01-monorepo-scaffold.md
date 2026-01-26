# Task: Scaffold the monorepo layout

## Summary
- Establish the top-level repo layout with `apps/api`, `apps/web`, `infra`, and `docs`.
- Add base tooling (pnpm workspace + Makefile) to run API and web together.

**Status:** Completed.

## Acceptance Criteria
- [x] Repo contains `apps/api`, `apps/web`, `infra`, and `docs` at the top level.
- [x] Root `pnpm-workspace.yaml` and `package.json` wire the web workspace.
- [x] `Makefile` offers install/dev/build/test convenience targets.

## Notes
- Align with ADR 0001 and keep paths stable for future milestones.
