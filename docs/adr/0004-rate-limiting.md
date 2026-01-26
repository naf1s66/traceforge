# ADR 0004: Rate limiting

## Decision
Apply strict rate limiting to protect free-tier deployments.

- Write: default 60/min per API key
- Public read: default 10/5min per IP + key

## Why
- Limits abuse/cost on free tiers.
- Shows production thinking.
