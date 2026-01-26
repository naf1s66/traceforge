# Task: Add idempotency key handling

## Summary
- Support `Idempotency-Key` on writes to prevent duplicate events.
- Return the original event response on duplicate keys.

**Status:** New.

## Acceptance Criteria
- [ ] Duplicate `(workspace_id, idempotency_key)` returns the original event without creating a new row.
- [ ] `idempotency_keys` are recorded alongside created events.
- [ ] Unit/integration tests cover duplicate submission behavior.

## Notes
- Idempotency is required for safe ingestion retries.
