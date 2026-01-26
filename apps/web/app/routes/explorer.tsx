import type { LoaderFunctionArgs } from "@remix-run/node";
import { json } from "@remix-run/node";
import { Form, useLoaderData } from "@remix-run/react";

type EventRow = {
  id: string;
  service: string;
  actor: string;
  action: string;
  resource: string;
  created_at: string;
};

export async function loader({ request }: LoaderFunctionArgs) {
  const url = new URL(request.url);
  const q = url.searchParams;

  const base = process.env.API_BASE_URL ?? "http://localhost:8080";
  const publicKey = process.env.PUBLIC_READ_API_KEY ?? "dev-public-read-key";
  const workspaceId = process.env.PUBLIC_WORKSPACE_ID ?? "public";

  const params = new URLSearchParams();
  params.set("workspace_id", workspaceId);
  for (const k of ["service", "actor", "action", "from", "to", "limit", "cursor"]) {
    const v = q.get(k);
    if (v) params.set(k, v);
  }

  // NOTE: endpoint implemented in Milestone 3. Until then, show placeholder UI.
  let rows: EventRow[] = [];
  let notReady = false;
  try {
    const res = await fetch(`${base}/api/traceforge/v1/events?${params.toString()}`, {
      headers: { "X-Public-Read-Key": publicKey },
    });
    if (res.status === 501 || res.status === 404) notReady = true;
    else if (res.ok) rows = (await res.json()).items ?? [];
    else notReady = true;
  } catch {
    notReady = true;
  }

  return json({ rows, notReady });
}

export default function Explorer() {
  const { rows, notReady } = useLoaderData<typeof loader>();
  const repo = process.env.GITHUB_REPO_URL ?? "https://github.com/naf1s66/traceforge";

  return (
    <main className="mx-auto max-w-6xl px-6 py-10">
      <header className="flex items-center justify-between">
        <div>
          <h1 className="text-xl font-semibold tracking-tight">API Explorer</h1>
          <p className="mt-1 text-sm text-neutral-400">
            Public read-only explorer (workspace is fixed for demo). Filters behave like a slim Swagger UI.
          </p>
        </div>
        <a className="text-sm text-neutral-300 hover:text-white" href={repo} target="_blank" rel="noreferrer">
          GitHub ↗
        </a>
      </header>

      <section className="mt-6 rounded-2xl border border-neutral-800 bg-neutral-900/20 p-6">
        <Form method="get" className="grid gap-3 md:grid-cols-6">
          <input name="service" placeholder="service" className="rounded-xl border border-neutral-800 bg-neutral-950 px-3 py-2 text-sm" />
          <input name="actor" placeholder="actor" className="rounded-xl border border-neutral-800 bg-neutral-950 px-3 py-2 text-sm" />
          <input name="action" placeholder="action" className="rounded-xl border border-neutral-800 bg-neutral-950 px-3 py-2 text-sm" />
          <input name="from" placeholder="from (ISO)" className="rounded-xl border border-neutral-800 bg-neutral-950 px-3 py-2 text-sm md:col-span-1" />
          <input name="to" placeholder="to (ISO)" className="rounded-xl border border-neutral-800 bg-neutral-950 px-3 py-2 text-sm md:col-span-1" />
          <button className="rounded-xl bg-white px-4 py-2 text-sm font-medium text-neutral-900 hover:bg-neutral-200">
            Query
          </button>
        </Form>

        {notReady ? (
          <div className="mt-5 rounded-xl border border-amber-900/40 bg-amber-950/20 p-4 text-sm text-amber-200">
            Explorer backend endpoint isn’t ready yet (Milestone 3). UI is scaffolded and will light up once
            <code className="mx-1">GET /events</code> is implemented.
          </div>
        ) : null}

        <div className="mt-6 overflow-hidden rounded-2xl border border-neutral-800">
          <table className="w-full text-left text-sm">
            <thead className="bg-neutral-900/40 text-neutral-300">
              <tr>
                <th className="px-4 py-3">time</th>
                <th className="px-4 py-3">service</th>
                <th className="px-4 py-3">actor</th>
                <th className="px-4 py-3">action</th>
                <th className="px-4 py-3">resource</th>
              </tr>
            </thead>
            <tbody>
              {rows.length === 0 ? (
                <tr>
                  <td className="px-4 py-4 text-neutral-500" colSpan={5}>No results.</td>
                </tr>
              ) : rows.map((r) => (
                <tr key={r.id} className="border-t border-neutral-800 hover:bg-neutral-900/20">
                  <td className="px-4 py-3 text-neutral-400">{new Date(r.created_at).toLocaleString()}</td>
                  <td className="px-4 py-3">{r.service}</td>
                  <td className="px-4 py-3">{r.actor}</td>
                  <td className="px-4 py-3">{r.action}</td>
                  <td className="px-4 py-3 text-neutral-300">{r.resource}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </section>
    </main>
  );
}
