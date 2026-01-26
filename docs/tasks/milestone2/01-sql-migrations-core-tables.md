# Task: Create SQL migrations for core tables

## Summary
- Define Postgres tables for workspaces, api_keys, events, and idempotency_keys.
- Add indexes for common event queries.

**Status:** Completed.

## Acceptance Criteria
- [x] `apps/api/migrations/001_init.sql` creates `workspaces`, `api_keys`, `events`, and `idempotency_keys`.
- [x] `workspace_id` is enforced on all tenant-scoped records.
- [x] Indexes exist for common event queries (time, service, actor, action).

## Notes
- Keep migrations as plain SQL for Railway compatibility.
