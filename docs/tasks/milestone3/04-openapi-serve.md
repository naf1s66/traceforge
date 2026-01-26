# Task: Update and serve OpenAPI spec

## Summary
- Expand `openapi.yaml` to document the events API.
- Serve the spec from the API for the Explorer UI.

**Status:** In progress - openapi.yaml only includes /health today.

## Acceptance Criteria
- [ ] `openapi.yaml` includes `POST /events` and `GET /events` schemas.
- [ ] API serves the spec at `/api/traceforge/v1/openapi`.
- [ ] The spec stays in sync with request/response payloads.

## Notes
- Update OpenAPI any time endpoints change.
