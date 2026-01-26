# Task: Enforce write rate limits

## Summary
- Add rate limiting middleware for write endpoints.
- Use a per-key limit with sane defaults.

**Status:** New.

## Acceptance Criteria
- [ ] Write endpoints enforce `RATE_LIMIT_WRITE_PER_MINUTE` per API key.
- [ ] Limit breaches return 429 with a helpful error message.
- [ ] Tests cover allowed vs blocked traffic.

## Notes
- Default to free-tier safe limits; make limits configurable.
