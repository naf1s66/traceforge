package httpapi

import (
  "net/http"

  "github.com/go-chi/chi/v5"
)

const base = "/api/traceforge/v1"

func Mount(r chi.Router) {
  r.Route(base, func(r chi.Router) {
    r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusOK)
      w.Write([]byte(`{"ok":true,"service":"traceforge","version":"v1"}`))
    })

    // TODO (Milestone 2+):
    // POST /events (write, API key + idempotency + rate limit)
    // GET /events (read, public read key + strict rate limit)
    // GET /openapi (serve checked-in openapi.yaml)
  })
}
