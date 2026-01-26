# Task: Add Docker Compose Postgres

## Summary
- Provide a local Postgres service for the API using Docker Compose.
- Persist data in a named volume and expose port 5432.

**Status:** Completed.

## Acceptance Criteria
- [x] `infra/docker-compose.yml` defines a Postgres 16 service with credentials.
- [x] Port 5432 is mapped for local development.
- [x] Data persists in a named Docker volume.

## Notes
- Keep credentials in env examples; no secrets in commits.
