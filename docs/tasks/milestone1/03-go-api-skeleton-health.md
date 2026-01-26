# Task: Build Go API skeleton and health check

## Summary
- Create the Go API server with routing, middleware, and a versioned base path.
- Add a health endpoint for uptime and CI checks.

**Status:** Completed.

## Acceptance Criteria
- [x] `cmd/server` boots the API and mounts routes under `/api/traceforge/v1`.
- [x] `GET /api/traceforge/v1/health` returns a JSON OK payload.
- [x] A basic unit test asserts the health response.

## Notes
- Keep handlers small and prepare for auth, rate limiting, and DB wiring.
