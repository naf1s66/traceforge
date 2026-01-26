# ADR 0002: Multi-tenant workspaces

## Decision
Introduce `workspace_id` as a first-class scope for all data and API access.

## Why
- Signals SaaS/system design maturity to recruiters.
- Enables per-workspace API keys, quotas, and future org features.
- Prevents a costly schema redesign later.

## Demo UX
The public UI uses a dedicated `PUBLIC_WORKSPACE_ID=public` and does not ask visitors to supply workspace IDs.
