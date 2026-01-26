# TraceForge

TraceForge is a small, production-minded **audit log / event ledger** service:

- **Go + PostgreSQL** ingestion & query API (append-only)
- **Idempotency keys** to prevent duplicate events
- **API-key auth** for write endpoints with strict rate limits (cost-safety)
- **Public read** mode for the hosted demo UI (read endpoints require a *public read key*)
- **Remix + Tailwind + shadcn/ui** web app:
  - Landing page (product-style)
  - Custom **API Explorer** (Swagger-like UX) with filters + event detail drawer + CSV export
  - Prominent link to the GitHub repo

## Monorepo layout

- `apps/api` — Go API (OpenAPI spec generated from a checked-in YAML for now)
- `apps/web` — Remix app (landing + explorer)
- `infra` — docker compose (Postgres)
- `docs` — PRD, agents, ADRs, and milestone task docs

## Quickstart (local)

### 1) Requirements

- Node 20+
- pnpm 8+
- Go 1.22+
- Docker Desktop

### 2) Install

```bash
pnpm install
cd apps/api && go mod download
```

### 3) Start infra

```bash
make up
```

### 4) Configure env

Copy env examples:

```bash
cp infra/env/api.env.example apps/api/.env
cp infra/env/web.env.example apps/web/.env
```

### 5) Run dev

In one terminal:

```bash
make dev
```

- API: `http://localhost:8080/api/traceforge/v1/health`
- Web: `http://localhost:3000`

## Env vars

### API (`apps/api/.env`)

- `DATABASE_URL` (required)
- `WRITE_API_KEY` (required for POST ingestion)
- `PUBLIC_READ_API_KEY` (required for read endpoints used by the public UI)
- `RATE_LIMIT_WRITE_PER_MINUTE` (default 60)
- `RATE_LIMIT_PUBLIC_READ_PER_5M` (default 10)

### Web (`apps/web/.env`)

- `API_BASE_URL` (e.g. `http://localhost:8080`)
- `PUBLIC_READ_API_KEY` (same as API)
- `GITHUB_REPO_URL` (shown in UI)

## Deployment notes (free-tier)

- **API + Postgres**: Railway (free tier)
- **Web**: Netlify
- Hosted demo uses a dedicated **public workspace** and **public read key**. In production SaaS usage, each customer would have its own workspace + keys (see ADRs).
