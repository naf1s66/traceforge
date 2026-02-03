package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"traceforge/api/internal/handler"
)

const basePath = "/api/traceforge/v1"

func Mount(r chi.Router) {
	store := handler.NewEventStore()
	writeLimiter := handler.NewWriteRateLimiter()

	r.Route(basePath, func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"ok":true,"service":"traceforge","version":"v1"}`))
		})

		r.With(writeLimiter.Middleware).Post("/events", handler.HandleCreateEvent(store))

		// TODO (Milestone 2+):
		// GET /events (read, public read key + strict rate limit)
		// GET /openapi (serve checked-in openapi.yaml)
	})
}
