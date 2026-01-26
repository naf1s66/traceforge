import { Link } from "@remix-run/react";

export default function Index() {
  const repo = process.env.GITHUB_REPO_URL ?? "https://github.com/naf1s66/traceforge";
  return (
    <main className="mx-auto max-w-5xl px-6 py-12">
      <header className="flex items-center justify-between">
        <h1 className="text-2xl font-semibold tracking-tight">TraceForge</h1>
        <a className="text-sm text-neutral-300 hover:text-white" href={repo} target="_blank" rel="noreferrer">
          GitHub ↗
        </a>
      </header>

      <section className="mt-10 rounded-2xl border border-neutral-800 bg-neutral-900/30 p-8 shadow-sm">
        <p className="text-neutral-200">
          A lightweight, production-minded <span className="font-medium">audit log / event ledger</span> service.
          Append-only events, idempotent ingestion, strict rate-limits, and a public read-only explorer.
        </p>
        <div className="mt-6 flex gap-3">
          <Link to="/explorer" className="rounded-xl bg-white px-4 py-2 text-sm font-medium text-neutral-900 hover:bg-neutral-200">
            Open API Explorer
          </Link>
          <a href={repo} target="_blank" rel="noreferrer" className="rounded-xl border border-neutral-700 px-4 py-2 text-sm text-neutral-200 hover:border-neutral-500">
            View source
          </a>
        </div>
      </section>

      <section className="mt-10 grid gap-4 sm:grid-cols-3">
        {[
          { title: "Append-only", body: "Events are immutable. No update/delete endpoints." },
          { title: "Idempotent writes", body: "Idempotency keys prevent duplicates on retries." },
          { title: "Free-tier safe", body: "Strict rate limits and rotatable public read key." },
        ].map((c) => (
          <div key={c.title} className="rounded-2xl border border-neutral-800 bg-neutral-900/20 p-5">
            <div className="text-sm font-semibold">{c.title}</div>
            <div className="mt-2 text-sm text-neutral-300">{c.body}</div>
          </div>
        ))}
      </section>

      <footer className="mt-14 text-xs text-neutral-500">
        Sleek, minimal, modern (dark default). See docs in <code>docs/</code>.
      </footer>
    </main>
  );
}
