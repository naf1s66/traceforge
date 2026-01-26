# Task: Wire Explorer to GET /events

## Summary
- Fetch events from the API using the public read key.
- Map query params to API filters in the loader.

**Status:** In progress - loader fetch exists, but API is not implemented.

## Acceptance Criteria
- [x] Loader calls `GET /events` with `workspace_id` and filter params.
- [x] Public read key is sent via `X-Public-Read-Key`.
- [ ] Successful responses populate the results table.

## Notes
- Default workspace to the public demo value from env.
