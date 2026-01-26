# Task: Add HTTP smoke tests for ingestion

## Summary
- Provide `.http` examples for creating and querying events.
- Include an idempotency duplicate scenario.

**Status:** New.

## Acceptance Criteria
- [ ] `apps/api/tests/requests.http` includes `POST /events` and `GET /events` examples.
- [ ] Requests show required headers (`X-API-Key`, `Idempotency-Key`, `X-Public-Read-Key`).
- [ ] Examples document expected responses for duplicate idempotent writes.

## Notes
- Keep examples copy/paste friendly for manual testing.
