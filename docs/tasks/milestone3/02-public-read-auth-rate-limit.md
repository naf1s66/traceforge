# Task: Enforce public read auth and rate limits

## Summary
- Protect read endpoints with a public read key.
- Apply strict rate limits for demo safety.

**Status:** New.

## Acceptance Criteria
- [ ] Read endpoints require `X-Public-Read-Key` and reject missing/invalid keys.
- [ ] Rate limiting enforces `RATE_LIMIT_PUBLIC_READ_PER_5M` per IP + key.
- [ ] Tests cover authorized reads and throttling behavior.

## Notes
- Public reads are demo-only; keep controls easy to rotate.
