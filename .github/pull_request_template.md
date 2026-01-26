# Title
<!-- e.g., milestone-1: monorepo scaffold + API + web foundations -->

## Summary
Explain the purpose of this PR and the outcome in 2-4 sentences.

## What's Included
- [ ] Monorepo / workspace setup
- [ ] Docker Compose (Postgres)
- [ ] Go API health endpoint + OpenAPI
- [ ] Remix + Tailwind UI
- [ ] Docs updated (README, PRD, ADRs, tasks)

## How to Test
1. **Infra**
   ```bash
   make up
   docker ps
   ```
2. **API**
   ```bash
   make dev
   # in another terminal
   curl http://localhost:8080/api/traceforge/v1/health
   ```
3. **Web**
   ```bash
   pnpm -C apps/web dev
   # open http://localhost:3000
   ```

## Screenshots / Logs
<!-- Landing page, explorer UI, curl health output, docker ps -->

## Checklist
- [ ] Lints pass (`make lint` or `pnpm -r run lint`)
- [ ] Tests pass (`make test` or `go test ./...`)
- [ ] Updated docs where needed
- [ ] No secrets committed

## Notes / Follow-ups
<!-- TODOs for next milestone, known limitations, decisions -->

---

If anything in the output looks noisy after this, paste the exact error lines and I'll give you the fix line-by-line.
