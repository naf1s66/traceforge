# ADR 0005: Public read key strategy

## Decision
Read endpoints require a rotatable `PUBLIC_READ_API_KEY` (header `X-Public-Read-Key`).

## Why
- UI can be public without login.
- Key can be rotated if leaked.
- Rate limits can be tied to the key + IP.

## Notes
Hosted demo uses read-only data and does not expose write keys.
