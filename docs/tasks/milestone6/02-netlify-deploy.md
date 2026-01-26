# Task: Configure Netlify deployment for Remix

## Summary
- Set up Netlify build settings for the Remix app.
- Document the hosted URL and environment variables.

**Status:** In progress - netlify.toml exists.

## Acceptance Criteria
- [x] Netlify build uses `pnpm -C apps/web build` and publishes `apps/web/public`.
- [ ] SSR is wired via the Remix serverless function.
- [ ] README includes the live Netlify URL and env var list.

## Notes
- Ensure demo uses the public read key only.
