# Task: Implement POST /events ingestion

## Summary
- Add append-only event ingestion with validation and API key auth.
- Persist events scoped to `workspace_id`.

**Status:** New.

## Acceptance Criteria
- [ ] `POST /api/traceforge/v1/events` validates required fields and payload shape.
- [ ] Write requests require `X-API-Key` and return 401/403 when invalid.
- [ ] Successful writes return the created event (201) with server timestamps.

## Notes
- Enforce append-only behavior; no update/delete endpoints.
