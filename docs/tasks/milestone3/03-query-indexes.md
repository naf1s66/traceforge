# Task: Add query indexes for events

## Summary
- Create indexes to keep common event queries fast.
- Match the expected filter fields from the API.

**Status:** Completed.

## Acceptance Criteria
- [x] Indexes exist for `(workspace_id, created_at)` and for `service`, `actor`, `action`.
- [x] Indexes are defined in `apps/api/migrations/001_init.sql`.
- [x] Index naming and coverage matches the planned query filters.

## Notes
- Revisit indexes after real query patterns emerge.
