# Task: Implement GET /events query API

## Summary
- Add read endpoint to list events with filters and pagination.
- Require workspace scoping on every request.

**Status:** New.

## Acceptance Criteria
- [ ] `GET /api/traceforge/v1/events` supports filters: service, actor, action, from, to.
- [ ] Pagination uses `limit` + `cursor` (or similar) and returns a next cursor.
- [ ] Responses return a stable `items` array with consistent ordering.

## Notes
- Align filtering and pagination with existing indexes.
