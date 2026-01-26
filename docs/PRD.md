# TraceForge — PRD

## One-liner
TraceForge is a lightweight, production-minded **audit logging / event ledger** service with an **interactive public API Explorer UI**.

## Goals
- Recruiter-ready: demonstrates senior backend patterns (**idempotency**, multi-tenant scoping, rate limiting, append-only ledger).
- Shippable in ≤ 1 week, including a hosted demo (free tiers).
- Sleek, minimal, modern UI (dark default) with a landing page + explorer.

## Non-goals (Week 1)
- Full org/user auth, RBAC, billing.
- Advanced analytics, full-text search.
- Complex “workspace provisioning” UX (we use a dedicated public demo workspace).

## Personas
1. **Developer (Demo visitor)**: wants to understand the product quickly and try queries in-browser.
2. **Developer (Integrator)**: wants to ingest events into TraceForge reliably and query them later.

## Core requirements

### API (Go + Postgres)
- Base path: `/api/traceforge/v1`
- Append-only events. No update/delete endpoints.
- Multi-tenant scoping via `workspace_id` on all records and queries.
- Auth:
  - **Write endpoints** require `X-API-Key: <WRITE_API_KEY>`
  - **Read endpoints** require `X-Public-Read-Key: <PUBLIC_READ_API_KEY>` (demo-friendly but controllable/rotatable)
- Rate limiting:
  - Write: default `60/min` per API key.
  - Public read: default `10/5min` per IP + key (configurable).
- Idempotency:
  - Header: `Idempotency-Key`
  - Unique: `(workspace_id, idempotency_key)`
  - If duplicate: return original event response (same ID).
- OpenAPI:
  - Checked-in `openapi.yaml` and served at `/api/traceforge/v1/openapi`.

### Web UI (Remix + Tailwind + shadcn/ui)
- Dark default with toggle.
- Landing page:
  - What it is, who it’s for, key features.
  - Architecture section (simple diagram + bullets).
  - Visible GitHub link.
- API Explorer:
  - Filters (service, actor, action, date range).
  - Events table + detail drawer.
  - “Send test event” form (write) — uses WRITE key *only locally*; in hosted demo this is disabled by default.
  - CSV export for current filtered result set.
- Public demo workspace:
  - UI defaults `workspace_id = public` and never asks users to enter it.

## Data model (week 1)
- `workspaces` (id, name)
- `api_keys` (id, workspace_id, label, key_hash, last_used_at, revoked_at)
- `events` (id, workspace_id, service, actor, action, resource, metadata jsonb, created_at)
- `idempotency_keys` (workspace_id, key, event_id, created_at)

## Milestones (day-by-day)

### Day 1 — Foundations
- Monorepo scaffold (apps/api, apps/web, infra, docs).
- Docker compose Postgres.
- Go API skeleton + health endpoint.
- Remix app skeleton + landing placeholder.
- CI (main-only) for Go test/build + web build.

### Day 2 — DB + Core ingestion
- SQL migrations for workspaces/api_keys/events/idempotency.
- POST `/events` (append-only) with validation.
- Idempotency implementation + tests.
- Rate limit middleware (write).
- `.http` files for quick manual testing.

### Day 3 — Query API + public read safety
- GET `/events` with filters + pagination.
- Public read key enforcement + strict read rate limits.
- Indexes for common queries.
- OpenAPI yaml updated and served.

### Day 4 — API Explorer UI (read)
- Explorer page: filters + table + drawer.
- Wire to GET `/events`.
- Loading + empty/error states.
- CSV export.

### Day 5 — UI polish + landing
- Real landing page content and sections.
- Dark mode toggle, animations, shadcn polish.
- GitHub repo link in header/footer.

### Day 6 — Deployment wiring (free tiers)
- Railway API deploy + Railway Postgres.
- Netlify deploy for Remix.
- Document env vars + demo limitations (no write in hosted demo unless explicitly enabled).

### Day 7 — Hardening + portfolio readiness
- Tighten README, add screenshots, add “How to demo”.
- Add ADRs and clean up.
- Final smoke tests + CI green.

## Demo limitations (documented)
- Hosted demo uses a **public read-only explorer**. Write testing is local-only by default to avoid abuse/cost.
