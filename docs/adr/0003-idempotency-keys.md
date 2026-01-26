# ADR 0003: Idempotency keys

## Decision
Support `Idempotency-Key` header on write ingestion.
Uniqueness: `(workspace_id, idempotency_key)`.

## Why
- Common production requirement for event ingestion.
- Prevents duplicates when clients retry.

## Behavior
- First request creates event and stores idempotency mapping.
- Duplicate key returns the original event (same event id).
