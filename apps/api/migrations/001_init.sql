-- Milestone 2: apply via your preferred migration tool (or psql) on Railway/local.
-- This repo keeps migrations as plain SQL for clarity.

create table if not exists workspaces (
  id text primary key,
  name text not null,
  created_at timestamptz not null default now()
);

create table if not exists api_keys (
  id uuid primary key default gen_random_uuid(),
  workspace_id text not null references workspaces(id) on delete cascade,
  label text not null,
  key_hash text not null,
  last_used_at timestamptz,
  revoked_at timestamptz,
  created_at timestamptz not null default now()
);

create table if not exists events (
  id uuid primary key default gen_random_uuid(),
  workspace_id text not null references workspaces(id) on delete cascade,
  service text not null,
  actor text not null,
  action text not null,
  resource text not null,
  metadata jsonb not null default '{}'::jsonb,
  created_at timestamptz not null default now()
);

create table if not exists idempotency_keys (
  workspace_id text not null references workspaces(id) on delete cascade,
  key text not null,
  event_id uuid not null references events(id) on delete cascade,
  created_at timestamptz not null default now(),
  primary key (workspace_id, key)
);

create index if not exists idx_events_ws_time on events (workspace_id, created_at desc);
create index if not exists idx_events_ws_service on events (workspace_id, service);
create index if not exists idx_events_ws_action on events (workspace_id, action);
create index if not exists idx_events_ws_actor on events (workspace_id, actor);
